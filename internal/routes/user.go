package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahildotexe/go-gin-dynamodb/internal/controllers"
)

func UserRoutes(r *gin.Engine) {
	public := r.Group("/api")
	public.GET("/user", controllers.DescribeTable)

}
