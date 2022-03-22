package main

import (
    "net/http"
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/gorm"
    "URL-shortener/src/model"
    "URL-shortener/src/persistence"
    "URL-shortener/src/config"
    "time"
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

    var db *gorm.DB
    var db_err error
    dsn := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbName
	db, db_err = persistence.Initialize(dsn)
    if db_err != nil {
        log.Fatal("Error loading db")
    }

    t, _ := time.Parse(time.RFC3339, "2023-02-08T09:20:41")
    url := model.Url{
        Original_url: "https://www.google.com.tw/?hl=zh_TW",
        Expired_date: t,
	}
    db.AutoMigrate(&model.Url{})
    
    // Insert
    db.Model(&model.Url{}).Create(&url)


    domain := os.Getenv("DOMAIN_NAME")
    router := gin.Default()
	router.GET("/test", test)
    config.Routes(router)

    router.Run(domain)
}