package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func usage() {
	fmt.Printf(`Usage of %s:
gig install /path/to/repo
gig update  /path/to/repo
`, os.Args[0])
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 2 {
		usage()
	}

	switch flag.Arg(0) {
	case "install":
		repo := "https://github.com/" + flag.Arg(1) + ".git"
		dst := os.Getenv("HOME") + "/.gig/" + flag.Arg(1)
		cmd := exec.Command("git", "clone", repo, dst)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	case "update":
		usage()
	default:
		usage()
	}
}
