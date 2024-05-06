package imageCompression

import (
	"encoding/base64"
	"github.com/disintegration/imaging"
	"image/gif"
	"strings"
)

// 这个是一个测试注释
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"
)

func Compression() {
	// Open a test image.
	src, err := imaging.Open("E:\\git-code\\a.jpg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 1000, 0, imaging.Lanczos)

	// Save the resulting image as JPEG.
	err = imaging.Save(src, "E:\\git-code\\a-yasuo01.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func FileUploadThumbnail(c *gin.Context) {
	// 单文件
	start := time.Now()
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	fff, _ := file.Open()
	fileType := file.Header.Get("content-type")
	var img image.Image
	if fileType == "image/jpeg" {
		img, _ = jpeg.Decode(fff)
	}
	if fileType == "image/png" {
		img, _ = png.Decode(fff)
	}
	m := resize.Thumbnail(800, 800, img, resize.NearestNeighbor)
	out, err := os.Create("e:/gin-upload-file/yasuo-Thumbnail-" + file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	log.Println("耗费时间：", time.Since(start), " ms")
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func Imaging01(c *gin.Context) {
	start := time.Now()
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	fff, _ := file.Open()
	fileType := file.Header.Get("content-type")
	var img image.Image
	if fileType == "image/jpeg" {
		img, _ = jpeg.Decode(fff)
	}
	if fileType == "image/gif" {
		img, _ = gif.Decode(fff)
	}
	if fileType == "image/png" {
		img, _ = png.Decode(fff)
		encoder := png.Encoder{CompressionLevel: 9}
		out, _ := os.Create("e:/gin-upload-file/yasuo-imaging0122-" + file.Filename)
		defer out.Close()
		encoder.Encode(out, img)
		return
	}
	//m := imaging.Blur(img, 0.75)
	m := imaging.Resize(img, 800, 0, imaging.Bartlett)
	out, err := os.Create("e:/gin-upload-file/yasuo-imaging01-" + file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	log.Println("耗费时间：", time.Since(start), " ms")
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func Base64Image(base64Str string) {
	imgBase64 := strings.Replace(base64Str, "data:image/png;base64,", "", 1)
	imageBatys, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		log.Println("图片base64解码失败！")
		return
	}
	img, _, err := image.Decode(strings.NewReader(string(imageBatys)))
	if err != nil {
		log.Println("读取图片失败！")
		return
	}
	outFile, err := os.Create("d://base64-01.jpg")
	if err != nil {
		log.Println("写入图片流失败！")
		return
	}
	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		log.Println("写入图片流失败！")
		return
	}
}
