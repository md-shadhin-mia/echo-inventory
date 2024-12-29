package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/md-shadhin-mia/echo-inventory/models"
	"github.com/md-shadhin-mia/echo-inventory/services"
	"github.com/md-shadhin-mia/echo-inventory/utils"
)

type CategoriesController struct {
	*BaseController
	categoryService *services.CategoryService
}

func NewCategoriesController(route *gin.RouterGroup) *CategoriesController {
	categoryController := CategoriesController{
		BaseController:  NewBaseController(route),
		categoryService: services.NewCategoryService(),
	}

	// func NewProductTypeController(route *gin.RouterGroup) *ProductTypeController {
	/* productType := ProductTypeController{
		BaseController: NewBaseController(route),
	} */
	// categoryController.SetDB(initialize.DB)
	categoryController.setup()

	return &categoryController
}

func (c *CategoriesController) setup() {
	c.Route.GET("/", c.getAll)
	c.Route.POST("", c.create)
	c.Route.GET("/:id", c.getOne)
	c.Route.PUT("/:id", c.update)
	c.Route.DELETE("/:id", c.delete)
}

func (c *CategoriesController) getAll(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	data, err := c.categoryService.GetAllCategories(limit, offset)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, gin.H{"data": data, "message": "success"})
}

func (c *CategoriesController) getOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	data, err := c.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"data": data, "message": "success"})
}

func (c *CategoriesController) create(ctx *gin.Context) {
	type CategoryRequest struct {
		Name string `json:"name" validate:"required,min=3,max=32"`
	}
	categoryReq := CategoryRequest{}
	err := ctx.BindJSON(&categoryReq)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	model := models.Category{}
	err = utils.ValidateAndTransfer(&categoryReq, &model)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	data, err := c.categoryService.CreateCategory(&model)
	log.Print(categoryReq.Name)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(201, gin.H{"data": data, "message": "success"})
}

func (c *CategoriesController) update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	type CategoryRequest struct {
		Name string `json:"name" validate:"required,min=3,max=32"`
	}
	categoryReq := CategoryRequest{}
	var category models.Category
	err = ctx.BindJSON(&categoryReq)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	err = utils.ValidateAndTransfer(&categoryReq, &category)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
		}
	data, err := c.categoryService.UpdateCategory(uint(id), &category)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"data": data, "message": "success"})
}

func (c *CategoriesController) delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	err = c.categoryService.DeleteCategory(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"message": "success"})
}
