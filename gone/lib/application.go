package lib

import (
	"errors"
	"path/filepath"

	"strings"

	"os"

	"fmt"

	"github.com/saturn4er/go-parse-types"
	"gopkg.in/ini.v1"
)

var ControllerTypeName = "Controller"
var ControllerPkgName = "gone"
var ControllerPkgPath = "github.com/go-one/gone"

type ApplicationConfig struct {
	BuildPath        string
	ControllersPaths []string
	RoutesPaths      []string
}
type Application struct {
	path        string
	config      ApplicationConfig
	Controllers []*Controller
	Routes      []*Route
	Packages    map[string]string
}

func (a *Application) LoadConfig() {
	InfoLog("Loading config %s", "build.conf")
	IncrLogOffset()
	defer DecrLogOffset()
	a.setDefaultConfig()
	cfg, err := ini.LoadSources(ini.LoadOptions{Insensitive: true}, filepath.Join(a.path, "build.conf"))
	if err != nil {
		InfoLog("Using default config due to error: %v", err)
		a.logConfig()
		return
	}
	s, err := cfg.GetSection("")
	if err != nil {
		InfoLog("Using default config due to error: %v", err)
		a.logConfig()
		return
	}
	// Build output
	if s.HasKey("output") {
		oKey, _ := s.GetKey("output")
		a.config.BuildPath = addExeIfWindows(oKey.String())
	}

	// Controllers paths
	if s.HasKey("controllers") {
		oKey, _ := s.GetKey("controllers")
		a.config.ControllersPaths = append(a.config.ControllersPaths, oKey.Strings(",")...)
	}

	// Routes file
	if s.HasKey("routes") {
		oKey, _ := s.GetKey("routes")
		a.config.RoutesPaths = append(a.config.RoutesPaths, oKey.Strings(",")...)
	}
	a.logConfig()
}
func (a *Application) setDefaultConfig() {
	a.config.BuildPath = addExeIfWindows(filepath.Base(a.path))
	a.config.RoutesPaths = []string{"config/routes"}
	a.config.ControllersPaths = []string{"app/controllers"}
}
func (a *Application) logConfig() {
	InfoLog("Build output:                 %s", a.config.BuildPath)
	InfoLog("Controllers directories:      %s", strings.Join(a.config.ControllersPaths, ", "))
	InfoLog("Routes:                       %s", strings.Join(a.config.RoutesPaths, ", "))
}
func (a *Application) parseRoutes() {
	InfoLog("Parsing routes")
	IncrLogOffset()
	defer DecrLogOffset()
	a.Routes = []*Route{}
	for _, p := range a.config.RoutesPaths {
		r, err := ParseRoutes(filepath.Join(a.path, p))
		if err != nil {
			ErrorLog("Error while parsing routes: %v", err)
			continue
		}
		a.Routes = append(a.Routes, r...)
	}
}
func (a *Application) mapRoutesToControllers() error {
	InfoLog("Mapping controllers to routes")
	for _, r := range a.Routes {
		var foundController bool
		var foundAction bool
		for _, c := range a.Controllers {
			if r.HandlerPackage != "" && r.HandlerPackage == c.Type.PkgName {
				continue
			}
			if c.Name == r.HandlerController {
				foundController = true
				var m *tparser.Type
				var ok bool
				if m, ok = c.Type.GetMethod(r.HandlerAction); !ok {
					continue
				}
				if ok := checkControllerMethod(m); !ok {
					continue
				}
				r.Controller = c
				foundAction = true

			}
		}
		if !foundController {
			if r.HandlerPackage == "" {
				ErrorLog("Can't find controller: `%s`", r.HandlerController)
			} else {
				ErrorLog("Can't find controller `%s.%s`", r.HandlerPackage, r.HandlerController)
			}
			return errors.New("can't find controller")
		}
		if !foundAction {
			if r.HandlerPackage == "" {
				ErrorLog("Can't find action `%s` in controller `%s`", r.HandlerAction, r.HandlerController)
			} else {
				ErrorLog("Can't find action `%s` in controller `%s.%s`", r.HandlerAction, r.HandlerPackage, r.HandlerController)
			}
			return errors.New("can't find action")
		}

	}
	return nil
}

func (a *Application) registerPackage(path string) string {
	if _, ok := a.Packages[path]; !ok {
		a.Packages[path] = randString(5)
	}
	return a.Packages[path]
}
func (a *Application) parseSources() error {
	a.parseRoutes()
	a.parseControllers()
	err := a.mapRoutesToControllers()
	return err
}
func (a *Application) parseControllers() {
	InfoLog("Parsing controllers types")
	IncrLogOffset()
	defer DecrLogOffset()
	a.Controllers = []*Controller{}
	for _, p := range a.config.ControllersPaths {
		walkFolders(p, func(fp string, _ os.FileInfo) {
			path := filepath.Join(a.path, fp)
			InfoLog("Parsing controllers from %s", fp)
			IncrLogOffset()
			p, err := tparser.New(path)
			if err != nil {
				ErrorLog("Parsing types in %s error: %v", path, err)
				DecrLogOffset()
				return
			}
			for _, t := range p.Types {
				b, p, ok := TypeContainsController(t)
				if !ok {
					continue
				}
				pkgAlias := a.registerPackage(t.PkgPath)
				a.Controllers = append(a.Controllers, &Controller{
					PathToController: p,
					PkgPath:          t.PkgPath,
					PkgAlias:         pkgAlias,
					PtrController:    b,
					Type:             t,
					Name:             t.Name,
				})
				InfoLog("Found %s %v", t.Name, b)
			}
			DecrLogOffset()
		})
	}
}
func (a *Application) generate() error {
	code, err := GenerateSource(a)
	fmt.Println(code)
	return err
}

func (a *Application) Build() error {
	a.LoadConfig()
	err := a.parseSources()
	if err != nil {
		return nil
	}
	a.generate()
	return nil
}
func (a *Application) Run() error {
	panic(errors.New("Not implemented yet"))
	return nil
}

func NewApplication(path string) *Application {
	result := new(Application)
	result.path = path
	result.Packages = make(map[string]string)
	return result
}

func TypeContainsController(t *tparser.Type) (bool, string, bool) {
	b, p, err := t.FindFieldWithType(ControllerPkgName, ControllerPkgPath, ControllerTypeName)
	if err != nil {
		return false, "", false
	}
	return b, p, true
}
