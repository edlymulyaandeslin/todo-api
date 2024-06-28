package controller

import (
	"fmt"
	"net/http"
	"todo-challange/model/dto"
	"todo-challange/service"

	"github.com/gin-gonic/gin"
)

// struct
type TaskController struct {
	service service.TaskService
	rg      *gin.RouterGroup
}

func (c *TaskController) getTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.service.FindById(id)
	fmt.Println("running controller")
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) getAllTask(ctx *gin.Context) {
	tasks, err := c.service.FindAllTask()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(tasks) < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "tasks not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) createTaskHandler(ctx *gin.Context) {
	var payload dto.TaskRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	response, err := c.service.CreateNewTask(payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *TaskController) updateTaskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var payload dto.TaskUpdated

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	response, err := c.service.UpdatedTask(id, payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Updated successfully",
		"data":    response,
	})
}

func (c *TaskController) deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.service.DeleteTask(id)
	if err != nil {

		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "deleted successfully",
	})
}

func (c *TaskController) Route() {
	router := c.rg.Group("/todos")
	router.GET("/", c.getAllTask)
	router.GET("/:id", c.getTaskById)
	router.POST("/", c.createTaskHandler)
	router.PUT("/:id/update", c.updateTaskHandler)
	router.DELETE("/:id", c.deleteTask)
}

// constructor
func NewTaskController(service service.TaskService, rg *gin.RouterGroup) *TaskController {
	return &TaskController{
		service: service,
		rg:      rg,
	}
}
