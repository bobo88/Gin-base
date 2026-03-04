package controllers

import (
	"net/http"
	"todo-list/models"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{DB: db}
}

// 创建待办事项
// 修改 Create 方法示例
func (tc *TodoController) Create(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := tc.DB.Create(&todo).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建待办事项失败")
		return
	}

	utils.SuccessResponse(c, todo)
}

// 获取所有待办事项
func (tc *TodoController) GetAll(c *gin.Context) {
	var todos []models.Todo
	if err := tc.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取待办事项失败"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// 获取单个待办事项
func (tc *TodoController) GetOne(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := tc.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "待办事项不存在"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// 更新待办事项
func (tc *TodoController) Update(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := tc.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "待办事项不存在"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新待办事项失败"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// 删除待办事项
func (tc *TodoController) Delete(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := tc.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "待办事项不存在"})
		return
	}

	if err := tc.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除待办事项失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}