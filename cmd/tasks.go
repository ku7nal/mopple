package cmd

// Depenendencies
import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/mergestat/timediff"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

// Data Management
type task struct {
	taskId int
	desc   string
	time   string
	status bool
}

var tasks []task

const filename string = "data.csv"

var noOfTasks = 0

// Creating the Task
func moppleCreate(d string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	errorCheck(err, nil)

	// tracking the no of tasks
	noOfTasks++

	// creating a new slice with appended elements
	newTask := task{taskId: noOfTasks, desc: d, time: time.Now().Format("2006-01-02 15:04:05"), status: false}
	tasks = append(tasks, newTask)

	// writing the task to the csv file
	writer := csv.NewWriter(file)
	var inputline []string
	inputline = append(inputline, strconv.Itoa(newTask.taskId))
	inputline = append(inputline, newTask.desc)
	inputline = append(inputline, newTask.time)
	inputline = append(inputline, strconv.FormatBool(newTask.status))
	writer.Write(inputline)
	writer.Flush()
	file.Close()

	return errors.New("task sucessfully created")
}

// Deleting the Task
func moppleDelete(taskId int) error {

	if taskId > len(tasks) {
		return errors.New("nothing to delete")
	}

	tasks = removeElement(taskId, tasks)
	writeToFile()
	return errors.New("task sucessfully deleted")
}

// Listing the Tasks
func moppleList() error {

	// specified format
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "TaskId\tDescription\tTime\tStatus")

	// parsing data to render with formatting
	for _, task := range tasks {
		loc := time.Now().Location()
		gettimetype, err := time.ParseInLocation("2006-01-02 15:04:05", task.time, loc)
		errorCheck(err, nil)

		diff := timediff.TimeDiff(gettimetype)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", strconv.Itoa(task.taskId), task.desc, diff, strconv.FormatBool(task.status))
	}
	w.Flush()

	return errors.New("list of tasks retrieved")
}

// Checking for the csv file | Creating the csv file
func getFile() {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		errorCheck(err, nil)

		file.Close()
	} else {
		file, err := os.Open(filename)
		errorCheck(err, nil)

		file.Close()
	}
}

// Loading the Data into the Slice [csv => tasks]
func loadTasks() {
	file, err := os.Open(filename)
	errorCheck(err, nil)

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	errorCheck(err, nil)

	for _, row := range data {
		idfrom, _ := strconv.Atoi(row[0])
		statfrom, _ := strconv.ParseBool(row[3])
		newData := task{
			taskId: idfrom,
			desc:   row[1],
			time:   row[2],
			status: statfrom,
		}
		tasks = append(tasks, newData)
	}

	// if there is no data
	if len(data) > 0{
		noOfTasks, _ = strconv.Atoi(data[len(data)-1][0]) // updating the task count
	}else {
		noOfTasks = 0
	}

	file.Close()
}

// Writing the data from Slice to csv file [tasks => csv]
func writeToFile() {
	file, err := os.Create(filename)

	errorCheck(err, nil)

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

// Marking task as Completed
func doneTask(id int) {
	for i := range tasks {
		if tasks[i].taskId == id {
			tasks[i].status = true
		}
	}
	writeToFile()
}

// Helper Functions =================================================
func removeElement(t int, tslice []task) []task {
	for idx, v := range tslice {
		if v.taskId == t {
			return append(tslice[:idx], tslice[idx+1:]...)
		}
	}
	return tslice
}

func errorCheck(err, want error) {
	if err != want {
		panic(err)
	}
}
