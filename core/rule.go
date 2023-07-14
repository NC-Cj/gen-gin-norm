package core

/*
SUCCESS The program met expectations and was successfully processed
FAILURE Exception status code captured within the service
InternalError Service internal exception, not caught, fatal error code
*/
const (
	SUCCESS       = 1000
	FAILURE       = 1001
	InternalError = 1002
)

/*
PRINT_ERROR_STACK The exception thrown in the program will print the error stack
OUTPUT_INTERNAL_ERROR Uncaptured service internal exception will be returned in the response
ADD_LINK_TRACKING_MIDDLEWARE Add request headers middleware
ADD_REQUEST_LOGGING_MIDDLEWARE Add request logging middleware
*/
const (
	PRINT_ERROR_STACK              = false
	OUTPUT_INTERNAL_ERROR          = true
	ADD_LINK_TRACKING_MIDDLEWARE   = true
	ADD_REQUEST_LOGGING_MIDDLEWARE = true
)
