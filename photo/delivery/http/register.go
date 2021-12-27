package http

import (
	"github.com/gin-gonic/gin"
	"rodny/image-saver/photo"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc photo.UseCase)  {
	h := newHandler(uc)

	photos := router.Group("/photos")
	{
		photos.POST("", h.Create)
	}
}
