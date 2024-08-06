package admin

import (
	"encoding/json"
	"fmt"
	"ginShop/common"
	"ginShop/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (LoginController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (login LoginController) DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	captchaId := c.PostForm("captchaId")
	verify := c.PostForm("verifyValue")

	if flag := common.VerifyCaptcha(captchaId, verify); flag {
		managerList := []models.Manager{}
		password = common.Md5(password)
		err := common.DB.Where("username = ? AND password = ?", username, password).Find(&managerList).Error
		if len(managerList) > 0 && err == nil {
			managerJson, err := json.Marshal(managerList)
			if err != nil {
				log.Fatal(err)
			}
			eaa := common.RedisClient{}.Set("manager_info", string(managerJson), 0)
			fmt.Println(eaa)
			login.Success(c, "登录成功", "/admin")
		} else {
			login.Error(c, "用户名或密码错误", "/admin/login")
		}
	} else {
		login.Error(c, "验证码验证失败", "/admin/login")
	}
}

func (LoginController) Captcha(c *gin.Context) {
	id, base64, err := common.MakeCaptcha()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"captchaId":  id,
		"captchaImg": base64,
	})
}

func (login LoginController) LoginOut(c *gin.Context) {
	err := common.RedisClient{}.Del("manager_info")
	if err != nil {
		login.Error(c, "退出登录失败", "/admin")
	} else {
		login.Success(c, "退出登录成功", "/admin/login")
	}
}
