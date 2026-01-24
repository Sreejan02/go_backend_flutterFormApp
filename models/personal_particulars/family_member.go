package personal_particulars
import "time"

type FamilyMember struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PersonalParticularsID uint `gorm:"index;not null" json:"personalParticularsId"`

	Name         string    `json:"name"`
	Relationship string    `json:"relationship"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	Occupation   string    `json:"occupation"`
}
