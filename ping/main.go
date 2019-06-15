package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Mailchimp verifies webhook url by sending a GET request and looking for 200 response
func Handler(ctx context.Context) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "OK",
		Headers:         map[string]string{"Content-Type": "text/plain"},
	}
	return &resp, nil
}

func main() {
	lambda.Start(Handler)
}
