package model

type AdminRole struct {
	ID      int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	AdminId int64 `gorm:"column:admin_id" db:"admin_id" json:"admin_id" form:"admin_id"`
	RoleId  int64 `gorm:"column:role_id" db:"role_id" json:"role_id" form:"role_id"`
}

func (AdminRole) TabName() string {
	return "admin_role"
}
