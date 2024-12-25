package main

import (
	"github.com/dtg-lucifer/go-todo-cli/cmd"
	"github.com/dtg-lucifer/go-todo-cli/storage"
	"github.com/dtg-lucifer/go-todo-cli/todo"
)

func main() {
todos := todo.Todos{}

storage := storage.NewStorage[todo.Todos]("todos.json")
storage.Load(&todos)

cmdFlags := cmd.NewCmdFlags()
cmdFlags.Execute(&todos)

storage.Save(todos)
}
