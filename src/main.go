package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const PORT = 9990

func getAbsLocation() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

/*
visualization
relaxation
reflection
expression
breathing
meditation
*/

func visualization(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "visualization!")
}

func relaxation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "relaxation!")
}

func reflection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "reflection!")
}

func expression(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "expression!")
}

func breathing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "breathing!")
}

func meditation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "meditation!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func webserver(port string) {
	fmt.Printf("Server on port %s\n", port)

	// var fileServer http.Handler

	// http.Handle("/", fileServer)

	http.HandleFunc("/", handler)

	http.HandleFunc("/visualization", visualization)
	http.HandleFunc("/relaxation", relaxation)
	http.HandleFunc("/reflection", reflection)
	http.HandleFunc("/expression", expression)
	http.HandleFunc("/breathing", breathing)
	http.HandleFunc("/meditation", meditation)

	err := http.ListenAndServe(":"+port, nil)

	log.Fatal(err)
}

func main() {

	port := strconv.Itoa(PORT)

	webserver(port)
}
