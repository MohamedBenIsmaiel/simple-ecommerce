package handlError

import(
	"net/http"
)

type HandleError struct{
	Message 	string		`json:"message"`
	Code 		int			`json:"code"`
	Error 		string		`json:"error"`
}

func BadRequest(message string) *HandleError{
	return &HandleError{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "Bad_Request",
	}
}

func NotFound(message string) *HandleError{
	return &HandleError{
		Message: message,
		Code: http.StatusNotFound,
		Error: "Not_Found",
	}
}

func InternalServerError(message string) *HandleError{
	return &HandleError{
		Message: message,
		Code: http.StatusInternalServerError,
		Error: "Internal_Server_Error",
	}
}
