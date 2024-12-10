package main

import (
	"fmt"
)

// Define Task struct
type Task struct {
	ID          uint
	Title       string
	Description string
	IsCompleted bool
}

// Slice to store tasks
var tasks []Task
var taskID uint

// Function to add a task
func addTask(title, description string) {
	taskID++
	task := Task{
		ID:          taskID,
		Title:       title,
		Description: description,
		IsCompleted: false,
	}
	tasks = append(tasks, task)
	fmt.Println("Task added successfully!")
}

// Function to display all tasks
func displayTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range tasks {
		status := "Incomplete"
		if task.IsCompleted {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, status)
	}
}

// Function to mark a task as completed
func markTaskCompleted(id uint) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].IsCompleted = true
			fmt.Println("Task marked as completed.")
			return
		}
	}
	fmt.Println("Task not found!")
}

// Function to display the menu and handle user input
func main() {
	for {
		// Display menu options
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task Completed")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		// Control flow for user choice and input
		switch choice {
		case 1:
			var title, description string
			fmt.Println("Enter task title:")
			fmt.Scan(&title)
			fmt.Println("Enter task description:")
			fmt.Scan(&description)
			addTask(title, description)
		case 2:
			displayTasks()
		case 3:
			var taskID uint
			fmt.Println("Enter task ID to mark as completed:")
			fmt.Scan(&taskID)
			markTaskCompleted(taskID)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
