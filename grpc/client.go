package grpc

import (
	"context"
	"fmt"
	"github.com/JobNing/framework/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(ctx context.Context, toService string) (*grpc.ClientConn, error) {
	conn, err := consul.AgentHealthService(ctx, toService)
	if err != nil {
		return nil, err
	}
	fmt.Println(conn)
	return grpc.Dial(conn, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
