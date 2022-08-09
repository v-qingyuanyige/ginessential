package main

import (
	"ginEssential/common"
	"ginEssential/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化并连接数据库
	common.InitDB()
	db := common.GetDB()
	// 延迟关闭数据库
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	r := gin.Default()
	r = routes.CollectRoute(r)
	// 监听并在 0.0.0.0:8080 上启动服务
	panic(r.Run(":8080"))
}
