// Package libiec61850go provides Go bindings for the libIEC61850 library,
// which is an open-source implementation of the IEC 61850 standard for
// communication with intelligent electronic devices in power systems.
package libiec61850go

/*
#include <stdlib.h>
#include "mms_common.h"

*/
import "C"
import "errors"

// MmsError represents MMS error codes.
type MmsError int32

// MMS error code constants.
const (
	/* generic error codes */
	// MMS_ERROR_NONE indicates no error occurred.
	MMS_ERROR_NONE MmsError = C.MMS_ERROR_NONE
	// MMS_ERROR_CONNECTION_REJECTED indicates the connection was rejected.
	MMS_ERROR_CONNECTION_REJECTED MmsError = C.MMS_ERROR_CONNECTION_REJECTED
	// MMS_ERROR_CONNECTION_LOST indicates the connection was lost.
	MMS_ERROR_CONNECTION_LOST MmsError = C.MMS_ERROR_CONNECTION_LOST
	// MMS_ERROR_SERVICE_TIMEOUT indicates a service timeout occurred.
	MMS_ERROR_SERVICE_TIMEOUT MmsError = C.MMS_ERROR_SERVICE_TIMEOUT
	// MMS_ERROR_PARSING_RESPONSE indicates an error occurred while parsing the response.
	MMS_ERROR_PARSING_RESPONSE MmsError = C.MMS_ERROR_PARSING_RESPONSE
	// MMS_ERROR_HARDWARE_FAULT indicates a hardware fault occurred.
	MMS_ERROR_HARDWARE_FAULT MmsError = C.MMS_ERROR_HARDWARE_FAULT
	// MMS_ERROR_CONCLUDE_REJECTED indicates the conclude operation was rejected.
	MMS_ERROR_CONCLUDE_REJECTED MmsError = C.MMS_ERROR_CONCLUDE_REJECTED
	// MMS_ERROR_INVALID_ARGUMENTS indicates invalid arguments were provided.
	MMS_ERROR_INVALID_ARGUMENTS MmsError = C.MMS_ERROR_INVALID_ARGUMENTS
	// MMS_ERROR_OUTSTANDING_CALL_LIMIT indicates the outstanding call limit was exceeded.
	MMS_ERROR_OUTSTANDING_CALL_LIMIT MmsError = C.MMS_ERROR_OUTSTANDING_CALL_LIMIT

	// MMS_ERROR_OTHER indicates an unspecified error.
	MMS_ERROR_OTHER MmsError = C.MMS_ERROR_OTHER

	/* confirmed error PDU codes */
	// MMS_ERROR_VMDSTATE_OTHER indicates an unspecified VMD state error.
	MMS_ERROR_VMDSTATE_OTHER MmsError = C.MMS_ERROR_VMDSTATE_OTHER

	// MMS_ERROR_APPLICATION_REFERENCE_OTHER indicates an unspecified application reference error.
	MMS_ERROR_APPLICATION_REFERENCE_OTHER MmsError = C.MMS_ERROR_APPLICATION_REFERENCE_OTHER

	// MMS_ERROR_DEFINITION_OTHER indicates an unspecified definition error.
	MMS_ERROR_DEFINITION_OTHER MmsError = C.MMS_ERROR_DEFINITION_OTHER
	// MMS_ERROR_DEFINITION_INVALID_ADDRESS indicates an invalid address was provided.
	MMS_ERROR_DEFINITION_INVALID_ADDRESS MmsError = C.MMS_ERROR_DEFINITION_INVALID_ADDRESS
	// MMS_ERROR_DEFINITION_TYPE_UNSUPPORTED indicates the type is unsupported.
	MMS_ERROR_DEFINITION_TYPE_UNSUPPORTED MmsError = C.MMS_ERROR_DEFINITION_TYPE_UNSUPPORTED
	// MMS_ERROR_DEFINITION_TYPE_INCONSISTENT indicates the type is inconsistent.
	MMS_ERROR_DEFINITION_TYPE_INCONSISTENT MmsError = C.MMS_ERROR_DEFINITION_TYPE_INCONSISTENT
	// MMS_ERROR_DEFINITION_OBJECT_UNDEFINED indicates the object is undefined.
	MMS_ERROR_DEFINITION_OBJECT_UNDEFINED MmsError = C.MMS_ERROR_DEFINITION_OBJECT_UNDEFINED
	// MMS_ERROR_DEFINITION_OBJECT_EXISTS indicates the object already exists.
	MMS_ERROR_DEFINITION_OBJECT_EXISTS MmsError = C.MMS_ERROR_DEFINITION_OBJECT_EXISTS
	// MMS_ERROR_DEFINITION_OBJECT_ATTRIBUTE_INCONSISTENT indicates the object attribute is inconsistent.
	MMS_ERROR_DEFINITION_OBJECT_ATTRIBUTE_INCONSISTENT MmsError = C.MMS_ERROR_DEFINITION_OBJECT_ATTRIBUTE_INCONSISTENT

	// MMS_ERROR_RESOURCE_OTHER indicates an unspecified resource error.
	MMS_ERROR_RESOURCE_OTHER MmsError = C.MMS_ERROR_RESOURCE_OTHER
	// MMS_ERROR_RESOURCE_CAPABILITY_UNAVAILABLE indicates a capability is unavailable.
	MMS_ERROR_RESOURCE_CAPABILITY_UNAVAILABLE MmsError = C.MMS_ERROR_RESOURCE_CAPABILITY_UNAVAILABLE

	// MMS_ERROR_SERVICE_OTHER indicates an unspecified service error.
	MMS_ERROR_SERVICE_OTHER MmsError = C.MMS_ERROR_SERVICE_OTHER
	// MMS_ERROR_SERVICE_OBJECT_CONSTRAINT_CONFLICT indicates an object constraint conflict.
	MMS_ERROR_SERVICE_OBJECT_CONSTRAINT_CONFLICT MmsError = C.MMS_ERROR_SERVICE_OBJECT_CONSTRAINT_CONFLICT

	// MMS_ERROR_SERVICE_PREEMPT_OTHER indicates an unspecified service preempt error.
	MMS_ERROR_SERVICE_PREEMPT_OTHER MmsError = C.MMS_ERROR_SERVICE_PREEMPT_OTHER

	// MMS_ERROR_TIME_RESOLUTION_OTHER indicates an unspecified time resolution error.
	MMS_ERROR_TIME_RESOLUTION_OTHER MmsError = C.MMS_ERROR_TIME_RESOLUTION_OTHER

	// MMS_ERROR_ACCESS_OTHER indicates an unspecified access error.
	MMS_ERROR_ACCESS_OTHER MmsError = C.MMS_ERROR_ACCESS_OTHER
	// MMS_ERROR_ACCESS_OBJECT_NON_EXISTENT indicates the object does not exist.
	MMS_ERROR_ACCESS_OBJECT_NON_EXISTENT MmsError = C.MMS_ERROR_ACCESS_OBJECT_NON_EXISTENT
	// MMS_ERROR_ACCESS_OBJECT_ACCESS_UNSUPPORTED indicates object access is unsupported.
	MMS_ERROR_ACCESS_OBJECT_ACCESS_UNSUPPORTED MmsError = C.MMS_ERROR_ACCESS_OBJECT_ACCESS_UNSUPPORTED
	// MMS_ERROR_ACCESS_OBJECT_ACCESS_DENIED indicates object access is denied.
	MMS_ERROR_ACCESS_OBJECT_ACCESS_DENIED MmsError = C.MMS_ERROR_ACCESS_OBJECT_ACCESS_DENIED
	// MMS_ERROR_ACCESS_OBJECT_INVALIDATED indicates the object is invalidated.
	MMS_ERROR_ACCESS_OBJECT_INVALIDATED MmsError = C.MMS_ERROR_ACCESS_OBJECT_INVALIDATED
	// MMS_ERROR_ACCESS_OBJECT_VALUE_INVALID indicates the object value is invalid (for DataAccessError 11).
	MMS_ERROR_ACCESS_OBJECT_VALUE_INVALID MmsError = C.MMS_ERROR_ACCESS_OBJECT_VALUE_INVALID
	// MMS_ERROR_ACCESS_TEMPORARILY_UNAVAILABLE indicates the object is temporarily unavailable (for DataAccessError 2).
	MMS_ERROR_ACCESS_TEMPORARILY_UNAVAILABLE MmsError = C.MMS_ERROR_ACCESS_TEMPORARILY_UNAVAILABLE

	// MMS_ERROR_FILE_OTHER indicates an unspecified file error.
	MMS_ERROR_FILE_OTHER MmsError = C.MMS_ERROR_FILE_OTHER
	// MMS_ERROR_FILE_FILENAME_AMBIGUOUS indicates the filename is ambiguous.
	MMS_ERROR_FILE_FILENAME_AMBIGUOUS MmsError = C.MMS_ERROR_FILE_FILENAME_AMBIGUOUS
	// MMS_ERROR_FILE_FILE_BUSY indicates the file is busy.
	MMS_ERROR_FILE_FILE_BUSY MmsError = C.MMS_ERROR_FILE_FILE_BUSY
	// MMS_ERROR_FILE_FILENAME_SYNTAX_ERROR indicates a filename syntax error.
	MMS_ERROR_FILE_FILENAME_SYNTAX_ERROR MmsError = C.MMS_ERROR_FILE_FILENAME_SYNTAX_ERROR
	// MMS_ERROR_FILE_CONTENT_TYPE_INVALID indicates the content type is invalid.
	MMS_ERROR_FILE_CONTENT_TYPE_INVALID MmsError = C.MMS_ERROR_FILE_CONTENT_TYPE_INVALID
	// MMS_ERROR_FILE_POSITION_INVALID indicates the position is invalid.
	MMS_ERROR_FILE_POSITION_INVALID MmsError = C.MMS_ERROR_FILE_POSITION_INVALID
	// MMS_ERROR_FILE_FILE_ACCESS_DENIED indicates file access is denied.
	MMS_ERROR_FILE_FILE_ACCESS_DENIED MmsError = C.MMS_ERROR_FILE_FILE_ACCESS_DENIED
	// MMS_ERROR_FILE_FILE_NON_EXISTENT indicates the file does not exist.
	MMS_ERROR_FILE_FILE_NON_EXISTENT MmsError = C.MMS_ERROR_FILE_FILE_NON_EXISTENT
	// MMS_ERROR_FILE_DUPLICATE_FILENAME indicates a duplicate filename.
	MMS_ERROR_FILE_DUPLICATE_FILENAME MmsError = C.MMS_ERROR_FILE_DUPLICATE_FILENAME
	// MMS_ERROR_FILE_INSUFFICIENT_SPACE_IN_FILESTORE indicates insufficient space in filestore.
	MMS_ERROR_FILE_INSUFFICIENT_SPACE_IN_FILESTORE MmsError = C.MMS_ERROR_FILE_INSUFFICIENT_SPACE_IN_FILESTORE

	/* reject codes */
	// MMS_ERROR_REJECT_OTHER indicates an unspecified reject error.
	MMS_ERROR_REJECT_OTHER MmsError = C.MMS_ERROR_REJECT_OTHER
	// MMS_ERROR_REJECT_UNKNOWN_PDU_TYPE indicates an unknown PDU type.
	MMS_ERROR_REJECT_UNKNOWN_PDU_TYPE MmsError = C.MMS_ERROR_REJECT_UNKNOWN_PDU_TYPE
	// MMS_ERROR_REJECT_INVALID_PDU indicates an invalid PDU.
	MMS_ERROR_REJECT_INVALID_PDU MmsError = C.MMS_ERROR_REJECT_INVALID_PDU
	// MMS_ERROR_REJECT_UNRECOGNIZED_SERVICE indicates an unrecognized service.
	MMS_ERROR_REJECT_UNRECOGNIZED_SERVICE MmsError = C.MMS_ERROR_REJECT_UNRECOGNIZED_SERVICE
	// MMS_ERROR_REJECT_UNRECOGNIZED_MODIFIER indicates an unrecognized modifier.
	MMS_ERROR_REJECT_UNRECOGNIZED_MODIFIER MmsError = C.MMS_ERROR_REJECT_UNRECOGNIZED_MODIFIER
	// MMS_ERROR_REJECT_REQUEST_INVALID_ARGUMENT indicates an invalid argument in the request.
	MMS_ERROR_REJECT_REQUEST_INVALID_ARGUMENT MmsError = C.MMS_ERROR_REJECT_REQUEST_INVALID_ARGUMENT
)

