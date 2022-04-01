package main

import (
	"practice/cache"
	"practice/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	cache.New()
	r := setupRouter()
	_ = r.Run(":8000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	memberRepo := controllers.New()
	r.GET("/members", memberRepo.GetMembers)
	r.POST("/members", memberRepo.CreateMember)
	r.GET("/members/:id", memberRepo.GetMember)
	r.PUT("/members/:id", memberRepo.UpdateMember)
	r.DELETE("/members/:id", memberRepo.DeleteMember)

	return r
}
