package helper

import (
	"cloud_homestay/models"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

// 发送验证码
func SendCodeByEmail(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <wd345901051@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请及时输入"
	e.HTML = []byte("您的验证码: <b>" + code + "</b>")
	return e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "wd345901051@163.com", "DQLSRYXLNGDNNZSU", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}

// 生成验证码
func GetRandCode() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

// 生成md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

const KEY = "MiNa"

// 生成session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("session", store)
}

// 获取session
func GetUserSessionHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	c.Set("user", user)
	c.Next()
}

// 文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 生成UUID,唯一标识
func GetUUid() string {
	return uuid.NewV4().String()
}

func GetUserStruct(us string) *models.UserBasic {
	user := new(models.UserBasic)
	err := models.DB.Model(new(models.UserBasic)).Where("email = ?", us).First(&user).Error
	if err != nil {
		fmt.Println("Get User Struct Error:" + err.Error())
		return nil
	}
	return user
}