func (x MmsError) Error() error {
	if x == MMS_ERROR_NONE {
		return nil
	}
	return errors.New(x.String())
}

func (x MmsError) String() string {
	return C.GoString(C.MmsError_toString(C.MmsError(x)))
}

// MmsType represents MMS data types.
type MmsType int32

// MMS type constants.
const (
	// MMS_ARRAY represents all MMS array types (arrays contain uniform elements).
	MMS_ARRAY MmsType = C.MMS_ARRAY
	// MMS_STRUCTURE represents all complex MMS types (structures).
	MMS_STRUCTURE MmsType = C.MMS_STRUCTURE
	// MMS_BOOLEAN represents a boolean value.
	MMS_BOOLEAN MmsType = C.MMS_BOOLEAN
	// MMS_BIT_STRING represents a bit string.
	MMS_BIT_STRING MmsType = C.MMS_BIT_STRING
	// MMS_INTEGER represents all signed integer types.
	MMS_INTEGER MmsType = C.MMS_INTEGER
	// MMS_UNSIGNED represents all unsigned integer types.
	MMS_UNSIGNED MmsType = C.MMS_UNSIGNED
	// MMS_FLOAT represents all float type (32 and 64 bit).
	MMS_FLOAT MmsType = C.MMS_FLOAT
	// MMS_OCTET_STRING represents an octet string (unstructured bytes).
	MMS_OCTET_STRING MmsType = C.MMS_OCTET_STRING
	// MMS_VISIBLE_STRING represents an MMS visible string.
	MMS_VISIBLE_STRING MmsType = C.MMS_VISIBLE_STRING
	// MMS_GENERALIZED_TIME represents an MMS generalized time type.
	MMS_GENERALIZED_TIME MmsType = C.MMS_GENERALIZED_TIME
	// MMS_BINARY_TIME represents an MMS binary time type.
	MMS_BINARY_TIME MmsType = C.MMS_BINARY_TIME
	// MMS_BCD represents an MMS BCD type.
	MMS_BCD MmsType = C.MMS_BCD
	// MMS_OBJ_ID represents an MMS object identifier type.
	MMS_OBJ_ID MmsType = C.MMS_OBJ_ID
	// MMS_STRING represents an MMS unicode string.
	MMS_STRING MmsType = C.MMS_STRING
	// MMS_UTC_TIME represents an MMS UTC time type.
	MMS_UTC_TIME MmsType = C.MMS_UTC_TIME
	// MMS_DATA_ACCESS_ERROR represents an error code as returned by MMS read services.
	MMS_DATA_ACCESS_ERROR MmsType = C.MMS_DATA_ACCESS_ERROR
)

