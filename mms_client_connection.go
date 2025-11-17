package libiec61850go

/*
#include <stdlib.h>
#include "mms_client_connection.h"

extern void fMmsRawMessageHandlerGo(void* parameter, uint8_t* message, int messageLength, bool received);
extern void fMmsConnectionStateChangedHandlerGo(MmsConnection connection, void* parameter, MmsConnectionState newState);
extern void fMmsInformationReportHandlerGo(void* parameter, char* domainName, char* variableListName, MmsValue* value, bool isVariableListName);
extern void fMmsConnectionLostHandlerGo(MmsConnection connection, void* parameter);
*/
import "C"
import (
	"sync"
	"unsafe"
)

type MmsConnectionParameters struct {
	MaxServOutstandingCalling int
	MaxServOutstandingCalled  int
	DataStructureNestingLevel int
	MaxPduSize                int
	ServicesSupported         [11]uint8
}

func (x *MmsConnectionParameters) toC() *C.MmsConnectionParameters {
	servicesSupported := [11]C.uint8_t{}
	for i := range 11 {
		servicesSupported[i] = C.uint8_t(x.ServicesSupported[i])
	}
	return &C.MmsConnectionParameters{
		maxServOutstandingCalling: C.int(x.MaxServOutstandingCalling),
		maxServOutstandingCalled:  C.int(x.MaxServOutstandingCalled),
		dataStructureNestingLevel: C.int(x.DataStructureNestingLevel),
		maxPduSize:                C.int(x.MaxPduSize),
		servicesSupported:         servicesSupported,
	}
}

func (x *MmsConnectionParameters) fromC(in *C.MmsConnectionParameters) {
	x.MaxServOutstandingCalling = int(in.maxServOutstandingCalling)
	x.MaxServOutstandingCalled = int(in.maxServOutstandingCalled)
	x.DataStructureNestingLevel = int(in.dataStructureNestingLevel)
	x.MaxPduSize = int(in.maxPduSize)
	for i := range 11 {
		x.ServicesSupported[i] = uint8(in.servicesSupported[i])
	}
}

type MmsServerIdentity struct {
	VendorName string
	ModelName  string
	Revision   string
}

func (x *MmsServerIdentity) toC() *C.MmsServerIdentity {
	return &C.MmsServerIdentity{
		vendorName: C.CString(x.VendorName),
		modelName:  C.CString(x.ModelName),
		revision:   C.CString(x.Revision),
	}
}

func (x *MmsServerIdentity) fromC(in *C.MmsServerIdentity) {
	x.VendorName = C.GoString(in.vendorName)
	x.ModelName = C.GoString(in.modelName)
	x.Revision = C.GoString(in.revision)
}

type MmsConnectionState int32

const (
	MMS_CONNECTION_STATE_CLOSED     MmsConnectionState = C.MMS_CONNECTION_STATE_CLOSED
	MMS_CONNECTION_STATE_CONNECTING MmsConnectionState = C.MMS_CONNECTION_STATE_CONNECTING
	MMS_CONNECTION_STATE_CONNECTED  MmsConnectionState = C.MMS_CONNECTION_STATE_CONNECTED
	MMS_CONNECTION_STATE_CLOSING    MmsConnectionState = C.MMS_CONNECTION_STATE_CLOSING
)

type MmsInformationReportHandler func(parameter unsafe.Pointer, domainName string, variableListName string, value *MmsValue, isVariableListName bool)

var mapMmsInformationReportHandlers = sync.Map{}

//export fMmsInformationReportHandlerGo
func fMmsInformationReportHandlerGo(parameter unsafe.Pointer, domainName *C.char, variableListName *C.char, value *C.MmsValue, isVariableListName C._Bool) {
	mapMmsInformationReportHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(MmsInformationReportHandler); ok {
			fn(parameter, C.GoString(domainName), C.GoString(variableListName), &MmsValue{ctx: value}, bool(isVariableListName))
		}
		return true
	})
}

type MmsConnection struct {
	ctx C.MmsConnection
}

func MmsConnectionCreate() *MmsConnection {
	return &MmsConnection{ctx: C.MmsConnection_create()}
}

// func MmsConnectionCreateSecure(tlsConfig *TLSConfiguration) *MmsConnection {
// 	return &MmsConnection{ctx: C.MmsConnection_createSecure(tlsConfig.ctx)}
// }

// func MmsConnectionCreateNonThreaded(tlsConfig *TLSConfiguration) *MmsConnection {
// 	return &MmsConnection{ctx: C.MmsConnection_createNonThreaded(tlsConfig.ctx)}
// }

