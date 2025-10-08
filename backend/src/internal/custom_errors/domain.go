package custom_errors

import "errors"

type domainErrorsMap = struct {
	Recipe recipeErrorsMap
}

type recipeErrorsMap = struct {
	NotFound error
}

var DomainErrors = domainErrorsMap{
	Recipe: recipeErrorsMap{
		NotFound: errors.New("recipe not found"),
	},
}
