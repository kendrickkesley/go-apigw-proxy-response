package response

import (
	"github.com/aws/aws-lambda-go/events"
)

const (
	//BadRequestRequired error code for required fields
	BadRequestRequired = "REQUIRED"
	//BadRequestDuplicate error code for duplicate fields
	BadRequestDuplicate = "DUPLICATE"
	//BadRequestInvalid error code for invalid fields
	BadRequestInvalid = "INVALID"
	//BadRequestNoMatch error code for no matching found
	BadRequestNoMatch = "NO_MATCH"
)

//BadRequest return 400 code
func BadRequest() (events.APIGatewayProxyResponse, error) {
	return Custom(400, "Bad Request", nil)
}

//BadRequestWithCode return 400 code
func BadRequestWithCode(codes map[string]string) (events.APIGatewayProxyResponse, error) {
	return CustomJSON(400, codes, nil)
}

//Unauthorized return 401 code
func Unauthorized() (events.APIGatewayProxyResponse, error) {
	return Custom(401, "Unauthorized", nil)
}

//Forbidden return 403 code
func Forbidden() (events.APIGatewayProxyResponse, error) {
	return Custom(403, "Forbidden", nil)
}

//NotFound return 404 code
func NotFound() (events.APIGatewayProxyResponse, error) {
	return Custom(404, "Resource Not Found", nil)
}

//MethodNotAllowed return 405 code
func MethodNotAllowed() (events.APIGatewayProxyResponse, error) {
	return Custom(405, "Method Not Allowed", nil)
}

//PreconditionFailed return 412 code
func PreconditionFailed() (events.APIGatewayProxyResponse, error) {
	return Custom(412, "Precondition Failed", nil)
}

//Locked return 423 code
func Locked() (events.APIGatewayProxyResponse, error) {
	return Custom(423, "Locked", nil)
}
