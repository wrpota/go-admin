package model

type RoleRule struct {
	ID     int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	RoleId int64 `gorm:"column:role_id" db:"role_id" json:"role_id" form:"role_id"`
	RuleId int64 `gorm:"column:rule_id" db:"rule_id" json:"rule_id" form:"rule_id"`
}

func (RoleRule) TabName() string {
	return "role_rule"
}
