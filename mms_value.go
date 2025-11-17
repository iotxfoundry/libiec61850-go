package libiec61850go

/*
#include <stdlib.h>
#include "mms_value.h"

*/
import "C"
import (
	"time"
	"unsafe"
)

type MmsValueIndication int32

const (
	MMS_VALUE_NO_RESPONSE               MmsValueIndication = C.MMS_VALUE_NO_RESPONSE
	MMS_VALUE_OK                        MmsValueIndication = C.MMS_VALUE_OK
	MMS_VALUE_ACCESS_DENIED             MmsValueIndication = C.MMS_VALUE_ACCESS_DENIED
	MMS_VALUE_VALUE_INVALID             MmsValueIndication = C.MMS_VALUE_VALUE_INVALID
	MMS_VALUE_TEMPORARILY_UNAVAILABLE   MmsValueIndication = C.MMS_VALUE_TEMPORARILY_UNAVAILABLE
	MMS_VALUE_OBJECT_ACCESS_UNSUPPORTED MmsValueIndication = C.MMS_VALUE_OBJECT_ACCESS_UNSUPPORTED
)

type MmsVariableSpecification struct {
	ctx *C.MmsVariableSpecification
}

type MmsDataAccessError int32

const (
	DATA_ACCESS_ERROR_SUCCESS_NO_UPDATE             MmsDataAccessError = C.DATA_ACCESS_ERROR_SUCCESS_NO_UPDATE
	DATA_ACCESS_ERROR_NO_RESPONSE                   MmsDataAccessError = C.DATA_ACCESS_ERROR_NO_RESPONSE
	DATA_ACCESS_ERROR_SUCCESS                       MmsDataAccessError = C.DATA_ACCESS_ERROR_SUCCESS
	DATA_ACCESS_ERROR_OBJECT_INVALIDATED            MmsDataAccessError = C.DATA_ACCESS_ERROR_OBJECT_INVALIDATED
	DATA_ACCESS_ERROR_HARDWARE_FAULT                MmsDataAccessError = C.DATA_ACCESS_ERROR_HARDWARE_FAULT
	DATA_ACCESS_ERROR_TEMPORARILY_UNAVAILABLE       MmsDataAccessError = C.DATA_ACCESS_ERROR_TEMPORARILY_UNAVAILABLE
	DATA_ACCESS_ERROR_OBJECT_ACCESS_DENIED          MmsDataAccessError = C.DATA_ACCESS_ERROR_OBJECT_ACCESS_DENIED
	DATA_ACCESS_ERROR_OBJECT_UNDEFINED              MmsDataAccessError = C.DATA_ACCESS_ERROR_OBJECT_UNDEFINED
	DATA_ACCESS_ERROR_INVALID_ADDRESS               MmsDataAccessError = C.DATA_ACCESS_ERROR_INVALID_ADDRESS
	DATA_ACCESS_ERROR_TYPE_UNSUPPORTED              MmsDataAccessError = C.DATA_ACCESS_ERROR_TYPE_UNSUPPORTED
	DATA_ACCESS_ERROR_TYPE_INCONSISTENT             MmsDataAccessError = C.DATA_ACCESS_ERROR_TYPE_INCONSISTENT
	DATA_ACCESS_ERROR_OBJECT_ATTRIBUTE_INCONSISTENT MmsDataAccessError = C.DATA_ACCESS_ERROR_OBJECT_ATTRIBUTE_INCONSISTENT
	DATA_ACCESS_ERROR_OBJECT_ACCESS_UNSUPPORTED     MmsDataAccessError = C.DATA_ACCESS_ERROR_OBJECT_ACCESS_UNSUPPORTED
	DATA_ACCESS_ERROR_OBJECT_NONE_EXISTENT          MmsDataAccessError = C.DATA_ACCESS_ERROR_OBJECT_NONE_EXISTENT
	DATA_ACCESS_ERROR_OBJECT_VALUE_INVALID          MmsDataAccessError = C.DATA_ACCESS_ERROR_OBJECT_VALUE_INVALID
	DATA_ACCESS_ERROR_UNKNOWN                       MmsDataAccessError = C.DATA_ACCESS_ERROR_UNKNOWN
)

