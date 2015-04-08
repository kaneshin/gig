package main

import (
	"fmt"
	"os"
)

var cmdInstall = &command{
	usage: "install <repository> [options...]",
}

func init() {
	cmdInstall.run = install
}

func install(args []string) error {
	switch len(args) {
	case 0:
		return cmdInstall.err()
	case 1:
		repo := args[0]
		uri := githubURI + "/" + repo + ".git"
		dir := os.Getenv("HOME") + "/" + gigDir + "/" + repo
		if err := run([]string{"git", "clone", "--quiet", uri, dir}); err != nil {
			return err
		}
		fmt.Fprintln(stdout, "Cloned into "+dir)
	default:
	}
	return nil
}
