package logrus

import (
	"context"

	"github.com/sirupsen/logrus"
)

var key = "logrus_fields"

func ToContext(ctx context.Context, entry *logrus.Entry) context.Context {

	fields := entry.Data

	ctxFields := fieldsFromContext(ctx)

	if ctxFields == nil {
		ctxFields = map[string]interface{}{}
	}

	for k, v := range fields {
		ctxFields[k] = v
	}

	return context.WithValue(ctx, key, ctxFields)
}

func FromContext(ctx context.Context) *logrus.Entry {
	fields := fieldsFromContext(ctx)
	return logrus.WithContext(ctx).WithFields(fields)
}

func fieldsFromContext(ctx context.Context) logrus.Fields {

	var fields logrus.Fields

	if ctx == nil {
		return logrus.Fields{}
	}

	if param := ctx.Value(key); param != nil {
		fields = ctx.Value(key).(logrus.Fields)
	}

	return fields
}
