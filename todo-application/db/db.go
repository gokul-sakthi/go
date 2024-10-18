package db

import (
	"fmt"
	"sync"
)

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type TodoApp struct {
	NextId int
	Mutex  sync.Mutex
	Todos  []Todo
}

func NewTodoApp() *TodoApp {
	return &TodoApp{
		NextId: 1,
		Mutex:  sync.Mutex{},
		Todos:  []Todo{},
	}
}

func (app *TodoApp) Add(t Todo) bool {
	app.Mutex.Lock()
	defer app.Mutex.Unlock()

	app.Todos = append(app.Todos, t)
	app.NextId++
	return true
}

func (app *TodoApp) Remove(id int) bool {
	app.Mutex.Lock()
	defer app.Mutex.Unlock()

	for i, item := range app.Todos {
		if item.Id == id {
			app.Todos = append(app.Todos[:i], app.Todos[i+1:]...)
			fmt.Printf("Deleted Todo with ID (%v)", id)
			return true
		}

	}
	return false
}

func (app *TodoApp) List() {
	app.Mutex.Lock()
	defer app.Mutex.Unlock()

	if len(app.Todos) == 0 {
		fmt.Println("The List is empty")
		return
	}

	for i, item := range app.Todos {
		fmt.Printf("S.No: %v | ID: %v | Name: %v | Description: %v | Status: %v\n", i, item.Id, item.Name, item.Description, item.Status)
	}
}

func (app *TodoApp) Update(id int, completed bool) bool {
	app.Mutex.Lock()
	defer app.Mutex.Unlock()

	for i, item := range app.Todos {
		if item.Id == id {
			if completed {
				app.Todos[i].Status = "Completed"
			} else {
				app.Todos[i].Status = "Pending"
			}
			return true
		}
	}
	return false
}

func (app *TodoApp) Clear() {
	app.Todos = []Todo{}
}