// MmsDomain represents an MMS domain.
type MmsDomain struct {
	ctx *C.MmsDomain
}

// MmsAccessSpecifier specifies access to an MMS variable.
type MmsAccessSpecifier struct {
	ctx *C.MmsAccessSpecifier
}

// Domain returns the domain associated with the access specifier.
func (x *MmsAccessSpecifier) Domain() *MmsDomain {
	return &MmsDomain{ctx: x.ctx.domain}
}

// VariableName returns the variable name from the access specifier.
func (x *MmsAccessSpecifier) VariableName() string {
	return C.GoString(x.ctx.variableName)
}

// ArrayIndex returns the array index from the access specifier.
// Returns -1 if no index is present or the index should be ignored.
func (x *MmsAccessSpecifier) ArrayIndex() int {
	return int(x.ctx.arrayIndex)
}

// ComponentName returns the component name from the access specifier.
func (x *MmsAccessSpecifier) ComponentName() string {
	return C.GoString(x.ctx.componentName)
}

// MmsVariableAccessSpecification specifies access to an MMS variable.
type MmsVariableAccessSpecification struct {
	ctx *C.MmsVariableAccessSpecification
}

// DomainId returns the domain ID from the variable access specification.
func (x *MmsVariableAccessSpecification) DomainId() string {
	return C.GoString(x.ctx.domainId)
}

