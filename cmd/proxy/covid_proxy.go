package main

import (
	"flag"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/labstack/echo/v4"
)

var (
	urlApi, _ = url.Parse("https://xx9p7hp1p7.execute-api.us-east-1.amazonaws.com")
	bind string
)

func init() {
	flag.StringVar(&bind, "bind", "0.0.0.0:8080", "bind address")
	flag.Parse()
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Any("*", handler)

	log.Println("running on " + bind)
	e.Logger.Fatal(
		e.Start(bind),
		)
}

func handler(c echo.Context) error {
	rProxy := proxyHTTP(urlApi)
	req := c.Request()

	// change request headers
	req.Host =urlApi.Host
	req.URL.Scheme = urlApi.Scheme
	req.URL.Host = urlApi.Host

	rProxy.ServeHTTP(c.Response(), c.Request())
	return nil
}

func proxyHTTP(t *url.URL) http.Handler {
	p := httputil.NewSingleHostReverseProxy(t)
	p.ModifyResponse = func(r *http.Response) error {

		// alterar os headers de reposta caso o metodo seha options
		if r.Request.Method == http.MethodOptions {
			r.Header.Set("Access-Control-Allow-Origin", "*")
			r.Header.Set("Access-Control-Allow-Methods", "*")
			r.Header.Set("Access-Control-Allow-Headers", "*")
		}
		return nil
	}
	return p
}
