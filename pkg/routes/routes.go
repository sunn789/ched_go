package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sunn789/ched.ir/main"
)

func (app *main.Application) routes() http.Handler {
	Mux := echo.New()
	Mux.Use(middleware.Recover())
	Mux.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "error on start up app",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			log.Println(c.Path())
		},
		Timeout: 60 * time.Second,
	}))
	return Mux
}
