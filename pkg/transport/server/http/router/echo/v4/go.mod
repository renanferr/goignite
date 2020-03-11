module github.com/b2wdigital/goignite/pkg/transport/server/http/router/echo/v4

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../../../../config
	github.com/b2wdigital/goignite/pkg/health => ../../../../../../health
	github.com/b2wdigital/goignite/pkg/log => ../../../../../../log
	github.com/b2wdigital/goignite/pkg/transport/server/http/router => ../../../router
)
