package main

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	base64Upload "github.com/heliojuniorkroger/golang-base64-upload"
	"net/http"
)

type image struct {
	Code string `json:"code"`
}

var images = []image{
	{Code: "first image"},
	{Code: "second image"},
	{Code: "Third image"},
}

func getImages(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, images)
}

func postImage(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var newImage image

	if err := c.BindJSON(&newImage); err != nil {
		return
	}

	fileName, _ := getRandomFileName()
	err := base64Upload.Upload("./images/"+fileName+".png", newImage.Code)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, fileName+".png")
}


func getRandomFileName() (s string, err error) {
	b := make([]byte, 8)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	s = fmt.Sprintf("%x", b)
	return
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/images", getImages)
	router.POST("/images", postImage)

	router.Run("localhost:8080")
}
