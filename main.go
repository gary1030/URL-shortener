package main

import (
    "net/http"
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
	"gorm.io/gorm"
    "github.com/joho/godotenv"
    "URL-shortener/src/model"
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
	var db_err error
    var db *gorm.DB
	db, db_err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if db_err != nil {
        log.Fatal("Error loading db")
    }

    db.AutoMigrate(&model.URL{})

    router := gin.Default()
	router.GET("/test", test)

    router.Run("localhost:8080")
}