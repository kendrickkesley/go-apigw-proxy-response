package response

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

//ServerError return 500 code
func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	log.Print(fmt.Errorf("ERROR: %v", err))
	return Custom(500, "Internal Server Error", nil)
}
