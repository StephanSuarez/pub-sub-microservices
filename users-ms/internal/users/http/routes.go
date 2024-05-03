package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, uDep *UsersDependencies) {
	routesUsers := r.Group("/v1/users")

	routesUsers.POST("", uDep.uh.CreateUser)
	routesUsers.GET("", uDep.uh.GetUsers)
	routesUsers.GET(":id", uDep.uh.GetUserByID)
	routesUsers.PUT(":id", uDep.uh.UpdateUser)
	routesUsers.DELETE(":id", uDep.uh.DeleteUser)
	routesUsers.GET("74abc", uDep.uh.GetUserByUserNameOrEmail)

}
