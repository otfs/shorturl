package shorturl

import (
	"github.com/jmoiron/sqlx"
	"shorturl/domain"
)

type shortUrlRepository struct {
	db *sqlx.DB
}

func NewShortUrlRepository(db *sqlx.DB) *shortUrlRepository {
	return &shortUrlRepository{
		db: db,
	}
}

func (t *shortUrlRepository) Save(shortUrl *domain.ShortUrl) error {
	result, err := t.db.Exec("INSERT INTO short_url(app_id, slug, url, remark, expire_at, create_time) VALUES(?, ?, ?, ?, ?, ?)",
		shortUrl.AppId, shortUrl.Slug, shortUrl.Url, shortUrl.Remark, shortUrl.ExpireAt, shortUrl.CreateTime)
	if err != nil {
		return err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	shortUrl.Id = lastId
	return nil
}

func (t *shortUrlRepository) UpdateSlugById(id int64, slug string) (int64, error) {
	result, err := t.db.Exec("UPDATE short_url SET slug = ? WHERE id = ?", slug, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
