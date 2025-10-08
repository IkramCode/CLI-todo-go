package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var cmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "Use to delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		l, err := loadTaks()
		if err != nil {
			log.Fatal("Cannot load tasks")
		}
		if err := l.deleteTask(args[0]); err != nil {
			log.Fatal("could not delete")
		}
		fmt.Printf("Deleted the task")
	},
}
