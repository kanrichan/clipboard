package model

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
