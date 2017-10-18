package h_db

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-oci8"
)

func main_test() {
	fmt.Println("oracle driver example")

	db, err := sql.Open("oci8", "song/123456@192.168.0.105:1521/ORCL")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err1 := db.Query("select 3.14, 'foo' from dual")
	if err1 != nil {
		fmt.Println(err1)
	}

	for rows.Next() {
		var f1 float64
		var f2 string
		rows.Scan(&f1, &f2)
		fmt.Println(f1, "heheeh", f2)
	}
	rows.Close()
}