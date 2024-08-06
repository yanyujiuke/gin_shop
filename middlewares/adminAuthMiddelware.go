package middlewares

import (
	"encoding/json"
	"fmt"
	"ginShop/common"
	"ginShop/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
	"strings"
)

func InitAdminAuthMiddlewarec(c *gin.Context) {
	// 获取Url访问的地址   /admin/captcha?t=0.8706946438889653
	pathname := strings.Split(c.Request.URL.String(), "?")[0]

	managerJson, ok := common.RedisClient{}.Get("manager_info")

	var username string
	var isSuper int
	var roleId int
	if ok {
		var managerList []models.Manager
		json.Unmarshal([]byte(managerJson), &managerList)
		username = managerList[0].Username
		isSuper = managerList[0].IsSuper
		roleId = managerList[0].RoleId

		if username == "" {
			c.Redirect(http.StatusFound, "/admin/login")
		} else {
			// 用户登录成功 权限判断
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			if isSuper == 0 && !excludeAuthPath("/"+urlPath) {
				// 根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
				roleAccess := []models.RoleAccess{}
				common.DB.Where("role_id = ?", roleId).Find(&roleAccess)

				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}

				// 获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
				// pathname   /admin/manager
				access := models.Access{}
				common.DB.Where("url = ?", urlPath).Find(&access)

				// 判断当前访问的url对应的权限id 是否在权限列表的id中
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort()
				}
			}
		}
	} else {
		c.Redirect(http.StatusFound, "/admin/login")
	}
}

// 排除权限判断的方法
func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()

	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	// return true
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
