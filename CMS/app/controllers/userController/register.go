package usercontroller

import (
	"CMS/app/models"
	"CMS/app/services/userService"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterData struct {
	Username        string `json:"username" binding:"required"`
	Sex             string `json:"sex" binding:"required"`
	PhoneNum        string `json:"phone_num" binding:"required"`
	Major           string `json:"major" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// 注册
func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 判断手机号是否已经注册
	err = userService.CheckUserExistByPhoneNum(data.PhoneNum)
	if err == nil {
		utils.JsonErrorResponse(c, 200504, "手机号已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 判断密码是否一致
	flag := userService.ComparePwd(data.Password, data.ConfirmPassword)
	if !flag {
		utils.JsonErrorResponse(c, 200505, "密码不一致")
		return
	}

	// 注册用户
	err = userService.Register(models.User{
		Username: data.Username,
		Sex:      data.Sex,
		PhoneNum: data.PhoneNum,
		Major:    data.Major,
		Password: data.Password,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
