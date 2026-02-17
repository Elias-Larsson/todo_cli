package main

import (
	// "time"
	"fmt"
)

type Task struct {
	Completed bool
	Title     string
	// CreatedAt time.Time
}

func main() {
	var task1 Task

	task1.Completed = false
	task1.Title = "test if this works"

	fmt.Println(task1)
}
