package tcp

type CommandType = string

const (
	CMD_REQUEST  CommandType = "Request"
	CMD_SOLUTION CommandType = "Solution"
)

type Header = string

const (
	COMMAND       Header = "X-Command:"
	SOLUTION      Header = "X-Solution:"
	CHALLENGE     Header = "X-Challenge:"
	ACCESS        Header = "X-Access:"
	TOKEN         Header = "X-Token:"
	QUOTE         Header = "X-Quote:"
	REQUEST_ID    Header = "X-Request-id:"
	REQUEST_TIME  Header = "X-Request-time:"
	MESSAGE_START Header = "START:"
	MESSAGE_END   Header = "END:"
)

const (
	MESSAGE_SIZE_LIMIT = 4096 * 4
)
