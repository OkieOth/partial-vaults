package jsonreader

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/okieoth/pvault/pkg/types"
)

// This code is AI proposed

// OrderedValue is a generic value that can be a primitive, object, or array
type OrderedValue struct {
	Type  types.ValueType
	Value interface{}
}

func traversToValue[T any](o *OrderedValue, errorValue T, fn func(*OrderedValue) (T, error)) (T, error) {
	switch o.Type {
	case types.OBJECT:
		orderedObject, ok := o.Value.(OrderedObject)
		if !ok {
			return errorValue, fmt.Errorf("error while cast object")
		}
		return traversToValue(orderedObject[0].Value, errorValue, fn)
	case types.BOOL, types.INTEGER, types.STRING, types.NUMBER:
		return fn(o)
	default:
		return errorValue, fmt.Errorf("wrong OrderedValue type: %v", o.Type)
	}
}

func (o *OrderedValue) StringValue() (string, error) {
	castFn := func(ov *OrderedValue) (string, error) {
		if ov.Type == types.STRING {
			s, ok := ov.Value.(string)
			if !ok {
				return "", fmt.Errorf("error while cast to string")
			}
			return s, nil
		} else {
			return "", fmt.Errorf("no string value")
		}
	}
	return traversToValue(o, "", castFn)
}

func (o *OrderedValue) IntValue() (int64, error) {
	castFn := func(ov *OrderedValue) (int64, error) {
		switch ov.Type {
		case types.INTEGER, types.NUMBER:
			n, ok := ov.Value.(json.Number)
			if !ok {
				return 0, fmt.Errorf("error while cast to json.Number")
			}
			return n.Int64()
		default:
			return 0, fmt.Errorf("no int value")
		}
	}
	return traversToValue(o, 0, castFn)
}

func (o *OrderedValue) BoolValue() (bool, error) {
	castFn := func(ov *OrderedValue) (bool, error) {
		if ov.Type == types.BOOL {
			n, ok := ov.Value.(bool)
			if !ok {
				return false, fmt.Errorf("error while cast to bool")
			}
			return n, nil
		} else {
			return false, fmt.Errorf("no bool value")
		}
	}
	return traversToValue(o, false, castFn)
}

func (o *OrderedValue) NumberValue() (float64, error) {
	castFn := func(ov *OrderedValue) (float64, error) {
		if ov.Type == types.NUMBER {
			n, ok := ov.Value.(json.Number)
			if !ok {
				return 0, fmt.Errorf("error while cast to json.Number")
			}
			return n.Float64()
		} else {
			return 0, fmt.Errorf("no number value")
		}
	}
	return traversToValue(o, 0.0, castFn)
}

// OrderedPair represents a single key-value pair in a JSON object
type OrderedPair struct {
	Key   string
	Value *OrderedValue
}

// OrderedObject is a JSON object with key order preserved
type OrderedObject []OrderedPair

// OrderedArray is a JSON array that can hold nested OrderedValues
type OrderedArray []*OrderedValue

// UnmarshalJSON for OrderedValue handles all JSON types
func (ov *OrderedValue) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	switch data[0] {
	case '{':
		var obj OrderedObject
		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		ov.Type = types.OBJECT
		ov.Value = obj
	case '[':
		var arr OrderedArray
		if err := json.Unmarshal(data, &arr); err != nil {
			return err
		}
		ov.Type = types.ARRAY
		ov.Value = arr
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		ov.Type = types.STRING
		ov.Value = s
	case 't', 'f':
		var b bool
		if err := json.Unmarshal(data, &b); err != nil {
			return err
		}
		ov.Type = types.BOOL
		ov.Value = b
	case 'n':
		ov.Type = types.NULL
		ov.Value = nil
	default:
		var num json.Number
		if err := json.Unmarshal(data, &num); err != nil {
			return err
		}
		ov.Type = types.NUMBER
		ov.Value = num
	}
	return nil
}

// MarshalJSON for OrderedValue writes it back correctly
func (ov *OrderedValue) MarshalJSON() ([]byte, error) {
	switch ov.Type {
	case types.OBJECT:
		return json.Marshal(ov.Value)
	case types.ARRAY:
		return json.Marshal(ov.Value)
	case types.STRING:
		return json.Marshal(ov.Value.(string))
	case types.BOOL:
		return json.Marshal(ov.Value.(bool))
	case types.NULL:
		return []byte("null"), nil
	case types.NUMBER:
		return []byte(ov.Value.(json.Number)), nil
	default:
		return nil, fmt.Errorf("unknown type: %s", ov.Type)
	}
}

// UnmarshalJSON for OrderedObject preserves key order
func (o *OrderedObject) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	tok, err := dec.Token()
	if err != nil || tok != json.Delim('{') {
		return fmt.Errorf("expected '{', got %v", tok)
	}

	var obj OrderedObject
	for dec.More() {
		tok, err := dec.Token()
		if err != nil {
			return err
		}
		key := tok.(string)

		var val OrderedValue
		if err := dec.Decode(&val); err != nil {
			return err
		}
		obj = append(obj, OrderedPair{Key: key, Value: &val})
	}

	// Consume '}'
	if tok, err := dec.Token(); err != nil || tok != json.Delim('}') {
		return fmt.Errorf("expected '}', got %v", tok)
	}

	*o = obj
	return nil
}

// MarshalJSON for OrderedObject preserves key order
func (o OrderedObject) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, pair := range o {
		if i > 0 {
			buf.WriteByte(',')
		}
		keyBytes, _ := json.Marshal(pair.Key)
		buf.Write(keyBytes)
		buf.WriteByte(':')
		valBytes, err := json.Marshal(pair.Value)
		if err != nil {
			return nil, err
		}
		buf.Write(valBytes)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON for OrderedArray
func (a *OrderedArray) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	tok, err := dec.Token()
	if err != nil || tok != json.Delim('[') {
		return fmt.Errorf("expected '[', got %v", tok)
	}

	var arr OrderedArray
	for dec.More() {
		var val OrderedValue
		if err := dec.Decode(&val); err != nil {
			return err
		}
		arr = append(arr, &val)
	}

	// Consume ']'
	if tok, err := dec.Token(); err != nil || tok != json.Delim(']') {
		return fmt.Errorf("expected ']', got %v", tok)
	}

	*a = arr
	return nil
}

func PrintOrdered(val *OrderedValue, indent int) {
	prefix := bytes.Repeat([]byte("  "), indent)
	switch val.Type {
	case types.OBJECT:
		fmt.Println(string(prefix) + "{")
		for _, pair := range val.Value.(OrderedObject) {
			fmt.Printf("%s  \"%s\":\n", prefix, pair.Key)
			PrintOrdered(pair.Value, indent+2)
		}
		fmt.Println(string(prefix) + "}")
	case types.ARRAY:
		fmt.Println(string(prefix) + "[")
		for _, v := range val.Value.(OrderedArray) {
			PrintOrdered(v, indent+1)
		}
		fmt.Println(string(prefix) + "]")
	default:
		out, _ := json.Marshal(val.Value)
		fmt.Println(string(prefix) + string(out))
	}
}
