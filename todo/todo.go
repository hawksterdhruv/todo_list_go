package todo

// todo : move all the todo items to todo module
type TodoItem struct {
	Title       	string 	`json:"title"`
	Description 	string 	`json:"description`
	Status 			bool	`json:"done"`//expand to include multiple statuses, Not Started, Started, Done, Stalled 
}
type TodoList struct {
	Title 	string 		`json:"title"`
	Items 	[]TodoItem 	`json:"items`
}

func (todoList *TodoList) AddItem(title, description string){
	todoItem := TodoItem{Title: title, Description: description, Status: false}
	todoList.Items = append(todoList.Items, todoItem)
}

func (todoList *TodoList) addItemWithStatus(title, description string, done bool){
	todoItem := TodoItem{Title: title, Description: description, Status: done}
	todoList.Items = append(todoList.Items, todoItem)
}