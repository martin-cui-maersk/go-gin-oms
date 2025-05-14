package models

type SysRoleMenu struct {
	RoleId uint `json:"roleId" gorm:"size:10;comment:角色ID"`
	MenuId uint `json:"menuId" gorm:"size:10;comment:菜单ID"`
}

// TableName 添加表前缀
func (*SysRoleMenu) TableName() string {
	return "oms_sys_role_menu"
}
