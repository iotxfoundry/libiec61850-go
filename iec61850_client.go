package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_client.h"

extern void fIedConnectionClosedHandlerGo(void* parameter, IedConnection connection);
extern void fStateChangedHandlerGo(void* parameter, IedConnection connection, IedConnectionState newState);
extern void fGetGoCBValuesHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, ClientGooseControlBlock goCB);
extern void fReadObjectHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, MmsValue* value);
extern void fGetRCBValuesHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, ClientReportControlBlock rcb);
extern void fGenericServiceHandlerGo(uint32_t invokeId, void* parameter, IedClientError err);
extern void fReportCallbackFunctionGo(void* parameter, ClientReport report);
extern void fReadDataSetHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, ClientDataSet dataSet);
extern void fGetDataSetDirectoryHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, LinkedList dataSetDirectory, bool isDeletable);
extern void fWriteDataSetHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, LinkedList accessResults);
extern void fControlActionHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, ControlActionType type_, bool success);
extern void fCommandTerminationHandlerGo(void* parameter, ControlObjectClient controlClient);
extern void fGetNameListHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, LinkedList nameList, bool moreFollows);
extern void fGetVariableSpecificationHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, MmsVariableSpecification* variableSpecification);
extern void fQueryLogHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, LinkedList journalEntries, bool moreFollows);
extern void fFileDirectoryEntryHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, char* filename, uint32_t size, uint64_t lastModfified, bool moreFollows);
extern bool fIedClientGetFileHandlerGo(void* parameter, uint8_t* buffer, uint32_t bytesRead);
extern bool fGetFileAsyncHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, uint32_t originalInvokeId, uint8_t* buffer, uint32_t bytesRead, bool moreFollows);
*/
import "C"
import (
	"sync"
	"unsafe"
)

type ClientDataSet struct {
	ctx C.ClientDataSet
}

type ClientReport struct {
	ctx C.ClientReport
}

type ClientReportControlBlock struct {
	ctx C.ClientReportControlBlock
}

type ClientGooseControlBlock struct {
	ctx C.ClientGooseControlBlock
}

type IedConnection struct {
	ctx C.IedConnection
}

type LastApplError struct {
	ctx *C.LastApplError
}

func (x *LastApplError) CtlNum() int {
	return int(x.ctx.ctlNum)
}

func (x *LastApplError) Error() ControlLastApplError {
	return ControlLastApplError(x.ctx.error)
}

func (x *LastApplError) AddCause() ControlAddCause {
	return ControlAddCause(x.ctx.addCause)
}

type IedConnectionState int32

const (
	IED_STATE_CLOSED     IedConnectionState = C.IED_STATE_CLOSED
	IED_STATE_CONNECTING IedConnectionState = C.IED_STATE_CONNECTING
	IED_STATE_CONNECTED  IedConnectionState = C.IED_STATE_CONNECTED
	IED_STATE_CLOSING    IedConnectionState = C.IED_STATE_CLOSING
)

type IedClientError int32

const (
	/* general errors */

	/** No error occurred - service request has been successful */
	IED_ERROR_OK IedClientError = C.IED_ERROR_OK

	/** The service request can not be executed because the client is not yet connected */
	IED_ERROR_NOT_CONNECTED IedClientError = C.IED_ERROR_NOT_CONNECTED

	/** Connect service not execute because the client is already connected */
	IED_ERROR_ALREADY_CONNECTED IedClientError = C.IED_ERROR_ALREADY_CONNECTED

	/** The service request can not be executed caused by a loss of connection */
	IED_ERROR_CONNECTION_LOST IedClientError = C.IED_ERROR_CONNECTION_LOST

	/** The service or some given parameters are not supported by the client stack or by the server */
	IED_ERROR_SERVICE_NOT_SUPPORTED IedClientError = C.IED_ERROR_SERVICE_NOT_SUPPORTED

	/** Connection rejected by server */
	IED_ERROR_CONNECTION_REJECTED IedClientError = C.IED_ERROR_CONNECTION_REJECTED

	/** Cannot send request because outstanding call limit is reached */
	IED_ERROR_OUTSTANDING_CALL_LIMIT_REACHED IedClientError = C.IED_ERROR_OUTSTANDING_CALL_LIMIT_REACHED

	/* client side errors */

	/** API function has been called with an invalid argument */
	IED_ERROR_USER_PROVIDED_INVALID_ARGUMENT IedClientError = C.IED_ERROR_USER_PROVIDED_INVALID_ARGUMENT

	/** The server reported that the requested report control block is not enabled */
	IED_ERROR_ENABLE_REPORT_FAILED_DATASET_MISMATCH IedClientError = C.IED_ERROR_ENABLE_REPORT_FAILED_DATASET_MISMATCH

	/** The object provided object reference is invalid (there is a syntactical error). */
	IED_ERROR_OBJECT_REFERENCE_INVALID IedClientError = C.IED_ERROR_OBJECT_REFERENCE_INVALID

	/** Received object is of unexpected type */
	IED_ERROR_UNEXPECTED_VALUE_RECEIVED IedClientError = C.IED_ERROR_UNEXPECTED_VALUE_RECEIVED

	/* service error - error reported by server */

	/** The communication to the server failed with a timeout */
	IED_ERROR_TIMEOUT IedClientError = C.IED_ERROR_TIMEOUT

	/** The server rejected the access to the requested object/service due to access control */
	IED_ERROR_ACCESS_DENIED IedClientError = C.IED_ERROR_ACCESS_DENIED

	/** The server reported that the requested object does not exist (returned by server) */
	IED_ERROR_OBJECT_DOES_NOT_EXIST IedClientError = C.IED_ERROR_OBJECT_DOES_NOT_EXIST

	/** The server reported that the requested object already exists */
	IED_ERROR_OBJECT_EXISTS IedClientError = C.IED_ERROR_OBJECT_EXISTS

	/** The server does not support the requested access method (returned by server) */
	IED_ERROR_OBJECT_ACCESS_UNSUPPORTED IedClientError = C.IED_ERROR_OBJECT_ACCESS_UNSUPPORTED

	/** The server expected an object of another type (returned by server) */
	IED_ERROR_TYPE_INCONSISTENT IedClientError = C.IED_ERROR_TYPE_INCONSISTENT

	/** The object or service is temporarily unavailable (returned by server) */
	IED_ERROR_TEMPORARILY_UNAVAILABLE IedClientError = C.IED_ERROR_TEMPORARILY_UNAVAILABLE

	/** The specified object is not defined in the server (returned by server) */
	IED_ERROR_OBJECT_UNDEFINED IedClientError = C.IED_ERROR_OBJECT_UNDEFINED

	/** The specified address is invalid (returned by server) */
	IED_ERROR_INVALID_ADDRESS IedClientError = C.IED_ERROR_INVALID_ADDRESS

	/** Service failed due to a hardware fault (returned by server) */
	IED_ERROR_HARDWARE_FAULT IedClientError = C.IED_ERROR_HARDWARE_FAULT

	/** The requested data type is not supported by the server (returned by server) */
	IED_ERROR_TYPE_UNSUPPORTED IedClientError = C.IED_ERROR_TYPE_UNSUPPORTED

	/** The provided attributes are inconsistent (returned by server) */
	IED_ERROR_OBJECT_ATTRIBUTE_INCONSISTENT IedClientError = C.IED_ERROR_OBJECT_ATTRIBUTE_INCONSISTENT

	/** The provided object value is invalid (returned by server) */
	IED_ERROR_OBJECT_VALUE_INVALID IedClientError = C.IED_ERROR_OBJECT_VALUE_INVALID

	/** The object is invalidated (returned by server) */
	IED_ERROR_OBJECT_INVALIDATED IedClientError = C.IED_ERROR_OBJECT_INVALIDATED

	/** Received an invalid response message from the server */
	IED_ERROR_MALFORMED_MESSAGE IedClientError = C.IED_ERROR_MALFORMED_MESSAGE

	/** Service was not executed because required resource is still in use */
	IED_ERROR_OBJECT_CONSTRAINT_CONFLICT IedClientError = C.IED_ERROR_OBJECT_CONSTRAINT_CONFLICT

	/** Service not implemented */
	IED_ERROR_SERVICE_NOT_IMPLEMENTED IedClientError = C.IED_ERROR_SERVICE_NOT_IMPLEMENTED

	/** unknown error */
	IED_ERROR_UNKNOWN IedClientError = C.IED_ERROR_UNKNOWN
)

func (e IedClientError) Error() string {
	return e.String()
}

func (x IedClientError) String() string {
	return C.GoString(C.IedClientError_toString(C.IedClientError(x)))
}

func IedConnectionCreate() *IedConnection {
	return &IedConnection{
		ctx: C.IedConnection_create(),
	}
}

// func IedConnectionCreateEx(tlsConfig *TLSConfiguration, useThreads bool) *IedConnection {
// 	return &IedConnection{
// 		ctx: C.IedConnection_createEx(tlsConfig.ctx, C.bool(useThreads)),
// 	}
// }

// func IedConnectionCreateWithTlsSupport(tlsConfig *TLSConfiguration) *IedConnection {
// 	return &IedConnection{
// 		ctx: C.IedConnection_createWithTlsSupport(tlsConfig.ctx),
// 	}
// }

func (x *IedConnection) Destroy() {
	C.IedConnection_destroy(x.ctx)
}

