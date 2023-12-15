package router

import (
	"os"

	"github.com/DarcoProgramador/url-shortener/internal/url"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func Init(urlHandler *url.Handler) {
	e = echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/", urlHandler.SaveUrl)
	e.GET("/:shortUrl", urlHandler.GetUrl)
}

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
