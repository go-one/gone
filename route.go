package gone

import (
	"github.com/gorilla/pat"
)
var Router *pat.Router

func init() {
	Router = pat.New()
}
