package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) GetCourse(ctx context.Context, in *GetCourseRequest) (*GetCourseReply, error) {
	return &GetCourseReply{Message: "*List of all courses*"}, nil
}

func main() {
	list, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("Failed to listen on port 8008: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(list); err != nil {

	}
}
