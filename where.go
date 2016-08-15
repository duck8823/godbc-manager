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
	Column   interface{}
	Value    interface{}
	Operator Operator
}

func (where *Where) toString() (string, error) {
	if where.Column == nil && where.Value == nil {
		return "", nil
	} else if (where.Column == nil && where.Value != nil) || (where.Column != nil && where.Value == nil) {
		return "", errors.New(fmt.Sprintf("error: %s", where))
	}
	value, err := toString(where.Value)
	if err != nil {
		return "", err
	}
	switch where.Operator {
	case LIKE:
		value = "%" + value + "%"
	}
	return fmt.Sprintf(`WHERE %s %s '%s'`, where.Column, where.Operator, value), nil
}