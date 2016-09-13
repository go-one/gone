package lib

import (
	"github.com/saturn4er/go-parse-types"
)

type Controller struct {
	Name             string
	PkgPath          string
	PkgAlias         string
	PathToController string
	PtrController    bool
	Type             *tparser.Type
}