func (x *IedConnection) SetLocalAddress(localIpAddress string, localPort int) {
	caddr := C.CString(localIpAddress)
	defer C.free(unsafe.Pointer(caddr))
	C.IedConnection_setLocalAddress(x.ctx, caddr, C.int(localPort))
}

func (x *IedConnection) SetConnectTimeout(timeoutInMs uint32) {
	C.IedConnection_setConnectTimeout(x.ctx, C.uint32_t(timeoutInMs))
}

func (x *IedConnection) SetMaxOutstandingCalls(calling int, called int) {
	C.IedConnection_setMaxOutstandingCalls(x.ctx, C.int(calling), C.int(called))
}

func (x *IedConnection) SetRequestTimeout(timeoutInMs uint32) {
	C.IedConnection_setRequestTimeout(x.ctx, C.uint32_t(timeoutInMs))
}

func (x *IedConnection) GetRequestTimeout() uint32 {
	return uint32(C.IedConnection_getRequestTimeout(x.ctx))
}

func (x *IedConnection) SetTimeQuality(leapSecondKnown bool, clockFailure bool, clockNotSynchronized bool, subsecondPrecision int) {
	C.IedConnection_setTimeQuality(x.ctx, C.bool(leapSecondKnown), C.bool(clockFailure), C.bool(clockNotSynchronized), C.int(subsecondPrecision))
}

func (x *IedConnection) Tick() bool {
	return bool(C.IedConnection_tick(x.ctx))
}

func (x *IedConnection) Connect(hostname string, tcpPort int) error {
	caddr := C.CString(hostname)
	defer C.free(unsafe.Pointer(caddr))
	err := IED_ERROR_OK
	C.IedConnection_connect(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), caddr, C.int(tcpPort))
	return err
}

func (x *IedConnection) ConnectAsync(hostname string, tcpPort int) error {
	caddr := C.CString(hostname)
	defer C.free(unsafe.Pointer(caddr))
	err := IED_ERROR_OK
	C.IedConnection_connectAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), caddr, C.int(tcpPort))
	return err
}

func (x *IedConnection) Abort() error {
	err := IED_ERROR_OK
	C.IedConnection_abort(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)))
	return err
}

func (x *IedConnection) AbortAsync() error {
	err := IED_ERROR_OK
	C.IedConnection_abortAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)))
	return err
}

func (x *IedConnection) Release() error {
	err := IED_ERROR_OK
	C.IedConnection_release(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)))
	return err
}

func (x *IedConnection) ReleaseAsync() error {
	err := IED_ERROR_OK
	C.IedConnection_releaseAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)))
	return err
}

func (x *IedConnection) Close() {
	C.IedConnection_close(x.ctx)
}

func (x *IedConnection) GetState() IedConnectionState {
	return IedConnectionState(C.IedConnection_getState(x.ctx))
}

func (x *IedConnection) GetLastApplError() *LastApplError {
	ctx := C.IedConnection_getLastApplError(x.ctx)
	return &LastApplError{
		ctx: &ctx,
	}
}

type IedConnectionClosedHandler func(parameter any, x *IedConnection)

var mapIedConnectionClosedHandlers = sync.Map{}

//export fIedConnectionClosedHandlerGo
func fIedConnectionClosedHandlerGo(parameter unsafe.Pointer, connection C.IedConnection) {
	mapIedConnectionClosedHandlers.Range(func(key, value any) bool {
		if key.(C.IedConnection) == connection {
			if fn, ok := value.(IedConnectionClosedHandler); ok {
				fn(parameter, &IedConnection{ctx: connection})
			}
		}
		return true
	})
}

func (x *IedConnection) InstallConnectionClosedHandler(handler IedConnectionClosedHandler, parameter unsafe.Pointer) {
	mapIedConnectionClosedHandlers.Store(x.ctx, handler)
	C.IedConnection_installConnectionClosedHandler(x.ctx, (C.IedConnectionClosedHandler)(C.fIedConnectionClosedHandlerGo), parameter)
}

type StateChangedHandler func(parameter any, connection *IedConnection, newState IedConnectionState)

var mapStateChangedHandlers = sync.Map{}

//export fStateChangedHandlerGo
func fStateChangedHandlerGo(parameter unsafe.Pointer, connection C.IedConnection, newState C.IedConnectionState) {
	mapStateChangedHandlers.Range(func(key, value any) bool {
		if key.(C.IedConnection) == connection {
			if fn, ok := value.(StateChangedHandler); ok {
				fn(parameter, &IedConnection{ctx: connection}, IedConnectionState(newState))
			}
		}
		return true
	})
}

func (x *IedConnection) InstallStateChangedHandler(handler StateChangedHandler, parameter unsafe.Pointer) {
	mapStateChangedHandlers.Store(x.ctx, handler)
	C.IedConnection_installStateChangedHandler(x.ctx, (C.IedConnection_StateChangedHandler)(C.fStateChangedHandlerGo), parameter)
}

func (x *IedConnection) GetMmsConnection() *MmsConnection {
	return &MmsConnection{
		ctx: C.IedConnection_getMmsConnection(x.ctx),
	}
}

const (
	/** SV ASDU contains attribute RefrTm */
	IEC61850_SV_OPT_REFRESH_TIME = C.IEC61850_SV_OPT_REFRESH_TIME

	/** SV ASDU contains attribute SmpSynch */
	IEC61850_SV_OPT_SAMPLE_SYNC = C.IEC61850_SV_OPT_SAMPLE_SYNC

	/** SV ASDU contains attribute SmpRate */
	IEC61850_SV_OPT_SAMPLE_RATE = C.IEC61850_SV_OPT_SAMPLE_RATE

	/** SV ASDU contains attribute DatSet */
	IEC61850_SV_OPT_DATA_SET = C.IEC61850_SV_OPT_DATA_SET

	/** SV ASDU contains attribute Security */
	IEC61850_SV_OPT_SECURITY = C.IEC61850_SV_OPT_SECURITY

	/** SV sampling mode: samples per period */
	IEC61850_SV_SMPMOD_SAMPLES_PER_PERIOD = C.IEC61850_SV_SMPMOD_SAMPLES_PER_PERIOD

	/** SV sampling mode: samples per second */
	IEC61850_SV_SMPMOD_SAMPLES_PER_SECOND = C.IEC61850_SV_SMPMOD_SAMPLES_PER_SECOND

	/** SV sampling mode: seconds per sample */
	IEC61850_SV_SMPMOD_SECONDS_PER_SAMPLE = C.IEC61850_SV_SMPMOD_SECONDS_PER_SAMPLE
)

type ClientSVControlBlock struct {
	ctx C.ClientSVControlBlock
}

func ClientSvControlBlockCreate(connection *IedConnection, reference string) *ClientSVControlBlock {
	cref := C.CString(reference)
	defer C.free(unsafe.Pointer(cref))
	return &ClientSVControlBlock{
		ctx: C.ClientSVControlBlock_create(connection.ctx, cref),
	}
}

func (x *ClientSVControlBlock) Destroy() {
	C.ClientSVControlBlock_destroy(x.ctx)
}

func (x *ClientSVControlBlock) IsMulticast() bool {
	return (bool)(C.ClientSVControlBlock_isMulticast(x.ctx))
}

func (x *ClientSVControlBlock) GetLastComError() error {
	err := C.ClientSVControlBlock_getLastComError(x.ctx)
	return IedClientError(err)
}

func (x *ClientSVControlBlock) SetSvEna(value bool) bool {
	return (bool)(C.ClientSVControlBlock_setSvEna(x.ctx, (C.bool)(value)))
}

func (x *ClientSVControlBlock) GetSvEna() bool {
	return (bool)(C.ClientSVControlBlock_getSvEna(x.ctx))
}

func (x *ClientSVControlBlock) SetResv(value bool) bool {
	return (bool)(C.ClientSVControlBlock_setResv(x.ctx, (C.bool)(value)))
}

func (x *ClientSVControlBlock) GetResv() bool {
	return (bool)(C.ClientSVControlBlock_getResv(x.ctx))
}

func (x *ClientSVControlBlock) GetMsvID() string {
	return C.GoString(C.ClientSVControlBlock_getMsvID(x.ctx))
}

func (x *ClientSVControlBlock) GetDatSet() string {
	return C.GoString(C.ClientSVControlBlock_getDatSet(x.ctx))
}

func (x *ClientSVControlBlock) GetConfRev() uint32 {
	return (uint32)(C.ClientSVControlBlock_getConfRev(x.ctx))
}

func (x *ClientSVControlBlock) GetSmpRate() uint16 {
	return (uint16)(C.ClientSVControlBlock_getSmpRate(x.ctx))
}

func (x *ClientSVControlBlock) GetDstAddress() *PhyComAddress {
	ctx := C.ClientSVControlBlock_getDstAddress(x.ctx)
	return &PhyComAddress{ctx: &ctx}
}

func (x *ClientSVControlBlock) GetOptFlds() int {
	return (int)(C.ClientSVControlBlock_getOptFlds(x.ctx))
}

func (x *ClientSVControlBlock) GetSmpMod() uint8 {
	return (uint8)(C.ClientSVControlBlock_getSmpMod(x.ctx))
}

func (x *ClientSVControlBlock) GetNoASDU() int {
	return (int)(C.ClientSVControlBlock_getNoASDU(x.ctx))
}

