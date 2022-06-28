package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

func GetMajorList(c *gin.Context) {
	fmt.Println("getMajorList()")
	sql := "SELECT * FROM zjw_major"
	result := util.QueryAndParseRows(util.Db, sql)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
