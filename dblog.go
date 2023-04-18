package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"
)

type DocusaurusBlog struct {
	InPath  string // input path
	OutPath string // output path
	Slug    string
	Title   string
	Date    string // 2023-04-17 for filename
	Authors []string
	Tags    []string //
	Content string   // rendered by template
}

func NewDocusaurusBlog(path string) *DocusaurusBlog {
	return &DocusaurusBlog{
		OutPath: path,
	}
}

// ReadBlog read blog from file and parse it to DocusaurusBlog
func (b *DocusaurusBlog) ReadBlog() error {
	return nil
}

// WriteBlog write blog to file from DocusaurusBlog
func (b *DocusaurusBlog) WriteBlog(dryrun bool) error {
	//fmt.Printf("DocusaurusBlog: %+v\n", b)
	fileName := strings.Split(b.Date, " ")[0] + "-" + b.Slug + ".md"
	outputFilePath := b.OutPath + "/" + fileName
	tags := ""
	for _, tag := range b.Tags {
		tags += "- " + tag + "\n"
	}
	outputContent := fmt.Sprintf("---\n"+
		"slug: %v\n"+
		"title: %v\n"+
		"authors: %v\n"+
		"tags: \n%v"+
		"---\n"+
		"%v", b.Slug, b.Title, b.Authors, tags, b.Content)
	if dryrun {
		fmt.Printf("INFO: [DRY-RUN] Write Docusaurus Blog [%s] Success!\n", b.Slug)
	} else {
		if err := ioutil.WriteFile(outputFilePath, []byte(outputContent), fs.ModePerm); err != nil {
			return err
		}
		fmt.Printf("INFO: Write Docusaurus Blog [%s] Success!\n", b.Slug)
	}
	return nil
}
