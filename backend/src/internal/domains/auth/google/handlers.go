package google

import (
	"gordon-raptor/src/internal/contracts"
	"gordon-raptor/src/internal/domains/auth"
	"gordon-raptor/src/internal/domains/users"
	"gordon-raptor/src/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func NewGoogleLoginHandler(cfg *oauth2.Config) gin.HandlerFunc {
	randomStateString := utils.GenerateRandomString(8)

	return func(context *gin.Context) {
		url := cfg.AuthCodeURL(randomStateString, oauth2.AccessTypeOffline)
		context.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func NewGoogleCallbackHandler(
	cfg *oauth2.Config,
	googleService GoogleService,
	userService users.UserService,
	authService auth.AuthService,
) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()

		googleUserData, customError := googleService.GetUserData(context)
		if customError != nil {
			context.JSON(http.StatusBadRequest, customError)
		}

		user, _ := userService.GetUserByEmail(googleUserData.Email, ctx)
		if user == nil {
			createUserDto := auth.MapGoogleUserToCreateUserDto(googleUserData)
			var err error
			user, err = userService.CreateUser(createUserDto, ctx)
			if err != nil {
				context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: "failed to create a user"})
				return
			}
		}

		token, err := authService.GenerateJWT(user)
		if err != nil || token == "" {
			context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: "failed to login"})
		}

		context.JSON(http.StatusOK, contracts.LoginResponseDto{Token: token})
	}
}
