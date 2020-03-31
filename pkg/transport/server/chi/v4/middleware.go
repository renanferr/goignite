package chi

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/b2wdigital/goignite/pkg/info"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/client/newrelic/v3"
	"github.com/go-chi/chi/middleware"
	nr "github.com/newrelic/go-agent/v3/newrelic"
	uuid "github.com/satori/go.uuid"
)

// NewTidMiddleware is a middleware that looks for a XTID value inside the http.Request
// and generate one if it does not exists.
func NewTidMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tid := r.Header.Get("x-tid")
			if tid == "" {
				tid = info.AppName + "-" + uuid.NewV4().String()
			}
			w.Header().Set("x-tid", tid)
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// NewLogMiddleware returns a middleware that logs HTTP requests.
func NewLogMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		t1 := time.Now()
		reqId := middleware.GetReqID(ctx)
		preReqContent := log.Fields{
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

		logger := log.FromContext(ctx).WithFields(preReqContent)
		ctx = logger.ToContext(ctx)
		r = r.WithContext(ctx)
		logger.Info("request started")

		defer func() {
			if err := recover(); err != nil {
				log.WithFields(
					log.Fields{
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
		postReqContent := log.Fields{
			"requestId":     reqId,
			"duration":      time.Since(t1),
			"contentLength": ww.BytesWritten(),
			"status":        status,
		}

		if cache := ww.Header().Get("x-cache"); cache != "" {
			postReqContent["cache"] = cache
		}

		logger = log.WithFields(postReqContent)
		if status >= 100 && status < 400 {
			logger.Info("request finished")
		} else if status == 500 {
			logger.WithField("stacktrace",
				string(debug.Stack())).Info("internal error during request")
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

func NewNewRelicMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		url := r.URL.String()
		path := r.URL.Path

		txn := newrelic.Application().StartTransaction(path)
		defer txn.End()

		txn.SetWebRequestHTTP(r)
		w = txn.SetWebResponse(w)

		txn.AddAttribute("request.url", fmt.Sprintf("http://%s%s", r.Host, url))

		qs := r.URL.Query()
		for key, value := range qs {
			txn.AddAttribute(key, strings.Join(value, "|"))
		}

		if reqID := middleware.GetReqID(ctx); reqID != "" {
			txn.AddAttribute("request.id", reqID)
		}

		nrCtx := nr.NewContext(ctx, txn)
		r = r.WithContext(nrCtx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
