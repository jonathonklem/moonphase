# Configuration

To deploy locally, you can use .env.  If deploying to lambda, be sure to set all the configurations to match the serverless file: `aws ssm put-parameter --name "YYYY" --type "String" --value "XXXXX"`