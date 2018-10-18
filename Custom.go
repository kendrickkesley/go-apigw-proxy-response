package response

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kkesley/go-apigw-httpparser"

	"github.com/aws/aws-lambda-go/events"
)

//Custom return custom code and message
func Custom(code int, message string, err error) (events.APIGatewayProxyResponse, error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	}
	fmt.Fprintf(os.Stderr, "RESPONSE-%d: %s", code, message)
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

//CustomJSON return custom code and message
func CustomJSON(code int, responseObject interface{}, extErr error) (events.APIGatewayProxyResponse, error) {
	resByte, err := json.Marshal(responseObject)
	if err != nil {
		return ServerError(err)
	}
	httpparser.LogResponse(code, responseObject)
	return events.APIGatewayProxyResponse{
		Body:            string(resByte),
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: code,
	}, nil
}
