package personal_particulars

import "time"

type PersonalParticulars struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// =====================
	// LINK TO USER (ADMIN CREATED)
	// =====================
	InviteUID string `json:"inviteUid"`


	// =====================
	// HEADER
	// =====================
	PlantLocation  string `json:"plantLocation"`
	ReferenceNo    string `json:"referenceNo"`
	PostAppliedFor string `json:"postAppliedFor"`

	// =====================
	// NAME
	// =====================
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`

	// =====================
	// CONTACT
	// =====================
	PresentAddress        string `json:"presentAddress"`
	PresentPhoneResidence string `json:"presentPhoneResidence"`
	Mobile                string `json:"mobile"`
	Email                 string `json:"email"`

	PermanentAddress        string `json:"permanentAddress"`
	PermanentPhoneResidence string `json:"permanentPhoneResidence"`

	EmergencyAddress string `json:"emergencyAddress"`
	EmergencyPhone   string `json:"emergencyPhone"`

	// =====================
	// PERSONAL INFO
	// =====================
	FatherOrHusbandName            string `json:"fatherOrHusbandName"`
	FatherOrHusbandAddress         string `json:"fatherOrHusbandAddress"`
	FatherOrHusbandOccupation      string `json:"fatherOrHusbandOccupation"`
	FatherOrHusbandDesignation     string `json:"fatherOrHusbandDesignation"`
	FatherOrHusbandOfficialAddress string `json:"fatherOrHusbandOfficialAddress"`
	FatherOrHusbandLastOccupation  string `json:"fatherOrHusbandLastOccupation"`

	DateOfBirth   time.Time `json:"dateOfBirth"`
	AgeYears      int       `json:"ageYears"`
	PlaceOfBirth  string    `json:"placeOfBirth"`
	PlaceOfOrigin string    `json:"placeOfOrigin"`

	MaritalStatus string `json:"maritalStatus"`
	HeightCm      int    `json:"heightCm"`
	WeightKg      int    `json:"weightKg"`


	// =====================
	// FLAG DETAILS (PAGE 6)
	// =====================
	AppearedForTestOrInterviewEarlierDetails string `json:"appearedForTestOrInterviewEarlierDetails" gorm:"type:text"`


	DetailedScopeOfResponsibilitiesPresent string `json:"detailedScopeOfResponsibilitiesPresent" gorm:"type:text"`
	RelatedToAnyDirectorDetails string `json:"relatedToAnyDirectorDetails" gorm:"type:text"`
	AppearedForTestOrInterviewEarlier    bool `json:"appearedForTestOrInterviewEarlier"`
	PresentEmployerAwareOfApplication   bool `json:"presentEmployerAwareOfApplication"`
	RelatedToAnyDirector                bool `json:"relatedToAnyDirector"`
	AllowRetainNameOnFileIfUnsuccessful bool `json:"allowRetainNameOnFileIfUnsuccessful"`
	NoticePeriodToJoin                  string `json:"noticePeriodToJoin"`


	// =====================
	// EXPERIENCE & REASONS (PAGE 6)
	// =====================
	
	ImportantAspectsOfExperience           string `json:"importantAspectsOfExperience" gorm:"type:text"`
	ReasonForSeekingNewAppointment         string `json:"reasonForSeekingNewAppointment" gorm:"type:text"`


	// =====================
	// PRESENT EMPLOYMENT (PAGE 5)
	// =====================
	PresentEmployerNameAddress        string `json:"presentEmployerNameAddress" gorm:"type:text"`
	PresentEmploymentDateOfAppointment string `json:"presentEmploymentDateOfAppointment"`
	DesignationOnJoining              string `json:"designationOnJoining"`
	PresentDesignation                string `json:"presentDesignation"`
	PresentPositionInHierarchy        string `json:"presentPositionInHierarchy"`



	// =====================
	// EXTRA / MEDICAL
	// =====================
	HobbiesInterests        string `json:"hobbiesInterests"`
	ExtraCurricularLiteraryCulturalArts string `json:"extraCurricularLiteraryCulturalArts"`
	ExtraCurricularSocial              string `json:"extraCurricularSocial"`
	
	CourtProceedingsDetails string `json:"courtProceedingsDetails"`
	SeriousIllness          string `json:"seriousIllness"`
	PhysicalDisability     string `json:"physicalDisability"`

	// =====================
	// STATUS
	// =====================
	ApplicationStatus string `gorm:"default:submitted;index" json:"applicationStatus"`

	// =====================
	// CHILD TABLE RELATIONS (ARRAYS)
	// =====================
	// =====================
// CHILD TABLE RELATIONS (ARRAYS)
// =====================
	FamilyMembers        []FamilyMember        `gorm:"foreignKey:PersonalParticularsID" json:"familyMembers"`
	EducationHistory     []Education           `gorm:"foreignKey:PersonalParticularsID" json:"educationHistory"`
	PastEmployment       []PastEmployment      `gorm:"foreignKey:PersonalParticularsID" json:"pastEmployment"`

	Dependents           []Dependent           `gorm:"foreignKey:PersonalParticularsID" json:"dependents"`
	ProfessionalTraining []ProfessionalTraining`gorm:"foreignKey:PersonalParticularsID" json:"professionalTraining"`
	Promotions           []Promotion           `gorm:"foreignKey:PersonalParticularsID" json:"promotions"`
	LanguagesKnown       []LanguageKnown       `gorm:"foreignKey:PersonalParticularsID" json:"languagesKnown"`
	LastThreeEmployment  []LastThreeEmployment  `gorm:"foreignKey:PersonalParticularsID" json:"lastThreeEmployment"`
}
