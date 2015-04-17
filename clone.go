package main

import (
	"fmt"
	"github.com/kaneshin/gig/repository"
	"strings"
)

var cmdClone = &command{
	usage: "clone <repository> [options...]",
}

func init() {
	cmdClone.run = clone
}

func clone(args []string) error {
	var isGitHub = true
	var transact = func(rname string) error {
		info := strings.Split(rname, "/")
		if len(info) < 2 {
			return cmdClone.err()
		}
		var rep repository.Repository
		if isGitHub {
			rep = repository.NewGitHubRepository(info[0], info[1])
		} else {
			rep = repository.NewBitBucketRepository(info[0], info[1])
		}
		dest := gigDir(rep.Identifier())
		elm := []string{
			"git", "clone", rep.URL(), dest,
		}
		if _, err := run(elm); err != nil {
			return err
		}
		fmt.Fprintln(stdout, "Cloned into "+dest)
		return nil
	}
	switch len(args) {
	case 0:
		return cmdClone.err()
	case 1:
		break
	default:
		opt := args[1:]
		var skip = false
		for idx, val := range opt {
			if skip {
				skip = false
				continue
			}
			switch val {
			case "--protocol":
				fallthrough
			case "-p":
				if len(opt) > idx+1 {
					if s := opt[idx+1]; s == "ssh" {
						repository.UseSSH()
					} else {
						repository.UseHTTPS()
					}
					skip = true
				}
			case "--host":
				fallthrough
			case "-h":
				if len(opt) > idx+1 {
					if s := opt[idx+1]; s == "bitbucket" {
						isGitHub = false
					} else {
						isGitHub = true
					}
					skip = true
				}
			}
		}
	}
	return transact(args[0])
}
