package giecho

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/b2wdigital/goignite/echo/v4/event"
	mware "github.com/b2wdigital/goignite/echo/v4/middleware"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/b2wdigital/goignite/rest/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	instance *echo.Echo
)

const (
	TopicInstance = "topic:giecho:instance"
	TopicError    = "topic:giecho:error"
)

func Start(ctx context.Context) *echo.Echo {

	instance = echo.New()

	instance.HideBanner = GetHideBanner()
	instance.Logger = Wrap(gilog.GetLogger())

	instance.HTTPErrorHandler = customHTTPErrorHandler

	setDefaultMiddlewares(instance)

	gieventbus.Publish(TopicInstance, instance)

	setDefaultRouters(ctx, instance)

	return instance
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var msg interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else {
		msg = err.Error()
	}

	resp := response.Error{HttpStatusCode: code, Message: fmt.Sprintf("%v", msg)}
	if err := json(c, code, resp); err != nil {
		c.Logger().Error(err)
	}

	requestErr := event.RequestError{
		Context: c,
		Error:   err,
	}

	gieventbus.Publish(TopicError, requestErr)
}

func setDefaultMiddlewares(instance *echo.Echo) {

	if GetMiddlewareRecoverEnabled() {
		instance.Use(middleware.Recover())
	}

	if GetMiddlewareLogEnabled() {
		instance.Use(mware.Logger())
	}

	if GetMiddlewareCORSEnabled() {
		instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     GetMiddlewareCORSAllowOrigins(),
			AllowMethods:     GetMiddlewareCORSAllowMethods(),
			AllowHeaders:     GetMiddlewareCORSAllowHeaders(),
			AllowCredentials: GetMiddlewareCORSAllowCredentials(),
			ExposeHeaders:    GetMiddlewareCORSExposeHeaders(),
			MaxAge:           GetMiddlewareCORSMaxAge(),
		}))
	}

	if GetMiddlewareSemaphoreEnabled() {
		instance.Use(mware.Semaphore(int64(GetMiddlewareSemaphoreLimit())))
	}

	if GetMiddlewareRequestIDEnabled() {
		instance.Use(middleware.RequestID())
	}

	if GetMiddlewareBodyDumpEnabled() {
		instance.Use(middleware.BodyDump(bodyDump))
	}

	if GetMiddlewareBodyLimitEnabled() {
		instance.Use(middleware.BodyLimit(GetMiddlewareBodyLimitSize()))
	}

}

func setDefaultRouters(ctx context.Context, instance *echo.Echo) {

	logger := gilog.FromContext(ctx)

	statusRoute := GetStatusRoute()

	logger.Infof("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

	healthRoute := GetHealthRoute()

	logger.Infof("configuring health router on %s", healthRoute)

	healthHandler := NewHealthHandler()
	instance.GET(healthRoute, healthHandler.Get)

}

func Serve(ctx context.Context) {
	l := gilog.FromContext(ctx)
	l.Infof("starting echo server. https://echo.labstack.com/")
	err := instance.Start(serverPort())
	if err != nil {
		instance.Logger.Fatalf(err.Error())
	}
}

func serverPort() string {
	return ":" + strconv.Itoa(GetPort())
}