type MmsRawMessageHandler func(parameter unsafe.Pointer, message []byte, received bool)

var mapMmsRawMessageHandlers = sync.Map{}

//export fMmsRawMessageHandlerGo
func fMmsRawMessageHandlerGo(parameter unsafe.Pointer, message *C.uint8_t, messageLength C.int, received C._Bool) {
	mapMmsRawMessageHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(MmsRawMessageHandler); ok {
			fn(parameter, C.GoBytes(unsafe.Pointer(message), messageLength), bool(received))
		}
		return true
	})
}

func (x *MmsConnection) SetRawMessageHandler(handler MmsRawMessageHandler, parameter unsafe.Pointer) {
	C.MmsConnection_setRawMessageHandler(x.ctx, (C.MmsRawMessageHandler)(C.fMmsRawMessageHandlerGo), parameter)
}

func (x *MmsConnection) SetFilestoreBasepath(basepath string) {
	C.MmsConnection_setFilestoreBasepath(x.ctx, C.CString(basepath))
}

func (x *MmsConnection) SetRequestTimeout(timeoutInMs uint32) {
	C.MmsConnection_setRequestTimeout(x.ctx, C.uint32_t(timeoutInMs))
}

func (x *MmsConnection) SetMaxOutstandingCalls(calling int, called int) {
	C.MmsConnnection_setMaxOutstandingCalls(x.ctx, C.int(calling), C.int(called))
}

func (x *MmsConnection) GetRequestTimeout() uint32 {
	return uint32(C.MmsConnection_getRequestTimeout(x.ctx))
}

func (x *MmsConnection) SetConnectTimeout(timeoutInMs uint32) {
	C.MmsConnection_setConnectTimeout(x.ctx, C.uint32_t(timeoutInMs))
}

func (x *MmsConnection) SetInformationReportHandler(handler MmsInformationReportHandler, parameter unsafe.Pointer) {
	C.MmsConnection_setInformationReportHandler(x.ctx, (C.MmsInformationReportHandler)(C.fMmsInformationReportHandlerGo), parameter)
}

func (x *MmsConnection) GetIsoConnectionParameters() *IsoConnectionParameters {
	return &IsoConnectionParameters{ctx: C.MmsConnection_getIsoConnectionParameters(x.ctx)}
}

func (x *MmsConnection) GetMmsConnectionParameters() *MmsConnectionParameters {
	out := &MmsConnectionParameters{}
	params := C.MmsConnection_getMmsConnectionParameters(x.ctx)
	out.fromC(&params)
	return out
}

type MmsConnectionStateChangedHandler func(connection *MmsConnection, parameter unsafe.Pointer, newState MmsConnectionState)

var mapMmsConnectionStateChangedHandlers = sync.Map{}

//export fMmsConnectionStateChangedHandlerGo
func fMmsConnectionStateChangedHandlerGo(connection C.MmsConnection, parameter unsafe.Pointer, newState C.MmsConnectionState) {
	mapMmsConnectionStateChangedHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(MmsConnectionStateChangedHandler); ok {
			fn(&MmsConnection{ctx: connection}, parameter, MmsConnectionState(newState))
		}
		return true
	})
}

func (x *MmsConnection) SetConnectionStateChangedHandler(handler MmsConnectionStateChangedHandler, parameter unsafe.Pointer) {
	C.MmsConnection_setConnectionStateChangedHandler(x.ctx, (C.MmsConnectionStateChangedHandler)(C.fMmsConnectionStateChangedHandlerGo), parameter)
}

type MmsConnectionLostHandler func(connection *MmsConnection, parameter unsafe.Pointer)

var mapMmsConnectionLostHandlers = sync.Map{}

//export fMmsConnectionLostHandlerGo
func fMmsConnectionLostHandlerGo(connection C.MmsConnection, parameter unsafe.Pointer) {
	mapMmsConnectionLostHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(MmsConnectionLostHandler); ok {
			fn(&MmsConnection{ctx: connection}, parameter)
		}
		return true
	})
}

func (x *MmsConnection) SetConnectionLostHandler(handler MmsConnectionLostHandler, parameter unsafe.Pointer) {
	C.MmsConnection_setConnectionLostHandler(x.ctx, (C.MmsConnectionLostHandler)(C.fMmsConnectionLostHandlerGo), parameter)
}

func (x *MmsConnection) SetIsoConnectionParameters(params *IsoConnectionParameters) {
	C.MmsConnection_setIsoConnectionParameters(x.ctx, &params.ctx)
}

func (x *MmsConnection) Destroy() {
	C.MmsConnection_destroy(x.ctx)
}
