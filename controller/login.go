package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login login and store session
func Login(c *gin.Context) {
	var userForm UserForm
	err := c.BindJSON(&userForm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userForm)

	if userForm.Username != "admin" || userForm.Password != "admin" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user", userForm.Username)
	session.Save()
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "登录成功",
	})
}
