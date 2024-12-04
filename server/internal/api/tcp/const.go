package tcp

type CommandType = string

const (
	CmdToken    CommandType = "Token"
	CmdSolution CommandType = "Solution"
	CmdQuote    CommandType = "Quote"
)

type Header = string

const (
	Command      Header = "X-Command:"
	Solution     Header = "X-Solution:"
	Challenge    Header = "X-Challenge:"
	Access       Header = "X-Access:"
	Token        Header = "X-Token:"
	Quote        Header = "X-Quote:"
	RequestID    Header = "X-Request-id:"
	RequestTime  Header = "X-Request-time:"
	MessageStart Header = "START:"
	MessageEnd   Header = "END:"
)

const (
	MessageSizeLimit = 4096 * 4
)
