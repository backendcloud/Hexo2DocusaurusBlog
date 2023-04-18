
# Hexo2DocusaurusBlog

## 1. 简介

此工具用于将 Hexo 博客转换为 Docusaurus 博客。

## 2. 用法

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

## 3. 例子

### 3.1. Input

Hexo 博客的markdown文件位于: input/hexo-blog-for-test.md:

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

执行工具，转换后的 Docusaurus 博客的markdown文件位于: output/2023-04-17-hexo-blog-for-test.md