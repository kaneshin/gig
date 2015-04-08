package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const (
	gigDir = ".gig"
)

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
	if flag.NArg() == 0 {
		usage()
	}

	name := flag.Arg(0)
	for _, cmd := range commands {
		if cmd.name() == name {
			if err := cmd.run(flag.Args()[1:]); err != nil {
				fmt.Fprintln(stderr, "ERROR: ", err)
				os.Exit(1)
			}
			os.Exit(0)
		}
	}
	fmt.Fprintln(stderr, "ERROR: Command Not Found")
	os.Exit(1)
}
