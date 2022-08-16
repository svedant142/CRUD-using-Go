package layer3

import (
	g "crud/getinfo"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type IBusinessRepo interface {
	CreateUser(int, string, int) bool
	GetAllUserInfo() []g.Getuser
	DeleteUser(int) bool
	UpdateUser(int, int) bool
	
}
type DbLayer3 struct {
	Db *sql.DB
}

func Initialize_database() IBusinessRepo {
	datab, err_db := sql.Open("mysql", "root:1234@123@tcp(127.0.0.1:3306)/dotDB")
	if err_db != nil {
		fmt.Println("error accessing database")
	}
	d := DbLayer3{datab}
	return d
}
func (l DbLayer3) GetAllUserInfo() []g.Getuser {

	rows, err_row := l.Db.Query("SELECT * FROM dotDB.new_users")
	if err_row != nil {
		fmt.Println(err_row)
		log.Fatal(err_row)
	}
	defer rows.Close()
	var get []g.Getuser

	var t g.Getuser
	for rows.Next() {

		err_col := rows.Scan(&t.Id, &t.Name, &t.Salary)
		if err_col != nil {
			fmt.Println("error while accessing row values")
			log.Fatal(err_col)
		}
		get = append(get, t)
	}
	err2 := rows.Err()
	if err2 != nil {
		fmt.Println("error after accessing row values")
		log.Fatal(err2)
	}
	fmt.Print(get)
	return get
}

func (l DbLayer3) CreateUser(id int, name string, salary int) bool {
	stmt, err := l.Db.Prepare("INSERT INTO dotDB.new_users VALUES(?,?,?)")
	fmt.Println(id, name, salary)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err1 := stmt.Exec(id, name, salary)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	return true
}

func (l DbLayer3) UpdateUser(id, salary int) bool {
	stmt, err := l.Db.Prepare("UPDATE dotDB.new_users SET salary = ? where id = ?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err1 := stmt.Exec(salary, id)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	return true
}
func (l DbLayer3) DeleteUser(id int) bool {
	stmt, err := l.Db.Prepare("DELETE FROM dotDB.new_users WHERE ID=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err1 := stmt.Exec(id)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	return true
}
