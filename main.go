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

// TODO: CHANGE SIGNATURE !!
func display(todo todo.TodoList) {
	// Displays todo list
	fmt.Println(strings.Repeat("=", len(todo.Title)))
	fmt.Println(todo.Title)
	fmt.Println(strings.Repeat("=", len(todo.Title)))
	for i, item := range todo.Items  {
		fmt.Printf("%d : %s\n", i+1, toString(item))
	}
}

// TODO: CHANGE SIGNATURE !!
func  toString(todoItem todo.TodoItem) string {
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
func addItem(todo *todo.TodoList){
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
	view TodoView
	data []todo.TodoList // live session data
	filename string
}

func (app App) run(){
	for {
		app.view.display(app.data)
		
		option := commons.Input("Select Option: ")
		listIndex, err := strconv.Atoi(option)
		if err == nil  {
			if listIndex>0 && listIndex <= len(app.data) {
				currentTodoList := &app.data[listIndex - 1]

				display(*currentTodoList)
				options := "[+]AddItem, [x]Back, [d]Delete list"
				fmt.Printf("%s : ", options)
				todoOptions := commons.Input("Select Option: ")
				switch todoOptions{
				case "+":
					addItem(currentTodoList)
					// WRONG !! This will take us back to the main menu.
				case "x","X":
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
			case "+" :
				// Call method to add todo list
			case "x", "X":
				// Call method for exiting program
				// Alternative runtime.Goexit()
				// defer os.Exit(0)
				fp := openCreateFile(app.filename)
				defer closeFile(fp)

				todoJson, err := json.Marshal(app.data)
				logAndFatal(err)
				_, err = fp.WriteString(string(todoJson))
				logAndFatal(err)
				
				return 
			default:
				commons.Input("You have selected an unknown option, Please choose again.")
				continue
			}
		}
	}
}

func initTodoApp(filename string) (App)  {
	log.Printf("Initializing App")
	// Select which View to intiate
	var view CLI
	var todoLists []todo.TodoList
	// Select database to intiate
	// todo : currently harcoding file name
	jsonInput, err := os.ReadFile(filename)
	logAndFatal(err)
	err = json.Unmarshal(jsonInput, &todoLists)
	logAndFatal(err)

	return App{view, todoLists, filename}
}


func main() {
	app := initTodoApp("todo.json")
	app.run()

	// var todo Todo
	// See all todo lists
	// todo.Title = "Current"
	
	// cli.addItem(&todo)
	
	// cli.display(todo)
	// todoJson, err := json.Marshal(todo)
	// if err != nil{
		// log.Println("Failed to Convert to Json : ", err)
	// }
	// fmt.Println(string(todoJson))
	// var jsonInput = `{"Title":"Current","Items":[{"title":"refactoring","Description":"study","done":false}]}`
	// var todo Todo
	
	// json.Unmarshal([]byte(jsonInput), &todo)
	// fmt.Printf("%T : %v\n", todo, todo)
	// fmt.Println(todo)

	// fp := openCreateFile("todo.json")
	// defer closeFile(fp)
	// // fp.WriteString(jsonInput)
	// var jsonInput []byte
	// n,err := fp.Read(jsonInput)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Printf("Read %d characters", n)
	// }
	// jsonInput, err := os.ReadFile("todo.json")
	// logAndFatal(err)
	// fmt.Printf("input bytes %d \n %s\n", len(jsonInput), jsonInput)
	// err = json.Unmarshal(jsonInput, &todo)
	// logAndFatal(err)
	// fmt.Printf("Unmarshalled output : %T : %v\n", todo, todo)

}
func logAndFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
} 
func openCreateFile(filename string) (*os.File){
	
	fp, err := os.Open(filename)
	
	if err != nil {
		log.Println(err)
		log.Printf("Will try to create %s file", filename)
	} else {
		// Super suspect implementation
		// Need to find better coding standard
		return fp
	}

	fp, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fp
}
func closeFile(fp *os.File) {
	err := fp.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Closed file Successfully")
	}

}

