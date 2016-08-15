package godbc

import "testing"

func TestExecutable_init(t *testing.T) {
	manager, _ = Connection("sqlite3", "./test.db")
}

func TestExecutable_Execute(t *testing.T) {
	type Fail struct {
		Id int
		Fail []byte
	}
	err := manager.Create(Fail{}).Execute()
	if err == nil {
		t.Fatal("error should not be nil.")
	}

	type Success struct {
		Id int
		Name string
	}
	_   = manager.Drop(Success{}).Execute()
	err = manager.Create(Success{}).Execute()
	if err != nil {
		t.Fatal(err)
	}
	err = manager.Create(Success{}).Execute()
	if err == nil {
		t.Fatal("error should not be nil.")
	}
}

func TestExecutable_GetSQL(t *testing.T) {
	type Hoge struct {
		Id int
		Name string
	}
	actual, _ := manager.Create(Hoge{}).GetSQL()
	expect := `CREATE TABLE Hoge ('Id' INTEGER,'Name' TEXT)`
	if actual != expect {
		t.Fatalf("fatal.\nactual: %s, expect:%s", actual, expect)
	}
}
