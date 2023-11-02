package main

import (
	initializers "luizafer.com/pvd/inicializers"
	"luizafer.com/pvd/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	categories := []model.Categories{
		{Description: "Informática"},
		{Description: "Celulares"},
		{Description: "Beleza e Perfumaria"},
		{Description: "Mercado"},
		{Description: "Livros e Papelaria"},
		{Description: "Brinquedos"},
		{Description: "Moda"},
		{Description: "Bebê"},
		{Description: "Games"},
	}

	initializers.DB.AutoMigrate(&model.Categories{})
	initializers.DB.Create(categories)
}
