package system

import (
	"github.com/gin-gonic/gin"
	"github.com/martin-cui-maersk/go-gin-oms/models"
	"github.com/martin-cui-maersk/go-gin-oms/utils"
	"github.com/martin-cui-maersk/go-gin-oms/utils/result"
)

func GetRoleList(c *gin.Context) {
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
	params := map[string]interface{}{
		// "page":     c.DefaultQuery("page", "1"),
		"page":     page,
		"pageSize": pageSize,
		"roleName": c.Query("roleName"),
		"roleCode": c.Query("roleCode"),
		"status":   status,
	}
	total, list := models.GetRoleList(params)
	data := result.DataList{List: list, Total: total}
	result.Response().SetData(data).Build(c)
}

func AddRole(c *gin.Context)       {}
func UpdateRole(c *gin.Context)    {}
func SetRoleStatus(c *gin.Context) {}
