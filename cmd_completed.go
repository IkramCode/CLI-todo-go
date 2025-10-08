package main

import (
	"log"

	"github.com/spf13/cobra"
)

var cmdComplete = &cobra.Command{
	Use:   "done",
	Short: "marks the task completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		l, err := loadTaks()
		if err != nil {
			log.Fatal("unable to load tasks")
		}
		if err := l.completeTask(args[0]); err != nil {
			log.Fatal(err)
		}
		log.Printf("marked the task completed")
	},
}
