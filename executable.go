package godbc

import "fmt"

type executable struct {
	manager *GodbcManager
	sql string
	err error
}

func (executable *executable) Execute() (error) {
	if executable.err != nil {
		return executable.err
	}
	_, err := executable.manager.db.Exec(
		fmt.Sprintf(executable.sql),
	)
	return err
}

func (executable *executable) GetSQL() (string, error) {
	return executable.sql, executable.err
}