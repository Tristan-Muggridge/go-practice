package cli

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"task-manager/internal/repo"
)

func formatEditInputString(input string) string {
	input = strings.TrimSpace(input)              // --title "Update Title" --description "Updated Description" --completed: true
	input = strings.ReplaceAll(input, "--", "")   // title "Update Title" description "Updated Description" completed: true
	input = strings.ReplaceAll(input, " \"", ":") // title:Update Title" description:Updated Description" completed: true
	input = strings.ReplaceAll(input, "\" ", ",") // title:Update Title,description:Updated Description,completed: true
	input = strings.ReplaceAll(input, ": ", ":")  // title:Update Title,description:Updated Description,completed:true

	return input
}

func getTaskIDInput() int {
	fmt.Println("Please enter the ID of the task you want to edit:")
	id := captureInput()
	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID")
		return -1
	}
	return id_int
}

func onEditTask(repo repo.TaskRepo) {

	id := getTaskIDInput()
	task, err := repo.GetTaskById(id)
	if err != nil {
		log.Fatalln(err)
	}

	if task == nil {
		fmt.Println("Task not found")
		return
	}

	fmt.Println("Please enter your edits using the following structure:")
	fmt.Println("--title \"Update Title\" --description \"Updated Description\" --completed: true")

	edits := formatEditInputString(captureInput()) // title:Update Title,description:Updated Description,completed:true

	for _, edit := range strings.Split(edits, ",") {
		edit = strings.TrimSpace(edit)
		editParts := strings.Split(edit, ":")
		if len(editParts) != 2 {
			fmt.Println("Invalid edit structure")
			return
		}

		field, value := editParts[0], editParts[1]
		switch field {
		case "title":
			task.Title = value
		case "description":
			task.Description = value
		case "completed":
			completed, err := strconv.ParseBool(value)
			if err != nil {
				fmt.Println("Invalid completed value")
				return
			}
			task.Completed = completed
		default:
			fmt.Println("Invalid field")
			return
		}
	}

	fmt.Println("Updating task...")

	err = repo.UpdateTask(task)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(task)
}
