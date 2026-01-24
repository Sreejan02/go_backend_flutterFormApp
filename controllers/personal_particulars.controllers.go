package controllers

import (
	"net/http"
	"time"

	"backend/database"
	"backend/models"
	pp "backend/models/personal_particulars"

	"github.com/gin-gonic/gin"
)

//
// ===============================
// INPUT DTOs
// ===============================
//

type FamilyMemberInput struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
	DateOfBirth  string `json:"dateOfBirth"`
	Occupation   string `json:"occupation"`
}

type EducationInput struct {
	InstituteName string `json:"instituteName"`
	University    string `json:"university"`
	DegreeOrExam  string `json:"degreeOrExam"`

	MainSubjects string `json:"mainSubjects"` // ✅ NEW
	Division     string `json:"division"`     // ✅ NEW

	YearFrom     string `json:"yearFrom"`
	YearTo       string `json:"yearTo"`
	MarksPercent string `json:"marksPercent"`
}


type PastEmploymentInput struct {
	EmployerName     string `json:"employerName"`
	Designation      string `json:"designation"`
	EmployedFrom     string `json:"employedFrom"`
	EmployedTo       string `json:"employedTo"`
	ReasonForLeaving string `json:"reasonForLeaving"`
}

type DependentInput struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
	DateOfBirth  string `json:"dateOfBirth"`
	Reason       string `json:"reason"`
}


type LastThreeEmploymentInput struct {
	DesignationScope           string `json:"designationScope"`
	SupervisorNameDesignation string `json:"supervisorNameDesignation"`
}

type ProfessionalTrainingInput struct {
	Subject  string `json:"subject"`
	Duration string `json:"duration"`
}

type PromotionInput struct {
	PromotionFrom   string `json:"promotionFrom"`
	PromotionAs     string `json:"promotionAs"`
	DateOfPromotion string `json:"dateOfPromotion"`
}

type LanguageKnownInput struct {
	Language string `json:"language"`
	Read     bool   `json:"read"`
	Write    bool   `json:"write"`
	Speak    bool   `json:"speak"`
}

