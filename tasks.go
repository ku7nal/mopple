package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

type task struct { // this made life so much easier
	taskId int
	desc   string
	time   string
	status bool
}

var tasks []task

const filename string = "data.csv"

var noOfTasks int = 0

func getFile() {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		file.Close()
	} else {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
}

func loadTasks() {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range data {
		idfrom, _:= strconv.Atoi(row[0])
		statfrom, _ := strconv.ParseBool(row[3])
		newData := task{
			taskId: idfrom,
			desc:   row[1],
			time:   row[2],
			status: statfrom,
		}
		tasks = append(tasks, newData)
	}
	file.Close()
}

func writeToFile() {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	for _, task := range tasks {
		var inputline []string
		inputline = append(inputline, strconv.Itoa(task.taskId))
		inputline = append(inputline, task.desc)
		inputline = append(inputline, task.time)
		inputline = append(inputline, strconv.FormatBool(task.status))
		writer.Write(inputline)
	}
	writer.Flush()
	file.Close()
}

func moppleCreate(d string) error { // this will recreate the file with modified (appended) content

	noOfTasks++
	newTask := task{taskId: noOfTasks, desc: d, time: time.Now().Format("2006-01-02 15:04:05"), status: false}

	tasks = append(tasks, newTask) // this creates a new slice with appended elements
	return errors.New("task sucessfully created")
}

// func moppleDelete(taskId int) error { -> most probably to delete the struct in the array which is easy
// 	if taskId > len(tasks) {
// 		return errors.New("nothing to delete")
// 	}
// 	tasks = slices.Delete(tasks, taskId-1, taskId)
// 	return errors.New("task sucessfully deleted")
// }

func letmesee() {
	fmt.Println(tasks)
}
func moppleList() error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintln(w, "TaskId\tDescription\tTime\tStatus")
	for _, task := range tasks {
		loc := time.Now().Location()
		gettimetype, err := time.ParseInLocation("2006-01-02 15:04:05", task.time, loc)
		if err != nil {
			fmt.Println("Error parsing date:", err)
		}
		diff := timediff.TimeDiff(gettimetype)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", strconv.Itoa(task.taskId), task.desc, diff, strconv.FormatBool(task.status))
	}
	w.Flush() // this is critical
	return errors.New("list of tasks retrieved")
}
