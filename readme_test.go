package godbc

import (
	"testing"
	_ "github.com/lib/pq"
)

func TestGodbcManager_Readme(t *testing.T) {
	// 構造体の定義
	type Hoge struct {
		Id int
		Name string
		Flg bool
	}
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