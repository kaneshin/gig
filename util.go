package main

import (
	"os"
)

var (
	stdout = os.Stdout
	stderr = os.Stderr
	stdin  = os.Stdin
)

func gigDir(path string) string {
	return os.Getenv("HOME") + "/.gig/" + path
}
