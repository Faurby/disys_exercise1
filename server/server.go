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

// Since we passed the Server struct along earlier, we can now call these methods from "within" it. As far as I understand, we "overwrite" /
// build on the already established GetCourse function from grpc.pg.go (maybe... not sure)
// Basically its very importan how this method signature is written and how the reutrn function is formatted.
func (s *Server) GetCourse(ctx context.Context, in *course.GetCourseRequest) (*course.GetCourseReply, error) {
	fmt.Println("Received GetCourse request")
	return &course.GetCourseReply{Message: "*List of all courses*"}, nil
}

func (s *Server) PostCourse(ctx context.Context, in *course.PostCourseRequest) (*course.PostCourseReply, error) {
	fmt.Printf("Received PostCourses request with course name: %s", in.Name)
	// Add courses to list / database of courses
	return &course.PostCourseReply{Message: "Successfully posted new course"}, nil
}

func main() {
	// Create listener tcp on port 8008
	list, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("Failed to listen on port 8008: %v", err)
	}

	// Create a new grpc server.
	// Register our CourseService on our grpcServer, and passing our Server{} struct with
	// Needs the struct for implementing the standard functions on it(?)
	grpcServer := grpc.NewServer()
	course.RegisterCourseServiceServer(grpcServer, &Server{})

	// Accept incoming connections on the listener.
	//  RegisterCourseServiceServer MUST be called before Serve()
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
