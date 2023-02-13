package service

import (
	"cloud_homestay/helper"
	"cloud_homestay/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 注册
func UserLogup(c *gin.Context) {
	email := c.PostForm("email1")
	getcode := c.PostForm("code1")
	password := c.PostForm("password")
	var cnt int64
	err := models.DB.Model(new(models.UserBasic)).Where("email = ?", email).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User Error:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
			"code": -1,
			"msg":  "邮箱已被注册，请重新输入邮箱地址!",
			"href": "login",
		})
		return
	}
	code, _ := models.RDB.Get(email).Result()
	if err != nil {
		c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
			"code": -1,
			"msg":  "请重新获取验证码",
			"href": "login",
		})
		return
	}
	if strings.Compare(getcode, code) != 0 {
		c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
			"code": -1,
			"msg":  "邮箱验证码输入错误，请重新输入",
			"href": "login",
		})
		return
	}
	password = helper.GetMd5(password)
	user := &models.UserBasic{
		Identity: helper.GetUUid(),
		Name:     email,
		Email:    email,
		Password: password,
	}
	err = models.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create User Error:" + err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
		"code": 200,
		"msg":  "注册成功，请重新登录",
		"href": "login",
	})
}

// 登录
func UserLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var cnt1 int64
	var cnt2 int64
	err := models.DB.Model(new(models.UserBasic)).Where("email = ?", email).Count(&cnt1).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserEmain Error:" + err.Error(),
		})
		return
	}
	if cnt1 > 0 {
		password = helper.GetMd5(password)
		err := models.DB.Model(new(models.UserBasic)).Where("email = ? AND password = ?", email, password).Count(&cnt2).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "Get UserPassword Error:" + err.Error(),
			})
			return
		}
		if cnt2 > 0 {
			// 成功
			// 获取identity
			var us models.UserBasic
			err = models.DB.Model(new(models.UserBasic)).Where("email = ?", email).First(&us).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "Get Identity Error:" + err.Error(),
				})
				return
			}
			if us.UserStatus == 1 {
				c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
					"code": -1,
					"msg":  "该用户已被禁用",
					"href": "login",
				})
				return
			}
			sesssion := sessions.Default(c)
			// 将us结构体存入session中
			sesssion.Set("user", us.Email)
			sesssion.Save()
			c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
				"code": 200,
				"msg":  "登录成功!",
				"href": "home",
			})
		} else {
			c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
				"code": -1,
				"msg":  "密码错误，请重新输入",
				"href": "login",
			})
		}
	} else {
		c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
			"code": -1,
			"msg":  "email不存在，请注册",
			"href": "login",
		})
	}
}

// 发送验证码
func LoginPage(c *gin.Context) {
	email := c.Query("email2")
	if email != "" {
		code := helper.GetRandCode()
		fmt.Println(code)
		models.RDB.Set(email, code, time.Minute*2)
		fmt.Println(models.RDB.Get(email))
		helper.SendCodeByEmail(email, code)
		c.HTML(http.StatusOK, "login/login.html", gin.H{})
	} else {
		c.HTML(http.StatusOK, "login/login.html", gin.H{})
	}
}

// 获取用户列表
func GetUsers(c *gin.Context) {
	users := make([]*models.UserBasic, 0)
	err := models.DB.Model(new(models.UserBasic)).Find(&users).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User List Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"users": users,
	})
}

func User(c *gin.Context) {
	users := make([]*models.UserBasic, 0)
	us, _ := c.Get("user")
	user := helper.GetUserStruct(us.(string))
	err := models.DB.Model(new(models.UserBasic)).
		Find(&users).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Admin Room List Error:" + err.Error(),
		})
		return
	}
	a := user
	c.HTML(http.StatusOK, "admin/user.html", gin.H{
		"user":  a,
		"users": users,
	})
}
func Message(c *gin.Context) {
	messages := make([]*models.CommentRoom, 0)
	us, _ := c.Get("user")
	user := helper.GetUserStruct(us.(string))
	err := models.DB.Model(new(models.CommentRoom)).Preload("RoomInfo").Preload("UserInfo").
		Find(&messages).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Admin Room List Error:" + err.Error(),
		})
		return
	}
	a := user
	c.HTML(http.StatusOK, "admin/message.html", gin.H{
		"user":     a,
		"messages": messages,
	})
}

func Rmsession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
		"code": 200,
		"msg":  "退出成功！",
		"href": "../login",
	})
}

func UserCT(c *gin.Context) {
	us, _ := c.Get("user")
	user := helper.GetUserStruct(us.(string))
	a := user
	c.HTML(http.StatusOK, "admin/userct.html", gin.H{
		"user": a,
	})
}

func UserInfo(c *gin.Context) {
	us, _ := c.Get("user")
	user := helper.GetUserStruct(us.(string))
	a := user
	c.HTML(http.StatusOK, "admin/userinfo.html", gin.H{
		"user": a,
	})
}

// 处理ajax发送来的banuser请求
func BanUser(c *gin.Context) {
	usname := c.PostForm("user")
	ust := c.PostForm("ust")
	status := 0
	st := "启用"
	if strings.Compare(ust, st) == 0 {
		status = 0
	} else {
		status = 1
	}
	err := models.DB.Model(new(models.UserBasic)).Where("name = ?", usname).Update("user_status", status).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Update User_tatus Error:" + err.Error(),
		})
		return
	}
}

func Personal(c *gin.Context) {
	us, _ := c.Get("user")
	user := new(models.UserBasic)
	if us != nil {
		user = helper.GetUserStruct(us.(string))
	}
	comments := make([]*models.CommentRoom, 0)
	var cnt int64
	err := models.DB.Model(new(models.CommentRoom)).Preload("RoomInfo").Preload("UserInfo").
		Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Personal Comment Error:" + err.Error(),
		})
		return
	}
	orders := make([]*models.OrderBasic, 0)
	err = models.DB.Model(new(models.OrderBasic)).Where("user_id = ?", user.ID).Find(&orders).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Personal Order Error:" + err.Error(),
		})
		return
	}
	a := user
	if user.Name != "" {
		c.HTML(http.StatusOK, "home/personal_info.html", gin.H{
			"name":     a.Name,
			"isadmin":  a.Isadmin,
			"user":     a,
			"comments": comments,
			"cnt":      cnt,
			"orders":   orders,
		})
	} else {
		c.HTML(http.StatusOK, "home/personal_info.html", gin.H{
			"name":    "登录",
			"isadmin": 0,
			"user":    a,
		})
	}
}

func Comment(c *gin.Context) {
	roomid := c.PostForm("roomid")
	userid := c.PostForm("userid")
	content := c.PostForm("content")
	usericon := c.PostForm("usericon")
	comment := &models.CommentRoom{
		RoomId:   roomid,
		UserId:   userid,
		Content:  content,
		UserIcon: usericon,
	}
	models.DB.Create(&comment)
}

func DeleteComment(c *gin.Context) {
	id := c.PostForm("id")
	comment := &models.CommentRoom{}
	models.DB.Where("id = ?", id).Delete(&comment)
}
