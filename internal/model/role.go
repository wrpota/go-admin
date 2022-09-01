package model

import "time"

type Role struct {
	ID        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name      string    `gorm:"column:name" db:"name" json:"name" form:"name"`         //  用户组名
	Remark    string    `gorm:"column:remark" db:"remark" json:"remark" form:"remark"` //  备注
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" db:"updated_at" json:"updated_at" form:"updated_at"`
}

func (Role) TabName() string {
	return "role"
}
