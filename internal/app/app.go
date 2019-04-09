package app

import (
	"github.com/PrakharSrivastav/api-example/internal/api/blog"
	"github.com/PrakharSrivastav/api-example/internal/api/comment"
	"github.com/PrakharSrivastav/api-example/internal/api/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type App struct {
	user    *user.Api
	blog    *blog.Api
	comment *comment.Api
}

// New cerates a new app instance
func New() *App {
	return &App{
		user:    user.New(),
		blog:    blog.New(),
		comment: comment.New(),
	}
}

func (a *App) Configure(e *echo.Echo) {
	e.HideBanner = true
	e.DisableHTTP2 = true
	e.Debug = true
	e.HidePort = true
}

func (a *App) Middleware(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// add your middlewares
}

func (a *App) Routes(e *echo.Echo) {
	log.Info("Registering routes")
	root := e.Group("api/v1/")
	a.user.Routes(root)
	a.comment.Routes(root)
	a.blog.Routes(root)
}

func (a *App) Server() *http.Server {
	log.Info("Setting up http server")

	return &http.Server{
		Addr:         ":1313",          // port
		ReadTimeout:  20 * time.Minute, // read timeout
		WriteTimeout: 20 * time.Minute, // write timeout
	}
}
