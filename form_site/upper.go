package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/evaluation", eval)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
func root(w http.ResponseWriter, r *http.Request) {
	rootForm, err := ioutil.ReadFile("html/rootForm.html")
	if err != nil {
		fmt.Fprint(w, "404 Not Found")
	}
	fmt.Fprint(w, string(rootForm))
}

var evaluation, _ = ioutil.ReadFile("html/evaluation.html")
var evaluationTemplate = template.Must(template.New("evaluation").Parse(string(evaluation)))

func eval(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("str")
	if strEntered == ("Andy" || "andy") {
		err := evaluationTemplate.Execute(w, "Hey, nice name buddy.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := evaluationTemplate.Execute(w, "Hey, not so nice name dude.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4748"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