const (
	/** Enable GOOSE publisher GoCB block element */
	GOCB_ELEMENT_GO_ENA = C.GOCB_ELEMENT_GO_ENA

	/** GOOSE ID GoCB block element */
	GOCB_ELEMENT_GO_ID = C.GOCB_ELEMENT_GO_ID

	/** Data set GoCB block element */
	GOCB_ELEMENT_DATSET = C.GOCB_ELEMENT_DATSET

	/** Configuration revision GoCB block element (this is usually read-only) */
	GOCB_ELEMENT_CONF_REV = C.GOCB_ELEMENT_CONF_REV

	/** Need commission GoCB block element (read-only according to 61850-7-2) */
	GOCB_ELEMENT_NDS_COMM = C.GOCB_ELEMENT_NDS_COMM

	/** Destination address GoCB block element (read-only according to 61850-7-2) */
	GOCB_ELEMENT_DST_ADDRESS = C.GOCB_ELEMENT_DST_ADDRESS

	/** Minimum time GoCB block element (read-only according to 61850-7-2) */
	GOCB_ELEMENT_MIN_TIME = C.GOCB_ELEMENT_MIN_TIME

	/** Maximum time GoCB block element (read-only according to 61850-7-2) */
	GOCB_ELEMENT_MAX_TIME = C.GOCB_ELEMENT_MAX_TIME

	/** Fixed offsets GoCB block element (read-only according to 61850-7-2) */
	GOCB_ELEMENT_FIXED_OFFS = C.GOCB_ELEMENT_FIXED_OFFS

	/** select all elements of the GoCB */
	GOCB_ELEMENT_ALL = C.GOCB_ELEMENT_ALL
)

func ClientGooseControlBlockCreate(dataAttributeReference string) *ClientGooseControlBlock {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	return &ClientGooseControlBlock{
		ctx: C.ClientGooseControlBlock_create(cref),
	}
}

func (x *ClientGooseControlBlock) Destroy() {
	C.ClientGooseControlBlock_destroy(x.ctx)
}

func (x *ClientGooseControlBlock) GetGoEna() bool {
	return (bool)(C.ClientGooseControlBlock_getGoEna(x.ctx))
}

func (x *ClientGooseControlBlock) SetGoEna(value bool) {
	C.ClientGooseControlBlock_setGoEna(x.ctx, (C.bool)(value))
}

func (x *ClientGooseControlBlock) GetGoID() string {
	return C.GoString(C.ClientGooseControlBlock_getGoID(x.ctx))
}

func (x *ClientGooseControlBlock) SetGoID(value string) {
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	C.ClientGooseControlBlock_setGoID(x.ctx, cvalue)
}

func (x *ClientGooseControlBlock) GetDatSet() string {
	return C.GoString(C.ClientGooseControlBlock_getDatSet(x.ctx))
}

func (x *ClientGooseControlBlock) SetDatSet(value string) {
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	C.ClientGooseControlBlock_setDatSet(x.ctx, cvalue)
}

func (x *ClientGooseControlBlock) GetConfRev() uint32 {
	return (uint32)(C.ClientGooseControlBlock_getConfRev(x.ctx))
}

func (x *ClientGooseControlBlock) GetNdsComm() bool {
	return (bool)(C.ClientGooseControlBlock_getNdsComm(x.ctx))
}

func (x *ClientGooseControlBlock) GetMinTime() uint32 {
	return (uint32)(C.ClientGooseControlBlock_getMinTime(x.ctx))
}

func (x *ClientGooseControlBlock) GetMaxTime() uint32 {
	return (uint32)(C.ClientGooseControlBlock_getMaxTime(x.ctx))
}

func (x *ClientGooseControlBlock) GetFixedOffs() bool {
	return (bool)(C.ClientGooseControlBlock_getFixedOffs(x.ctx))
}

func (x *ClientGooseControlBlock) GetDstAddress() *PhyComAddress {
	ctx := C.ClientGooseControlBlock_getDstAddress(x.ctx)
	return &PhyComAddress{ctx: &ctx}
}

func (x *ClientGooseControlBlock) SetDstAddress(value *PhyComAddress) {
	C.ClientGooseControlBlock_setDstAddress(x.ctx, *value.ctx)
}

func (x *ClientGooseControlBlock) GetDstAddress_Addr() *MmsValue {
	ctx := C.ClientGooseControlBlock_getDstAddress_addr(x.ctx)
	return &MmsValue{ctx: ctx}
}

func (x *ClientGooseControlBlock) SetDstAddress_Addr(value *MmsValue) {
	C.ClientGooseControlBlock_setDstAddress_addr(x.ctx, value.ctx)
}

func (x *ClientGooseControlBlock) GetDstAddress_Priority() uint8 {
	return (uint8)(C.ClientGooseControlBlock_getDstAddress_priority(x.ctx))
}

func (x *ClientGooseControlBlock) SetDstAddress_Priority(value uint8) {
	C.ClientGooseControlBlock_setDstAddress_priority(x.ctx, (C.uint8_t)(value))
}

func (x *ClientGooseControlBlock) GetDstAddress_Vid() uint16 {
	return (uint16)(C.ClientGooseControlBlock_getDstAddress_vid(x.ctx))
}

func (x *ClientGooseControlBlock) SetDstAddress_Vid(value uint16) {
	C.ClientGooseControlBlock_setDstAddress_vid(x.ctx, (C.uint16_t)(value))
}

func (x *ClientGooseControlBlock) GetDstAddress_Appid() uint16 {
	return (uint16)(C.ClientGooseControlBlock_getDstAddress_appid(x.ctx))
}

func (x *ClientGooseControlBlock) SetDstAddress_Appid(value uint16) {
	C.ClientGooseControlBlock_setDstAddress_appid(x.ctx, (C.uint16_t)(value))
}

func (x *IedConnection) GetGoCBValues(goCBReference string, updateGoCB *ClientGooseControlBlock) (*ClientGooseControlBlock, error) {
	cref := C.CString(goCBReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_getGoCBValues(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, updateGoCB.ctx)
	goCB := &ClientGooseControlBlock{ctx: out}
	return goCB, IedClientError(err)
}

type GetGoCBValuesHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, goCB *ClientGooseControlBlock)

var mapGetGoCBValuesHandlers = sync.Map{}

//export fGetGoCBValuesHandlerGo
func fGetGoCBValuesHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, goCB C.ClientGooseControlBlock) {
	mapGetGoCBValuesHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(GetGoCBValuesHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &ClientGooseControlBlock{ctx: goCB})
		}
		return true
	})
}

func (x *IedConnection) GetGoCBValuesAsync(goCBReference string, updateGoCB *ClientGooseControlBlock, handler GetGoCBValuesHandler, parameter unsafe.Pointer) (uint32, error) {
	cref := C.CString(goCBReference)
	defer C.free(unsafe.Pointer(cref))
	mapGetGoCBValuesHandlers.Store(handler, handler)
	defer mapGetGoCBValuesHandlers.Delete(handler)
	err := IED_ERROR_OK
	out := C.IedConnection_getGoCBValuesAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, updateGoCB.ctx, (C.IedConnection_GetGoCBValuesHandler)(C.fGetGoCBValuesHandlerGo), parameter)
	return (uint32)(out), IedClientError(err)
}

func (x *IedConnection) SetGoCBValues(goCB *ClientGooseControlBlock, parametersMask uint32, singleRequest bool) error {
	err := IED_ERROR_OK
	C.IedConnection_setGoCBValues(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), goCB.ctx, (C.uint32_t)(parametersMask), (C.bool)(singleRequest))
	return IedClientError(err)
}

type GenericServiceHandler func(invokeId uint32, parameter any, err IedClientError)

var mapGenericServiceHandlers = sync.Map{}

//export fGenericServiceHandlerGo
func fGenericServiceHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError) {
	mapGenericServiceHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(GenericServiceHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err))
		}
		return true
	})
}

func (x *IedConnection) SetGoCBValuesAsync(goCB *ClientGooseControlBlock, parametersMask uint32, singleRequest bool, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := IED_ERROR_OK
	mapGenericServiceHandlers.Store(handler, handler)
	defer mapGenericServiceHandlers.Delete(handler)
	out := C.IedConnection_setGoCBValuesAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), goCB.ctx, (C.uint32_t)(parametersMask), (C.bool)(singleRequest), (C.IedConnection_GenericServiceHandler)(C.fGenericServiceHandlerGo), parameter)
	return (uint32)(out), IedClientError(err)
}

func (x *IedConnection) ReadObject(dataAttributeReference string, fc FunctionalConstraint) (*MmsValue, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_readObject(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc))
	mmsValue := &MmsValue{ctx: out}
	return mmsValue, IedClientError(err)
}

type ReadObjectHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, value *MmsValue)

var mapReadObjectHandlers = sync.Map{}

//export fReadObjectHandlerGo
func fReadObjectHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, value *C.MmsValue) {
	mapReadObjectHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(ReadObjectHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &MmsValue{ctx: value.(*C.MmsValue)})
		}
		return true
	})
}

func (x *IedConnection) ReadObjectAsync(dataAttributeReference string, fc FunctionalConstraint, handler ReadObjectHandler, parameter unsafe.Pointer) (uint32, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	mapReadObjectHandlers.Store(handler, handler)
	defer mapReadObjectHandlers.Delete(handler)
	out := C.IedConnection_readObjectAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), (C.IedConnection_ReadObjectHandler)(C.fReadObjectHandlerGo), parameter)
	return (uint32)(out), IedClientError(err)
}

