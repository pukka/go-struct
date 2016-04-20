package main

import (
	"fmt"
)

type ID int
type TaskID int

type Task struct {
	Code   int
	Detail string
	done   bool
}

func ProcessTask(id ID, taskId TaskID) {
	fmt.Printf("id: %v\n", id)
	fmt.Printf("taskId: %v\n", taskId)
}

func main() {
	var id ID = 3
	var taskId TaskID = 5
	ProcessTask(id, taskId)

	var task Task = Task{
		Code:   1,
		Detail: "enjoy go",
		done:   true,
	}

	fmt.Println(task.Code)
	fmt.Println(task.Detail)
	fmt.Println(task.done)
}
