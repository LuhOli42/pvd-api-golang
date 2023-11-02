package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
	initializers "luizafer.com/pvd/inicializers"
	"luizafer.com/pvd/model"
)

func ListCategories(c *gin.Context) {
	var categories []model.Categories

	initializers.DB.Find(&categories)

	c.JSON(200, gin.H{
		"categories": categories,
	})
}

func CreateCategory(c *gin.Context) {
	var body struct {
		Description string
	}

	c.Bind(&body)

	if strings.TrimSpace(body.Description) == "" {
		c.JSON(400, gin.H{
			"error": "Please put a description",
		})
		return
	}

	category := model.Categories{Description: body.Description}

	result := initializers.DB.Create(&category)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"category": category,
	})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Description string
	}

	c.Bind(&body)

	if strings.TrimSpace(body.Description) == "" {
		c.JSON(400, gin.H{
			"error": "Please put a description",
		})
		return
	}

	var category model.Categories

	initializers.DB.First(&category, id)

	if category.ID == 0 {
		c.JSON(400, gin.H{
			"error": "category not found",
		})
	}

	initializers.DB.Model(&category).Updates(model.Categories{Description: body.Description})

	c.JSON(200, gin.H{
		"category": category,
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var category model.Categories
	initializers.DB.First(&category, id)

	if category.ID == 0 {
		c.JSON(400, gin.H{
			"error": "category not found",
		})
	}

	initializers.DB.Delete(&model.Categories{}, id)

	c.Status(201)
}
