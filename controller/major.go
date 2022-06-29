package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

// /major

func GetMajorList(c *gin.Context) {
	fmt.Println("getMajorList()")
	sql := "SELECT * FROM zjw_major"
	rows, _ := util.Db.Query(sql)
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
