module github.com/b2wdigital/goignite/examples/cloud/aws/s3

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/cloud/aws/v2 => ../../../../pkg/cloud/aws/v2
	github.com/b2wdigital/goignite/pkg/config => ../../../../pkg/config
	github.com/b2wdigital/goignite/pkg/log => ../../../../pkg/log
)

require (
	github.com/b2wdigital/goignite/pkg/log v0.0.0-00010101000000-000000000000 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20200310193758-2437e8417af5 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
)
