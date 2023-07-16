package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// todo : move all the todo items to todo module
type TodoItem struct {
	title       string
	description string
	status bool //expand to include multiple statuses, Not Started, Started, Done, Stalled 
}
type Todo struct {
	title string
	items []TodoItem
}
func (todoItem TodoItem) toString() string {
	return fmt.Sprintf("%s : %s", todoItem.title, todoItem.description)
}
func (todo Todo) cliDisplay() {
	// Displays todo list
	// Part of CLI interface
	fmt.Println(strings.Repeat("=", len(todo.title)))
	fmt.Println(todo.title)
	fmt.Println(strings.Repeat("=", len(todo.title)))
	for i := 0; i <len(todo.items); i++ {
		fmt.Printf("%d : %s\n", i+1, todo.items[i].toString())
	}
}

func (todo *Todo) addItem(title, description string){
	todoItem := TodoItem{title: title, description: description, status: false}
	todo.items = append(todo.items, todoItem)
}

func (todo *Todo) addItemWithStatus(title, description string, done bool){
	todoItem := TodoItem{title: title, description: description, status: false}
	todo.items = append(todo.items, todoItem)
}

func main() {
	var todo Todo
	// See all todo lists
	todo.title = "Current"
	// var itemTitle string
	// fmt.Scanf("%s\n",&itemTitle)
	itemTitle := input("Enter New Task")
	itemDescription := input("Enter Task Description")
	// fmt.Scanf("%s\n",&itemDescription)
	todo.addItem(string(itemTitle), string(itemDescription))
	todo.cliDisplay()


}

func input(prompt string) (string) {
	fmt.Printf("%s : ", prompt)
	in := bufio.NewReader(os.Stdin)
	result, _, _ := in.ReadLine()
	return string(result)
}