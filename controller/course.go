package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

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