func (x *IedConnection) WriteObject(dataAttributeReference string, fc FunctionalConstraint, value *MmsValue) error {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_writeObject(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), value.ctx)
	return IedClientError(err)
}

func (x *IedConnection) WriteObjectAsync(dataAttributeReference string, fc FunctionalConstraint, value *MmsValue, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	mapGenericServiceHandlers.Store(handler, handler)
	defer mapGenericServiceHandlers.Delete(handler)
	out := C.IedConnection_writeObjectAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), value.ctx, (C.IedConnection_GenericServiceHandler)(C.fGenericServiceHandlerGo), parameter)
	return (uint32)(out), IedClientError(err)
}

func (x *IedConnection) ReadBooleanValue(dataAttributeReference string, fc FunctionalConstraint) (bool, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_readBooleanValue(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc))
	return (bool)(out), IedClientError(err)
}

func (x *IedConnection) ReadFloatValue(dataAttributeReference string, fc FunctionalConstraint) (float32, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_readFloatValue(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc))
	return (float32)(out), IedClientError(err)
}

func (x *IedConnection) ReadStringValue(dataAttributeReference string, fc FunctionalConstraint) (string, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_readStringValue(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc))
	return C.GoString(out), IedClientError(err)
}

func (x *IedConnection) ReadInt32Value(dataAttributeReference string, fc FunctionalConstraint) (int32, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_readInt32Value(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc))
	return (int32)(out), IedClientError(err)
}

func (x *IedConnection) ReadInt64Value(dataAttributeReference string, fc FunctionalConstraint) (int64, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_readInt64Value(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc))
	return (int64)(out), IedClientError(err)
}

func (x *IedConnection) ReadUnsigned32Value(dataAttributeReference string, fc FunctionalConstraint) (uint32, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_readUnsigned32Value(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc))
	return (uint32)(out), IedClientError(err)
}

func (x *IedConnection) ReadTimestampValue(dataAttributeReference string, fc FunctionalConstraint, timestamp *Timestamp) (*Timestamp, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	ctimestamp := &C.Timestamp{}
	if timestamp != nil {
		ctimestamp = timestamp.ctx
	}
	out := C.IedConnection_readTimestampValue(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), ctimestamp)
	timestamp.ctx = ctimestamp
	return &Timestamp{ctx: out}, IedClientError(err)
}

func (x *IedConnection) ReadQualityValue(dataAttributeReference string, fc FunctionalConstraint) (Quality, error) {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	out := C.IedConnection_readQualityValue(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc))
	return (Quality)(out), IedClientError(err)
}

func (x *IedConnection) WriteBooleanValue(dataAttributeReference string, fc FunctionalConstraint, value bool) error {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_writeBooleanValue(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), (C.bool)(value))
	return IedClientError(err)
}

func (x *IedConnection) WriteInt32Value(dataAttributeReference string, fc FunctionalConstraint, value int32) error {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_writeInt32Value(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), (C.int32_t)(value))
	return IedClientError(err)
}

func (x *IedConnection) WriteUnsigned32Value(dataAttributeReference string, fc FunctionalConstraint, value uint32) error {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_writeUnsigned32Value(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), (C.uint32_t)(value))
	return IedClientError(err)
}

func (x *IedConnection) WriteFloatValue(dataAttributeReference string, fc FunctionalConstraint, value float32) error {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_writeFloatValue(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), (C.float)(value))
	return IedClientError(err)
}

func (x *IedConnection) WriteVisibleStringValue(dataAttributeReference string, fc FunctionalConstraint, value string) error {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_writeVisibleStringValue(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), C.CString(value))
	return IedClientError(err)
}

func (x *IedConnection) WriteOctetStringValue(dataAttributeReference string, fc FunctionalConstraint, value []byte) error {
	cref := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_writeOctetString(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), (*C.uint8_t)(unsafe.SliceData(value)), (C.int)(len(value)))
	return IedClientError(err)
}

func (x *IedConnection) GetRCBValues(rcbReference string, updateRcb *ClientReportControlBlock) (ClientReportControlBlock, error) {
	cref := C.CString(rcbReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getRCBValues(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, updateRcb.ctx)
	return ClientReportControlBlock{ctx: ctx}, IedClientError(err)
}

type GetRCBValuesHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, rcb *ClientReportControlBlock)

var mapGetRCBValuesHandlers = sync.Map{}

//export fGetRCBValuesHandlerGo
func fGetRCBValuesHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, rcb C.ClientReportControlBlock) {
	mapGetRCBValuesHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(GetRCBValuesHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &ClientReportControlBlock{ctx: rcb})
		}
		return true
	})
}

func (x *IedConnection) GetRCBValuesAsync(rcbReference string, updateRcb *ClientReportControlBlock, handler GetRCBValuesHandler, parameter unsafe.Pointer) (uint32, error) {
	cref := C.CString(rcbReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	invokeId := C.IedConnection_getRCBValuesAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, updateRcb.ctx, (C.IedConnection_GetRCBValuesHandler)(C.fGetRCBValuesHandlerGo), parameter)
	return (uint32)(invokeId), IedClientError(err)
}

type ReasonForInclusion int

const (
	/** the element is not included in the received report */
	IEC61850_REASON_NOT_INCLUDED = C.IEC61850_REASON_NOT_INCLUDED

	/** the element is included due to a change of the data value */
	IEC61850_REASON_DATA_CHANGE = C.IEC61850_REASON_DATA_CHANGE

	/** the element is included due to a change in the quality of data */
	IEC61850_REASON_QUALITY_CHANGE = C.IEC61850_REASON_QUALITY_CHANGE

	/** the element is included due to an update of the data value */
	IEC61850_REASON_DATA_UPDATE = C.IEC61850_REASON_DATA_UPDATE

	/** the element is included due to a periodic integrity report task */
	IEC61850_REASON_INTEGRITY = C.IEC61850_REASON_INTEGRITY

	/** the element is included due to a general interrogation by the client */
	IEC61850_REASON_GI = C.IEC61850_REASON_GI

	/** the reason for inclusion is unknown (e.g. report is not configured to include reason-for-inclusion) */
	IEC61850_REASON_UNKNOWN = C.IEC61850_REASON_UNKNOWN

	/* Element encoding mask values for ClientReportControlBlock */

	/** include the report ID into the setRCB request */
	RCB_ELEMENT_RPT_ID = C.RCB_ELEMENT_RPT_ID

	/** include the report enable element into the setRCB request */
	RCB_ELEMENT_RPT_ENA = C.RCB_ELEMENT_RPT_ENA

	/** include the reservation element into the setRCB request (only available in unbuffered RCBs!) */
	RCB_ELEMENT_RESV = C.RCB_ELEMENT_RESV

	/** include the data set element into the setRCB request */
	RCB_ELEMENT_DATSET = C.RCB_ELEMENT_DATSET

	/** include the configuration revision element into the setRCB request */
	RCB_ELEMENT_CONF_REV = C.RCB_ELEMENT_CONF_REV

	/** include the option fields element into the setRCB request */
	RCB_ELEMENT_OPT_FLDS = C.RCB_ELEMENT_OPT_FLDS

	/** include the bufTm (event buffering time) element into the setRCB request */
	RCB_ELEMENT_BUF_TM = C.RCB_ELEMENT_BUF_TM

	/** include the sequence number element into the setRCB request (should be used!) */
	RCB_ELEMENT_SQ_NUM = C.RCB_ELEMENT_SQ_NUM

	/** include the trigger options element into the setRCB request */
	RCB_ELEMENT_TRG_OPS = C.RCB_ELEMENT_TRG_OPS

	/** include the integrity period element into the setRCB request */
	RCB_ELEMENT_INTG_PD = C.RCB_ELEMENT_INTG_PD

	/** include the GI (general interrogation) element into the setRCB request */
	RCB_ELEMENT_GI = C.RCB_ELEMENT_GI

	/** include the purge buffer element into the setRCB request (only available in buffered RCBs) */
	RCB_ELEMENT_PURGE_BUF = C.RCB_ELEMENT_PURGE_BUF

	/** include the entry ID element into the setRCB request (only available in buffered RCBs) */
	RCB_ELEMENT_ENTRY_ID = C.RCB_ELEMENT_ENTRY_ID

	/** include the time of entry element into the setRCB request (only available in buffered RCBs) */
	RCB_ELEMENT_TIME_OF_ENTRY = C.RCB_ELEMENT_TIME_OF_ENTRY

	/** include the reservation time element into the setRCB request (only available in buffered RCBs) */
	RCB_ELEMENT_RESV_TMS = C.RCB_ELEMENT_RESV_TMS

	/** include the owner element into the setRCB request */
	RCB_ELEMENT_OWNER = C.RCB_ELEMENT_OWNER
)

func (x ReasonForInclusion) String() string {
	return C.GoString(C.ReasonForInclusion_getValueAsString(C.ReasonForInclusion(x)))
}

func (x *IedConnection) SetRCBValues(rcbReference string, rcb *ClientReportControlBlock, parametersMask uint32, singleRequest bool) error {
	cref := C.CString(rcbReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_setRCBValues(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), rcb.ctx, C.uint32_t(parametersMask), C.bool(singleRequest))
	return IedClientError(err)
}

