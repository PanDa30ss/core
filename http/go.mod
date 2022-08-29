module http

go 1.16

require core/service v0.0.0

require core/logManager v0.0.0

replace core/service => ../service

replace core/logManager => ../logManager
