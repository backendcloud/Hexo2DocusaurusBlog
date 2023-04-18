package main

import (
	"bytes"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
	"time"
)

type HexoBlog struct {
	InPath  string // input path
	OutPath string // output path
	Slug    string // Similar to uuid, use the file name as the unique identifier
	Title   string
	Date    string   // 2023-04-17 12:09:29
	Tags    []string //
	Content string   // blog content
}

func NewHexoBlog(path string) *HexoBlog {
	return &HexoBlog{
		InPath: path,
	}
}

// ReadBlog read blog from file and parse it to HexoBlog
func (b *HexoBlog) ReadBlog() error {
	// get the filename with suffix
	fileNameWithSuffix := path.Base(b.InPath)
	b.Slug = fileNameWithSuffix[:len(fileNameWithSuffix)-len(path.Ext(fileNameWithSuffix))]
	//fmt.Printf("b %+v\n", b)
	fileContentBs, err := os.ReadFile(b.InPath)
	if err != nil {
		return err
	}

	fileContent := string(fileContentBs)
	//fmt.Printf("fileContent %v\n", fileContent)
	sep := "---"
	headerIdx := strings.Index(fileContent, sep)
	if headerIdx == -1 {
		return ErrInvalidBlogFormat
	}
	lastHeaderIdx := strings.LastIndex(fileContent, sep)
	if lastHeaderIdx == -1 {
		return ErrInvalidBlogFormat
	}

	start := headerIdx + len(sep)
	header := fileContent[start:lastHeaderIdx]

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer([]byte(header))); err != nil {
		return err
	}
	//fmt.Println(viper.AllSettings())
	b.Title = viper.Get("title").(string)
	b.Date = viper.Get("date").(time.Time).Format("2006-01-02 15:04:05")
	var tags []interface{}
	tags = viper.Get("tags").([]interface{})
	b.Tags = make([]string, len(tags))
	for i, tag := range tags {
		b.Tags[i] = tag.(string)
	}

	//hds := strings.Split(header, "\n")
	//var title, date string
	//for _, hd := range hds {
	//	hd = strings.TrimSpace(hd)
	//
	//	switch {
	//	case strings.Contains(hd, "title"):
	//		title = strings.Split(hd, ": ")[1]
	//	case strings.Contains(hd, "date"):
	//		date = strings.Split(hd, ": ")[1]
	//	}
	//}
	//b.Title = title
	//b.Date = date

	start = lastHeaderIdx + len(sep)
	b.Content = fileContent[start:]
	//fmt.Printf("b %+v\n", b) // for debug
	return nil
}

// WriteBlog write blog to file from HexoBlog
func (b *HexoBlog) WriteBlog(dryrun bool) error {
	return nil
}
