package models

import (
	models "github.com/martin-cui-maersk/go-gin-oms/models/common"
)

type SysMenu struct {
	//gorm.Model
	MenuId                 uint   `gorm:"primaryKey;autoIncrement;comment:菜单ID"  json:"menuId"`
	MenuType               uint   `json:"menuType" gorm:"size:10;comment:菜单类型：0-目录|1-菜单|2-按钮"`
	MenuName               string `json:"menuName" gorm:"size:100;comment:菜单名称"`
	MetaTitle              string `json:"metaTitle" gorm:"size:100;comment:菜单语言包"`
	MenuPath               string `json:"menuPath" gorm:"size:100;comment:菜单路径"`
	ApiPath                string `json:"apiPath" gorm:"size:100;comment:接口路径"`
	HideMenu               uint   `json:"hideMenu" gorm:"size:10;comment:是否是隐藏菜单"`
	Component              string `json:"component" gorm:"size:100;comment:前端Component"`
	Redirect               string `json:"redirect" gorm:"size:100;comment:前端跳转路径"`
	MetaIcon               string `json:"metaIcon" gorm:"size:50;comment:前端菜单图标"`
	MetaPermission         string `json:"metaPermission" gorm:"size:50;comment:前端权限标识"`
	MetaHideChildrenInMenu uint   `json:"metaHideChildrenInMenu" gorm:"size:10;comment:"`
	MetaIgnoreKeepAlive    uint   `json:"metaIgnoreKeepAlive" gorm:"size:10;comment:是否缓存"`
	MetaHideBreadCrumb     uint   `json:"metaHideBreadCrumb" gorm:"size:10;comment:是否缓存"`
	MetaCurrentActiveMenu  string `json:"metaCurrentActiveMenu" gorm:"size:100;comment:当前高亮菜单路径"`
	ParentId               uint   `json:"parentId" gorm:"size:10;comment:父级菜单ID"`
	Sort                   uint   `json:"sort" gorm:"size:10;comment:菜单排序"`
	Status                 uint   `json:"status" gorm:"index;comment:1-启用|2-禁用"`
	models.IsActive
	models.ControlBy
	models.ModelTime
}

// TableName 添加表前缀
func (*SysMenu) TableName() string {
	return "oms_sys_menu"
}
