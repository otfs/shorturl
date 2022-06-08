package shorturl

import (
	"github.com/gin-gonic/gin"
	"log"
	"shorturl/config"
	"shorturl/domain"
)

//
// shortUrlController 短链控制层
//
type shortUrlController struct {
	shortUrlService domain.ShortUrlService
	settings        *config.Config
}

func NewShortUrlController(g *gin.Engine, shortUrlService domain.ShortUrlService,
	settings *config.Config) *shortUrlController {
	ctl := &shortUrlController{
		shortUrlService: shortUrlService,
		settings:        settings,
	}
	g.POST("/shorturl/create", ctl.CreateShortUrl)

	return ctl
}

// CreateShortUrl 创建短链
func (t *shortUrlController) CreateShortUrl(c *gin.Context) {
	request := new(domain.CreateShortUrlRequest)
	if err := c.BindJSON(request); err != nil {
		c.JSON(400, domain.NewResultFail(domain.ResultCodeBadRequest, err.Error()))
		return
	}
	shortUrl, err := t.shortUrlService.CreateShortUrl(request)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(500, domain.NewResultFail(domain.ResultCodeInternalError, err.Error()))
		return
	}

	response := domain.CreateShortUrlResponse{
		Url: t.calculateUrl(shortUrl.Slug),
	}
	c.JSON(200, domain.NewResultOk(response))
}

func (t *shortUrlController) calculateUrl(slug string) string {
	return t.settings.BaseUrl + "/" + slug
}
