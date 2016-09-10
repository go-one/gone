package lib

import (
	"errors"
	"path/filepath"

	"strings"

	"gopkg.in/ini.v1"
)

type ApplicationConfig struct {
	BuildPath        string
	ControllersPaths []string
	RoutesPaths      []string
}
type Application struct {
	path   string
	config ApplicationConfig
}

func (a *Application) loadConfig() error {

	InfoLog("Loading config %s", "build.conf")
	IncrLogOffset()
	defer DecrLogOffset()
	cfg, err := ini.LoadSources(ini.LoadOptions{Insensitive: true}, filepath.Join(a.path, "build.conf"))
	if err != nil {
		InfoLog("Using default config due to error: %v", err)
		return err
	}
	s, err := cfg.GetSection("")
	if err != nil {
		InfoLog("Using default config due to error: %v", err)
	}
	// Build output
	if s.HasKey("output") {
		oKey, _ := s.GetKey("output")
		a.config.BuildPath = addExeIfWindows(oKey.String())
	} else {
		a.config.BuildPath = addExeIfWindows(filepath.Base(a.path))
	}
	InfoLog("Build output:                 %s", a.config.BuildPath)

	// Controllers paths
	a.config.ControllersPaths = []string{"controllers"}
	if s.HasKey("controllers") {
		oKey, _ := s.GetKey("controllers")
		a.config.ControllersPaths = append(a.config.ControllersPaths, oKey.Strings(",")...)
	}
	InfoLog("Controllers directories:      %v", strings.Join(a.config.ControllersPaths, ", "))

	// Routes file
	a.config.RoutesPaths = []string{"config/routes"}
	if s.HasKey("routes") {
		oKey, _ := s.GetKey("routes")
		a.config.RoutesPaths = append(a.config.RoutesPaths, oKey.Strings(",")...)
	}
	InfoLog("Routes:                       %v", strings.Join(a.config.RoutesPaths, ", "))
	return nil
}
func (a *Application) parseRoutes() {
	InfoLog("Parsing routes")
	IncrLogOffset()
	defer DecrLogOffset()
	for _, p := range a.config.RoutesPaths {
		_, err := ParseRoutes(p)
		if err != nil {
			ErrorLog("Error while parsing routes: %v", err)
		}
	}
}
func (a *Application) Build() error {

	a.loadConfig()
	a.parseRoutes()
	InfoLog("Building:")
	IncrLogOffset()
	defer DecrLogOffset()
	//panic(errors.New("Not implemented yet"))
	return nil
}
func (a *Application) Run() error {
	panic(errors.New("Not implemented yet"))
	return nil
}

func NewApplication(path string) *Application {
	result := new(Application)
	result.path = path
	return result
}

const AppTemplate = `

`
