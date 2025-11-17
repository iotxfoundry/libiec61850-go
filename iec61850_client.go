package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_client.h"

extern void fIedConnectionClosedHandlerGo(void* parameter, IedConnection connection);
extern void fStateChangedHandlerGo(void* parameter, IedConnection connection, IedConnectionState newState);
extern void fGetGoCBValuesHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, ClientGooseControlBlock goCB);
extern void fSetGoCBValuesHandlerGo(uint32_t invokeId, void* parameter, IedClientError err);
extern void fReadObjectHandlerGo(uint32_t invokeId, void* parameter, IedClientError err, MmsValue* value);
*/
import "C"
import (
	"sync"
	"unsafe"

	"github.com/rs/xid"
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

// func (x *IedConnection) GetMmsConnection() *MmsConnection {
// 	return &MmsConnection{
// 		ctx: C.IedConnection_getMmsConnection(x.ctx),
// 	}
// }

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
	id := xid.New().String()
	mapGetGoCBValuesHandlers.Store(id, handler)
	defer mapGetGoCBValuesHandlers.Delete(id)
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

var mapSetGoCBValuesHandlers = sync.Map{}

//export fGenericServiceHandlerGo
func fGenericServiceHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, err C.IedClientError) {
	mapSetGoCBValuesHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(GenericServiceHandler); ok {
			fn(uint32(invokeId), parameter, IedClientError(err))
		}
		return true
	})
}

func (x *IedConnection) SetGoCBValuesAsync(goCB *ClientGooseControlBlock, parametersMask uint32, singleRequest bool, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := IED_ERROR_OK
	id := xid.New().String()
	mapSetGoCBValuesHandlers.Store(id, handler)
	defer mapSetGoCBValuesHandlers.Delete(id)
	out := C.IedConnection_setGoCBValuesAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), goCB.ctx, (C.uint32_t)(parametersMask), (C.bool)(singleRequest), (C.IedConnection_GenericServiceHandler)(C.fSetGoCBValuesHandlerGo), parameter)
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
	id := xid.New().String()
	mapReadObjectHandlers.Store(id, handler)
	defer mapReadObjectHandlers.Delete(id)
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
	id := xid.New().String()
	mapSetGoCBValuesHandlers.Store(id, handler)
	defer mapSetGoCBValuesHandlers.Delete(id)
	out := C.IedConnection_writeObjectAsync(x.ctx, (*C.IedClientError)(unsafe.Pointer(&err)), cref, (C.FunctionalConstraint)(fc), value.ctx, (C.IedConnection_GenericServiceHandler)(C.fSetGoCBValuesHandlerGo), parameter)
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
