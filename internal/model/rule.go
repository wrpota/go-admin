package model

type Rule struct {
	ID         int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name       string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Path       string `gorm:"column:path" db:"path" json:"path" form:"path"`             //  路径
	Code       string `gorm:"column:code" db:"code" json:"code" form:"code"`             //  权限标示码
	IsShow     int64  `gorm:"column:is_show" db:"is_show" json:"is_show" form:"is_show"` //  是否显示 0否 1是
	Pid        int64  `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`                 //  父级id
	Type       int64  `gorm:"column:type" db:"type" json:"type" form:"type"`             //  1菜单  0按钮
	Sort       int64  `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`             //  排序
	Icon       string `gorm:"column:icon" db:"icon" json:"icon" form:"icon"`
	Cache      int64  `gorm:"column:cache" db:"cache" json:"cache" form:"cache"`
	Title      string `gorm:"column:title" db:"title" json:"title" form:"title"`
	Permission string `gorm:"column:permission" db:"permission" json:"permission" form:"permission"` //  权限别名
}

func (Rule) TabName() string {
	return "rule"
}
