package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	db := createdb()
	defer db.Close()

	for {

		fmt.Println("1-- Enter user details")
		fmt.Println("2-- Display all the user details ")
		fmt.Println("3-- Enter project details")
		fmt.Println("6-- retrieve project details of specified project")
		fmt.Println("4-- Retrieve all the data in of projects  ")
		fmt.Println("5-- Retrieve only the detailss of specified person ")
		fmt.Println("7-- exit")
		fmt.Println("8 --would like to update details of userdetails ")
		fmt.Println("9-- would you like you to delete user detail ..")
		fmt.Println("ENTER YOUR CHOICE : ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			u, p, e, pn := details()

			er := inserting(db, u, p, e, pn)
			if er != nil {
				fmt.Println("error aat inserting user details.", er)
			}
		case 2:

			result := retrive(db)
			for _, v := range result {
				fmt.Println(v)
				//defer db.Close()
			}
		case 3:
			pn, tn, st, et := projectdetails()
			err := insertprojectd(db, pn, tn, st, et)
			if err != nil {
				log.Fatal("error inserting project details ", err)
			}
		case 4:
			result, err := retrieveprojectd(db)
			for _, v := range result {
				fmt.Println(v)
				//defer db.close()
			}
			if err != nil {
				log.Println(err)
			}

		case 5:
			fmt.Println("enter the user name to display all the user details :")
			var name string
			fmt.Scan(&name)

			result, _ := onlyuserdetails(db, name)
			for _, v := range result {
				fmt.Println(v)
			}
		case 6:
			var projectname string
			fmt.Println("enter which project detail would like to display :")
			fmt.Scan(&projectname)
			err, result := retrspecifiedpro(db, projectname)
			if err != nil {
				log.Println("errror while retriving data of projects ", err)
			}
			for _, d := range result {
				fmt.Println(d)
			}
			fmt.Println("data retrieved succesfully .")

		case 7:
			fmt.Println("exiting ..")
			db.Close()
			os.Exit(0)
		case 8:
			var userdata string
			fmt.Println("enter which data you would like to ")
			fmt.Scan(&userdata)
			fmt.Println("enter the new phnumvber ")
			var phno int64
			fmt.Scan(&phno)
			fmt.Println("enter the new mailid")
			var email string

			fmt.Scan(&email)

			updateuserphno(db, userdata, phno, email)
			fmt.Println("all the detailed are updated !!!")
		case 9:
			var username string
			fmt.Println("enter user name you want to delete ")
			fmt.Scan(&username)
			deletinguserdetails(db, username)

		}

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
	condetail := "user=santhosha dbname=mydb_1 password=santhosha@123 "
	db, err := sql.Open("postgres", condetail)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database !")
	return db
}

func inserting(db *sql.DB, username string, phno int, email string, pname string) error {
	rows := "insert into empdetails(username,phno,email,pname) values ($1,$2,$3,$4)"
	_, err := db.Exec(rows, username, phno, email, pname)
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
func projectdetails() (string, string, string, string) {
	var pname string
	var taskname string
	var starttime string
	var endtime string
	fmt.Println("enter the project name :")
	fmt.Scanln(&pname)
	fmt.Println("enter the task name :")
	fmt.Scanln(&taskname)
	fmt.Println("enter starttime :")
	fmt.Scanln(&starttime)
	fmt.Println("enter endtime : ")
	fmt.Scanln(&endtime)
	return pname, taskname, starttime, endtime
}

func insertprojectd(db *sql.DB, pn string, tn string, st string, et string) error {
	rows := "insert into projectdetails(pname ,taskname,starttime,endtime) values ($1,$2,$3,$4)"
	_, err := db.Exec(rows, pn, tn, st, et)
	if err != nil {
		return err

	} else {
		fmt.Println("inserted data sucesfully")
	}
	return nil
}

func retrieveprojectd(db *sql.DB) ([]string, error) {
	rows, err := db.Query("select * from projectdetails ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var pnme, tsknm, sttm, entm string
	var data []string
	for rows.Next() {
		err := rows.Scan(&pnme, &tsknm, &sttm, &entm)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, fmt.Sprintf("projectname : %s , taskname : %s , starttime : %s , endtime : %s ", pnme, tsknm, sttm, entm))

	}
	return data, nil

}

func onlyuserdetails(db *sql.DB, name string) ([]string, error) {
	rows, err := db.Query("select username, phno , email ,pname from empdetails where username =$1", name)
	if err != nil {
		return nil, err

	}
	defer rows.Close()
	var data []string
	for rows.Next() {
		var un string
		var phno int
		var email string
		var pname string
		err := rows.Scan(&un, &phno, &email, &pname)
		if err != nil {
			return nil, err
		}
		data = append(data, fmt.Sprintf("name : %s ,phno : %d ,email : %s ,pname : %s", un, phno, email, pname))

	}
	return data, nil

}

func retrspecifiedpro(db *sql.DB, pname string) (error, []string) {
	rows, err := db.Query("select pname ,taskname,starttime,endtime from projectdetails where pname =$1", pname)
	if err != nil {
		log.Fatal(err)
	}

	var data []string
	var fn, tknm, st, et string
	for rows.Next() {
		err := rows.Scan(&fn, &tknm, &st, &et)
		if err != nil {
			log.Println(err)
		}
		data = append(data, fmt.Sprintf("projectname : %s , taskname : %s , starttime : %s ,endtime : %s", fn, tknm, st, et))

	}
	return nil, data
}

func updateuserphno(db *sql.DB, username string, phno int64, email string) {
	_, err := db.Exec("update empdetails set phno = $1 , email=$2 where username=$3", phno, email, username)
	if err != nil {
		log.Fatal("the error is because of :", err.Error())
	} else {
		fmt.Println("updated successfully!")
	}

}

func deletinguserdetails(db *sql.DB, username string) {
	_, err := db.Exec("delete from empdetails where username =$1", username)
	if err != nil {
		log.Fatal(err)
	}

}
