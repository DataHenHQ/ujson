package ujson

import (
	"bytes"
	"encoding/json"
)

// JSONPretty does marshalling without escaping html
func JSONPretty(t interface{}, indent string) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if indent != "" {
		encoder.SetIndent("", indent)
	}
	err := encoder.Encode(t)
	if err != nil {
		return []byte{}
	}

	return bytes.TrimSpace(buffer.Bytes())
}

// MarshalIndent is like the json package's MarshalIndent but it uses unescaped html
func MarshalIndent(t interface{}, prefix, indent string) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if indent != "" {
		encoder.SetIndent(prefix, indent)
	}
	err := encoder.Encode(t)
	if err != nil {
		return nil, err
	}

	return bytes.TrimSpace(buffer.Bytes()), nil
}

// Marshal is like the json package's Marshal but it uses unescaped html
func Marshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(t)
	if err != nil {
		return nil, err
	}

	return bytes.TrimSpace(buffer.Bytes()), nil
}
