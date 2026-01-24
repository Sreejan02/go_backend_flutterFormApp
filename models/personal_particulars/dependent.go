package personal_particulars

import "time"

type Dependent struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PersonalParticularsID uint `gorm:"index;not null" json:"personalParticularsId"`

	Name         string    `json:"name"`
	Relationship string    `json:"relationship"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	Reason       string    `json:"reason"`
}
