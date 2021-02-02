package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkerr(err)

	stmt, err := db.Prepare("insert userinfo SET username=?, departname=?,created=?")
	checkerr(err)

	res, err := stmt.Exec("William", "Computer Science Guy", "2019-08-24")
	checkerr(err)

	id, err := res.LastInsertId()
	checkerr(err)
	fmt.Println(id)

	stmt, err = db.Prepare("update userinfo SET username=? where uid=?")
	checkerr(err)

	res, err = stmt.Exec("William-updated", id)
	checkerr(err)

	affect, err := res.RowsAffected()
	checkerr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * from userinfo")
	checkerr(err)

	for rows.Next() {
		var uid int
		var username, department, created string
		err := rows.Scan(&uid, &username, &department, &created)
		checkerr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)

	}

	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkerr(err)

	res, err = stmt.Exec(id)
	checkerr(err)

	db.Close()
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
