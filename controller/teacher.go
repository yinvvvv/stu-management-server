package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

func GetTeacherList(c *gin.Context) {
	fmt.Println("getTeacherList()")
	sql := "select * from zjw_teacher;"
	rows, _ := util.Db.Query(sql)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}

func GetTeacherCourse(c *gin.Context) {
	fmt.Println("getTeacherCourse()")
	id := c.Param("id")
	sql := `
		select zjw_course.*
		from zjw_course, zjw_teacher, zjw_teach
		where zjw_course.id = zjw_teach.course_id
		and zjw_teacher.id = zjw_teach.teacher_id
		and zjw_teacher.id = ?;
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
