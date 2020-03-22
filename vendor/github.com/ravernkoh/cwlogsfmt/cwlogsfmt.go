package cwlogsfmt

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

// CloudWatchLogsFormatter provides formatting similar to the default
// CloudWatch Logs format (from AWS Lambda). Uses the field Message to
// log the message.
//
//     // Examples of the format
//     START RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Version: $LATEST
//     PANIC RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Message: A serious crash
//     FATAL RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Message: A crash
//     ERROR RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Message: An error
//     WARNING RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Message: A warning
//     INFO RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Message: Some information
//     DEBUG RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Message: Some more information
//     END RequestId: 66389135-fd00-11e7-a1f9-8945479469b0
//     REPORT RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Duration: 0.96 ms Billed Duration: 100 ms Memory Size: 128 MB Max Memory Used: 24 MB
type CloudWatchLogsFormatter struct {
	// These fields will always be placed in front of the other fields, in the
	// order given in the slice. The log message will always be the first field
	// after the PrefixFields.
	PrefixFields []string

	// The fields are sorted by default for a consistent output. For applications
	// that log extremely frequently and don't use the JSON formatter this may not
	// be desired.
	DisableSorting bool

	// QuoteEmptyFields will wrap empty fields in quotes if true
	QuoteEmptyFields bool
}

// Format formats the given log entry.
func (f *CloudWatchLogsFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteByte(' ')

	for _, field := range f.PrefixFields {
		if val, exist := entry.Data[field]; exist {
			f.appendKeyValue(b, field, val)
			delete(entry.Data, field)
		}
	}

	keys := make([]string, 0, len(entry.Data))
	for key := range entry.Data {
		keys = append(keys, key)
	}

	if !f.DisableSorting {
		sort.Strings(keys)
	}

	f.appendKeyValue(b, "Message", entry.Message)

	for _, key := range keys {
		f.appendKeyValue(b, key, entry.Data[key])
	}

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *CloudWatchLogsFormatter) needsQuoting(text string) bool {
	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' ||
			ch == '.' ||
			ch == '_' ||
			ch == '/' ||
			ch == '@' ||
			ch == '^' ||
			ch == '+') {
			return true
		}
	}
	return false

}

func (f *CloudWatchLogsFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	b.WriteString(key)
	b.WriteString(": ")
	f.appendValue(b, value)
	b.WriteByte(' ')
}

func (f *CloudWatchLogsFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	if !f.needsQuoting(stringVal) {
		b.WriteString(stringVal)
	} else {
		b.WriteString(fmt.Sprintf("%q", stringVal))
	}
}
