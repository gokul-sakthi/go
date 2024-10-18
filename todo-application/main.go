package main

import "todo-app/db"

func main() {
	run()
}

func run() {
	app := db.NewTodoApp()

	app.Add(db.Todo{
		Id:          app.NextId + 1,
		Name:        "Gello",
		Description: "Something",
		Status:      "Pending",
	})

	app.List()

	app.Add(db.Todo{
		Id:     app.NextId + 1,
		Name:   "Sfer",
		Status: "Pending",
	})

	app.List()

	app.Remove(2)

	app.List()
	app.List()

	app.Update(3, true)

	app.List()
	app.List()

	app.Clear()

	app.List()

}
