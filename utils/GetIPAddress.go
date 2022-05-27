package utils

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetSourceIP(req *http.Request) string {
	fwdAddress := req.Header.Get(echo.HeaderXForwardedFor)
	if fwdAddress != "" {
		// Got X-Forwarded-For. If we got an array... grab the first IP
		pos := strings.Index(fwdAddress, ",")
		if pos > -1 {
			return fwdAddress[0:pos]
		}

		return fwdAddress
	}

	// No X-Forward-For header, so return the actual IP address...
	return req.RemoteAddr
}
