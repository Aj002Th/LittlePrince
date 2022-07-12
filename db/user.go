package db

import "github.com/Aj002Th/LittlePrince/model"

func InsertUser(user model.User) (uint, error) {
	err := db.Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func DeleteUser(id uint) error {
	return db.Where("id = ?", id).Delete(&model.User{}).Error
}

func UpdateUserName(id uint, name string) error {
	return db.Model(&model.User{}).Where("id = ?", id).Update("name", name).Error
}

func UpdateUserPwd(id uint, pwd string) error {
	return db.Model(&model.User{}).Where("id = ?", id).Update("password", pwd).Error
}

func SelectUserNameByID(id uint) (string, error) {
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

func SelectUserPasswordByID(id uint) (string, error) {
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Pwd, nil
}

func SelectUserPasswordByName(name string) (string, error) {
	var user model.User
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Pwd, nil
}

func GetUser(id uint) (model.User, error) {
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetUserByName(name string) (model.User, error) {
	var user model.User
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
