package custom_errors

import (
	"net/http"
)

type httpError struct {
	Status  int    `json:"code"`
	Message string `json:"message"`
}

type httpErrorsMap = struct {
	Auth authErrorsMap
}

type authErrorsMap = struct {
	Forbidden httpError
}

var HttpErrors = httpErrorsMap{
	Auth: authErrorsMap{
		Forbidden: httpError{Status: http.StatusForbidden, Message: "You're not allowed to perform this action"},
	},
}
