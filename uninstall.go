package main

import (
	"fmt"
	"os"
)

var cmdUninstall = &command{
	usage: "gig uninstall kaneshin/gig",
}

func init() {
	cmdUninstall.run = uninstall
}

func uninstall(args []string) error {
	switch len(args) {
	case 0:
		return cmdUninstall.err()
	case 1:
		repo := args[0]
		dir := os.Getenv("HOME") + "/" + gigDir + "/" + repo
		if err := run([]string{"rm", "-rf", dir}); err != nil {
			return err
		}
		fmt.Fprintln(stdout, "Deleted "+dir)
	default:
		return cmdUninstall.err()
	}
	return nil
}
