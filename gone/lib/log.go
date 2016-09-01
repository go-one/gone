package lib

import (
	"fmt"
	"strings"

	"time"

	"errors"
	"runtime"

	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

const logDebug = false
const buildDate = false
const asciiName = `
 ██████╗  ██████╗ ███╗   ██╗███████╗
██╔════╝ ██╔═══██╗████╗  ██║██╔════╝
██║  ███╗██║   ██║██╔██╗ ██║█████╗
██║   ██║██║   ██║██║╚██╗██║██╔══╝
╚██████╔╝╚██████╔╝██║ ╚████║███████╗
 ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚══════╝
 Fastest go web framework
 `

var logOffset = ""
var logOffsetN = 0

func IncrLogOffset() {
	logOffsetN++
	logOffset = strings.Repeat("\t", logOffsetN)
}
func DecrLogOffset() {
	if logOffsetN >= 0 {
		logOffsetN--
		logOffset = strings.Repeat("\t", logOffsetN)
	}
}
func PrefixString(place, s string, attrs ...color.Attribute) string {
	timeS := time.Now().Format("2006/01/02 15:04:05")
	return timeS + " " + place + "\t" + logOffset
}
func InfoLog(s string, args ...interface{}) {
	var prefix = ""
	if logDebug {
		prefix = PrefixString(LogPlace(), "INFO", color.Bold, color.FgWhite)
	}
	Log(prefix, fmt.Sprintf(s, args...))
}

// Log split string by \n, and
func Log(prefix, s string) {
	strs := strings.Split(s, "\n")
	for _, value := range strs {
		fmt.Println(prefix, value)
	}
}
func LogPlace() string {
	_, cf, cl, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Something went wrong"))
	}
	return filepath.Base(cf) + ":" + strconv.FormatInt(int64(cl), 10)
}

func ShowBanner(c *cli.Context) error {
	InfoLog(asciiName)
	InfoLog("Version: " + c.App.Version)
	InfoLog("Build date: " + BuildDate)
	return nil
}
