package tcp

type CommandType = string

const (
	CMD_REQUEST  CommandType = "Request"
	CMD_SOLUTION CommandType = "Solution"
)

type Header = string

const (
	COMMAND      Header = "X-Command:"
	SOLUTION     Header = "X-Solution:"
	CHALLENGE    Header = "X-Challenge:"
	ACCESS       Header = "X-Access:"
	TOKEN        Header = "X-Token:"
	QUOTE        Header = "X-Quote:"
	REQUEST_ID   Header = "X-Request-id:"
	REQUEST_TIME Header = "X-Request-time:"

	MESSAGE_SIZE_LIMIT = 4096 * 4
	MESSAGE_START      = "START"
	MESSAGE_END        = "END"
)

func isPayload(msg string) bool {
	return msg != MESSAGE_START && msg != MESSAGE_END
}
