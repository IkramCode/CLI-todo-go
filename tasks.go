package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const fileName = "task.json"

type task struct {
	Name      string `json:"name"`
	Completed bool   `json:"status"`
}

type list struct {
	Tasks []task `json:"tasks"`
}

func loadTaks() (*list, error) {
	file, err := os.ReadFile(fileName)
	if os.IsNotExist(err) {
		return &list{}, nil
	}
	if err != nil {
		return nil, err
	}
	if len(file) == 0 {
		return &list{}, nil
	}
	var myList list
	if err := json.Unmarshal(file, &myList); err != nil {
		return nil, err
	}
	return &myList, nil
}

func saveTasks(l *list) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}
	os.WriteFile(fileName, data, 0644)
	return nil
}

func (l *list) addTask(name string) error {
	if name == "" {
		return fmt.Errorf("task name cannot be empty")
	}
	newTask := task{Name: name, Completed: false}
	l.Tasks = append(l.Tasks, newTask)
	return saveTasks(l)
}

func (l *list) listTasks() error {
	if len(l.Tasks) == 0 {
		return fmt.Errorf("no tasks found")
	}
	fmt.Println("Your Tasks:")
	for i, t := range l.Tasks {
		status := "not completed"
		if t.Completed {
			status = "completed"
		}
		fmt.Printf("%d. %s [%s]\n", i, t.Name, status)
	}
	return nil
}

func (l *list) deleteTask(index string) error {
	if len(l.Tasks) == 0 {
		return errors.New("len is zero of list")
	}

	ind, err := strconv.Atoi(index)
	if err != nil {
		log.Fatal("cannot convert to int")
	}
	l.Tasks = append(l.Tasks[:ind], l.Tasks[ind+1:]...)
	saveTasks(l)
	return nil
}

func (l *list) completeTask(index string) error {
	if len(l.Tasks) == 0 {
		return errors.New("len is zero of list")
	}
	ind, err := strconv.Atoi(index)
	if err != nil {
		log.Fatal("cannot convert to int")
	}
	l.Tasks[ind].Completed = true
	saveTasks(l)
	return nil
}
