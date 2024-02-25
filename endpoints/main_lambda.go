package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"moonalert/communication"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	godotenv.Load()

	
	communication.ProcessStates();

	message := fmt.Sprintf("Execution complete")
	return &message, nil
}