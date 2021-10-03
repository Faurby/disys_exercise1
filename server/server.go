package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"disys_exercise1/course"

	"google.golang.org/grpc"
)

type Server struct {
	course.UnimplementedCourseServiceServer
}

func (s *Server) GetCourse(ctx context.Context, in *course.GetCourseRequest) (*course.GetCourseReply, error) {
	fmt.Println("Recieved GetCourse request")
	return &course.GetCourseReply{Message: "*List of all courses*"}, nil
}

func (s *Server) PostCourse(ctx context.Context, in *course.PostCourseRequest) (*course.PostCourseReply, error) {
	fmt.Printf("Recieved PostCourses request with course name: %s", in.Name)
	// Add courses to list / database of courses
	return &course.PostCourseReply{Message: "Successfully posted new course"}, nil
}

func main() {
	// Create listener tcp on port 8008
	list, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("Failed to listen on port 8008: %v", err)
	}

	// Create a new grpc server with the generated server from the proto
	grpcServer := grpc.NewServer()
	course.RegisterCourseServiceServer(grpcServer, &Server{})

	// Accept incoming connections on the listener
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
