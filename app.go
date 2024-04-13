package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

type aa int8

func main() {
	query := fmt.Sprintf(`from(bucket: "%v") |> range(start: -1d)`, "my-bucket")
	fmt.Println(query)
}

func aaa() {
	funmap := make(map[string]func(), 23)
	funmap["name"] = getName
	funmap["funcName"] = getPerson
	funcName := "name"
	funmap[funcName]()

	funmap1 := make(map[string]aa, 23)
	funmap1["name"] = 89
	if f, ok := funmap1["funcName"]; ok {
		fmt.Println(f)
	}
}
func getPerson() {
	fmt.Println("getPerson")
}
func getName() {
	fmt.Println("getname")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
