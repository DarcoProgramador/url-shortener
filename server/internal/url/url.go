package url

type Url struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}

type Repository interface {
	SaveUrl(shortUrl string, originalUrl string) (*Url, error)
	GetUrl(shortUrl string) (string, error)
}

type UrlCreateRequest struct {
	OriginalUrl string `json:"original_url"`
}

type UrlCreateResponse struct {
	Id       uint   `json:"id"`
	ShortUrl string `json:"short_url"`
}

type Service interface {
	SaveUrl(req *UrlCreateRequest) (*UrlCreateResponse, error)
	GetUrl(shortUrl string) (string, error)
}
