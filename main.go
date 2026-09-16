package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

var dict = map[string]map[string]string{
	"ru": {
		"switch": "en", "switch_text": "English",
		"name": "Имя (string)", "age": "Возраст (int)", "role": "Роль (enum)",
		"agree": "Согласие (bool)", "submit": "Отправить", "result": "Результат POST-запроса:",
	},
	"en": {
		"switch": "ru", "switch_text": "Русский",
		"name": "Name (string)", "age": "Age (int)", "role": "Role (enum)",
		"agree": "Agree (bool)", "submit": "Submit", "result": "POST Result:",
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	lang := r.URL.Query().Get("lang")
	if lang != "en" && lang != "ru" {
		lang = "ru"
	}
	t := dict[lang]

	getName := r.URL.Query().Get("name")

	fmt.Fprintf(w, "<p><a href=\"/?lang=%s&name=%s\">%s</a></p>", t["switch"], getName, t["switch_text"])

	if getName != "" {
		fmt.Fprintf(w, "<p>GET name: <b>%s</b></p>", getName)
	}

	fmt.Fprintf(w, "<form method=\"POST\" action=\"/?lang=%s\">"+
		"<p>%s: <input type=\"text\" name=\"name\" value=\"%s\" required></p>"+
		"<p>%s: <input type=\"number\" name=\"age\" required></p>"+
		"<p>%s: "+
		"<select name=\"role\">"+
		"<option value=\"Developer\">Developer</option>"+
		"<option value=\"QA\">QA</option>"+
		"</select>"+
		"</p>"+
		"<p><label><input type=\"checkbox\" name=\"agree\" value=\"true\"> %s</label></p>"+
		"<button type=\"submit\">%s</button>"+
		"</form>",
		lang, t["name"], getName, t["age"], t["role"], t["agree"], t["submit"])

	if r.Method == http.MethodPost {
		r.ParseForm()

		name := r.FormValue("name")
		age, _ := strconv.Atoi(r.FormValue("age"))
		role := r.FormValue("role")
		agree := r.FormValue("agree") == "true"

		fmt.Fprintf(w, "<hr><h3>%s</h3>", t["result"])
		fmt.Fprintf(w, "<p>name: %s | age: %d | role: %s | agree: %t</p>", name, age, role, agree)
	}
}

func main() {
	port := flag.String("port", "8443", "порт сервера")
	flag.Parse()

	http.HandleFunc("/", handler)

	addr := ":" + *port
	fmt.Printf("Сервер запущен: https://localhost%s\n", addr)
	http.ListenAndServeTLS(addr, "server.crt", "server.key", nil)
}
