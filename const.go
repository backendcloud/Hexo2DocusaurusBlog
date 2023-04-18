package main

import "fmt"

var (
	ErrInvalidBlogFormat = fmt.Errorf("invalid blog format")
	ErrHexoBlogStructNil = fmt.Errorf("Hexo Blog Struct is nil")
)
