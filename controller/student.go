package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCourseByStudent(c *gin.Context) {
	fmt.Println("getCourseByStudent()")
}

func GetStudentList(c *gin.Context) {
	fmt.Println("getStudentList()")
}
