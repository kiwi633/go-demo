package main

import (
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/nfnt/resize"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"
	"xorm.io/xorm"
)

func createRouteHandler(e *echo.Echo, routeName string, handler func(c echo.Context) error) {
	e.GET("/api/"+routeName, func(c echo.Context) error {
		span := opentracing.StartSpan(routeName)
		defer span.Finish()
		return handler(c)
	})
}
func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 32 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		yasuoType := c.Query("yasuoType")
		if yasuoType == "resize" {
			imaging01(c)
		}
		if yasuoType == "thumbnail" {
			fileUploadThumbnail(c)
		}
	})
	router.Run(":8081")

}

func fileUploadResize(c *gin.Context) {

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
	m := resize.Resize(0, 0, img, resize.Bilinear)
	out, err := os.Create("e:/gin-upload-file/yasuo-Resize-" + file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	log.Println("===========================   ", time.Since(start), "   ************")
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}

func fileUploadThumbnail(c *gin.Context) {
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
	log.Println("===========================   ", time.Since(start), "   ************")
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func imaging01(c *gin.Context) {
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
	log.Println("===========================   ", time.Since(start), "   ************")
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
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

type Person struct {
	PersonId   int8
	Name       string
	Idno       string
	ProjectId  string
	CreateTime time.Time
	UpdateTime time.Time
}

func test() {
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

var first *Node

type Node struct {
	Next  *Node
	Value interface{}
}

func jaeger() {
	// 创建 Jaeger 配置
	cfg := config.Configuration{
		ServiceName: "my-echo-app",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}

	// 初始化 Jaeger 追踪器
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("Error initializing Jaeger: %v", err))
	}
	defer closer.Close()

	// 设置全局追踪器
	opentracing.SetGlobalTracer(tracer)

	// 创建 Echo 实例
	e := echo.New()

	// 创建多个路由和 span
	createRouteHandler(e, "route1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Response from Route 1")
	})

	createRouteHandler(e, "route2", func(c echo.Context) error {
		return c.String(http.StatusOK, "Response from Route 2")
	})

	// 启动 Echo 服务器
	e.Start(":1234")
}
func influxdbLog() {
	client := influxdb2.NewClient("http://192.168.0.136:8086", "MCm4grDIAfkj1Jvam9AILVqVIjsgC4H46zcg0BslgH-lR9DJPqVcfLA7s9hdpySTICGqkOJVVHCg8QShUXWKJg==")
	// get non-blocking write client
	writeAPI := client.WriteAPI("aaaa", "user_log")
	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	// Flush writes
	writeAPI.Flush()

	client.Close()
}

func influxdbLog2() {
	client := influxdb2.NewClient("http://192.168.0.136:8086", "MCm4grDIAfkj1Jvam9AILVqVIjsgC4H46zcg0BslgH-lR9DJPqVcfLA7s9hdpySTICGqkOJVVHCg8QShUXWKJg==")

	writeAPI := client.WriteAPI("aaaa", "user_log")

	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45, "cpu_data": "suntong"},
		time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 6565).
		AddField("max", 900000).
		AddField("cpu_data", "suntong").
		SetTime(time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
}

var eg *xorm.Engine

func postgresGET() {

	/**
	var err error
	    master, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
	    if err != nil {
			return
		}
	*/

	eg, _ = xorm.NewEngine("postgres", "postgres://postgres:admin123456@192.168.74.129:5432/mydatabase?sslmode=disable")
	//eg, err = xorm.NewEngine("postgres", "postgres://postgres:admin123456@192.168.74.129:5432/mydatabase?sslmode=disable")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}

	sql := "select * from person where name = ?"
	_, err1 := eg.Exec(sql, "xiaolun")
	if err1 != nil {
		fmt.Println("=======", err1)
		return
	}

	var p Person
	p = Person{
		PersonId:   -1,
		Name:       "suntong",
		Idno:       "123456789",
		ProjectId:  "2",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	eg.Insert(p)
	_, err := eg.SQL("select * from person t where t.name = ? ", `suntong`).Get(&p)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes))
	println(p.CreateTime.Format(time.DateTime))
}
