package main

import (
	"log"
	"todo-list/config"
	"todo-list/routes"
)

func main() {
	// 初始化数据库
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("数据库初始化失败:", err)
	}
	
	// 确保在main函数结束时关闭数据库连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库实例失败:", err)
	}
	defer sqlDB.Close()

	// 初始化路由
	r := routes.SetupRouter()

	// 启动服务器
	r.Run(":9090")
}