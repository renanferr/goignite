package giecho

import (
	"log"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/labstack/echo/v4"
)

const (
	HideBanner                     = "gi.echo.hidebanner"
	Port                           = "gi.echo.port"
	StatusRoute                    = "gi.echo.route.status"
	HealthRoute                    = "gi.echo.route.health"
	PProfEnabled                   = "gi.echo.pprof.enabled"
	JSONPrettyEnabled              = "gi.echo.json.pretty.enabled"
	MiddlewareRequestIDEnabled     = "gi.echo.middleware.requestid.enabled"
	MiddlewareLogEnabled           = "gi.echo.middleware.log.enabled"
	MiddlewareSwaggerEnabled       = "gi.echo.middleware.swagger.enabled"
	MiddlewareRecoverEnabled       = "gi.echo.middleware.recover.enabled"
	MiddlewareNewRelicEnabled      = "gi.echo.middleware.newrelic.enabled"
	MiddlewarePrometheusEnabled    = "gi.echo.middleware.prometheus.enabled"
	PrometheusRoute                = "gi.echo.route.prometheus"
	SwaggerRoute                   = "gi.echo.route.swagger"
	MiddlewareBodyDumpEnabled      = "gi.echo.middleware.bodydump.enabled"
	MiddlewareBodyLimitEnabled     = "gi.echo.middleware.bodylimit.enabled"
	MiddlewareBodyLimitSize        = "gi.echo.middleware.bodylimit.size"
	MiddlewareCORSEnabled          = "gi.echo.middleware.cors.enabled"
	MiddlewareCORSAllowOrigins     = "gi.echo.middleware.cors.allow.origins"
	MiddlewareCORSAllowHeaders     = "gi.echo.middleware.cors.allow.headers"
	MiddlewareCORSAllowMethods     = "gi.echo.middleware.cors.allow.methods"
	MiddlewareCORSAllowCredentials = "gi.echo.middleware.cors.allow.credentials"
	MiddlewareCORSExposeHeaders    = "gi.echo.middleware.cors.expose.headers"
	MiddlewareCORSMaxAge           = "gi.echo.middleware.cors.maxage"
)

func init() {

	log.Println("getting configurations for echo")

	giconfig.Add(HideBanner, true, "echo hide/show banner")
	giconfig.Add(Port, 8080, "server http port")
	giconfig.Add(StatusRoute, "/resource-status", "define status url")
	giconfig.Add(HealthRoute, "/health", "define health url")
	giconfig.Add(MiddlewareSwaggerEnabled, false, "enable/disable swagger middleware")
	giconfig.Add(SwaggerRoute, "/swagger", "define prometheus metrics url")
	giconfig.Add(MiddlewareRequestIDEnabled, false, "enable/disable request id middleware")
	giconfig.Add(MiddlewareLogEnabled, false, "enable/disable logging request middleware")
	giconfig.Add(MiddlewareRecoverEnabled, true, "enable/disable recover middleware")
	giconfig.Add(PProfEnabled, false, "enable/disable pprof")
	giconfig.Add(JSONPrettyEnabled, false, "enable/disable json pretty response")
	giconfig.Add(MiddlewareNewRelicEnabled, false, "enable/disable newrelic middleware")
	giconfig.Add(MiddlewarePrometheusEnabled, false, "enable/disable prometheus middleware")
	giconfig.Add(PrometheusRoute, "/metrics", "define prometheus metrics url")
	giconfig.Add(MiddlewareBodyDumpEnabled, false, "enable/disable body dump middleware")
	giconfig.Add(MiddlewareBodyLimitEnabled, false, "enable/disable body limit middleware")
	giconfig.Add(MiddlewareBodyLimitSize, "8M", "body limit size")
	giconfig.Add(MiddlewareCORSEnabled, false, "enable/disable cors middleware")
	giconfig.Add(MiddlewareCORSAllowOrigins, []string{"*"}, "cors allow origins")
	giconfig.Add(MiddlewareCORSAllowHeaders, []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		"cors allow headers")
	giconfig.Add(MiddlewareCORSAllowMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	giconfig.Add(MiddlewareCORSAllowCredentials, true, "cors allow credentials")
	giconfig.Add(MiddlewareCORSExposeHeaders, []string{}, "cors expose headers")
	giconfig.Add(MiddlewareCORSMaxAge, 5200, "cors max age (seconds)")
}

func GetHideBanner() bool {
	return giconfig.Bool(HideBanner)
}

func GetPort() int {
	return giconfig.Int(Port)
}

func GetStatusRoute() string {
	return giconfig.String(StatusRoute)
}

func GetHealthRoute() string {
	return giconfig.String(HealthRoute)
}

func GetMiddlewareSwaggerEnabled() bool {
	return giconfig.Bool(MiddlewareSwaggerEnabled)
}

func GetMiddlewareRequestIDEnabled() bool {
	return giconfig.Bool(MiddlewareRequestIDEnabled)
}

func GetMiddlewareLogEnabled() bool {
	return giconfig.Bool(MiddlewareLogEnabled)
}

func GetMiddlewareRecoverEnabled() bool {
	return giconfig.Bool(MiddlewareRecoverEnabled)
}

func GetMiddlewareNewRelicEnabled() bool {
	return giconfig.Bool(MiddlewareNewRelicEnabled)
}

func GetMiddlewarePrometheusEnabled() bool {
	return giconfig.Bool(MiddlewarePrometheusEnabled)
}

func GetPrometheusRoute() string {
	return giconfig.String(PrometheusRoute)
}

func GetSwaggerRoute() string {
	return giconfig.String(SwaggerRoute)
}

func GetMiddlewareBodyDumpEnabled() bool {
	return giconfig.Bool(MiddlewareBodyDumpEnabled)
}

func GetMiddlewareBodyLimitEnabled() bool {
	return giconfig.Bool(MiddlewareBodyLimitEnabled)
}

func GetMiddlewareBodyLimitSize() string {
	return giconfig.String(MiddlewareBodyLimitSize)
}

func GetMiddlewareCORSEnabled() bool {
	return giconfig.Bool(MiddlewareCORSEnabled)
}

func GetMiddlewareCORSAllowOrigins() []string {
	return giconfig.Strings(MiddlewareCORSAllowOrigins)
}

func GetMiddlewareCORSAllowMethods() []string {
	return giconfig.Strings(MiddlewareCORSAllowMethods)
}

func GetMiddlewareCORSAllowHeaders() []string {
	return giconfig.Strings(MiddlewareCORSAllowHeaders)
}

func GetMiddlewareCORSAllowCredentials() bool {
	return giconfig.Bool(MiddlewareCORSAllowCredentials)
}

func GetMiddlewareCORSExposeHeaders() []string {
	return giconfig.Strings(MiddlewareCORSExposeHeaders)
}

func GetMiddlewareCORSMaxAge() int {
	return giconfig.Int(MiddlewareCORSMaxAge)
}

func GetPProfEnabled() bool {
	return giconfig.Bool(PProfEnabled)
}

func GetJSONPrettyEnabled() bool {
	return giconfig.Bool(JSONPrettyEnabled)
}
