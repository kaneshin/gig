package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf(`gig

  Usage:
    gig -h/--help
    gig <command> [arguments...] [options...]

  Commands:
    clone                       Clone repository.
    uninstall                   Uninstall repository.
    list                        List installed repository.
    upgrade                     Upgrade repository.
    version                     Show version and exit.

  Examples:
    gig clone kaneshin/gig --host bitbucket

  General Options:
    --verbose                   Give more output.
    -h,--host
    -p,--protocol
`)
	os.Exit(1)
}
