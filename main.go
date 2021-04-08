package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	UserID    string
	StartDate string
	EndDate   string
	Title     string
	Status    string
}

type User struct {
	UserID   string
	Password string
}

func main() {
	todo := Todo{
		UserID:    "userid",
		StartDate: "2021-04-08",
		EndDate:   "2021-04-08",
		Title:     "test",
		Status:    "done",
	}

	// todo2 := Todo{
	// 	UserID:    "userid",
	// 	StartDate: "2021-04-09",
	// 	EndDate:   "2021-04-09",
	// 	Title:     "test2",
	// 	Status:    "done2",
	// }

	user := User{
		UserID:   "userid",
		Password: "password",
	}

	// err := insertTodo(todo)
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())
	// }

	// err = insertUser(user)
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())
	// }

	// updateTodo("userid", todo2)

	todos, err := selectTodo("userid")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println(todos)

	fmt.Println("todo: ", todo)
	fmt.Println("user: ", user)
}

func insertTodo(todo Todo) error {
	fmt.Println("insertTodo: ", todo)

	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/todo2")
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("DB Connected")

	_, err = db.Query(`
		insert into 
			todo 
		values (
			"` + todo.UserID + `", 
			"` + todo.StartDate + `", 
			"` + todo.EndDate + `", 
			"` + todo.Title + `", 
			"` + todo.Status + `"
		);`)

	if err != nil {
		return err
	}

	fmt.Println("data insert success")

	return nil
}

func insertUser(user User) error {
	fmt.Println("insertUser: ", user)

	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/todo2")
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("DB Connected")

	_, err = db.Query(`insert into user values ("` + user.UserID + `", "` + user.Password + `");`)

	if err != nil {
		return err
	}

	fmt.Println("data insert success")

	return nil
}

func updateTodo(userID string, todo Todo) error {
	fmt.Println("target: ", userID)
	fmt.Println("data: ", todo)

	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/todo2")
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("DB Connected")

	_, err = db.Query(`
		update 
			todo 
		set 
			user_id = "` + todo.UserID + `", 
			start_date="` + todo.StartDate + `", 
			end_date="` + todo.EndDate + `", 
			title="` + todo.Title + `",
			status="` + todo.Status + `" 
		where 
			user_id="` + userID + `"`)

	if err != nil {
		return err
	}

	fmt.Println("data insert success")

	return nil
}

func deleteTodo(userID string) error {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/todo2")
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("DB Connected")

	_, err = db.Query(`delete from todo where user_id="` + userID + `"`)

	if err != nil {
		return err
	}

	fmt.Println("data insert success")

	return nil
}

func selectTodo(userID string) ([]Todo, error) {
	db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/todo2")
	if err != nil {
		return []Todo{}, err
	}
	defer db.Close()

	fmt.Println("DB Connected")

	row, err := db.Query("select * from todo")

	if err != nil {
		return []Todo{}, err
	}
	fmt.Println("data selected")

	var todos []Todo

	for row.Next() {
		var data1, data2, data3, data4, data5 string
		err := row.Scan(&data1, &data2, &data3, &data4, &data5)
		if err != nil {
			return []Todo{}, err
		}
		fmt.Println("data1", data1)
		fmt.Println("data2", data2)
		fmt.Println("data3", data3)
		fmt.Println("data4", data4)
		fmt.Println("data5", data5)
		todos = append(todos, Todo{
			UserID:    data1,
			StartDate: data2,
			EndDate:   data3,
			Title:     data4,
			Status:    data5,
		})
	}

	return todos, nil
}
