package gifiber

import (
	"log"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/gofiber/fiber/v2"
)

const (
	fiberRoot                      = "gi.fiber"
	statusRoute                    = fiberRoot + ".route.status"
	healthRoute                    = fiberRoot + ".route.health"
	port                           = fiberRoot + ".port"
	configRoot                     = fiberRoot + ".config"
	prefork                        = configRoot + ".prefork"
	serverHeader                   = configRoot + ".serverHeader"
	strictRouting                  = configRoot + ".strictRouting"
	caseSensitive                  = configRoot + ".caseSensitive"
	immutable                      = configRoot + ".immutable"
	unescapePath                   = configRoot + ".unescapePath"
	ETag                           = configRoot + ".ETag"
	bodyLimit                      = configRoot + ".bodyLimit"
	concurrency                    = configRoot + ".concurrency"
	readTimeout                    = configRoot + ".readTimeout"
	writeTimeout                   = configRoot + ".writeTimeout"
	idleTimeout                    = configRoot + ".idleTimeout"
	readBufferSize                 = configRoot + ".readBufferSize"
	writeBufferSize                = configRoot + ".writeBufferSize"
	compressedFileSuffix           = configRoot + ".compressedFileSuffix"
	proxyHeader                    = configRoot + ".proxyHeader"
	GETOnly                        = configRoot + ".GETOnly"
	reduceMemoryUsage              = configRoot + ".reduceMemoryUsage"
	network                        = configRoot + ".network"
	disableKeepalive               = configRoot + ".disableKeepalive"
	disableDefaultDate             = configRoot + ".disableDefaultDate"
	disableDefaultContentType      = configRoot + ".disableDefaultContentType"
	disableHeaderNormalizing       = configRoot + ".disableHeaderNormalizing"
	disableStartupMessage          = configRoot + ".disableStartupMessage"
	middlewareRoot                 = fiberRoot + ".middleware"
	middlewareRequestIDEnabled     = middlewareRoot + ".requestid.enabled"
	middlewareLogEnabled           = middlewareRoot + ".log.enabled"
	middlewareRecoverEnabled       = middlewareRoot + ".recover.enabled"
	middlewareCORSEnabled          = middlewareRoot + ".cors.enabled"
	middlewareCORSAllowOrigins     = middlewareRoot + ".cors.allow.origins"
	middlewareCORSAllowHeaders     = middlewareRoot + ".cors.allow.headers"
	middlewareCORSAllowMethods     = middlewareRoot + ".cors.allow.methods"
	middlewareCORSAllowCredentials = middlewareRoot + ".cors.allow.credentials"
	middlewareCORSExposeHeaders    = middlewareRoot + ".cors.expose.headers"
	middlewareCORSMaxAge           = middlewareRoot + ".cors.maxage"
	middlewarePprofEnabled         = middlewareRoot + ".pprof.enabled"
	middlewareMonitorEnabled       = middlewareRoot + ".monitor.enabled"
	middlewareETagEnabled          = middlewareRoot + ".etag.enabled"
)

