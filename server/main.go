package main

import (
	"github.com/DarcoProgramador/url-shortener/config"
	"github.com/DarcoProgramador/url-shortener/db"
	"github.com/DarcoProgramador/url-shortener/internal/url"
	"github.com/DarcoProgramador/url-shortener/router"
)

func main() {
	conf := config.New(true)
	DB := db.Init(conf)

	urlRepo := url.NewRepository(DB)
	urlService := url.NewService(urlRepo)
	urlHandler := url.NewHandler(urlService)

	router.Init(urlHandler)
	router.Start()
}
