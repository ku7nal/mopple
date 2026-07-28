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

type task struct{ // this made life so much easier
	taskId int
	desc string
	time time.Time
	status bool
}

var tasks []task

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
	for i, row := range data {
		if i != 0 {
			for j, col := range row{
				switch j{
				case 0:
					tasks[i].taskId, _ = strconv.Atoi(col)
				case 1:
					tasks[i].desc = col
				case 2:
					layout := "2006-01-02 15:04:05"
					tasks[i].time, _ = time.Parse(layout, col)
				case 3:
					tasks[i].status, _ = strconv.ParseBool(col)

				}
			}
		}
	}
	file.Close()
}

func moppleCreate(d string) error {

	noOfTasks++
	newTask := task{taskId: noOfTasks, desc: d,time: time.Now(), status: false}

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

func moppleList() error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintln(w, "TaskId\tDescription\tTime\tStatus")
	for _, task := range tasks{
		time := timediff.TimeDiff(task.time)
		fmt.Fprintf(w,"%s\t%s\t%s\t%s\n", strconv.Itoa(task.taskId), task.desc, time, strconv.FormatBool(task.status))
	}
	w.Flush() // this is critical
	return errors.New("list of tasks retrieved")
}
