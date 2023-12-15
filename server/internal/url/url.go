package url

type Url struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}

//go:generate mockery --name=repository --output=urlrepositorymock --inpackage
type Repository interface {
	SaveUrl(shortUrl string, originalUrl string) (*Url, error)
	GetUrl(shortUrl string) (string, error)
}

type UrlCreateRequest struct {
	OriginalUrl string `json:"original_url"`
}

type UrlCreateResponse struct {
	Id       uint   `json:"-"`
	ShortUrl string `json:"short_url"`
}

// go:generate mockery --name=service --output=urlservicemock --inpackage
type Service interface {
	SaveUrl(req *UrlCreateRequest) (*UrlCreateResponse, error)
	GetUrl(shortUrl string) (string, error)
}
