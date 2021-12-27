package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rodny/image-saver/photo"
)

type Handler struct {
	useCase photo.UseCase
}

type imageInput struct {
	Code string `json:"code"`
}

func newHandler(useCase photo.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create(c *gin.Context)  {
	inp := new(imageInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	photoPath, err := h.useCase.CreatePhoto(c.Request.Context(), inp.Code)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusCreated, photoPath)
}
