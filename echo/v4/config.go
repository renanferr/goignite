package giecho

import (
	"log"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/labstack/echo/v4"
)

const (
	echoRoot                       = "gi.echo"
	hideBanner                     = echoRoot + ".hidebanner"
	port                           = echoRoot + ".port"
	statusRoute                    = echoRoot + ".route.status"
	healthRoute                    = echoRoot + ".route.health"
	jsonPrettyEnabled              = echoRoot + ".json.pretty.enabled"
	middlewareRoot                 = echoRoot + ".middleware"
	middlewareRequestIDEnabled     = middlewareRoot + ".requestid.enabled"
	middlewareLogEnabled           = middlewareRoot + ".log.enabled"
	middlewareSemaphoreEnabled     = middlewareRoot + ".semaphore.enabled"
	middlewareSemaphoreLimit       = middlewareRoot + ".semaphore.limit"
	middlewareRecoverEnabled       = middlewareRoot + ".recover.enabled"
	middlewareBodyDumpEnabled      = middlewareRoot + ".bodydump.enabled"
	middlewareBodyLimitEnabled     = middlewareRoot + ".bodylimit.enabled"
	middlewareBodyLimitSize        = middlewareRoot + ".bodylimit.size"
	middlewareCORSEnabled          = middlewareRoot + ".cors.enabled"
	middlewareCORSAllowOrigins     = middlewareRoot + ".cors.allow.origins"
	middlewareCORSAllowHeaders     = middlewareRoot + ".cors.allow.headers"
	middlewareCORSAllowMethods     = middlewareRoot + ".cors.allow.methods"
	middlewareCORSAllowCredentials = middlewareRoot + ".cors.allow.credentials"
	middlewareCORSExposeHeaders    = middlewareRoot + ".cors.expose.headers"
	middlewareCORSMaxAge           = middlewareRoot + ".cors.maxage"
)

func init() {

	log.Println("getting configurations for echo")

	giconfig.Add(hideBanner, true, "echo hide/show banner")
	giconfig.Add(port, 8080, "server http port")
	giconfig.Add(statusRoute, "/resource-status", "define status url")
	giconfig.Add(healthRoute, "/health", "define health url")
	giconfig.Add(middlewareLogEnabled, false, "enable/disable logging request middleware")
	giconfig.Add(middlewareSemaphoreEnabled, false, "enable/disable semaphore middleware")
	giconfig.Add(middlewareSemaphoreLimit, 10000, "defines numbers for concurrent connections")
	giconfig.Add(middlewareRecoverEnabled, true, "enable/disable recover middleware")
	giconfig.Add(jsonPrettyEnabled, false, "enable/disable json pretty response")
	giconfig.Add(middlewareBodyDumpEnabled, false, "enable/disable body dump middleware")
	giconfig.Add(middlewareBodyLimitEnabled, false, "enable/disable body limit middleware")
	giconfig.Add(middlewareBodyLimitSize, "8M", "body limit size")
	giconfig.Add(middlewareCORSEnabled, false, "enable/disable cors middleware")
	giconfig.Add(middlewareCORSAllowOrigins, []string{"*"}, "cors allow origins")
	giconfig.Add(middlewareCORSAllowHeaders, []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		"cors allow headers")
	giconfig.Add(middlewareCORSAllowMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	giconfig.Add(middlewareCORSAllowCredentials, true, "cors allow credentials")
	giconfig.Add(middlewareCORSExposeHeaders, []string{}, "cors expose headers")
	giconfig.Add(middlewareCORSMaxAge, 5200, "cors max age (seconds)")
}

func GetHideBanner() bool {
	return giconfig.Bool(hideBanner)
}

func GetPort() int {
	return giconfig.Int(port)
}

func GetStatusRoute() string {
	return giconfig.String(statusRoute)
}

func GetHealthRoute() string {
	return giconfig.String(healthRoute)
}

func GetMiddlewareRequestIDEnabled() bool {
	return giconfig.Bool(middlewareRequestIDEnabled)
}
func GetMiddlewareLogEnabled() bool {
	return giconfig.Bool(middlewareLogEnabled)
}

func GetMiddlewareRecoverEnabled() bool {
	return giconfig.Bool(middlewareRecoverEnabled)
}

func GetMiddlewareBodyDumpEnabled() bool {
	return giconfig.Bool(middlewareBodyDumpEnabled)
}

func GetMiddlewareBodyLimitEnabled() bool {
	return giconfig.Bool(middlewareBodyLimitEnabled)
}

func GetMiddlewareBodyLimitSize() string {
	return giconfig.String(middlewareBodyLimitSize)
}

func GetMiddlewareCORSEnabled() bool {
	return giconfig.Bool(middlewareCORSEnabled)
}

func GetMiddlewareCORSAllowOrigins() []string {
	return giconfig.Strings(middlewareCORSAllowOrigins)
}

func GetMiddlewareCORSAllowMethods() []string {
	return giconfig.Strings(middlewareCORSAllowMethods)
}

func GetMiddlewareCORSAllowHeaders() []string {
	return giconfig.Strings(middlewareCORSAllowHeaders)
}

func GetMiddlewareCORSAllowCredentials() bool {
	return giconfig.Bool(middlewareCORSAllowCredentials)
}

func GetMiddlewareCORSExposeHeaders() []string {
	return giconfig.Strings(middlewareCORSExposeHeaders)
}

func GetMiddlewareCORSMaxAge() int {
	return giconfig.Int(middlewareCORSMaxAge)
}

func GetMiddlewareSemaphoreEnabled() bool {
	return giconfig.Bool(middlewareSemaphoreEnabled)
}

func GetMiddlewareSemaphoreLimit() int {
	return giconfig.Int(middlewareSemaphoreLimit)
}

func GetJSONPrettyEnabled() bool {
	return giconfig.Bool(jsonPrettyEnabled)
}
