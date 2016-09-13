package lib

import (
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"github.com/saturn4er/go-parse-types"
)

func addExeIfWindows(path string) string {
	if runtime.GOOS == "windows" {
		return path + ".exe"
	}
	return path
}

func walkFolders(path string, c func(string, os.FileInfo)) {
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			ErrorLog("Error while walking throught %s: %v", path, err)
			return err
		}
		if f.IsDir() {
			c(path, f)
		}
		return nil
	})
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randString(rslen int) string {
	b := make([]rune, rslen)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func checkControllerMethod(t *tparser.Type) bool {
	for _, value := range t.In {
		switch value.Type.Kind{
		case tparser.Bool:
		case tparser.Int:
		case tparser.Int8:
		case tparser.Int16:
		case tparser.Int32:
		case tparser.Int64:
		case tparser.Uint:
		case tparser.Uint8:
		case tparser.Uint16:
		case tparser.Uint32:
		case tparser.Uint64:
		case tparser.String:
		default:
			ErrorLog("Unsupported action in parameter type: %v", value.Type.Kind)
			return false
		}
	}

	return true
}