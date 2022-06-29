package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/controller"
	"stu-management-server/util"
)

func main() {

	dbAddress := "lambo77.world:3306"
	dbName := "202001_zjw_mis"
	dbUsername := "lambo77"
	dbPassword := "rootroot"

	util.InitDB(fmt.Sprintf("%s:%s@(%s)/%s", dbUsername, dbPassword, dbAddress, dbName))
	defer util.Db.Close()

	router := gin.Default()

	// major
	major := router.Group("/major")
	{
		major.GET("/list", controller.GetMajorList)
	}

	// class
	class := router.Group("/class")
	{
		class.GET("/course/:id", controller.GetCourseByClass)
		class.GET("/rank/:id", controller.GetRankByClass)
		class.GET("/list/:major_id", controller.GetClassList)
	}

	// score
	score := router.Group("/score")
	{
		score.GET("/course/:course_id", controller.GetAvgScoreByCourse)
		score.POST("/add", controller.AddScore)
		score.GET("/student/:student_id", controller.GetStudentScore)
	}

	// course
	course := router.Group("/course")
	{
		course.GET("/list", controller.GetCourseList)
		course.GET("/rank/:id", controller.GetCourseRank)
	}

	// student
	student := router.Group("/student")
	{
		student.GET("/course/:id", controller.GetCourseByStudent)
		student.GET("/list/:class_id", controller.GetStudentList)
		student.GET("/score/:id/:year", controller.GetStudentScoreByYear)
	}

	// teacher
	teacher := router.Group("/teacher")
	{
		teacher.GET("/list", controller.GetTeacherList)
		teacher.GET("/course/:id", controller.GetTeacherCourse)
	}

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
