package repository

const bitbucketDomainName = "bitbucket.org"

func NewBitBucketRepository(username, reponame string) Repository {
	r := newRepository(username, reponame)
	r.id = bitbucket
	return r
}
