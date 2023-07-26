package contactController

import (
	"CMS/app/models"
	"CMS/app/services/contactService"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateContactData struct {
	OwnerID   uint   `json:"owner_id" binding:"required"`
	StudentID string `json:"student_id"`
	Name      string `json:"name" binding:"required"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num" binding:"required"`
	Major     string `json:"major"`
	Blacklist bool   `json:"blacklist" binding:"required"`
}

// 添加联系人
func CreateContact(c *gin.Context) {
	var data CreateContactData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	err = contactService.CreateContact(models.Contact{
		OwnerID:   data.OwnerID,
		StudentID: data.StudentID,
		Name:      data.Name,
		Sex:       data.Sex,
		PhoneNum:  data.PhoneNum,
		Major:     data.Major,
		Blacklist: data.Blacklist,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type UpdateContactData struct {
	ID        uint   `json:"id" binding:"required"`
	StudentID string `json:"student_id"`
	Name      string `json:"name" binding:"required"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num" binding:"required"`
	Major     string `json:"major"`
	Blacklist bool   `json:"blacklist" binding:"required"`
}

// 更新联系人信息
func UpdateContact(c *gin.Context) {
	var data UpdateContactData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	err = contactService.UpdateContact(models.Contact{
		ID:        data.ID,
		StudentID: data.StudentID,
		Name:      data.Name,
		Sex:       data.Sex,
		PhoneNum:  data.PhoneNum,
		Major:     data.Major,
		Blacklist: data.Blacklist,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type DeleteContactData struct {
	ID uint `json:"id" binding:"required"`
}

// 删除联系人
func DeleteContact(c *gin.Context) {
	var data DeleteContactData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	err = contactService.DeleteContact(data.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type GetContactData struct {
	OwnerID uint `json:"owner_id" binding:"required"`
}

// 获取联系人列表
func GetContact(c *gin.Context) {
	var data GetContactData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	var contactList []models.Contact
	contactList, err = contactService.GetContactList(data.OwnerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "联系人列表为空")
			return
		}
	}

	utils.JsonSuccessResponse(c, gin.H{
		"contact_list": contactList,
		"num":          len(contactList),
	})
}
