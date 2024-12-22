package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BaseController struct {
    Route *gin.RouterGroup
	DB *gorm.DB
}

//set db
func (c *BaseController) SetDB(db *gorm.DB) {	
	c.DB = db
}

func NewBaseController(route *gin.RouterGroup) *BaseController {
    return &BaseController{
        Route: route,
    }
}

type CategoryController struct {
    *BaseController
}

func NewCategoryController(route *gin.RouterGroup) *CategoryController {
	categoryController := CategoryController{
        BaseController: NewBaseController(route),
    }
	categoryController.setup()
    return &categoryController
}

func (c *CategoryController) setup(){
	c.Route.GET("/", c.getAll)
	c.Route.GET("/:id", c.getOne)
	c.Route.POST("/", c.create)
	c.Route.PUT("/:id", c.update)
	c.Route.DELETE("/:id", c.delete)
}

func (c *CategoryController) getAll(ctx *gin.Context) {
    ctx.JSON(200, gin.H{"message": "Get all categories"})
}

func (c *CategoryController) getOne(ctx *gin.Context) {
    id := ctx.Param("id")
    ctx.JSON(200, gin.H{"message": "Get category by id", "id": id})
}

func (c *CategoryController) create(ctx *gin.Context) {
    var newCategory struct {
        Name string `json:"name"`
    }
    if err := ctx.BindJSON(&newCategory); err != nil {
        ctx.JSON(400, gin.H{"message": "Invalid request"})
        return
    }
    ctx.JSON(201, gin.H{"message": "Category created", "name": newCategory.Name})
}

func (c *CategoryController) update(ctx *gin.Context) {
    id := ctx.Param("id")
    var updatedCategory struct {
        Name string `json:"name"`
    }
    if err := ctx.BindJSON(&updatedCategory); err != nil {
        ctx.JSON(400, gin.H{"message": "Invalid request"})
        return
    }
    ctx.JSON(200, gin.H{"message": "Category updated", "id": id, "name": updatedCategory.Name})
}

func (c *CategoryController) delete(ctx *gin.Context) {
    id := ctx.Param("id")
    ctx.JSON(200, gin.H{"message": "Category deleted", "id": id})
}