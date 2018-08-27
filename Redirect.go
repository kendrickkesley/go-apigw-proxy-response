package response

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

//Redirect return 300 code
func Redirect() (events.APIGatewayProxyResponse, error) {
	return Custom(300, "Redirect", nil)
}

//RedirectWithLocation return 300 code and move to a location
func RedirectWithLocation(location string) (events.APIGatewayProxyResponse, error) {
	fmt.Fprintf(os.Stderr, "RESPONSE-%d: %s", 301, location)
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Location":                    location,
		},
		StatusCode: 301,
	}, nil
}

//MovedPermanently return 301 code
func MovedPermanently() (events.APIGatewayProxyResponse, error) {
	return Custom(301, "Moved Permanently", nil)
}
