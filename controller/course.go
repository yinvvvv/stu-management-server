package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

// /course

func GetCourseList(c *gin.Context) {
	fmt.Println("getCourseList()")
	sql := "Select * from zjw_course"
	rows, _ := util.Db.Query(sql)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}

func GetCourseRank(c *gin.Context) {
	fmt.Println("getCourseRank()")
	id := c.Param("id")
	sql := `
		select * from zjw_course_rank
		where course_id = ?;
	`
	stmt, _ := util.Db.Prepare(sql)
	defer stmt.Close()
	rows, _ := stmt.Query(id)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
