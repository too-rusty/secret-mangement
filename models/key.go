package models

type Key struct {
	ID        int64  `gorm:"primary_key;auto_increment;not_null"`
	Username  string `gorm:"unique;not null"`
	Public    []byte `gorm:"type:bytea;not null"`
	Private   []byte `gorm:"type:bytea;not null"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;"`
	CreatedAt int64  `gorm:"autoUpdateTime:milli;"`
}
