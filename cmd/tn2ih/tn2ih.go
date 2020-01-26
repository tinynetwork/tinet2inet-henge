package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/ak1ra24/tn2inethenge/pkg/converter"
)

var configFile string

func Handler(w http.ResponseWriter, r *http.Request) {

	ihJson, err := converter.Tn2inethenge(configFile)
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.ParseFiles("./templates/index.tpl"))

	if err := t.Execute(w, ihJson); err != nil {
		log.Fatal(err)
	}
}

func debug() {
	ihJson, err := converter.Tn2inethenge(configFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ihJson)

	t := template.Must(template.ParseFiles("./templates/index.tpl"))

	if err := t.Execute(os.Stdout, ihJson); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("tn2ih <tn config file>")
		os.Exit(1)
	} else {
		configFile = args[0]
		fmt.Println("Config File: ", configFile)
	}
	log.Println("Running: http://localhost:8080/")

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
	// debug()

}
