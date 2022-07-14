package jobcontroller

import (
  "fmt"
  "log"
  "time"
  "sync"
  "bytes"
  "net/http"
	"io/ioutil"
  "encoding/json"

  "fl/http_server/http/response"
)


func TransferJob(ipAndPort string, uri string, f interface{}, w *sync.WaitGroup) bool {
  defer w.Done()
  url := fmt.Sprintf("http://%s%s", ipAndPort, uri)
  b, e := json.Marshal(f)
  if e != nil {
		log.Fatalf("new request to '%s' failed: %v\n", url, e)
	}
  request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
  if err != nil {
    log.Fatalf("new request to '%s' failed: %v\n", url, err)
  }
  var client = &http.Client{
    Timeout:   time.Second * 30,
  }
  response_, err := client.Do(request)
  if err != nil {
    log.Fatalf("request for '%s' failed: %v\n", url, err)
  }
  defer response_.Body.Close()
  body, _ := ioutil.ReadAll(response_.Body)
  if response_.StatusCode != 200 {
    log.Println("request for '%s' status : %v\n body: %v\n", url, response_.StatusCode, string(body))
    return false
  }
  var ret response.Ret
  err_ := json.Unmarshal([]byte(body), &ret)
  if err_ != nil {
    log.Fatalf("data json loads error:  %v\n", err_)
  }
  return ret.Code == 0
}
