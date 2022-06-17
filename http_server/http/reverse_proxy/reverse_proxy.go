package reverseproxy

import(
  "log"
  "fmt"
  "net/url"
  "net/http"
  "net/http/httputil"
  "github.com/gin-gonic/gin"
)


func ReverseProxy(c *gin.Context, remote string, schema string) {
    // target := "127.0.0.1:8080"
    // schema "http"
    u := &url.URL{}
    u.Scheme = schema
    u.Host = remote
    proxy := httputil.NewSingleHostReverseProxy(u)
    proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
      log.Printf("http: proxy error: %v", err)
      ret := fmt.Sprintf("http proxy error %v", err)
      rw.Write([]byte(ret))
    }
    proxy.ServeHTTP(c.Writer, c.Request)
}