func (x *IedConnection) SetRCBValuesAsync(rcbReference string, rcb *ClientReportControlBlock, parametersMask uint32, singleRequest bool, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	cref := C.CString(rcbReference)
	defer C.free(unsafe.Pointer(cref))
	mapGenericServiceHandlers.Store(handler, handler)
	defer mapGenericServiceHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_setRCBValuesAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), rcb.ctx, C.uint32_t(parametersMask), C.bool(singleRequest), (C.IedConnection_GenericServiceHandler)(C.fGenericServiceHandlerGo), parameter)
	return (uint32)(invokeId), IedClientError(err)
}

type ReportCallbackFunction func(parameter unsafe.Pointer, report *ClientReport)

var mapReportCallbackFunctions = sync.Map{}

//export fReportCallbackFunctionGo
func fReportCallbackFunctionGo(parameter unsafe.Pointer, report C.ClientReport) {
	mapReportCallbackFunctions.Range(func(key, value any) bool {
		if fn, ok := value.(ReportCallbackFunction); ok {
			fn(parameter, &ClientReport{ctx: report})
		}
		return true
	})
}

func (x *IedConnection) InstallReportHandler(rcbReference string, rptId string, handler ReportCallbackFunction, parameter unsafe.Pointer) {
	cref := C.CString(rcbReference)
	defer C.free(unsafe.Pointer(cref))
	cid := C.CString(rptId)
	defer C.free(unsafe.Pointer(cid))
	mapReportCallbackFunctions.Store(rcbReference, handler)
	C.IedConnection_installReportHandler(x.ctx, cref, cid, (C.ReportCallbackFunction)(C.fReportCallbackFunctionGo), parameter)
}

func (x *IedConnection) UninstallReportHandler(rcbReference string) {
	cref := C.CString(rcbReference)
	defer C.free(unsafe.Pointer(cref))
	C.IedConnection_uninstallReportHandler(x.ctx, cref)
	mapReportCallbackFunctions.Delete(rcbReference)
}

func (x *IedConnection) TriggerGIReport(rcbReference string) error {
	cref := C.CString(rcbReference)
	defer C.free(unsafe.Pointer(cref))
	err := IED_ERROR_OK
	C.IedConnection_triggerGIReport(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref)
	return IedClientError(err)
}

func (x *ClientReport) GetDataSetName() string {
	return C.GoString(C.ClientReport_getDataSetName(x.ctx))
}

func (x *ClientReport) GetDataSetValues() *MmsValue {
	return &MmsValue{ctx: C.ClientReport_getDataSetValues(x.ctx)}
}

func (x *ClientReport) GetRcbReference() string {
	return C.GoString(C.ClientReport_getRcbReference(x.ctx))
}

func (x *ClientReport) GetRptId() string {
	return C.GoString(C.ClientReport_getRptId(x.ctx))
}

func (x *ClientReport) GetReasonForInclusion(elementIndex int) ReasonForInclusion {
	return ReasonForInclusion(C.ClientReport_getReasonForInclusion(x.ctx, C.int(elementIndex)))
}

func (x *ClientReport) GetEntryId() *MmsValue {
	return &MmsValue{ctx: C.ClientReport_getEntryId(x.ctx)}
}

func (x *ClientReport) HasTimestamp() bool {
	return bool(C.ClientReport_hasTimestamp(x.ctx))
}

func (x *ClientReport) HasSeqNum() bool {
	return bool(C.ClientReport_hasSeqNum(x.ctx))
}

func (x *ClientReport) GetSeqNum() uint16 {
	return (uint16)(C.ClientReport_getSeqNum(x.ctx))
}

func (x *ClientReport) HasDataSetName() bool {
	return bool(C.ClientReport_hasDataSetName(x.ctx))
}

func (x *ClientReport) HasReasonForInclusion() bool {
	return bool(C.ClientReport_hasReasonForInclusion(x.ctx))
}

func (x *ClientReport) HasConfRev() bool {
	return bool(C.ClientReport_hasConfRev(x.ctx))
}

func (x *ClientReport) GetConfRev() uint32 {
	return (uint32)(C.ClientReport_getConfRev(x.ctx))
}

func (x *ClientReport) HasBufOvfl() bool {
	return bool(C.ClientReport_hasBufOvfl(x.ctx))
}

func (x *ClientReport) GetBufOvfl() bool {
	return bool(C.ClientReport_getBufOvfl(x.ctx))
}

func (x *ClientReport) HasDataReference() bool {
	return bool(C.ClientReport_hasDataReference(x.ctx))
}

func (x *ClientReport) GetDataReference(elementIndex int) string {
	return C.GoString(C.ClientReport_getDataReference(x.ctx, C.int(elementIndex)))
}

func (x *ClientReport) GetTimestamp() uint64 {
	return (uint64)(C.ClientReport_getTimestamp(x.ctx))
}

func (x *ClientReport) HasSubSeqNum() bool {
	return bool(C.ClientReport_hasSubSeqNum(x.ctx))
}

func (x *ClientReport) GetSubSeqNum() uint16 {
	return (uint16)(C.ClientReport_getSubSeqNum(x.ctx))
}

func (x *ClientReport) GetMoreSeqmentsFollow() bool {
	return bool(C.ClientReport_getMoreSeqmentsFollow(x.ctx))
}

func ClientReportControlBlockCreate(rcbReference string) *ClientReportControlBlock {
	cref := C.CString(rcbReference)
	defer C.free(unsafe.Pointer(cref))
	return &ClientReportControlBlock{ctx: C.ClientReportControlBlock_create(cref)}
}

func (x *ClientReportControlBlock) Destroy() {
	C.ClientReportControlBlock_destroy(x.ctx)
}

func (x *ClientReportControlBlock) GetObjectReference() string {
	return C.GoString(C.ClientReportControlBlock_getObjectReference(x.ctx))
}

func (x *ClientReportControlBlock) IsBuffered() bool {
	return bool(C.ClientReportControlBlock_isBuffered(x.ctx))
}

func (x *ClientReportControlBlock) GetRptId() string {
	return C.GoString(C.ClientReportControlBlock_getRptId(x.ctx))
}

func (x *ClientReportControlBlock) SetRptId(rptId string) {
	cid := C.CString(rptId)
	defer C.free(unsafe.Pointer(cid))
	C.ClientReportControlBlock_setRptId(x.ctx, cid)
}

func (x *ClientReportControlBlock) GetRptEna() bool {
	return bool(C.ClientReportControlBlock_getRptEna(x.ctx))
}

func (x *ClientReportControlBlock) SetRptEna(rptEna bool) {
	C.ClientReportControlBlock_setRptEna(x.ctx, C.bool(rptEna))
}

func (x *ClientReportControlBlock) GetResv() bool {
	return bool(C.ClientReportControlBlock_getResv(x.ctx))
}

func (x *ClientReportControlBlock) SetResv(resv bool) {
	C.ClientReportControlBlock_setResv(x.ctx, C.bool(resv))
}

func (x *ClientReportControlBlock) GetDataSetReference() string {
	return C.GoString(C.ClientReportControlBlock_getDataSetReference(x.ctx))
}

func (x *ClientReportControlBlock) SetDataSetReference(dataSetReference string) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	C.ClientReportControlBlock_setDataSetReference(x.ctx, cid)
}

func (x *ClientReportControlBlock) GetConfRev() uint32 {
	return (uint32)(C.ClientReportControlBlock_getConfRev(x.ctx))
}

func (x *ClientReportControlBlock) GetOptFlds() int {
	return (int)(C.ClientReportControlBlock_getOptFlds(x.ctx))
}

func (x *ClientReportControlBlock) SetOptFlds(optFlds int) {
	C.ClientReportControlBlock_setOptFlds(x.ctx, C.int(optFlds))
}

func (x *ClientReportControlBlock) GetBufTm() uint32 {
	return (uint32)(C.ClientReportControlBlock_getBufTm(x.ctx))
}

func (x *ClientReportControlBlock) SetBufTm(bufTm uint32) {
	C.ClientReportControlBlock_setBufTm(x.ctx, C.uint32_t(bufTm))
}

func (x *ClientReportControlBlock) GetSqNum() uint16 {
	return (uint16)(C.ClientReportControlBlock_getSqNum(x.ctx))
}

func (x *ClientReportControlBlock) GetTrgOps() int {
	return (int)(C.ClientReportControlBlock_getTrgOps(x.ctx))
}

func (x *ClientReportControlBlock) SetTrgOps(trgOps int) {
	C.ClientReportControlBlock_setTrgOps(x.ctx, C.int(trgOps))
}

func (x *ClientReportControlBlock) GetIntgPd() uint32 {
	return (uint32)(C.ClientReportControlBlock_getIntgPd(x.ctx))
}

func (x *ClientReportControlBlock) SetIntgPd(intgPd uint32) {
	C.ClientReportControlBlock_setIntgPd(x.ctx, C.uint32_t(intgPd))
}

func (x *ClientReportControlBlock) GetGI() bool {
	return bool(C.ClientReportControlBlock_getGI(x.ctx))
}

func (x *ClientReportControlBlock) SetGI(gi bool) {
	C.ClientReportControlBlock_setGI(x.ctx, C.bool(gi))
}

func (x *ClientReportControlBlock) GetPurgeBuf() bool {
	return bool(C.ClientReportControlBlock_getPurgeBuf(x.ctx))
}

