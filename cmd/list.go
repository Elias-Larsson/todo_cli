package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  "List all tasks in your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := listTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing tasks: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(list)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTasks() (string, error) {
	file, err := os.OpenFile("tasks.csv", os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return "Tasks:\n(no tasks yet)\n", nil
		}
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	if len(records) == 0 {
		return "Tasks:\n(no tasks yet)\n", nil
	}

	var builder strings.Builder
	builder.WriteString("Tasks:\n")
	for i, record := range records {
		title := ""
		completed := "false"
		if len(record) > 0 {
			title = record[0]
		}
		if len(record) > 1 {
			completed = record[1]
		}

		mark := " "
		if strings.EqualFold(completed, "true") {
			mark = "x"
		}

		builder.WriteString(fmt.Sprintf("%d. [%s] %s\n", i+1, mark, title))
	}

	return builder.String(), nil
}
