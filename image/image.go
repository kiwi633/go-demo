package imageCompression

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/adrium/goheif"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
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
	if fileType == "image/heic" {
		err := convertHeicToJpg("D:/IMG_1450.heic", "D:/IMG_1450.jpg")
		if err != nil {
			return
		}
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

func convertHeicToJpg(in, out string) error {
	fin, fout := flag.Arg(0), flag.Arg(1)
	fi, err := os.Open(fin)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	exif, err := goheif.ExtractExif(fi)
	if err != nil {
		log.Printf("Warning: no EXIF from %s: %v\n", fin, err)
	}

	img, err := goheif.Decode(fi)
	if err != nil {
		log.Fatalf("Failed to parse %s: %v\n", fin, err)
	}

	fo, err := os.OpenFile(fout, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to create output file %s: %v\n", fout, err)
	}
	defer fo.Close()

	w, _ := newWriterExif(fo, exif)
	err = jpeg.Encode(w, img, nil)
	if err != nil {
		log.Fatalf("Failed to encode %s: %v\n", fout, err)
	}

	log.Printf("Convert %s to %s successfully\n", fin, fout)
	return nil
}
