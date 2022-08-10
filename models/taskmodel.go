package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ahmjaelani/task_crud_golang/config"
	"github.com/ahmjaelani/task_crud_golang/entities"
)

type TaskModel struct {
	conn *sql.DB
}

func NewTaskModel() *TaskModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &TaskModel{
		conn: conn,
	}
}

func (p *TaskModel) FindAll() ([]entities.Task, error) {

	rows, err := p.conn.Query("select * from task")
	if err != nil {
		return []entities.Task{}, err
	}
	defer rows.Close()

	var dataTask []entities.Task
	for rows.Next() {
		var task entities.Task
		rows.Scan(&task.Id,
			&task.InputTask,
			&task.Name,
			&task.Deadline)

		// 2006-01-02 => yyyy-mm-dd
		deadline, _ := time.Parse("2006-01-02", task.Deadline)
		// 02-01-2006 => dd-mm-yyyy
		task.Deadline = deadline.Format("02-01-2006")

		dataTask = append(dataTask, task)
	}

	return dataTask, nil
}

func (p *TaskModel) Create(task entities.Task) bool {
	result, err := p.conn.Exec("insert into task (input_task, name, deadline) values(?,?,?)",
		task.InputTask, task.Name, task.Deadline)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0

}

func (p *TaskModel) Find(id int64, task *entities.Task) error {

	return p.conn.QueryRow("select * from task where id = ?", id).Scan(
		&task.Id,
		&task.InputTask,
		&task.Name,
		&task.Deadline)
}

func (p *TaskModel) Update(task entities.Task) error {

	_, err := p.conn.Exec(
		"update task set input_task = ?, name = ?, deadline = ? where id = ?",
		task.InputTask, task.Name, task.Deadline, task.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *TaskModel) Delete(id int64) {
	p.conn.Exec("delete from task where id = ?", id)
}
