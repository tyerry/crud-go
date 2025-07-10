package view

import (
	"github.com/tyerry/crud-go/src/controller/model/response"
	"github.com/tyerry/crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomainInterface model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID: 	"",
		Email: 	userDomainInterface.GetEmail(),
		Name:	userDomainInterface.GetName(),
		Age:	userDomainInterface.GetAge(),
	}
}