func (x *ClientReportControlBlock) SetPurgeBuf(purgeBuf bool) {
	C.ClientReportControlBlock_setPurgeBuf(x.ctx, C.bool(purgeBuf))
}

func (x *ClientReportControlBlock) HasResvTms() bool {
	return bool(C.ClientReportControlBlock_hasResvTms(x.ctx))
}

func (x *ClientReportControlBlock) GetResvTms() int16 {
	return (int16)(C.ClientReportControlBlock_getResvTms(x.ctx))
}

func (x *ClientReportControlBlock) SetResvTms(resvTms int16) {
	C.ClientReportControlBlock_setResvTms(x.ctx, C.int16_t(resvTms))
}

func (x *ClientReportControlBlock) GetEntryId() *MmsValue {
	return &MmsValue{ctx: C.ClientReportControlBlock_getEntryId(x.ctx)}
}

func (x *ClientReportControlBlock) SetEntryId(entryId *MmsValue) {
	C.ClientReportControlBlock_setEntryId(x.ctx, entryId.ctx)
}

func (x *ClientReportControlBlock) GetEntryTime() uint64 {
	return (uint64)(C.ClientReportControlBlock_getEntryTime(x.ctx))
}

func (x *ClientReportControlBlock) GetOwner() *MmsValue {
	return &MmsValue{ctx: C.ClientReportControlBlock_getOwner(x.ctx)}
}

func (x *IedConnection) ReadDataSetValues(dataSetReference string, dataSet *ClientDataSet) (*ClientDataSet, error) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	err := IED_ERROR_OK
	out := C.IedConnection_readDataSetValues(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, dataSet.ctx)
	return &ClientDataSet{ctx: out}, IedClientError(err)
}

type ReadDataSetHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, dataSet *ClientDataSet)

var mapReadDataSetHandlers = sync.Map{}

//export fReadDataSetHandlerGo
func fReadDataSetHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, dataSet C.ClientDataSet) {
	mapReadDataSetHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(ReadDataSetHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &ClientDataSet{ctx: dataSet})
		}
		return true
	})
}

func (x *IedConnection) ReadDataSetValuesAsync(dataSetReference string, dataSet *ClientDataSet, handler ReadDataSetHandler, parameter unsafe.Pointer) (uint32, error) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	mapReadDataSetHandlers.Store(handler, handler)
	defer mapReadDataSetHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_readDataSetValuesAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, dataSet.ctx, (C.IedConnection_ReadDataSetHandler)(C.fReadDataSetHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) CreateDataSet(dataSetReference string, dataSetElements *LinkedList) error {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	err := IED_ERROR_OK
	C.IedConnection_createDataSet(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, dataSetElements.ctx)
	return IedClientError(err)
}

func (x *IedConnection) CreateDataSetAsync(dataSetReference string, dataSetElements *LinkedList, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	mapGenericServiceHandlers.Store(handler, handler)
	defer mapGenericServiceHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_createDataSetAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, dataSetElements.ctx, (C.IedConnection_GenericServiceHandler)(C.fGenericServiceHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) DeleteDataSet(dataSetReference string) error {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	err := IED_ERROR_OK
	C.IedConnection_deleteDataSet(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid)
	return IedClientError(err)
}

func (x *IedConnection) DeleteDataSetAsync(dataSetReference string, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	mapGenericServiceHandlers.Store(handler, handler)
	defer mapGenericServiceHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_deleteDataSetAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, (C.IedConnection_GenericServiceHandler)(C.fGenericServiceHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) GetDataSetDirectory(dataSetReference string, isDeletable *bool) (*LinkedList, error) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	err := IED_ERROR_OK
	out := C.IedConnection_getDataSetDirectory(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, (*C.bool)(unsafe.Pointer(isDeletable)))
	return &LinkedList{ctx: out}, IedClientError(err)
}

type GetDataSetDirectoryHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, dataSetDirectory *LinkedList, isDeletable bool)

var mapGetDataSetDirectoryHandlers = sync.Map{}

//export fGetDataSetDirectoryHandlerGo
func fGetDataSetDirectoryHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, dataSetDirectory C.LinkedList, isDeletable C._Bool) {
	mapGetDataSetDirectoryHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(GetDataSetDirectoryHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &LinkedList{ctx: dataSetDirectory}, bool(isDeletable))
		}
		return true
	})
}

func (x *IedConnection) GetDataSetDirectoryAsync(dataSetReference string, handler GetDataSetDirectoryHandler, parameter unsafe.Pointer) (uint32, error) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	mapGetDataSetDirectoryHandlers.Store(handler, handler)
	defer mapGetDataSetDirectoryHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_getDataSetDirectoryAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, (C.IedConnection_GetDataSetDirectoryHandler)(C.fGetDataSetDirectoryHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) WriteDataSetValues(dataSetReference string, values *LinkedList) (*LinkedList, error) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	var accessResults *C.LinkedList
	err := IED_ERROR_OK
	C.IedConnection_writeDataSetValues(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, values.ctx, (*C.LinkedList)(unsafe.Pointer(accessResults)))
	return &LinkedList{ctx: *accessResults}, IedClientError(err)
}

type WriteDataSetHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, accessResults *LinkedList)

var mapWriteDataSetHandlers = sync.Map{}

//export fWriteDataSetHandlerGo
func fWriteDataSetHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, accessResults C.LinkedList) {
	mapWriteDataSetHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(WriteDataSetHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &LinkedList{ctx: accessResults})
		}
		return true
	})
}

func (x *IedConnection) WriteDataSetValuesAsync(dataSetReference string, values *LinkedList, handler WriteDataSetHandler, parameter unsafe.Pointer) (uint32, error) {
	cid := C.CString(dataSetReference)
	defer C.free(unsafe.Pointer(cid))
	mapWriteDataSetHandlers.Store(handler, handler)
	defer mapWriteDataSetHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_writeDataSetValuesAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cid, values.ctx, (C.IedConnection_WriteDataSetHandler)(C.fWriteDataSetHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *ClientDataSet) Destroy() {
	C.ClientDataSet_destroy(x.ctx)
}

func (x *ClientDataSet) GetValues() *MmsValue {
	return &MmsValue{ctx: C.ClientDataSet_getValues(x.ctx)}
}

func (x *ClientDataSet) GetReference() string {
	return C.GoString(C.ClientDataSet_getReference(x.ctx))
}

func (x *ClientDataSet) GetDataSetSize() int {
	return int(C.ClientDataSet_getDataSetSize(x.ctx))
}

type ControlObjectClient struct {
	ctx C.ControlObjectClient
}

func ControlObjectClientCreate(objectReference string, connection *IedConnection) *ControlObjectClient {
	cid := C.CString(objectReference)
	defer C.free(unsafe.Pointer(cid))
	return &ControlObjectClient{ctx: C.ControlObjectClient_create(cid, connection.ctx)}
}

func ControlObjectClientCreateEx(objectReference string, connection *IedConnection, ctlModel ControlModel, controlObjectSpec *MmsVariableSpecification) *ControlObjectClient {
	cid := C.CString(objectReference)
	defer C.free(unsafe.Pointer(cid))
	return &ControlObjectClient{ctx: C.ControlObjectClient_createEx(cid, connection.ctx, C.ControlModel(ctlModel), controlObjectSpec.ctx)}
}

func (x *ControlObjectClient) Destroy() {
	C.ControlObjectClient_destroy(x.ctx)
}

type ControlActionType int32

const (
	CONTROL_ACTION_TYPE_SELECT  ControlActionType = C.CONTROL_ACTION_TYPE_SELECT
	CONTROL_ACTION_TYPE_OPERATE ControlActionType = C.CONTROL_ACTION_TYPE_OPERATE
	CONTROL_ACTION_TYPE_CANCEL  ControlActionType = C.CONTROL_ACTION_TYPE_CANCEL
)

type ControlActionHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, type_ ControlActionType, success bool)

var mapControlActionHandlers = sync.Map{}

//export fControlActionHandlerGo
func fControlActionHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, type_ C.ControlActionType, success C._Bool) {
	mapControlActionHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(ControlActionHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), ControlActionType(type_), bool(success))
		}
		return true
	})
}

func (x *ControlObjectClient) GetObjectReference() string {
	return C.GoString(C.ControlObjectClient_getObjectReference(x.ctx))
}

func (x *ControlObjectClient) GetControlModel() ControlModel {
	return ControlModel(C.ControlObjectClient_getControlModel(x.ctx))
}

func (x *ControlObjectClient) SetControlModel(ctlModel ControlModel) {
	C.ControlObjectClient_setControlModel(x.ctx, C.ControlModel(ctlModel))
}

func (x *ControlObjectClient) ChangeServerControlModel(ctlModel ControlModel) {
	C.ControlObjectClient_changeServerControlModel(x.ctx, C.ControlModel(ctlModel))
}

func (x *ControlObjectClient) GetCtlValType() MmsType {
	return MmsType(C.ControlObjectClient_getCtlValType(x.ctx))
}

func (x *ControlObjectClient) GetLastError() IedClientError {
	return IedClientError(C.ControlObjectClient_getLastError(x.ctx))
}

func (x *ControlObjectClient) Operate(ctlVal *MmsValue, operTime uint64) bool {
	return bool(C.ControlObjectClient_operate(x.ctx, ctlVal.ctx, C.uint64_t(operTime)))
}

func (x *ControlObjectClient) Select() bool {
	return bool(C.ControlObjectClient_select(x.ctx))
}

