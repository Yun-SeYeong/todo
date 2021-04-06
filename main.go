package main
 
import (
    "database/sql" 
    _ "github.com/go-sql-driver/mysql"
    "log"
	"fmt"
)
 
func main() {
    // db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer db.Close()

	// var host []uint8
    // var user string


    // rows, err := db.Query("select host, user from mysql.user;")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer rows.Close()
	
    // for rows.Next() {
    //     err := rows.Scan(&host, &user)
    //     if err != nil {
    //         log.Fatal(err)
    //     }
    //     fmt.Println(string(host), user)
    // }

	// createDatabase()
	// createTable()
	// insertData()
	// selectData()
	dropData()
}

func createDatabase() {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	result, err := db.Exec("create database todo2")
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("n: ", n)
}

func createTable() {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	result, err := db.Exec(`
		create table todo2.users (
			user_id varchar(10),
			password varchar(10)
		)`)

	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("n: ", n)
}

func insertData() {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	result, err := db.Exec(`
		insert into todo2.users (
			user_id, 
			password
		) values (
			"user1", 
			"pwd1"
		)`)

	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("n: ", n)
}

func selectData() {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	var host []uint8
    var user string


    rows, err := db.Query("select * from todo2.users;")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
	
    for rows.Next() {
        err := rows.Scan(&host, &user)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(host), user)
    }
}

func dropData() {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	result, err := db.Exec(`drop database todo2`)

	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("n: ", n)
}