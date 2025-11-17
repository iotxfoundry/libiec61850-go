package libiec61850go

/*
#include <stdlib.h>
#include "mms_common.h"

*/
import "C"

type MmsError int32

const (
	/* generic error codes */
	MMS_ERROR_NONE                   MmsError = C.MMS_ERROR_NONE
	MMS_ERROR_CONNECTION_REJECTED    MmsError = C.MMS_ERROR_CONNECTION_REJECTED
	MMS_ERROR_CONNECTION_LOST        MmsError = C.MMS_ERROR_CONNECTION_LOST
	MMS_ERROR_SERVICE_TIMEOUT        MmsError = C.MMS_ERROR_SERVICE_TIMEOUT
	MMS_ERROR_PARSING_RESPONSE       MmsError = C.MMS_ERROR_PARSING_RESPONSE
	MMS_ERROR_HARDWARE_FAULT         MmsError = C.MMS_ERROR_HARDWARE_FAULT
	MMS_ERROR_CONCLUDE_REJECTED      MmsError = C.MMS_ERROR_CONCLUDE_REJECTED
	MMS_ERROR_INVALID_ARGUMENTS      MmsError = C.MMS_ERROR_INVALID_ARGUMENTS
	MMS_ERROR_OUTSTANDING_CALL_LIMIT MmsError = C.MMS_ERROR_OUTSTANDING_CALL_LIMIT

	MMS_ERROR_OTHER MmsError = C.MMS_ERROR_OTHER

	/* confirmed error PDU codes */
	MMS_ERROR_VMDSTATE_OTHER MmsError = C.MMS_ERROR_VMDSTATE_OTHER

	MMS_ERROR_APPLICATION_REFERENCE_OTHER MmsError = C.MMS_ERROR_APPLICATION_REFERENCE_OTHER

	MMS_ERROR_DEFINITION_OTHER                         MmsError = C.MMS_ERROR_DEFINITION_OTHER
	MMS_ERROR_DEFINITION_INVALID_ADDRESS               MmsError = C.MMS_ERROR_DEFINITION_INVALID_ADDRESS
	MMS_ERROR_DEFINITION_TYPE_UNSUPPORTED              MmsError = C.MMS_ERROR_DEFINITION_TYPE_UNSUPPORTED
	MMS_ERROR_DEFINITION_TYPE_INCONSISTENT             MmsError = C.MMS_ERROR_DEFINITION_TYPE_INCONSISTENT
	MMS_ERROR_DEFINITION_OBJECT_UNDEFINED              MmsError = C.MMS_ERROR_DEFINITION_OBJECT_UNDEFINED
	MMS_ERROR_DEFINITION_OBJECT_EXISTS                 MmsError = C.MMS_ERROR_DEFINITION_OBJECT_EXISTS
	MMS_ERROR_DEFINITION_OBJECT_ATTRIBUTE_INCONSISTENT MmsError = C.MMS_ERROR_DEFINITION_OBJECT_ATTRIBUTE_INCONSISTENT

	MMS_ERROR_RESOURCE_OTHER                  MmsError = C.MMS_ERROR_RESOURCE_OTHER
	MMS_ERROR_RESOURCE_CAPABILITY_UNAVAILABLE MmsError = C.MMS_ERROR_RESOURCE_CAPABILITY_UNAVAILABLE

	MMS_ERROR_SERVICE_OTHER                      MmsError = C.MMS_ERROR_SERVICE_OTHER
	MMS_ERROR_SERVICE_OBJECT_CONSTRAINT_CONFLICT MmsError = C.MMS_ERROR_SERVICE_OBJECT_CONSTRAINT_CONFLICT

	MMS_ERROR_SERVICE_PREEMPT_OTHER MmsError = C.MMS_ERROR_SERVICE_PREEMPT_OTHER

	MMS_ERROR_TIME_RESOLUTION_OTHER MmsError = C.MMS_ERROR_TIME_RESOLUTION_OTHER

	MMS_ERROR_ACCESS_OTHER                     MmsError = C.MMS_ERROR_ACCESS_OTHER
	MMS_ERROR_ACCESS_OBJECT_NON_EXISTENT       MmsError = C.MMS_ERROR_ACCESS_OBJECT_NON_EXISTENT
	MMS_ERROR_ACCESS_OBJECT_ACCESS_UNSUPPORTED MmsError = C.MMS_ERROR_ACCESS_OBJECT_ACCESS_UNSUPPORTED
	MMS_ERROR_ACCESS_OBJECT_ACCESS_DENIED      MmsError = C.MMS_ERROR_ACCESS_OBJECT_ACCESS_DENIED
	MMS_ERROR_ACCESS_OBJECT_INVALIDATED        MmsError = C.MMS_ERROR_ACCESS_OBJECT_INVALIDATED
	MMS_ERROR_ACCESS_OBJECT_VALUE_INVALID      MmsError = C.MMS_ERROR_ACCESS_OBJECT_VALUE_INVALID    /* for DataAccessError 11 */
	MMS_ERROR_ACCESS_TEMPORARILY_UNAVAILABLE   MmsError = C.MMS_ERROR_ACCESS_TEMPORARILY_UNAVAILABLE /* for DataAccessError 2 */

	MMS_ERROR_FILE_OTHER                           MmsError = C.MMS_ERROR_FILE_OTHER
	MMS_ERROR_FILE_FILENAME_AMBIGUOUS              MmsError = C.MMS_ERROR_FILE_FILENAME_AMBIGUOUS
	MMS_ERROR_FILE_FILE_BUSY                       MmsError = C.MMS_ERROR_FILE_FILE_BUSY
	MMS_ERROR_FILE_FILENAME_SYNTAX_ERROR           MmsError = C.MMS_ERROR_FILE_FILENAME_SYNTAX_ERROR
	MMS_ERROR_FILE_CONTENT_TYPE_INVALID            MmsError = C.MMS_ERROR_FILE_CONTENT_TYPE_INVALID
	MMS_ERROR_FILE_POSITION_INVALID                MmsError = C.MMS_ERROR_FILE_POSITION_INVALID
	MMS_ERROR_FILE_FILE_ACCESS_DENIED              MmsError = C.MMS_ERROR_FILE_FILE_ACCESS_DENIED
	MMS_ERROR_FILE_FILE_NON_EXISTENT               MmsError = C.MMS_ERROR_FILE_FILE_NON_EXISTENT
	MMS_ERROR_FILE_DUPLICATE_FILENAME              MmsError = C.MMS_ERROR_FILE_DUPLICATE_FILENAME
	MMS_ERROR_FILE_INSUFFICIENT_SPACE_IN_FILESTORE MmsError = C.MMS_ERROR_FILE_INSUFFICIENT_SPACE_IN_FILESTORE

	/* reject codes */
	MMS_ERROR_REJECT_OTHER                    MmsError = C.MMS_ERROR_REJECT_OTHER
	MMS_ERROR_REJECT_UNKNOWN_PDU_TYPE         MmsError = C.MMS_ERROR_REJECT_UNKNOWN_PDU_TYPE
	MMS_ERROR_REJECT_INVALID_PDU              MmsError = C.MMS_ERROR_REJECT_INVALID_PDU
	MMS_ERROR_REJECT_UNRECOGNIZED_SERVICE     MmsError = C.MMS_ERROR_REJECT_UNRECOGNIZED_SERVICE
	MMS_ERROR_REJECT_UNRECOGNIZED_MODIFIER    MmsError = C.MMS_ERROR_REJECT_UNRECOGNIZED_MODIFIER
	MMS_ERROR_REJECT_REQUEST_INVALID_ARGUMENT MmsError = C.MMS_ERROR_REJECT_REQUEST_INVALID_ARGUMENT
)

