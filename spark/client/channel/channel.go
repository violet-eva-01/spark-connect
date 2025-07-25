package channel

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/violet-eva-01/spark-connect/spark"

	"github.com/google/uuid"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/violet-eva-01/spark-connect/spark/sparkerrors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

// Builder is the interface that is used to implement different patterns that
// create the GRPC connection.
//
// This allows other consumers to plugin custom authentication and authorization
// handlers without having to extend directly the Spark Connect code.
type Builder interface {
	// Build creates the grpc.ClientConn according to the configuration of the builder.
	// Implementations are free to provide additional paramters in their implementation
	// and simply must satisfy this minimal set of requirements.
	Build(ctx context.Context) (*grpc.ClientConn, error)
	// User identifies the username passed as part of the Spark Connect requests.
	User() string
	// Headers refers to the request metadata that is passed for every request from the
	// client to the server.
	Headers() map[string]string
	// SessionId identifies the client side session identifier. This value must be a UUID formatted
	// as a string.
	SessionId() string
	// UserAgent identifies the user agent string that is passed as part of the request. It contains
	// information about the operating system, Go version etc.
	UserAgent() string
}

// BaseBuilder is used to parse the different parameters of the connection
// string according to the specification documented here:
//
//	https://github.com/apache/spark/blob/master/connector/connect/docs/client-connection-string.md
type BaseBuilder struct {
	host      string
	port      int
	token     string
	user      string
	headers   map[string]string
	sessionId string
	userAgent string
	clientIP  string
	conn      *grpc.ClientConn
}

func (cb *BaseBuilder) Host() string {
	return cb.host
}

func (cb *BaseBuilder) Port() int {
	return cb.port
}

func (cb *BaseBuilder) Token() string {
	return cb.token
}

func (cb *BaseBuilder) User() string {
	return cb.user
}

func (cb *BaseBuilder) Headers() map[string]string {
	return cb.headers
}

func (cb *BaseBuilder) SessionId() string {
	return cb.sessionId
}

func (cb *BaseBuilder) UserAgent() string {
	return cb.userAgent
}

// Build finalizes the creation of the gprc.ClientConn by creating a GRPC channel
// with the necessary options extracted from the connection string. For
// TLS connections, this function will load the system certificates.
func (cb *BaseBuilder) Build(ctx context.Context) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithAuthority(cb.host))
	if cb.token == "" {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		// Note: On the Windows platform, use of x509.SystemCertPool() requires
		// go version 1.18 or higher.
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
		ts := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: cb.token,
			TokenType:   "bearer",
		})
		opts = append(opts, grpc.WithPerRPCCredentials(oauth.TokenSource{TokenSource: ts}))
	}

	remote := fmt.Sprintf("%v:%v", cb.host, cb.port)
	conn, err := grpc.NewClient(remote, opts...)
	if err != nil {
		return nil, sparkerrors.WithType(fmt.Errorf("failed to connect to remote %s: %w",
			remote, err), sparkerrors.ConnectionError)
	}
	return conn, nil
}

// NewBuilder creates a new instance of the BaseBuilder. This constructor effectively
// parses the connection string and extracts the relevant parameters directly.
//
// The following parameters to the connection string are reserved: user_id, session_id, use_ssl,
// and token. These parameters are not allowed to be injected as headers.
func NewBuilder(connection string) (*BaseBuilder, error) {
	u, err := url.Parse(connection)
	if err != nil {
		return nil, err
	}

	if u.Hostname() == "" {
		return nil, sparkerrors.WithType(errors.New("URL must contain a hostname"), sparkerrors.InvalidInputError)
	}

	if u.Scheme != "sc" {
		return nil, sparkerrors.WithType(errors.New("URL schema must be set to `sc`"), sparkerrors.InvalidInputError)
	}

	port := 15002
	host := u.Host
	// Check if the host part of the URL contains a port and extract.
	if strings.Contains(u.Host, ":") {
		// We can ignore the error here already since the url parsing
		// raises the error about invalid port.
		hostStr, portStr, _ := net.SplitHostPort(u.Host)
		host = hostStr
		if len(portStr) != 0 {
			port, err = strconv.Atoi(portStr)
			if err != nil {
				return nil, err
			}
		}
	}

	// Validate that the URL path is empty or follows the right format.
	if u.Path != "" && !strings.HasPrefix(u.Path, "/;") {
		return nil, sparkerrors.WithType(
			fmt.Errorf("the URL path (%v) must be empty or have a proper parameter syntax", u.Path),
			sparkerrors.InvalidInputError)
	}

	cb := &BaseBuilder{
		host:      host,
		port:      port,
		headers:   map[string]string{},
		sessionId: uuid.NewString(),
		userAgent: "",
	}

	elements := strings.Split(u.Path, ";")
	for _, e := range elements {
		props := strings.Split(e, "=")
		if len(props) == 2 {
			if props[0] == "token" {
				cb.token = props[1]
			} else if props[0] == "user_id" {
				cb.user = props[1]
			} else if props[0] == "session_id" {
				cb.sessionId = props[1]
			} else if props[0] == "user_agent" {
				cb.userAgent = props[1]
			} else {
				cb.headers[props[0]] = props[1]
			}
		}
	}

	// Set default user ID if not set.
	if cb.user == "" {
		cb.user = os.Getenv("USER")
		if cb.user == "" {
			cb.user = "na"
		}
	}

	if cb.clientIP == "" {
		if cb.clientIP = getClientIp(); cb.clientIP == "" {
			cb.clientIP = "127.0.0.1"
		}
	}

	// Update the user agent if it is not set or set to a custom value.
	val := os.Getenv("SPARK_CONNECT_USER_AGENT")
	if cb.userAgent == "" && val != "" {
		cb.userAgent = os.Getenv("SPARK_CONNECT_USER_AGENT")
	} else if cb.userAgent == "" {
		cb.userAgent = "_SPARK_CONNECT_GO"
	}

	// In addition, to the specified user agent, we need to append information about the
	// host encoded as user agent components.
	cb.userAgent = fmt.Sprintf("%s spark/%s os/%s go/%s", cb.userAgent, spark.Version(), runtime.GOOS, runtime.Version())

	return cb, nil
}

func getClientIp() string {
	dial, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	localAddr := dial.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
