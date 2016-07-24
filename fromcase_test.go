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
}

func TestGodbcManager_FindSingleResult(t *testing.T) {
	manager.Drop(Test{}).Execute()
	manager.Create(Test{}).Execute()

	manager.Insert(Test{1, "name_1", true}).Execute()
	manager.Insert(Test{2, "name_2", false}).Execute()

	actual, _ := manager.From(&Test{}).Where(Where{"Id", 1, EQUAL}).SingleResult()
	expect := Test{1, "name_1", true}

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("データが一致しません.\nactual: %s, expect: %s", actual, expect)
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