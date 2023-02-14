package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FullName  string    `gorm:"type:character varying;not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"type:character varying;not null"`
	Salt      string    `gorm:"type:character varying;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUpdate struct {
}
