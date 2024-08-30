package route

import (
	"github.com/gin-gonic/gin"
	"github.com/susilo001/Wallet-System/user/handler"
	"github.com/susilo001/Wallet-System/user/service"
)

func InitRoutes() {
	r := gin.Default()
 
	userService := service.NewUserService() // Assuming you have a NewUserService() constructor
	userHandler := handler.NewUserHandler(userService)

	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.Run(":8080")
}