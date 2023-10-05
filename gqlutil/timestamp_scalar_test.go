package gqlutil_test

import (
	"testing"

	"github.com/clavinjune/gokit/gqlutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	millis  int64 = 1696469420103
	seconds int64 = 1696469420
	nanos   int32 = 103000000
)

func TestTimestampScalarSerialize(t *testing.T) {
	r := require.New(t)

	x := gqlutil.TimestampScalarSerialize(timestamppb.Timestamp{
		Seconds: seconds,
		Nanos:   nanos,
	})
	n, ok := x.(int64)
	r.True(ok)
	r.Equal(millis, n)
}

func TestTimestampScalarParseValue(t *testing.T) {
	r := require.New(t)

	x := gqlutil.TimestampScalarParseValue(1696469420103)
	pb, ok := x.(*timestamppb.Timestamp)
	r.True(ok)
	r.Equal(seconds, pb.GetSeconds())
	r.Equal(nanos, pb.GetNanos())
}
