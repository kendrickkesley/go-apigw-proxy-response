package response

import "github.com/aws/aws-lambda-go/events"

//Redirect return 300 code
func Redirect() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Redirect",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 300,
	}, nil
}

//RedirectWithLocation return 300 code and move to a location
func RedirectWithLocation(location string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Redirect",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
			"Location":                    location,
		},
		StatusCode: 300,
	}, nil
}

//MovedPermanently return 301 code
func MovedPermanently() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Moved Permanently",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 301,
	}, nil
}
