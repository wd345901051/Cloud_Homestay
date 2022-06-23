package service

import (
	"cloud_homestay/helper"
	"cloud_homestay/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	us, _ := c.Get("user")
	rooms := make([]*models.RoomBasic, 0)
	user := new(models.UserBasic)
	if us != nil {
		user = helper.GetUserStruct(us.(string))
	}

	var cnt int64
	var err error
	err = models.DB.Model(new(models.RoomBasic)).Find(&rooms).Offset(1).Limit(6).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Rooms Error:" + err.Error(),
		})
		return
	}
	if user.Name != "" {
		isadmin := user.Isadmin
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"name":    user.Name,
			"rooms":   rooms,
			"isadmin": isadmin,
		})
	} else {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"name":    "登录",
			"rooms":   rooms,
			"isadmin": 0,
		})
	}
}

func Goods(c *gin.Context) {
	goods := make([]*models.GoodBasic, 0)
	us, _ := c.Get("user")
	user := helper.GetUserStruct(us.(string))
	err := models.DB.Model(new(models.GoodBasic)).
		Find(&goods).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Admin Goods List Error:" + err.Error(),
		})
		return
	}
	a := user
	c.HTML(http.StatusOK, "admin/goods.html", gin.H{
		"user":  a,
		"goods": goods,
	})
}

func Equip(c *gin.Context) {
	equips := make([]*models.EquipBasic, 0)
	us, _ := c.Get("user")
	user := helper.GetUserStruct(us.(string))
	err := models.DB.Model(new(models.EquipBasic)).
		Find(&equips).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Admin Equip List Error:" + err.Error(),
		})
		return
	}
	a := user
	c.HTML(http.StatusOK, "admin/equip.html", gin.H{
		"user":   a,
		"equips": equips,
	})
}

func DeleteGood(c *gin.Context) {
	good := c.PostForm("good")
	gd := new(models.GoodBasic)
	models.DB.Where("good_name", good).Delete(&gd)
}

func GoodAdd(c *gin.Context) {
	good_name := c.PostForm("good_name")
	good_price := c.PostForm("good_price")
	gp, _ := strconv.Atoi(good_price)
	good_num := c.PostForm("good_num")
	gn, _ := strconv.Atoi(good_num)
	good_content := c.PostForm("good_content")
	good := &models.GoodBasic{
		GoodName:    good_name,
		GoodPrice:   gp,
		GoodNum:     gn,
		GoodContent: good_content,
	}
	models.DB.Create(&good)
}
