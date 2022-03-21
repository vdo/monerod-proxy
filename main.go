package main

import (
	"digitalcashtools/monerod-proxy/endpoints"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Failed to read config.ini")
		os.Exit(1)
	}

	http_port := cfg.Section("").Key("http_port").Value()
	fmt.Println("Port from config: ", http_port)

	e := echo.New()
	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	fmt.Println("Request Body Dump")
	// 	fmt.Println(string(reqBody))
	// }))
	endpoints.ConfigurePing(e)
	endpoints.ConfigureMonerodProxyHandler(e, cfg)

	e.GET("*", func(c echo.Context) error {
		reqDump := "GET Request received: " + c.Path() + c.QueryString()
		fmt.Println(reqDump)
		return c.String(http.StatusOK, reqDump)
	})

	e.POST("*", func(c echo.Context) error {
		reqDump := "POST Request received: " + c.Path() + c.QueryString()
		fmt.Println(reqDump)
		return c.String(http.StatusOK, reqDump)
	})

	fmt.Println("Server running, test by visiting localhost:", http_port, "/ping")
	e.Logger.Fatal(e.Start(":" + http_port))
}
