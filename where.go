package godbc

import (
	"fmt"
	"errors"
)

type Operator string

const (
	EQUAL Operator = "="
	NOT_EQUAL Operator = "<>"
	LIKE Operator = "LIKE"
)

type Where struct {
	column interface{}
	value  interface{}
	operator Operator
}

func (where *Where) toString() (string, error) {
	if where.column == nil && where.value == nil {
		return "", nil
	} else if (where.column == nil && where.value != nil) || (where.column != nil && where.value == nil) {
		return "", errors.New(fmt.Sprintf("error: %s", where))
	}
	value, err := toString(where.value)
	if err != nil {
		return "", err
	}
	switch where.operator {
	case LIKE:
		value = "%" + value + "%"
	}
	return fmt.Sprintf(`WHERE %s %s "%s"`, where.column, where.operator, value), nil
}