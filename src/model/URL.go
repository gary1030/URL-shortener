package model

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	gorm.Model
	ID              int64     `gorm:"primary_key;auto_increment" json:"id"`
	Original_url    string    `gorm:"not null" json:"original_url"`
	Expired_date    time.Time `gorm:"not null" json:"time"`
}
