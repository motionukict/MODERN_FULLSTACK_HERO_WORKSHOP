package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	_ "modernc.org/sqlite"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.Write([]byte("<marquee>hello world</marquee>"))
}
func getFormHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	var html = `
		<form action="/form" method="POST">
			<label for="name" >First Name</label>
			<input type="text" name="name" id="name" placeholder="Enter your name" />
			<input type="submit" value="Submit" />
		</form>`
	writer.Write([]byte(html))
}

func postFormHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var db = Connect()
	InsertName(db, request.Form.Get("name"))
	db.Close()
	var html = `<p> Hello ` + request.Form.Get("name") + `</p>`
	writer.Write([]byte(html))
}
func getNamesHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	var strBuilder strings.Builder
	strBuilder.WriteString("<ul>")
	var db = Connect()
	var names = GetNames(db)
	for _, name := range names {
		strBuilder.WriteString("<li>")
		strBuilder.WriteString(name)
		strBuilder.WriteString("</li>")
	}
	strBuilder.WriteString("</ul>")

	writer.Write([]byte(strBuilder.String()))
}
func main() {
	var router = chi.NewRouter()
	router.Get("/hello", helloHandler)
	router.Get("/form", getFormHandler)
	router.Post("/form", postFormHandler)
	router.Get("/names", getNamesHandler)
	router.Get("/api/names", getApiNamesHandler)
	http.ListenAndServe(":7070", router)

}

func getApiNamesHandler(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var db = Connect()
	var names = GetNames(db)
	var bytes, _ = json.Marshal(names)
	writer.Write(bytes)
}
