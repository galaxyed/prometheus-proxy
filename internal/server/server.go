package server

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/galaxyed/prometheus-proxy/internal/conf"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string, cfg *conf.Config) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req, cfg)
	}

	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func modifyRequest(req *http.Request, cfg *conf.Config) {
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

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}
