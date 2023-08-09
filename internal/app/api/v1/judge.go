package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func JudgeProxy(c *gin.Context) {
	// 从负载均衡算法中选择一个目标服务器
	w := srv.W.Next()
	targetURL := w.Item
	fmt.Println("targetURL", targetURL)
	target, _ := url.Parse(targetURL)
	fmt.Println("JudgeProxy", singleJoiningSlash(target.Path, c.Param("path")))
	// 创建反向代理
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			target, _ := url.Parse(targetURL)
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = singleJoiningSlash(target.Path, c.Param("path"))
			req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
			req.Host = target.Host
		},
		ModifyResponse: func(res *http.Response) error {
			if res.StatusCode == 404 {
				Fail(c)
				w.OnInvokeFault()
				return nil
			} else {
				w.OnInvokeSuccess()
			}
			return nil
		},
		Transport: &http.Transport{},
	}
	// 代理请求到目标服务器
	proxy.ServeHTTP(c.Writer, c.Request)
}

func singleJoiningSlash(a, b string) string {
	aslash := len(a) > 0 && a[len(a)-1] == '/'
	bslash := len(b) > 0 && b[0] == '/'
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		if b != "" {
			return a + "/" + b
		}
		return a
	}
	return a + b
}
