package config

import (
	"github.com/gin-gonic/gin"
	"URL-shortener/src/controller"
)

func Routes(r *gin.Engine) {
	url := r.Group("")
	{
		url.POST("/api/v1/urls", controller.UploadUrl)
		// url.GET("/post/:post_id/comment", controller.BrowseCommentController().BrowseComment)
	}
}
