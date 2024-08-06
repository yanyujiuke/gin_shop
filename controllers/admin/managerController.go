package admin

import (
	"ginShop/common"
	"ginShop/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	managerList := []models.Manager{}

	common.DB.Preload("Role").Find(&managerList)

	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})

}

func (con ManagerController) Add(c *gin.Context) {
	roleList := []models.Role{}

	common.DB.Find(&roleList)
	if len(roleList) <= 0 {
		con.Error(c, "请先创建角色", "/admin/role/add")
	}

	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	roleId, _ := common.StringToInt(c.PostForm("role_id"))

	// 用户名和密码长度是否合法
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或者密码的长度不合法", "/admin/manager/add")
		return
	}

	// 判断管理是否存在
	managerList := []models.Manager{}
	common.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "此管理员已存在", "/admin/manager/add")
		return
	}

	manager := models.Manager{
		Username: username,
		Password: common.Md5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   roleId,
		Status:   1,
		AddTime:  int(common.GetUnix()),
	}

	err := common.DB.Create(&manager).Error
	if err != nil {
		con.Error(c, "添加管理员失败", "/admin/manager/add")
		return
	}
	con.Success(c, "添加管理员成功", "/admin/manager")
}

func (con ManagerController) Edit(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	manager := models.Manager{Id: id}
	common.DB.Find(&manager)

	roleList := []models.Role{}
	common.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err1 := common.StringToInt(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	roleId, err2 := common.StringToInt(c.PostForm("role_id"))
	if err2 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")

	if len(mobile) > 11 {
		con.Error(c, "mobile长度不合法", "/admin/manager/edit?id="+common.IntToString(id))
		return
	}

	// 执行修改
	manager := models.Manager{Id: id}
	common.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId

	// 注意：判断密码是否为空 为空表示不修改密码 不为空表示修改密码
	if password != "" {
		// 判断密码长度是否合法
		if len(password) < 6 {
			con.Error(c, "密码的长度不合法 密码长度不能小于6位", "/admin/manager/edit?id="+common.IntToString(id))
			return
		}
		manager.Password = common.Md5(password)
	}
	err3 := common.DB.Save(&manager).Error
	if err3 != nil {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+common.IntToString(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/manager")
}

func (con ManagerController) Delete(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	manager := models.Manager{Id: id}
	err := common.DB.Delete(&manager).Error
	if err != nil {
		con.Error(c, "管理员删除失败", "/admin/manager")
		return
	}

	con.Success(c, "管理员删除成功", "/admin/manager")
}
