package admin

import (
	"github.com/go-one/gone"
)

type AdminController struct {
	*gone.Controller
}
type AA AdminController
func (a *AA) Index(l, b int) error {
	return nil
}
