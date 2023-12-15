package url

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

type Message struct {
	Message string `json:"message"`
}

func (h *Handler) SaveUrl(c echo.Context) error {
	var u UrlCreateRequest

	if err := c.Bind(&u); err != nil {
		return err
	}

	_, err := url.ParseRequestURI(u.OriginalUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Message{Message: "Invalid url"})
	}

	url, err := h.Service.SaveUrl(&u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, url)
}

func (h *Handler) GetUrl(c echo.Context) error {
	shortUrl := c.Param("shortUrl")

	url, err := h.Service.GetUrl(shortUrl)
	if err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Url not found"})
	}

	return c.Redirect(http.StatusMovedPermanently, url)
}
