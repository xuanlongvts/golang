package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataVectorsResponse struct {
	Result map[string][]float64 `json:"result"`
}

type EachFile struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ContentReq map[string]string

type DataReq struct {
	Data ContentReq `json:"data"`
}

func main() {
	var contReq = ContentReq{}
	for i := 0; i < 1; i++ {
		iStr := fmt.Sprintf("%d", i)
		contReq["title_"+iStr] = "title_" + iStr
		contReq["content_"+iStr] = "content_ " + iStr
	}
	dataReq := DataReq{Data: contReq}
	bodyReq, err := json.Marshal(dataReq)
	request, err := http.NewRequest("POST", "http://ai.hasaki.vn:8001/embedding-batch", bytes.NewBuffer(bodyReq))
	if err != nil {
		fmt.Println("err : ", err)
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("err 1: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err 2: ", err)
	}

	var result = &DataVectorsResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:", result.Result)
	for k, v := range result.Result {
		fmt.Println("k: ", k)
		fmt.Println("v: ", v)
	}
}
