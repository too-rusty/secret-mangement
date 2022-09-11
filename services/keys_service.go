package services

import "gorm.io/gorm"

type KeysService struct {
	db *gorm.DB
}

func NewKeysService(db *gorm.DB) *KeysService {
	return &KeysService{db: db}
}

func (k KeysService) GenerateKeys() {}
