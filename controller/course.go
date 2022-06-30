package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stu-management-server/util"
)

// /course

func GetCourseList(c *gin.Context) {
	fmt.Println("getCourseList()")
	class_id := c.Query("class_id")
	sql := "Select * from zjw_course where class_id = ?;"
	stmt, _ := util.Db.Prepare(sql)
	defer stmt.Close()
	rows, _ := stmt.Query(class_id)
	defer rows.Close()
	result := util.QueryAndParseRows(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}

func GetCourseRank(c *gin.Context) {
	fmt.Println("getCourseRank()")
	id := c.Query("id")
	sql := `
		SELECT *, RANK() OVER (ORDER BY score DESC) rank_num
		FROM (SELECT student_id, student_name, score FROM zjw_course_rank
		WHERE course_id = ?) r;
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

func GetCourseAvg(c *gin.Context) {
	fmt.Println("getCourseAvg()")
	id := c.Query("id")
	sql := `SELECT AVG(score) as avg
		FROM (SELECT student_id, student_name, score FROM zjw_course_rank
		WHERE course_id = ?) r;`
	stmt, _ := util.Db.Prepare(sql)
	defer stmt.Close()
	rows, _ := stmt.Query(id)
	defer rows.Close()
	result := util.QueryAndParse(rows)
	c.JSON(200, gin.H{
		"code": 0,
		"data": result,
	})
}
