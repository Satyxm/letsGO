package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		tasks = append(tasks, args[0])
		saveTasks(tasks)
		fmt.Println("Task added:", args[0])
	},
}
