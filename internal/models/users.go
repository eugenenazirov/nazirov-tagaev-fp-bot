package models

import (
	"gorm.io/gorm"
	"strconv"
)

// User struct declares DB table "users" in code
type User struct {
	gorm.Model
	Name       string `json:"name"`
	TelegramId int64  `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ChatId     int64  `json:"chat_id"`
}

// Recipient method is required for using in bot.SendMailing()
func (u User) Recipient() string {
	return strconv.FormatInt(u.ChatId, 10)
}

// UserModel links model to DB
type UserModel struct {
	Db *gorm.DB
}

// Create creates new user
// It is similar to "INSERT" query
func (m *UserModel) Create(user User) error {

	result := m.Db.Create(&user)

	return result.Error
}

// FindOne method finds user by telegram id
// returns link for User struct (the found user)
// It is similar to "SELECT * FROM users WHERE telegram_id = ... LIMIT 1"
func (m *UserModel) FindOne(telegramId int64) (*User, error) {
	existUser := User{}

	result := m.Db.First(&existUser, User{TelegramId: telegramId})

	if result.Error != nil {
		return nil, result.Error
	}

	return &existUser, nil
}

// FindAll finds all subscribed users
// returns slice of User structs
// It is similar to "SELECT * FROM users"
func (m *UserModel) FindAll() ([]User, error) {
	var existUsers []User

	result := m.Db.Find(&existUsers, User{})

	if result.Error != nil {
		return nil, result.Error
	}

	return existUsers, nil
}
