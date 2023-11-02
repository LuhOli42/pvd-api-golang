package main

import (
	"github.com/gin-gonic/gin"
	initializers "luizafer.com/pvd/inicializers"
	"luizafer.com/pvd/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	app := gin.Default()

	routes.CategoriesRoutes(app)

	app.Run()
}
