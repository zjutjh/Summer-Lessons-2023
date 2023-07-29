package userController

import (
	"CMS/app/models"
	"CMS/app/services/userService"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	PhoneNum string `json:"phone_num" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 登录
func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 判断用户是否存在
	err = userService.CheckUserExistByPhoneNum(data.PhoneNum)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 获取用户信息
	var user *models.User
	user, err = userService.GetUserByPhoneNum(data.PhoneNum)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 判断密码是否正确
	flag := userService.ComparePwd(data.Password, user.Password)
	if !flag {
		utils.JsonErrorResponse(c, 200503, "密码错误")
		return
	}

	// 返回用户信息
	utils.JsonSuccessResponse(c, user)
}
