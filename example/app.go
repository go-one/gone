package main

import (
	"net/http"
	"github.com/go-one/gone"
	"github.com/gorilla/pat"

	vlnmq "github.com/go-one/gone/example/app/controllers"

	nilwe "github.com/go-one/gone/example/app/controllers/admin"

)
//regexps
func main() {
	router := pat.New()


	router.Add("GET", "/test/{id}", http.HandlerFunc(vlnmq_ApplicationController_Index))


	http.Handle("/", router)
	http.ListenAndServe(":8085", nil)

}


func vlnmq_ApplicationController_Index(res http.ResponseWriter, req *http.Request){
	controller := new(vlnmq.ApplicationController)

	goneController := &controller.Controller
	goneController.ControllerName = "ApplicationController"
	goneController.ActionName = "Index"
	goneController.TplPath = "./ApplicationController/Index"

	goneController.Method = req.Method
	goneController.Request = req
	goneController.Response = res
	controller.Index()
}



func vlnmq_D_Index1(res http.ResponseWriter, req *http.Request){
	controller := new(vlnmq.D)

	goneController := &controller.ApplicationController.Controller
	goneController.ControllerName = "D"
	goneController.ActionName = "Index1"
	goneController.TplPath = "./D/Index1"

	goneController.Method = req.Method
	goneController.Request = req
	goneController.Response = res
	controller.Index1()
}

func vlnmq_D_Index2(res http.ResponseWriter, req *http.Request){
	controller := new(vlnmq.D)

	goneController := &controller.ApplicationController.Controller
	goneController.ControllerName = "D"
	goneController.ActionName = "Index2"
	goneController.TplPath = "./D/Index2"

	goneController.Method = req.Method
	goneController.Request = req
	goneController.Response = res
	controller.Index2()
}

func vlnmq_D_Index3(res http.ResponseWriter, req *http.Request){
	controller := new(vlnmq.D)

	goneController := &controller.ApplicationController.Controller
	goneController.ControllerName = "D"
	goneController.ActionName = "Index3"
	goneController.TplPath = "./D/Index3"

	goneController.Method = req.Method
	goneController.Request = req
	goneController.Response = res
	controller.Index3()
}





func nilwe_AA_Index(res http.ResponseWriter, req *http.Request){
	controller := new(nilwe.AA)

	goneController := &gone.Controller{
		ControllerName: "AA",
		ActionName: "Index",
		TplPath: "./AA/Index",
	}
	controller.Controller = goneController

	goneController.Method = req.Method
	goneController.Request = req
	goneController.Response = res

	controller.Index()
}
