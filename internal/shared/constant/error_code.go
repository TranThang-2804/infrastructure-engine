package constant

type ErrorCode int

const (
	NotBelongToAnyContext ErrorCode = iota
	UnexpectedResponseFormat
)

func (s ErrorCode) String() string {
	return [...]string{
		"Not Belong To Any Context",
		"Unexpected Response Format",
	}[s]
}
