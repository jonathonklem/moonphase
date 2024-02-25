package communication

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"moonalert/celestial"
)

func ProcessStates() {
	daysUntilFullMoon := celestial.GetDaysUntilFullMoon()
	fmt.Println("Days until full moon: ", daysUntilFullMoon)
	// days until full moon will be negative if we pass it but we're before new moon
	if daysUntilFullMoon <= 3 {
		if daysUntilFullMoon == 0 {
			sendEmail("Full moon is today!", "Please be mindful of the full moon today.")
		} else {
			sendEmail("Full moon is " + fmt.Sprint(daysUntilFullMoon) + " days away!", "Please be mindful of the upcoming full moon.")
		}
	}

	// next let's check if mercury is in retrograde
	if celestial.GetMercuryResponseToday() {
		sendEmail("Mercury is in retrograde", "Please be mindful of the current retrograde.")
	} else {
		fmt.Println("Mercury is not in retrograde")
		if celestial.GetMercuryResponseNextWeek() {
			fmt.Println("Mercury will be in retrograde within a week or less")
			sendEmail("Mercury will be in retrograde within a week or less", "Please be mindful of the upcoming retrograde.")
		} else {
			fmt.Println("Mercury will not be in retrograde within a week")
		}
	}
}

func sendEmail(subject string, body string) {
	accessKey := os.Getenv("MY_AWS_ACCESS_KEY_ID");
	secret := os.Getenv("MY_AWS_SECRET_ACCESS_KEY");
	region := os.Getenv("MY_AWS_REGION");

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