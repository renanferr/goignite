module github.com/b2wdigital/goignite/pkg/cloud/aws/v2

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/cloud/aws => ../
	github.com/b2wdigital/goignite/pkg/config => ../../../config
	github.com/b2wdigital/goignite/pkg/log => ../../../log
)

require (
	github.com/aws/aws-sdk-go-v2 v0.19.0
	github.com/b2wdigital/goignite/pkg/cloud/aws v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/config v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/log v0.0.0-00010101000000-000000000000
	github.com/jmespath/go-jmespath v0.0.0-20200310193758-2437e8417af5 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/text v0.3.2 // indirect
)
