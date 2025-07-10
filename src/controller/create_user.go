package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tyerry/crud-go/src/configuration/logger"
	"github.com/tyerry/crud-go/src/configuration/validation"
	"github.com/tyerry/crud-go/src/controller/model/request"
	"github.com/tyerry/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(ctx *gin.Context) {
	logger.Info("Init CreateUser controller",
	zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
		zap.String("journey", "createUser"),
		)

		restErr := validation.ValidateUserError(err)

		ctx.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
		)
		
	ctx.String(http.StatusOK, "")
}