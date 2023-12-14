package url

import "github.com/DarcoProgramador/url-shortener/util"

type serv struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serv{
		repo: repo,
	}
}

func (s *serv) SaveUrl(req *UrlCreateRequest) (*UrlCreateResponse, error) {
	shortUrl := util.GenerateShortUrl(req.OriginalUrl)

	url, err := s.repo.SaveUrl(shortUrl, req.OriginalUrl)
	if err != nil {
		return nil, err
	}

	return &UrlCreateResponse{
		Id:       url.Id,
		ShortUrl: url.ShortUrl,
	}, nil
}

func (s *serv) GetUrl(shortUrl string) (string, error) {
	return s.repo.GetUrl(shortUrl)
}
