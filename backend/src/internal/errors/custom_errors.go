package errors

import (
	"net/http"
)

type customError struct {
	Status  int    `json:"code"`
	Message string `json:"message"`
}

type customErrorsMap = struct {
	Auth authErrorsMap
}

type authErrorsMap = struct {
	BadCredentials customError
	Forbidden customError
}

var CustomErrors = customErrorsMap{
	Auth: authErrorsMap{
		BadCredentials: customError{Status: http.StatusUnauthorized, Message: "Bad credentials"},
		Forbidden: customError{Status: http.StatusForbidden, Message: "You're not allowed to perform this action"},
	},
}
