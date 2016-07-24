package godbc

import (
	"testing"
	"fmt"
)

func TestWhere_new(t *testing.T) {
	where := Where{"Id", 1, EQUAL}
	actual, _ := where.toString()
	expect := `WHERE Id = "1"`
	if actual != expect {
		t.Fatalf("fatal.\nactual: %s, expect: %s", actual, expect)
	}
}

func TestWhere_toString(t *testing.T) {
	where := Where{nil, 1, EQUAL}
	if actual, err := where.toString(); err == nil {
		t.Fatalf("error should not be empty. %s", actual)
	}
	where = Where{"Id", nil, EQUAL}
	if actual, err := where.toString(); err == nil {
		t.Fatalf("error should not be empty. %s", actual)
	}
	where = Where{nil, nil, EQUAL}
	if actual, _ := where.toString(); actual != "" {
		t.Fatalf("fatal. %s", actual)
	}

	expect := `WHERE Name LIKE "%name%"`
	where = Where{"Name", "name", LIKE}
	if actual, _ := where.toString(); actual != expect {
		t.Fatalf("fatal.\nactual: %s, expect: %s", actual, expect)
	}
}

func TestWhere_Operator(t *testing.T) {
	equal(t, fmt.Sprint(EQUAL), "=")
	equal(t, fmt.Sprint(NOT_EQUAL), "<>")
	equal(t, fmt.Sprint(LIKE), "LIKE")
}

func equal(t *testing.T, actual interface{}, expect interface{}) {
	if actual != expect {
		t.Fatalf("fatal.\nactual: %s, expect: %s", actual, expect)
	}
}