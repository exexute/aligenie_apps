package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// 	request.Header.Add("Content-Type", "application/json") //json请求
func SendPost(url string, headers map[string]string, data interface{}) []byte {
	dataStr, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(string(dataStr)))
	if err != nil {
		panic(err)
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return b
}
