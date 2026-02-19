package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	Title     string
	Completed string
	CreatedAt time.Time
}

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple todo CLI application",
	Long:  `A command line todo application built with Go and Cobra.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
