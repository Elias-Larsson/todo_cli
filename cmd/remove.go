package cmd

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"

    "github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
    Use:   "remove [task id]",
    Short: "Remove a task",
    Long:  `Remove a task from your todo list by its ID`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        taskID := args[0]
        err := removeTask(taskID)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error removing task: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("Task %s removed successfully\n", taskID)
    },
}

func init() {
    rootCmd.AddCommand(removeCmd)
}

func removeTask(taskID string) error {
    id, err := strconv.Atoi(taskID)
    if err != nil {
        return fmt.Errorf("invalid task ID: %s", taskID)
    }
    file, err := os.OpenFile("tasks.csv", os.O_RDONLY, 0644)
    if err != nil {
        if os.IsNotExist(err) {
            return fmt.Errorf("no tasks file found")
        }
        return err
    }

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    file.Close()
    if err != nil {
        return err
    }

    if len(records) == 0 {
        return fmt.Errorf("no tasks to remove")
    }

    if id < 1 || id > len(records) {
        return fmt.Errorf("task ID %d out of range (1-%d)", id, len(records))
    }

    records = append(records[:id-1], records[id:]...)

    // Write back to file
    file, err = os.OpenFile("tasks.csv", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, record := range records {
        if err := writer.Write(record); err != nil {
            return err
        }
    }

    return nil
}