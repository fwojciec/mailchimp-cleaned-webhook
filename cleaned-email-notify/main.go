package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fwojciec/email"
)

/*
Documentation: https://developer.mailchimp.com/documentation/mailchimp/guides/about-webhooks/
Format of Mailchimp's cleaned email webhook POST request:

	"type": "cleaned",
	"fired_at": "2009-03-26 22:01:00",
	"data[list_id]": "a6b5da1054",
	"data[campaign_id]": "4fjk2ma9xd",
	"data[reason]": "hard",
	"data[email]": "api+cleaned@mailchimp.com"

*/

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// parse the request data
	// serverless encodes application/x-www-form-urlencoded POST request data
	// as a query string (key1=val1&key2=val2) to the body of the request
	values, err := url.ParseQuery(event.Body)
	if err != nil {
		return nil, err
	}

	// get the relevant valuse
	dFired := values.Get("fired_at")
	dReason := values.Get("data[reason]")
	dEmail := values.Get("data[email]")

	// prepare email
	from := os.Getenv("FROM_EMAIL")
	to := os.Getenv("TO_EMAIL")
	subject := "Mailchimp Webhook: Cleaned Email Notification"
	body := fmt.Sprintf("Cleaned email: %s\nAt: %s\nReason: %s", dEmail, dFired, dReason)

	// send email
	err = email.New(
		email.From(from),
		email.To(to),
		email.Subject(subject),
		email.TextBody(body),
	).Send()
	if err != nil {
		return nil, err
	}

	// all good
	return &events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