type PersonalParticularsInput struct {
	InviteUID string `json:"inviteUid"`

	PlantLocation  string `json:"plantLocation"`
	ReferenceNo    string `json:"referenceNo"`
	PostAppliedFor string `json:"postAppliedFor"`

	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`

	PresentAddress        string `json:"presentAddress"`
	PresentPhoneResidence string `json:"presentPhoneResidence"`
	Mobile                string `json:"mobile"`
	Email                 string `json:"email"`

	PermanentAddress        string `json:"permanentAddress"`
	PermanentPhoneResidence string `json:"permanentPhoneResidence"`

	EmergencyAddress string `json:"emergencyAddress"`
	EmergencyPhone   string `json:"emergencyPhone"`

	FatherOrHusbandName            string `json:"fatherOrHusbandName"`
	FatherOrHusbandAddress         string `json:"fatherOrHusbandAddress"`
	FatherOrHusbandOccupation      string `json:"fatherOrHusbandOccupation"`
	FatherOrHusbandDesignation     string `json:"fatherOrHusbandDesignation"`
	FatherOrHusbandOfficialAddress string `json:"fatherOrHusbandOfficialAddress"`
	FatherOrHusbandLastOccupation  string `json:"fatherOrHusbandLastOccupation"`

	DateOfBirth   string `json:"dateOfBirth"`
	AgeYears      int    `json:"ageYears"`
	PlaceOfBirth  string `json:"placeOfBirth"`
	PlaceOfOrigin string `json:"placeOfOrigin"`

	MaritalStatus string `json:"maritalStatus"`
	HeightCm      int    `json:"heightCm"`
	WeightKg      int    `json:"weightKg"`

	AppearedForTestOrInterviewEarlier    bool `json:"appearedForTestOrInterviewEarlier"`
	PresentEmployerAwareOfApplication   bool `json:"presentEmployerAwareOfApplication"`
	RelatedToAnyDirector                bool `json:"relatedToAnyDirector"`
	AllowRetainNameOnFileIfUnsuccessful bool `json:"allowRetainNameOnFileIfUnsuccessful"`

	  // =====================
    // PAGE 5 – PRESENT EMPLOYMENT
    // =====================
    PresentEmployerNameAddress         string `json:"presentEmployerNameAddress"`
    PresentEmploymentDateOfAppointment string `json:"presentEmploymentDateOfAppointment"`
    DesignationOnJoining               string `json:"designationOnJoining"`
    PresentDesignation                 string `json:"presentDesignation"`
    PresentPositionInHierarchy         string `json:"presentPositionInHierarchy"`

    // =====================
    // PAGE 6 – DETAILS
    // =====================
    AppearedForTestOrInterviewEarlierDetails string `json:"appearedForTestOrInterviewEarlierDetails"`
    RelatedToAnyDirectorDetails             string `json:"relatedToAnyDirectorDetails"`

    DetailedScopeOfResponsibilitiesPresent  string `json:"detailedScopeOfResponsibilitiesPresent"`
    ImportantAspectsOfExperience            string `json:"importantAspectsOfExperience"`
    ReasonForSeekingNewAppointment          string `json:"reasonForSeekingNewAppointment"`
	NoticePeriodToJoin string `json:"noticePeriodToJoin"`



	HobbiesInterests                    string `json:"hobbiesInterests"`
	ExtraCurricularLiteraryCulturalArts string `json:"extraCurricularLiteraryCulturalArts"`
	ExtraCurricularSocial               string `json:"extraCurricularSocial"`
	CourtProceedingsDetails             string `json:"courtProceedingsDetails"`
	SeriousIllness                      string `json:"seriousIllness"`
	PhysicalDisability                  string `json:"physicalDisability"`

	FamilyMembers        []FamilyMemberInput         `json:"familyMembers"`
	EducationHistory     []EducationInput            `json:"educationHistory"`
	PastEmployment       []PastEmploymentInput       `json:"pastEmployment"`
	Dependents           []DependentInput            `json:"dependents"`
	ProfessionalTraining []ProfessionalTrainingInput `json:"professionalTraining"`
	Promotions           []PromotionInput            `json:"promotions"`
	LanguagesKnown       []LanguageKnownInput        `json:"languagesKnown"`
	LastThreeEmployment  []LastThreeEmploymentInput `json:"lastThreeEmployment"`



}

//
// ===============================
// POST /personal-particulars
// ===============================
//

func CreatePersonalParticulars(c *gin.Context) {
	var input PersonalParticularsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate invite
	var user models.User
	if err := database.DB.Where("invite_uid = ?", input.InviteUID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired invite"})
		return
	}

	// Prevent duplicate
	var count int64
	database.DB.Model(&pp.PersonalParticulars{}).
		Where("invite_uid = ?", input.InviteUID).
		Count(&count)

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Application already submitted"})
		return
	}

	 dob, err := time.Parse("2006-01-02", input.DateOfBirth)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dateOfBirth"})
		return
	}

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	main := pp.PersonalParticulars{
		InviteUID: input.InviteUID,

		PlantLocation:  input.PlantLocation,
		ReferenceNo:    input.ReferenceNo,
		PostAppliedFor: input.PostAppliedFor,

		FirstName:  input.FirstName,
		MiddleName: input.MiddleName,
		LastName:   input.LastName,

		PresentAddress:        input.PresentAddress,
		PresentPhoneResidence: input.PresentPhoneResidence,
		Mobile:                input.Mobile,
		Email:                 input.Email,

		PermanentAddress:        input.PermanentAddress,
		PermanentPhoneResidence: input.PermanentPhoneResidence,

		EmergencyAddress: input.EmergencyAddress,
		EmergencyPhone:   input.EmergencyPhone,

		FatherOrHusbandName:            input.FatherOrHusbandName,
		FatherOrHusbandAddress:         input.FatherOrHusbandAddress,
		FatherOrHusbandOccupation:      input.FatherOrHusbandOccupation,
		FatherOrHusbandDesignation:     input.FatherOrHusbandDesignation,
		FatherOrHusbandOfficialAddress: input.FatherOrHusbandOfficialAddress,
		FatherOrHusbandLastOccupation:  input.FatherOrHusbandLastOccupation,

		DateOfBirth:   dob,
		AgeYears:      input.AgeYears,
		PlaceOfBirth:  input.PlaceOfBirth,
		PlaceOfOrigin: input.PlaceOfOrigin,

		MaritalStatus: input.MaritalStatus,
		HeightCm:      input.HeightCm,
		WeightKg:      input.WeightKg,

		AppearedForTestOrInterviewEarlier:    input.AppearedForTestOrInterviewEarlier,
		PresentEmployerAwareOfApplication:   input.PresentEmployerAwareOfApplication,
		RelatedToAnyDirector:                input.RelatedToAnyDirector,
		AllowRetainNameOnFileIfUnsuccessful: input.AllowRetainNameOnFileIfUnsuccessful,

		// =====================
		// PAGE 5 – PRESENT EMPLOYMENT
		// =====================
		PresentEmployerNameAddress:         input.PresentEmployerNameAddress,
		PresentEmploymentDateOfAppointment: input.PresentEmploymentDateOfAppointment,
		DesignationOnJoining:               input.DesignationOnJoining,
		PresentDesignation:                 input.PresentDesignation,
		PresentPositionInHierarchy:         input.PresentPositionInHierarchy,

		// =====================
		// PAGE 6 – DETAILS
		// =====================
		AppearedForTestOrInterviewEarlierDetails: input.AppearedForTestOrInterviewEarlierDetails,
		RelatedToAnyDirectorDetails:             input.RelatedToAnyDirectorDetails,

		DetailedScopeOfResponsibilitiesPresent:  input.DetailedScopeOfResponsibilitiesPresent,
		ImportantAspectsOfExperience:            input.ImportantAspectsOfExperience,
		ReasonForSeekingNewAppointment:          input.ReasonForSeekingNewAppointment,

		HobbiesInterests:                    input.HobbiesInterests,
		ExtraCurricularLiteraryCulturalArts: input.ExtraCurricularLiteraryCulturalArts,
		ExtraCurricularSocial:               input.ExtraCurricularSocial,
		CourtProceedingsDetails:             input.CourtProceedingsDetails,
		SeriousIllness:                      input.SeriousIllness,
		PhysicalDisability:                  input.PhysicalDisability,

		ApplicationStatus: "submitted",
	}

	if err := tx.Create(&main).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Helper for safe create
	safeCreate := func(obj interface{}) {
		if err := tx.Create(obj).Error; err != nil {
			tx.Rollback()
			panic(err)
		}
	}


	for _, f := range input.FamilyMembers {
		d, err := time.Parse("2006-01-02", f.DateOfBirth)
		if err != nil {
			tx.Rollback()
			c.JSON(400, gin.H{"error": "Invalid family member DOB"})
			return
		}
		safeCreate(&pp.FamilyMember{
			PersonalParticularsID: main.ID,
			Name:                 f.Name,
			Relationship:         f.Relationship,
			DateOfBirth:          d,
			Occupation:           f.Occupation,
		})
	}

	for _, e := range input.EducationHistory {
		safeCreate(&pp.Education{
			PersonalParticularsID: main.ID,
			InstituteName:         e.InstituteName,
			University:            e.University,
			DegreeOrExam:          e.DegreeOrExam,
			MainSubjects:          e.MainSubjects, // ✅ NEW
			Division:              e.Division,     // ✅ NEW
			YearFrom:              e.YearFrom,
			YearTo:                e.YearTo,
			MarksPercent:          e.MarksPercent,
		})
	}



// =====================
// LAST THREE EMPLOYMENT (MAX 3)
// =====================


if len(input.LastThreeEmployment) > 3 {
	tx.Rollback()
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Maximum 3 last three employment records allowed",
	})
	return
}

for _, l := range input.LastThreeEmployment {
	if l.DesignationScope == "" || l.SupervisorNameDesignation == "" {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Designation scope and supervisor details are required",
		})
		return
	}

	if err := tx.Create(&pp.LastThreeEmployment{
		PersonalParticularsID:     main.ID,
		DesignationScope:          l.DesignationScope,
		SupervisorNameDesignation: l.SupervisorNameDesignation,
	}).Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

	





	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Application submitted successfully",
		"id":      main.ID,
	})
}

//
// ===============================
// GET APIs
// ===============================
//

func GetAllPersonalParticulars(c *gin.Context) {
	var list []pp.PersonalParticulars

	if err := database.DB.
		Preload("FamilyMembers").
		Preload("EducationHistory").
		Preload("PastEmployment").
		Preload("Dependents").
		Preload("ProfessionalTraining").
		Preload("Promotions").
		Preload("LanguagesKnown").
		Preload("LastThreeEmployment").

		Find(&list).Error; err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, list)
}

func GetPersonalParticularsByInvite(c *gin.Context) {
	inviteUID := c.Param("inviteUid")

	var data pp.PersonalParticulars
	if err := database.DB.
		Where("invite_uid = ?", inviteUID).
		Preload("FamilyMembers").
		Preload("EducationHistory").
		Preload("PastEmployment").
		Preload("Dependents").
		Preload("ProfessionalTraining").
		Preload("Promotions").
		Preload("LanguagesKnown").
		Preload("LastThreeEmployment").


		First(&data).Error; err != nil {

		c.JSON(404, gin.H{"error": "record not found"})
		return
	}

	c.JSON(200, data)
}

//
// ===============================
// PUT /personal-particulars/:id
// ===============================
//

type PersonalParticularsUpdateInput struct {

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
	DateOfBirth   string `json:"dateOfBirth"`
	AgeYears      *int   `json:"ageYears"`
	PlaceOfBirth  string `json:"placeOfBirth"`
	PlaceOfOrigin string `json:"placeOfOrigin"`

	MaritalStatus string `json:"maritalStatus"`
	HeightCm      *int   `json:"heightCm"`
	WeightKg      *int   `json:"weightKg"`

	// =====================
	// FATHER / HUSBAND
	// =====================
	FatherOrHusbandName            string `json:"fatherOrHusbandName"`
	FatherOrHusbandAddress         string `json:"fatherOrHusbandAddress"`
	FatherOrHusbandOccupation      string `json:"fatherOrHusbandOccupation"`
	FatherOrHusbandDesignation     string `json:"fatherOrHusbandDesignation"`
	FatherOrHusbandOfficialAddress string `json:"fatherOrHusbandOfficialAddress"`
	FatherOrHusbandLastOccupation  string `json:"fatherOrHusbandLastOccupation"`

	// =====================
	// DECLARATIONS (PAGE 6)
	// =====================
	AppearedForTestOrInterviewEarlier    *bool `json:"appearedForTestOrInterviewEarlier"`
	PresentEmployerAwareOfApplication   *bool `json:"presentEmployerAwareOfApplication"`
	RelatedToAnyDirector                *bool `json:"relatedToAnyDirector"`
	AllowRetainNameOnFileIfUnsuccessful *bool `json:"allowRetainNameOnFileIfUnsuccessful"`

	AppearedForTestOrInterviewEarlierDetails string `json:"appearedForTestOrInterviewEarlierDetails"`
	RelatedToAnyDirectorDetails             string `json:"relatedToAnyDirectorDetails"`

	NoticePeriodToJoin string `json:"noticePeriodToJoin"`

	// =====================
	// EXPERIENCE & REASONS (PAGE 6)
	// =====================
	DetailedScopeOfResponsibilitiesPresent string `json:"detailedScopeOfResponsibilitiesPresent"`
	ImportantAspectsOfExperience           string `json:"importantAspectsOfExperience"`
	ReasonForSeekingNewAppointment         string `json:"reasonForSeekingNewAppointment"`

	// =====================
	// PRESENT EMPLOYMENT (PAGE 5)
	// =====================
	PresentEmployerNameAddress         string `json:"presentEmployerNameAddress"`
	PresentEmploymentDateOfAppointment string `json:"presentEmploymentDateOfAppointment"`
	DesignationOnJoining               string `json:"designationOnJoining"`
	PresentDesignation                 string `json:"presentDesignation"`
	PresentPositionInHierarchy         string `json:"presentPositionInHierarchy"`

	// =====================
	// EXTRA / MEDICAL
	// =====================
	HobbiesInterests                    string `json:"hobbiesInterests"`
	ExtraCurricularLiteraryCulturalArts string `json:"extraCurricularLiteraryCulturalArts"`
	ExtraCurricularSocial               string `json:"extraCurricularSocial"`

	CourtProceedingsDetails string `json:"courtProceedingsDetails"`
	SeriousIllness          string `json:"seriousIllness"`
	PhysicalDisability     string `json:"physicalDisability"`

	// =====================
	// CHILD TABLES (FULL REPLACE)
	// =====================
	FamilyMembers        []FamilyMemberInput         `json:"familyMembers"`
	Dependents           []DependentInput            `json:"dependents"`
	EducationHistory     []EducationInput            `json:"educationHistory"`
	PastEmployment       []PastEmploymentInput       `json:"pastEmployment"`
	ProfessionalTraining []ProfessionalTrainingInput `json:"professionalTraining"`
	Promotions           []PromotionInput            `json:"promotions"`
	LanguagesKnown       []LanguageKnownInput        `json:"languagesKnown"`
	LastThreeEmployment  []LastThreeEmploymentInput  `json:"lastThreeEmployment"`
}



func UpdatePersonalParticulars(c *gin.Context) {
	id := c.Param("id")

	var input PersonalParticularsUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}

	if input.Mobile != "" {
		updates["mobile"] = input.Mobile
	}
	if input.Email != "" {
		updates["email"] = input.Email
	}
	if input.PresentAddress != "" {
		updates["present_address"] = input.PresentAddress
	}
	if input.PermanentAddress != "" {
		updates["permanent_address"] = input.PermanentAddress
	}
	if input.HeightCm != nil {
		updates["height_cm"] = *input.HeightCm
	}
	if input.WeightKg != nil {
		updates["weight_kg"] = *input.WeightKg
	}

	if input.PlantLocation != "" {
	updates["plant_location"] = input.PlantLocation
}
if input.ReferenceNo != "" {
	updates["reference_no"] = input.ReferenceNo
}
if input.PostAppliedFor != "" {
	updates["post_applied_for"] = input.PostAppliedFor
}
if input.PresentEmployerNameAddress != "" {
	updates["present_employer_name_address"] = input.PresentEmployerNameAddress
}
if input.PresentDesignation != "" {
	updates["present_designation"] = input.PresentDesignation
}
if input.PresentPositionInHierarchy != "" {
	updates["present_position_in_hierarchy"] = input.PresentPositionInHierarchy
}
if input.DetailedScopeOfResponsibilitiesPresent != "" {
	updates["detailed_scope_of_responsibilities_present"] =
		input.DetailedScopeOfResponsibilitiesPresent
}
if input.ImportantAspectsOfExperience != "" {
	updates["important_aspects_of_experience"] =
		input.ImportantAspectsOfExperience
}
if input.ReasonForSeekingNewAppointment != "" {
	updates["reason_for_seeking_new_appointment"] =
		input.ReasonForSeekingNewAppointment
		
}

if input.FirstName != "" {
	updates["first_name"] = input.FirstName
}
if input.MiddleName != "" {
	updates["middle_name"] = input.MiddleName
}
if input.LastName != "" {
	updates["last_name"] = input.LastName
}

if input.DateOfBirth != "" {
	d, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err == nil {
		updates["date_of_birth"] = d
	}
}

if input.AgeYears != nil {
	updates["age_years"] = *input.AgeYears
}

if input.MaritalStatus != "" {
	updates["marital_status"] = input.MaritalStatus
}
if input.PlaceOfBirth != "" {
	updates["place_of_birth"] = input.PlaceOfBirth
}
if input.PlaceOfOrigin != "" {
	updates["place_of_origin"] = input.PlaceOfOrigin
}



// =====================
// UPDATE LAST THREE EMPLOYMENT
// =====================
if input.LastThreeEmployment != nil {

	// Convert id string → uint
	var main pp.PersonalParticulars
	if err := database.DB.First(&main, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}

	// 1️⃣ Delete old rows
	database.DB.
		Where("personal_particulars_id = ?", main.ID).
		Delete(&pp.LastThreeEmployment{})

	// 2️⃣ Insert new rows (max 3)
	for _, l := range input.LastThreeEmployment {

		if l.DesignationScope == "" && l.SupervisorNameDesignation == "" {
			continue
		}

		database.DB.Create(&pp.LastThreeEmployment{
			PersonalParticularsID:     main.ID,
			DesignationScope:          l.DesignationScope,
			SupervisorNameDesignation: l.SupervisorNameDesignation,
		})
	}
}


	if len(updates) == 0 {
		c.JSON(400, gin.H{"error": "no fields to update"})
		return
	}

	if err := database.DB.
		Model(&pp.PersonalParticulars{}).
		Where("id = ?", id).
		Updates(updates).Error; err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "updated successfully"})
}


