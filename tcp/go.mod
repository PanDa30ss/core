module tcp

go 1.16

require core/logManager v0.0.0

require core/message v0.0.0

require core/service v0.0.0

replace core/logManager => ../logManager

replace core/message => ../message

replace core/service => ../service