// ItemId returns the item ID from the variable access specification.
func (x *MmsVariableAccessSpecification) ItemId() string {
	return C.GoString(x.ctx.itemId)
}

// ArrayIndex returns the array index from the variable access specification.
// Returns -1 if no index is present or the index should be ignored.
func (x *MmsVariableAccessSpecification) ArrayIndex() int {
	return int(x.ctx.arrayIndex)
}

// ComponentName returns the component name from the variable access specification.
func (x *MmsVariableAccessSpecification) ComponentName() string {
	return C.GoString(x.ctx.componentName)
}

// MmsNamedVariableList represents a named variable list.
type MmsNamedVariableList struct {
	ctx C.MmsNamedVariableList
}

// MmsNamedVariableListEntry represents an entry in a named variable list.
type MmsNamedVariableListEntry struct {
	ctx C.MmsNamedVariableListEntry
}

// ItuObjectIdentifier represents an ITU (International Telecommunication Union) object identifier (OID).
type ItuObjectIdentifier struct {
	ctx *C.ItuObjectIdentifier
}

// ArcCount returns the count of arcs in the ITU object identifier.
func (x *ItuObjectIdentifier) ArcCount() int {
	return int(x.ctx.arcCount)
}

// Arc returns the arcs of the ITU object identifier as an array of 16 uint16 values.
func (x *ItuObjectIdentifier) Arc() [16]uint16 {
	arc := [16]uint16{}
	for i := range 16 {
		arc[i] = uint16(x.ctx.arc[i])
	}
	return arc
}

// IsoApplicationReference represents an ISO application reference
// (specifies an ISO application endpoint).
type IsoApplicationReference struct {
	ctx *C.IsoApplicationReference
}

// ApTitle returns the application title from the ISO application reference.
func (x *IsoApplicationReference) ApTitle() ItuObjectIdentifier {
	return ItuObjectIdentifier{ctx: &x.ctx.apTitle}
}

// AeQualifier returns the AE qualifier from the ISO application reference.
func (x *IsoApplicationReference) AeQualifier() int {
	return int(x.ctx.aeQualifier)
}
