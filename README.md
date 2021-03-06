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
	_ "github.com/lib/pq"
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
	manager, _ := Connect("postgres", "dbname=test host=localhost user=postgres")
	// テーブルの作成
	manager.Create(Hoge{}).Execute()
	// データの挿入
	manager.Insert(Hoge{1, "name1", true}).Execute()
	manager.Insert(Hoge{2, "name2", false}).Execute()
	// データの取得(リスト)
	manager.From(&Hoge{}).List()
	manager.From(&Hoge{}).Where(Where{"name", "name", LIKE}).List()
	// データの取得(一意)
	manager.From(&Hoge{}).Where(Where{"Id", 1, EQUAL}).SingleResult()
	// データの削除
	manager.From(&Hoge{}).Where(Where{"Id", 1, EQUAL}).Delete().Execute()
	// テーブルの削除
	manager.Drop(Hoge{}).Execute()
	// SQLの取得
	manager.Create(Hoge{}).GetSQL()
	manager.Insert(Hoge{1, "name1", true}).GetSQL()
	manager.From(&Hoge{}).Where(Where{"Id", 1, EQUAL}).Delete().GetSQL()
	manager.Drop(Hoge{}).GetSQL()
}
```

## License
MIT License
