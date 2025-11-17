// Package libiec61850go provides Go bindings for the libIEC61850 library,
// which is an open-source implementation of the IEC 61850 standard for
// communication with intelligent electronic devices in power systems.
package libiec61850go

/*
#include <stdlib.h>
#include "linked_list.h"

extern void fLinkedListValueDeleteFunctionGo(void* data);
*/
import "C"
import (
	"sync"
	"unsafe"
)

// LinkedListValueDeleteFunction is a function type used to delete data elements
// in a deep destroy operation.
type LinkedListValueDeleteFunction func(value unsafe.Pointer)

// LinkedList represents a linked list data structure.
type LinkedList struct {
	ctx C.LinkedList
}

// LinkedListCreate creates a new LinkedList object.
//
// Returns the newly created LinkedList instance.
func LinkedListCreate() *LinkedList {
	return &LinkedList{ctx: C.LinkedList_create()}
}

// Destroy deletes a LinkedList object.
//
// This function destroys the LinkedList object. It will free all data structures
// used by the LinkedList instance. It will call free for all elements of the
// linked list. This function should only be used if simple objects (like
// dynamically allocated strings) are stored in the linked list.
func (l *LinkedList) Destroy() {
	C.LinkedList_destroy(l.ctx)
}

var mapLinkedListValueDeleteFunctionCallbacks = sync.Map{} //

//export fLinkedListValueDeleteFunctionGo
func fLinkedListValueDeleteFunctionGo(data unsafe.Pointer) {
	mapLinkedListValueDeleteFunctionCallbacks.Range(func(k, v any) bool {
		cb, ok := v.(LinkedListValueDeleteFunction)
		if ok {
			cb(data)
		}
		return true
	})
}

// DestroyDeep deletes a LinkedList object with a custom delete function.
//
// This function destroys the LinkedList object. It will free all data structures
// used by the LinkedList instance. It will call a user provided function for
// each data element. This user provided function is responsible to properly
// free the data element.
//
// Parameters:
//   - fn: a function that is called for each data element of the LinkedList
//     with the pointer to the linked list data element.
func (l *LinkedList) DestroyDeep(fn LinkedListValueDeleteFunction) {
	mapLinkedListValueDeleteFunctionCallbacks.Store(l.ctx, fn)
	defer mapLinkedListValueDeleteFunctionCallbacks.Delete(l.ctx)
	C.LinkedList_destroyDeep(l.ctx, C.LinkedListValueDeleteFunction(C.fLinkedListValueDeleteFunctionGo))
}

// DestroyStatic deletes a LinkedList object without freeing the element data.
//
// This function should be used when statically allocated data objects are stored
// in the LinkedList instance. Other use cases would be if the data elements in
// the list should not be deleted.
func (l *LinkedList) DestroyStatic() {
	C.LinkedList_destroyStatic(l.ctx)
}

// Add adds a new element to the list.
//
// This function will add a new data element to the list. The new element will
// be the last element in the list.
//
// Parameters:
//   - data: data to append to the LinkedList instance
func (l *LinkedList) Add(data unsafe.Pointer) {
	C.LinkedList_add(l.ctx, data)
}

// Contains checks if the specified data is contained in the list.
//
// Parameters:
//   - data: data to check for in the LinkedList instance
//
// Returns true if data is part of the list, false otherwise.
func (l *LinkedList) Contains(data unsafe.Pointer) bool {
	return (bool)(C.LinkedList_contains(l.ctx, data))
}

// Remove removes the specified element from the list.
//
// Parameters:
//   - data: data to remove from the LinkedList instance
//
// Returns true if data has been removed from the list, false otherwise.
func (l *LinkedList) Remove(data unsafe.Pointer) bool {
	return (bool)(C.LinkedList_remove(l.ctx, data))
}

// Get returns the list element specified by index (starting with 0).
//
// Parameters:
//   - index: index of the requested element
//
// Returns the LinkedList element at the specified index.
func (l *LinkedList) Get(index int) *LinkedList {
	return &LinkedList{
		ctx: C.LinkedList_get(l.ctx, C.int(index)),
	}
}

// GetNext returns the next element in the list (iterator).
//
// Returns the next LinkedList element.
func (l *LinkedList) GetNext() *LinkedList {
	return &LinkedList{
		ctx: C.LinkedList_getNext(l.ctx),
	}
}

// GetLastElement returns the last element in the list.
//
// Returns the last LinkedList element.
func (l *LinkedList) GetLastElement() *LinkedList {
	return &LinkedList{
		ctx: C.LinkedList_getLastElement(l.ctx),
	}
}

// InsertAfter inserts a new element in the list after the current element.
//
// Parameters:
//   - data: data to insert in the LinkedList instance
func (l *LinkedList) InsertAfter(data unsafe.Pointer) {
	C.LinkedList_insertAfter(l.ctx, data)
}

// Size returns the size of the list.
//
// Returns the number of data elements stored in the list.
func (l *LinkedList) Size() int {
	return int(C.LinkedList_size(l.ctx))
}

// GetData returns the data stored in the current list element.
//
// Returns a pointer to the data stored in the current list element.
func (l *LinkedList) GetData() unsafe.Pointer {
	return C.LinkedList_getData(l.ctx)
}

// PrintStringList prints the string list.
//
// This function is primarily used for debugging purposes.
func (l *LinkedList) PrintStringList() {
	C.LinkedList_printStringList(l.ctx)
}
