package DB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func SqlOpen() error {
	var err error
	db, err = sql.Open("postgres", "port=5432 user=postgres password=000824 dbname=CC_Reciteword sslmode=disable")
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;

	//还可以是这种方式打开
	//db, err := Sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	if err != nil {
		return err
	}
	return nil
}

func SqlSelect() {
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	println("-----------")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid = ", uid, "\nname = ", username, "\ndep = ", department, "\ncreated = ", created, "\n-----------")
	}
}
func SqlUpdate() {
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("ficow", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}
func SqlClose() {
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
