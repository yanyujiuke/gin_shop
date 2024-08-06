package admin

import (
	"encoding/json"
	"fmt"
	"ginShop/common"
	"ginShop/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type IndexController struct {
	BaseController
}

func (con IndexController) Index(c *gin.Context) {
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
	}

	// 获取所有的权限
	accessList := []models.Access{}
	common.DB.Where("module_id = ?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
		return db.Order("access.sort DESC")
	}).Order("sort DESC").Find(&accessList)

	// 获取当前角色拥有的权限
	roleAccessList := []models.RoleAccess{}
	common.DB.Where("role_id = ?", roleId).Find(&roleAccessList)

	roleAccessMap := make(map[int]int)
	for _, v := range roleAccessList {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	// 循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
		accessItem := accessList[i].AccessItem
		for j := 0; j < len(accessItem); j++ {
			if _, ok := roleAccessMap[accessItem[j].Id]; ok {
				accessItem[j].Checked = true
			}
		}
	}

	c.HTML(http.StatusOK, "admin/index/index.html", gin.H{
		"username":   username,
		"isSuper":    isSuper,
		"accessList": accessList,
	})
}

func (con IndexController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/welcome.html", gin.H{})
}

func (con IndexController) ChangeStatus(c *gin.Context) {
	id, _ := common.StringToInt(c.PostForm("id"))
	table := strings.Trim(c.PostForm("table"), " ")
	field := strings.Trim(c.PostForm("field"), " ")

	err := common.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

// 修改单元格数据
func (con IndexController) ChangeCellValue(c *gin.Context) {
	id, _ := common.StringToInt(c.PostForm("id"))
	table := strings.Trim(c.PostForm("table"), " ")
	field := strings.Trim(c.PostForm("field"), " ")
	value := strings.Trim(c.PostForm("value"), " ")

	err := common.DB.Exec("update "+table+" set "+field+"="+value+" where id=?", id).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}
