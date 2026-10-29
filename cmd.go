package main

import (
	"fmt"
	"time"
)

const helptext string = `
Использование:
  todo <команда> [флаги]

Команды и доступыне для них флаги:
  add        добавить задачу
  list       показать список задач
  		-p, --priority		вывести задачи определенного приоритета
		-d, --due_to		вывести задачи до определенной даты включительно
		-f, --filter		вывести задачи по убыванию или возрастанию приоритета (desc/asc)
		-s, --single		вывести задачу с определенным номером
  done       отметить задачу выполненной
  delete     удалить задачу
  help       показать справку
`

type Task struct {
	Id int
	Title string
	Description string
	Completed bool
	Created_at time.Time
	Due_to time.Time
	Priority int
	Is_deleted bool
}

func NewTask(id int, title string, description string, 
	created_at time.Time, due_to time.Time, priority int) *Task {
	
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

func ShowAllTasks(tasks []Task) {
	for _, task := range tasks {
		fmt.Printf("Номер задачи: %d, название задачи: %s, дедлайн: %v, приоритет: %d\n", task.Id, task.Title, task.Due_to, task.Priority)
	}
}

func Help() (string) {
	return helptext
}

func UpdateTaskList() {

}