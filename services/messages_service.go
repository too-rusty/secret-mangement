package services

import "gorm.io/gorm"

type MessagesService struct {
	db *gorm.DB
}

func NewMessagesService(db *gorm.DB) *MessagesService {
	return &MessagesService{db: db}
}

func (MessagesService) SendMessage() {}