func (x *ControlObjectClient) SelectWithValue(ctlVal *MmsValue) bool {
	return bool(C.ControlObjectClient_selectWithValue(x.ctx, ctlVal.ctx))
}

func (x *ControlObjectClient) Cancel() bool {
	return bool(C.ControlObjectClient_cancel(x.ctx))
}
func (x *ControlObjectClient) OperateAsync(ctlVal *MmsValue, operTime uint64, handler ControlActionHandler, parameter unsafe.Pointer) (uint32, error) {
	mapControlActionHandlers.Store(handler, handler)
	defer mapControlActionHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.ControlObjectClient_operateAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), ctlVal.ctx, C.uint64_t(operTime),
		(C.ControlObjectClient_ControlActionHandler)(C.fControlActionHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *ControlObjectClient) SelectAsync(handler ControlActionHandler, parameter unsafe.Pointer) (uint32, error) {
	mapControlActionHandlers.Store(handler, handler)
	defer mapControlActionHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.ControlObjectClient_selectAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)),
		(C.ControlObjectClient_ControlActionHandler)(C.fControlActionHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *ControlObjectClient) SelectWithValueAsync(ctlVal *MmsValue, handler ControlActionHandler, parameter unsafe.Pointer) (uint32, error) {
	mapControlActionHandlers.Store(handler, handler)
	defer mapControlActionHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.ControlObjectClient_selectWithValueAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), ctlVal.ctx, (C.ControlObjectClient_ControlActionHandler)(C.fControlActionHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *ControlObjectClient) CancelAsync(handler ControlActionHandler, parameter unsafe.Pointer) (uint32, error) {
	mapControlActionHandlers.Store(handler, handler)
	defer mapControlActionHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.ControlObjectClient_cancelAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)),
		(C.ControlObjectClient_ControlActionHandler)(C.fControlActionHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *ControlObjectClient) GetLastApplError() *LastApplError {
	ctx := C.ControlObjectClient_getLastApplError(x.ctx)
	return &LastApplError{ctx: &ctx}
}

func (x *ControlObjectClient) SetTestMode(value bool) {
	C.ControlObjectClient_setTestMode(x.ctx, C._Bool(value))
}

func (x *ControlObjectClient) SetOrigin(orIdent string, orCat int) {
	C.ControlObjectClient_setOrigin(x.ctx, C.CString(orIdent), C.int(orCat))
}

func (x *ControlObjectClient) UseConstantT(useConstantT bool) {
	C.ControlObjectClient_useConstantT(x.ctx, C._Bool(useConstantT))
}

// EnableInterlockCheck
//
// Deprecated: use ControlObjectClient_setInterlockCheck instead
func (x *ControlObjectClient) EnableInterlockCheck() {
	C.ControlObjectClient_enableInterlockCheck(x.ctx)
}

// EnableSynchroCheck
//
// Deprecated: use ControlObjectClient_setSynchroCheck instead
func (x *ControlObjectClient) EnableSynchroCheck() {
	C.ControlObjectClient_enableSynchroCheck(x.ctx)
}

// SetCtlNum
//
// Deprecated: Do not use (ctlNum is handled automatically by the library)! Intended for test purposes only.
func (x *ControlObjectClient) SetCtlNum(ctlNum uint8) {
	C.ControlObjectClient_setCtlNum(x.ctx, C.uint8_t(ctlNum))
}

func (x *ControlObjectClient) SetInterlockCheck(value bool) {
	C.ControlObjectClient_setInterlockCheck(x.ctx, C._Bool(value))
}

func (x *ControlObjectClient) SetSynchroCheck(value bool) {
	C.ControlObjectClient_setSynchroCheck(x.ctx, C._Bool(value))
}

type CommandTerminationHandler func(parameter unsafe.Pointer, controlClient *ControlObjectClient)

var mapCommandTerminationHandlers = sync.Map{}

//export fCommandTerminationHandlerGo
func fCommandTerminationHandlerGo(parameter unsafe.Pointer, controlClient C.ControlObjectClient) {
	mapCommandTerminationHandlers.Range(func(key, value interface{}) bool {
		if fn, ok := value.(CommandTerminationHandler); ok {
			fn(parameter, &ControlObjectClient{ctx: controlClient})
		}
		return true
	})
}

func (x *ControlObjectClient) SetCommandTerminationHandler(handler CommandTerminationHandler, parameter unsafe.Pointer) {
	mapCommandTerminationHandlers.Store(handler, handler)
	defer mapCommandTerminationHandlers.Delete(handler)
	C.ControlObjectClient_setCommandTerminationHandler(x.ctx, (C.CommandTerminationHandler)(C.fCommandTerminationHandlerGo), parameter)
}

func (x *IedConnection) GetDeviceModelFromServer() error {
	err := IED_ERROR_OK
	C.IedConnection_getDeviceModelFromServer(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)))
	return IedClientError(err)
}

