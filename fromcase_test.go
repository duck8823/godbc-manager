package godbc

import (
	"testing"
	"reflect"
)

func TestGodbcManager_List(t *testing.T) {
	manager.Drop(Test{}).Execute()
	manager.Create(Test{}).Execute()

	manager.Insert(Test{1, "name_1", true}).Execute()
	manager.Insert(Test{2, "name_2", false}).Execute()

	actual, err := manager.From(&Test{}).List()
	if err != nil {
		t.Fatal(err)
	}
	expect := []Test{{1, "name_1", true}, {2, "name_2", false}}

	var errs []interface{}
	for i := range actual {
		if !reflect.DeepEqual(actual[i], expect[i]) {
			errs = append(errs, "データが一致しません.", actual[i], expect[i])
		}
		i++
	}
	if len(errs) > 0 {
		t.Fatal(errs)
	}

	if actual, err := manager.From(&Test{}).Where(Where{"Id", []byte{}, EQUAL}).List(); err == nil {
		t.Fatalf("error should not be nil.: %s", actual)
	}

	type NotExist struct {
		Id int
	}
	if actual, err := manager.From(&NotExist{}).List(); err == nil {
		t.Fatalf("error should not be nil.: %s", actual)
	}
}

func TestGodbcManager_SingleResult(t *testing.T) {
	manager.Drop(Test{}).Execute()
	manager.Create(Test{}).Execute()

	manager.Insert(Test{1, "name_1", true}).Execute()
	manager.Insert(Test{2, "name_2", false}).Execute()

	actual, _ := manager.From(&Test{}).Where(Where{"Id", 1, EQUAL}).SingleResult()
	expect := Test{1, "name_1", true}

	if !reflect.DeepEqual(actual, expect) {
		t.Fatalf("データが一致しません.\nactual: %s, expect: %s", actual, expect)
	}

	if actual, err := manager.From(&Test{}).Where(Where{"Id", []byte{}, EQUAL}).SingleResult(); err == nil {
		t.Fatalf("error should not be nil.\nactual: %s", actual)
	}

	if actual, err := manager.From(&Test{}).SingleResult(); err == nil {
		t.Fatalf("error should not be nil.\nactual: %s", actual)
	}
}

func TestGodbcManager_Delete(t *testing.T) {
	manager.Drop(Test{}).Execute()
	manager.Create(Test{}).Execute()

	manager.Insert(Test{1, "name_1", true}).Execute()
	manager.Insert(Test{2, "name_2", false}).Execute()

	manager.From(&Test{}).Where(Where{"Id", 1, EQUAL}).Delete().Execute()
	actual, _ := manager.From(&Test{}).SingleResult()
	expect := Test{2, "name_2", false}

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("データが一致しません.\nactual: %s, expect: %s", actual, expect)
	}
}