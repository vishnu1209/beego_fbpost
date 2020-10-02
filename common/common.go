package common

import (
	"os"
	"strings"
)

func init() {
	cwd := os.Getenv("PWD")
	rootDir := strings.Split(cwd, "tests/")
	os.Args[0] = rootDir[0] // path to you dir
}
