package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const (
	defURI = "https://github.com"
	defDst = ".gig"
)

func usage() {
	fmt.Printf(`
Usage:
  gig <command> [options] <repository-path>

Commands:
  install                     Install repository.
  uninstall                   Uninstall repository.
  list                        List installed repository.
  upgrade                     Upgrade repository.

General Options:
  -v, --verbose               Give more output.
  -V, --version               Show version and exit.
`)
	os.Exit(1)
}

var (
	stdout = os.Stdout
	stderr = os.Stderr
	stdin  = os.Stdin
)

func run(args []string) error {
	if len(args) == 0 {
		usage()
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Stdin = stdin
	err := cmd.Run()
	return err
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
	}

	switch flag.Arg(0) {
	case "install":
		repos := flag.Args()[1:]
		for _, repo := range repos {
			uri := defURI + "/" + repo + ".git"
			dst := os.Getenv("HOME") + "/" + defDst + "/" + repo
			if err := run([]string{"git", "clone", "--quiet", uri, dst}); err != nil {
				fmt.Fprintln(stderr, os.Args[0], ": ", err)
				os.Exit(1)
			}
			fmt.Fprintln(stdout, "Cloned into "+dst)
		}
	case "uninstall":
		repos := flag.Args()[1:]
		for _, repo := range repos {
			dst := os.Getenv("HOME") + "/" + defDst + "/" + repo
			if err := run([]string{"rm", "-rf", dst}); err != nil {
				fmt.Fprintln(os.Stderr, os.Args[0], ": ", err)
			}
		}
	case "list":

	case "upgrade":

	default:
		usage()
	}
}
