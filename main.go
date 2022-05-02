package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

const message = "Hello from the other side!"

func main() {
	if err := godotenv.Load(); err != nil {
		return
	}

	client := twilio.NewRestClient()
	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(message)

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
