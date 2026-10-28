package main

import (
	// "encoding/json"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	tasks := make([]Task, 5)	
	for {
		reader := bufio.NewReader(os.Stdin)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "todo add":
			fmt.Println("Enter some details")

			fmt.Print("Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace((description))

			fmt.Print("Due to: ")
			due_toStr, _ := reader.ReadString('\n')
			due_toStr = strings.TrimSpace(due_toStr)
			due_to, err := time.Parse("", due_toStr)
			if err != nil {
				err = fmt.Errorf("Error in parsing time. Enter time correctly. Returning to main menu")
				fmt.Println(err)
				continue
			}

			fmt.Print("Priority: ")
			priorityStr, _ := reader.ReadString('\n')
			priorityStr = strings.TrimSpace(priorityStr)
			priority, err := strconv.Atoi(priorityStr)
			if err != nil {
				err = fmt.Errorf("Error in parsing priority. Enter priority correctly. Returning to main menu")
				fmt.Println(err)
				continue
			}


			task := NewTask(len(tasks) + 1, title, description, time.Now(), due_to, priority)
			tasks = append(tasks, *task)
		
			fmt.Println("Task created successfully")

		case "todo list":
			ShowAllTasks(tasks)

		case "todo done":
			fmt.Println("ID of task: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(idStr)
			if err != nil {
				err = fmt.Errorf("Error in parsing ID. Enter id correctly")
			}

			task := tasks[id]
			task.MarkAsDone()

		case "todo delete":
			fmt.Println("ID of task: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(idStr)
			if err != nil {
				err = fmt.Errorf("Error in parsing ID. Enter id correctly")
			}
			task := tasks[id]
			task.MarkAsDeleted()
		default:
			fmt.Println("Uknown command. Type 'commands' to see all commands and flags")
		}
	}
}