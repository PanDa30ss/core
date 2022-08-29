package tcp

const NETWORK_BUF_SIZE int = 1024 * 16
const BLOCK_SIZE_RATE int = 10
const MAX_BLOCKS int = 5
const BLOCK_SIZE = BLOCK_SIZE_RATE * NETWORK_BUF_SIZE
const SEND_MAX_SIZE = MAX_BLOCKS * BLOCK_SIZE
const INVALID_CONV int = ^0
const MSG_HEADER_SIZE int = 4
const CMD_SIZE int = 0x00010000

const (
	ERROR_OK              = 0
	ERROR_DEFAULT         = -1
	ERROR_REMOTECLOSED    = -2
	ERROR_PINGTIMEOUT     = -3
	ERROR_SENDBUFFER      = -4
	ERROR_PROTOCOLVERSION = -5
	ERROR_UNPACK          = -6
	ERROR_PARSEMESSAGE    = -7
	ERROR_SENDMESSAGE     = -8
	ERROR_RECVFAILED      = -9
	ERROR_SENDFAILED      = -10
)

const (
	WaitConnect = iota // value --> 0
	Connecting         // value --> 1
	Fail               // value --> 2
	Connected          // value --> 3
	Open
)

const PING_REQ uint16 = 0xffff
const PING_ACK uint16 = 0xfffe

var DefaultPingTimeData *[]int = &[]int{500, 500, 500}
