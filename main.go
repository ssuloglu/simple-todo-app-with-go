package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpl *template.Template

type Todo struct {
	Item    string
	DueDate string
	Done    bool
}

type PageData struct {
	Title        string
	Date         string
	DoneItems    int
	WaitingItems int
	Todos        []Todo
}

func getDoneItems(todos []Todo) int {
	var count int = 0
	for _, t := range todos {
		if t.Done {
			count = count + 1
		}
	}
	return count
}

func index(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{Item: "Initialize go", DueDate: "19-7-2023", Done: true},
		{Item: "Create mux to handle requests", DueDate: "19-7-2023", Done: true},
		{Item: "Create UI templates", DueDate: "19-7-2023", Done: true},
		{Item: "Demo your code", DueDate: "20-7-2023", Done: false},
	}
	var nDoneItems int = getDoneItems(todos)

	data := PageData{
		Title:        "Todo List",
		Date:         time.Now().Format("01-02-2006"),
		DoneItems:    nDoneItems,
		WaitingItems: len(todos) - nDoneItems,
		Todos:        todos,
	}
	tmpl.Execute(w, data)
}

func main() {

	// server request handler
	mux := http.NewServeMux()

	//index page template
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	// static file storage for css and js files
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// handle index page
	mux.HandleFunc("/", index)

	// start web server listening port 8083
	log.Fatal(http.ListenAndServe(":8083", mux))
}
