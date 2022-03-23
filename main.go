package main

import (
	"URL-shortener/src/config"
	"URL-shortener/src/model"
	"URL-shortener/src/persistence"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func test(c *gin.Context) {
	var message = "Hello world!"
	c.IndentedJSON(http.StatusOK, message)
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DBNAME")

	dsn := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbName
	db, db_err := persistence.Initialize(dsn)
	if db_err != nil {
		log.Fatal("Error loading db")
	}

	db.AutoMigrate(&model.Url{})

	t, _ := time.Parse(time.RFC3339, "2030-12-31T23:59:59")
	url := model.Url{
		Original_url: "https://www.google.com.tw/",
		Expired_date: t,
	}
	// Insert
	db.Model(&model.Url{}).Create(&url)

	domain := os.Getenv("DOMAIN_NAME")
	router := gin.Default()
	router.GET("/test", test)
	config.Routes(router)

	router.Run(domain)
}
