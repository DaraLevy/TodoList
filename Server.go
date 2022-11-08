package main

//CITE: https://github.com/NerdCademyDev/golang/tree/main/05_http_server
//DESC: Starter code for the todo list and html, .js and .css files used for todo list server

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

var data = PageData{
	Title: "TODO List",
	Todos: []Todo{
		{Item: "Install GO", Done: true},
		{Item: "Finish Term Project", Done: false},
	},
}

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func appends(input string) http.Handler { // Creates page on sever
	fn := func(w http.ResponseWriter, r *http.Request) {
		data.Todos = append(data.Todos, Todo{input, false})
		tmpl.Execute(w, data)
	}
	return http.HandlerFunc(fn)
}

func delete(input string) http.Handler {
	var new_Todo []Todo
	fn := func(w http.ResponseWriter, r *http.Request) {
		for _, task := range data.Todos {
			if task.Item != input {
				new_Todo = append(new_Todo, task)
			}
		}
		data.Todos = new_Todo
		tmpl.Execute(w, data)
	}
	return http.HandlerFunc(fn)
}

func main() {

	fmt.Println("Append or Delete?:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	if input == "Append" {
		fmt.Println("Enter Task:")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		mux.Handle("/todo", appends(input))
	} else {
		fmt.Println("Enter Task:")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		mux.Handle("/todo", delete(input))
	}

	log.Fatal(http.ListenAndServe(":9091", mux))
}
