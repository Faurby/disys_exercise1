package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.POST("/courses", postCourse)
	router.GET("/courses/:id", getCourseID)
	router.GET("/courses/:id/students", getStudentsFromCourse)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getCourses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"List": "List of all courses!",
	})
}

func getCourseID(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"courseID": id,
	})
}

func postCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Course": "POST Success",
	})
}

func getStudentsFromCourse(c *gin.Context) {
	courseName := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"CourseName":     courseName,
		"ListOfStudents": "list",
	})
}
