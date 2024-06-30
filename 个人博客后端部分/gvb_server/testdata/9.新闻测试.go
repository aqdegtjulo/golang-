package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type NewsResponse struct {
	Code int `json:"code"`
	Data []struct {
		Index    interface{} `json:"index"`
		Title    string      `json:"title"`
		HotValue string      `json:"hotValue"`
		Link     string      `json:"link"`
	} `json:"data"`
	Msg string `json:"msg"`
}

const newsAPI = "https://api.codelife.cc/api/top/list"
const timeout = 2 * time.Second

func main() {
	var params = Params{
		ID:   "KqndgxeLl9",
		Size: 3,
	}
	reqParam, _ := json.Marshal(params)
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", newsAPI, reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("signaturekey", "U2FsdGVkX193W25A9yLmeFA3K5aAyDG18PsSy8gsn90=")
	httpReq.Header.Add("version", "1.3.52")
	client := http.Client{
		Timeout: timeout,
	}

	//http请求
	httpResp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	var response NewsResponse
	byteData, err := io.ReadAll(httpResp.Body)
	fmt.Println(string(byteData))
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.Code != 200 {
		fmt.Println("状态码非200")
		return
	}
	fmt.Println(response.Data)
}
