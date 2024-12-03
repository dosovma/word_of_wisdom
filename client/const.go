package app

const (
	REQUEST  = "X-Request-quote"
	TASK     = "X-Task"
	SOLUTION = "X-Solution"
	DENIED   = "X-Denied"
	QUOTE    = "X-Quote"
)

type command string

func (c command) String() string {
	return string(c)
}

var quote command = "quote"
