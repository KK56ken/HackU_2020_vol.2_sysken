package model

import (
	"database/sql"
	"log"
	"fmt"
)

type Task struct {
	Task_id       int
	User_id       int
	Subject_id    int
	Name          string
	Limit         string
	End_flag      int
	Register_date string
	End_date      string
}

type Tasks []Task

func SelectGettingTodo(user_id int) (*Tasks, error) {

	fmt.Println("11")

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	fmt.Println("22")

	// rows, err := db.Query("SELECT * FROM tasks WHERE user_id = ?", user_id)
	rows, err := db.Query("SELECT * FROM tasks")
	// if err == nil {
	// 	panic(err.Error())
	// 	return nil, err
	// }
	defer rows.Close()

	fmt.Println("33")

	var tasks Tasks

	for rows.Next() {
		task := Task{}
		if err = rows.Scan(&task.Task_id, &task.User_id ,&task.Subject_id, &task.Name, &task.Limit, &task.End_flag, &task.Register_date, &task.End_date); err != nil {
			log.Fatal(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	fmt.Println("44")
	return &tasks, nil
}
