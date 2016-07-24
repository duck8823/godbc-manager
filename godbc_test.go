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
	manager.Drop(Test{}).Execute()
	manager.Create(Test{}).Execute()

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
	manager.Drop(Test{}).Execute()
	manager.Create(Test{}).Execute()

	manager.Drop(Test{}).Execute()
	_ , err := manager.db.Query(
		`SELECT * FROM test`,
	)
	if err == nil {
		t.Fatal("カラムが存在します.")
	}
}

func TestGodbcManager_Insert(t *testing.T) {
	manager.Drop(Test{}).Execute()
	manager.Create(Test{}).Execute()

	manager.Insert(Test{1, "name_1", true}).Execute()
	manager.Insert(Test{2, "name_2", false}).Execute()

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
		test := Test{}
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

func TestGodbcManager_toString(t *testing.T) {
	if actual, _ := toString(1); actual != "1" {
		t.Fatalf("fatal.\nactual: %s", actual)
	}
	if actual, _ := toString("string"); actual != "string" {
		t.Fatalf("fatal.\nactual: %s", actual)
	}
	if actual, _ := toString(true); actual != "true" {
		t.Fatalf("fatal.\nactual: %s", actual)
	}
	if actual, err := toString([]byte{}); err == nil {
		t.Fatalf("fatal.\nactual: %s", actual)
	}
}

func TestGodbcManager_createSentence(t *testing.T) {
	type Success struct {
		Id int
		Name string
		Flg bool
	}
	expect := ""
	if actual, _ := createSentence(Success{}); actual == expect {
		t.Fatalf("fail.\nactual: %s, expect: %s", actual, expect)
	}

	type Fail struct {
		FailField []byte
	}
	if _, err := createSentence(Fail{}); err == nil {
		t.Fatalf("should have error.")
	}
}

func TestGodbcManager_Readme(t *testing.T) {
	// 構造体の定義
	type Hoge struct {
		Id int
		Name string
		Flg bool
	}
	// データベースへの接続
	manager, _ := Connection("sqlite3", "./test.db")
	// テーブルの作成
	manager.Create(Hoge{}).Execute()
	// データの挿入
	manager.Insert(Hoge{1, "name1", true}).Execute()
	manager.Insert(Hoge{2, "name2", false}).Execute()
	// データの取得(リスト)
	manager.From(&Hoge{}).List()
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