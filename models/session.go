package models

import "time"

type Session struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relation
	UserID uint `gorm:"index;not null" json:"userId"`

	// Session security
	TokenID    string `gorm:"uniqueIndex;not null" json:"tokenId"`
	SecretHash string `gorm:"not null" json:"-"`

	UserAgent string `json:"userAgent"`
	IP        string `json:"ip"`

	Active bool `gorm:"default:true" json:"active"`

	LastUsedAt time.Time `json:"lastUsedAt"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
