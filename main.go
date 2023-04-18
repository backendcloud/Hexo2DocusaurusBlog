package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	//h := NewHexoBlog("./input/hexo-blog-for-test.md")
	//err := h.ReadBlog()
	//if err != nil {
	//	panic(err)
	//}
	//d, err := h.toDocusaurusBlog([]string{"cloud", "Bob"})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("d %+v\n", d)

	hexoBlogDir := flag.String("hexo", "", "hexo blog directory")
	docBlogDir := flag.String("docusaurus", "", "docusaurus blog directory")
	author := flag.String("author", "", "blog author name")
	dryrun := flag.Bool("dry-run", false, "perform a trial run with no changes made")
	flag.Parse()

	if len(*hexoBlogDir) == 0 || len(*docBlogDir) == 0 {
		fmt.Fprintln(os.Stderr, "blog dir is empty")
		return
	}
	if len(*author) == 0 {
		fmt.Fprintln(os.Stderr, "author is empty")
		return
	}

	authors := []string{}
	// 将string类型的author，用逗号分割，转换为[]string类型
	for _, a := range strings.Split(*author, ",") {
		authors = append(authors, strings.TrimSpace(a))
	}

	fs, err := os.ReadDir(*hexoBlogDir)
	if err != nil {
		panic(err)
	}

	var blogFiles []string
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".DS_Store") {
			continue
		}

		if f.IsDir() {
			continue
		}

		if strings.HasSuffix(f.Name(), ".md") {
			blogFiles = append(blogFiles, f.Name())
		}
	}

	var hexoBlogs []*HexoBlog
	for _, filename := range blogFiles {
		hexoblog := NewHexoBlog(*hexoBlogDir + "/" + filename)
		if err := hexoblog.ReadBlog(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
		hexoBlogs = append(hexoBlogs, hexoblog)
		docusaurusblog, err := hexoblog.toDocusaurusBlog(authors, *docBlogDir)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
		if err := docusaurusblog.WriteBlog(*dryrun); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
	}

}
