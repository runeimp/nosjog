package nosjog

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
)

/*
 * CONSTANTS
 */
const (
	Name    = "NOSJoG"
	Version = "0.1.0"
)

// Dynamic converts JSON bytes to an interface{} object
func Dynamic(jsonBytes []byte) interface{} {
	var anything interface{}
	json.Unmarshal(jsonBytes, &anything)
	return anything
}

// Marshal encodes objects without HTML escaping the data
func Marshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// PrettyJSON converts JSON bytes to a JSON string
func PrettyJSON(jsonBytes []byte) string {
	buffer := new(bytes.Buffer)
	json.Indent(buffer, jsonBytes, "", "  ")
	return fmt.Sprintf("%s", buffer)
}

// Pretty formats JSON bytes to be pretty
func Pretty(jsonBytes []byte) []byte {
	prettyOptions := pretty.DefaultOptions
	prettyOptions.SortKeys = true
	return pretty.PrettyOptions(jsonBytes, prettyOptions)
}

// MarshalPrintln takes an object and converts it to JSON bytes then colorizes it for terminal output before printing
func MarshalPrintln(content interface{}) {
	jsonBytes, _ := Marshal(content)
	jsonBytes = Terminal(jsonBytes)
	fmt.Println(string(jsonBytes))
}

// Println takes JSON bytes then colorizes it for terminal output before printing
func Println(jsonBytes []byte) {
	jsonBytes = Terminal(jsonBytes)
	fmt.Println(string(jsonBytes))
}

// Terminal colorizes JSON bytes with terminal escape codes
func Terminal(jsonBytes []byte) []byte {
	prettyJSON := Pretty(jsonBytes)
	return pretty.Color(prettyJSON, nil)
}
