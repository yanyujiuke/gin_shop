package admin

import (
	"ginShop/common"
	"ginShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	accessList := []models.Access{}
	common.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})

}

func (con AccessController) Add(c *gin.Context) {
	accessList := []models.Access{}

	common.DB.Where("module_id = ?", 0).Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) DoAdd(c *gin.Context) {
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, typeErr := common.StringToInt(c.PostForm("type"))
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	url := strings.Trim(c.PostForm("url"), " ")
	moduleId, moduleIdErr := common.StringToInt(c.PostForm("module_id"))
	sort, sortErr := common.StringToInt(c.PostForm("sort"))
	description := strings.Trim(c.PostForm("description"), " ")
	status, statusErr := common.StringToInt(c.PostForm("status"))
	if typeErr != nil || moduleIdErr != nil || sortErr != nil || statusErr != nil {
		con.Error(c, "参数错误", "/admin/access/add")
		return
	}

	// 用户名和密码长度是否合法
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}

	access := models.Access{
		ModuleName:  moduleName,
		Type:        accessType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}

	createErr := common.DB.Create(&access).Error
	if createErr != nil {
		con.Error(c, "添加权限失败", "/admin/access/add")
		return
	}

	con.Success(c, "添加权限成功", "/admin/access")
}

func (con AccessController) Edit(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	access := models.Access{Id: id}
	common.DB.Find(&access)

	roleList := []models.Role{}
	common.DB.Find(&roleList)

	// 获取顶级模块
	accessList := []models.Access{}
	common.DB.Where("module_id=?", 0).Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access,
		"accessList": accessList,
	})
}

func (con AccessController) DoEdit(c *gin.Context) {
	id, idErr := common.StringToInt(c.PostForm("id"))
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, typeErr := common.StringToInt(c.PostForm("type"))
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	url := strings.Trim(c.PostForm("url"), " ")
	moduleId, moduleIdErr := common.StringToInt(c.PostForm("module_id"))
	sort, sortErr := common.StringToInt(c.PostForm("sort"))
	description := strings.Trim(c.PostForm("description"), " ")
	status, statusErr := common.StringToInt(c.PostForm("status"))

	if idErr != nil || typeErr != nil || moduleIdErr != nil || sortErr != nil || statusErr != nil {
		con.Error(c, "参数错误", "/admin/access/add")
		return
	}

	// 用户名和密码长度是否合法
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}

	// 执行修改
	access := models.Access{Id: id}
	common.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = accessType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status

	saveErr := common.DB.Save(&access).Error
	if saveErr != nil {
		con.Error(c, "修改数据失败", "/admin/access/edit?id="+common.IntToString(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/access")
}

func (con AccessController) Delete(c *gin.Context) {
	id, idErr := common.StringToInt(c.Query("id"))
	if idErr != nil {
		con.Error(c, "参数错误", "/admin/access")
		return
	}

	access := models.Access{Id: id}
	common.DB.Find(&access)

	// 顶级菜单，判断是否有子菜单
	if access.ModuleId == 0 {
		accessList := []models.Access{}
		common.DB.Where("module_id = ?", access.Id).Find(&accessList)
		if len(accessList) > 0 {
			con.Error(c, "该菜单还有子菜单，无法删除", "/admin/access")
			return
		}
	}

	common.DB.Delete(&access)
	con.Success(c, "权限删除成功", "/admin/access")
}
