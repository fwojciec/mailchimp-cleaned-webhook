package ses

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/matryer/is"
	"github.com/pkg/errors"
)

type MockSESApi struct {
	sesiface.SESAPI
	Err bool
}

func (s MockSESApi) SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	if s.Err {
		return nil, errors.New("test error")
	}
	return nil, nil
}

func TestNewEmailSvc(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	fromEmail := "from@email.com"
	res1 := NewEmailSvc(fromEmail)
	is.Equal(fromEmail, res1.fromEmail) // expected fromEmail to match
}

func TestSendHappyPath(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	ms := EmailSvc{Svc: MockSESApi{Err: false}}
	err := ms.Send("", "", "")
	is.NoErr(err) // expected no error
}

func TestSendError(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	ms := EmailSvc{Svc: MockSESApi{Err: true}}
	err := ms.Send("", "", "")
	is.True(err != nil) // expected an error
}

func TestCreateInput(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	res := createInput("to@email.com", "body", "subject", "from@email.com")
	exp := ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{aws.String("to@email.com")},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("body")},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("subject"),
			},
		},
		Source: aws.String("from@email.com"),
	}
	is.True(reflect.DeepEqual(*res, exp))
}
