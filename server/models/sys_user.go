package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-gin-oms/server/global"
	models "go-gin-oms/server/models/common"
	"go-gin-oms/server/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"regexp"
	"strings"
)

type SysUser struct {
	//gorm.Model
	UserId   uint   `gorm:"primaryKey;autoIncrement;comment:用户ID"  json:"userId"`
	UserName string `json:"userName" gorm:"size:100;uniqueIndex:unique_index;comment:用户名"`
	Password string `json:"-" gorm:"size:100;comment:密码"`
	Salt     string `json:"salt" gorm:"size:100;comment:密码盐"`
	Email    string `json:"email" gorm:"size:100;uniqueIndex:unique_index;comment:邮箱"`
	RoleId   uint   `json:"roleId" gorm:"size:10;comment:角色ID"`
	models.IsActive
	models.ControlBy
	models.ModelTime
}

// TableName 添加表前缀
func (*SysUser) TableName() string {
	return "oms_sys_user"
}

func (u *SysUser) SaveUser() (*SysUser, error) {
	err := global.DB.Create(&u).Error
	if err != nil {
		return &SysUser{}, err
	}
	return u, nil
}

// BeforeCreate 使用gorm的hook在保存密码前对密码进行hash
func (u *SysUser) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string) (string, error) {
	var err error
	u := SysUser{}

	// 邮箱正则表达式（支持大多数常见格式）
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(username) {
		err = global.DB.Model(SysUser{}).Where("email = ?", username).Take(&u).Error
	} else {
		err = global.DB.Model(SysUser{}).Where("user_name = ?", username).Take(&u).Error
	}

	if err != nil {
		return "", err
	}
	//err = VerifyPassword(password+""+u.Salt, u.Password)
	err = VerifyPassword(fmt.Sprintf("%s%s", password, u.Salt), u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.UserId)
	if err != nil {
		return "", err
	}
	return token, nil
}

// PrepareGive 返回前将用户密码置空
func (u *SysUser) PrepareGive() {
	u.Password = ""
}

// GetUserInfoByID 通过ID获取用户信息
func GetUserInfoByID(uid uint) (SysUser, error) {
	var u SysUser
	if err := global.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}

	//u.PrepareGive()
	return u, nil
}

