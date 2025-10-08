package main

import (
	"log"

	"github.com/spf13/cobra"
)

var cmdAdd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task",
	Long:  "Add a task to your task list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		l, err := loadTaks()
		if err != nil {
			log.Fatal(err)
		}
		if err := l.addTask(taskName); err != nil {
			log.Fatal(err)
		}
		log.Printf("Task added %v ", taskName)
	},
}
