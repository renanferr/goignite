module github.com/b2wdigital/goignite/pkg/log/zap/v1

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../config
	github.com/b2wdigital/goignite/pkg/log => ../../../log
	github.com/b2wdigital/goignite/pkg/log/zap => ../
)

require (
	github.com/b2wdigital/goignite/pkg/config v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/log v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/log/zap v0.0.0-00010101000000-000000000000
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.14.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200311035431-43e3193a9bc5 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)
