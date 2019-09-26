package models

type User struct {
	Uid         int64	`json:"uid",gorm:"primary_key;auto_increment"`
	Username    string	`json:"username"`
	Password    string	`json:"-"`
	Nickname    string	`json:"nickname"`
	Role        string	`json:"role"`
	Ischeck     string	`json:"ischeck"`
	Photo       string	`json:"photo"`
	Email       string	`json:"email"`
	Description string	`json:"description"`
}
