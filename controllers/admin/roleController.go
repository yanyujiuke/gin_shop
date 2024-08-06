package admin

import (
	"ginShop/common"
	"ginShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	common.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(common.GetUnix())
	err := common.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "添加失败", "/admin/role/add")
	} else {
		con.Success(c, "添加成功", "/admin/role")
	}
}

func (con RoleController) Edit(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	role := models.Role{Id: id}
	err := common.DB.Find(&role).Error
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role": role,
	})
}

func (con RoleController) DoEdit(c *gin.Context) {
	id, _ := common.StringToInt(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	role := models.Role{Id: id}
	common.DB.Find(&role)
	role.Title = title
	role.Description = description

	err := common.DB.Save(&role).Error
	if err != nil {
		con.Error(c, "编辑失败", "/admin/role")
	} else {
		con.Success(c, "编辑成功", "/admin/role")
	}
}

func (con RoleController) Delete(c *gin.Context) {
	id, err := common.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		common.DB.Delete(&role)
		con.Success(c, "删除成功", "/admin/role")
	}
}

func (con RoleController) Auth(c *gin.Context) {
	roleId, _ := common.StringToInt(c.Query("role_id"))

	// 获取所有的权限
	accessList := []models.Access{}
	common.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)

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

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": accessList,
	})
}

func (con RoleController) DoAuth(c *gin.Context) {
	// 获取角色id
	roleId, _ := common.StringToInt(c.PostForm("role_id"))
	// 获取权限id  切片
	accessIds := c.PostFormArray("access_node[]")

	// 删除当前觉得的权限
	roleAccess := models.RoleAccess{}
	common.DB.Where("role_id = ?", roleId).Delete(&roleAccess)

	// 添加觉得权限
	for _, id := range accessIds {
		accessId, _ := common.StringToInt(id)
		roleAccess.RoleId = roleId
		roleAccess.AccessId = accessId
		common.DB.Create(&roleAccess)
	}

	con.Success(c, "授权成功", "/admin/role/auth?role_id="+common.IntToString(roleId))
}
