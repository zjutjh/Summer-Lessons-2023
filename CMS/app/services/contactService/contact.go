package contactService

import (
	"CMS/app/models"
	"CMS/config/database"
)

func CreateContact(contact models.Contact) error {
	result := database.DB.Create(&contact)
	return result.Error
}

func UpdateContact(contact models.Contact) error {
	result := database.DB.Save(&contact)
	return result.Error
}

func DeleteContact(id uint) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Contact{})
	return result.Error
}

func GetContactList(ownerID uint) ([]models.Contact, error) {
	result := database.DB.Where("owner_id = ?", ownerID).First(&models.Contact{})
	if result.Error != nil {
		return nil, result.Error
	}
	var contactList []models.Contact
	result = database.DB.Where("owner_id = ?", ownerID).Find(&contactList)
	if result.Error != nil {
		return nil, result.Error
	}
	return contactList, nil
}
