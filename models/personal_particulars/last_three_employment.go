package personal_particulars

type LastThreeEmployment struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PersonalParticularsID uint `gorm:"index;not null" json:"personalParticularsId"`

	DesignationScope           string `json:"designationScope" gorm:"type:text"`
	SupervisorNameDesignation string `json:"supervisorNameDesignation" gorm:"type:text"`
}
