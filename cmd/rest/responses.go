package rest

type GlobalError struct {
	Error string `json:"error"`
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
