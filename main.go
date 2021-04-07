package main
 
import (
    "database/sql" 
    _ "github.com/go-sql-driver/mysql"
    "log"
	"fmt"

	"encoding/json"
)


type Todo struct {
	UserID string
	StartDate string
	EndDate string
	Title string
	Status string
}

type User struct {
	UserID string
	Password string
}

func insertTodo(todo Todo) {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	result, err := db.Exec(`
		insert into todo2.todo  
		values (
			"` + todo.UserID + `", 
			"` + todo.StartDate + `",
			"` + todo.EndDate + `",
			"` + todo.Title + `",
			"` + todo.Status + `"
		)`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func insertUser(user User) {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	result, err := db.Exec(`
		insert into todo2.user 
		values (
			"` + user.UserID + `", 
			"` + user.Password + `"
		)`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	todo := Todo {
		UserID: "userid",
		StartDate: "2021-04-07",
		EndDate: "2021-04-07",
		Title: "userid",
		Status: "ready",
	}
	// err := updateTodo("userid", todo)
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())
	// }
	insertTodo(todo)

	// err := selectTodo("userID")
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())
	// }

	// err := deleteTodo("userid")
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())
	// }


	// insertUser(User{
	// 	UserID: "userid",
	// 	Password: "password",
	// })
	// err := updateUser("userid", User{
	// 	UserID: "userid",
	// 	Password: "password2",
	// })
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())
	// }

	datas, err := selectTodo("userid")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}

	fmt.Println("datas: ", datas)


	// err = deleteUser("userid")
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())
	// }

	j, _ := json.Marshal(&todo)
	fmt.Println(string(j))

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
	// dropData()
}

func selectTodo(userID string) ([]Todo, error){
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        return nil, err
    }
    defer db.Close()


	rows, err := db.Query("select * from todo2.todo")
	if err != nil {
		return nil, err
	}

	todos := make([]Todo, 0)

	defer rows.Close()
	for rows.Next() {
		todo := Todo{}

		err := rows.Scan(&todo.UserID, &todo.StartDate, &todo.EndDate, &todo.Title, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	fmt.Println(todos)

	return todos, nil
}

func MysqlExec(query string) error{
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        return err
    }
    defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func updateTodo(userID string, todo Todo) error {
	err := MysqlExec(`
		UPDATE todo2.todo
		SET 
			user_id = "` + todo.UserID + `",
			start_date = "` + todo.StartDate + `",
			end_date = "` + todo.EndDate + `",
			title = "` + todo.Title + `",
			status = "` + todo.Status + `"
		WHERE
			user_id = "` + userID + `"`)

	if err != nil {
		return err
	}

	return nil
}

func updateUser(userID string, user User) error {
	err := MysqlExec(`
		UPDATE todo2.user
		SET 
			user_id = "` + user.UserID + `",
			password = "` + user.Password + `"
		WHERE
			user_id = "` + user.UserID + `"`)

	if err != nil {
		return err
	}

	return nil
}

func deleteTodo(userID string) error {
	err := MysqlExec(`
		DELETE FROM 
			todo2.todo
		WHERE
			user_id = "` + userID + `"`)

	if err != nil {
		return err
	}

	return nil
}

func deleteUser(userID string) error {
	err := MysqlExec(`
		DELETE FROM 
			todo2.user
		WHERE
			user_id = "` + userID + `"`)

	if err != nil {
		return err
	}

	return nil
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