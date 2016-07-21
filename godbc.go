package godbc

import (
	"database/sql"
	"strings"
	"strconv"
	"reflect"
	"errors"
)

type GodbcManager struct {
	db sql.DB
}

func Connection(driverName string, dataSourceName string) (GodbcManager, error) {
	db, err := sql.Open(driverName, dataSourceName)
	return GodbcManager{db: *db}, err
}

func (manager GodbcManager) Drop(entity interface{}) (error) {
	name := reflect.TypeOf(entity).Name()
	_, err := manager.db.Exec(
		`DROP TABLE IF EXISTS ` + name,
	)
	return err
}

func (manager GodbcManager) Create(entity interface{}) (error) {
	t := reflect.TypeOf(entity)
	v := reflect.ValueOf(entity)
	schema := make([]string, t.NumField());
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i).Interface()
		switch fv.(type) {
		default:
			return errors.New("次の型は対応していません. :" + reflect.ValueOf(fv).Type().Name())
		case string:
			schema[i] = "\"" + f.Name + "\" TEXT"
		case int:
			schema[i] = "\"" + f.Name + "\" INTEGER"
		case bool:
			schema[i] = "\"" + f.Name + "\" BOOLEAN"
		}
	}

	_, err := manager.db.Exec(
		`CREATE TABLE ` + t.Name() + ` (` + strings.Join(schema, ",") + `)`,
	)
	return err;
}

func (manager GodbcManager) Insert(data interface{}) (error) {
	names, values := toStrings(data)
	_, err := manager.db.Exec(
		`INSERT INTO ` + reflect.TypeOf(data).Name() + ` (` + names + `) VALUES (` + values + `)`,
	)
	return err
}

func toStrings(data interface{}) (string, string) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	keys := make([]string, t.NumField())
	values := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i).Interface()
		str, _ := toString(fv)
		keys[i] = f.Name
		values[i] = "\"" + str + "\""
	}
	return strings.Join(keys, ","), strings.Join(values, ",")
}

func (manager GodbcManager) FindAll(entity interface{}) ([]interface{}, error) {
	result := []interface{}{}

	t := reflect.TypeOf(entity).Elem()
	v := reflect.ValueOf(entity).Elem()
	cols := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		cols[i] = t.Field(i).Name
	}

	rows, err := manager.db.Query(
		`SELECT ` + strings.Join(cols, ",")  + ` FROM ` + t.Name(),
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
			f := v.Field(i)
			value := values[i]
			switch v.Field(i).Kind() {
			case reflect.Int:
				value = int(value.(int64))
			case reflect.String:
				value = string(value.([]byte))
			case reflect.Bool:
				value, _ = strconv.ParseBool(string(value.([]byte)))
			}
			f.Set(reflect.ValueOf(value))
		}
		result = append(result, reflect.ValueOf(entity).Elem().Interface())
	}
	return result, nil;
}

func toString(args ...interface{}) (string, error) {
	result := ""
	for _, arg := range args {
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
	}
	return result, nil
}