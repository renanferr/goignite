module github.com/b2wdigital/goignite/pkg/transport/client/ftp/jlaffaye/v0

go 1.13

replace (
	github.com/b2wdigital/goignite/pkg/config => ../../../../../config
	github.com/b2wdigital/goignite/pkg/transport/client/ftp => ../../

)

require (
	github.com/b2wdigital/goignite/pkg/config v0.0.0-00010101000000-000000000000
	github.com/b2wdigital/goignite/pkg/transport/client/ftp v0.0.0-00010101000000-000000000000
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/jlaffaye/ftp v0.0.0-20200309171336-6841a2daa0d5
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2 // indirect
	gopkg.in/matryer/try.v1 v1.0.0-20150601225556-312d2599e12e
)
