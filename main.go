package main

import (
	"log"
	"net/http"
	"regexp"
	"text/template"
)

var (
	// Regular expression to replace invalid characters
	onlyNumber = regexp.MustCompile("[^0-9]")
)

// Provides the page application on port 8080
func main() {
	http.HandleFunc("/", ZipCodeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// renderPage renders the app page
func renderPage(w http.ResponseWriter, data map[string]interface{}) error {
	t := template.Must(template.ParseFiles("index.html"))
	return t.Execute(w, data)
}

// ZipCodeHandler is the main handler of application
// When it is access with GET, renders the search zipcode page.
// When it is access with POST, gets the address from Digipix api and shows the result
func ZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	logServer("Incoming request...")
	switch r.Method {
	case "GET":
		// GET: renders the page
		logServer("Render page...")
		if err := renderPage(w, nil); err != nil {
			log.Printf("%v", err)
			http.Error(w, "Erro ao exibir a página", http.StatusInternalServerError)
			return
		}
	case "POST":
		zipParam := r.FormValue("zipcode")
		logServer("ZipCode received: %s", zipParam)
		pageData := make(map[string]interface{}, 0)
		pageData["Zipcode"] = zipParam
		// POST: search the address data given a zipcode
		zip := onlyNumber.ReplaceAllString(zipParam, "")
		if len(zip) != 8 {
			logServer("Invalid ZipCode")
			pageData["Error"] = "CEP inválido"
			renderPage(w, pageData)
			return
		}
		addr, err := GetAddress(zip)
		if err != nil {
			logServer("Unable to get address: %s", err.Error())
			pageData["Error"] = "Ocorreu um erro. Por favor tente novamente"
			renderPage(w, pageData)
			return
		}
		if addr == nil {
			logServer("Address not found")
			pageData["Error"] = "Endereço não encontrado"
			renderPage(w, pageData)
			return
		}
		logServer("Address found!")
		pageData["Address"] = *addr
		renderPage(w, pageData)
	default:
		http.Error(w, "Erro interno", http.StatusMethodNotAllowed)
	}
}
