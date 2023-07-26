package userService

import (
	"CMS/app/models"
	"CMS/config/database"
)

func CheckUserExistByPhoneNum(phoneNum string) error {
	result := database.DB.Where("phone_num = ?", phoneNum).First(&models.User{})
	return result.Error
}

func GetUserByPhoneNum(phoneNum string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("phone_num = ?", phoneNum).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}
