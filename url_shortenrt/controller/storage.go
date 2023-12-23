package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortener/model"
	"urlshortener/service"
	"urlshortener/utils"
)

func SetUrl(c *gin.Context) {
	requestUrl := model.RequestUrl{}
	c.ShouldBindJSON(&requestUrl)
	originUrl := requestUrl.OriginUrl
	shortUrl := utils.GenerateShortenUrl(originUrl)
	service.SetUrl(originUrl, shortUrl)

	// TODO err
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "generate short url successfully!",
		"data": gin.H{
			"shortUrl":  "http://localhost:8080/" + shortUrl,
			"originUrl": originUrl,
		},
	})
}

func GetUrl(c *gin.Context) {
	//shortUrl, _ := c.Get("short_url")
	shortUrl := c.Params.ByName("short_url")
	originUrl, err := service.GetUrl(shortUrl)

	if err != nil {
		errors.New("get url error!")
	} else {
		// Perform a redirect to the specified URL
		c.Redirect(http.StatusMovedPermanently, originUrl)
	}
}