func (x *IedConnection) GetLogicalDeviceList() (*LinkedList, error) {
	err := IED_ERROR_OK
	ctx := C.IedConnection_getLogicalDeviceList(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)))
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetServerDirectory(getFileNames bool) (*LinkedList, error) {
	err := IED_ERROR_OK
	ctx := C.IedConnection_getServerDirectory(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), C._Bool(getFileNames))
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetLogicalDeviceDirectory(logicalDeviceName string) (*LinkedList, error) {
	cstr := C.CString(logicalDeviceName)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getLogicalDeviceDirectory(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr)
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetLogicalNodeVariables(logicalNodeReference string) (*LinkedList, error) {
	cstr := C.CString(logicalNodeReference)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getLogicalNodeVariables(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr)
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetLogicalNodeDirectory(logicalNodeReference string, acsiClass ACSIClass) (*LinkedList, error) {
	cstr := C.CString(logicalNodeReference)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getLogicalNodeDirectory(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, C.ACSIClass(acsiClass))
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetDataDirectory(dataReference string) (*LinkedList, error) {
	cstr := C.CString(dataReference)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getDataDirectory(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr)
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetDataDirectoryFC(dataReference string) (*LinkedList, error) {
	cstr := C.CString(dataReference)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getDataDirectoryFC(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr)
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetDataDirectoryByFC(dataReference string, fc FunctionalConstraint) (*LinkedList, error) {
	cstr := C.CString(dataReference)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getDataDirectoryByFC(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, C.FunctionalConstraint(fc))
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetVariableSpecification(dataAttributeReference string, fc FunctionalConstraint) (*MmsVariableSpecification, error) {
	cstr := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getVariableSpecification(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, C.FunctionalConstraint(fc))
	return &MmsVariableSpecification{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetLogicalDeviceVariables(ldName string) (*LinkedList, error) {
	cstr := C.CString(ldName)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getLogicalDeviceVariables(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr)
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetLogicalDeviceDataSets(ldName string) (*LinkedList, error) {
	cstr := C.CString(ldName)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getLogicalDeviceDataSets(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr)
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

type GetNameListHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, nameList *LinkedList, moreFollows bool)

var mapGetNameListHandlers = sync.Map{}

//export fGetNameListHandlerGo
func fGetNameListHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, nameList C.LinkedList, moreFollows C._Bool) {
	if fn, ok := mapGetNameListHandlers.Load(parameter); ok {
		if fn, ok := fn.(GetNameListHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &LinkedList{ctx: nameList}, bool(moreFollows))
		}
	}
}

func (x *IedConnection) GetServerDirectoryAsync(continueAfter string, result *LinkedList, handler GetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(continueAfter)
	defer C.free(unsafe.Pointer(cstr))
	mapGetNameListHandlers.Store(handler, handler)
	defer mapGetNameListHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_getServerDirectoryAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, result.ctx, (C.IedConnection_GetNameListHandler)(C.fGetNameListHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) GetLogicalDeviceVariablesAsync(ldName string, continueAfter string, result *LinkedList, handler GetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(ldName)
	defer C.free(unsafe.Pointer(cstr))
	cstr2 := C.CString(continueAfter)
	defer C.free(unsafe.Pointer(cstr2))
	mapGetNameListHandlers.Store(handler, handler)
	defer mapGetNameListHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_getLogicalDeviceVariablesAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, cstr2, result.ctx, (C.IedConnection_GetNameListHandler)(C.fGetNameListHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) GetLogicalDeviceDataSetsAsync(ldName string, continueAfter string, result *LinkedList, handler GetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(ldName)
	defer C.free(unsafe.Pointer(cstr))
	cstr2 := C.CString(continueAfter)
	defer C.free(unsafe.Pointer(cstr2))
	mapGetNameListHandlers.Store(handler, handler)
	defer mapGetNameListHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_getLogicalDeviceDataSetsAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, cstr2, result.ctx, (C.IedConnection_GetNameListHandler)(C.fGetNameListHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

type GetVariableSpecificationHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, variableSpecification *MmsVariableSpecification)

var mapGetVariableSpecificationHandlers = sync.Map{}

//export fGetVariableSpecificationHandlerGo
func fGetVariableSpecificationHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, variableSpecification *C.MmsVariableSpecification) {
	if fn, ok := mapGetVariableSpecificationHandlers.Load(parameter); ok {
		if fn, ok := fn.(GetVariableSpecificationHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &MmsVariableSpecification{ctx: variableSpecification})
		}
	}
}

func (x *IedConnection) GetVariableSpecificationAsync(dataAttributeReference string, fc FunctionalConstraint, handler GetVariableSpecificationHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(dataAttributeReference)
	defer C.free(unsafe.Pointer(cstr))
	mapGetVariableSpecificationHandlers.Store(handler, handler)
	defer mapGetVariableSpecificationHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_getVariableSpecificationAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, C.FunctionalConstraint(fc), (C.IedConnection_GetVariableSpecificationHandler)(C.fGetVariableSpecificationHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) QueryLogByTime(logReference string, startTime uint64, endTime uint64, moreFollows *bool) (*LinkedList, error) {
	cstr := C.CString(logReference)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_queryLogByTime(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, C.uint64_t(startTime), C.uint64_t(endTime), (*C._Bool)(unsafe.Pointer(moreFollows)))
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) QueryLogAfter(logReference string, entryID *MmsValue, timeStamp uint64, moreFollows *bool) (*LinkedList, error) {
	cstr := C.CString(logReference)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_queryLogAfter(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, entryID.ctx, C.uint64_t(timeStamp), (*C._Bool)(unsafe.Pointer(moreFollows)))
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

type QueryLogHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, journalEntries *LinkedList, moreFollows bool)

var mapQueryLogHandlers = sync.Map{}

//export fQueryLogHandlerGo
func fQueryLogHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, journalEntries C.LinkedList, moreFollows C._Bool) {
	if fn, ok := mapQueryLogHandlers.Load(parameter); ok {
		if fn, ok := fn.(QueryLogHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), &LinkedList{ctx: journalEntries}, bool(moreFollows))
		}
	}
}

func (x *IedConnection) QueryLogByTimeAsync(logReference string, startTime uint64, endTime uint64, handler QueryLogHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(logReference)
	defer C.free(unsafe.Pointer(cstr))
	mapQueryLogHandlers.Store(handler, handler)
	defer mapQueryLogHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_queryLogByTimeAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, C.uint64_t(startTime), C.uint64_t(endTime), (C.IedConnection_QueryLogHandler)(C.fQueryLogHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) QueryLogAfterAsync(logReference string, entryID *MmsValue, timeStamp uint64, handler QueryLogHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(logReference)
	defer C.free(unsafe.Pointer(cstr))
	mapQueryLogHandlers.Store(handler, handler)
	defer mapQueryLogHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_queryLogAfterAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, entryID.ctx, C.uint64_t(timeStamp), (C.IedConnection_QueryLogHandler)(C.fQueryLogHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

type FileDirectoryEntry struct {
	ctx C.FileDirectoryEntry
}

func FileDirectoryEntryCreate(fileName string, fileSize uint32, lastModified uint64) *FileDirectoryEntry {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	ctx := C.FileDirectoryEntry_create(cstr, C.uint32_t(fileSize), C.uint64_t(lastModified))
	return &FileDirectoryEntry{ctx: ctx}
}

func (x *FileDirectoryEntry) Destroy() {
	C.FileDirectoryEntry_destroy(x.ctx)
}

func (x *FileDirectoryEntry) GetFileName() string {
	return C.GoString(C.FileDirectoryEntry_getFileName(x.ctx))
}

func (x *FileDirectoryEntry) GetFileSize() uint32 {
	return uint32(C.FileDirectoryEntry_getFileSize(x.ctx))
}

func (x *FileDirectoryEntry) GetLastModified() uint64 {
	return uint64(C.FileDirectoryEntry_getLastModified(x.ctx))
}

func (x *IedConnection) GetFileDirectory(directoryName string) (*LinkedList, error) {
	cstr := C.CString(directoryName)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getFileDirectory(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr)
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

func (x *IedConnection) GetFileDirectoryEx(directoryName string, continueAfter string, moreFollows *bool) (*LinkedList, error) {
	cstr1 := C.CString(directoryName)
	defer C.free(unsafe.Pointer(cstr1))
	cstr2 := C.CString(continueAfter)
	defer C.free(unsafe.Pointer(cstr2))
	err := IED_ERROR_OK
	ctx := C.IedConnection_getFileDirectoryEx(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr1, cstr2, (*C._Bool)(unsafe.Pointer(moreFollows)))
	return &LinkedList{ctx: ctx}, IedClientError(err)
}

type FileDirectoryEntryHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, filename string, size uint32, lastModfified uint64, moreFollows bool)

var mapFileDirectoryEntryHandlers = sync.Map{}

//export fFileDirectoryEntryHandlerGo
func fFileDirectoryEntryHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, filename *C.char, size C.uint32_t, lastModfified C.uint64_t, moreFollows C._Bool) {
	if fn, ok := mapFileDirectoryEntryHandlers.Load(parameter); ok {
		if fn, ok := fn.(FileDirectoryEntryHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err), C.GoString(filename), uint32(size), uint64(lastModfified), bool(moreFollows))
		}
	}
}

func (x *IedConnection) GetFileDirectoryAsyncEx(directoryName string, continueAfter string, handler FileDirectoryEntryHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr1 := C.CString(directoryName)
	defer C.free(unsafe.Pointer(cstr1))
	cstr2 := C.CString(continueAfter)
	defer C.free(unsafe.Pointer(cstr2))
	mapFileDirectoryEntryHandlers.Store(handler, handler)
	defer mapFileDirectoryEntryHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_getFileDirectoryAsyncEx(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr1, cstr2, (C.IedConnection_FileDirectoryEntryHandler)(C.fFileDirectoryEntryHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

type IedClientGetFileHandler func(parameter unsafe.Pointer, buffer []byte) bool

var mapIedClientGetFileHandlers = sync.Map{}

//export fIedClientGetFileHandlerGo
func fIedClientGetFileHandlerGo(parameter unsafe.Pointer, buffer *C.uint8_t, bufferSize C.uint32_t) C._Bool {
	if fn, ok := mapIedClientGetFileHandlers.Load(parameter); ok {
		if fn, ok := fn.(IedClientGetFileHandler); ok {
			return C._Bool(fn(parameter, C.GoBytes(unsafe.Pointer(buffer), C.int(bufferSize))))
		}
	}
	return C._Bool(false)
}

func (x *IedConnection) GetFile(fileName string, handler IedClientGetFileHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	mapIedClientGetFileHandlers.Store(handler, handler)
	defer mapIedClientGetFileHandlers.Delete(handler)
	err := IED_ERROR_OK
	bytesRead := C.IedConnection_getFile(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, (C.IedConnection_GetFileAsyncHandler)(C.fIedClientGetFileHandlerGo), parameter)
	return uint32(bytesRead), IedClientError(err)
}

type GetFileAsyncHandler func(invokeId uint32, parameter unsafe.Pointer, err IedClientError, originalInvokeId uint32, buffer []byte, moreFollows bool) bool

var mapGetFileAsyncHandlers = sync.Map{}

//export fGetFileAsyncHandlerGo
func fGetFileAsyncHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError, originalInvokeId C.uint32_t, buffer *C.uint8_t, bytesRead C.uint32_t, moreFollows C._Bool) C._Bool {
	if fn, ok := mapGetFileAsyncHandlers.Load(parameter); ok {
		if fn, ok := fn.(GetFileAsyncHandler); ok {
			return C._Bool(fn(uint32(invokeId), parameter, IedClientError(err), uint32(originalInvokeId), C.GoBytes(unsafe.Pointer(buffer), C.int(bytesRead)), bool(moreFollows)))
		}
	}
	return C._Bool(false)
}

func (x *IedConnection) GetFileAsync(fileName string, handler GetFileAsyncHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	mapGetFileAsyncHandlers.Store(handler, handler)
	defer mapGetFileAsyncHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_getFileAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, (C.IedConnection_GetFileAsyncHandler)(C.fGetFileAsyncHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) SetFilestoreBasepath(basepath string) {
	cstr := C.CString(basepath)
	defer C.free(unsafe.Pointer(cstr))
	C.IedConnection_setFilestoreBasepath(x.ctx, cstr)
}

func (x *IedConnection) SetFile(sourceFilename string, destinationFilename string) error {
	cstr1 := C.CString(sourceFilename)
	defer C.free(unsafe.Pointer(cstr1))
	cstr2 := C.CString(destinationFilename)
	defer C.free(unsafe.Pointer(cstr2))
	err := IED_ERROR_OK
	C.IedConnection_setFile(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr1, cstr2)
	return IedClientError(err)
}

func (x *IedConnection) SetFileAsync(sourceFilename string, destinationFilename string, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr1 := C.CString(sourceFilename)
	defer C.free(unsafe.Pointer(cstr1))
	cstr2 := C.CString(destinationFilename)
	defer C.free(unsafe.Pointer(cstr2))
	mapGenericServiceHandlers.Store(handler, handler)
	defer mapGenericServiceHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_setFileAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr1, cstr2, (C.IedConnection_GenericServiceHandler)(C.fGenericServiceHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}

func (x *IedConnection) DeleteFile(fileName string) error {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	err := IED_ERROR_OK
	C.IedConnection_deleteFile(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr)
	return IedClientError(err)
}

func (x *IedConnection) DeleteFileAsync(fileName string, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	mapGenericServiceHandlers.Store(handler, handler)
	defer mapGenericServiceHandlers.Delete(handler)
	err := IED_ERROR_OK
	invokeId := C.IedConnection_deleteFileAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cstr, (C.IedConnection_GenericServiceHandler)(C.fGenericServiceHandlerGo), parameter)
	return uint32(invokeId), IedClientError(err)
}
