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
