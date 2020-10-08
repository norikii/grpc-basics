package serializer_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tatrasoft/grpc-basics/pb"
	"github.com/tatrasoft/grpc-basics/sample"
	"github.com/tatrasoft/grpc-basics/serializer"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobugFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	assert.True(t, proto.Equal(laptop1, laptop2))
}

