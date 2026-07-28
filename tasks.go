package main

import (
	// "encoding/csv"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	// "time"
	"github.com/mergestat/timediff"
)

var tasks []string

const filename string = "data.csv"

var noOfTasks int = 0

func getFile() {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// initial header parsing
	headers := []string{"taskid", "task", "time", "status"}
	writer := csv.NewWriter(file)
	writer.Write(headers)
	writer.Flush()
	file.Close()
}

func moppleCreate(task string) error {
	f1, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	row := []string{strconv.Itoa(noOfTasks + 1), task, "9:19", "pending"}
	writer := csv.NewWriter(f1)
	writer.Write(row)
	writer.Flush()
	f1.Close()

	tasks = append(tasks, task) // this creates a new slice with appended elements
	// time := time.Now() // this gives type time.Time
	noOfTasks++
	return errors.New("task sucessfully created")
}

func moppleDelete(taskId int) error {
	if taskId > len(tasks) {
		return errors.New("nothing to delete")
	}
	tasks = slices.Delete(tasks, taskId-1, taskId)
	return errors.New("task sucessfully deleted")
}

func moppleList() error {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil{
		panic(err)
	}
	for _, row := range data{
		for _, col := range row{
			fmt.Printf("%s, ", col)
		}
		fmt.Println()
	}
	
	for index, task := range tasks {
		fmt.Println(index+1, "->", task)
	}
	return errors.New("list of tasks retrieved")
}
