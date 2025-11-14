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

type LinkedListValueDeleteFunction func(value unsafe.Pointer)

type LinkedList struct {
	ctx C.LinkedList
}

func LinkedListCreate() *LinkedList {
	return &LinkedList{ctx: C.LinkedList_create()}
}

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

func (l *LinkedList) DestroyDeep(fn LinkedListValueDeleteFunction) {
	mapLinkedListValueDeleteFunctionCallbacks.Store(l.ctx, fn)
	defer mapLinkedListValueDeleteFunctionCallbacks.Delete(l.ctx)
	C.LinkedList_destroyDeep(l.ctx, C.LinkedListValueDeleteFunction(C.fLinkedListValueDeleteFunctionGo))
}

func (l *LinkedList) DestroyStatic() {
	C.LinkedList_destroyStatic(l.ctx)
}

func (l *LinkedList) Add(data unsafe.Pointer) {
	C.LinkedList_add(l.ctx, data)
}

func (l *LinkedList) Contains(data unsafe.Pointer) bool {
	return (bool)(C.LinkedList_contains(l.ctx, data))
}

func (l *LinkedList) Remove(data unsafe.Pointer) bool {
	return (bool)(C.LinkedList_remove(l.ctx, data))
}

func (l *LinkedList) Get(index int) *LinkedList {
	return &LinkedList{
		ctx: C.LinkedList_get(l.ctx, C.int(index)),
	}
}

func (l *LinkedList) GetNext() *LinkedList {
	return &LinkedList{
		ctx: C.LinkedList_getNext(l.ctx),
	}
}

func (l *LinkedList) GetLastElement() *LinkedList {
	return &LinkedList{
		ctx: C.LinkedList_getLastElement(l.ctx),
	}
}

func (l *LinkedList) InsertAfter(data unsafe.Pointer) {
	C.LinkedList_insertAfter(l.ctx, data)
}

func (l *LinkedList) Size() int {
	return int(C.LinkedList_size(l.ctx))
}

func (l *LinkedList) GetData() unsafe.Pointer {
	return C.LinkedList_getData(l.ctx)
}

func (l *LinkedList) PrintStringList() {
	C.LinkedList_printStringList(l.ctx)
}
