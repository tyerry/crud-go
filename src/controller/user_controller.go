package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tyerry/crud-go/src/model/service"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserById(ctx *gin.Context)
	FindUserByEmail(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}