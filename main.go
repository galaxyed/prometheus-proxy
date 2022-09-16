package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req)
	}

	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func modifyRequest(req *http.Request) {
	if req.URL.Path == "/api/v1/label/__name__/values" {
		q := req.URL.Query()
		q.Add("match[]", "{project=\"DataV2\"}")
		req.URL.RawQuery = q.Encode()
	}

	if req.URL.Path == "/api/v1/labels" {
		q := req.URL.Query()
		q.Add("match[]", "{project=\"DataV2\"}")
		req.URL.RawQuery = q.Encode()
	}
	if req.URL.Path == "/api/v1/query" {
		req.ParseForm()
		req.ContentLength = 0
		q := req.URL.Query()
		for k, v := range req.Form {
			if k == "query" {
				q.Add(k, v[0]+"{project=\"DataV2\"}")
				continue
			}
			q.Add(k, v[0])
		}
		req.URL.RawQuery = q.Encode()
	}

	if req.URL.Path == "/api/v1/query_range" {
		req.ParseForm()
		req.ContentLength = 0
		q := req.URL.Query()
		for k, v := range req.Form {
			if k == "query" {
				q.Add(k, v[0]+"{project=\"DataV2\"}")
				continue
			}
			q.Add(k, v[0])
		}
		req.URL.RawQuery = q.Encode()
	}

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST:\n%s", string(reqDump))
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		return errors.New("response body is invalid")
	}
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	// initialize a reverse proxy and pass the actual backend server url here
	proxy, err := NewProxy("http://10.100.0.52:9090")
	if err != nil {
		panic(err)
	}

	// handle all requests to your server using the proxy
	http.HandleFunc("/", ProxyRequestHandler(proxy))
	log.Println("Server Started")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
