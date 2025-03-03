package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	},
}
