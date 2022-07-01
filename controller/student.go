package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

// /student

func GetStudentClassCourse(c *gin.Context) {
	fmt.Println("getStudentClassCourse()")
	studentId := c.Query("student_id")
	sql := `
		SELECT zjw_course.* FROM zjw_student, zjw_class, zjw_course
		WHERE zjw_student.class_id = zjw_class.id
		AND zjw_course.class_id = zjw_class.id
		AND zjw_student.id = ?;
	`
	stmt, _ := util.Db.Prepare(sql)
	defer stmt.Close()
	rows, _ := stmt.Query(studentId)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}

func GetCourseByStudent(c *gin.Context) {
	fmt.Println("getCourseByStudent()")
	id := c.Query("id")
	sql := `
		select zjw_course.* from zjw_score, zjw_course
		where zjw_score.course_id = zjw_course.id
		and zjw_score.student_id = ?;
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
	classId := c.Query("class_id")
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
	id := c.Query("id")
	year := c.Query("year")
	sql := `
		select * from zjw_student_score_by_year
		where class_year + course_year = ? + 1
		and student_id = ?;
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

func GetStudentGPAByYear(c *gin.Context) {
	fmt.Println("getStudentGPAByYear()")
	id := c.Query("id")
	year := c.Query("year")
	sql := `
		select (SUM(score) / COUNT(*)) / 20 AS GPA from zjw_student_score_by_year
		where class_year + course_year = ? + 1
		and student_id = ?;
	`
	stmt, _ := util.Db.Prepare(sql)
	defer stmt.Close()
	rows, _ := stmt.Query(year, id)
	defer rows.Close()
	result := util.QueryAndParse(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
