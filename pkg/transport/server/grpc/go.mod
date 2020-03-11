module github.com/b2wdigital/goignite/pkg/transport/server/grpc/v1

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../../../config
	github.com/b2wdigital/goignite/pkg/log => ../../../../../log
)

require (
	github.com/b2wdigital/goignite v0.0.17
	github.com/golang/protobuf v1.3.4
)