type MmsValue struct {
	ctx *C.MmsValue
}

func MmsValueCreateArray(elementType *MmsVariableSpecification, size int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_createArray(elementType.ctx, C.int(size))}
}

func MmsValueCreateEmptyArray(size int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_createEmptyArray(C.int(size))}
}

func (x *MmsValue) GetArraySize() uint32 {
	return uint32(C.MmsValue_getArraySize(x.ctx))
}

func (x *MmsValue) GetElement(index uint32) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_getElement(x.ctx, C.int(index))}
}

func (x *MmsValue) SetElement(index uint32, element *MmsValue) {
	C.MmsValue_setElement(x.ctx, C.int(index), element.ctx)
}

func (x *MmsValue) GetDataAccessError() MmsDataAccessError {
	return MmsDataAccessError(C.MmsValue_getDataAccessError(x.ctx))
}

func (x *MmsValue) ToInt64() int64 {
	return int64(C.MmsValue_toInt64(x.ctx))
}

func (x *MmsValue) ToInt32() int32 {
	return int32(C.MmsValue_toInt32(x.ctx))
}

func (x *MmsValue) ToUint32() uint32 {
	return uint32(C.MmsValue_toUint32(x.ctx))
}

func (x *MmsValue) ToDouble() float64 {
	return float64(C.MmsValue_toDouble(x.ctx))
}

func (x *MmsValue) ToFloat() float32 {
	return float32(C.MmsValue_toFloat(x.ctx))
}

func (x *MmsValue) ToUnixTimestamp() uint32 {
	return uint32(C.MmsValue_toUnixTimestamp(x.ctx))
}

func (x *MmsValue) SetFloat(value float32) {
	C.MmsValue_setFloat(x.ctx, C.float(value))
}

func (x *MmsValue) SetDouble(value float64) {
	C.MmsValue_setDouble(x.ctx, C.double(value))
}

func (x *MmsValue) SetInt8(value int8) {
	C.MmsValue_setInt8(x.ctx, C.int8_t(value))
}

func (x *MmsValue) SetInt16(value int16) {
	C.MmsValue_setInt16(x.ctx, C.int16_t(value))
}

func (x *MmsValue) SetInt32(value int32) {
	C.MmsValue_setInt32(x.ctx, C.int32_t(value))
}

func (x *MmsValue) SetInt64(value int64) {
	C.MmsValue_setInt64(x.ctx, C.int64_t(value))
}

func (x *MmsValue) SetUint8(value uint8) {
	C.MmsValue_setUint8(x.ctx, C.uint8_t(value))
}

func (x *MmsValue) SetUint16(value uint16) {
	C.MmsValue_setUint16(x.ctx, C.uint16_t(value))
}

func (x *MmsValue) SetUint32(value uint32) {
	C.MmsValue_setUint32(x.ctx, C.uint32_t(value))
}

func (x *MmsValue) SetBoolean(value bool) {
	C.MmsValue_setBoolean(x.ctx, C.bool(value))
}

func (x *MmsValue) GetBoolean() bool {
	return bool(C.MmsValue_getBoolean(x.ctx))
}

func (x *MmsValue) ToString() string {
	return C.GoString(C.MmsValue_toString(x.ctx))
}

func (x *MmsValue) GetStringSize() int {
	return int(C.MmsValue_getStringSize(x.ctx))
}

func (x *MmsValue) SetVisibleString(value string) {
	cval := C.CString(value)
	defer C.free(unsafe.Pointer(cval))
	C.MmsValue_setVisibleString(x.ctx, cval)
}

func (x *MmsValue) SetBitStringBit(index uint32, value bool) {
	C.MmsValue_setBitStringBit(x.ctx, C.int(index), C.bool(value))
}

func (x *MmsValue) GetBitStringBit(index uint32) bool {
	return bool(C.MmsValue_getBitStringBit(x.ctx, C.int(index)))
}

func (x *MmsValue) DeleteAllBitStringBits() {
	C.MmsValue_deleteAllBitStringBits(x.ctx)
}

func (x *MmsValue) GetBitStringSize() int {
	return int(C.MmsValue_getBitStringSize(x.ctx))
}

