package controllers

import "github.com/gin-gonic/gin"

type BaseController struct {
    route *gin.RouterGroup
}

func NewBaseController(route *gin.RouterGroup) *BaseController {
    return &BaseController{
        route: route,
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
	c.route.GET("/", c.getAll)
	c.route.GET("/:id", c.getOne)
	c.route.POST("/", c.create)
	c.route.PUT("/:id", c.update)
	c.route.DELETE("/:id", c.delete)
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