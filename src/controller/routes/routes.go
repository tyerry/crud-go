package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tyerry/crud-go/src/controller"
)

func InitRoutes(route *gin.RouterGroup, userController controller.UserControllerInterface) {
	route.GET("/getUserById/:userId", userController.FindUserById)
	route.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	route.POST("/createUser", userController.CreateUser)
	route.PUT("/updateUser/:userId", userController.UpdateUser)
	route.DELETE("/deleteUser/:userId", userController.DeleteUser)
}