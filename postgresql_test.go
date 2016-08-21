package godbc

import (
	"testing"
	_ "github.com/lib/pq"
	"reflect"
)

type Hoge struct {
	Id int
	Name string
	Flg bool
}

func TestPostgresql(t *testing.T) {
	manager, err := Connect("postgres", "dbname=test host=localhost user=postgres sslmode=disable");
	if err != nil {
		t.Fatal(err)
	}
	manager.Drop(Hoge{}).Execute()
	manager.Create(Hoge{}).Execute()

	err = manager.Insert(Hoge{1, "name_1", true}).Execute()
	if err != nil {
		t.Fatal(err)
	}
	manager.Insert(Hoge{2, "name_2", false}).Execute()

	actual, err := manager.From(&Hoge{}).List()
	if err != nil {
		t.Fatal(err)
	}
	expect := [...]Hoge{{1, "name_1", true}, {2, "name_2", false}}
	if len(actual) != 2 {
		t.Fatalf("値が一致しません.\nactual: %s\nexpect: %s", len(actual), 2)
	}
	if !reflect.DeepEqual(actual[0], expect[0]) {
		t.Fatalf("値が一致しません.\nactual: %s\nexpect: %s", actual[0], expect[0])
	}
	if !reflect.DeepEqual(actual[1], expect[1]) {
		t.Fatalf("値が一致しません.\nactual: %s\nexpect: %s", actual[1], expect[1])
	}
}