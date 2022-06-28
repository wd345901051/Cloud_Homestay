package service

import (
	"cloud_homestay/helper"
	"cloud_homestay/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Attention(c *gin.Context) {
	// 直接处理订阅任务
	us, _ := c.Get("user")
	user := new(models.UserBasic)
	if us != nil {
		user = helper.GetUserStruct(us.(string))
	}
	err := models.DB.Model(new(models.UserBasic)).Where("email = ?", user.Email).Update("attention", 1).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Update Attention Error:" + err.Error(),
		})
		return
	} else {
		c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
			"code": 200,
			"msg":  "订阅成功!",
			"href": "home",
		})
		helper.SendCodeByEmail(user.Email, "欢迎您订阅本频道，(๑*◡*๑) Cloud Homestay Happy~")
	}
}

// 检查是否为管理员
func CheckAdmin(c *gin.Context) {
	us, _ := c.Get("user")
	user := new(models.UserBasic)
	if us != nil {
		user = helper.GetUserStruct(us.(string))
	}
	if us == nil {
		c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
			"code": -1,
			"msg":  "请先登录",
			"href": "../login",
		})
	}
	if user.Isadmin == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Not Have rout To This Page!",
		})
		c.Abort()
	}
	c.Next()
}

func CheckLogin(c *gin.Context) {
	us, _ := c.Get("user")
	if us == nil {
		c.HTML(http.StatusOK, "temp/login_info.html", gin.H{
			"code": 200,
			"msg":  "请先登录!",
			"href": "login",
		})
		c.Abort()
		return
	}
	c.Next()
}

func OrderPush(c *gin.Context) {
	us, _ := c.Get("user")
	user := new(models.UserBasic)
	if us != nil {
		user = helper.GetUserStruct(us.(string))
	}
	name := c.PostForm("name")
	phone := c.PostForm("tel")
	num := c.PostForm("id-number")
	cnt := c.PostForm("number")
	ct, _ := strconv.Atoi(cnt)
	timein := c.PostForm("timein")
	timeout := c.PostForm("timeout")
	content := c.PostForm("content")
	roomid := c.PostForm("roomid")
	ri, _ := strconv.Atoi(roomid)
	order := &models.OrderBasic{
		OrderNo:   helper.GetUUid(),
		UserName:  name,
		UserPhone: phone,
		UserNum:   num,
		PeopleCnt: ct,
		RoomIn:    timein,
		RoomOut:   timeout,
		Content:   content,
		UserId:    int(user.ID),
		RoomId:    ri,
	}
	models.DB.Create(&order)
}
