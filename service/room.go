package service

import (
	"cloud_homestay/helper"
	"cloud_homestay/models"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取房间列表
func GetRooms(c *gin.Context) {
	od := c.Query("type")
	rooms := make([]*models.RoomBasic, 0)
	var err error
	var cnt int64
	err = models.DB.Model(new(models.RoomBasic)).Find(&rooms).Offset(1).Limit(6).Count(&cnt).Error
	if od == "0" {
		// 默认排序
		err = models.DB.Model(new(models.RoomBasic)).Find(&rooms).Offset(1).Limit(6).Count(&cnt).Error
	} else if od == "1" {
		// 销量升序
		err = models.DB.Model(new(models.RoomBasic)).Find(&rooms).Offset(1).Limit(6).Order("sales ASC").Count(&cnt).Error
	} else if od == "2" {
		// 销量降序
		err = models.DB.Model(new(models.RoomBasic)).Find(&rooms).Offset(1).Limit(6).Order("sales DESC").Count(&cnt).Error
	} else if od == "3" {
		// 价格升序
		err = models.DB.Model(new(models.RoomBasic)).Find(&rooms).Offset(1).Limit(6).Order("price ASC").Count(&cnt).Error
	} else if od == "4" {
		// 价格降序
		err = models.DB.Model(new(models.RoomBasic)).Find(&rooms).Offset(1).Limit(6).Order("price DESC").Count(&cnt).Error
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Room List Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"rooms": rooms,
	})
}

// 创建房间的提交页面
func Creat(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/creatroom.html", gin.H{})
}

// 修改房间的提交页面
func Change(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/changeroom.html", gin.H{})
}

// 删除房间业务
func DeleteRoom(c *gin.Context) {
	title := c.PostForm("title")
	room := models.RoomBasic{}
	models.DB.Model(&room).Where("title = ?", title).Delete(&room)
}

// 获取房间销量排行榜
func GetRankRoomBySales(c *gin.Context) {
	rooms := make([]*models.RoomBasic, 0)
	err := models.DB.Model(new(models.RoomBasic)).Find(&rooms).Offset(1).Limit(10).Order("sales ASC").Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Room Rank List Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"rooms": rooms,
	})
}

func Room(c *gin.Context) {
	rooms := make([]*models.RoomBasic, 0)
	us, _ := c.Get("user")
	user := helper.GetUserStruct(us.(string))
	err := models.DB.Model(new(models.RoomBasic)).Find(&rooms).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Admin Room List Error:" + err.Error(),
		})
		return
	}
	a := user
	c.HTML(http.StatusOK, "admin/room.html", gin.H{
		"user":  a,
		"rooms": rooms,
	})
}

func Order(c *gin.Context) {
	orders := make([]*models.OrderBasic, 0)
	us, _ := c.Get("user")
	user := helper.GetUserStruct(us.(string))

	err := models.DB.Model(new(models.OrderBasic)).Preload("RoomInfo").Preload("UserInfo").
		Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Admin Room List Error:" + err.Error(),
		})
		return
	}
	a := user
	c.HTML(http.StatusOK, "admin/order.html", gin.H{
		"user":   a,
		"orders": orders,
	})
}

func RoomAdd(c *gin.Context) {
	title := c.PostForm("title")
	info := c.PostForm("info")
	price := c.PostForm("price")
	pr, _ := strconv.Atoi(price)
	ty := c.PostForm("type")
	t, _ := strconv.Atoi(ty)
	content := c.PostForm("content")
	room_num := c.PostForm("room_num")

	dst := path.Join("./static/images", "room-"+room_num+".png")
	img, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Creat Img Error" + err.Error(),
		})
		return
	} else {
		c.SaveUploadedFile(img, dst)
	}
	room := &models.RoomBasic{
		Title:   title,
		Info:    info,
		Content: content,
		Price:   pr,
		Type:    t,
		Img:     "../" + dst,
		RoomNum: room_num,
	}
	models.DB.Create(&room)
}

func RoomList(c *gin.Context) {
	us, _ := c.Get("user")
	user := new(models.UserBasic)
	roomrank := make([]*models.RoomBasic, 0)
	if us != nil {
		user = helper.GetUserStruct(us.(string))
	}
	err := models.DB.Model(new(models.RoomBasic)).Limit(10).Offset(0).Order("sales DESC").Find(&roomrank).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get RoomRank Error" + err.Error(),
		})
		return
	}
	a := user
	if user.Name != "" {
		c.HTML(http.StatusOK, "home/room.html", gin.H{
			"name":     a.Name,
			"isadmin":  a.Isadmin,
			"user":     a,
			"roomrank": roomrank,
		})
	} else {
		c.HTML(http.StatusOK, "home/room.html", gin.H{
			"name":     "登录",
			"isadmin":  0,
			"user":     a,
			"roomrank": roomrank,
		})
	}
}

func Room1(c *gin.Context) {
	roomrank := make([]*models.RoomBasic, 0)
	err := models.DB.Model(new(models.RoomBasic)).Find(&roomrank).Error
	//.Order("sales DESC")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get RoomRank Error" + err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "home/room-1.html", gin.H{
		"roomrank": roomrank,
	})
}

func Room2(c *gin.Context) {
	roomrank := make([]*models.RoomBasic, 0)
	err := models.DB.Model(new(models.RoomBasic)).Order("sales DESC").Find(&roomrank).Error
	//.Order("sales DESC")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get RoomRank Error" + err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "home/room-sales.html", gin.H{
		"roomrank": roomrank,
	})
}
func Room3(c *gin.Context) {
	roomrank := make([]*models.RoomBasic, 0)
	err := models.DB.Model(new(models.RoomBasic)).Order("price DESC").Find(&roomrank).Error
	//.Order("sales DESC")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get RoomRank Error" + err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "home/room-price.html", gin.H{
		"roomrank": roomrank,
	})
}

func RoomInfo(c *gin.Context) {
	title := c.Query("title")
	us, _ := c.Get("user")
	user := new(models.UserBasic)
	comments := make([]*models.CommentRoom, 0)
	if us != nil {
		user = helper.GetUserStruct(us.(string))
	}
	room := new(models.RoomBasic)
	err := models.DB.Model(new(models.RoomBasic)).Where("title = ?", title).First(&room).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get RoomInfo Error:" + err.Error(),
		})
		return
	}
	var cnt int64
	err = models.DB.Model(new(models.CommentRoom)).Where("room_id = ?", room.ID).Find(&comments).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Comments Error:" + err.Error(),
		})
		return
	}
	a := user
	c.HTML(http.StatusOK, "home/roominfo.html", gin.H{
		"room":     room,
		"name":     a.Name,
		"isadmin":  a.Isadmin,
		"user":     a,
		"comments": comments,
		"cnt":      cnt,
	})
}
