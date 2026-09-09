package main

import (
	"flag"
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Сообщаем браузеру, что передаем HTML-разметку
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	name := r.URL.Query().Get("name")

	fmt.Fprintf(w, `
		<form method="GET">
			<input type="text" name="name" placeholder="Enter your name" value="%s">
			<button type="submit">Submit</button>
		</form>
	`, name)

	if name != "" {
		fmt.Fprintf(w, "<h3>Hello, %s!</h3>", name)
	}
}

func main() {
	port := flag.String("port", "8080", "порт для запуска сервера")
	flag.Parse()

	http.HandleFunc("/form", formHandler)

	addr := ":" + *port
	fmt.Printf("Сервер запущен. Откройте в браузере http://localhost%s/form\n", addr)

	http.ListenAndServe(addr, nil)
}
