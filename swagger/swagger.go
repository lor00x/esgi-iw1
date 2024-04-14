package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/option"
)

func handle(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, fmt.Sprintf("[%s] Hello World!", r.Method))
}

func main() {
	api := swag.New(
		option.Title("Example API Doc"),
		option.Security("petstore_auth", "read:pets"),
		option.SecurityScheme("petstore_auth",
			option.OAuth2Security("accessCode", "http://example.com/oauth/authorize", "http://example.com/oauth/token"),
			option.OAuth2Scope("write:pets", "modify pets in your account"),
			option.OAuth2Scope("read:pets", "read your pets"),
		),
	)

	router := echo.New()
	api.Walk(func(path string, e *swag.Endpoint) {
		h := echo.WrapHandler(e.Handler.(http.Handler))
		path = swag.ColonPath(path)

		switch strings.ToLower(e.Method) {
		case "get":
			router.GET(path, h)
		case "head":
			router.HEAD(path, h)
		case "options":
			router.OPTIONS(path, h)
		case "delete":
			router.DELETE(path, h)
		case "put":
			router.PUT(path, h)
		case "post":
			router.POST(path, h)
		case "trace":
			router.TRACE(path, h)
		case "patch":
			router.PATCH(path, h)
		case "connect":
			router.CONNECT(path, h)
		}
	})

	router.GET("/swagger/json", echo.WrapHandler(api.Handler()))
	router.GET("/swagger/ui/*", echo.WrapHandler(swag.UIHandler("/swagger/ui", "/swagger/json", true)))

	log.Fatal(http.ListenAndServe(":8080", router))
}