type MmsType int32

const (
	/*! this represents all MMS array types (arrays contain uniform elements) */
	MMS_ARRAY MmsType = C.MMS_ARRAY
	/*! this represents all complex MMS types (structures) */
	MMS_STRUCTURE MmsType = C.MMS_STRUCTURE
	/*! boolean value */
	MMS_BOOLEAN MmsType = C.MMS_BOOLEAN
	/*! bit string */
	MMS_BIT_STRING MmsType = C.MMS_BIT_STRING
	/*! represents all signed integer types */
	MMS_INTEGER MmsType = C.MMS_INTEGER
	/*! represents all unsigned integer types */
	MMS_UNSIGNED MmsType = C.MMS_UNSIGNED
	/*! represents all float type (32 and 64 bit) */
	MMS_FLOAT MmsType = C.MMS_FLOAT
	/*! octet string (unstructured bytes) */
	MMS_OCTET_STRING MmsType = C.MMS_OCTET_STRING
	/*! MMS visible string */
	MMS_VISIBLE_STRING MmsType = C.MMS_VISIBLE_STRING
	/*! MMS generalized time type */
	MMS_GENERALIZED_TIME MmsType = C.MMS_GENERALIZED_TIME
	/*! MMS binary time type */
	MMS_BINARY_TIME MmsType = C.MMS_BINARY_TIME
	/*! MMS BCD type */
	MMS_BCD MmsType = C.MMS_BCD
	/*! MMS object identifier type */
	MMS_OBJ_ID MmsType = C.MMS_OBJ_ID
	/*! MMS unicode string */
	MMS_STRING MmsType = C.MMS_STRING
	/*! MMS UTC time type */
	MMS_UTC_TIME MmsType = C.MMS_UTC_TIME
	/*! This represents an error code as returned by MMS read services */
	MMS_DATA_ACCESS_ERROR MmsType = C.MMS_DATA_ACCESS_ERROR
)

