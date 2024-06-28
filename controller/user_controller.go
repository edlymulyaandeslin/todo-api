package controller

import (
	"net/http"
	"todo-challange/service"

	"github.com/gin-gonic/gin"
)

// struct
type UserController struct {
	service service.UserService
	rg      *gin.RouterGroup
}

func (c *UserController) getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Route() {
	router := c.rg.Group("/users")
	router.GET("/:id", c.getUserById)
}

// constructor
func NewUserController(service service.UserService, rg *gin.RouterGroup) *UserController {
	return &UserController{
		service: service,
		rg:      rg,
	}
}
