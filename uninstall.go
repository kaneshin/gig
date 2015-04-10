package main

import (
	"fmt"
)

var cmdUninstall = &command{
	usage: "uninstall <respository>",
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
		dest := gigDir(repo)
		if _, err := run([]string{"rm", "-rf", dest}); err != nil {
			return err
		}
		fmt.Fprintln(stdout, "Deleted "+dest)
	default:
		return cmdUninstall.err()
	}
	return nil
}
