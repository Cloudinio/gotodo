package main

import (
	"fmt"
	"time"
)
type Task struct {
	Id int
	Title string
	Description string
	Completed bool
	Created_at time.Time
	Due_to time.Time
	Priority int8
	Is_deleted bool
}

func NewTask(id int, title string, description string, completed bool, 
	created_at time.Time, due_to time.Time, priority int8) *Task {
	
	newTask := Task{
		Id: id,
		Title: title,
		Completed: false,
		Created_at: created_at,
		Due_to: due_to,
		Priority: priority,
		Is_deleted: false,
	}

	if description == "-" {
		newTask.Description = ""
	} else {
		newTask.Description = description
	}	

	return &newTask
}

func (t* Task) MarkAsDone() {
	if t != nil {
		t.Completed = true
	}
}

func (t* Task) MarkAsDeleted() {
	if t != nil {
		t.Is_deleted = true
	}
}

func ShowAllTasks(tasks *[]Task) {
	for _, task := range *tasks {
		fmt.Printf("Номер задачи: %d, название задачи: %s, дедлайн: %v\n, приоритет: %d", task.Id, task.Title, task.Due_to, task.Priority)
	}
}