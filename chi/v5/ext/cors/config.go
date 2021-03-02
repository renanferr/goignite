package cors

import (
	"net/http"

	gichi "github.com/b2wdigital/goignite/chi/v5"
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root               = gichi.ExtRoot + ".cors"
	enabled            = root + ".enabled"
	allowedOrigins     = root + ".allowed.origins"
	allowedHeaders     = root + ".allowed.headers"
	allowedMethods     = root + ".allowed.methods"
	allowedCredentials = root + ".allowed.credentials"
	exposedHeaders     = root + ".exposed.headers"
	maxAge             = root + ".maxage"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable cors middleware")
	giconfig.Add(allowedOrigins, []string{"*"}, "cors allow origins")
	giconfig.Add(allowedHeaders, []string{"Origin", "Content-Type", "Accept"},
		"cors allow headers")
	giconfig.Add(allowedMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	giconfig.Add(allowedCredentials, true, "cors allow credentials")
	giconfig.Add(exposedHeaders, []string{}, "cors expose headers")
	giconfig.Add(maxAge, 5200, "cors max age (seconds)")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}

func getAllowedOrigins() []string {
	return giconfig.Strings(allowedOrigins)
}

func getAllowedMethods() []string {
	return giconfig.Strings(allowedMethods)
}

func getAllowedHeaders() []string {
	return giconfig.Strings(allowedHeaders)
}

func getAllowedCredentials() bool {
	return giconfig.Bool(allowedCredentials)
}

func getExposedHeaders() []string {
	return giconfig.Strings(exposedHeaders)
}

func getMaxAge() int {
	return giconfig.Int(maxAge)
}
