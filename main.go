package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the to-do list application")
	scanner := bufio.NewScanner(os.Stdin)

	// Define the struct for a to-do task
	type todoTask struct {
		name string
		time string
	}
	// Create a slice to store the to-do tasks
	var todoList []todoTask

	for {
		fmt.Println("Enter `ESC` to Exit the program or click any keyword to Continue ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		// Check if the user wants to exit or continue
		if choice == "ESC" {
			break
		}

		fmt.Println("Enter the name of the task:")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())

		fmt.Println("Enter the estimated time to complete the task:")
		scanner.Scan()
		time := strings.TrimSpace(scanner.Text())

		// Create a new to-do task
		todo := todoTask{name, time}
		// Append the new task to the to-do list
		todoList = append(todoList, todo)
		fmt.Println("Task added:", name, time)
	}

	fmt.Println("\nYour To-Do List:")
	fmt.Println("----------------------")
	for i, todo := range todoList {
		fmt.Println(
			"The",
			i+1,
			"task is",
			todo.name,
			"& it takes this much time:",
			todo.time,
		)
	}
}
