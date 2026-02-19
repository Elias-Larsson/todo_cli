package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task",
	Long:  `Add a new task to your todo list`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskDescription := args[0]
		task, err := createTask(taskDescription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task added: %s\n", task.Title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func createTask(taskString string) (Task, error) {
	task := Task{
		Title:     taskString,
		Completed: "false",
		CreatedAt: time.Now(),
	}
	file, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return Task{}, err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{task.Title, task.Completed, task.CreatedAt.Format(time.RFC3339)})

	if err != nil {
		return Task{}, err
	}

	return task, nil
}
