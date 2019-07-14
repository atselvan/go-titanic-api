package model

// Message represents the https status and message
//
// swagger:response Message
type Message struct {
	// The status of a request.
	//
	// required: true
	Status string `json:"status"`
	// Details message explaining the request.
	//
	// required: true
	Message string `json:"message"`
}

// Success
//
// swagger:response success
type SuccessResp struct {
	// Status of a request
	Status string `json:"status"`
	// Message explaining the status
	Message string `json:"message"`
}

// Bad Request
//
// swagger:response badRequest
type BadReqResp struct {
	// Status of a request
	Status string `json:"status"`
	// Message explaining the status
	Message string `json:"message"`
}

// Not Found
//
// swagger:response notFound
type NotFoundResp struct {
	// Status of a request
	Status string `json:"status"`
	// Message explaining the status
	Message string `json:"message"`
}

// Internal Server Error
//
// swagger:response internalServerError
type IntServErrResp struct {
	// Status of a request
	Status string `json:"status"`
	// Message explaining the error
	Message string `json:"message"`
}
