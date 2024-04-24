package qdrant

// This file contains methods to convert a generic map to Map of string to *grpc.Value(Qdrant payload type).
// This is a custom implementatation based on "google.golang.org/protobuf/types/known/structpb".
// It extends the original implementation to support IntegerValue and DoubleValue instead of a single NumberValue.
//
// USAGE:
//
// jsonMap := map[string]interface{}{
// 	"some_null":    nil,
// 	"some_bool":    true,
// 	"some_int":     42,
// 	"some_float":   3.14,
// 	"some_string":  "hello",
// 	"some_bytes":   []byte("world"),
// 	"some_nested":  map[string]interface{}{"key": "value"},
// 	"some_list":    []interface{}{"foo", 32},
// }
//
// valueMap := NewValueMap(jsonMap)

import (
	"encoding/base64"
	"fmt"
	"unicode/utf8"

	"github.com/qdrant/go-client/grpc"
)

// Converts a map of string to interface{} to a map of string to *grpc.Value
//
//	╔════════════════════════╤════════════════════════════════════════════╗
//	║ Go type                │ Conversion                                 ║
//	╠════════════════════════╪════════════════════════════════════════════╣
//	║ nil                    │ stored as NullValue                        ║
//	║ bool                   │ stored as BoolValue                        ║
//	║ int, int32, int64      │ stored as IntegerValue                     ║
//	║ uint, uint32, uint64   │ stored as IntegerValue                     ║
//	║ float32, float64       │ stored as DoubleValue                      ║
//	║ string                 │ stored as StringValue; must be valid UTF-8 ║
//	║ []byte                 │ stored as StringValue; base64-encoded      ║
//	║ map[string]interface{} │ stored as StructValue                      ║
//	║ []interface{}          │ stored as ListValue                        ║
//	╚════════════════════════╧════════════════════════════════════════════╝

func NewValueMap(inputMap map[string]interface{}) map[string]*grpc.Value {
	valueMap := make(map[string]*grpc.Value)
	for key, val := range inputMap {
		value, err := NewValue(val)
		if err != nil {
			panic(err)
		}
		valueMap[key] = value
	}
	return valueMap
}

// NewValue constructs a *grpc.Value from a general-purpose Go interface.
func NewValue(v interface{}) (*grpc.Value, error) {
	switch v := v.(type) {
	case nil:
		return NewNullValue(), nil
	case bool:
		return NewBoolValue(v), nil
	case int:
		return NewIntegerValue(int64(v)), nil
	case int32:
		return NewIntegerValue(int64(v)), nil
	case int64:
		return NewIntegerValue(int64(v)), nil
	case uint:
		return NewIntegerValue(int64(v)), nil
	case uint32:
		return NewIntegerValue(int64(v)), nil
	case uint64:
		return NewIntegerValue(int64(v)), nil
	case float32:
		return NewDoubleValue(float64(v)), nil
	case float64:
		return NewDoubleValue(float64(v)), nil
	case string:
		if !utf8.ValidString(v) {
			return nil, fmt.Errorf("invalid UTF-8 in string: %q", v)
		}
		return NewStringValue(v), nil
	case []byte:
		s := base64.StdEncoding.EncodeToString(v)
		return NewStringValue(s), nil
	case map[string]interface{}:
		v2, err := NewStruct(v)
		if err != nil {
			return nil, err
		}
		return NewStructValue(v2), nil
	case []interface{}:
		v2, err := NewList(v)
		if err != nil {
			return nil, err
		}
		return NewListValue(v2), nil
	default:
		return nil, fmt.Errorf("invalid type: %T", v)
	}
}

// NewNullValue constructs a new null Value.
func NewNullValue() *grpc.Value {
	return &grpc.Value{Kind: &grpc.Value_NullValue{NullValue: grpc.NullValue_NULL_VALUE}}
}

// NewBoolValue constructs a new boolean Value.
func NewBoolValue(v bool) *grpc.Value {
	return &grpc.Value{Kind: &grpc.Value_BoolValue{BoolValue: v}}
}

// NewInteger constructs a new number Value.
func NewIntegerValue(v int64) *grpc.Value {
	return &grpc.Value{Kind: &grpc.Value_IntegerValue{IntegerValue: v}}
}

// NewNumberValue constructs a new number Value.
func NewDoubleValue(v float64) *grpc.Value {
	return &grpc.Value{Kind: &grpc.Value_DoubleValue{DoubleValue: v}}
}

// NewStringValue constructs a new string Value.
func NewStringValue(v string) *grpc.Value {
	return &grpc.Value{Kind: &grpc.Value_StringValue{StringValue: v}}
}

// NewStructValue constructs a new struct Value.
func NewStructValue(v *grpc.Struct) *grpc.Value {
	return &grpc.Value{Kind: &grpc.Value_StructValue{StructValue: v}}
}

// NewListValue constructs a new list Value.
func NewListValue(v *grpc.ListValue) *grpc.Value {
	return &grpc.Value{Kind: &grpc.Value_ListValue{ListValue: v}}
}

// NewList constructs a ListValue from a general-purpose Go slice.
// The slice elements are converted using NewValue.
func NewList(v []interface{}) (*grpc.ListValue, error) {
	x := &grpc.ListValue{Values: make([]*grpc.Value, len(v))}
	for i, v := range v {
		var err error
		x.Values[i], err = NewValue(v)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}

// NewStruct constructs a Struct from a general-purpose Go map.
// The map keys must be valid UTF-8.
// The map values are converted using NewValue.
func NewStruct(v map[string]interface{}) (*grpc.Struct, error) {
	x := &grpc.Struct{Fields: make(map[string]*grpc.Value, len(v))}
	for k, v := range v {
		if !utf8.ValidString(k) {
			return nil, fmt.Errorf("invalid UTF-8 in string: %q", k)
		}
		var err error
		x.Fields[k], err = NewValue(v)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}
