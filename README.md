# GodbcManager
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

type Hoge struct {
	Id int
	Name string
	Flg bool
}

func main() {
	manager, _ := godbc.Connection("sqlite3", "./test.db")
	manager.Create(Hoge{})
	manager.Insert(Hoge{1, "name1", true})
	manager.Insert(Hoge{2, "name2", false})

	rows, _ := manager.FindAll(&Hoge{})
	for i := range rows {
		fmt.Println(rows[i].(Hoge))
	}

	manager.Drop(Hoge{})
}
```

## License
MIT License