func (x *MmsValue) GetBitStringByteSize() int {
	return int(C.MmsValue_getBitStringByteSize(x.ctx))
}

func (x *MmsValue) GetNumberOfSetBits() int {
	return int(C.MmsValue_getNumberOfSetBits(x.ctx))
}

func (x *MmsValue) SetAllBitStringBits() {
	C.MmsValue_setAllBitStringBits(x.ctx)
}

func (x *MmsValue) GetBitStringAsInteger() uint32 {
	return uint32(C.MmsValue_getBitStringAsInteger(x.ctx))
}

func (x *MmsValue) SetBitStringFromInteger(value uint32) {
	C.MmsValue_setBitStringFromInteger(x.ctx, C.uint32_t(value))
}

func (x *MmsValue) GetBitStringAsIntegerBigEndian() uint32 {
	return uint32(C.MmsValue_getBitStringAsIntegerBigEndian(x.ctx))
}

func (x *MmsValue) SetBitStringFromIntegerBigEndian(value uint32) {
	C.MmsValue_setBitStringFromIntegerBigEndian(x.ctx, C.uint32_t(value))
}

func (x *MmsValue) SetUtcTime(timeval time.Time) *MmsValue {
	x.ctx = C.MmsValue_setUtcTime(x.ctx, C.uint32_t(timeval.Unix()))
	return &MmsValue{ctx: x.ctx}
}

func (x *MmsValue) SetUtcTimeMs(timeval time.Time) *MmsValue {
	x.ctx = C.MmsValue_setUtcTimeMs(x.ctx, C.uint64_t(timeval.UnixNano()/1e6))
	return &MmsValue{ctx: x.ctx}
}

func (x *MmsValue) SetUtcTimeByBuffer(timeval [8]byte) {
	if len(timeval) != 8 {
		panic("timeval must be 8 bytes long")
	}
	ctimeval := C.CBytes(timeval[:])
	defer C.free(ctimeval)
	C.MmsValue_setUtcTimeByBuffer(x.ctx, (*C.uint8_t)(ctimeval))
}

func (x *MmsValue) GetUtcTimeBuffer() [8]byte {
	buf := C.MmsValue_getUtcTimeBuffer(x.ctx)
	return *(*[8]byte)(unsafe.Pointer(buf))
}

func (x *MmsValue) GetUtcTimeInMs() uint64 {
	return uint64(C.MmsValue_getUtcTimeInMs(x.ctx))
}

func (x *MmsValue) GetUtcTimeInMsWithUs(usec *uint32) uint64 {
	out := C.MmsValue_getUtcTimeInMsWithUs(x.ctx, (*C.uint32_t)(unsafe.Pointer(usec)))
	return uint64(out)
}

func (x *MmsValue) SetUtcTimeQuality(timeQuality byte) {
	C.MmsValue_setUtcTimeQuality(x.ctx, C.uint8_t(timeQuality))
}

func (x *MmsValue) SetUtcTimeMsEx(timeval time.Time, timeQuality uint8) *MmsValue {
	x.ctx = C.MmsValue_setUtcTimeMsEx(x.ctx, C.uint64_t(timeval.UnixNano()/1e6), C.uint8_t(timeQuality))
	return &MmsValue{ctx: x.ctx}
}

func (x *MmsValue) GetUtcTimeQuality() uint8 {
	return uint8(C.MmsValue_getUtcTimeQuality(x.ctx))
}

func (x *MmsValue) SetBinaryTime(timestamp uint64) {
	C.MmsValue_setBinaryTime(x.ctx, C.uint64_t(timestamp))
}

func (x *MmsValue) GetBinaryTimeAsUtcMs() uint64 {
	return uint64(C.MmsValue_getBinaryTimeAsUtcMs(x.ctx))
}

func (x *MmsValue) SetOctetString(value []byte) {
	cval := C.CBytes(value)
	defer C.free(unsafe.Pointer(cval))
	C.MmsValue_setOctetString(x.ctx, (*C.uint8_t)(cval), C.int(len(value)))
}

