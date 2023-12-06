package common

import (
	"encoding/base64"
	"final-project-booking-room/config"
	"final-project-booking-room/utils/modelutil"

	"fmt"
	"io"
	"net/smtp"
	"os"
)

type EmailService interface {
	SendEmail(payload modelutil.BodySender) error
	SendEmailFile(payload modelutil.BodySender) error
}

type emailService struct {
	cfg *config.Config
}

func (e *emailService) SendEmail(payload modelutil.BodySender) error {
	message := "From: " + e.cfg.EmailFrom + "\n" +
		"To: " + payload.To[0] + "\n" +
		"Subject: " + payload.Subject + "\n\n" +
		payload.Body

	auth := smtp.PlainAuth("", e.cfg.EmailFrom, e.cfg.EmailConfig.Password, e.cfg.Server)
	err := smtp.SendMail(e.cfg.Server+":"+e.cfg.EmailConfig.Port, auth, e.cfg.EmailFrom, payload.To, []byte(message))
	if err != nil {
		return err
	}

	return nil

}

func (e *emailService) SendEmailFile(payload modelutil.BodySender) error {
	file, err := os.Open(payload.CSVFilePath)
	if err != nil {
		return fmt.Errorf("error opening csv file: %v", err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading CSV file: %v", err)
	}

	encodedContents := base64.StdEncoding.EncodeToString(contents)

	message := "From: " + e.cfg.EmailFrom + "\n" +
		"To: " + payload.To[0] + "\n" +
		"Subject: " + payload.Subject + "\n" +
		"Content-Type: multipart/mixed; boundary=boundarystring\n\n" +
		"--boundarystring\n" +
		"Content-Type: text/plain; charset=UTF-8\n\n" +
		payload.Body + "\n\n" +
		"--boundarystring\n" +
		"Content-Type: application/csv\n" +
		"Content-Transfer-Encoding: base64\n" +
		"Content-Disposition: attachment; filename=\"" + payload.CSVFilePath + "\"\n\n"

	message += encodedContents + "\n\n" +
		"--boundarystring--"

	auth := smtp.PlainAuth("", e.cfg.EmailFrom, e.cfg.EmailConfig.Password, e.cfg.Server)
	err = smtp.SendMail(e.cfg.Server+":"+e.cfg.EmailConfig.Port, auth, e.cfg.EmailFrom, payload.To, []byte(message))
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil

}

func NewEmailService(cfg *config.Config) EmailService {
	return &emailService{cfg: cfg}
}
