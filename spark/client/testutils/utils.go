package testutils

import (
	"context"
	"testing"

	proto "github.com/violet-eva-01/spark-connect/internal/generated"
	"google.golang.org/grpc"
)

// connectServiceClient is a mock implementation of the SparkConnectServiceClient interface.
type connectServiceClient struct {
	t *testing.T

	analysePlanResponse *proto.AnalyzePlanResponse
	executePlanClient   proto.SparkConnectService_ExecutePlanClient

	err error
}

func (c *connectServiceClient) ExecutePlan(ctx context.Context, in *proto.ExecutePlanRequest,
	opts ...grpc.CallOption,
) (proto.SparkConnectService_ExecutePlanClient, error) {
	return c.executePlanClient, c.err
}

func (c *connectServiceClient) AnalyzePlan(ctx context.Context, in *proto.AnalyzePlanRequest,
	opts ...grpc.CallOption,
) (*proto.AnalyzePlanResponse, error) {
	return c.analysePlanResponse, c.err
}

func (c *connectServiceClient) Config(ctx context.Context, in *proto.ConfigRequest, opts ...grpc.CallOption) (*proto.ConfigResponse, error) {
	return nil, c.err
}

func (c *connectServiceClient) AddArtifacts(ctx context.Context, opts ...grpc.CallOption) (proto.SparkConnectService_AddArtifactsClient, error) {
	return nil, c.err
}

func (c *connectServiceClient) ArtifactStatus(ctx context.Context,
	in *proto.ArtifactStatusesRequest, opts ...grpc.CallOption,
) (*proto.ArtifactStatusesResponse, error) {
	return nil, c.err
}

func (c *connectServiceClient) Interrupt(ctx context.Context, in *proto.InterruptRequest,
	opts ...grpc.CallOption,
) (*proto.InterruptResponse, error) {
	return nil, c.err
}

func (c *connectServiceClient) ReattachExecute(ctx context.Context,
	in *proto.ReattachExecuteRequest, opts ...grpc.CallOption,
) (proto.SparkConnectService_ReattachExecuteClient, error) {
	return c.executePlanClient, c.err
}

func (c *connectServiceClient) ReleaseExecute(ctx context.Context, in *proto.ReleaseExecuteRequest,
	opts ...grpc.CallOption,
) (*proto.ReleaseExecuteResponse, error) {
	return nil, c.err
}

func NewConnectServiceClientMock(epc proto.SparkConnectService_ExecutePlanClient,
	apr *proto.AnalyzePlanResponse, err error, t *testing.T,
) proto.SparkConnectServiceClient {
	return &connectServiceClient{
		t:                   t,
		analysePlanResponse: apr,
		executePlanClient:   epc,
		err:                 err,
	}
}
