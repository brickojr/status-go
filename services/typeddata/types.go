package typeddata

import (
	"errors"
	"fmt"
)

const (
	eip712Domain = "EIP712Domain"
)

type Types map[string][]Field

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (f Field) Validate() error {
	if len(f.Name) == 0 {
		return errors.New("`name` is required")
	}
	if len(f.Type) == 0 {
		return errors.New("`type` is required")
	}
	return nil
}

type TypedData struct {
	Types       Types                  `json:"types"`
	PrimaryType string                 `json:"primaryType"`
	Domain      map[string]interface{} `json:"domain"`
	Message     map[string]interface{} `json:"message"`
}

func (t TypedData) Validate() error {
	if _, exist := t.Types[eip712Domain]; !exist {
		return fmt.Errorf("`%s` must be in `types`", eip712Domain)
	}
	if t.PrimaryType == "" {
		return errors.New("`primaryType` is required")
	}
	if _, exist := t.Types[t.PrimaryType]; !exist {
		return fmt.Errorf("primary type `%s` not defined in types", t.PrimaryType)
	}
	if t.Domain == nil {
		return errors.New("`domain` is required")
	}
	if t.Message == nil {
		return errors.New("`message` is required")
	}
	for typ := range t.Types {
		fields := t.Types[typ]
		for i := range fields {
			if err := fields[i].Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}