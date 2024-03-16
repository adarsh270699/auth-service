package routers

import (
	"auth-service/internal/api/controllers"
	"auth-service/internal/api/middlewares"
	"auth-service/internal/api/urls"
	"fmt"
	"net/http"
)

func LoadRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	urlPatterns := urls.LoadUrlPatterns()
	fmt.Println(urlPatterns)

	//public routes
	mux.Handle(
		urlPatterns.Base,
		middlewares.BuildMiddlewareChain(
			&controllers.HomeController{},
			middlewares.LogMiddleware,
			middlewares.PublicMiddleware))

	mux.Handle(
		urlPatterns.GoogleLogin,
		middlewares.BuildMiddlewareChain(
			&controllers.GoogleLoginController{},
			middlewares.LogMiddleware,
			middlewares.PublicMiddleware))

	mux.Handle(
		urlPatterns.GoogleCallback,
		middlewares.BuildMiddlewareChain(&controllers.GoogleCallbackController{},
			middlewares.LogMiddleware,
			middlewares.PublicMiddleware))

	//private routes
	return mux
}
