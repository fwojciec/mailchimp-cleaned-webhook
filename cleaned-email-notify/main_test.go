package main

import (
	"context"
	"net/url"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fwojciec/mailchimp-cleaned-webhook/ses"
	"github.com/matryer/is"
)

func TestCleanedEmailNotifyHandler(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	var pBody string
	mockSes := &ses.SenderMock{
		SendFunc: func(toEmail string, body string, subject string) error {
			pBody = body
			return nil
		},
	}
	h := handler{mockSes}
	params := url.Values{}
	params.Add("type", "cleaned")
	params.Add("data[email]", "test@email.com")
	params.Add("fired_at", "2009-03-26 22:01:00")
	params.Add("data[reason]", "hard")
	err := h.Handler(context.Background(), events.APIGatewayProxyRequest{Body: params.Encode()})
	is.NoErr(err)                                                                           // expected no error
	is.Equal(pBody, "Cleaned email: test@email.com\nAt: 2009-03-26 22:01:00\nReason: hard") // expected different body
}
