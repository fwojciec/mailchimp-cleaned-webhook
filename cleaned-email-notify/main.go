package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fwojciec/mailchimp-cleaned-webhook/ses"
)

type handler struct {
	ses.Sender
}

func (h *handler) Handler(ctx context.Context, event events.APIGatewayProxyRequest) error {
	// parse the request data
	values, err := url.ParseQuery(event.Body)
	if err != nil {
		return err
	}
	// get the relevant valuse
	dFired := values.Get("fired_at")
	dReason := values.Get("data[reason]")
	dEmail := values.Get("data[email]")
	// prepare the email
	toEmail := os.Getenv("TO_EMAIL")
	body := fmt.Sprintf("Cleaned email: %s\nAt: %s\nReason: %s", dEmail, dFired, dReason)
	subject := "Mailchimp Webhook: Cleaned Email Notification"
	if err := h.Send(toEmail, body, subject); err != nil {
		return err
	}
	// all went well
	return nil
}

func main() {
	em := ses.NewEmailSvc(os.Getenv("FROM_EMAIL"))
	h := handler{em}
	lambda.Start(h.Handler)
}
