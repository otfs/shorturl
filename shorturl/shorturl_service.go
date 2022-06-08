package shorturl

import (
	"github.com/speps/go-hashids/v2"
	"shorturl/config"
	"shorturl/domain"
	"time"
)

type shortUrlService struct {
	ShortUrlRepo domain.ShortUrlRepository
	Settings     *config.Config
}

func NewShortUrlService(shortUrlRepo domain.ShortUrlRepository,
	settings *config.Config) *shortUrlService {
	return &shortUrlService{
		ShortUrlRepo: shortUrlRepo,
		Settings:     settings,
	}
}

func (t *shortUrlService) CreateShortUrl(request *domain.CreateShortUrlRequest) (*domain.ShortUrl, error) {
	shortUrl := &domain.ShortUrl{
		AppId:      request.AppId,
		Url:        request.Url,
		Remark:     request.Remark,
		CreateTime: time.Now(),
	}
	err := t.ShortUrlRepo.Save(shortUrl)
	if err != nil {
		return nil, err
	}

	// 短链标识
	slug := t.generateSlug(shortUrl.Id)
	shortUrl.Slug = slug
	_, err = t.ShortUrlRepo.UpdateSlugById(shortUrl.Id, shortUrl.Slug)
	if err != nil {
		return nil, err
	}

	return shortUrl, nil
}

func (t *shortUrlService) generateSlug(id int64) string {
	data := hashids.NewData()
	data.Salt = t.Settings.Salt
	data.MinLength = 6
	hashId, _ := hashids.NewWithData(data)
	slug, _ := hashId.EncodeInt64([]int64{id})
	return slug
}
