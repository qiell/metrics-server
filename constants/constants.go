package constants

// constant for server port
const (
	Port = ":8080"
)

// constants for request methods
const (
	Post = "POST"
	Get  = "GET"
)

// constants for routes supported by server
const (
	Metrics = "/metrics"
	Report  = "/report"
)

// constants for header keys and values
const (
	AllowOriginHeaderKey   = "Access-Control-Allow-Origin"
	AllowOriginHeaderValue = "*"
	ContentTypeHeader      = "Content-Type"
	ApplicationJSON        = "application/json"
)
