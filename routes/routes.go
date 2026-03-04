package routes

import (
	"todo-list/config"
	"todo-list/controllers"
	"todo-list/middleware" // 添加这行导入

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())

	db, _ := config.InitDB()
	todoController := controllers.NewTodoController(db)
	userController := controllers.NewUserController(db)

	v1 := r.Group("/api/v1")
	{
		// 用户相关路由
		users := v1.Group("/users")
		{
			users.POST("/register", userController.Create)
			users.POST("/login", userController.Login)
			users.GET("/info", middleware.AuthMiddleware(), userController.GetInfo)
		}

		// 待办事项路由，添加认证中间件
		todos := v1.Group("/todos", middleware.AuthMiddleware())
		{
			todos.POST("/", todoController.Create)
			todos.GET("/", todoController.GetAll)
			todos.GET("/:id", todoController.GetOne)
			todos.PUT("/:id", todoController.Update)
			todos.DELETE("/:id", todoController.Delete)
		}
	}

	return r
}
