package cors

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
	"github.com/labstack/echo/v4"
)

const (
	root             = giecho.ExtRoot + ".cors"
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

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetAllowOrigins() []string {
	return giconfig.Strings(allowOrigins)
}

func GetAllowMethods() []string {
	return giconfig.Strings(allowMethods)
}

func GetAllowHeaders() []string {
	return giconfig.Strings(allowHeaders)
}

func GetAllowCredentials() bool {
	return giconfig.Bool(allowCredentials)
}

func GetExposeHeaders() []string {
	return giconfig.Strings(exposeHeaders)
}

func GetMaxAge() int {
	return giconfig.Int(maxAge)
}
