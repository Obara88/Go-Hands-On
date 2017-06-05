package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("http://localhost:8080")

	http.HandleFunc("/", list)
	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/save", save)
	http.HandleFunc("/delete", delete)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list")
}

func create(w http.ResponseWriter, r *http.Request) {
	form := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Title of the document</title>
</head>

<body>

<form action="/save" method="POST">
  Nome: <input type="text" name="nome"><br>
  <input type="submit" value="Enviar">
</form>

</body>

</html>`
	fmt.Fprintf(w, form)

}

func read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "read")
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete")
}

func save(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		msg := http.StatusText(http.StatusMethodNotAllowed)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Apenas metodo POST permitido.")
		return
	}

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	nome := strings.Join(r.Form["nome"], "")
	if nome == "" {
		msg := http.StatusText(http.StatusForbidden)
		http.Error(w, msg, http.StatusForbidden)
		fmt.Fprintf(w, "Campo Nome deve ser preenchido.")
		return
	}

	fmt.Fprintf(w, "save")

}
