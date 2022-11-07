package main

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

func main() {
	fmt.Println("Enter Task:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/todo", appends(input))

	log.Fatal(http.ListenAndServe(":9091", mux))
}

//The code below is just the original test code from the last commit, Doesn't work

//package main
//
//import (
//	_ "fmt"
//	"github.com/gorilla/mux"
//	"html/template"
//	"net/http"
//)
//
//var tmpl *template.Template
//
//type Todo struct {
//	Item string
//	Done bool
//}
//
//type PageData struct {
//	Title string
//	Todos []Todo
//}
//
//var data = PageData{
//	Title: "Todo List",
//	Todos: []Todo{},
//}
//
////func todo(w http.ResponseWriter, r *http.Request) {
////	data = PageData{
////		Title: "Todo List",
////		Todos: []Todo{
////			{Item: "Hi", Done: false},
////		},
////	}
////	tmpl.Execute(w, data)
////}
//
////func createTodo() http.Handler { // Creates page on sever
////	fn := func(w http.ResponseWriter, r *http.Request) {
////		data = PageData{
////			Title: "Todo List",
////			Todos: []Todo{
////				{Item: "Go Home", Done: false},
////			},
////		}
////	}
////	return http.HandlerFunc(fn)
////}
//
//func appendTodo(page PageData, adding Todo) http.Handler { //Currently will probably completely wipe out the old page but we can probably fix that
//
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		page.Todos = append(page.Todos, adding)
//	}
//	return http.HandlerFunc(fn)
//}
//
//func CreateTodo(w http.ResponseWriter, r *http.Request) {
//	data.Todos = append(data.Todos, Todo{"Added", false})
//	tmpl.Execute(w, data)
//}
//
//func main() {
//	Router := mux.NewRouter()
//	tmpl = template.Must(template.ParseFiles("index.gohtml"))
//
//	//Router.HandleFunc("/todo-completed", GetCompletedItems).Methods("GET")
//	//Router.HandleFunc("/todo-incomplete", GetIncompleteItems).Methods("GET")
//	Router.HandleFunc("/todo", CreateTodo).Methods("POST")
//
//	fs := http.FileServer(http.Dir("./static"))
//	Router.Handle("/static/", http.StripPrefix("/static/", fs))
//	http.ListenAndServe(":9091", Router)
//
//}
