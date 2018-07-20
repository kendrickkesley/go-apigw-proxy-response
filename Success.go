package response

import (
	"github.com/aws/aws-lambda-go/events"
)

//SuccessJSON return 200 code and a JSON body
func SuccessJSON(responseObject interface{}) (events.APIGatewayProxyResponse, error) {
	return CustomJSON(200, responseObject, nil)
}

//SuccessText return 200 code and a Text body
func SuccessText(responseText string) (events.APIGatewayProxyResponse, error) {
	return Custom(200, responseText, nil)
}