func (x *MmsValue) SetOctetStringOctet(octetPos int, value uint8) {
	C.MmsValue_setOctetStringOctet(x.ctx, C.int(octetPos), C.uint8_t(value))
}

func (x *MmsValue) GetOctetStringSize() uint16 {
	return uint16(C.MmsValue_getOctetStringSize(x.ctx))
}

func (x *MmsValue) GetOctetStringMaxSize() uint16 {
	return uint16(C.MmsValue_getOctetStringMaxSize(x.ctx))
}

func (x *MmsValue) GetOctetStringBuffer() []byte {
	buf := C.MmsValue_getOctetStringBuffer(x.ctx)
	return C.GoBytes(unsafe.Pointer(buf), C.int(C.MmsValue_getOctetStringSize(x.ctx)))
}

func (x *MmsValue) GetOctetStringOctet(octetPos int) uint8 {
	return uint8(C.MmsValue_getOctetStringOctet(x.ctx, C.int(octetPos)))
}

func (x *MmsValue) Update(source *MmsValue) bool {
	return bool(C.MmsValue_update(x.ctx, source.ctx))
}

func (x *MmsValue) Equals(otherValue *MmsValue) bool {
	return bool(C.MmsValue_equals(x.ctx, otherValue.ctx))
}

func (x *MmsValue) EqualTypes(otherValue *MmsValue) bool {
	return bool(C.MmsValue_equalTypes(x.ctx, otherValue.ctx))
}

func MmsValueNewDataAccessError(errorCode MmsDataAccessError) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newDataAccessError(C.MmsDataAccessError(errorCode))}
}

func MmsValueNewInteger(size int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newInteger(C.int(size))}
}

func MmsValueNewUnsigned(size int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newUnsigned(C.int(size))}
}

func MmsValueNewBoolean(boolean bool) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newBoolean(C.bool(boolean))}
}

func MmsValueNewBitString(bitSize int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newBitString(C.int(bitSize))}
}

func MmsValueNewOctetString(size int, maxSize int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newOctetString(C.int(size), C.int(maxSize))}
}

func MmsValueNewStructure(typeSpec *MmsVariableSpecification) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newStructure(typeSpec.ctx)}
}

func MmsValueCreateEmptyStructure(size int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_createEmptyStructure(C.int(size))}
}

func MmsValueNewDefaultValue(typeSpec *MmsVariableSpecification) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newDefaultValue(typeSpec.ctx)}
}

func MmsValueNewIntegerFromInt8(integer int8) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newIntegerFromInt8(C.int8_t(integer))}
}

func MmsValueNewIntegerFromInt16(integer int16) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newIntegerFromInt16(C.int16_t(integer))}
}

func MmsValueNewIntegerFromInt32(integer int32) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newIntegerFromInt32(C.int32_t(integer))}
}

func MmsValueNewIntegerFromInt64(integer int64) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newIntegerFromInt64(C.int64_t(integer))}
}

func MmsValueNewUnsignedFromUint32(integer uint32) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newUnsignedFromUint32(C.uint32_t(integer))}
}

func MmsValueNewFloat32(floatVal float32) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newFloat(C.float(floatVal))}
}

func MmsValueNewFloat64(floatVal float64) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newDouble(C.double(floatVal))}
}

func (x *MmsValue) Clone() *MmsValue {
	return &MmsValue{ctx: C.MmsValue_clone(x.ctx)}
}

func (x *MmsValue) CloneToBuffer(destinationAddress []byte) []byte {
	C.MmsValue_cloneToBuffer(x.ctx, (*C.uint8_t)(unsafe.SliceData(destinationAddress)))
	return destinationAddress
}

func (x *MmsValue) GetSizeInMemory() int {
	return int(C.MmsValue_getSizeInMemory(x.ctx))
}

func (x *MmsValue) Delete() {
	C.MmsValue_delete(x.ctx)
}

func (x *MmsValue) DeleteConditional() {
	C.MmsValue_deleteConditional(x.ctx)
}

func MmsValueNewVisibleString(value string) *MmsValue {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	return &MmsValue{ctx: C.MmsValue_newVisibleString(cstr)}
}

func MmsValueNewVisibleStringWithSize(size int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newVisibleStringWithSize(C.int(size))}
}

