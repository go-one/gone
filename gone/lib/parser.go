package lib

// Package main provides
import (
	"errors"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"unicode"
)

type Alloha struct {
}

func (a *Alloha) name() int {
	fmt.Println(123)
	return 1
}
func (a *Alloha) asd(dasd int) {
	fmt.Println(dasd)
}

type MethodArgument struct {
	Name string
	Type string
}
type Method struct {
	Name, GGs string
	Arguments []MethodArgument
	Doc       string
}

func (m ((*Method))) HHH(a int) {
	fmt.Println(123)
}

type Type struct {
	PkgName string
	Name    string
	Doc     string
	Methods []Method
}

// GetFile asd
func GetTypesInDir(path string) []Type {
	fset := token.NewFileSet() // positions are relative to fset

	d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	result := []Type{}

	for _, f := range d {
		ParsePackageTypes(path, fset, f)
	}
	return result

}
func ParsePackageTypes(path string, fset *token.FileSet, pkg *ast.Package) {
	p := doc.New(pkg, path, 0)
	ast.Inspect(pkg, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Recv == nil || len(x.Recv.List) == 0 {
				// We are looking only for functions with receiver
				return true
			}
			ParseMethod(x)
			os.Exit(1)
		//fmt.Printf("%s:\tFuncDecl %s\t%s\n", fset.Position(n.Pos()).String(), x.Name, x.Recv.List[0].Names, x.Doc.Text())
		case *ast.TypeSpec:
		//fmt.Printf("%s:\tTypeSpec %s\t%s\n", fset.Position(n.Pos()), x.Name, x.Doc.Text())
		case *ast.Field:
		//fmt.Printf("%s:\tField %s\t%s\n", fset.Position(n.Pos()), x.Names, x.Doc.Text())
		case *ast.GenDecl:
		//fmt.Printf("%s:\tGenDecl %s\n", fset.Position(n.Pos()), x.Doc.Text())
		case *ast.FuncType:
			//fmt.Printf("%s:\tGenDecl %s\n", fset.Position(n.Pos()), x.Func)
		}
		return true
	})
	for _, t := range p.Types {
		fmt.Println(t.Decl.Pos(), t.Decl.End())
		fmt.Println("  type", t.Name)
		fmt.Println("    docs:", t.Doc)
	}
}
func ParseMethod(f *ast.FuncDecl) (string, *Method, error) {
	var typeName string
	result := new(Method)

	// Parsing receiver type
	receiver := f.Recv.List[0]
	switch t := receiver.Type.(type) {
	case *ast.StarExpr:
		if i, ok := t.X.(*ast.Ident); ok {
			typeName = i.Name
		} else {
			fmt.Println("Unknown *ast.StarExpr in receiver")
		}
	case *ast.Ident:
		typeName = t.Name
	default:
		fmt.Printf("Found unknown receiver Type %T", receiver.Type)
	}
	if len(typeName) == 0 {
		return "", nil, errors.New("Can't parse function receiver type")
	}
	if unicode.IsLower(rune(typeName[0])) {
		return "", nil, errors.New("We doesn't support non-exportable controllers")
	}
	// Parsing method arguments
	if f.Type.Params != nil {
		result.Arguments = ParseMethodParams(f.Type.Params)
	}
	fmt.Println("Found method", f.Name, "for type", typeName, result.Arguments)
	return typeName, result, nil
}
func ParseMethodParams(f *ast.FieldList) []MethodArgument {
	result := []MethodArgument{}
	for _, arg := range f.List {
		r := MethodArgument{}
		r.Type = fmt.Sprintf("%v", arg.Type)
		if len(arg.Names) == 0 {
			result = append(result, r)
			continue
		}
		for _, n := range arg.Names {
			r.Name = n.Name
			result = append(result, r)
		}
	}
	return result
}
func main() {
	GetTypesInDir("./")
}
