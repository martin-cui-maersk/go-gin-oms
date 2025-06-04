package models

import (
	"go-gin-oms/server/global"
	models "go-gin-oms/server/models/common"
	"gorm.io/gorm"
)

type SysRole struct {
	RoleId     uint   `gorm:"primaryKey;autoIncrement;comment:角色ID"  json:"roleId"`
	RoleName   string `json:"roleName" gorm:"size:100;comment:角色名称"`
	RoleCode   string `json:"roleCode" gorm:"size:100;comment:角色Code"`
	MerchantId uint   `json:"merchantId" gorm:"size:10;comment:是否是隐藏菜单"`
	Status     uint   `json:"status" gorm:"index;comment:1-启用|2-禁用"`
	Remarks    string `json:"remarks" gorm:"size:500;comment:备注"`
	models.IsActive
	models.ControlBy
	models.ModelTime
}

// TableName 添加表前缀
func (*SysRole) TableName() string {
	return "oms_sys_role"
}

// AfterFind 钩子会在查询完成后自动调用
func (s *SysRole) AfterFind(tx *gorm.DB) (err error) {
	// 将 int64 时间戳转换为 time.Time
	//t1 := time.Unix(s.CreateAt/1000, 0)
	//t2 := time.Unix(s.UpdateAt/1000, 0)

	// 格式化为 YYYY-MM-DD HH:MM:SS
	s.FormattedCreatedAt = models.FormatUnixTime(s.CreateAt)
	s.FormattedUpdatedAt = models.FormatUnixTime(s.UpdateAt)
	return nil

}

// GetRoleList 获取角色列表
func GetRoleList(params map[string]interface{}) (int64, []SysRole) {
	var result []SysRole
	var count int64
	page := params["page"].(int)
	pageSize := params["pageSize"].(int)
	offset := (page - 1) * pageSize
	query := global.DB.Model(&SysRole{})
	if params["roleName"].(string) != "" {
		query = query.Where("role_name like ?", "%"+params["roleName"].(string)+"%")
	}
	if params["roleCode"].(string) != "" {
		query = query.Where("role_code like ?", "%"+params["roleCode"].(string)+"%")
	}
	if params["status"].(int) > 0 {
		query = query.Where("status = ?", params["status"])
	}
	// 先计数再查询
	query.Count(&count).Limit(pageSize).Offset(offset).Order("role_id desc, update_at desc").Find(&result)
	return count, result
}

func SetRoleStatus() {

}
