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

func onCreateTask() {
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
	fmt.Println(task)
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
			onCreateTask()
		case controlOptions[1]:
			fmt.Println("Listing tasks...") // TODO: Implement
		case controlOptions[2]:
			fmt.Println("Editing task...") // TODO: Implement
		case controlOptions[3]:
			fmt.Println("Deleting task...") // TODO: Implement
		default:
			fmt.Println("Invalid option") // TODO: Implement
		}

		printNumberedList(controlOptions)
	}
}
