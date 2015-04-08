package main

import (
	"errors"
	"strings"
)

type command struct {
	usage string
	run   func([]string) error
}

func (c *command) name() string {
	return strings.Fields(c.usage)[0]
}

func (c *command) err() error {
	return errors.New(`While executing gig ` + c.name() + `
Usage: ` + c.usage)
}

var commands = []*command{
	cmdInstall,
	cmdUninstall,
}
