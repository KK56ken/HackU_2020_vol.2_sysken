package model

import (
	"database/sql"
	"errors"
	"log"
)

type Task struct {
	TaskId       int
	UserId       int
	SubjectId    int
	Name         string
	Limit        string
	EndFlag      int
	RegisterDate string
	EndDate      string
}

type Tasks []*Task

func SelectGettingTodo(user_id int) (Tasks, error) {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM tasks")
	//rows, err := db.Query("SELECT * FROM tasks WHERE user_id = ?", user_id)
	if err != nil {
		return nil, err
	}
	return convertToTodos(rows)
}

func convertToTodos(rows *sql.Rows) (Tasks, error) {
	var tasks Tasks
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.TaskId, &task.UserId, &task.SubjectId, &task.Name, &task.Limit, &task.EndFlag, &task.RegisterDate, &task.EndDate); err != nil {
			log.Println(errors.New("scan failed"))
			return nil, err
		}
		//要素を追加
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		log.Println(errors.New("rows scan failed"))
		return nil, err
	}
	return tasks, nil
}
func InsertTodo(record *Task) error {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	// userテーブルへのレコードの登録を行うSQLを入力する
	stmt, err := db.Prepare("INSERT INTO tasks (user_id, subject_id,name, time_limit, end_flag , end_date ) VALUES ( ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.UserId, record.SubjectId, record.Name, record.Limit, record.EndFlag, record.EndDate)
	return err
}
