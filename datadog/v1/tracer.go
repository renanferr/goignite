package gidatadog

import (
	"context"
	"net"
	"os"
	"sync"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/giinfo"
	gihttp "github.com/b2wdigital/goignite/v2/http/v1/client"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var once sync.Once

func NewTracer(ctx context.Context, opts ...tracer.StartOption) {

	if !IsEnabled() {

		once.Do(func() {

			logger := gilog.FromContext(ctx)

			svc := giconfig.String(service)
			if v := giinfo.AppName; v != "" {
				svc = v
			}
			if v := os.Getenv("DD_SERVICE"); v != "" {
				svc = v
			}

			host := giconfig.String(host)
			if v := os.Getenv("DD_AGENT_HOST"); v != "" {
				host = v
			}

			port := giconfig.String(port)
			if v := os.Getenv("DD_TRACE_AGENT_PORT"); v != "" {
				port = v
			}

			env := giconfig.String(env)
			if v := os.Getenv("DD_ENV"); v != "" {
				env = v
			}

			var version string
			if v := giinfo.Version; v != "" {
				version = v
			}
			if v := os.Getenv("DD_VERSION"); v != "" {
				version = v
			}

			addr := net.JoinHostPort(host, port)

			httpClientOpt := &gihttp.Options{}

			err := giconfig.UnmarshalWithPath(httpClientRoot, httpClientOpt)
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
				tracer.WithAnalytics(giconfig.Bool(analytics)),
				tracer.WithAnalyticsRate(giconfig.Float64(analyticsRate)),
				tracer.WithLambdaMode(giconfig.Bool(lambdaMode)),
				tracer.WithDebugMode(giconfig.Bool(debugMode)),
				tracer.WithDebugStack(giconfig.Bool(debugStack)),
			}

			for k, v := range giconfig.StringMap(tags) {
				startOptions = append(startOptions, tracer.WithGlobalTag(k, v))
			}

			startOptions = append(startOptions, opts...)

			tracer.Start(startOptions...)

			logger.Infof("started a datadog tracer: %s", svc)

		})

	}

}
