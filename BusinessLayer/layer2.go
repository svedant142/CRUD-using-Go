package layer2

import (
	g "crud/getinfo"
	"crud/layer3"
	"fmt"
)

type IBusiness interface {
	CreatefromLayer3(int, string, int) bool
	DeletefromLayer3(int) bool
	UpdatefromLayer3(int, int) bool
	ReadfromLayer3() []g.Getuser
}

type BusinessLayer2 struct {
	repo layer3.IBusinessRepo
}

func NewBusinessLogic(r layer3.IBusinessRepo) IBusiness {

	return BusinessLayer2{repo: r}
}

func (l BusinessLayer2) CreatefromLayer3(id int, name string, salary int) bool {
	fmt.Println(id, name, salary)
	created := l.repo.CreateUser(id, name, salary)
	return created
}

func (l BusinessLayer2) ReadfromLayer3() []g.Getuser {
	a := l.repo.GetAllUserInfo()
	fmt.Println(a)
	return a
}

func (l BusinessLayer2) DeletefromLayer3(id int) bool {
	//	fmt.Println(id)
	created := l.repo.DeleteUser(id)
	return created
}

func (l BusinessLayer2) UpdatefromLayer3(id, salary int) bool {
	//	fmt.Println(id)
	created := l.repo.UpdateUser(id, salary)
	return created
}
