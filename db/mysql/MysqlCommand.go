package mysql

type MysqlCallBack func(result *MysqlResult, params ...interface{})
type mysqlCommand struct {
	callback MysqlCallBack
	result   *MysqlResult
	params   []interface{}
}

func (this *mysqlCommand) Execute() {
	if this.callback == nil {
		return
	}
	this.callback(this.result, this.params...)
	this.result.Result.Close()
}
