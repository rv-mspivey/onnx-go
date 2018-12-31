package onnx

import (
	fmt "fmt"

	"github.com/pkg/errors"
)

var (
	// ErrNotYetImplemented ...
	ErrNotYetImplemented = errors.New("Not Yet Implemented")
	// ErrNoDataFound ...
	ErrNoDataFound = errors.New("No data found")
	// ErrCorruptedData ...
	ErrCorruptedData = errors.New("Unable to decode data")
)

// ErrNotImplemented is returned for any operator or attribute
type ErrNotImplemented struct {
	Operator       string
	AttributeName  string
	AttributeValue interface{}
	Message        string
}

func (e *ErrNotImplemented) Error() string {
	if e.AttributeName != "" {
		return fmt.Sprintf("onnx: operator %v. Implementation error for attribute %v (%#v): %v",
			e.Operator,
			e.AttributeName,
			e.AttributeValue,
			e.Message)
	}
	return "onnx: operator " + e.Operator + " not implemented"
}

// ErrInvalidModel is raised if we are not able to unmarshal the model because it is invalid
type ErrInvalidModel struct {
	NodeNotDefined string // Happens if the corresponding node is used in the graph, but not present in the input or output fields
}

func (e *ErrInvalidModel) Error() string {
	err := "Invalid graph: "
	if e.NodeNotDefined != "" {
		return fmt.Sprintf("%v: node %v is referenced in the Node lists but not defined nor in []input nor in []output", err, e.NodeNotDefined)
	}
	return err

}
