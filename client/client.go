package main

import (
	"context"
	"disys_exercise1/course"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// Creat a virtual RPC Client Connection on port  8008 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := course.NewCourseServiceClient(conn)

	SendGetCourseRequest(c)
	SendPostCourseRequst(c)

}

func SendGetCourseRequest(c course.CourseServiceClient) {
	// This is tricky. Since I only want to get all courses, I dont need to send any data.
	// So im sending an empty GetCourseRequest struct (look in .proto for more info)
	message := course.GetCourseRequest{}

	//This is the responses from the GetCourse method in server.go.
	// Here I sending my request, and getting response back. Error if something bad happens
	response, err := c.GetCourse(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling GetCourse: %s", err)
	}

	// Prints the response
	fmt.Printf("Response from the Server: %s \n", response.Message)
}

func SendPostCourseRequst(c course.CourseServiceClient) {
	// Now I wanna post something to the server, so I cant send empty request.
	message := course.PostCourseRequest{
		Name: "Algorithms and Data Structures",
		Ects: 7.5,
		Teachers: []*course.Teacher{{Id: 99, Name: "Thore Husfeldt", Score: 100},
			{Id: 5, Name: "Troels Bjerre Lund", Score: 55}},
		Rating: 100,
	}

	// Same as above
	response, err := c.PostCourse(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling PostCourse: %s", err)
	}

	// Same as above
	fmt.Printf("Response from  the Server: %s \n", response.Message)
}
