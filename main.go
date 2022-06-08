package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"shorturl/config"
	"shorturl/shorturl"
)

func main() {
	config.Init()
	g := gin.Default()
	route(g)
	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}

func route(g *gin.Engine) {
	db := sqlx.MustOpen(config.Settings.DbDriverName, config.Settings.DbDataSourceName)
	shortUrlRepo := shorturl.NewShortUrlRepository(db)
	shortUrlService := shorturl.NewShortUrlService(shortUrlRepo, config.Settings)
	shorturl.NewShortUrlController(g, shortUrlService, config.Settings)
}
