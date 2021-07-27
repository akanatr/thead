[![build](https://github.com/akanatr/thead/actions/workflows/build.yml/badge.svg)](https://github.com/akanatr/thead/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/akanatr/thead/badge.svg?branch=main)](https://coveralls.io/github/akanatr/thead?branch=main)
[![codebeat badge](https://codebeat.co/badges/dd29b1a2-7c3c-4c2a-b08e-87e152043f57)](https://codebeat.co/projects/github-com-akanatr-thead-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/akanatr/thead)](https://goreportcard.com/report/github.com/akanatr/thead)
[![GitHub license](https://img.shields.io/github/license/akanatr/thead)](https://github.com/akanatr/thead/blob/main/LICENSE)
[![Dockerfile](https://img.shields.io/badge/Docker-ghcr.io%2Fakanatr%2Fthead%3A1.0.0-green?logo=docker)](https://hub.docker.com/repository/docker/akanatr/thead_image)
[![DOI](https://zenodo.org/badge/370276517.svg)](https://zenodo.org/badge/latestdoi/370276517)

# thead
headコマンドの機能拡張 

<img src="./img/thead_icon.svg" width="100">

## Description
既存のheadコマンドは複数のファイル指定し，各ファイルごとの先頭10行分の内容を出力することが可能である．
しかし，特定のディレクトリ直下のファイル全てを指定する際には，正規表現を工夫する必要があり面倒である．
そこでtheadではheadコマンドの機能の一部に加え，ディレクトリを指定するだけでその直下のファイルをtheadコマンドの対象とする機能を追加する．
ディレクトリ内のディレクトリにも再帰的に処理が行われる．


## Usage
```
 thead [OPTIONS...] <FILEs...|DIRs...>  
 OPTIONS:  
    -l, --lines <LINES>             指定された行数を各ファイルごとに出力する．
    -q, --quiet                     複数ファイルを出力する際の各ファイル名を表示しない．
    -h, --help                      このメッセージを出力する.  
    
 ARGUMENTS   
    FILEs...                 カウント対象を指定する．
    DIRs...                  指定したディレクトリ内のファイルを入力ファイルとする．
```

# Install

## 🍺Homebrew

```sh
$ brew tap akanatr/thead
$ brew install thead
```

## 💪Compiling Yourself
```sh
$ git clone https://github.com/akanatr/thead.git
$ cd thead
$ make 
```
  
## icon
<img src="./img/thead_icon.svg" width="100">  

アイコンの画像は [freesvg.org](https://freesvg.org/rejons-head-vector)から取得.
