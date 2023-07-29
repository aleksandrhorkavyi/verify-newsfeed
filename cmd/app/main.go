package main

import (
	"fmt"
	"github.com/aleksandrhorkavyi/verify-newsfeed/internal/article"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = DB()
	r := gin.Default()

	var (
		articleRepo = article.NewRepository()
		articleSvc  = article.NewService(articleRepo)
	)

	article.NewController(r, articleSvc).InitRoutes()

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func DB() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "newsfeed"
		password = "newsfeed"
		dbname   = "newsfeed"
		timeZone = "Europe/Kiev"
		sslmode  = "disable"
	)
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
		timeZone,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,  // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
