package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Phone string `gorm:"uniqueIndex;not null" json:"phone"`

	Role string `gorm:"default:user;not null" json:"role"`

	PostAppliedFor string `json:"postAppliedFor"`

	// 🔑 Invite link identifier
	InviteUID *string `gorm:"uniqueIndex" json:"inviteUid"`

	CreatedAt time.Time `json:"createdAt"`
}
