package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	username = "hasaki_service"
	password = "12345_@!#"
)

type facesData struct {
	Bbox         []float64 `json:"bbox"`
	Kps          []float64 `json:"kps"`
	AppearedTime float64   `json:"appeared_time"`
}

type GreetingsRequest struct {
	Ip          string        `json:"ip"`
	Faces       []facesData   `json:"faces"`
	IsCashier   bool          `json:"is_cashier"`
	FrameShapes []interface{} `json:"frame_shapes"`
}

func main() {
	//storeInfor := GreetingsRequest{
	//	Ip:          "192.168.15.6",
	//	Faces:       []interface{}{1, 2, 3},
	//	IsCashier:   true,
	//	FrameShapes: []interface{}{"a", "b"},
	//}

	var metaData = `{"ip": "117.2.24.90:555@3", "faces": [{"bbox": [1537.7679443359375, 579.882568359375, 1689.93310546875, 792.1209716796875], "kps": [[1567.9617919921875, 659.8635864257812], [1636.20849609375, 663.2447509765625], [1593.085693359375, 699.9969482421875], [1573.24853515625, 734.5160522460938], [1629.9600830078125, 736.993408203125]], "appeared_time": 1699322301.3791187}], "is_cashier": false, "frame_shapes": [[1080, 1920, 3]]}`
	metaDataReq := GreetingsRequest{}
	_ = json.Unmarshal([]byte(metaData), &metaDataReq)

	lengthFaces := len(metaDataReq.Faces) - 1
	getAppearedTime := metaDataReq.Faces[lengthFaces].AppearedTime
	//intAppearedTime, _ := strconv.ParseInt(getAppearedTime, 10, 64)
	tm := time.Unix(int64(getAppearedTime), 0)
	fmt.Println("tm: ", tm)

	res, err := call("http://localhost:1323/api/v1/greetings/store", "GET", metaDataReq)
	//res, err := call("http://localhost:1323/api/v1/stores", "GET")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("res: ", *res)
}

func call(url, method string, storeInfoReq GreetingsRequest) (*string, error) {
	client := &http.Client{
		Timeout: time.Microsecond * 500,
	}

	marshalled, _ := json.Marshal(storeInfoReq)
	req, err := http.NewRequest(method, url, bytes.NewReader(marshalled))
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()

	fmt.Println("response: ", response)

	body, err := io.ReadAll(response.Body)
	stringBody := string(body)

	return &stringBody, nil
}
