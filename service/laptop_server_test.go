package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tatrasoft/grpc-basics/pb"
	"github.com/tatrasoft/grpc-basics/sample"
	"github.com/tatrasoft/grpc-basics/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "invalid_ID"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicatedID := service.NewInMemoryLaptopStore()
	err := storeDuplicatedID.Save(laptopDuplicateID)
	require.Nil(t, err)

	testCases := []struct{
		name string
		laptop *pb.Laptop
		store service.LaptopStore
		errCode codes.Code
	}{
		{
			name: "success_with_id",
			laptop : sample.NewLaptop(),
			store: service.NewInMemoryLaptopStore(),
			errCode: codes.OK,
		},
		{
			name: "success_with_no_id",
			laptop : laptopNoID,
			store: service.NewInMemoryLaptopStore(),
			errCode: codes.OK,
		},
		{
			name: "failure_invalid_id",
			laptop : laptopInvalidId,
			store: service.NewInMemoryLaptopStore(),
			errCode: codes.InvalidArgument,
		},
		{
			name: "failure_duplicated_id",
			laptop : laptopDuplicateID,
			store: storeDuplicatedID,
			errCode: codes.AlreadyExists,
		},
	}
	
	for i := range testCases {
		tc := testCases[i]
		
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{Laptop:tc.laptop}

			server := service.NewLaptopServer(tc.store, nil, nil)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.errCode == codes.OK {
				assert.NoError(t, err)
				assert.NotNil(t, res)
				assert.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					assert.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.errCode, st.Code())
			}
		})
	}
}
	

