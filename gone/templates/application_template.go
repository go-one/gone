package main

import (
	"net/http"
	"strings"
)

//regexps

func handler(res http.ResponseWriter, req *http.Request) {
	method := strings.ToLower(req.Method)
	if method == ""{
		method = "get"
	}
	// handlers

}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":{Port}", nil)
}

// handler functions
{{range .Handlers}}
func {{.Name}}(res http.ResponseWriter, req *http.Request) error{
	{{range .Lines}}

	{{end}}
}
{{end}}