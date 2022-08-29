package mysql

import (
	log "core/logManager"
	"core/service"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jmoiron/sqlx"
	"database/sql"
)

var defaultMaxConn int = 10
var defaultMaxIdle int = 5

type Mysql struct {
	db      *sql.DB
	maxConn int
	maxIdle int
	url     string
	isOpen  bool
}

func (this *Mysql) InitDB() {
	this.url = ""
	this.maxConn = defaultMaxConn
	this.maxIdle = defaultMaxIdle
	this.db = nil
	this.isOpen = false
}

func (this *Mysql) SetUrl(url string) {
	this.url = url
}

func (this *Mysql) SetMaxConn(maxConn int) {
	this.maxConn = maxConn
}

func (this *Mysql) SetMaxIdle(maxIdle int) {
	this.maxIdle = maxIdle
}

func (this *Mysql) Open() bool {
	if this.url == "" {
		return false
	}
	// db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mog")
	db, err := sql.Open("mysql", this.url)
	if err != nil {
		return false
	}
	this.db = db
	this.isOpen = true
	log.Info("mysql open")
	return true
}

func (this *Mysql) Query(cmd *mysqlCommand, query string) {
	if !this.isOpen {
		return
	}
	go func() {
		rows, err := this.db.Query(query)
		cmd.result = &MysqlResult{err, rows}
		service.Post(cmd)

	}()
}
