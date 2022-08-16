package routes

import (
	layer2 "crud/BusinessLayer"
	layer1 "crud/Layer1"
	layer3 "crud/layer3"

	"github.com/gin-gonic/gin"
)

func Running() {
	r3 := layer3.Initialize_database()

	r2 := layer2.NewBusinessLogic(r3)

	r1 := layer1.Initialize(r2)
	r := gin.Default()
	r.GET("api/create", r1.Create)
	r.GET("api/read", r1.Read)
	r.GET("api/update", r1.Update)
	r.GET("api/delete", r1.Delete)
	r.Run()
}

//Db.Close() where?
