module github.com/b2wdigital/goignite/pkg/transport/client/pubsub/nats/v1

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../../../config
	github.com/b2wdigital/goignite/pkg/health => ../../../../../health
	github.com/b2wdigital/goignite/pkg/log => ../../../../../log
	github.com/b2wdigital/goignite/pkg/transport/client/pubsub/nats => ../

)

require (
	github.com/b2wdigital/goignite/pkg/config v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/health v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/log v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/transport/client/pubsub/nats v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/nats-io/nats-server/v2 v2.1.4 // indirect
	github.com/nats-io/nats.go v1.9.1
)
