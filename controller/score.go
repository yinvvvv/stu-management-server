package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"stu-management-server/util"
)

type ScoreForm struct {
	StudentId int `json:"student_id"`
	CourseId  int `json:"course_id"`
	Score     int `json:"score"`
}

func GetAvgScoreByCourse(c *gin.Context) {
	fmt.Println("getAvgScoreByCourse()")
}

func AddScore(c *gin.Context) {
	fmt.Println("addScore()")
	var scoreForm ScoreForm
	err := c.BindJSON(&scoreForm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(scoreForm)
	sql := "insert into zjw_score (student_id, course_id, score) values (?, ?, ?);"
	stmt, _ := util.Db.Prepare(sql)
	defer stmt.Close()
	rs, err2 := stmt.Exec(scoreForm.StudentId, scoreForm.CourseId, scoreForm.Score)

	code := 0
	msg := ""

	if err2 != nil {
		fmt.Println(err2)
		code = 1
		msg = err2.Error()
	} else {
		ra, _ := rs.RowsAffected()
		msg := "Affected rows: " + strconv.FormatInt(ra, 10)
		fmt.Println(msg)
	}

	c.JSON(200, gin.H{
		"code": code,
		"data": "",
		"msg":  msg,
	})
}

func GetStudentScore(c *gin.Context) {
	fmt.Println("getStudentScore()")
}

func GetStudentScoreByYear(c *gin.Context) {
	fmt.Println("getStudentScoreByYear()")
}
