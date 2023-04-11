package config

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConn() *sql.DB {
	db, err := sql.Open("mysql", "root:twentyfour@/test?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ReadData(file *os.File) {
	db := CreateConn()
	db.Exec("drop table employee")
	db.Exec("create table employee(ID int, Name varchar(50), Dept varchar(50), Title varchar(50), Comp varchar(20))")

	csvFile, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvFile {

		_, err = db.Exec("insert into employee values(?,?,?,?,?)", line[0], line[1], line[2], line[3], line[4])
		if err != nil {
			fmt.Println(err)
		}

	}
}
