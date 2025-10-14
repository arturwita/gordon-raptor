package custom_errors

import "errors"

type domainErrorsMap = struct {
	Recipe recipeErrorsMap
	User   userErrorsMap
}

type recipeErrorsMap = struct {
	NotFound error
}

type userErrorsMap = struct {
	NotFound error
}

var DomainErrors = domainErrorsMap{
	Recipe: recipeErrorsMap{
		NotFound: errors.New("recipe not found"),
	},
	User: userErrorsMap{
		NotFound: errors.New("user not found"),
	},
}
