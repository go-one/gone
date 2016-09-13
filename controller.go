package gone

import "net/http"

type Controller struct {
	ControllerName string
	ActionName     string
	Method         string
	Request        *http.Request
	Response       http.ResponseWriter
	TplPath        string
}