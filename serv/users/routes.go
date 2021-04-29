package users

import (
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	grpUser := g.Group("/users")

	grpUser.GET("/test", PublishUserCreated)
	grpUser.GET("", GetUsers)
	// grpUser.GET(":id", GetUserByID)
	grpUser.POST("", CreateUser)
	grpUser.PUT("/:id", UpdateUser)
	grpUser.DELETE("/:id", DeleteUser)

}
