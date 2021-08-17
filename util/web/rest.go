package web

import (
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 25 * time.Second}

func Get(url string) []byte {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	return body
}
