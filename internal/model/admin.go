package model

import "time"

type Admin struct {
	ID        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Username  string    `gorm:"column:username" db:"username" json:"username" form:"username"` //  用户名
	Password  string    `gorm:"column:password" db:"password" json:"password" form:"password"` //  密码
	Tel       string    `gorm:"column:tel" db:"tel" json:"tel" form:"tel"`                     //  手机号
	Email     string    `gorm:"column:email" db:"email" json:"email" form:"email"`             //  邮箱
	Status    int64     `gorm:"column:status" db:"status" json:"status" form:"status"`         //  账号状态
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" db:"updated_at" json:"updated_at" form:"updated_at"`
}

func (Admin) TabName() string {
	return "admin"
}
