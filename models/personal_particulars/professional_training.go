package personal_particulars


type ProfessionalTraining struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PersonalParticularsID uint `gorm:"index;not null;constraint:OnDelete:CASCADE"`

	Subject  string `json:"subject" gorm:"type:text"`
	Duration string `json:"duration"`
}
