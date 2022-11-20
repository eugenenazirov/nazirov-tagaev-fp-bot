package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	MsgText string `json:"msg_text"`
}

type MessageModel struct {
	Db *gorm.DB
}

func (m *MessageModel) Create(message Message) error {

	result := m.Db.Create(&message)

	return result.Error
}

func (m *MessageModel) FindOne(msgId uint64) (*Message, error) {
	existMessage := Message{}

	result := m.Db.First(&existMessage, msgId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existMessage, nil
}
