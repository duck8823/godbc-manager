package godbc

import (
	"fmt"
	"database/sql"
)

type executable struct {
	db sql.DB
	sql string
	err error
}

func (executable *executable) Execute() (error) {
	if executable.err != nil {
		return executable.err
	}
	_, err := executable.db.Exec(
		fmt.Sprintf(executable.sql),
	)
	return err
}

func (executable *executable) GetSQL() (string, error) {
	return executable.sql, executable.err
}