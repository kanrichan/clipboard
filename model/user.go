package model

import "github.com/jinzhu/gorm"

type User struct {
	ID       int64  `gorm:"column:id;primary_key" json:"id"`
	Username string `gorm:"column:username;unique" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

func NewUser(username, password string) (int64, error) {
	var user = &User{
		Username: username,
		Password: password,
	}
	err := DB.Create(&user).Error
	return user.ID, err
}

func GetUserByID(id int64) (User, error) {
	var user User
	return user, DB.First(&user, "id = ? ", id).Error
}

func UpdUserByID(id int64, username, password string) error {
	var user = User{
		ID:       id,
		Username: username,
		Password: password,
	}
	return DB.Model(&user).Where("id = ? ", user.ID).Update(user).Error
}

func DelUserByID(id int64) error {
	var user = &User{ID: id}
	err := DB.Delete(&user).Error
	return err
}

func VerifyUserByUsernameAndPassword(username, password string) (int64, error) {
	var user User
	err := DB.First(&user, "username = ? AND password = ?", username, password).Error
	if err == gorm.ErrRecordNotFound {
		return -1, nil
	}
	return user.ID, err
}
