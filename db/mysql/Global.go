package mysql

func MakeMysqlCommand(callback MysqlCallBack, params ...interface{}) *mysqlCommand {
	ret := &mysqlCommand{}
	ret.callback = callback
	ret.params = params
	return ret
}
