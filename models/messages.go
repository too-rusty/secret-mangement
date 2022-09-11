package models

type Message struct {
	ID        int64  `gorm:"primary_key;auto_increment;not_null"`
	From      string `gorm:"unique;not null"`
	To        string `gorm:"unique;not null"`
	CreatedAt int64  `gorm:"autoUpdateTime:milli;"`
}
