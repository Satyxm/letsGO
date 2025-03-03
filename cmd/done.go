package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 || index > len(tasks) {
			fmt.Println("Invalid task number.")
			return
		}
		tasks = append(tasks[:index-1], tasks[index:]...)
		saveTasks(tasks)
		fmt.Println("Task marked as done.")
	},
}
