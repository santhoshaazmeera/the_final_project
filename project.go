package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	u, p, e, pn := details()

	db := createdb()
	defer db.Close()
	inserting(db, u, p, e, pn)

}
func details() (string, int, string, string) {
	var username string
	var phno int
	var email string
	var pname string
	fmt.Println("enter the user name :")
	fmt.Scan(&username)
	fmt.Println("enter phone number :")
	fmt.Scan(&phno)
	fmt.Println("enter email :")
	fmt.Scan(&email)
	fmt.Println("enter project name :")
	fmt.Scan(&pname)
	return username, phno, email, pname

}
func createdb() *sql.DB {
	condetail := "user=santhosha dbname=mydb_1 password=santhosha@123 "
	db, err := sql.Open("postgres", condetail)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database !")
	return db
}

func inserting(d *sql.DB, username string, phno int, email string, pname string) {
	rows := "insert into empdetails(username,phno,email) values ($1,$2,$3,$4)"
	_, err := d.Exec(rows, username, phno, email, pname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data inserted ")

}
