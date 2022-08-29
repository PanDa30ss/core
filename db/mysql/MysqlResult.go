package mysql

// import "github.com/jmoiron/sqlx"
import (
	"database/sql"
	"reflect"
)

type MysqlResult struct {
	Err    error
	Result *sql.Rows
}

func (result *MysqlResult) ScanRow() map[string]reflect.Value {
	names, _ := result.Result.Columns()
	types, _ := result.Result.ColumnTypes()
	interfaces := make([]interface{}, len(types))
	instances := make([]interface{}, len(types))
	for index, v := range types {
		c := reflect.New(v.ScanType())
		instances[index] = c.Elem()
		interfaces[index] = c.Interface()
	}
	result.Result.Scan(interfaces...)
	ret := make(map[string]reflect.Value)
	for index, name := range names {
		ret[name] = instances[index].(reflect.Value)
	}
	return ret
}

func (result *MysqlResult) Next() bool {
	return result.Result.Next()
}

func (result *MysqlResult) NextResultSet() bool {
	return result.Result.NextResultSet()
}
