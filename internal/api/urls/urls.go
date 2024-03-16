package urls

import (
	"auth-service/config"
)

type TUrlPatterns struct {
	Base           string
	GoogleLogin    string
	GoogleCallback string
}

func LoadUrlPatterns() *TUrlPatterns {
	urls := &TUrlPatterns{}
	urls.Base = "/auth"
	urls.GoogleLogin = urls.Base + "/google-login"
	urls.GoogleCallback = config.Config.GoogleOAuthConfig.CustomConfig.GoogleCallbackRelativeUri
	return urls
}
