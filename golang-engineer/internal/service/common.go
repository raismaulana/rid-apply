package service

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/rid-apply/golang-engineer/internal/config"
	"github.com/raismaulana/rid-apply/golang-engineer/internal/helper/page"
	"golang.org/x/oauth2"
)

/*
HandleMain Function renders the index page when the application index route is called
*/
func HandleMain(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(page.IndexPage))
}

/*
HandleLogin Function
*/
func HandleLogin(c *gin.Context, oauthConf *oauth2.Config, oauthStateString string) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		config.Log.Error("Parse: " + err.Error())
	}
	config.Log.Info(URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	config.Log.Info(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
