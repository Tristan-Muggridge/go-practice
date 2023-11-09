package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
