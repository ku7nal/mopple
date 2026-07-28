package main

import (
	"errors"
	"fmt"
	"slices"
)

var tasks []string

func moppleCreate(task string) error {
	tasks = append(tasks, task) // this creates a new slice with appended elements
	return errors.New("task sucessfully created")
}

func moppleDelete(taskId int) error {
	tasks = slices.Delete(tasks, taskId-1, taskId)
	return errors.New("task sucessfully deleted")
}

func moppleList() error {
	fmt.Println(tasks)
	return errors.New("list of tasks retrieved")
}