func init() {

	log.Println("getting configurations for fiber")

	giconfig.Add(port, 8080, "server http port")
	giconfig.Add(statusRoute, "/resource-status", "define status url")
	giconfig.Add(healthRoute, "/health", "define health url")
	giconfig.Add(prefork, false, "Enables use of the SO_REUSEPORT socket option. This will spawn multiple Go processes listening on the same port. learn more about socket sharding.")
	giconfig.Add(serverHeader, "", "Enables the Server HTTP header with the given value.")
	giconfig.Add(strictRouting, false, "When enabled, the router treats /foo and /foo/ as different. Otherwise, the router treats /foo and /foo/ as the same.")
	giconfig.Add(caseSensitive, false, "When enabled, /Foo and /foo are different routes. When disabled, /Fooand /foo are treated the same.")
	giconfig.Add(immutable, false, "When enabled, all values returned by context methods are immutable. By default, they are valid until you return from the handler; see issue #185.")
	giconfig.Add(unescapePath, false, "Converts all encoded characters in the route back before setting the path for the context, so that the routing can also work with URL encoded special characters")
	giconfig.Add(ETag, false, "Enable or disable ETag header generation, since both weak and strong etags are generated using the same hashing method (CRC-32). Weak ETags are the default when enabled.")
	giconfig.Add(bodyLimit, 4*1024*1024, "Sets the maximum allowed size for a request body, if the size exceeds the configured limit, it sends 413 - Request Entity Too Large response.")
	giconfig.Add(concurrency, 256*1024, "Maximum number of concurrent connections.")
	giconfig.Add(readTimeout, "0s", "The amount of time allowed to read the full request, including the body. The default timeout is unlimited.")
	giconfig.Add(writeTimeout, "0s", "The maximum duration before timing out writes of the response. The default timeout is unlimited.")
	giconfig.Add(idleTimeout, "0s", "The maximum amount of time to wait for the next request when keep-alive is enabled. If IdleTimeout is zero, the value of ReadTimeout is used.")
	giconfig.Add(readBufferSize, 4096, "per-connection buffer size for requests' reading. This also limits the maximum header size. Increase this buffer if your clients send multi-KB RequestURIs and/or multi-KB headers (for example, BIG cookies).")
	giconfig.Add(writeBufferSize, 4096, "Per-connection buffer size for responses' writing.")
	giconfig.Add(compressedFileSuffix, ".fiber.gz", "Adds a suffix to the original file name and tries saving the resulting compressed file under the new file name.")
	giconfig.Add(proxyHeader, "", "This will enable c.IP() to return the value of the given header key. By default c.IP()will return the Remote IP from the TCP connection, this property can be useful if you are behind a load balancer e.g. X-Forwarded-*.")
	giconfig.Add(GETOnly, false, "Rejects all non-GET requests if set to true. This option is useful as anti-DoS protection for servers accepting only GET requests. The request size is limited by ReadBufferSize if GETOnly is set.")
	giconfig.Add(reduceMemoryUsage, false, "Aggressively reduces memory usage at the cost of higher CPU usage if set to true")
	giconfig.Add(network, fiber.NetworkTCP4, "Known networks are \"tcp\", \"tcp4\" (IPv4-only), \"tcp6\" (IPv6-only)")
	giconfig.Add(disableKeepalive, false, "Disable keep-alive connections, the server will close incoming connections after sending the first response to the client")
	giconfig.Add(disableDefaultDate, false, "When set to true causes the default date header to be excluded from the response.")
	giconfig.Add(disableDefaultContentType, false, "When set to true, causes the default Content-Type header to be excluded from the Response.")
	giconfig.Add(disableHeaderNormalizing, false, "By default all header names are normalized: conteNT-tYPE -> Content-Type")
	giconfig.Add(disableStartupMessage, false, "When set to true, it will not print out debug information")
	giconfig.Add(middlewareLogEnabled, false, "enable/disable logging request middleware")
	giconfig.Add(middlewareRecoverEnabled, true, "enable/disable recover middleware")
	giconfig.Add(middlewareCORSEnabled, false, "enable/disable cors middleware")
	giconfig.Add(middlewareCORSAllowOrigins, []string{"*"}, "cors allow origins")
	giconfig.Add(middlewareCORSAllowHeaders, []string{fiber.HeaderOrigin, fiber.HeaderContentType, fiber.HeaderAccept},
		"cors allow headers")
	giconfig.Add(middlewareCORSAllowMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	giconfig.Add(middlewareCORSAllowCredentials, true, "cors allow credentials")
	giconfig.Add(middlewareCORSExposeHeaders, []string{}, "cors expose headers")
	giconfig.Add(middlewareCORSMaxAge, 5200, "cors max age (seconds)")
	giconfig.Add(middlewarePprofEnabled, false, "enable/disable pprof middleware")
	giconfig.Add(middlewareMonitorEnabled, false, "enable/disable monitor middleware")
	giconfig.Add(middlewareETagEnabled, false, "enable/disable etag middleware")

}

func Port() int {
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

func GetMiddlewarePprofEnabled() bool {
	return giconfig.Bool(middlewarePprofEnabled)
}

func GetMiddlewareMonitorEnabled() bool {
	return giconfig.Bool(middlewareMonitorEnabled)
}

func GetMiddlewareETagEnabled() bool {
	return giconfig.Bool(middlewareETagEnabled)
}

func AppConfig() (*fiber.Config, error) {

	o := &fiber.Config{}

	err := giconfig.UnmarshalWithPath(configRoot, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
