package lib

import (
	"bytes"
	"text/template"
)

type AppSourceGenerator struct {
	*Application
}

func GenerateSource(a *Application) (string, error) {
	InfoLog("Building source code:")
	IncrLogOffset()
	defer DecrLogOffset()
	g := AppSourceGenerator{a}
	t, err := template.New("gone").Parse(AppTemplate)
	if err != nil {
		panic(err)
	}
	outB := bytes.NewBuffer([]byte{})
	err = t.Execute(outB, g)
	if err != nil {
		ErrorLog("Error: %v", err)
	}
	return outB.String(), err
}

const AppTemplate = `package main

import (
	"net/http"
	"github.com/go-one/gone"
	"github.com/gorilla/pat"
	{{  range $path, $alias := .Packages }}
	{{$alias}} "{{$path}}"
	{{ end }}
)
//regexps
func main() {
	router := pat.New()
	{{ range $r := .Routes}}
	{{ range $m := $r.HTPPMethods }}
	router.Add("{{$m}}", "{{$r.Route}}", http.HandlerFunc({{$r.Controller.PkgAlias}}_{{$r.HandlerController}}_{{$r.HandlerAction}}))
	{{ end }}
	{{ end }}
	http.Handle("/", router)
	http.ListenAndServe(":8085", nil)

}
{{ range $c := .Controllers }}
{{ range $m := .Type.Methods }}
func {{$c.PkgAlias}}_{{$c.Name}}_{{$m.Name}}(res http.ResponseWriter, req *http.Request){
	controller := new({{$c.PkgAlias}}.{{$c.Name}})
	{{ if $c.PtrController }}
	goneController := &gone.Controller{
		ControllerName: "{{$c.Name}}",
		ActionName: "{{$m.Name}}",
		TplPath: "./{{$c.Name}}/{{$m.Name}}",
	}
	controller{{$c.PathToController}} = goneController
	{{ else }}
	goneController := &controller{{$c.PathToController}}
	goneController.ControllerName = "{{$c.Name}}"
	goneController.ActionName = "{{$m.Name}}"
	goneController.TplPath = "./{{$c.Name}}/{{$m.Name}}"
	{{ end }}
	goneController.Method = req.Method
	goneController.Request = req
	goneController.Response = res
	controller.{{$m.Name}}()
}
{{ end }}
{{ end }}
`
