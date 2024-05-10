package repository

import (
	"trademarkia/common"

	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	config common.Config
	err    error
)

func init() {
	config, err = common.LoadConfig()
	DB, err = ConnectPsqlDB(&config)
}

func AddUserData(input UserData) (UserData, error) {
	if err := DB.Create(&input).Error; err != nil {
		return UserData{}, err
	} else {
		return UserData{}, nil
	}
}

func EditUserData(input UserData) error {
	var existingUser UserData
	if err := DB.Where("email = ?", input.Email).First(&existingUser).Error; err != nil {
		return err
	}
	if input.FirstName != "" {
		existingUser.FirstName = input.FirstName
	}
	if input.LastName != "" {
		existingUser.LastName = input.LastName
	}
	if input.Age != 0 {
		existingUser.Age = input.Age
	}
	if input.Married != true {
		existingUser.Married = input.Married
	}
	if err := DB.Save(&existingUser).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserData(email string) error {
	if err := DB.Where("email = ?", email).Delete(&UserData{}).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByMail(email string) (UserData, error) {
	var user UserData
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return UserData{}, err
	}
	return user, nil
}

func FindUserIDByMail(email string) (int, error) {
	var userID int
	if err := DB.Model(UserData{}).Where("email = ?", email).Pluck("id", &userID).Error; err != nil {
		return 0, err
	}
	return userID, nil
}
