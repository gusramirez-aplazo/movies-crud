package main

func newSuccessResponse(content any) *MovieResponse {
	return &MovieResponse{
		Ok:      true,
		Message: "",
		Status:  200,
		Content: content,
	}
}

func newErrorResponse(message string, statusCode uint) *MovieResponse {
	return &MovieResponse{
		Ok:      false,
		Message: message,
		Status:  statusCode,
		Content: nil,
	}
}
