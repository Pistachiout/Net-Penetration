package main

import (
	"Net-Penetration/constant"
	"encoding/json"
	"log"
	"net/http"
)

// 本地应用，用于测试内网穿透
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()
		b, err := json.Marshal(q)
		if err != nil {
			log.Println(err)
		}
		writer.Write(b)
	})
	log.Printf("本地服务已启动：%s\n", constant.AppPort)
	http.ListenAndServe(constant.AppPort, nil)
}
