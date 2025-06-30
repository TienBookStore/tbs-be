package main

import (
	"backend/internal/common"
	"backend/internal/config"
	"backend/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Không thể tải cấu hình: ", err)
	}
	container, err := common.NewContainer(cfg)
	if err != nil {
		log.Fatal("Khởi tạo vùng chứa thất bại: ", err)
	}
	r := gin.Default()

	config.SetUpCors(r)
	
	api := r.Group("/api/bao-tien")
	router.SetupAuthRoute(api, container.AuthHandler)

	log.Printf("Server đang chạy ở cổng: %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Lỗi chạy server:", err)
	}
}
