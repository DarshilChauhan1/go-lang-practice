package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
    ID          int32  `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Done        bool   `json:"done"`
}

var todos []Todo

func main() {
	// Load existing todos from file
	loadTodos()

	// Main menu
	fmt.Printf("Welcome to the Todo App\n")
	fmt.Printf("1. Add Todo\n")
	fmt.Printf("2. List Todos\n")
	fmt.Printf("3. Update Todo\n")
	fmt.Printf("4. Delete Todo\n")
	fmt.Printf("5. Exit\n")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		addTodo()
	case 2:
		listTodos()
	case 3:
		var id int32
		var updateDto Todo
		fmt.Printf("Enter the id of the todo to update: \n")
		fmt.Scanln(&id)
		fmt.Printf("Enter title: \n")
		fmt.Scanln(&updateDto.Title)
		fmt.Printf("Enter description: \n")
		fmt.Scanln(&updateDto.Description)
		updateTodo(id, updateDto)
	case 4:
		var id int32
		fmt.Printf("Enter the id of the todo to delete: \n")
		fmt.Scanln(&id)
		deleteTodo(id)
	case 5:
		fmt.Printf("Exiting the Todo App\n")
		// Save todos to file before exit
		saveTodos()
	default:
		fmt.Printf("Invalid choice\n")
	}
}

func addTodo() {
	bytesTodo, err := os.ReadFile("todos.json")
    if err != nil {
        fmt.Printf("Error reading todos file: %v\n", err)
        return
    }

    err = json.Unmarshal(bytesTodo, &todos)
    if err != nil {
        fmt.Printf("Error unmarshalling todos: %v\n", err)
        return
    }
	fmt.Printf("bytesTodo %v", bytesTodo)
	fmt.Printf("todos  %v", todos)
	

	var todo Todo
	fmt.Printf("Enter title: \n")
	fmt.Scanln(&todo.Title)

	fmt.Printf("Enter description: \n")
	fmt.Scanln(&todo.Description)

	fmt.Printf("Is the todo completed? (yes/no):")
	var done string
	fmt.Scanln(&done)

	if done == "yes" {
		todo.Done = true
	} else {
		todo.Done = false
	}
	if len(todos) > 0 {
		todo.ID = todos[len(todos)-1].ID + 1
	} else {
		todo.ID = 1
	}
	todos = append(todos, todo)
	// Save todos to file after adding a new todo
	saveTodos()
}

func listTodos() {
	for _, todo := range todos {
		fmt.Printf("ID: %d\n", todo.ID)
		fmt.Printf("Title: %s\n", todo.Title)
		fmt.Printf("Description: %s\n", todo.Description)
		fmt.Printf("Done: %t\n", todo.Done)
	}
}

func updateTodo(id int32, updateDto Todo) (Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updateDto
			saveTodos()
			return todos[i], nil
		}
	}
	return Todo{}, errors.New("Todo not found")
}

func deleteTodo(id int32) error {
	// find the todo with the given id
	// find the todo index
	// delete the todo
	// return nil
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos()
			return nil
		}
	}
	return errors.New("Todo not found")
}

func saveTodos() {
	// Save todos slice to a file
	data, err := json.MarshalIndent(todos, "", "")
	fmt.Printf("data: %v\n", data)
	if err != nil {
		fmt.Println("Error saving todos:", err)
		return
	}

	err = os.WriteFile("todos.json", data, 0644)
	if err != nil {
		fmt.Println("Error saving todos to file:", err)
	}
}

func loadTodos() {
	// Load todos from file
	file, err := os.ReadFile("todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, don't worry about it
			return
		}
		fmt.Println("Error reading todos file:", err)
		return
	}

	err = json.Unmarshal(file, &todos)
	if err != nil {
		fmt.Println("Error unmarshalling todos:", err)
	}
}
