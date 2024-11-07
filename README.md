# go-file-xor

[![Typing SVG](https://readme-typing-svg.demolab.com?font=Fira+Code&pause=1000&width=435&lines=Go+File+XOR+%F0%9F%91%86%F0%9F%A4%93)](https://github.com/liuyuhe666/go-file-xor)

![GitHub watchers](https://img.shields.io/github/watchers/liuyuhe666/go-file-xor?style=social) ![GitHub stars](https://img.shields.io/github/stars/liuyuhe666/go-file-xor?style=social) ![GitHub forks](https://img.shields.io/github/forks/liuyuhe666/go-file-xor?style=social)

![GitHub top language](https://img.shields.io/github/languages/top/liuyuhe666/go-file-xor?style=flat-square)  ![GitHub](https://img.shields.io/github/license/liuyuhe666/go-file-xor?style=flat-square) ![GitHub open issues](https://img.shields.io/github/issues/liuyuhe666/go-file-xor?style=flat-square) ![GitHub closed issues](https://img.shields.io/github/issues-closed/liuyuhe666/go-file-xor) ![GitHub last commit](https://img.shields.io/github/last-commit/liuyuhe666/go-file-xor?style=flat-square) ![GitHub repo size](https://img.shields.io/github/repo-size/liuyuhe666/go-file-xor?style=flat-square)

[报告 Bug](https://github.com/liuyuhe666/go-file-xor/issues) · [提出新特性](https://github.com/liuyuhe666/go-file-xor/pulls)

## 使用说明

1. 设置环境变量 `XOR_SECRET`

PowerShell:

```ps1
$env:XOR_SECRET="YOUR_XOR_SECRET"
```
Bash:

```bash
export XOR_SECRET="YOUR_XOR_SECRET"
```

2. 下载 `go-file-xor`

如果已经配置好 `Go` 语言开发环境，可以直接使用 `go install` 命令来安装

```bash
go install github.com/liuyuhe666/go-file-xor@latest
```
如果没有配置 `Go` 语言开发环境，可以在 [`Release`](https://github.com/liuyuhe666/go-file-xor/releases) 页面下载对应操作系统的可执行文件

3. 在命令行中运行 `go-file-xor`

```bash
go file xor

Usage:
  go-file-xor <filename> [flags]

Flags:
  -h, --help   help for go-file-xor
```

XOR 加密:

```bash
go-file-xor test.txt
```

XOR 解密:

```bash
go-file-xor test.txt.xor
```

## 作者

- [GitHub@liuyuhe666](https://github.com/liuyuhe666)

## 版权说明

`GPL-3.0 license` ，详情请参阅 [LICENSE](./LICENSE)

## 鸣谢

- [https://github.com/spf13/cobra](https://github.com/spf13/cobra)