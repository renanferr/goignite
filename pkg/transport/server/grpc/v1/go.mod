module github.com/b2wdigital/goignite/pkg/transport/server/grpc/v2

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../../config
	github.com/b2wdigital/goignite/pkg/log => ../../../../log
	github.com/b2wdigital/goignite/pkg/transport/server/grpc => ../
)

require (
	github.com/b2wdigital/goignite/pkg/log v0.0.0-00010101000000-000000000000 // indirect
	github.com/b2wdigital/goignite/pkg/transport/server/grpc v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
	google.golang.org/grpc v1.28.0
)
