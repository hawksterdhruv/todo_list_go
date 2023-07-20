# Todo List

## Interfaces
1. CLI (In Progress)
2. Browser (In Progress)

## Learnings
- Implemented a TodoView interface
- Capitalizing allows export
    - Capitalizing `struct` and `function` names allows them to be exported from the module
- Mechanism for using local modules
    ```zsh
    go mod edit -replace github.com/hawksterdhruv/todo_list_golang/commons=./commons
    go mod tidy
    ```

### Questions / Design decision confusions
- Does opening the file fall under the purview of `Initapp()`
- Does the inital display fall under the purview of the `Initapp()`
- Is showing options responsibility of the struct? 
    - Eg. todo list display should display operations on the list??
    - Eg. app options should be displayed by the `app.run` or `app.display` ? 

## Roadmap
1. Create a config file to allow for changing databases
