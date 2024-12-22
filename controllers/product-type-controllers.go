package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/md-shadhin-mia/echo-inventory/initialize"
	"github.com/md-shadhin-mia/echo-inventory/models"
)

type ProductTypeController struct {
	*BaseController
	
}

func NewProductTypeController(route *gin.RouterGroup) *ProductTypeController {
	productType := ProductTypeController{
		BaseController: NewBaseController(route),
	}
	productType.SetDB(initialize.DB)
	productType.setup()

	return &productType
}

func (c *ProductTypeController) setup() {
	c.Route.GET("/", c.getAll)
	c.Route.GET("/:id", c.getOne)
	c.Route.POST("/", c.create)
	c.Route.PUT("/:id", c.update)
	c.Route.DELETE("/:id", c.delete)
}
func (c *ProductTypeController) getAll(ctx *gin.Context) {
	var productTypes []models.ProductType
	if err := c.DB.Find(&productTypes).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, productTypes)
}
func (c *ProductTypeController) getOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var productType models.ProductType
	if err := c.DB.First(&productType, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product type not found"})
		return
	}
	ctx.JSON(200, productType)
}

func (c *ProductTypeController) create(ctx *gin.Context) {
	var productType models.ProductType
	if err := ctx.ShouldBindJSON(&productType); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.DB.Create(&productType).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, productType)
}

func (c *ProductTypeController) update(ctx *gin.Context) {
	id := ctx.Param("id")
	var productType models.ProductType
	if err := c.DB.First(&productType, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product type not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&productType); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.DB.Save(&productType)
	ctx.JSON(200, productType)
}

func (c *ProductTypeController) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.DB.Delete(&models.ProductType{}, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product type not found"})
		return
	}
	ctx.JSON(204, nil)
}

