package google

import (
	"encoding/json"
	"gordon-raptor/src/internal/contracts"
	"gordon-raptor/src/internal/domains/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type GoogleService interface {
	GetUserData(context *gin.Context) (*auth.GoogleOauthUser, *contracts.ErrorResponse)
}

type googleService struct {
	googleOauthConfig *oauth2.Config
}

func NewGoogleService(googleOauthConfig *oauth2.Config) (GoogleService, error) {
	return &googleService{googleOauthConfig}, nil
}

func (service *googleService) GetUserData(context *gin.Context) (*auth.GoogleOauthUser, *contracts.ErrorResponse) {
	code := context.Query("code")
	if code == "" {
		return nil, &contracts.ErrorResponse{Message: "code not found"}
	}

	token, err := service.googleOauthConfig.Exchange(context, code)
	if err != nil {
		return nil, &contracts.ErrorResponse{Message: "failed to exchange oauth token"}
	}

	client := service.googleOauthConfig.Client(context, token)

	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, &contracts.ErrorResponse{Message: "failed to get user info"}
	}
	defer response.Body.Close()

	var userData auth.GoogleOauthUser
	if err := json.NewDecoder(response.Body).Decode(&userData); err != nil {
		return nil, &contracts.ErrorResponse{Message: "failed to decode user payload"}
	}

	return &userData, nil
}
