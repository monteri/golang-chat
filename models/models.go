package models

import _ "gorm.io/gorm"

type UserRoom struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex:user_room_idx;not null"`
	Room     string `gorm:"uniqueIndex:user_room_idx;not null"`
}
