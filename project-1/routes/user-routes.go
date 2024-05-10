package routes

import (
	"trademarkia/common"
	"trademarkia/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("signup", handlers.UserSignup)
	r.POST("login", handlers.UserLogin)
	r.GET("user-data", common.RetriveJwtToken, handlers.GetUserData)
	r.PATCH("edit", common.RetriveJwtToken, handlers.EditUserData)
	r.DELETE("delete", common.RetriveJwtToken, handlers.DeleteUserData)
	r.DELETE("logout", common.RetriveJwtToken, handlers.Logout)
}
