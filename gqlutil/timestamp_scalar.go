package gqlutil

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	TimestampScalar = graphql.NewScalar(graphql.ScalarConfig{
		Name:         "ProtobufTimestamp",
		Description:  "protobuf Timestamp scalar type, serialize Timestamp as Unix millis",
		Serialize:    TimestampScalarSerialize,
		ParseValue:   TimestampScalarParseValue,
		ParseLiteral: graphql.Boolean.ParseLiteral,
	})
)

func TimestampScalarSerialize(v any) any {
	switch t := v.(type) {
	case timestamppb.Timestamp:
		return t.AsTime().UnixMilli()
	case *timestamppb.Timestamp:
		return t.AsTime().UnixMilli()
	}

	return nil
}

func TimestampScalarParseValue(v any) any {
	switch t := v.(type) {
	case int:
		return timestamppb.New(
			time.UnixMilli(int64(t)),
		)
	case *int:
		return timestamppb.New(
			time.UnixMilli(int64(*t)),
		)
	}

	return nil
}

func TimestampScalarParseLiteral(v ast.Value) any {
	return timestamppb.New(
		time.UnixMilli(
			int64(
				graphql.Int.ParseLiteral(v).(int),
			),
		),
	)
}
