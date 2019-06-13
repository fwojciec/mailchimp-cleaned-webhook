package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type handler struct{}

func (h *handler) Handler(ctx context.Context) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "OK",
		Headers:         map[string]string{"Content-Type": "text/plain"},
	}
	return &resp, nil
}

func main() {
	h := handler{}
	lambda.Start(h.Handler)
}
