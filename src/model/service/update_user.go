package service

import (
	"github.com/tyerry/crud-go/src/configuration/rest_err"
	"github.com/tyerry/crud-go/src/model"
)

func (*userDomainService) UpdateUser(
	userId string, 
	userDomainInterface model.UserDomainInterface,
	) *rest_err.RestErr {
	return nil
}