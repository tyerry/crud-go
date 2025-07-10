package service

import (
	"github.com/tyerry/crud-go/src/configuration/logger"
	"github.com/tyerry/crud-go/src/configuration/rest_err"
	"github.com/tyerry/crud-go/src/model"
	"go.uber.org/zap"
)

func (userDomain *userDomainService) CreateUser(
	UserDomainInterface model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey","createUser"))	

	UserDomainInterface.EncryptPassword()
	return nil
}