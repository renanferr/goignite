package jlaffaye

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	rootftp "github.com/b2wdigital/goignite/pkg/transport/client/ftp"
	"github.com/jlaffaye/ftp"
	"gopkg.in/matryer/try.v1"
)

func NewClient(options *rootftp.Options) (*ftp.ServerConn, error) {

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

	options := rootftp.OptionsBuilder.
		Addr(config.String(rootftp.Addr)).
		User(config.String(rootftp.Username)).
		Password(config.String(rootftp.Password)).
		Timeout(config.Int(rootftp.Timeout)).
		Build()

	return NewClient(&options)

}
