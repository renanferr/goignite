package echo

import (
	"context"
	"strconv"

	"github.com/b2wdigital/goignite/v2/log"
	"github.com/labstack/echo/v4"
)

type Ext func(context.Context, *echo.Echo) error

type Server struct {
	instance *echo.Echo
	options  *Options
}

func NewDefault(ctx context.Context, exts ...Ext) *Server {
	opt, err := DefaultOptions()
	if err != nil {
		panic(err)
	}
	return New(ctx, opt, exts...)
}

func New(ctx context.Context, opt *Options, exts ...Ext) *Server {

	instance := echo.New()

	instance.HideBanner = opt.HideBanner
	instance.Logger = WrapLogger(log.GetLogger())

	for _, ext := range exts {
		if err := ext(ctx, instance); err != nil {
			panic(err)
		}
	}

	return &Server{instance: instance, options: opt}
}

func (s *Server) Echo() *echo.Echo {
	return s.instance
}

func (s *Server) Serve(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Infof("starting echo Server. https://echo.labstack.com/")
	address := ":" + strconv.Itoa(s.options.Port)
	if err := s.instance.Start(address); err != nil {
		s.instance.Logger.Fatalf(err.Error())
	}
}
