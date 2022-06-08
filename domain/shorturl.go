package domain

import (
	"time"
)

// ShortUrl 短链实体
type ShortUrl struct {
	Id         int64     `db:"id"`          // 短链Id
	AppId      string    `db:"app_id"`      // 所属应用
	Slug       string    `db:"slug"`        // 短链地址
	Url        string    `db:"url"`         // 原始url地址
	Remark     string    `db:"remark"`      // 备注
	Hints      int       `db:"hints"`       // 命中次数
	ExpireAt   int64     `db:"expire_at"`   // 过期时间轴（毫秒）
	CreateTime time.Time `db:"create_time"` // 创建时间
}

type ShortUrlRepository interface {
	Save(shortUrl *ShortUrl) error

	UpdateSlugById(id int64, slug string) (int64, error)
}

type ShortUrlService interface {
	CreateShortUrl(request *CreateShortUrlRequest) (*ShortUrl, error)
}

//-------------------------------------------------------------
// 请求和响应
//-------------------------------------------------------------

// CreateShortUrlRequest 创建短链请求数据
type CreateShortUrlRequest struct {
	AppId    string `json:"appId" form:"appId"`                // 所属应用
	Url      string `json:"url" form:"url" binding:"required"` // 原始url地址
	Remark   string `json:"remark" form:"remark"`              // 备注
	ExpireAt int64  `json:"expireAt" form:"expireAt"`          // 过期时间轴（毫秒）
}

// CreateShortUrlResponse 创建短链响应数据
type CreateShortUrlResponse struct {
	Url string `json:"url"` // 短链
}
