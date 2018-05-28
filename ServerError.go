package response

import (
	"github.com/aws/aws-lambda-go/events"
)

//ServerError return 500 code
func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Internal Server Error",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 500,
	}, err
}
