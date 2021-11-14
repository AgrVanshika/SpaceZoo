package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	L "github.com/AOrps/spaceZoo/src/lib"
)

const PORT = 9990

/*
visualization
relaxation
reflection
expression
breathing
meditation
*/

func visualization(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "start", nil)

	tmpl.ExecuteTemplate(w, "header", nil)

	tmpl.ExecuteTemplate(w, "startbody", nil)

	navlinks := L.NavBarFill()
	tmpl.ExecuteTemplate(w, "navbar", navlinks)

	tmpl.ExecuteTemplate(w, "startcontain", nil)

	tmpl.ExecuteTemplate(w, "startcontent", nil)

	techniques := L.ListMethods("visualization")
	tmpl.ExecuteTemplate(w, "method", techniques)

	tmpl.ExecuteTemplate(w, "endcontent", nil)

	tmpl.ExecuteTemplate(w, "endcontain", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)

	tmpl.ExecuteTemplate(w, "end", nil)
}

func relaxation(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "start", nil)

	tmpl.ExecuteTemplate(w, "header", nil)

	tmpl.ExecuteTemplate(w, "startbody", nil)

	navlinks := L.NavBarFill()
	tmpl.ExecuteTemplate(w, "navbar", navlinks)

	tmpl.ExecuteTemplate(w, "startcontain", nil)

	tmpl.ExecuteTemplate(w, "startcontent", nil)

	techniques := L.ListMethods("relaxation")
	tmpl.ExecuteTemplate(w, "method", techniques)

	tmpl.ExecuteTemplate(w, "endcontent", nil)

	tmpl.ExecuteTemplate(w, "endcontain", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)

	tmpl.ExecuteTemplate(w, "end", nil)
}

func reflection(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "start", nil)

	tmpl.ExecuteTemplate(w, "header", nil)

	tmpl.ExecuteTemplate(w, "startbody", nil)

	navlinks := L.NavBarFill()
	tmpl.ExecuteTemplate(w, "navbar", navlinks)

	tmpl.ExecuteTemplate(w, "startcontain", nil)

	tmpl.ExecuteTemplate(w, "startcontent", nil)

	techniques := L.ListMethods("reflection")
	tmpl.ExecuteTemplate(w, "method", techniques)

	tmpl.ExecuteTemplate(w, "endcontent", nil)

	tmpl.ExecuteTemplate(w, "endcontain", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)

	tmpl.ExecuteTemplate(w, "end", nil)
}

func expression(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "start", nil)

	tmpl.ExecuteTemplate(w, "header", nil)

	tmpl.ExecuteTemplate(w, "startbody", nil)

	navlinks := L.NavBarFill()
	tmpl.ExecuteTemplate(w, "navbar", navlinks)

	tmpl.ExecuteTemplate(w, "startcontain", nil)

	tmpl.ExecuteTemplate(w, "startcontent", nil)

	techniques := L.ListMethods("expression")
	tmpl.ExecuteTemplate(w, "method", techniques)

	tmpl.ExecuteTemplate(w, "endcontent", nil)

	tmpl.ExecuteTemplate(w, "endcontain", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)

	tmpl.ExecuteTemplate(w, "end", nil)
}

func breathing(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "start", nil)

	tmpl.ExecuteTemplate(w, "header", nil)

	tmpl.ExecuteTemplate(w, "startbody", nil)

	navlinks := L.NavBarFill()
	tmpl.ExecuteTemplate(w, "navbar", navlinks)

	tmpl.ExecuteTemplate(w, "startcontain", nil)

	tmpl.ExecuteTemplate(w, "startcontent", nil)

	techniques := L.ListMethods("breathing")
	tmpl.ExecuteTemplate(w, "method", techniques)

	tmpl.ExecuteTemplate(w, "endcontent", nil)

	tmpl.ExecuteTemplate(w, "endcontain", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)

	tmpl.ExecuteTemplate(w, "end", nil)
}

func meditation(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "start", nil)

	tmpl.ExecuteTemplate(w, "header", nil)

	tmpl.ExecuteTemplate(w, "startbody", nil)

	navlinks := L.NavBarFill()
	tmpl.ExecuteTemplate(w, "navbar", navlinks)

	tmpl.ExecuteTemplate(w, "startcontain", nil)

	tmpl.ExecuteTemplate(w, "startcontent", nil)

	techniques := L.ListMethods("meditation")
	tmpl.ExecuteTemplate(w, "method", techniques)

	tmpl.ExecuteTemplate(w, "endcontent", nil)

	tmpl.ExecuteTemplate(w, "endcontain", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)

	tmpl.ExecuteTemplate(w, "end", nil)
}

func mainsite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "start", nil)

	tmpl.ExecuteTemplate(w, "header", nil)

	tmpl.ExecuteTemplate(w, "startbody", nil)

	navlinks := L.NavBarFill()
	tmpl.ExecuteTemplate(w, "navbar", navlinks)

	tmpl.ExecuteTemplate(w, "startcontain", nil)

	tmpl.ExecuteTemplate(w, "startcontent", nil)

	con := L.FillContent()
	tmpl.ExecuteTemplate(w, "content", con)

	tmpl.ExecuteTemplate(w, "endcontain", nil)

	tmpl.ExecuteTemplate(w, "endcontain", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)

	tmpl.ExecuteTemplate(w, "end", nil)
}

func amongus(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "/assets/website/img/logo")
	// tmpl.ExecuteTemplate(w, "amongus", nil)
}

func webserver(port string) {
	fmt.Printf("Server on port %s\n", port)

	fs := http.FileServer(http.Dir("."))

	// Puts everything from File Server into a /assets/ directory
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", mainsite)

	http.HandleFunc("/visualization", visualization)
	http.HandleFunc("/relaxation", relaxation)
	http.HandleFunc("/reflection", reflection)
	http.HandleFunc("/expression", expression)
	http.HandleFunc("/breathing", breathing)
	http.HandleFunc("/meditation", meditation)

	http.HandleFunc("/amongus")

	err := http.ListenAndServe(":"+port, nil)

	log.Fatal(err)
}

func main() {

	port := strconv.Itoa(PORT)

	webserver(port)
}
