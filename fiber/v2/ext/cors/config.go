package cors

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
	"github.com/gofiber/fiber/v2"
)

const (
	root             = gifiber.ExtRoot + ".cors"
	enabled          = root + ".enabled"
	allowOrigins     = root + ".allow.origins"
	allowHeaders     = root + ".allow.headers"
	allowMethods     = root + ".allow.methods"
	allowCredentials = root + ".allow.credentials"
	exposeHeaders    = root + ".expose.headers"
	maxAge           = root + ".maxAge"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable cors middleware")
	giconfig.Add(allowOrigins, []string{"*"}, "cors allow origins")
	giconfig.Add(allowHeaders, []string{fiber.HeaderOrigin, fiber.HeaderContentType, fiber.HeaderAccept},
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
