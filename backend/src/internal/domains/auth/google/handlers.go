package google

import (
	"encoding/json"
	"gordon-raptor/src/internal/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func NewGoogleLoginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		url := GoogleOauthConfig.AuthCodeURL("random-state-string", oauth2.AccessTypeOffline) // todo: replace random-string
		context.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func NewGoogleCallbackHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		code := context.Query("code")
		if code == "" {
			context.JSON(http.StatusBadRequest, contracts.ErrorResponse{Message: "code not found"})
			return
		}

		token, err := GoogleOauthConfig.Exchange(context, code)
		if err != nil {
			context.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Message: "failed to exchange oauth token"})
			return
		}

		client := GoogleOauthConfig.Client(context, token)

		response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			context.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Message: "failed to get user info"})
			return
		}
		defer response.Body.Close()

		var userData contracts.GoogleOauthUser
		if err := json.NewDecoder(response.Body).Decode(&userData); err != nil {
			context.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Message: "failed to decode user payload"})
			return
		}

		context.JSON(http.StatusOK, userData)
	}
}
