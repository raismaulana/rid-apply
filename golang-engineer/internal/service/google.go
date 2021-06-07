package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/rid-apply/golang-engineer/internal/config"
	"github.com/raismaulana/rid-apply/golang-engineer/internal/entity"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/callback-gl",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
	oauthStateStringGl = ""
	// repo               repository.UserRepository = repository.NewUserRepository()
)

/*
InitializeOAuthGoogle Function
*/
func InitializeOAuthGoogle() {
	oauthConfGl.ClientID = viper.GetString("google.clientID")
	oauthConfGl.ClientSecret = viper.GetString("google.clientSecret")
	oauthStateStringGl = viper.GetString("oauthStateString")
}

/*
HandleGoogleLogin Function
*/
func HandleGoogleLogin(c *gin.Context) {
	HandleLogin(c, oauthConfGl, oauthStateStringGl)
}

/*
CallBackFromGoogle Function
*/
func CallBackFromGoogle(c *gin.Context) {
	config.Log.Info("Callback-gl..")

	state := c.Query("state")
	config.Log.Info(state)
	if state != oauthStateStringGl {
		config.Log.Info("invalid oauth state, expected " + oauthStateStringGl + ", got " + state + "\n")
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid oauth state, expected "+oauthStateStringGl+", got "+state)
		return
	}

	code := c.Query("code")
	config.Log.Info(code)

	if code == "" {
		config.Log.Warn("Code not found..")
		msg := ("Code Not Found to provide AccessToken..")
		reason := c.Query("error_reason")
		if reason == "user_denied" {
			msg = ("User has denied Permission..")
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
	} else {
		token, err := oauthConfGl.Exchange(context.TODO(), code)
		if err != nil {
			config.Log.Error("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		config.Log.Info("TOKEN>> AccessToken>> " + token.AccessToken)
		config.Log.Info("TOKEN>> Expiration Time>> " + token.Expiry.String())
		config.Log.Info("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			config.Log.Error("Get: " + err.Error() + "\n")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			config.Log.Error("ReadAll: " + err.Error() + "\n")
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		var google entity.Google
		err1 := json.Unmarshal(response, &google)
		if err1 != nil {
			config.Log.Error("JSON UNMARSHAL: " + err.Error() + "\n")
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		config.Log.Info("parseResponseBody: " + string(response) + "\n")
		// if !isRegistered(google.ID) {
		// 	insert(google)
		// }

		msg := `{"access_token":"` + token.AccessToken + `"}`
		var objmap map[string]*json.RawMessage
		json.Unmarshal([]byte(msg), &objmap)
		c.JSON(http.StatusOK, objmap)
		return
	}
}

// func isRegistered(id string) bool {
// 	user := repo.GetUser(id)
// 	return user.ID != 0
// }

// func insert(google entity.Google) {
// 	var sso = []entity.SSO{
// 		{
// 			Provider: "GOOGLE",
// 			Userid:   google.ID,
// 		},
// 	}
// 	var email = []entity.Email{
// 		{
// 			Email: google.Email,
// 		},
// 	}
// 	var user = entity.User{
// 		Name:   google.Name,
// 		Emails: email,
// 		Sso:    sso,
// 	}
// 	repo.InsertUser(user)
// }
