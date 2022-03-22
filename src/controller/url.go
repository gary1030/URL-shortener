package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"URL-shortener/src/service"
	"time"
	"fmt"
	"os"
	"strconv"
)

// type UrlController struct {}

// func AddUrlController() UrlController {
// 	return UrlController{}
// }

// func GetUrlController() UrlController {
// 	return UrlController{}
// }

type AddUrlInput struct {
	Url      string    `json:"url" binding:"required"`
	ExpireAt time.Time `json:"expireAt" binding:"required"`
}

// AddUrl
// @Success 200 string successful return id & shortUrl
// @Router /api/v1/urls [post]
func UploadUrl(c *gin.Context) {
	var form AddUrlInput

	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   bindErr.Error(),
		})
		return
	}

	url, err := service.AddUrl(form.Url, form.ExpireAt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
		})
	} else {
		domain := os.Getenv("DOMAIN_NAME")
		shortUrl := fmt.Sprintf("http://%s/url/%d", domain, url.ID)
		c.JSON(http.StatusOK, gin.H{
			"id": url.ID,
			"shortUrl": shortUrl,
		})
	}
}

// Redirect Url
// @Success 200 redirect to original URL
// @Router /url/:url_id [post]
func RedirectUrl(c *gin.Context) {
	id := c.Params.ByName("url_id")
	urlId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	url, err := service.GetUrl(urlId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		cur_time := time.Now()
		if url.Expired_date.After(cur_time){
			c.Redirect(http.StatusMovedPermanently, url.Original_url)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Expired",
			})			
		}
	}
}
