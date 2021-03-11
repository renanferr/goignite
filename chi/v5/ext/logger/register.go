package gichilogger

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-chi/chi/v5/middleware"
)

func Register(ctx context.Context) (*gichi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("enabling logger middleware in chi")

	return &gichi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			loggerMiddleware,
		},
	}, nil

}

// loggerMiddleware returns a middleware that logs HTTP requests.
func loggerMiddleware(next http.Handler) http.Handler {

	level := Level()

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		t1 := time.Now()
		reqId := middleware.GetReqID(ctx)
		preReqContent := gilog.Fields{
			"time":      t1,
			"requestId": reqId,
			"method":    r.Method,
			"endpoint":  r.RequestURI,
			"protocol":  r.Proto,
		}

		if r.RemoteAddr != "" {
			preReqContent["ip"] = r.RemoteAddr
		}

		tid := r.Header.Get("X-TID")
		if tid != "" {
			preReqContent["tid"] = tid
		}

		logger := gilog.FromContext(ctx).WithFields(preReqContent)
		ctx = logger.ToContext(ctx)
		r = r.WithContext(ctx)
		logger.Info("request started")

		defer func() {
			if err := recover(); err != nil {
				gilog.WithFields(
					gilog.Fields{
						"requestId":  reqId,
						"duration":   time.Since(t1),
						"status":     500,
						"stacktrace": string(debug.Stack()),
					},
				).Error("request finished with panic")
				panic(err)
			}
		}()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		status := ww.Status()
		postReqContent := gilog.Fields{
			"requestId":     reqId,
			"duration":      time.Since(t1),
			"contentLength": ww.BytesWritten(),
			"status":        status,
		}

		if cache := ww.Header().Get("x-cache"); cache != "" {
			postReqContent["cache"] = cache
		}

		logger = gilog.FromContext(ctx).WithFields(postReqContent)
		if status >= 100 && status < 500 {

			var method func(format string, args ...interface{})

			switch level {
			case "TRACE":
				method = logger.Tracef
			case "DEBUG":
				method = logger.Debugf
			default:
				method = logger.Infof
			}

			method("request finished")
		} else if status == 500 {
			logger.WithField("stacktrace",
				string(debug.Stack())).Error("internal error during request")
		} else {
			message := "request finished"

			// FIX: For some reason, the 'context.deadlineExceededError{}' isn't getting into here, we
			// did a quick fix checking the status code and returing the same message as the error., but
			// something is wrong and we need fix it.
			if status == 504 {
				message += ": context deadline exceeded"
			} else {
				if err := ctx.Err(); err != nil {
					message += fmt.Sprintf(": %s", err.Error())
				}
			}
			logger.Error(message)
		}
	}

	return http.HandlerFunc(fn)
}
