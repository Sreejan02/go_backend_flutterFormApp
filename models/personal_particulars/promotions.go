package personal_particulars


import "time"

type Promotion struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PersonalParticularsID uint `gorm:"index;not null" json:"personalParticularsId"`

	PromotionFrom string    `json:"promotionFrom"`
	PromotionAs   string    `json:"promotionAs"`
	DateOfPromotion time.Time `json:"dateOfPromotion"`
}
