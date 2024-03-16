package controllers

import (
	"auth-service/config"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HomeController struct {
}

func (h *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth-service"))
}

type GoogleLoginController struct {
	url string
}

func (g *GoogleLoginController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.url = config.Config.GoogleOAuthConfig.OAuthConfig.AuthCodeURL("bridge-app")
	w.Write([]byte(g.url))
}

type GoogleCallbackController struct {
}

func (g *GoogleCallbackController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.GoogleOAuthConfig.CustomConfig.GoogleOathContextDeadline)
	defer cancel()
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")

	if state == "bridge-app" {
		googleConfig := config.Config.GoogleOAuthConfig.OAuthConfig
		token, err := googleConfig.Exchange(ctx, code)
		if err == nil {
			resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
			if err == nil {
				userinfo := make(map[string]string)
				decoder := json.NewDecoder(resp.Body)
				err := decoder.Decode(&userinfo)
				if err == nil {
					w.Write([]byte("ok"))
				}
			}
		}

		if err != nil {
			w.Write([]byte(fmt.Sprintf("error in oauth, error:%s", err)))
		}

	} else {
		w.Write([]byte("error in oauth, error:invalid state"))
	}

}
