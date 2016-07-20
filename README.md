# GodbcManager
構造体でデータベース操作する  
  
## SYNOPSIS
```go
import (
    "godbc"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
)

type Hoge struct {
    Id int
    Name string
    Flg bool
}

func main() {
    manager, _ = Connection("sqlite3", "./test.db")
    manager.Create(Hoge{})
    manager.Insert(Hoge{1, "name1", true})
    manager.Insert(Hoge{2, "name2", false})
    
    rows, _ = manager.FindAll(&Hoge{})
    for i := range rows {
        fmt.Println(row[i].(Hoge))
    }
    
    manager.Drop(Hoge{})
}
```