package layer1

import (
	layer2 "crud/BusinessLayer"
	g "crud/getinfo"
	layer3 "crud/layer3"
	"fmt"

	"github.com/gin-gonic/gin"
)

type layer1interface interface {
	Create(*gin.Context)
	Delete(*gin.Context)
	Update(*gin.Context)
	Read(c *gin.Context)
}
type layer1repo struct {
	repo layer2.IAddBusiness
}
type Request_to_create struct {
	Newid     int    `form:"id"`
	Newname   string `form:"name"`
	Newsalary int    `form:"salary"`
}
type Request_to_update struct {
	Newid     int `form:"id"`
	Newsalary int `form:"salary"`
}
type Request_to_delete struct {
	Newid int `form:"id"`
}
type Response_to_read struct {
	Info []g.Getuser
}

func Initialize() layer1interface {
	r3 := layer3.Initialize_database()

	//r3 = layer3.initialize_database()
	r2 := layer2.NewBusinessLogic(r3)
	return layer1repo{repo: r2}

}

func (l1 layer1repo) Create(c *gin.Context) {
	var req Request_to_create
	c.BindQuery(&req)
	fmt.Println(req.Newid, req.Newname, req.Newsalary)
	created := l1.repo.CreatefromLayer3(req.Newid, req.Newname, req.Newsalary)
	if created == true {
		c.JSON(200, gin.H{"status": "true"})
	} else {
		c.JSON(200, gin.H{"status": "false"})
	}

}

func (l1 layer1repo) Read(c *gin.Context) {
	var res Response_to_read
	res.Info = l1.repo.ReadfromLayer3()
	fmt.Println("here now")
	fmt.Print(res.Info)
	c.JSON(200, res)
}

func (l1 layer1repo) Update(c *gin.Context) {
	var req Request_to_update
	c.BindQuery(&req)
	//fmt.Println(req.Newid, req.Newname, req.Newsalary)
	created := l1.repo.UpdatefromLayer3(req.Newid, req.Newsalary)
	if created == true {
		c.JSON(200, gin.H{"status": "true"})
	} else {
		c.JSON(200, gin.H{"status": "false"})
	}

}
func (l1 layer1repo) Delete(c *gin.Context) {
	var req Request_to_delete
	c.BindQuery(&req)
	//fmt.Println(req.Newid, req.Newname, req.Newsalary)
	created := l1.repo.DeletefromLayer3(req.Newid)
	if created == true {
		c.JSON(200, gin.H{"status": "true"})
	} else {
		c.JSON(200, gin.H{"status": "false"})
	}

}
func Read() {

}
