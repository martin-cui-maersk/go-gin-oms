package system

import (
	"github.com/gin-gonic/gin"
	"github.com/martin-cui-maersk/go-gin-oms/models"
	"github.com/martin-cui-maersk/go-gin-oms/utils"
	"github.com/martin-cui-maersk/go-gin-oms/utils/result"
)

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	var err error
	var exists bool
	page, exists, err := utils.GetQueryInt(c, "page")
	if err != nil {
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}
	if !exists {
		page = 1 // 默认值
	}
	pageSize, exists, err := utils.GetQueryInt(c, "pageSize")
	if err != nil {
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}
	if !exists {
		pageSize = 10 // 默认值
	}
	status, exists, err := utils.GetQueryInt(c, "status")
	if err != nil {
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}
	if !exists {
		status = 0 // 默认值
	}
	roleId, exists, err := utils.GetQueryInt(c, "roleId")
	if err != nil {
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}
	if !exists {
		roleId = 0 // 默认值
	}
	params := map[string]interface{}{
		// "page":     c.DefaultQuery("page", "1"),
		"page":     page,
		"pageSize": pageSize,
		"roleName": c.Query("roleName"),
		"roleCode": c.Query("roleCode"),
		"userName": c.DefaultQuery("name", ""),
		"email":    c.DefaultQuery("email", ""),
		"status":   status,
		"roleId":   roleId,
	}
	total, list := models.GetUserList(params)
	data := result.DataList{List: list, Total: total}
	result.Response().SetData(data).Build(c)
}

// GetRoleSelect 用户列表select
func GetRoleSelect(c *gin.Context) {
	result.Response().SetData(models.GetRoleSelect()).Build(c)
}
