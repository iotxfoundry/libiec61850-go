package libiec61850go

/*
#include <stdlib.h>
#include "mms_type_spec.h"

*/
import "C"

func (x *MmsVariableSpecification) Destroy() {
	C.MmsVariableSpecification_destroy(x.ctx)
}

func (x *MmsVariableSpecification) GetChildValue(value *MmsValue, childId string) *MmsValue {
	return &MmsValue{
		ctx: C.MmsVariableSpecification_getChildValue(x.ctx, value.ctx, StringData(childId)),
	}
}

func (x *MmsVariableSpecification) GetNamedVariableRecursive(nameId string) *MmsVariableSpecification {
	return &MmsVariableSpecification{
		ctx: C.MmsVariableSpecification_getNamedVariableRecursive(x.ctx, StringData(nameId)),
	}
}

func (x *MmsVariableSpecification) GetType() MmsType {
	return MmsType(C.MmsVariableSpecification_getType(x.ctx))
}

func (x *MmsVariableSpecification) IsValueOfType(value *MmsValue) bool {
	return bool(C.MmsVariableSpecification_isValueOfType(x.ctx, value.ctx))
}

func (x *MmsVariableSpecification) GetName() string {
	return C.GoString(C.MmsVariableSpecification_getName(x.ctx))
}

func (x *MmsVariableSpecification) GetStructureElements() *LinkedList {
	return &LinkedList{
		ctx: C.MmsVariableSpecification_getStructureElements(x.ctx),
	}
}

func (x *MmsVariableSpecification) GetSize() int {
	return int(C.MmsVariableSpecification_getSize(x.ctx))
}

func (x *MmsVariableSpecification) GetChildSpecificationByIndex(index int) *MmsVariableSpecification {
	return &MmsVariableSpecification{
		ctx: C.MmsVariableSpecification_getChildSpecificationByIndex(x.ctx, C.int(index)),
	}
}

func (x *MmsVariableSpecification) GetChildSpecificationByName(name string) (*MmsVariableSpecification, int) {
	var index C.int
	return &MmsVariableSpecification{
		ctx: C.MmsVariableSpecification_getChildSpecificationByName(x.ctx, StringData(name), &index),
	}, int(index)
}

func (x *MmsVariableSpecification) GetArrayElementSpecification() *MmsVariableSpecification {
	return &MmsVariableSpecification{
		ctx: C.MmsVariableSpecification_getArrayElementSpecification(x.ctx),
	}
}

func (x *MmsVariableSpecification) GetExponentWidth() int {
	return int(C.MmsVariableSpecification_getExponentWidth(x.ctx))
}
