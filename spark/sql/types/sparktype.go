// Package types @author: Violet-Eva @date  : 2025/8/28 @notes :
package types

import (
	"reflect"
	"strings"
)

type SparkTag struct {
	Column string
	Type   string
	Format string
}

// 解析spark标签内容
func parseSparkTag(tag string) SparkTag {
	var st SparkTag
	if tag == "" {
		return st
	}

	parts := strings.Split(tag, ";")
	for _, part := range parts {
		kv := strings.SplitN(part, ":", 2)
		if len(kv) != 2 {
			continue
		}

		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "column":
			st.Column = value
		case "type":
			st.Type = value
		case "format":
			st.Format = value
		}
	}
	return st
}

func ExtractSparkTags(structType interface{}) map[string]SparkTag {
	result := make(map[string]SparkTag)
	t := reflect.TypeOf(structType)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("spark")
		if tag != "" {
			result[field.Name] = parseSparkTag(tag)
		}
	}

	return result
}
