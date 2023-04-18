package main

type Blog interface {
	// ReadBlog() error
	ReadBlog() error
	// WriteBlog() error
	WriteBlog(dryrun bool) error
}
