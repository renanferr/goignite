package datadog

import (
	"context"
	"net"
	"os"
	"sync"

	"github.com/b2wdigital/goignite/v2/config"
	gihttp "github.com/b2wdigital/goignite/v2/http/v1/client"
	"github.com/b2wdigital/goignite/v2/info"
	"github.com/b2wdigital/goignite/v2/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var once sync.Once

func NewTracer(ctx context.Context, opts ...tracer.StartOption) {

	if !IsEnabled() {
		return
	}

	once.Do(func() {

		logger := log.FromContext(ctx)

		svc := config.String(service)
		if v := info.AppName; v != "" {
			svc = v
		}
		if v := os.Getenv("DD_SERVICE"); v != "" {
			svc = v
		}

		host := config.String(host)
		if v := os.Getenv("DD_AGENT_HOST"); v != "" {
			host = v
		}

		port := config.String(port)
		if v := os.Getenv("DD_TRACE_AGENT_PORT"); v != "" {
			port = v
		}

		env := config.String(env)
		if v := os.Getenv("DD_ENV"); v != "" {
			env = v
		}

		var version string
		if v := info.Version; v != "" {
			version = v
		}
		if v := os.Getenv("DD_VERSION"); v != "" {
			version = v
		}

		addr := net.JoinHostPort(host, port)

		httpClientOpt := &gihttp.Options{}

		err := config.UnmarshalWithPath(httpClientRoot, httpClientOpt)
		if err != nil {
			logger.Panic(err)
		}

		httpClient := gihttp.NewClient(ctx, httpClientOpt)

		startOptions := []tracer.StartOption{
			tracer.WithAgentAddr(addr),
			tracer.WithEnv(env),
			tracer.WithService(svc),
			tracer.WithServiceVersion(version),
			tracer.WithLogger(NewLogger()),
			tracer.WithHTTPClient(httpClient),
			tracer.WithAnalytics(config.Bool(analytics)),
			tracer.WithAnalyticsRate(config.Float64(analyticsRate)),
			tracer.WithLambdaMode(config.Bool(lambdaMode)),
			tracer.WithDebugMode(config.Bool(debugMode)),
			tracer.WithDebugStack(config.Bool(debugStack)),
		}

		for k, v := range config.StringMap(tags) {
			startOptions = append(startOptions, tracer.WithGlobalTag(k, v))
		}

		startOptions = append(startOptions, opts...)

		tracer.Start(startOptions...)

		logger.Infof("started a datadog tracer: %s", svc)

	})

}
