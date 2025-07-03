package types

type ValueType int

const (
	OBJECT ValueType = iota
	ARRAY
	STRING
	BOOL
	NUMBER
	NULL
)

type ProcessFunc func([]byte, ValueType, string) (any, ValueType, error)
