package main

import (
	_ "fmt"
	"html/template"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

var data = PageData{
	Title: "Todo List",
	Todos: []Todo{
		{Item: "Hi", Done: false},
	},
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{Item: "Hi", Done: false},
		},
	}
	tmpl.Execute(w, data)
}

func createTodo() http.Handler { // Creates page on sever
	fn := func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Todo List",
			Todos: []Todo{},
		}
		tmpl.Execute(w, data)
	}
	return http.HandlerFunc(fn)
}

func appendTodo(page PageData, adding Todo) http.Handler { //Currently will probably completely wipe out the old page but we can probably fix that
	fn := func(w http.ResponseWriter, r *http.Request) {
		_ = page.Todos
		_ = append(page.Todos, adding)
	}
	return http.HandlerFunc(fn)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
	ts := createTodo()
	tn := appendTodo(data, Todo{Item: "hi", Done: false})
	mux.Handle("/todo", ts)
	mux.Handle("/todo", tn) //append part

	//fs := http.FileServer(http.Dir("./static"))
	//mux.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(":9091", mux))
}
