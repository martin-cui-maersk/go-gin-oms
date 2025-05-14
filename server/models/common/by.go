package models

type IsActive struct {
	IsActive int `json:"isActive" gorm:"index;comment:数据是否有效,1-有效|2-删除"`
}

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt int `json:"createAt" gorm:"comment:创建时间"`
	UpdatedAt int `json:"updateAt" gorm:"comment:最后更新时间"`
}
