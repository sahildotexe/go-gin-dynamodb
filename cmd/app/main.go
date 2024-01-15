package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sahildotexe/go-gin-dynamodb/config"
	"github.com/sahildotexe/go-gin-dynamodb/internal/routes"
	"github.com/sahildotexe/go-gin-dynamodb/utils/env"
)

func main() {
	env.LoadEnvVariables()
	config.ConnectDB()
	r := gin.New()
	routes.UserRoutes(r)
	r.Run()
}
