package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

// /student

func GetCourseByStudent(c *gin.Context) {
	fmt.Println("getCourseByStudent()")
	id := c.Param("id")
	sql := `
		
	`
	stmt, _ := util.Db.Prepare(sql)
	rows, _ := stmt.Query(id)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}

func GetStudentList(c *gin.Context) {
	fmt.Println("getStudentList()")
	classId := c.Param("class_id")
	sql := `
		select zjw_student.*
		from zjw_student, zjw_class
		where zjw_student.class_id = zjw_class.id
		and zjw_class.id = ?;
	`
	stmt, _ := util.Db.Prepare(sql)
	rows, _ := stmt.Query(classId)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}

func GetStudentScoreByYear(c *gin.Context) {
	fmt.Println("getStudentScoreByYear()")
	id := c.Param("id")
	year := c.Param("year")
	sql := `
		select zjw_course.title, zjw_score.score from zjw_score, zjw_course, zjw_class
		where zjw_score.course_id = zjw_course.id
		and zjw_course.class_id = zjw_class.id
		and zjw_class.year + zjw_course.year = ? + 1
		and zjw_score.student_id = ?;
	`
	stmt, _ := util.Db.Prepare(sql)
	defer stmt.Close()
	rows, _ := stmt.Query(year, id)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
