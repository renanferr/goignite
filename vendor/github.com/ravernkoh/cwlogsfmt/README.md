# cwlogsfmt
`cwlogsfmt` provides a formatter for [`logrus`](https://github.com/sirupsen/logrus) with the default format in CloudWatch Logs from AWS Lambda.

## Installation
Install via `go get`.
```bash
$ go get -u github.com/ravernkoh/cwlogsfmt
```

## Example
Use it with `logrus` to format logs.
```go
fmt := &cwlogsfmt.CloudWatchLogsFormatter{
	PrefixFields: []string{"RequestId"},
	QuoteEmptyFields: true,
}

log := &logrus.Logger{
	Formatter: fmt,
}

log.WithFields(logrus.Fields{
	"RequestId": "66389135-fd00-11e7-a1f9-8945479469b0",
	"OtherField": "Boom!",
}).Info("Hello!")
```
