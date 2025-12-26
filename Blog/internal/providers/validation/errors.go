package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "The Field is required",
		"email":    "The Field must be a valid email address",
		"min":      "The Field must be greater than the limit",
		"max":      "The Field may not be greater than the limit",
	}
}
