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
    install                     Install repository.
    uninstall                   Uninstall repository.
    list                        List installed repository.
    upgrade                     Upgrade repository.
    version                     Show version and exit.

  Examples:
    gig install kaneshin/gig --host github ~/.gig

  General Options:
    --verbose                   Give more output.
    --host
`)
	os.Exit(1)
}
