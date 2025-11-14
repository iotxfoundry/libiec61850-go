package libiec61850go

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo linux,amd64 LDFLAGS: -static -L${SRCDIR}/3rdParty/linux_amd64/iec61850/lib/libiec61850.a
#cgo linux,arm LDFLAGS: -static -L${SRCDIR}/3rdParty/linux_armv7/iec61850/lib/libiec61850.a
#cgo linux,arm64 LDFLAGS: -static -L${SRCDIR}/3rdParty/linux_armv8/iec61850/lib/libiec61850.a
#cgo linux,386 LDFLAGS: -static -L${SRCDIR}/3rdParty/linux_386/iec61850/lib/libiec61850.a
#cgo windows LDFLAGS: -static -L${SRCDIR}/3rdParty/windows_amd64/iec61850/lib/iec61850.lib

#include <stdlib.h>
#include "linked_list.h"

extern void fLinkedListValueDeleteFunctionGo(void* data);
*/
import "C"
import (
	"sync"
	"unsafe"
)

type LinkedListValueDeleteFunction func(value unsafe.Pointer)

type LinkedList struct {
	ptr C.LinkedList
}

func LinkedListCreate() *LinkedList {
	return &LinkedList{ptr: C.LinkedList_create()}
}

func (l *LinkedList) Destroy() {
	C.LinkedList_destroy(l.ptr)
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

func (l *LinkedList) DestroyDeep(fn LinkedListValueDeleteFunction) {
	mapLinkedListValueDeleteFunctionCallbacks.Store(l.ptr, fn)
	defer mapLinkedListValueDeleteFunctionCallbacks.Delete(l.ptr)
	C.LinkedList_destroyDeep(l.ptr, C.LinkedListValueDeleteFunction(C.fLinkedListValueDeleteFunctionGo))
}

func (l *LinkedList) DestroyStatic() {
	C.LinkedList_destroyStatic(l.ptr)
}

func (l *LinkedList) Add(data unsafe.Pointer) {
	C.LinkedList_add(l.ptr, data)
}

func (l *LinkedList) Contains(data unsafe.Pointer) bool {
	return (bool)(C.LinkedList_contains(l.ptr, data))
}

func (l *LinkedList) Remove(data unsafe.Pointer) bool {
	return (bool)(C.LinkedList_remove(l.ptr, data))
}

func (l *LinkedList) Get(index int) *LinkedList {
	return &LinkedList{
		ptr: C.LinkedList_get(l.ptr, C.int(index)),
	}
}

func (l *LinkedList) GetNext() *LinkedList {
	return &LinkedList{
		ptr: C.LinkedList_getNext(l.ptr),
	}
}

func (l *LinkedList) GetLastElement() *LinkedList {
	return &LinkedList{
		ptr: C.LinkedList_getLastElement(l.ptr),
	}
}

func (l *LinkedList) InsertAfter(data unsafe.Pointer) {
	C.LinkedList_insertAfter(l.ptr, data)
}

func (l *LinkedList) Size() int {
	return int(C.LinkedList_size(l.ptr))
}

func (l *LinkedList) GetData() unsafe.Pointer {
	return C.LinkedList_getData(l.ptr)
}

func (l *LinkedList) PrintStringList() {
	C.LinkedList_printStringList(l.ptr)
}
