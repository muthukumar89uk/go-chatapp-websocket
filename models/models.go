package models

type User struct {
	Id        string `json:"user_Id" gorm:"unique"`
	User_Name string `json:"user_Name" gorm:"unique"`
}

// func (User) TableName() string {
// 	return "check"
// }
