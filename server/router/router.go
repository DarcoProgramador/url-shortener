package router

import (
	"github.com/DarcoProgramador/url-shortener/internal/url"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func Init(urlHandler *url.Handler) {
	e = echo.New()
	e.POST("/", urlHandler.SaveUrl)
	e.GET("/:shortUrl", urlHandler.GetUrl)
}

func Start() {
	e.Logger.Fatal(e.Start(":1323"))
}