func MmsValueNewMmsStringWithSize(size int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newMmsStringWithSize(C.int(size))}
}

func MmsValueNewBinaryTime(timeOfDay bool) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newBinaryTime(C.bool(timeOfDay))}
}

func MmsValueNewVisibleStringFromByteArray(byteArray []byte) *MmsValue {
	cbytes := C.CBytes(byteArray)
	defer C.free(unsafe.Pointer(cbytes))
	return &MmsValue{ctx: C.MmsValue_newVisibleStringFromByteArray((*C.uint8_t)(cbytes), C.int(len(byteArray)))}
}

func MmsValueNewMmsStringFromByteArray(byteArray []byte) *MmsValue {
	cbytes := C.CBytes(byteArray)
	defer C.free(unsafe.Pointer(cbytes))
	return &MmsValue{ctx: C.MmsValue_newMmsStringFromByteArray((*C.uint8_t)(cbytes), C.int(len(byteArray)))}
}

func MmsValueNewMmsString(str string) *MmsValue {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return &MmsValue{ctx: C.MmsValue_newMmsString(cstr)}
}

func (x *MmsValue) SetMmsString(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.MmsValue_setMmsString(x.ctx, cstr)
}

func MmsValueNewUtcTime(timeval uint32) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newUtcTime(C.uint32_t(timeval))}
}

func MmsValueNewUtcTimeByMsTime(timeval uint64) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_newUtcTimeByMsTime(C.uint64_t(timeval))}
}

func (x *MmsValue) SetDeletable() {
	C.MmsValue_setDeletable(x.ctx)
}

func (x *MmsValue) SetDeletableRecursive() {
	C.MmsValue_setDeletableRecursive(x.ctx)
}

func (x *MmsValue) IsDeletable() bool {
	return C.MmsValue_isDeletable(x.ctx) != 0
}

func (x *MmsValue) GetType() MmsType {
	return MmsType(C.MmsValue_getType(x.ctx))
}

func (x *MmsValue) GetSubElement(varSpec *MmsVariableSpecification, mmsPath string) *MmsValue {
	cstr := C.CString(mmsPath)
	defer C.free(unsafe.Pointer(cstr))
	return &MmsValue{ctx: C.MmsValue_getSubElement(x.ctx, varSpec.ctx, cstr)}
}

func (x *MmsValue) GetTypeString() string {
	return C.GoString(C.MmsValue_getTypeString(x.ctx))
}

func (x *MmsValue) PrintToBuffer(buffer string) string {
	C.MmsValue_printToBuffer(x.ctx, (*C.char)(unsafe.Pointer(unsafe.StringData(buffer))), C.int(len(buffer)))
	return buffer
}

func MmsValueDecodeMmsData(buffer []byte, bufPos int, endBufPos *int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_decodeMmsData((*C.uint8_t)(unsafe.SliceData(buffer)), C.int(bufPos), C.int(len(buffer)), (*C.int)(unsafe.Pointer(endBufPos)))}
}

func MmsValueDecodeMmsDataMaxRecursion(buffer []byte, bufPos int, endBufPos *int, maxDepth int) *MmsValue {
	return &MmsValue{ctx: C.MmsValue_decodeMmsDataMaxRecursion((*C.uint8_t)(unsafe.SliceData(buffer)), C.int(bufPos), C.int(len(buffer)), (*C.int)(unsafe.Pointer(endBufPos)), C.int(maxDepth))}
}

func (x *MmsValue) EncodeMmsData(buffer []byte, bufPos int, encode bool) int {
	return int(C.MmsValue_encodeMmsData(x.ctx, (*C.uint8_t)(unsafe.SliceData(buffer)), C.int(bufPos), C.bool(encode)))
}

func (x *MmsValue) GetMaxEncodedSize() int {
	return int(C.MmsValue_getMaxEncodedSize(x.ctx))
}

func (x *MmsVariableSpecification) GetMaxEncodedSize() int {
	return int(C.MmsVariableSpecification_getMaxEncodedSize(x.ctx))
}

func (x *MmsError) String() string {
	return C.GoString(C.MmsError_toString(C.MmsError(x.ctx)))
}
