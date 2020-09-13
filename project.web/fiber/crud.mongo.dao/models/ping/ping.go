package mping

// Request is the Struct that mirror the JSON received when GET /ping
type Request struct {
	Message string `json:"message"`
}

// Response is the Struct that mirror the JSON returned when client send GET /ping request
type Response struct {
	Message string `json:"message"`
}
