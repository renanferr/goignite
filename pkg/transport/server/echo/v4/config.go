package echo

import (
	"log"
	"net/http"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/labstack/echo/v4"
)

const (
	HideBanner                     = "transport.server.echo.hidebanner"
	Port                           = "transport.server.echo.port"
	StatusRoute                    = "transport.server.echo.route.status"
	HealthRoute                    = "transport.server.echo.route.health"
	MiddlewareLogEnabled           = "transport.server.echo.middleware.log.enabled"
	MiddlewareRecoverEnabled       = "transport.server.echo.middleware.recover.enabled"
	MiddlewareNewRelicEnabled      = "transport.server.echo.middleware.newrelic.enabled"
	MiddlewarePrometheusEnabled    = "transport.server.echo.middleware.prometheus.enabled"
	PrometheusRoute                = "transport.server.echo.route.prometheus"
	MiddlewareBodyDumpEnabled      = "transport.server.echo.middleware.bodydump.enabled"
	MiddlewareBodyLimitEnabled     = "transport.server.echo.middleware.bodylimit.enabled"
	MiddlewareBodyLimitSize        = "transport.server.echo.middleware.bodylimit.size"
	MiddlewareCORSEnabled          = "transport.server.echo.middleware.cors.enabled"
	MiddlewareCORSAllowOrigins     = "transport.server.echo.middleware.cors.allow.origins"
	MiddlewareCORSAllowHeaders     = "transport.server.echo.middleware.cors.allow.headers"
	MiddlewareCORSAllowMethods     = "transport.server.echo.middleware.cors.allow.methods"
	MiddlewareCORSAllowCredentials = "transport.server.echo.middleware.cors.allow.credentials"
	MiddlewareCORSExposeHeaders    = "transport.server.echo.middleware.cors.expose.headers"
	MiddlewareCORSMaxAge           = "transport.server.echo.middleware.cors.maxage"
)

func init() {

	log.Println("getting configurations for echo")

	config.Add(HideBanner, true, "echo hide/show banner")
	config.Add(Port, 8080, "server http port")
	config.Add(StatusRoute, "/resource-status", "define status url")
	config.Add(HealthRoute, "/health", "define health url")
	config.Add(MiddlewareLogEnabled, false, "enable/disable logging request middleware")
	config.Add(MiddlewareRecoverEnabled, true, "enable/disable recover middleware")
	config.Add(MiddlewareNewRelicEnabled, false, "enable/disable newrelic middleware")
	config.Add(MiddlewarePrometheusEnabled, false, "enable/disable prometheus middleware")
	config.Add(PrometheusRoute, "/metrics", "define prometheus metrics url")
	config.Add(MiddlewareBodyDumpEnabled, false, "enable/disable body dump middleware")
	config.Add(MiddlewareBodyLimitEnabled, false, "enable/disable body limit middleware")
	config.Add(MiddlewareBodyLimitSize, "8M", "body limit size")
	config.Add(MiddlewareCORSEnabled, false, "enable/disable cors middleware")
	config.Add(MiddlewareCORSAllowOrigins, []string{"*"}, "cors allow origins")
	config.Add(MiddlewareCORSAllowHeaders, []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	"cors allow headers")
	config.Add(MiddlewareCORSAllowMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	config.Add(MiddlewareCORSAllowCredentials, true, "cors allow credentials")
	config.Add(MiddlewareCORSExposeHeaders, []string{}, "cors expose headers")
	config.Add(MiddlewareCORSMaxAge, 5200, "cors max age (seconds)")
}

func GetHideBanner() bool {
	return config.Bool(HideBanner)
}

func GetPort() int {
	return config.Int(Port)
}

func GetStatusRoute() string {
	return config.String(StatusRoute)
}

func GetHealthRoute() string {
	return config.String(HealthRoute)
}

func GetMiddlewareLogEnabled() bool {
	return config.Bool(MiddlewareLogEnabled)
}

func GetMiddlewareRecoverEnabled() bool {
	return config.Bool(MiddlewareRecoverEnabled)
}

func GetMiddlewareNewRelicEnabled() bool {
	return config.Bool(MiddlewareNewRelicEnabled)
}

func GetMiddlewarePrometheusEnabled() bool {
	return config.Bool(MiddlewarePrometheusEnabled)
}

func GetPrometheusRoute() string {
	return config.String(PrometheusRoute)
}

func GetMiddlewareBodyDumpEnabled() bool {
	return config.Bool(MiddlewareBodyDumpEnabled)
}

func GetMiddlewareBodyLimitEnabled() bool {
	return config.Bool(MiddlewareBodyLimitEnabled)
}

func GetMiddlewareBodyLimitSize() string {
	return config.String(MiddlewareBodyLimitSize)
}

func GetMiddlewareCORSEnabled() bool {
	return config.Bool(MiddlewareCORSEnabled)
}

func GetMiddlewareCORSAllowOrigins() []string {
	return config.Strings(MiddlewareCORSAllowOrigins)
}

func GetMiddlewareCORSAllowMethods() []string {
	return config.Strings(MiddlewareCORSAllowMethods)
}

func GetMiddlewareCORSAllowHeaders() []string {
	return config.Strings(MiddlewareCORSAllowHeaders)
}

func GetMiddlewareCORSAllowCredentials() bool {
	return config.Bool(MiddlewareCORSAllowCredentials)
}

func GetMiddlewareCORSExposeHeaders() []string {
	return config.Strings(MiddlewareCORSExposeHeaders)
}

func GetMiddlewareCORSMaxAge() int {
	return config.Int(MiddlewareCORSMaxAge)
}
