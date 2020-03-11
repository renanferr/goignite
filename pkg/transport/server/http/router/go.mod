module github.com/b2wdigital/goignite/pkg/transport/server/http/router

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../../config
	github.com/b2wdigital/goignite/pkg/health => ../../../../health
	github.com/b2wdigital/goignite/pkg/info => ../../../../info
)

require (
	github.com/b2wdigital/goignite/pkg/config v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/health v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/info v0.0.0-00010101000000-000000000000
	github.com/go-playground/validator/v10 v10.2.0
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
)
