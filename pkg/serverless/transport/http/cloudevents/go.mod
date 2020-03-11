module github.com/b2wdigital/goignite/pkg/serverless/transport/http/cloudevents

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../../config
	github.com/b2wdigital/goignite/pkg/log => ../../../../log
)

require (
	github.com/b2wdigital/goignite/pkg/config v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/log v0.0.0-00010101000000-000000000000
	github.com/cloudevents/sdk-go v1.1.2
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.opencensus.io v0.22.3 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.14.0 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200311035431-43e3193a9bc5 // indirect
)
