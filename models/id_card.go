package models

import "time"

type IDCard struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Relation
	UserID uint `gorm:"index;not null" json:"userId"`

	// Printed details
	Name string `gorm:"not null" json:"name"`

	BloodGroup string `gorm:"default:UNKNOWN" json:"bloodGroup"`
	// Possible values enforced at app-level (same as Mongo enum)

	ResidenceAddress string `gorm:"not null" json:"residenceAddress"`

	EmergencyContactNo string `gorm:"not null" json:"emergencyContactNo"`

	// File references (Mongo GridFS → nullable FK / string ref)
	SignatureFileID *string `json:"signatureFileId"`
	PhotoFileID     *string `json:"photoFileId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
