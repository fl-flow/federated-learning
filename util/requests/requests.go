package requests

import (
  "log"
  "time"
	"bytes"
  "net/http"
	"io/ioutil"
)


func Request(
  method string,
  url string,
  jsonData []byte,
  header map[string]string,
) []byte {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("new request to '%s' failed: %v\n", url, err)
	}
  if len(header) != 0 {
    for k, v := range header {
      request.Header.Set(k, v)
    }
  } else {
    request.Header.Set("Content-Type", "application/json; charset=UTF-8")
  }
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("request for '%s' failed: %v\n", url, err) // TODO:
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}


func Post(
  url string,
  jsonData []byte,
  header map[string]string,
) []byte {
  return Request(
    "POST",
    url,
    jsonData,
    header,
  )
}
