package main

import (
	"database/sql"
	"fmt" 
)

func main() {
	//conn, err := sql.Open("odbc", "driver={Microsoft Access Driver (*.mdb)};dbq=d:\\test.mdb")
	conn, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")
	if err != nil {
		fmt.Println("Connecting Error")
		fmt.Printf("%v\n",err)
		return
	}
	defer conn.Close()
	stmt, err := conn.Prepare("select * from test") //ALTER TABLE tb ALTER COLUMN aa Long
	if err != nil {
		fmt.Println("Query Error")
		return
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		fmt.Print(err)
		fmt.Println("Query Error")
		return
	}
	defer row.Close()
	for row.Next() {
		var ID string
		var SequenceNumber int
		var ValueCode string
		if err := row.Scan(&ID, &SequenceNumber, &ValueCode); err == nil {
			fmt.Println(ID, SequenceNumber, ValueCode)
		}
	}
	fmt.Printf("%s\n", "finish")
	return
}
