package config

import (
	"github.com/gin-gonic/gin"
	"URL-shortener/src/controller"
)

func Routes(r *gin.Engine) {
	url := r.Group("")
	{
		url.POST("/api/v1/urls", controller.UploadUrl)
		url.GET("/url/:url_id", controller.RedirectUrl)
	}
}
