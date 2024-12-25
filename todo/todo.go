package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"createdAT"`
	CompletedAt *time.Time `json:"completedAt"`
}

type Todos []Todo

func (todos *Todos) Add(title string) (t Todo) {
	todo := Todo{
		Title:       title,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
		Completed:   false,
	}

	*todos = append(*todos, todo)

	return todo
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	todo := &t[index]

	if !todo.Completed {
		todo.Completed = true
		completionTime := time.Now()
		todo.CompletedAt = &completionTime
	}

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	todo := &t[index]

	todo.Title = title

	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CompletedAt", "CreatedAt")

	for index, t := range *todos {
		completed := "❎"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, completedAt, t.CreatedAt.Format(time.RFC1123))
	}

	table.Render()
}
