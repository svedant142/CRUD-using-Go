package routes

import (
	layer1 "crud/Layer1"

	"github.com/gin-gonic/gin"
)

func Running() {
	r1 := layer1.Initialize()
	r := gin.Default()
	r.GET("api/create", r1.Create)
	r.GET("api/read", r1.Read)
	r.GET("api/update", r1.Update)
	r.GET("api/delete", r1.Delete)
	r.Run()
}
