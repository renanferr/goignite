package jlaffaye

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	. "github.com/b2wdigital/goignite/pkg/ftp/client/config"
	. "github.com/b2wdigital/goignite/pkg/ftp/client/model"
	"github.com/jlaffaye/ftp"
	"gopkg.in/matryer/try.v1"
)

func NewClient(options *Options) (*ftp.ServerConn, error) {

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

	options := OptionsBuilder.
		Addr(config.String(Addr)).
		User(config.String(Username)).
		Password(config.String(Password)).
		Timeout(config.Int(Timeout)).
		Build()

	return NewClient(&options)

}
