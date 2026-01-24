package personal_particulars

type Education struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PersonalParticularsID uint `gorm:"index;not null" json:"personalParticularsId"`

	InstituteName string `json:"instituteName"`
	University    string `json:"university"`
	DegreeOrExam  string `json:"degreeOrExam"`

	MainSubjects string `json:"mainSubjects"`   // ✅ NEW
	Division     string `json:"division"`       // ✅ NEW (First / Second / Third)

	YearFrom     string `json:"yearFrom"`
	YearTo       string `json:"yearTo"`
	MarksPercent string `json:"marksPercent"`
}
