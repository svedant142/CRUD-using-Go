package Getinfo

type Getuser struct {
	Id     int
	Name   string
	Salary int
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
	Info []Getuser
}

//commit
