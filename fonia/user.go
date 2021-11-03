package mainx

import (
    "fmt"
    "github.com/beego/beego/v2/client/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type FnUser struct {
    Id   int
}

func init() {
    // set default database
    orm.RegisterDataBase("default", "mysql", "fonia:root@123@tcp(127.0.0.1:3306)/goadmin?charset=utf8")

    // register model
    orm.RegisterModel(new(FnUser))
}

func main() {
    o := orm.NewOrm()


    var user FnUser
	err := o.Raw("SELECT id FROM fn_user WHERE 1").QueryRow(&user)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}

	fmt.Printf()
}