package gijlaffaye

import (
	"context"
	"time"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/jlaffaye/ftp"
	"gopkg.in/matryer/try.v1"
)

func NewServerConn(options *Options) (*ftp.ServerConn, error) {

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

func NewDefaultServerConn(ctx context.Context) (*ftp.ServerConn, error) {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewServerConn(o)

}
