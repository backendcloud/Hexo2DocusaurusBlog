# Hexo2DocusaurusBlog
 Golang transfers hexo blog to Docusaurus blog

English | [简体中文](./README_zh.md)


# Convert Hexo Blog to Docusaurus Blog

## 1. Introduction

This tool is used to convert Hexo blog to Docusaurus blog.

## 2. Usage

```bash
Hexo2DocusaurusBlog.exe -h
  -author string
        blog author name
  -docusaurus string
        docusaurus blog directory
  -dry-run
        perform a trial run with no changes made
  -hexo string
        hexo blog directory
```

## 3. Example

### 3.1. Input

from hexo blog markdown file: input/hexo-blog-for-test.md:

```bash
Hexo2DocusaurusBlog.exe -hexo ./input -docusaurus ./output -author backendcloud -dry-run
INFO: Convert Hexo Blog [hexo-blog-for-test] to Docusaurus Blogs Success!
INFO: [DRY-RUN] Write Docusaurus Blog [hexo-blog-for-test] Success!
```

> -dry-run : perform a trial run with no changes made


### 3.2. Output


```bash
Hexo2DocusaurusBlog.exe -hexo ./input -docusaurus ./output -author backendcloud
INFO: Convert Hexo Blog [hexo-blog-for-test] to Docusaurus Blogs Success!
INFO: Write Docusaurus Blog [hexo-blog-for-test] Success!
```

write to Docusaurus Blog markdown output file: output/2023-04-17-hexo-blog-for-test.md