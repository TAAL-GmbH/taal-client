package console

import (
	"embed"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed public/*

	res embed.FS
)

func AppHandler(c echo.Context) error {

	var resource string

	if c.Request().URL.Path == "/" {
		resource = "public/index.html"
	} else {
		resource = strings.Replace(c.Request().URL.Path, "/", "public/", 1)
	}

	b, err := res.ReadFile(resource)
	if err != nil {
		// Just in case we're missing the /index.html, add it and try again...
		resource += "/index.html"
		b, err = res.ReadFile(resource)
		if err != nil {
			return c.String(http.StatusNotFound, "Not found")
		}
	}

	var mimeType string
	if strings.HasSuffix(resource, ".css") {
		mimeType = "text/css"
	} else if strings.HasSuffix(resource, ".js") {
		mimeType = "text/javascript"
	} else if strings.HasSuffix(resource, ".png") {
		mimeType = "image/png"
	} else {
		mimeType = "text/html"
	}

	return c.Blob(http.StatusOK, mimeType, b)

}
