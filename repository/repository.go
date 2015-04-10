package repository

import (
	"bytes"
)

const (
	github = iota
	bitbucket
)

type Repository interface {
	URL() string
	Identifier() string
}

type repositoryInfo struct {
	id       int
	username string
	reponame string
}

func (r repositoryInfo) URL() string {
	return specifyURL(r.domainName(), r.username, r.reponame)
}

func (r repositoryInfo) Identifier() string {
	var buf bytes.Buffer
	buf.WriteString(r.username)
	buf.WriteString("/")
	buf.WriteString(r.reponame)
	return buf.String()
}

func (r repositoryInfo) domainName() string {
	switch r.id {
	case github:
		return githubDomainName
	case bitbucket:
		return bitbucketDomainName
	}
	return githubDomainName
}

func newRepository(username, reponame string) repositoryInfo {
	r := repositoryInfo{
		username: username,
		reponame: reponame,
	}
	return r
}
