package service

import (
	"github.com/tyerry/crud-go/src/configuration/rest_err"
	"github.com/tyerry/crud-go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}