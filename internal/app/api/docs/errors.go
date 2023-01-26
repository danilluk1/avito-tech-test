package docs

// NotFoundError Not Found
// swagger:response notFoundError
type NotFoundError struct {
	//in: body
	Body string
}

// InternalError Internal error
// Some internal error happened
// swagger:response internalError
type InternalError struct {
	//in: body
	Body string
}

// ValidationError is an error that used when the required input fails validation
//
//swagger:response validationError
type ValidationError struct {
	//The error message
	// in: body
	Body struct {
		//The validation message
		//
		// Required: true
		// Example: []
		Messages []string `json:"messages"`
	}
}
