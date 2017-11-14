package main

import (
	"flag"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"fmt"
)

var debug = flag.Bool("debug", true, "enable debugging")
var password = flag.String("password", "youhao", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "192.168.0.129", "the database server")
var user = flag.String("user", "sa", "the database user")
var database  = flag.String("database", "his", "the database name")

func GetDB() (*sql.DB, error){
	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable", *server, *database, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Print("PING:%s",err)
		return nil, err
	}
	return db, nil
}

func main()  {
	db, err := GetDB()

	fmt.Println(db, err)

	// 执行SQL语句
	rows, err1 := db.Query("select VAA01 from VAA1")
	if err1 != nil {
		fmt.Println("query: ", err, "---")
		return
	}
	for rows.Next() {
		var name string
		var number string
		rows.Scan(&name, &number)
		fmt.Printf("Name: %s \t Number: %s\n", name, number)
	}
}