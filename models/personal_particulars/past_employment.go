package personal_particulars

import "time"

type PastEmployment struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PersonalParticularsID uint `gorm:"index;not null" json:"personalParticularsId"`

	EmployerName     string    `json:"employerName"`
	Designation      string    `json:"designation"`
	EmployedFrom     time.Time `json:"employedFrom"`
	EmployedTo       time.Time `json:"employedTo"`
	ReasonForLeaving string    `json:"reasonForLeaving"`
}
