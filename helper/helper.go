package helper

// 1. create object response 
type Response struct {
	Meta Meta
	Data interface{} // why interface{}? bcoz value of the data can change
}

type Meta struct {
	Message string
	Code    int
	Status  string
}

// 2. Mapping value of response from handler
func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta {
		Message: message,
		Code: code,
		Status: status,
	}

	jsonResponse := Response {
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}