package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

// /class

func GetCourseByClass(c *gin.Context) {
	fmt.Println("getCourseByClass()")
	id := c.Query("id")
	sql := "select zjw_course.* " +
		"from zjw_class, zjw_course " +
		"where zjw_class.id = zjw_course.class_id " +
		"and zjw_class.id = ?; "
	stmt, err := util.Db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	rows, err2 := stmt.Query(id)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}

// GetRankByClass TODO:
func GetRankByClass(c *gin.Context) {
	fmt.Println("getRankByClass()")
}

func GetClassList(c *gin.Context) {
	fmt.Println("getClassList()")
	majorId := c.Query("major_id")
	sql := `
		select zjw_class.id, zjw_major.title, zjw_class.year, zjw_class.sub_id 
		from zjw_major, zjw_class 
		where zjw_class.major_id = zjw_major.id 
		and zjw_major.id = ?
	`
	stmt, _ := util.Db.Prepare(sql)
	rows, _ := stmt.Query(majorId)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
