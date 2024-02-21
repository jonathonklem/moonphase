package main

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"math"
	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ses"
)

type MoonPhase struct {
	CurrentConditions struct {
		Moonphase float64 `json:"moonphase"`
	} `json:"currentConditions"`
}

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY");
	location := os.Getenv("LOCATION");
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + location + "?unitGroup=metric&elements=moonphase&contentType=json&key=" + apiKey

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var moonPhase MoonPhase;
	err = json.NewDecoder(response.Body).Decode(&moonPhase)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// api returns 1 = new moon, .5= full moon. 1/28 = 1 day of moonphase
	daysUntilFullMoon  := int(math.Ceil(((.5-moonPhase.CurrentConditions.Moonphase)/.03571428571)))

	// days until full moon will be negative if we pass it but we're before new moon
	if daysUntilFullMoon > 0 && daysUntilFullMoon <= 3 {
		sendAlert(daysUntilFullMoon)
	}
}

func sendAlert(daysUntilFullMoon int) {
	accessKey := os.Getenv("MY_AWS_ACCESS_KEY_ID");
	secret := os.Getenv("MY_AWS_SECRET_ACCESS_KEY");
	region := os.Getenv("MY_AWS_REGION");
	fmt.Println("Access Key: " + accessKey)
	fmt.Println("Secret: " + secret)
	fmt.Println(region)

	sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(region),
        Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
    })
    if err != nil {
        panic(err)
    }

    // Create a new instance of the AWS SES service
    svc := ses.New(sess)

    // Set the sender and recipient email addresses
    from := os.Getenv("FROM_ADDRESS")
    to := os.Getenv("TO_ADDRESS")

    // Set the email subject and body
    subject := "Full moon is " + fmt.Sprint(daysUntilFullMoon) + " days away!"
    body := "Please be mindful of the upcoming full moon."

    // Send the email
    _, err = svc.SendEmail(&ses.SendEmailInput{
        Destination: &ses.Destination{
            ToAddresses: []*string{
                aws.String(to),
            },
        },
        Message: &ses.Message{
            Body: &ses.Body{
                Text: &ses.Content{
                    Charset: aws.String("UTF-8"),
                    Data:    aws.String(body),
                },
            },
            Subject: &ses.Content{
                Charset: aws.String("UTF-8"),
                Data:    aws.String(subject),
            },
        },
        Source: aws.String(from),
    })
    if err != nil {
        panic(err)
    }
}