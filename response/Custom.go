package response

import "github.com/aws/aws-lambda-go/events"

//Custom return custom code and message
func Custom(code int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            message,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: code,
	}, nil
}
