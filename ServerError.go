package response

import (
	"github.com/aws/aws-lambda-go/events"
)

//ServerError return 500 code
func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	return Custom(500, "Internal Server Error", err)
}
