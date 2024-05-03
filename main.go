package main

import (
	"github.com/gin-gonic/gin"
	imageCompression "github.com/kiwi633/go-demo/image"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 64 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		yasuoType := c.Query("yasuoType")
		if yasuoType == "resize" {
			imageCompression.Imaging01(c)
		}
		if yasuoType == "thumbnail" {
			imageCompression.FileUploadThumbnail(c)
		}
	})
	err := router.Run(":8081")
	if err != nil {
		return
	}
}
