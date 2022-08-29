module mysql

go 1.16

require github.com/go-sql-driver/mysql v1.6.0

require core/logManager v0.0.0

require core/service v0.0.0

replace core/logManager => ../../logManager

replace core/service => ../../service
