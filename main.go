package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// todo : move all the todo items to todo module
type TodoItem struct {
	Title       string `json:"title"`
	Description string `json:"description`
	Status bool `json:"done"`//expand to include multiple statuses, Not Started, Started, Done, Stalled 
}
type Todo struct {
	Title string `json:"title"`
	Items []TodoItem `json:"items`
}

func (todo Todo) display() {
	// Displays todo list
	fmt.Println(strings.Repeat("=", len(todo.Title)))
	fmt.Println(todo.Title)
	fmt.Println(strings.Repeat("=", len(todo.Title)))
	for i, item := range todo.Items  {
		fmt.Printf("%d : %s\n", i+1, item.toString())
	}
}

func (todoItem TodoItem) toString() string {
	return fmt.Sprintf("%s : %s", todoItem.Title, todoItem.Description)
}


func (todo *Todo) addItem(title, description string){
	todoItem := TodoItem{Title: title, Description: description, Status: false}
	todo.Items = append(todo.Items, todoItem)
}

func (todo *Todo) addItemWithStatus(title, description string, done bool){
	todoItem := TodoItem{Title: title, Description: description, Status: done}
	todo.Items = append(todo.Items, todoItem)
}

func input(prompt string) (string) {
	// Utility function for CLI but can be kept in commons
	fmt.Printf("%s : ", prompt)
	in := bufio.NewReader(os.Stdin)
	result, _, _ := in.ReadLine()
	return string(result)
}

type TodoView interface {
	addItem(todo *Todo) // Should the method signature be different??
	display(todo []Todo)
}

type TodoStorage interface {
	open()
	read()
	write()
	close()
}

func (cli CLI) addItem(todo *Todo){
	itemTitle := input("Enter New Task")
	itemDescription := input("Enter Task Description")
	todo.addItem(string(itemTitle), string(itemDescription))
}

func (cli CLI) display(todoLists []Todo) {
	// Displays todo lists titles
	fmt.Printf("TODO LISTS")
	for i, todo := range todoLists { 
		fmt.Println(i+1, todo.Title)
	}
	options := "Select Option : []View List, [+]New List, [x]Exit"
	fmt.Printf("%s : ", options)
}

type CLI struct {
	// Empty Struct ??
	// Implements todoview interface
}

type App struct {
	view TodoView
	data []Todo
}

func initTodoApp() (App)  {
	log.Printf("Initializing App")
	// Select which View to intiate
	var view CLI

	// Select database to intiate
	// todo : currently harcoding file name
	jsonInput, err := os.ReadFile("todo.json")
	logAndFatal(err)
	// fmt.Printf("input bytes %d \n %s\n", len(jsonInput), jsonInput)
	var todoLists []Todo
	err = json.Unmarshal(jsonInput, &todoLists)
	logAndFatal(err)
	view.display(todoLists)

	return App{view, todoLists}
}


func main() {
	initTodoApp()

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

