package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/taaraora/hh-stub/pkg/hh"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9876))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	hh.RegisterSessionForwarderServer(grpcServer, &dummyServer{})
	grpcServer.Serve(lis)
}

type dummyServer struct{}

func (d *dummyServer) ReportSessionsStats(ctx context.Context, request *hh.ReportSessionsStatsRequest) (*emptypb.Empty, error) {
	//TODO implement me
	b, _ := json.Marshal(request)
	fmt.Println(string(b))

	return nil, nil
}
