package models

type Url struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}
