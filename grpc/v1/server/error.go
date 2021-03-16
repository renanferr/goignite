package gigrpc

import (
	gierrors "github.com/b2wdigital/goignite/v2/errors"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Error(err error) error {

	if gierrors.IsNotFound(err) {
		return status.Errorf(codes.NotFound, err.Error())
	} else if gierrors.IsNotValid(err) || gierrors.IsBadRequest(err) {
		return status.Errorf(codes.InvalidArgument, err.Error())
	} else if gierrors.IsServiceUnavailable(err) {
		return status.Errorf(codes.Unavailable, err.Error())
	} else {
		switch t := err.(type) {
		case validator.ValidationErrors:
			return status.Errorf(codes.InvalidArgument, t.Error())
		default:
			return status.Errorf(codes.Internal, t.Error())
		}
	}
}
