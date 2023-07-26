package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Sex      string `json:"sex"`
	PhoneNum string `json:"phone_num"`
	Major    string `json:"major"`
	Password string `json:"-"`
}
