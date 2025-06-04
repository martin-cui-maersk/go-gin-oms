package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-oms/server/models"
	"go-gin-oms/server/utils/result"
	"go-gin-oms/server/utils/token"
)

// ReqRegister /api/register的请求体
type ReqRegister struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ReqLogin api/login 的请求体
type ReqLogin struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req ReqRegister

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"data": err.Error(),
		//})
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}

	u := models.SysUser{
		UserName: req.Account,
		Password: req.Password,
	}

	_, err := u.SaveUser()
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"data": err.Error(),
		//})
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "register success",
	//	"data":    req,
	//})
	result.Response().SetCode(200).SetMsg("register success").SetData(req).Build(c)
}

func Login(c *gin.Context) {
	var req ReqLogin
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		if err.Error() == "EOF" {
			result.Response().SetCode(403).SetMsg("params cannot be empty").SetData(nil).Build(c)
			return
		}
		result.Response().SetCode(503).SetMsg(err.Error()).Build(c)
		return
	}

	u := models.SysUser{
		UserName: req.Account,
		Password: req.Password,
	}
	// 调用 models.LoginCheck 对用户名和密码进行验证
	token, err := models.LoginCheck(u.UserName, u.Password)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": "account or password is incorrect.",
		//})
		//result.Response().SetCode(503).SetMsg("account or password is incorrect.").Build(c)
		result.Response().SetCode(503).SetMsg(err.Error()).Build(c)
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"token": token,
	//})
	result.Response().SetData(map[string]string{"token": token}).Build(c)
}

// CurrentUserInfo 获取当前用户信息
func CurrentUserInfo(c *gin.Context) {
	// 从token中解析出user_id
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": err.Error(),
		//})
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}

	// 根据user_id从数据库查询数据
	u, err := models.GetUserInfoByID(userId)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": err.Error(),
		//})
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "success",
	//	"data":    u,
	//})
	result.Response().SetData(u).Build(c)
}

// GetMyMenuList 获取角色的菜单
func GetMyMenuList(c *gin.Context) {
	var err error
	// 从token中解析出user_id
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": err.Error(),
		//})
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}
	userInfo, err := models.GetUserInfoByID(userId)
	if err != nil {
		result.Response().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}
	results := models.GetRoleMenu(userInfo.RoleId)
	result.Response().SetData(results).Build(c)
}

// GetPermissionCode 获取角色的 Permission Code
func GetPermissionCode(c *gin.Context) {
	result.Response().SetCode(200).SetData(models.GetPermissionCode()).Build(c)
}
