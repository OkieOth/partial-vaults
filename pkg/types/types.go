package types

import "fmt"

type ValueType int

const (
	OBJECT ValueType = iota
	ARRAY
	STRING
	BOOL
	NUMBER
	INTEGER
	NULL
)

func (t ValueType) String() string {
	switch t {
	case OBJECT:
		return "Object"
	case ARRAY:
		return "Array"
	case STRING:
		return "String"
	case BOOL:
		return "Bool"
	case NUMBER:
		return "Number"
	case INTEGER:
		return "Integer"
	case NULL:
		return "Null"
	default:
		return fmt.Sprintf("UnknownType: %d", t)
	}
}

type ProcessFunc func([]byte, ValueType, string) (any, ValueType, error)
