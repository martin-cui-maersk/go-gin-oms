package models

import "time"

type IsActive struct {
	IsActive uint `json:"isActive" gorm:"index;comment:数据是否有效,1-有效|2-删除"`
}

type ControlBy struct {
	CreateBy uint `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy uint `json:"updateBy" gorm:"index;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy uint) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy uint) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreateAt int64 `json:"-" gorm:"comment:创建时间"`
	UpdateAt int64 `json:"-" gorm:"comment:最后更新时间"`
	//CreateAt int64 `gorm:"column:create_at;comment:创建时间" json:"createAt"`
	//UpdateAt int64 `gorm:"column:update_at;comment:最后更新时间" json:"updateAt"`
	// 不映射到数据库的字段
	FormattedCreatedAt string `gorm:"-" json:"createTime"` // 格式化后的时间字符串
	FormattedUpdatedAt string `gorm:"-" json:"updateTime"` // 格式化后的时间字符串
}

// FormatUnixTime 格式时间
func FormatUnixTime(t int64) string {
	if t == 0 {
		return ""
	}
	// 毫秒级时间戳除以1000
	if Int64Length(t) == 13 {
		t = t / 1000
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

// Int64Length 判断int长度，13位的则位毫秒级时间戳
func Int64Length(n int64) int {
	if n == 0 {
		return 1 // 0 算作 1 位数
	}

	if n < 0 {
		n = -n // 处理负数
	}

	length := 0
	for n > 0 {
		length++
		n /= 10
	}
	return length
}
