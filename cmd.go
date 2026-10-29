package main

import (
	"flag"
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

type Arguments struct {
	Priority int
	Single int
	Filter string
	Due_to time.Time
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

func ParseFlags(args []string) (Arguments, error){
	result := Arguments{}

	fs := flag.NewFlagSet("task", flag.ContinueOnError)
	priorityPtr := fs.Int("priority", 0, "приоритет")
	priorityShortPtr := fs.Int("p", 0, "приоритет (коротко)")
	due_toPtr := fs.String("due_to", "", "дедлайн")
	due_toShortPtr := fs.String("d", "", "дедлайн (коротко)")
	filterPtr := fs.String("filter", "asc", "фильтр")
	filterShortPtr := fs.String("f", "asc", "фильтр (коротко)")
	singlePtr := fs.Int("single", 0, "номер задачи")
	singleShortPtr := fs.Int("s", 0, "номер задачи (коротко)")

	if err := fs.Parse(args); err != nil {
		err = fmt.Errorf("Error in parsing args: %w", err)
		return result, err
	}

	priority := *priorityPtr
	if *priorityShortPtr != 0 {
		priority = *priorityShortPtr
	}

	single := *singlePtr
	if *singleShortPtr != 0 {
		single = *singleShortPtr
	}

	filter := *filterPtr
	if *filterShortPtr != "asc" {
		filter = *filterShortPtr
	}
	if filter != "asc" && filter != "desc" {
		err := fmt.Errorf("Error in filter. Enter correct argument (asc/desc)")
		return result, err
	}

	due_toStr := *due_toPtr
	if *due_toShortPtr != "" {
		due_toStr = *due_toShortPtr
	}
	var due_to time.Time
	if due_toStr != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04", due_toStr)
		if err != nil {
				err = fmt.Errorf("неверный формат времени '%s'. Ожидается YYYY-MM-DD HH:MM", due_toStr)
				return result, err
			}
		due_to = parsedTime
	}

	result.Priority = priority
	result.Single = single
	result.Filter = filter
	result.Due_to = due_to

	return result, nil
}