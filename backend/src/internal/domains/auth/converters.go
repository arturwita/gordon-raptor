package auth

import (
	"gordon-raptor/src/internal/domains/users"
)

func MapGoogleUserToCreateUserDto(googleUser *GoogleOauthUser) *users.CreateUserDto {
	return &users.CreateUserDto{
		Email:     googleUser.Email,
		FirstName: &googleUser.GivenName,
		LastName:  &googleUser.FamilyName,
		Picture:   &googleUser.Picture,
	}
}
