package layer1

import (
	layer2 "crud/BusinessLayer"
	g "crud/getinfo"
	"fmt"

	"github.com/gin-gonic/gin"
)

type layer1interface interface {
	Create(*gin.Context)
	Delete(*gin.Context)
	Update(*gin.Context)
	Read(c *gin.Context)
}
type layer1Struct struct {
	repo layer2.IBusiness
}

func Initialize(r2 layer2.IBusiness) layer1interface {
	return layer1Struct{repo: r2}

}

func (l1 layer1Struct) Create(c *gin.Context) {
	var req g.Request_to_create
	c.BindQuery(&req)
	fmt.Println(req.Newid, req.Newname, req.Newsalary)
	created := l1.repo.CreatefromLayer3(req.Newid, req.Newname, req.Newsalary)
	if created == true {
		c.JSON(200, gin.H{"status": "true"})
	} else {
		c.JSON(200, gin.H{"status": "false"})
	}

}

func (l1 layer1Struct) Read(c *gin.Context) {
	var res g.Response_to_read
	res.Info = l1.repo.ReadfromLayer3()
	fmt.Println("here now")
	fmt.Print(res.Info)
	c.JSON(200, res)
}

func (l1 layer1Struct) Update(c *gin.Context) {
	var req g.Request_to_update
	c.BindQuery(&req)
	//fmt.Println(req.Newid, req.Newname, req.Newsalary)
	created := l1.repo.UpdatefromLayer3(req.Newid, req.Newsalary)
	if created == true {
		c.JSON(200, gin.H{"status": "true"})
	} else {
		c.JSON(200, gin.H{"status": "false"})
	}

}
func (l1 layer1Struct) Delete(c *gin.Context) {
	var req g.Request_to_delete
	c.BindQuery(&req)
	//fmt.Println(req.Newid, req.Newname, req.Newsalary)
	created := l1.repo.DeletefromLayer3(req.Newid)
	if created == true {
		c.JSON(200, gin.H{"status": "true"})
	} else {
		c.JSON(200, gin.H{"status": "false"})
	}

}