type MenuData struct {
	// column data
	Id                uint   `json:"id"`
	ParentId          uint   `json:"parentId"`
	Icon              string `json:"icon"`
	Sort              int    `json:"sort"`
	Permission        string `json:"permission"`
	ApiPath           string `json:"apiPath"`
	CreateTime        int    `json:"createTime"`
	UpdateTime        int    `json:"updateTime"`
	Show              int    `json:"show"`
	Status            int    `json:"status"`
	MenuTitle         string `json:"menuTitle"`
	CurrentActiveMenu string `json:"currentActiveMenu"`
	// tree data
	Type      uint       `json:"type"`
	MenuName  string     `json:"menuName"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Component string     `json:"component"`
	Redirect  string     `json:"redirect"`
	Meta      MenuMeta   `json:"meta"`
	Children  []MenuData `json:"children"`
}
type MenuMeta struct {
	Title             string `json:"title"`
	Icon              string `json:"icon"`
	HideBreadcrumb    bool   `json:"hideBreadcrumb"`
	IgnoreKeepAlive   bool   `json:"ignoreKeepAlive"`
	HideMenu          bool   `json:"hideMenu"`
	CurrentActiveMenu string `json:"currentActiveMenu"`
}

func GetRoleMenu(roleId uint) []MenuData {
	//var menuData []SysRoleMenu
	//result := DB.Find(&menuData, "role_id = ?", roleId)
	//// 检查错误
	//if result.Error != nil {
	//	panic("查询失败: " + result.Error.Error())
	//}
	//var selectedRoleMenuIds []uint
	//for _, row := range menuData {
	//	// fmt.Println("RoleId:", row.RoleId, "MenuId:", row.MenuId)
	//	selectedRoleMenuIds = append(selectedRoleMenuIds, row.MenuId)
	//}
	var selectedRoleMenuIds []uint
	err := global.DB.Model(&SysRoleMenu{}).Where("role_id = ?", roleId).Pluck("menu_id", &selectedRoleMenuIds)
	if err.Error != nil {
		panic("查询失败: " + err.Error.Error())
	}
	allMyMenuIds := getParentMenuIds(selectedRoleMenuIds)
	var roleMenuData []SysMenu
	err = global.DB.Where(" menu_id in ? and status = ? and is_active = ?", allMyMenuIds, 1, 1).Order("sort desc").Find(&roleMenuData)
	// 检查错误
	if err.Error != nil {
		panic("查询失败: " + err.Error.Error())
	}
	return handleMenuTree(roleMenuData)
}

// getParentMenuIds 根据菜单ID查找出所有的父级菜单ID
func getParentMenuIds(selectedRoleMenuIds []uint) []uint {
	var allMenuIds []uint
	var result []map[string]interface{}
	err := global.DB.Model(&SysMenu{}).Where("menu_id in ?", selectedRoleMenuIds).Select("menu_id", "parent_id").Find(&result)
	if err.Error != nil {
		panic("查询失败: " + err.Error.Error())
	}
	queryMenuIds := make([]uint, len(result))
	var queryParentMenuIds []uint
	for i := 0; i < len(result); i++ {
		queryMenuIds[i] = result[i]["menu_id"].(uint)
		if result[i]["parent_id"].(uint) > 0 {
			queryParentMenuIds = append(queryParentMenuIds, result[i]["parent_id"].(uint))
		}
	}
	allMenuIds = append(allMenuIds, queryMenuIds...)
	if len(queryParentMenuIds) > 0 {
		allMenuIds = append(allMenuIds, getParentMenuIds(queryParentMenuIds)...)
	}
	return allMenuIds
}

// handleMenuTree 将菜单处理成树结构
func handleMenuTree(roleMenuData []SysMenu) []MenuData {
	//var menuData []MenuData
	var menus []MenuData
	menuData := make(map[uint]*MenuData)
	for _, row := range roleMenuData {
		//fmt.Println(row.MenuName)
		if row.MenuType == 2 {
			row.MetaIcon = "ant-design:api-twotone"
		}
		menuData[row.MenuId] = &MenuData{
			Id: row.MenuId, ParentId: row.ParentId, Icon: row.MetaIcon, Name: row.MenuName,
			Type: row.MenuType, MenuName: row.MenuName, Path: row.MenuPath, Component: row.Component, Redirect: row.Redirect,
			Meta: MenuMeta{Title: row.MetaTitle, Icon: row.MetaIcon, HideBreadcrumb: false, IgnoreKeepAlive: false,
				HideMenu: row.HideMenu == 1, CurrentActiveMenu: row.MetaCurrentActiveMenu}}
		//if row.MetaCurrentActiveMenu == "" {
		//	value := menuData[row.MenuId]
		//	// 类型断言为 map[string]interface{}
		//	if m, ok := value.(map[string]interface{}); ok {
		//		delete(m, "MetaCurrentActiveMenu") // 删除 "age" 字段
		//		fmt.Println("更新后的值:", m)           // 输出: map[name:Alice]
		//	} else {
		//		fmt.Println("值不是 map 类型")
		//	}
		//}
		//menuData = append(menuData, MenuData{
		//	Id: row.MenuId, ParentId: row.ParentId, Icon: row.MetaIcon, Name: row.MenuName,
		//	Type: row.MenuType, MenuName: row.MenuName, Path: row.MenuPath, Component: row.Component, Redirect: row.Redirect,
		//	Meta: MenuMeta{Title: row.MetaTitle, Icon: row.MetaIcon, HideBreadcrumb: false, IgnoreKeepAlive: false,
		//		HideMenu: row.HideMenu == 1, CurrentActiveMenu: row.MetaCurrentActiveMenu}})
	}
	jsonBytes, err := json.Marshal(menuData)
	if err != nil {
		return nil
	}
	// 将字节切片转为字符串输出
	jsonStr := string(jsonBytes)
	fmt.Println(jsonStr)
	for menuId, row := range menuData {
		// fmt.Println(menuId, row.MenuName)
		// 检查是否存在
		if _, exists := menuData[row.ParentId]; exists {
			// fmt.Printf("ParentId %d 存在，菜单名称: %s\n", row.ParentId, a.Name)
			menuData[row.ParentId].Children = append(menuData[row.ParentId].Children, *menuData[menuId])
		} else {
			// fmt.Printf("ParentId %d 不存在\n", row.ParentId)
			menus = append(menus, *menuData[menuId])
		}
	}
	return menus
}
