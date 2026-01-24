package personal_particulars

type LanguageKnown struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PersonalParticularsID uint `gorm:"index;not null" json:"personalParticularsId"`

	Language string `json:"language"`
	Speak    bool   `json:"speak"`
	Read     bool   `json:"read"`
	Write    bool   `json:"write"`
}
