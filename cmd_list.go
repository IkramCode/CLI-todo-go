package main

import (
	"log"

	"github.com/spf13/cobra"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List all task",
	Run: func(cmd *cobra.Command, args []string) {
		l, err := loadTaks()
		if err != nil {
			log.Fatal(err)
		}
		if err := l.listTasks(); err != nil {
			log.Fatal(err)
		}
	},
}
