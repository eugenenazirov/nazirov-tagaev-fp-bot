package models

import "gorm.io/gorm"

// Message struct declares DB table "messages" in code
type Message struct {
	gorm.Model
	MsgText string `json:"msg_text"`
}

// MessageModel links model to DB
type MessageModel struct {
	Db *gorm.DB
}

// Create creates new message
// It is similar to "INSERT" query
func (m *MessageModel) Create(message Message) error {

	result := m.Db.Create(&message)

	return result.Error
}

// FindOne method finds message by message id
// returns link for Message struct (the found message)
// It is similar to "SELECT * FROM messages WHERE msg_id = ... LIMIT 1"
func (m *MessageModel) FindOne(msgId uint64) (*Message, error) {
	existMessage := Message{}

	result := m.Db.First(&existMessage, msgId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existMessage, nil
}
