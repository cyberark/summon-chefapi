package main

import (
	"os"
	"fmt"
	"regexp"
	"strings"
)

func readEnvVar(varName string) string  {
	out := os.Getenv(varName)
	if out == "" {
		printAndExit(fmt.Errorf("'%s' is not set", varName))
	}
	return out
}

func printAndExit(err error) {
	os.Stderr.Write([]byte(err.Error()))
	os.Exit(1)
}

// parsePath checks that a variable is correctly formatted to map to a Chef databag item key
// It parses a given path and returns a databag name, databag item name, and the property holding a secret
// If parsing is unsuccessful, an error is returned
func parsePath(path string) (string, string, string, error)  {
	var r = regexp.MustCompile("(\\w)+/(\\w)+/(\\w)+")
	if !r.MatchString(path) {
		return "", "", "", fmt.Errorf("'%s' is an invalid path", path)
	}

	tokens := strings.Split(path, "/")
	return tokens[0], tokens[1], tokens[2], nil
}