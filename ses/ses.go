package ses

//go:generate moq -out ses_mock.go . Sender

import (
	"log"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

// Sender represents email mailing functionaliry
type Sender interface {
	Send(toEmail, body, subject string) error
}

// EmailSvc represents AWS SES email message
type EmailSvc struct {
	fromEmail string
	Svc       sesiface.SESAPI
}

// Send sends an email using SES service
func (m *EmailSvc) Send(toEmail, body, subject string) error {
	input := createInput(toEmail, body, subject, m.fromEmail)
	_, err := m.Svc.SendEmail(input)
	if err != nil {
		return err
	}
	return nil
}

// NewEmailSvc returns a Sender implementation with a connected AWS SES client
func NewEmailSvc(fromEmail string) *EmailSvc {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(fromEmail) {
		log.Fatal("from email is invalid")
	}
	sess := session.Must(session.NewSession())
	svc := ses.New(sess)
	return &EmailSvc{
		fromEmail: fromEmail,
		Svc:       svc,
	}
}

func createInput(toEmail, body, subject, fromEmail string) *ses.SendEmailInput {
	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{aws.String(toEmail)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body)},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(fromEmail),
	}
}
