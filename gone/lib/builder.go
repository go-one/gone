package lib

import (
	"errors"
)

type Application struct {
	configPath string
}

func (a *Application) Build() error {
	InfoLog("Build:")
	IncrLogOffset()
	defer DecrLogOffset()
	InfoLog("Loading config...")
	panic(errors.New("Not implemented yet"))
	return nil
}

func (a *Application) Run() error {
	panic(errors.New("Not implemented yet"))
	return nil
}

func NewBuilder(configPath string) *Application {
	result := new(Application)
	result.configPath = configPath
	return result
}
