package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCourseByClass(c *gin.Context) {
	fmt.Println("getCourseByClass()")
}

func GetRankByClass(c *gin.Context) {
	fmt.Println("getRankByClass()")
}

func GetClassList(c *gin.Context) {
	fmt.Println("getClassList()")
}
