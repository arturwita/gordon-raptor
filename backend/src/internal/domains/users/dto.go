package users

type CreateUserDto struct {
	Email     string  `json:"email"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Picture   *string `json:"picture,omitempty"`
}
