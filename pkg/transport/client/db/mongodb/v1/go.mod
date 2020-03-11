module github.com/b2wdigital/goignite/pkg/transport/client/db/mongodb/v1

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../../../config
	github.com/b2wdigital/goignite/pkg/health => ../../../../../health
	github.com/b2wdigital/goignite/pkg/log => ../../../../../log
	github.com/b2wdigital/goignite/pkg/transport/client/db/mongodb => ../
)

require (
	github.com/b2wdigital/goignite/pkg/config v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/health v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/log v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/transport/client/db/mongodb v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.3.1
)