type MmsDomain struct {
	ctx *C.MmsDomain
}

type MmsAccessSpecifier struct {
	ctx *C.MmsAccessSpecifier
}

func (x *MmsAccessSpecifier) Domain() *MmsDomain {
	return &MmsDomain{ctx: x.ctx.domain}
}

func (x *MmsAccessSpecifier) VariableName() string {
	return C.GoString(x.ctx.variableName)
}

func (x *MmsAccessSpecifier) ArrayIndex() int {
	return int(x.ctx.arrayIndex)
}

func (x *MmsAccessSpecifier) ComponentName() string {
	return C.GoString(x.ctx.componentName)
}

type MmsVariableAccessSpecification struct {
	ctx *C.MmsVariableAccessSpecification
}

func (x *MmsVariableAccessSpecification) DomainId() string {
	return C.GoString(x.ctx.domainId)
}

func (x *MmsVariableAccessSpecification) ItemId() string {
	return C.GoString(x.ctx.itemId)
}

func (x *MmsVariableAccessSpecification) ArrayIndex() int {
	return int(x.ctx.arrayIndex)
}

func (x *MmsVariableAccessSpecification) ComponentName() string {
	return C.GoString(x.ctx.componentName)
}

type MmsNamedVariableList struct {
	ctx C.MmsNamedVariableList
}

type MmsNamedVariableListEntry struct {
	ctx C.MmsNamedVariableListEntry
}

type ItuObjectIdentifier struct {
	ctx *C.ItuObjectIdentifier
}

func (x *ItuObjectIdentifier) ArcCount() int {
	return int(x.ctx.arcCount)
}

func (x *ItuObjectIdentifier) Arc() [16]uint16 {
	arc := [16]uint16{}
	for i := range 16 {
		arc[i] = uint16(x.ctx.arc[i])
	}
	return arc
}

type IsoApplicationReference struct {
	ctx *C.IsoApplicationReference
}

func (x *IsoApplicationReference) ApTitle() ItuObjectIdentifier {
	return ItuObjectIdentifier{ctx: &x.ctx.apTitle}
}

func (x *IsoApplicationReference) AeQualifier() int {
	return int(x.ctx.aeQualifier)
}
