package main

import "fmt"

func (b *HexoBlog) toDocusaurusBlog(authors []string, outPath string) (*DocusaurusBlog, error) {
	if b == nil {
		return nil, ErrHexoBlogStructNil
	}

	//authors := []string{}
	//authors = append(authors, "cloud")
	tags := []string{}

	if len(b.Tags) > 0 {
		tags = b.Tags
	}
	fmt.Printf("INFO: Convert Hexo Blog [%s] to Docusaurus Blogs Success!\n", b.Slug)

	return &DocusaurusBlog{
		OutPath: outPath,
		Slug:    b.Slug,
		Title:   b.Title,
		Date:    b.Date,
		Authors: authors,
		Tags:    tags,
		Content: b.Content,
	}, nil
}
