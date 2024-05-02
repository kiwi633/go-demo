package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	imageCompression "github.com/kiwi633/go-demo/image"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/nfnt/resize"
	"github.com/opentracing/opentracing-go"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"time"
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
func createRouteHandler(e *echo.Echo, routeName string, handler func(c echo.Context) error) {
	e.GET("/api/"+routeName, func(c echo.Context) error {
		span := opentracing.StartSpan(routeName)
		defer span.Finish()
		return handler(c)
	})
}
func hello(c echo.Context) error {
	req := c.Request()
	format := `<code>
			Protocol: %s<br>
			Host: %s<br>
			Remote Address: %s<br>
			Method: %s<br>
			Path: %s<br>
			</code>
			`
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}

func yasuo111(f os.File) {
	start := time.Now()
	// open "test.jpg"
	file, err := os.Open("E:\\git-code\\cc.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	//images, _, _ := image.DecodeConfig(file)

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	//m := resize.Resize(uint(images.Width), uint(images.Height), img, resize.Lanczos3)
	m := resize.Resize(0, 0, img, resize.Lanczos3)

	out, err := os.Create("E:\\git-code\\cc-golang3.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	log.Println("===========================   ", time.Since(start), "   ************")
}
