package models

type User struct {
	Auditable
	Email    string `gorm:"uniqueIndex"`
	Username string `json:"username" gorm:"unique"`
	Password string
	Roles    []Role `gorm:"many2many:user_roles;"`
}
