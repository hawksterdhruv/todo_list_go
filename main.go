package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hawksterdhruv/todo_list_go/commons"
	"github.com/hawksterdhruv/todo_list_go/todo"
)

// Displays todo list
func display(todo todo.TodoList) {
	// TODO: CHANGE SIGNATURE !!
	fmt.Println(strings.Repeat("=", len(todo.Title)))
	fmt.Println(todo.Title)
	fmt.Println(strings.Repeat("=", len(todo.Title)))
	for i, item := range todo.Items {
		fmt.Printf("%d : %s\n", i+1, toString(item))
	}
}

// TODO: CHANGE SIGNATURE !!
func toString(todoItem todo.TodoItem) string {
	// TODO: CHANGE SIGNATURE !!
	return fmt.Sprintf("%s : %s", todoItem.Title, todoItem.Description)
}

type TodoView interface {
	// addItem(todo *todo.TodoList) // Should the method signature be different??
	display(todo []todo.TodoList)
}

type TodoStorage interface {
	open()
	read()
	write()
	close()
}

func addTodoList() todo.TodoList {
	// Pass app vs app interface method??
	// Modifies app/app.storage object ??
	listTitle := commons.Input("Enter name for new Todo List: ")
	return todo.TodoList{Title: listTitle}

}

// What level of abstaction should this be?
func addItem(todo *todo.TodoList) {
	itemTitle := commons.Input("Enter New Task: ")
	itemDescription := commons.Input("Enter Task Description: ")
	todo.AddItem(string(itemTitle), string(itemDescription))
}

func (cli CLI) display(todoLists []todo.TodoList) {
	// Displays todo lists titles
	commons.ClearScreen()
	fmt.Printf("TODO LISTS\n==========\n")
	for i, todo := range todoLists {
		fmt.Println(i+1, todo.Title)
	}
	options := "[]View List, [+]New List, [x]Exit"
	fmt.Printf("%s: ", options)
}

type CLI struct {
	// Empty Struct ??
	// Implements todoview interface
}

type App struct {
	view     TodoView
	data     []todo.TodoList // live session data
	filename string
}

// Runs the application in a loop
func (app App) run() {
	for {
		app.view.display(app.data)

		option := commons.Input("Select Option: ")
		listIndex, err := strconv.Atoi(option)
		if err == nil {
			if listIndex > 0 && listIndex <= len(app.data) {
				currentTodoList := &app.data[listIndex-1]

				display(*currentTodoList)
				options := "[+]AddItem, [x]Back, [d]Delete list"
				fmt.Printf("%s : ", options)
				todoOptions := commons.Input("Select Option: ")
				switch todoOptions {
				case "+":
					addItem(currentTodoList)
					// WRONG !! This will take us back to the main menu.
				case "x", "X":
					continue
				default:
					commons.Input("You have selected an unknown option, Please choose again.")
					continue
				}

			} else {
				commons.Input("You have selected an unknown option, Please choose again.")
				continue
			}

		} else {
			switch option {
			case "+":
				tempTodoList := addTodoList()
				app.data = append(app.data, tempTodoList)

			case "x", "X":
				// Call method for exiting program
				// Alternative runtime.Goexit()
				// defer os.Exit(0)
				fp, err := os.OpenFile(app.filename, os.O_CREATE|os.O_RDWR, 0666)
				commons.LogAndFatal(err)
				defer fp.Close()

				todoJson, err := json.Marshal(app.data)
				commons.LogAndFatal(err)
				_, err = fp.WriteString(string(todoJson))
				commons.LogAndFatal(err)

				return
			default:
				commons.Input("You have selected an unknown option, Please choose again.")
				continue
			}
		}
	}
}

// Initialises the applicationn
func initTodoApp(filename string) App {
	log.Printf("Initializing App")
	// Select which View to intiate
	var view CLI
	var todoLists []todo.TodoList

	// Select database to intiate
	jsonInput, err := os.ReadFile(filename)
	if _, ok := err.(*os.PathError); ok {

		log.Println(err)
		log.Println("Trying to create new file.")
		_, err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(jsonInput) > 0 {
		err = json.Unmarshal(jsonInput, &todoLists)
	}
	commons.LogAndFatal(err)

	return App{view, todoLists, filename}
}

func main() {
	app := initTodoApp("todo1.json")
	app.run()
}
