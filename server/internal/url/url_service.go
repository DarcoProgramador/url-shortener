package url

import (
	"errors"
	"os"

	"github.com/DarcoProgramador/url-shortener/util"
)

var (
	ErrInvalidUrl = errors.New("invalid url")
)

type serv struct {
	Repository
}

func NewService(repo Repository) Service {
	return &serv{
		Repository: repo,
	}
}

func (s *serv) SaveUrl(req *UrlCreateRequest) (*UrlCreateResponse, error) {
	shortUrl := util.GenerateShortUrl(req.OriginalUrl)

	if shortUrl == "" {
		return nil, ErrInvalidUrl
	}
	url, err := s.Repository.SaveUrl(shortUrl, req.OriginalUrl)
	if err != nil {
		return nil, err
	}

	host := os.Getenv("HOST")

	if host == "" {
		host = "http://localhost:1323/"
	}

	return &UrlCreateResponse{
		Id:       url.Id,
		ShortUrl: host + url.ShortUrl,
	}, nil
}

func (s *serv) GetUrl(shortUrl string) (string, error) {
	return s.Repository.GetUrl(shortUrl)
}
