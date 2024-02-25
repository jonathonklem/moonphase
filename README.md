# Moon Phase Reminder
Get email reminders of the current moon phase via a Lambda function.  You'll have to configure a custom trigger.  I've chosen daily.  This makes use of the https://www.visualcrossing.com/ API.  You'll need a free account to be able to plug in your API credentials.  It also uses Amazon SES to send email alerts.

## Configuration
To deploy locally, you can use .env.  If deploying to lambda, be sure to set all the configurations to match the serverless file: `aws ssm put-parameter --name "YYYY" --type "String" --value "XXXXX"`

The necessary configuration variables are:

```
API_KEY=__
LOCATION=__
MY_AWS_ACCESS_KEY_ID=__
MY_AWS_SECRET_ACCESS_KEY=__
MY_AWS_REGION=__
FROM_ADDRESS=__
TO_ADDRESS=__
```

LOCATION is url encoded, so it would be "evansville,%20in" for example.  Spaces can also be omitted.

## Running
Run locally with `go run ./endpoints/main_local.go`.  `make` builds an executable that can be deployed to lambda