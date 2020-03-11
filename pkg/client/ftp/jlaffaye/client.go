package jlaffaye

import (
	"time"

	ftp2 "github.com/b2wdigital/goignite/pkg/client/ftp"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/jlaffaye/ftp"
	"gopkg.in/matryer/try.v1"
)

func NewClient(options *ftp2.Options) (*ftp.ServerConn, error) {

	var conn *ftp.ServerConn

	err := try.Do(func(attempt int) (bool, error) {
		var err error
		conn, err = ftp.Dial(options.Addr, ftp.DialWithTimeout(time.Duration(options.Timeout)*time.Second))
		return attempt < options.Retry, err
	})
	if err != nil {
		return nil, err
	}

	err = conn.Login(options.User, options.Password)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewDefaultClient() (*ftp.ServerConn, error) {

	options := ftp2.OptionsBuilder.
		Addr(config.String(ftp2.Addr)).
		User(config.String(ftp2.Username)).
		Password(config.String(ftp2.Password)).
		Timeout(config.Int(ftp2.Timeout)).
		Build()

	return NewClient(&options)

}
