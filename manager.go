package godbc

import (
	"database/sql"
	"strings"
	"strconv"
	"reflect"
	"errors"
	"fmt"
)

type GodbcManager struct {
	db sql.DB
}

func Connect(driverName string, dataSourceName string) (GodbcManager, error) {
	db, err := sql.Open(driverName, dataSourceName)
	return GodbcManager{db: *db}, err
}

func (manager *GodbcManager) From(entity interface{}) (*fromCase) {
	return &fromCase{manager, entity, Where{}}
}

func (manager *GodbcManager) Drop(entity interface{}) (*executable) {
	name := reflect.TypeOf(entity).Name()
	return &executable{manager, fmt.Sprintf(`DROP TABLE IF EXISTS %s`, name), nil}
}

func (manager *GodbcManager) Create(entity interface{}) (*executable) {
	t := reflect.TypeOf(entity)
	v := reflect.ValueOf(entity)
	schema := make([]string, t.NumField());
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i).Interface()
		switch fv.(type) {
		default:
			return &executable{err: errors.New("次の型は対応していません. :" + reflect.ValueOf(fv).Type().Name())}
		case string:
			schema[i] = f.Name + " TEXT"
		case int:
			schema[i] = f.Name + " INTEGER"
		case bool:
			schema[i] = f.Name + " BOOLEAN"
		}
	}
	return &executable{manager, fmt.Sprintf(`CREATE TABLE %s (%s)`, t.Name(), strings.Join(schema, ",")), nil}
}

func (manager *GodbcManager) Insert(data interface{}) (*executable) {
	sentence, err := createSentence(data)
	return &executable{manager, fmt.Sprintf(`INSERT INTO %s %s`, reflect.TypeOf(data).Name(), sentence), err}
}

func createSentence(data interface{}) (string, error) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	keys := make([]string, t.NumField())
	values := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i).Interface()
		str, err := toString(fv)
		if err != nil {
			return str, err
		}
		keys[i] = f.Name
		values[i] = "'" + str + "'"
	}
	return fmt.Sprintf(`(%s) VALUES (%s)`, strings.Join(keys, ","), strings.Join(values, ",")), nil
}

func toString(arg interface{}) (string, error) {
	result := ""
	switch val := arg.(type) {
	default:
		return "", errors.New("次の型は対応していません. " + reflect.ValueOf(val).Type().Name())
	case int:
		result += strconv.Itoa(val)
	case string:
		result += val
	case bool:
		result += strconv.FormatBool(val)
	}
	return result, nil
}