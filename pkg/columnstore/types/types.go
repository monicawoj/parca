package types

import "fmt"

type Type interface {
	NewValue(data interface{}) Value
}

type MapType struct {
	Key   Type
	Value Type
}

func Map(keyType Type, valueType Type) Type {
	return &MapType{
		Key:   keyType,
		Value: valueType,
	}
}

func (m *MapType) NewValue(data interface{}) Value {
	return Value{
		Type: m,
		Data: data,
	}
}

type StaticType uint64

const (
	NullType StaticType = iota
	String
	Uint64
	Int64
	UUID
)

func (t StaticType) NewValue(data interface{}) Value {
	return Value{
		Type: t,
		Data: data,
	}
}

type Value struct {
	Type Type
	Data interface{}
}

func (v Value) String() string {
	return fmt.Sprint(v.Data)
}

func (v *Value) Less(other Value) bool {
	if v.Data == nil && other.Data == nil {
		return false
	}
	if v.Data == nil && other.Data != nil {
		return false
	}
	if v.Data != nil && other.Data == nil {
		return true
	}

	switch v.Data.(type) {
	case string:
		return v.Data.(string) < other.Data.(string)
	case int64:
		return v.Data.(int64) < other.Data.(int64)
	case uint64:
		return v.Data.(uint64) < other.Data.(uint64)
	default:
		panic("unsupported type")
	}
}

func (v *Value) Equal(other Value) bool {
	if v.Data == nil && other.Data == nil {
		return true
	}
	if v.Data == nil && other.Data != nil {
		return false
	}
	if v.Data != nil && other.Data == nil {
		return false
	}

	switch v.Data.(type) {
	case string:
		return v.Data.(string) == other.Data.(string)
	case int64:
		return v.Data.(int64) == other.Data.(int64)
	case uint64:
		return v.Data.(uint64) == other.Data.(uint64)
	default:
		panic("unsupported type")
	}
}
