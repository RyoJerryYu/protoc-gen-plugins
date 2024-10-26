//Code generated by protoc-gen-go-json. DO NOT EDIT.
//versions:
//- protoc-gen-go-json v1.0.7
//- protoc 5.28.2
//source: e2e_enum.proto

package e2e

import (
	bytes "bytes"
	json "encoding/json"
	fmt "fmt"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (enum Enum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (enum *Enum) UnmarshalJSON(b []byte) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	v, err := dec.Token()
	if err != nil {
		return err
	}
	switch v := v.(type) {
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return err
		}
		*enum = Enum(n)
	case float64:
		*enum = Enum(v)
	case string:
		*enum = Enum(Enum_value[v])
	default:
		return fmt.Errorf("invalid enum value %v", v)
	}
	return nil
}

// MarshalJSON implements json.Marshaler
func (msg *EnumParent) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *EnumParent) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (enum EnumParent_EnumChild) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (enum *EnumParent_EnumChild) UnmarshalJSON(b []byte) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	v, err := dec.Token()
	if err != nil {
		return err
	}
	switch v := v.(type) {
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return err
		}
		*enum = EnumParent_EnumChild(n)
	case float64:
		*enum = EnumParent_EnumChild(v)
	case string:
		*enum = EnumParent_EnumChild(EnumParent_EnumChild_value[v])
	default:
		return fmt.Errorf("invalid enum value %v", v)
	}
	return nil
}
