package utils

import (
	"bytes"
	"errors"
	"path/filepath"

	"backend/database"
	pp "backend/models/personal_particulars"

	"github.com/jung-kurt/gofpdf"
	"github.com/phpdave11/gofpdi"
)

func GeneratePersonalParticularsPDF(inviteUid string) ([]byte, error) {

	// ================= FETCH DATA =================
	var applicant pp.PersonalParticulars

	if err := database.DB.
		Where("invite_uid = ?", inviteUid).
		First(&applicant).Error; err != nil {
		return nil, errors.New("applicant not found")
	}

	// ================= PDF INIT =================
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// ================= TEMPLATE =================
	importer := gofpdi.NewImporter()

	templatePath := filepath.Join(".", "assets", "PERSONAL_PARTICULARS_FORM.PDF")

	// DEBUG (keep for now)
	absPath, _ := filepath.Abs(templatePath)
	println("USING PDF TEMPLATE:", absPath)

	importer.SetSourceFile(templatePath)

	tpl := importer.ImportPage(1, "/MediaBox")
	importer.UseTemplate(tpl, 0, 0, 210, 297)

	// ================= FONT =================
	pdf.SetFont("Helvetica", "", 9)
	pdf.SetTextColor(0, 0, 0)

	// ================= TEST OVERLAY =================
	pdf.Text(25, 30, applicant.FirstName)
	pdf.Text(80, 30, applicant.MiddleName)
	pdf.Text(130, 30, applicant.LastName)

	pdf.Text(25, 40, applicant.Mobile)
	pdf.Text(80, 40, applicant.Email)

	// ================= OUTPUT =================
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
