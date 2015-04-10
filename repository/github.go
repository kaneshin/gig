package repository

const githubDomainName = "github.com"

func NewGitHubRepository(username, reponame string) Repository {
	r := newRepository(username, reponame)
	r.id = github
	return r
}
