package url

import (
	"errors"

	"gorm.io/gorm"
)

var (
	errSaveUrl = errors.New("error saving url")
	errGetUrl  = errors.New("error getting url")
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) SaveUrl(shortUrl string, originalUrl string) (*Url, error) {

	url := &Url{}

	if err := r.db.Where("long_url = ?", originalUrl).First(url).Error; err == nil {
		return url, nil
	}

	url = &Url{
		ShortUrl: shortUrl,
		LongUrl:  originalUrl,
	}

	if err := r.db.Create(url).Error; err != nil {
		return nil, errSaveUrl
	}

	return url, nil
}

func (r *repo) GetUrl(shortUrl string) (string, error) {
	var url Url

	if err := r.db.Where("short_url = ?", shortUrl).First(&url).Error; err != nil {
		return "", errGetUrl
	}

	return url.LongUrl, nil
}
