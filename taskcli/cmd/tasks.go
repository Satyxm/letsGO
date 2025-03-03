package cmd

import (
	"encoding/json"
	"os"
)

var taskFile = "tasks.json"

func loadTasks() []string {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		return []string{}
	}
	var tasks []string
	json.Unmarshal(file, &tasks)
	return tasks
}

func saveTasks(tasks []string) {
	file, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(taskFile, file, 0644)
}
