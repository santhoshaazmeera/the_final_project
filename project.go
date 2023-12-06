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

	inserting(db, u, p, e, pn)
	fmt.Println("enter which details you want to retrive :")
	
	result := retrive(db)
	for _, v := range result {
		fmt.Println(v)
		defer db.Close()
	}

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
	condetail := "user=postgres dbname=postgres password=test@123 "
	db, err := sql.Open("postgres", condetail)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database !")
	return db
}

func inserting(d *sql.DB, username string, phno int, email string, pname string) error {
	rows := "insert into empdetails(username,phno,email,pname) values ($1,$2,$3,$4)"
	_, err := d.Exec(rows, username, phno, email, pname)
	if err != nil {
		return err
	} else {
		fmt.Println("data inserted ")
	}
	return nil
}
func retrive(db *sql.DB) []string {
	rows, err := db.Query("select * from empdetails ")
	if err != nil {
		log.Fatal(err)
	}
	// rows.Close()

	var data []string
	for rows.Next() {
		var nme string
		var pn int64
		var eml string
		var pnme string
		err := rows.Scan(&nme, &pn, &eml, &pnme)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, fmt.Sprintf("name : %s,phone number : %d , email : %s ,project name : %s", nme, pn, eml, pnme))

	}
	return data

}
