package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// https://gist.github.com/tkrajina/aec8d1b15b088c20f0df4afcd5f0c511
type Something interface{}

type Something1 struct {
	Aaa, Bbb, Ccc string
}

type Something2 struct {
	Ccc, Ddd string
}

// // The declaration does nothing, but if T does not implement the interface I, compile fails
// var _ I = (*T)(nil) // Verify that *T implements I.
// @see https://github.com/uber-go/guide/issues/25
// @see https://www.reddit.com/r/golang/comments/m1hfl7/what_does_this_syntax_mean_tnil/
var _ Something = (*Something1)(nil)
var _ Something = (*Something2)(nil)

type Container struct {
	Type  string    `json:"type"`
	Value Something `json:"value"`
}

func (c *Container) UnmarshalJSON(data []byte) error {
	value, err := UnmarshalCustomValue(data, "type", "value", map[string]reflect.Type{
		"something1": reflect.TypeOf(Something1{}),
		"something2": reflect.TypeOf(Something2{}),
	})
	if err != nil {
		return err
	}

	c.Value = value

	return nil
}

func UnmarshalCustomValue(data []byte, typeJsonField, valueJsonField string, customTypes map[string]reflect.Type) (interface{}, error) {
	m := map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	typeName := m[typeJsonField].(string)
	var value Something
	if ty, found := customTypes[typeName]; found {
		value = reflect.New(ty).Interface().(Something)
	}

	valueBytes, err := json.Marshal(m[valueJsonField])
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(valueBytes, &value); err != nil {
		return nil, err
	}

	return value, nil
}

var _ json.Unmarshaler = (*Container)(nil)

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	testUnmarshalling(`{"type":"something1","value":{"Aaa": "a", "bbb": "bb", "ccc": "cc"}}`)
	testUnmarshalling(`{"type":"something2","value":{"Ccc": "a"}}`)
}

func testUnmarshalling(jsonStr string) {
	var container Container
	err := json.Unmarshal([]byte(jsonStr), &container)
	panicIfErr(err)
	fmt.Printf("container=%+v\n", container)
	fmt.Printf("value=%#v\n", container.Value)
}
