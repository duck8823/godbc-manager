package godbc

import (
	"reflect"
	"strconv"
	"strings"
	"errors"
	"fmt"
	"database/sql"
)

type fromCase struct {
	db sql.DB
	from interface{}
	where Where
}

func (fromCase *fromCase) Where(where Where) (*fromCase) {
	fromCase.where = where
	return fromCase
}

func (fromCase *fromCase) List() ([]interface{}, error) {
	result := []interface{}{}

	t := reflect.TypeOf(fromCase.from).Elem()
	v := reflect.ValueOf(fromCase.from).Elem()
	cols := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		cols[i] = t.Field(i).Name
	}

	where, err := fromCase.where.toString()
	if err != nil {
		return nil, err
	}
	rows, err := fromCase.db.Query(
		fmt.Sprintf(`SELECT %s FROM %s %s`, strings.Join(cols, ","), t.Name(), where),
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		values := make([]interface{}, t.NumField())
		valuePointers := make([]interface{}, t.NumField())

		for i := 0; i < t.NumField(); i++ {
			valuePointers[i] = &values[i]
		}

		rows.Scan(valuePointers...)
		for i := 0; i < t.NumField(); i++ {
			field := v.Field(i)
			value := values[i]
			setValue(field, value)
		}
		result = append(result, reflect.ValueOf(fromCase.from).Elem().Interface())
	}
	return result, nil;
}

func (fromCase *fromCase) SingleResult() (interface{}, error) {
	result, err := fromCase.List()
	if err != nil {
		return nil, err
	}
	if len_ := len(result); len_ > 1 {
		return nil, errors.New("結果が一意でありません.")
	}
	return result[0], nil;
}

func (fromCase *fromCase) Delete() (*executable) {
	t := reflect.TypeOf(fromCase.from).Elem()
	where, err := fromCase.where.toString()
	return &executable{fromCase.db, fmt.Sprintf(`DELETE FROM %s %s`, t.Name(), where), err}
}

func setValue(field reflect.Value, value interface{}) {
	switch field.Kind() {
	case reflect.Int:
		value = int(value.(int64))
	case reflect.String:
		if _, success := value.(string); !success {
			value = string(value.([]byte))
		}
	case reflect.Bool:
		if _, success := value.(bool); !success {
			value, _ = strconv.ParseBool(string(value.([]byte)))
		}
	}
	field.Set(reflect.ValueOf(value))
}