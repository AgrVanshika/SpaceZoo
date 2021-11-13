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

func webserver(port string) {
	fmt.Printf("Server on port %s\n", port)

	var fileServer http.Handler

	http.Handle("/", fileServer)

	err := http.ListenAndServe(":"+port, nil)

	log.Fatal(err)
}

func main() {

	port := strconv.Itoa(PORT)

	webserver(port)
}
