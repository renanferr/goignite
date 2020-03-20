package log

import "context"

// constant
const key = "log_fields"

func ToContext(ctx context.Context) context.Context {
	return l.ToContext(ctx)
	//
	//fields := logger.GetFields()
	//
	//ctxFields := fieldsFromContext(ctx)
	//
	//if ctxFields == nil {
	//	ctxFields = map[string]interface{}{}
	//}
	//
	//for k, v := range fields {
	//	ctxFields[k] = v
	//}
	//
	//return context.WithValue(ctx, key, ctxFields)
}

func FromContext(ctx context.Context) Logger {
	return l.FromContext(ctx)
	//fields := fieldsFromContext(ctx)
	//return l.WithFields(fields)
}

func fieldsFromContext(ctx context.Context) Fields {

	var fields Fields

	if ctx == nil {
		return Fields{}
	}

	if param := ctx.Value(key); param != nil {
		fields = ctx.Value(key).(Fields)
	}

	return fields
}
