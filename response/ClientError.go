package response

import "github.com/aws/aws-lambda-go/events"

//BadRequest return 400 code
func BadRequest() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Bad Request",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 400,
	}, nil
}

//Unauthorized return 401 code
func Unauthorized() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Unauthorized",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 401,
	}, nil
}

//Forbidden return 403 code
func Forbidden() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Forbidden",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 403,
	}, nil
}

//NotFound return 404 code
func NotFound() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Resource Not Found",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 404,
	}, nil
}

//MethodNotAllowed return 405 code
func MethodNotAllowed() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Method Not Allowed",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 405,
	}, nil
}

//PreconditionFailed return 412 code
func PreconditionFailed() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:            "Precondition Failed",
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: 412,
	}, nil
}
