package cli

/*
	App Start:
		1. User runs app
		2. App connects to repo (database / file / etc)
		3. App displays menu options
			3.1. 'Create Task'
			3.2. 'List Tasks'
			3.3. 'Edit Task'
			3.4. 'Delete Task'
			3.5. 'Exit'

	User Journey: Create Task
	1-3. Same as App Start
	4. User selects 'Create Task'
	5. App prompts user for task title
	6. User enters task title
	7. App prompts user for task description
	8. User enters task description
	9. App creates task
	10. App displays menu options

	User Journey: List Tasks

	1-3. Same as App Start
	4. User selects 'List Tasks'
	5. App retrieves tasks from repo
	6. App displays tasks
	7. App displays menu options

	User Journey: Edit Task

	1-3. Same as App Start
	4. User selects 'Edit Task'
	5. App lists tasks with Title and ID
	6. App prompts user for task ID
	7. User enters task ID
	8. App prompts user for task details with relevant flags
		ie: --title "Update Title" --description "Updated Description" --completed: true
	9. App updates task
	10. App displays menu options

	User Journey: Delete Task

	1-3. Same as App Start
	4. User selects 'Delete Task'
	5. App lists tasks with Title and ID
	6. App prompts user for task ID
	7. User enters task ID
	8. App deletes task
	9. App displays menu options
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"task-manager/internal/models/task"
	"task-manager/internal/repo"
)

var controlOptions = []string{
	"Create Task",
	"List Tasks",
	"Edit Task",
	"Delete Task",
	"Exit",
}

func printNumberedList(items []string) {
	for i, items := range items {
		fmt.Printf("%d. %s\n", i+1, items)
	}
}

func captureInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) // remove the newline character at the end
}

func onCreateTask(repo repo.TaskRepo) {
	fmt.Println("Please enter a title:")
	title := captureInput()
	fmt.Println("Please enter a description:")
	description := captureInput()

	task := &task.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	fmt.Println("Creating task...")

	err := repo.CreateTask(task)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(task)
}

func onListTasks(repo repo.TaskRepo) {
	fmt.Println("Listing tasks...")
	tasks := repo.GetTasks()
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func onEditTask(repo repo.TaskRepo) {
	fmt.Println("Please enter the ID of the task you want to edit:")
	id := captureInput()
	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task, err := repo.GetTaskById(id_int)
	if err != nil {
		log.Fatalln(err)
	}
	if task == nil {
		fmt.Println("Task not found")
		return
	}

	fmt.Println("Please enter your edits using the following structure:")
	fmt.Println("--title \"Update Title\" --description \"Updated Description\" --completed: true")

	edits := captureInput()                       // --title "Update Title" --description "Updated Description" --completed: true
	edits = strings.TrimSpace(edits)              // --title "Update Title" --description "Updated Description" --completed: true
	edits = strings.ReplaceAll(edits, "--", "")   // title "Update Title" description "Updated Description" completed: true
	edits = strings.ReplaceAll(edits, " \"", ":") // title:Update Title" description:Updated Description" completed: true
	edits = strings.ReplaceAll(edits, "\" ", ",") // title:Update Title,description:Updated Description,completed: true
	edits = strings.ReplaceAll(edits, ": ", ":")  // title:Update Title,description:Updated Description,completed:true

	editsMap := make(map[string]string)

	for _, edit := range strings.Split(edits, ",") {
		edit = strings.TrimSpace(edit)
		editParts := strings.Split(edit, ":")
		if len(editParts) != 2 {
			fmt.Println("Invalid edit")
			return
		}
		editsMap[editParts[0]] = editParts[1]
	}

	if title, ok := editsMap["title"]; ok {
		task.Title = title
	}

	if description, ok := editsMap["description"]; ok {
		task.Description = description
	}

	if completed, ok := editsMap["completed"]; ok {
		completed_bool, err := strconv.ParseBool(completed)
		if err != nil {
			fmt.Println("Invalid completed value")
			return
		}
		task.Completed = completed_bool
	}

	fmt.Println("Updating task...")

	err = repo.UpdateTask(task)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(task)
}

func onDeleteTask(repo repo.TaskRepo) {
	fmt.Println("Please enter the ID of the task you want to delete:")
	id := captureInput()
	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task, err := repo.GetTaskById(id_int)
	if err != nil {
		log.Fatalln(err)
	}

	if task == nil {
		fmt.Println("Task not found")
		return
	}

	fmt.Println("Deleting task...")
	err = repo.DeleteTask(task)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Task deleted")
}

func NewCli(repo repo.TaskRepo) {
	fmt.Println("Welcome to the Task Manager")
	fmt.Println("Please select an option:")

	printNumberedList(controlOptions)

	for {
		option_input := captureInput()
		option_input_int, err := strconv.Atoi(option_input)
		if err != nil {
			fmt.Println("Invalid option")
			continue
		}

		selectedOption := controlOptions[option_input_int-1]

		if selectedOption == controlOptions[len(controlOptions)-1] {
			fmt.Println("Exiting...")
			break
		}

		switch selectedOption {
		case controlOptions[0]:
			onCreateTask(repo)
		case controlOptions[1]:
			onListTasks(repo)
		case controlOptions[2]:
			onEditTask(repo)
		case controlOptions[3]:
			onDeleteTask(repo)
		default:
			fmt.Println("Invalid option") // TODO: Implement
		}

		printNumberedList(controlOptions)
	}
}
