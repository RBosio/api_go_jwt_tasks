package models

import (
	"packages/db"
	"time"
)

type Task struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	IdUser      int       `json:"id_user"`
}

type Tasks []Task

func GetTasks(idUser int) (Tasks, error) {
	sql := "SELECT * FROM task WHERE id_user = ?"
	if rows, err := db.Query(sql, idUser); err != nil {
		return nil, err
	} else {
		tasks := Tasks{}
		for rows.Next() {
			task := Task{}
			rows.Scan(&task.Id, &task.Description, &task.Completed, &task.CreatedAt, &task.IdUser)
			tasks = append(tasks, task)
		}
		return tasks, nil
	}
}

func GetTask(id int) (*Task, error) {
	sql := "SELECT * FROM task WHERE id = ?"
	if rows, err := db.Query(sql, id); err != nil {
		return nil, err
	} else {
		task := Task{}
		for rows.Next() {
			rows.Scan(&task.Id, &task.Description, &task.Completed, &task.CreatedAt, &task.IdUser)
		}
		return &task, nil
	}
}

func (task *Task) insert() {
	sql := "INSERT INTO task (description, id_user) VALUES (?, ?)"
	result, _ := db.Exec(sql, task.Description, task.IdUser)
	task.Id, _ = result.LastInsertId()
	task.CreatedAt = time.Now()
}

func (task *Task) update() {
	sql := "UPDATE task SET description = ?, completed = ? WHERE id = ?"
	db.Exec(sql, task.Description, task.Completed, task.Id)
}

func (task *Task) Save() {
	if task.Id == 0 {
		task.insert()
	} else {
		task.update()
	}
}

func (task *Task) Delete() {
	sql := "DELETE FROM task WHERE id = ?"
	db.Exec(sql, task.Id)
}
