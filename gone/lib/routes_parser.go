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
	HandlerController string
	HandlerAction     string
	Alias             string
}

var routesRegex = regexp.MustCompile("^(.+?)(?:\t| )+(.+?)(?:\t| )+(.+?)\\.(.+?)(?:(?:\t| )+(.+?))?$")

func ParseRoutes(path string) ([]Route, error) {
	InfoLog("Parsing routes from %s", path)
	IncrLogOffset()
	defer DecrLogOffset()
	routesFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer routesFile.Close()
	result := []Route{}
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
		route := Route{}
		route.Route = sParams[1]
		if len(sParams) == 6 {
			route.Alias = sParams[5]
		}
		hm := strings.Split(sParams[2], ",")
		for i, method := range hm {
			hm[i] = strings.ToLower(method)
		}
		route.HTPPMethods = hm
		route.HandlerController = sParams[3]
		route.HandlerAction = sParams[4]
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
