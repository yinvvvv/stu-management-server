package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"stu-management-server/controller"
	"stu-management-server/middleware"
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
	router.Use(sessions.Sessions("SESSIONID", cookie.NewStore([]byte("secret"))))

	router.POST("/api/login", controller.Login)

	r := router.Group("/api")
	{

		r.Use(middleware.AuthCheck)

		// major
		major := r.Group("/major")
		{
			major.GET("/list", controller.GetMajorList)
		}

		// class
		class := r.Group("/class")
		{
			class.GET("/course", controller.GetCourseByClass)
			class.GET("/rank", controller.GetRankByClass)
			class.GET("/list", controller.GetClassList)
		}

		// score
		score := r.Group("/score")
		{
			score.GET("/course", controller.GetAvgScoreByCourse)
			score.POST("/add", controller.AddScore)
			score.GET("/student", controller.GetStudentScore)
		}

		// course
		course := r.Group("/course")
		{
			course.GET("/list", controller.GetCourseList)
			course.GET("/rank", controller.GetCourseRank)
			course.GET("/avg", controller.GetCourseAvg)
		}

		// student
		student := r.Group("/student")
		{
			student.GET("/class/course", controller.GetStudentClassCourse)
			student.GET("/course", controller.GetCourseByStudent)
			student.GET("/list", controller.GetStudentList)
			student.GET("/score", controller.GetStudentScoreByYear)
		}

		// teacher
		teacher := r.Group("/teacher")
		{
			teacher.GET("/list", controller.GetTeacherList)
			teacher.GET("/course", controller.GetTeacherCourse)
		}
	}

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
