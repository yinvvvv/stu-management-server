package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

func GetCourseByClass(c *gin.Context) {
	fmt.Println("getCourseByClass()")
}

func GetRankByClass(c *gin.Context) {
	fmt.Println("getRankByClass()")
}

func GetClassList(c *gin.Context) {
	fmt.Println("getClassList()")
	majorId := c.Param("major_id")
	sql := "select zjw_class.id, zjw_major.title, zjw_class.year, zjw_class.sub_id " +
		"from zjw_major, zjw_class " +
		"where zjw_class.major_id = zjw_major.id " +
		"and zjw_major.id = ?"
	stmt, _ := util.Db.Prepare(sql)
	rows, _ := stmt.Query(majorId)
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
