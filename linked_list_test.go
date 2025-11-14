package libiec61850go_test

import (
	"testing"
	"unsafe"

	libiec61850go "github.com/iotxfoundry/libiec61850-go"
)

func TestLinkedList_Add(t *testing.T) {

	i32v := int32(123456)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data unsafe.Pointer
	}{
		{
			name: "Add nil",
			data: nil,
		},
		{
			name: "Add non-nil",
			data: unsafe.Pointer(&i32v),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := libiec61850go.LinkedListCreate()
			l.Add(tt.data)
			l.PrintStringList()
		})
	}
}
