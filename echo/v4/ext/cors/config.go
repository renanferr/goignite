package cors

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
	"github.com/labstack/echo/v4"
)

const (
	root             = giecho.MiddlewareRoot + ".cors"
	enabled          = root + ".enabled"
	allowOrigins     = root + ".allow.origins"
	allowHeaders     = root + ".allow.headers"
	allowMethods     = root + ".allow.methods"
	allowCredentials = root + ".allow.credentials"
	exposeHeaders    = root + ".expose.headers"
	maxAge           = root + ".maxage"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable cors middleware")
	giconfig.Add(allowOrigins, []string{"*"}, "cors allow origins")
	giconfig.Add(allowHeaders, []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		"cors allow headers")
	giconfig.Add(allowMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	giconfig.Add(allowCredentials, true, "cors allow credentials")
	giconfig.Add(exposeHeaders, []string{}, "cors expose headers")
	giconfig.Add(maxAge, 5200, "cors max age (seconds)")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}

func getAllowOrigins() []string {
	return giconfig.Strings(allowOrigins)
}

func getAllowMethods() []string {
	return giconfig.Strings(allowMethods)
}

func getAllowHeaders() []string {
	return giconfig.Strings(allowHeaders)
}

func getAllowCredentials() bool {
	return giconfig.Bool(allowCredentials)
}

func getExposeHeaders() []string {
	return giconfig.Strings(exposeHeaders)
}

func getMaxAge() int {
	return giconfig.Int(maxAge)
}
