package godbc

import (
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
)

var manager GodbcManager
type Test struct {
	Id int
	Name string
	Flg bool
}

func TestConnection(t *testing.T) {
	manager, _ = Connection("sqlite3", "./test.db")
	var _manger interface{} = manager
	_, success := _manger.(GodbcManager)
	if !success {
		t.Fatal("型が違います.", reflect.TypeOf(manager))
	}
}

func TestGodbcManager_Create(t *testing.T) {
	manager.Drop(Test{})
	manager.Create(Test{})

	rows, _ := manager.db.Query(
		`SELECT * FROM test`,
	)

	expect := []string{"Id", "Name", "Flg"}
	actual, _ := rows.Columns()
	if !reflect.DeepEqual(expect, actual) {
		t.Fatal("カラムが一致しません.", actual, expect)
	}
}

func TestGodbcManager_Drop(t *testing.T) {
	manager.Drop(Test{})
	manager.Create(Test{})

	manager.Drop(Test{})
	_ , err := manager.db.Query(
		`SELECT * FROM test`,
	)
	if err == nil {
		t.Fatal("カラムが存在します.")
	}
}

func TestGodbcManager_Insert(t *testing.T) {
	manager.Drop(Test{})
	manager.Create(Test{})

	manager.Insert(Test{1, "name_1", true})
	manager.Insert(Test{2, "name_2", false})

	expect := []Test{{1, "name_1", true}, {2, "name_2", false}}

	row, _ := manager.db.Query(
		`SELECT count(*) FROM test`,
	)
	row.Next()

	var count int
	row.Scan(&count)

	if count != len(expect) {
		t.Fatal("数が一致しません.", count, len(expect))
	}
	row.Close()

	actual, _ := manager.db.Query(
		`SELECT * FROM test`,
	)

	var err []interface{}
	i := 0
	for actual.Next() {
		test := Test {}
		actual.Scan(&test.Id, &test.Name, &test.Flg)
		if !reflect.DeepEqual(test, expect[i]) {
			err = append(err, "データが一致しません.", test, expect[i])
		}
		i++
	}
	if len(err) > 0 {
		t.Fatal(err)
	}
}

func TestGodbcManager_FindAll(t *testing.T) {
	manager.Drop(Test{})
	manager.Create(Test{})

	manager.Insert(Test{1, "name_1", true})
	manager.Insert(Test{2, "name_2", false})

	actual, _ := manager.FindAll(&Test{})
	expect := []Test{{1, "name_1", true}, {2, "name_2", false}}

	var err []interface{}
	for i := range actual {
		if !reflect.DeepEqual(actual[i], expect[i]) {
			err = append(err, "データが一致しません.", actual[i], expect[i])
		}
		i++
	}
	if len(err) > 0 {
		t.Fatal(err)
	}
}