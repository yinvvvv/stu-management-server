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
}
