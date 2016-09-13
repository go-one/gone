package lib

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

type Route struct {
	Route             string
	HTPPMethods       []string
	HandlerPackage    string
	HandlerController string
	HandlerAction     string
	Alias             string
	Controller        *Controller
}

var routesRegex = regexp.MustCompile("^(.+?)(?:\t| )+(.+?)(?:\t| )+(.+?)(?:(?:\t| )+(.+?))?$")

func ParseRoutes(path string) ([]*Route, error) {
	InfoLog("Parsing routes in %s", path)
	IncrLogOffset()
	defer DecrLogOffset()
	routesFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer routesFile.Close()
	result := []*Route{}
	routesReader := bufio.NewReader(routesFile)
	i := 0
	// TODO: imports, subroutes
	for {
		i++
		l, _, err := routesReader.ReadLine()
		if err == io.EOF && len(l) == 0 {
			break
		}
		if len(l) == 0 {
			continue
		}
		line := strings.TrimSpace(string(l))
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}
		if !routesRegex.MatchString(line) {
			ErrorLog("Parsing route %s errror at line %d. Bad route.", path, i)
			continue
		}
		sParams := routesRegex.FindAllStringSubmatch(line, -1)[0]
		route := &Route{}
		route.Route = sParams[1]
		if len(sParams) == 6 {
			route.Alias = sParams[5]
		}
		hm := strings.Split(sParams[2], ",")
		for i, method := range hm {
			hm[i] = strings.ToUpper(method)
		}
		route.HTPPMethods = hm
		action := sParams[3]
		actionParts := strings.Split(action, ".")
		switch len(actionParts) {
		case 2:
			route.HandlerController = actionParts[0]
			route.HandlerAction = actionParts[1]
		case 3:
			route.HandlerPackage = actionParts[0]
			route.HandlerController = actionParts[1]
			route.HandlerAction = actionParts[2]
		default:
			ErrorLog("Parsing route %s errror at line %d. Bad action.", path, i)
			continue
		}
		result = append(result, route)
	}
	InfoLog("Found %d routes", len(result))
	return result, nil
}

func skipSpaces(line string) (int, string) {
	c := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			c++
			continue
		}
		if line[i] == '\t' {
			c += 4
			continue
		}
		return c, line[i:]
	}
	return c, ""
}
