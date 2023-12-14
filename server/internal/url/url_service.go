package url

import (
	"errors"

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

	return &UrlCreateResponse{
		Id:       url.Id,
		ShortUrl: url.ShortUrl,
	}, nil
}

func (s *serv) GetUrl(shortUrl string) (string, error) {
	return s.Repository.GetUrl(shortUrl)
}
