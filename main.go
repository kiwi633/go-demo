package main

import (
	"github.com/gin-gonic/gin"
	imageCompression "github.com/kiwi633/go-demo/image"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 32 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		yasuoType := c.Query("yasuoType")
		if yasuoType == "resize" {
			imageCompression.Imaging01(c)
		}
		if yasuoType == "thumbnail" {
			imageCompression.FileUploadThumbnail(c)
		}
	})
	router.Run(":8081")

}
