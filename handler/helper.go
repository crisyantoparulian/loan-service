package handler

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"time"

	apperror "github.com/crisyantoparulian/loansvc/utils/app_error"
	httphelper "github.com/crisyantoparulian/loansvc/utils/http_helper"
	"github.com/crisyantoparulian/loansvc/utils/str"
	utilvalidator "github.com/crisyantoparulian/loansvc/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/jordan-wright/email"
	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo/v4"
)

func generateLoanCode(borrowerType string) string {
	code := "RGL" // default regular
	switch borrowerType {
	case "premium":
		code = "PRM"
	case "ultimate":
		code = "ULT"
	}
	datePart := time.Now().Format("060102")   // YYMMDD format
	randomPart := str.GenerateRandomString(5) // Generate 5-character random string
	return fmt.Sprintf("%s%s%s", code, datePart, randomPart)
}

func (s *Server) bindAndValidatePayload(c echo.Context, payload interface{}) error {
	if err := c.Bind(&payload); err != nil {
		httphelper.HttpRespError(c,
			apperror.WrapWithCode(fmt.Errorf("failed to unmarshall request: %w", err),
				http.StatusBadRequest))
		return err
	}

	// Validate payload
	if err := s.Validator.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		validationErr := utilvalidator.FormatValidationErrors(validationErrors)
		c.JSON(http.StatusBadRequest, httphelper.ErrorResponse{
			Message: "Validation failed",
			Errors:  &validationErr,
		})
		return err
	}

	return nil
}

func (s *Server) sendMail(filename, recepientMail string) error {
	// SMTP Config
	smtpHost := s.Config.Mail.SMTPHost
	smtpPort := s.Config.Mail.SMTPPort
	senderEmail := s.Config.Mail.SMTPSender
	password := s.Config.Mail.SMTPPass
	recipient := recepientMail

	// Create Email
	e := email.NewEmail()
	e.From = "Loand Svc <" + senderEmail + ">"
	e.To = []string{recipient}
	e.Subject = "Agreement Document"
	e.Text = []byte("Dear User,\n\nPlease find attached the agreement document.\n\nBest regards,\nYour Company")

	// TODO: Attach PDF
	_, err := e.AttachFile(filename)
	if err != nil {
		return fmt.Errorf("failed to attach file: %v", err)
	}

	// Send Email
	auth := smtp.PlainAuth("", senderEmail, password, smtpHost)
	err = e.Send(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	// Remove the generated PDF after sending
	if err := os.Remove(filename); err != nil {
		fmt.Printf("Warning: Failed to remove file %s: %v\n", filename, err)
	} else {
		fmt.Println("PDF removed successfully.")
	}

	fmt.Println("Email sent successfully!")
	return nil
}

// Generate a simple agreement PDF
func generateAggrementPDF(name, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Agreement Document")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(190, 10, fmt.Sprintf("This is mail loan agreement %s.\n\nSuccessull invested on loan.", name), "", "L", false)

	return pdf.OutputFileAndClose(filename)
}
