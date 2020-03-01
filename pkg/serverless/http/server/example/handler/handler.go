package handler

import (
	"net/http"

	"github.com/jpfaria/goignite/pkg/serverless/http/server/example/model/request"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Message string
}

func Test1(request *http.Request) (interface{}, error) {
	return Response{Message: "Test 1!!"}, nil
}

func Test2(req *http.Request, body interface{}) (interface{}, error) {

	user := body.(*request.User)

	log.Info(user.Name)

	return Response{Message: "Test 2!!"}, nil
}