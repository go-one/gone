package lib

import "runtime"

func addExeIfWindows(path string) string {
	if runtime.GOOS == "windows" {
		return path + ".exe"
	}
	return path
}
