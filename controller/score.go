package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAvgScoreByCourse(c *gin.Context) {
	fmt.Println("getAvgScoreByCourse()")
}

func AddScore(c *gin.Context) {
	fmt.Println("addScore()")
}

func GetStudentScore(c *gin.Context) {
	fmt.Println("getStudentScore()")
}

func GetStudentScoreByYear(c *gin.Context) {
	fmt.Println("getStudentScoreByYear()")
}
