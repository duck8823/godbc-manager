# GodbcManager
[![Build Status](https://travis-ci.org/duck8823/godbc-manager.svg?branch=master)](https://travis-ci.org/duck8823/godbc-manager)
[![Coverage Status](http://coveralls.io/repos/github/duck8823/godbc-manager/badge.svg?branch=master)](https://coveralls.io/github/duck8823/godbc-manager?branch=master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)  
  
構造体でデータベース操作する  
  
## INSTALL
```sh
go get github.com/duck8823/godbc-manager
```
  
## SYNOPSIS
```go
package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/duck8823/godbc-manager"
)

// 構造体の定義
type Hoge struct {
	Id int
	Name string
	Flg bool
}

func main() {
	// データベースへの接続
	manager, _ := godbc.Connection("sqlite3", "./test.db")
	// テーブルの作成
	manager.Create(Hoge{}).Execute()
	// データの挿入
	manager.Insert(Hoge{1, "name1", true}).Execute()
	manager.Insert(Hoge{2, "name2", false}).Execute()
	// データの取得(リスト)
	rows, _ := manager.From(&Hoge{}).List()
	for i := range rows {
		fmt.Println(rows[i].(Hoge))
	}
	// データの取得(一意)
	row, _ := manager.From(&Hoge{}).Where(godbc.Where{"Id", 1, godbc.EQUAL}).SingleResult()
	fmt.Println(row.(Hoge))
	// データの削除
	manager.From(&Hoge{}).Where(godbc.Where{"Id", 1, godbc.EQUAL}).Delete().Execute()
	// テーブルの削除
	manager.Drop(Hoge{}).Execute()
	
	// SQLの取得
	createSQL, err := manager.Create(Hoge{}).GetSQL()
	insertSQL, err := manager.Insert(Hoge{1, "name1", true}).GetSQL()
	deleteSQL, err := manager.From(&Hoge{}).Where(godbc.Where{"Id", 1, godbc.EQUAL}).Delete().GetSQL()
	dropSQL,   err := manager.Drop(Hoge{}).GetSQL()
}
```

## License
MIT License
