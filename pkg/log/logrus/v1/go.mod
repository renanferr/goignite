module github.com/b2wdigital/goignite/pkg/log/logrus/v1

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../config
	github.com/b2wdigital/goignite/pkg/log => ../../../log
	github.com/b2wdigital/goignite/pkg/log/logrus => ../
)

require (
	github.com/b2wdigital/goignite/pkg/config v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/log v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/log/logrus v0.0.0-00010101000000-000000000000
	github.com/jpfaria/logrus-redis-hook v1.0.2
	github.com/ravernkoh/cwlogsfmt v0.0.0-20180121032441-917bad983b4c
	github.com/sirupsen/logrus v1.4.2
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
