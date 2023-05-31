package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "golang.cafe/protobuf/model"
	"google.golang.org/grpc"
)

// implement the RouteGuideServer interface
type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	return &pb.Feature{Name: "test", Location: point}, nil
}

func (s *routeGuideServer) Check(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	fmt.Println("Health Check")
	return &pb.Empty{}, nil
}

func main() {

	finish := make(chan bool)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := routeGuideServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
		}
	}()

	//----------------------------------------
	gwmux := runtime.NewServeMux()

	err = pb.RegisterRouteGuideHandlerServer(context.Background(), gwmux, &routeGuideServer{})

	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	go func() {
		log.Fatalln(http.ListenAndServe(":80", gwmux))
	}()

	<-finish
}
