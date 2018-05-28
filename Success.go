package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

//SuccessJSON return 200 code and a JSON body
func SuccessJSON(responseObject interface{}) (events.APIGatewayProxyResponse, error) {
	resByte, err := json.Marshal(responseObject)
	if err != nil {
		return ServerError(err)
	}
	return events.APIGatewayProxyResponse{
		Body:            string(resByte[:]),
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 200,
	}, nil
}

//SuccessText return 200 code and a Text body
func SuccessText(responseText string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            responseText,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 200,
	}, nil
}
