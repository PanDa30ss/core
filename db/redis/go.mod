module redis

go 1.16

require github.com/garyburd/redigo v1.6.2

require core/logManager v0.0.0

require core/service v0.0.0

replace core/logManager => ../../logManager

replace core/service => ../../service
