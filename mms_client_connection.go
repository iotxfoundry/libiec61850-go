package libiec61850go

/*
#include <stdlib.h>
#include "mms_client_connection.h"

extern void fMmsRawMessageHandlerGo(void* parameter, uint8_t* message, int messageLength, bool received);
extern void fMmsConnectionStateChangedHandlerGo(MmsConnection connection, void* parameter, MmsConnectionState newState);
extern void fMmsInformationReportHandlerGo(void* parameter, char* domainName, char* variableListName, MmsValue* value, bool isVariableListName);
extern void fMmsConnectionLostHandlerGo(MmsConnection connection, void* parameter);
extern void fConcludeAbortHandlerGo(void* parameter, MmsError mmsError, bool success);
extern void fMmsConnectionGenericServiceHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, bool success);
extern void fMmsConnectionGetNameListHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, LinkedList nameList, bool moreFollows);
extern void fReadVariableHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, MmsValue* value);
extern void fWriteVariableHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, MmsDataAccessError accessError);
extern void fWriteMultipleVariablesHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, LinkedList accessResults);
extern void fGetVariableAccessAttributesHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, MmsVariableSpecification* spec);
extern void fReadNVLDirectoryHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, LinkedList specs, bool deletable);
extern void fIdentifyHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, char* vendorName, char* modelName, char* revision);
extern void fGetServerStatusHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, int vmdLogicalStatus, int vmdPhysicalStatus);
extern void fMmsFileDirectoryHandlerGo(void* parameter, char* filename, uint32_t size, uint64_t lastModified);
extern void fFileDirectoryHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, char* filename, uint32_t size, uint64_t lastModified, bool moreFollows);
extern void fMmsFileReadHandlerGo(void* parameter, int32_t frsmId, uint8_t* buffer, uint32_t bufferSize);
extern void fFileReadHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, int32_t frsmId, uint8_t* buffer, uint32_t byteReceived, bool moreFollows);
extern void fFileOpenHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, int32_t frsmId, uint32_t fileSize, uint64_t lastModified);
extern void fReadJournalHandlerGo(uint32_t invokeId, void* parameter, MmsError mmsError, LinkedList journalEntries, bool moreFollows);
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
	ctx *C.MmsServerIdentity
}

func (x *MmsServerIdentity) VendorName() string {
	return C.GoString(x.ctx.vendorName)
}

func (x *MmsServerIdentity) ModelName() string {
	return C.GoString(x.ctx.modelName)
}
func (x *MmsServerIdentity) Revision() string {
	return C.GoString(x.ctx.revision)
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

func MmsConnectionCreateSecure(tlsConfig *TLSConfiguration) *MmsConnection {
	return &MmsConnection{ctx: C.MmsConnection_createSecure(tlsConfig.ctx)}
}

func MmsConnectionCreateNonThreaded(tlsConfig *TLSConfiguration) *MmsConnection {
	return &MmsConnection{ctx: C.MmsConnection_createNonThreaded(tlsConfig.ctx)}
}

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

// func (x *MmsConnection) SetIsoConnectionParameters(params *IsoConnectionParameters) {
// 	C.MmsConnection_setIsoConnectionParameters(x.ctx, &params.ctx)
// }

func (x *MmsConnection) Destroy() {
	C.MmsConnection_destroy(x.ctx)
}

func (x *MmsConnection) Connect(serverName string, serverPort int) (bool, error) {
	err := MMS_ERROR_NONE
	success := bool(C.MmsConnection_connect(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(serverName), C.int(serverPort)))
	return success, err.Error()
}

func (x *MmsConnection) ConnectAsync(mmsError *MmsError, serverName string, serverPort int) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_connectAsync(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(serverName), C.int(serverPort))
	return err.Error()
}

func (x *MmsConnection) Tick() bool {
	return bool(C.MmsConnection_tick(x.ctx))
}

func (x *MmsConnection) SendRawData(buffer []byte) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_sendRawData(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), SliceData(buffer), C.int(len(buffer)))
	return err.Error()
}

func (x *MmsConnection) Close() {
	C.MmsConnection_close(x.ctx)
}

type ConcludeAbortHandler func(parameter unsafe.Pointer, mmsError MmsError, success bool)

var mapConcludeAbortHandlers = sync.Map{}

//export fConcludeAbortHandlerGo
func fConcludeAbortHandlerGo(parameter unsafe.Pointer, mmsError C.MmsError, success C.bool) {
	mapConcludeAbortHandlers.Range(func(key, value any) bool {
		if fn, ok := value.(ConcludeAbortHandler); ok {
			fn(parameter, MmsError(mmsError), bool(success))
		}
		return true
	})
}

func (x *MmsConnection) Abort() error {
	err := MMS_ERROR_NONE
	C.MmsConnection_abort(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)))
	return err.Error()
}

func (x *MmsConnection) AbortAsync() error {
	err := MMS_ERROR_NONE
	C.MmsConnection_abortAsync(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)))
	return err.Error()
}

func (x *MmsConnection) Conclude() error {
	err := MMS_ERROR_NONE
	C.MmsConnection_conclude(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)))
	return err.Error()
}

func (x *MmsConnection) ConcludeAsync(handler ConcludeAbortHandler, parameter unsafe.Pointer) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_concludeAsync(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), (C.MmsConnection_ConcludeAbortHandler)(C.fConcludeAbortHandlerGo), parameter)
	return err.Error()
}

type MmsConnectionGenericServiceHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, success bool)

var mapMmsConnectionGenericServiceHandlers = sync.Map{}

//export fMmsConnectionGenericServiceHandlerGo
func fMmsConnectionGenericServiceHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, success C.bool) {
	mapMmsConnectionGenericServiceHandlers.Range(func(key, value any) bool {
		if uint32(invokeId) == key.(uint32) {
			if fn, ok := value.(MmsConnectionGenericServiceHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), bool(success))
			}
			mapMmsConnectionGenericServiceHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

type MmsConnectionGetNameListHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, nameList *LinkedList, moreFollows bool)

var mapMmsConnectionGetNameListHandlers = sync.Map{}

//export fMmsConnectionGetNameListHandlerGo
func fMmsConnectionGetNameListHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, nameList C.LinkedList, moreFollows C.bool) {
	mapMmsConnectionGetNameListHandlers.Range(func(key, value any) bool {
		if uint32(invokeId) == key.(uint32) {
			if fn, ok := value.(MmsConnectionGetNameListHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), &LinkedList{ctx: nameList}, bool(moreFollows))
			}
			mapMmsConnectionGetNameListHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) GetVMDVariableNames() (*LinkedList, error) {
	err := MMS_ERROR_NONE
	nameList := C.MmsConnection_getVMDVariableNames(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)))
	return &LinkedList{ctx: nameList}, err.Error()
}

func (x *MmsConnection) GetVMDVariableNamesAsync(continueAfter string, handler MmsConnectionGetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getVMDVariableNamesAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(continueAfter), (C.MmsConnection_GetNameListHandler)(C.fMmsConnectionGetNameListHandlerGo), parameter)
	mapMmsConnectionGetNameListHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) GetDomainNames() (*LinkedList, error) {
	err := MMS_ERROR_NONE
	nameList := C.MmsConnection_getDomainNames(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)))
	return &LinkedList{ctx: nameList}, err.Error()
}

func (x *MmsConnection) GetDomainNamesAsync(continueAfter string, result *LinkedList, handler MmsConnectionGetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getDomainNamesAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(continueAfter), result.ctx, (C.MmsConnection_GetNameListHandler)(C.fMmsConnectionGetNameListHandlerGo), parameter)
	mapMmsConnectionGetNameListHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) GetDomainVariableNames(domainId string) (*LinkedList, error) {
	err := MMS_ERROR_NONE
	nameList := C.MmsConnection_getDomainVariableNames(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId))
	return &LinkedList{ctx: nameList}, err.Error()
}

func (x *MmsConnection) GetDomainVariableNamesAsync(domainId string, continueAfter string, result *LinkedList, handler MmsConnectionGetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getDomainVariableNamesAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(continueAfter), result.ctx, (C.MmsConnection_GetNameListHandler)(C.fMmsConnectionGetNameListHandlerGo), parameter)
	mapMmsConnectionGetNameListHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) GetDomainVariableListNames(domainId string) (*LinkedList, error) {
	err := MMS_ERROR_NONE
	nameList := C.MmsConnection_getDomainVariableListNames(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId))
	return &LinkedList{ctx: nameList}, err.Error()
}

func (x *MmsConnection) GetDomainVariableListNamesAsync(domainId string, continueAfter string, result *LinkedList, handler MmsConnectionGetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getDomainVariableListNamesAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(continueAfter), result.ctx, (C.MmsConnection_GetNameListHandler)(C.fMmsConnectionGetNameListHandlerGo), parameter)
	mapMmsConnectionGetNameListHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) GetDomainJournals(domainId string) (*LinkedList, error) {
	err := MMS_ERROR_NONE
	nameList := C.MmsConnection_getDomainJournals(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId))
	return &LinkedList{ctx: nameList}, err.Error()
}

func (x *MmsConnection) GetDomainJournalsAsync(domainId string, continueAfter string, handler MmsConnectionGetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getDomainJournalsAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(continueAfter), (C.MmsConnection_GetNameListHandler)(C.fMmsConnectionGetNameListHandlerGo), parameter)
	mapMmsConnectionGetNameListHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) GetVariableListNamesAssociationSpecific() (*LinkedList, error) {
	err := MMS_ERROR_NONE
	nameList := C.MmsConnection_getVariableListNamesAssociationSpecific(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)))
	return &LinkedList{ctx: nameList}, err.Error()
}

func (x *MmsConnection) GetVariableListNamesAssociationSpecificAsync(continueAfter string, handler MmsConnectionGetNameListHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getVariableListNamesAssociationSpecificAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(continueAfter), (C.MmsConnection_GetNameListHandler)(C.fMmsConnectionGetNameListHandlerGo), parameter)
	mapMmsConnectionGetNameListHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadVariable(domainId string, itemId string) (*MmsValue, error) {
	err := MMS_ERROR_NONE
	value := C.MmsConnection_readVariable(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId))
	return &MmsValue{ctx: value}, err.Error()
}

type ReadVariableHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, value *MmsValue)

var mapReadVariableHandlers = sync.Map{}

//export fReadVariableHandlerGo
func fReadVariableHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, value *C.MmsValue) {
	mapReadVariableHandlers.Range(func(k, v any) bool {
		if uint32(invokeId) == k.(uint32) {
			if fn, ok := v.(ReadVariableHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), &MmsValue{ctx: value})
			}
			mapReadVariableHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) ReadVariableAsync(domainId string, itemId string, handler ReadVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readVariableAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), (C.MmsConnection_ReadVariableHandler)(C.fReadVariableHandlerGo), parameter)
	mapReadVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadVariableComponent(domainId string, itemId string, componentId string) (*MmsValue, error) {
	err := MMS_ERROR_NONE
	value := C.MmsConnection_readVariableComponent(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), StringData(componentId))
	return &MmsValue{ctx: value}, err.Error()
}

func (x *MmsConnection) ReadVariableComponentAsync(domainId string, itemId string, componentId string, handler ReadVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readVariableComponentAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), StringData(componentId), (C.MmsConnection_ReadVariableHandler)(C.fReadVariableHandlerGo), parameter)
	mapReadVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadArrayElements(domainId string, itemId string, startIndex uint32, numberOfElements uint32) (*MmsValue, error) {
	err := MMS_ERROR_NONE
	value := C.MmsConnection_readArrayElements(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), C.uint32_t(startIndex), C.uint32_t(numberOfElements))
	return &MmsValue{ctx: value}, err.Error()
}

func (x *MmsConnection) ReadArrayElementsAsync(domainId string, itemId string, startIndex uint32, numberOfElements uint32, handler ReadVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readArrayElementsAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), C.uint32_t(startIndex), C.uint32_t(numberOfElements), (C.MmsConnection_ReadVariableHandler)(C.fReadVariableHandlerGo), parameter)
	mapReadVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadSingleArrayElementWithComponent(domainId string, itemId string, index uint32, componentId string) (*MmsValue, error) {
	err := MMS_ERROR_NONE
	value := C.MmsConnection_readSingleArrayElementWithComponent(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), C.uint32_t(index), StringData(componentId))
	return &MmsValue{ctx: value}, err.Error()
}

func (x *MmsConnection) ReadSingleArrayElementWithComponentAsync(domainId string, itemId string, index uint32, componentId string, handler ReadVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readSingleArrayElementWithComponentAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), C.uint32_t(index), StringData(componentId), (C.MmsConnection_ReadVariableHandler)(C.fReadVariableHandlerGo), parameter)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadMultipleVariables(domainId string, items []string) (*MmsValue, error) {
	err := MMS_ERROR_NONE
	ll := LinkedListCreate()
	for _, item := range items {
		ll.Add((unsafe.Pointer)(C.CString(item)))
	}
	value := C.MmsConnection_readMultipleVariables(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), ll.ctx)
	ll.DestroyDeep(func(data unsafe.Pointer) {
		C.free(data)
	})
	return &MmsValue{ctx: value}, err.Error()
}

func (x *MmsConnection) ReadMultipleVariablesAsync(domainId string, items []string, handler ReadVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	ll := LinkedListCreate()
	for _, item := range items {
		ll.Add((unsafe.Pointer)(C.CString(item)))
	}
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readMultipleVariablesAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), ll.ctx, (C.MmsConnection_ReadVariableHandler)(C.fReadVariableHandlerGo), parameter)
	mapReadVariableHandlers.Store(uint32(usedInvokeId), handler)
	ll.DestroyDeep(func(data unsafe.Pointer) {
		C.free(data)
	})
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) WriteVariable(domainId string, itemId string, value *MmsValue) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_writeVariable(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), value.ctx)
	return err.Error()
}

type WriteVariableHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, accessError MmsDataAccessError)

var mapWriteVariableHandlers = sync.Map{}

//export fWriteVariableHandlerGo
func fWriteVariableHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, accessError C.MmsDataAccessError) {
	mapWriteVariableHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(WriteVariableHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), MmsDataAccessError(accessError))
			}
			mapWriteVariableHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) WriteVariableAsync(domainId string, itemId string, value *MmsValue, handler WriteVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_writeVariableAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), value.ctx, (C.MmsConnection_WriteVariableHandler)(C.fWriteVariableHandlerGo), parameter)
	mapWriteVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) WriteVariableComponent(domainId string, itemId string, componentId string, value *MmsValue) (MmsDataAccessError, error) {
	err := MMS_ERROR_NONE
	code := C.MmsConnection_writeVariableComponent(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), StringData(componentId), value.ctx)
	return MmsDataAccessError(code), err.Error()
}

func (x *MmsConnection) WriteSingleArrayElementWithComponent(domainId string, itemId string, arrayIndex uint32, componentId string, value *MmsValue) (MmsDataAccessError, error) {
	err := MMS_ERROR_NONE
	code := C.MmsConnection_writeSingleArrayElementWithComponent(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), C.uint32_t(arrayIndex), StringData(componentId), value.ctx)
	return MmsDataAccessError(code), err.Error()
}

func (x *MmsConnection) WriteSingleArrayElementWithComponentAsync(domainId string, itemId string, arrayIndex uint32, componentId string, value *MmsValue, handler WriteVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_writeSingleArrayElementWithComponentAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), C.uint32_t(arrayIndex), StringData(componentId), value.ctx, (C.MmsConnection_WriteVariableHandler)(C.fWriteVariableHandlerGo), parameter)
	mapWriteVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) WriteVariableComponentAsync(domainId string, itemId string, componentId string, value *MmsValue, handler WriteVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_writeVariableComponentAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), StringData(componentId), value.ctx, (C.MmsConnection_WriteVariableHandler)(C.fWriteVariableHandlerGo), parameter)
	mapWriteVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) WriteArrayElements(domainId string, itemId string, index int, numberOfElements int, value *MmsValue) (MmsDataAccessError, error) {
	err := MMS_ERROR_NONE
	code := C.MmsConnection_writeArrayElements(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), C.int(index), C.int(numberOfElements), value.ctx)
	return MmsDataAccessError(code), err.Error()
}

func (x *MmsConnection) WriteArrayElementsAsync(domainId string, itemId string, index int, numberOfElements int, value *MmsValue, handler WriteVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_writeArrayElementsAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), C.int(index), C.int(numberOfElements), value.ctx, (C.MmsConnection_WriteVariableHandler)(C.fWriteVariableHandlerGo), parameter)
	mapWriteVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

type WriteMultipleVariablesHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, accessResults *LinkedList)

var mapWriteMultipleVariablesHandlers = sync.Map{}

//export fWriteMultipleVariablesHandlerGo
func fWriteMultipleVariablesHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, accessResults C.LinkedList) {
	mapWriteMultipleVariablesHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(WriteMultipleVariablesHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), &LinkedList{ctx: accessResults})
			}
			mapWriteMultipleVariablesHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) WriteMultipleVariables(domainId string, items *LinkedList, values *LinkedList, accessResults *LinkedList) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_writeMultipleVariables(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), items.ctx, values.ctx, (*C.LinkedList)(unsafe.Pointer(accessResults.ctx)))
	return err.Error()
}

func (x *MmsConnection) WriteMultipleVariablesAsync(domainId string, items *LinkedList, values *LinkedList, handler WriteMultipleVariablesHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_writeMultipleVariablesAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), items.ctx, values.ctx, (C.MmsConnection_WriteMultipleVariablesHandler)(C.fWriteMultipleVariablesHandlerGo), parameter)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) WriteNamedVariableList(isAssociationSpecific bool, domainId string, itemId string, values *LinkedList, accessResults *LinkedList) (*LinkedList, error) {
	err := MMS_ERROR_NONE
	var out C.LinkedList
	C.MmsConnection_writeNamedVariableList(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), C.bool(isAssociationSpecific), StringData(domainId), StringData(itemId), values.ctx, &out)
	return &LinkedList{ctx: out}, err.Error()
}

func (x *MmsConnection) WriteNamedVariableListAsync(isAssociationSpecific bool, domainId string, itemId string, values *LinkedList, handler WriteMultipleVariablesHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_writeNamedVariableListAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), C.bool(isAssociationSpecific), StringData(domainId), StringData(itemId), values.ctx, (C.MmsConnection_WriteMultipleVariablesHandler)(C.fWriteMultipleVariablesHandlerGo), parameter)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) GetVariableAccessAttributes(domainId string, itemId string) (*MmsVariableSpecification, error) {
	err := MMS_ERROR_NONE
	out := C.MmsConnection_getVariableAccessAttributes(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId))
	return &MmsVariableSpecification{ctx: out}, err.Error()
}

type GetVariableAccessAttributesHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, spec *MmsVariableSpecification)

var mapGetVariableAccessAttributesHandlers = sync.Map{}

//export fGetVariableAccessAttributesHandlerGo
func fGetVariableAccessAttributesHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, spec *C.MmsVariableSpecification) {
	mapGetVariableAccessAttributesHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(GetVariableAccessAttributesHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), &MmsVariableSpecification{ctx: spec})
			}
			mapGetVariableAccessAttributesHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) GetVariableAccessAttributesAsync(domainId string, itemId string, handler GetVariableAccessAttributesHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getVariableAccessAttributesAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), (C.MmsConnection_GetVariableAccessAttributesHandler)(C.fGetVariableAccessAttributesHandlerGo), parameter)
	mapGetVariableAccessAttributesHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadNamedVariableListValues(domainId string, listName string, specWithResult bool) (*MmsValue, error) {
	err := MMS_ERROR_NONE
	out := C.MmsConnection_readNamedVariableListValues(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(listName), C.bool(specWithResult))
	return &MmsValue{ctx: out}, err.Error()
}

func (x *MmsConnection) ReadNamedVariableListValuesAsync(domainId string, listName string, specWithResult bool, handler ReadVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readNamedVariableListValuesAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(listName), C.bool(specWithResult), (C.MmsConnection_ReadVariableHandler)(C.fReadVariableHandlerGo), parameter)
	mapReadVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadNamedVariableListValuesAssociationSpecific(listName string, specWithResult bool) (*MmsValue, error) {
	err := MMS_ERROR_NONE
	out := C.MmsConnection_readNamedVariableListValuesAssociationSpecific(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(listName), C.bool(specWithResult))
	return &MmsValue{ctx: out}, err.Error()
}

func (x *MmsConnection) ReadNamedVariableListValuesAssociationSpecificAsync(listName string, specWithResult bool, handler ReadVariableHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readNamedVariableListValuesAssociationSpecificAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(listName), C.bool(specWithResult), (C.MmsConnection_ReadVariableHandler)(C.fReadVariableHandlerGo), parameter)
	mapReadVariableHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) DefineNamedVariableList(domainId string, listName string, variableSpecs *LinkedList) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_defineNamedVariableList(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(listName), variableSpecs.ctx)
	return err.Error()
}

func (x *MmsConnection) DefineNamedVariableListAsync(domainId string, listName string, variableSpecs *LinkedList, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_defineNamedVariableListAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(listName), variableSpecs.ctx, (C.MmsConnection_GenericServiceHandler)(C.fMmsConnectionGenericServiceHandlerGo), parameter)
	mapMmsConnectionGenericServiceHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) DefineNamedVariableListAssociationSpecific(listName string, variableSpecs *LinkedList) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_defineNamedVariableListAssociationSpecific(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(listName), variableSpecs.ctx)
	return err.Error()
}

func (x *MmsConnection) DefineNamedVariableListAssociationSpecificAsync(listName string, variableSpecs *LinkedList, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_defineNamedVariableListAssociationSpecificAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(listName), variableSpecs.ctx, (C.MmsConnection_GenericServiceHandler)(C.fMmsConnectionGenericServiceHandlerGo), parameter)
	mapMmsConnectionGenericServiceHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadNamedVariableListDirectory(domainId string, listName string, deletable *bool) (*LinkedList, error) {
	err := MMS_ERROR_NONE
	out := C.MmsConnection_readNamedVariableListDirectory(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(listName), (*C.bool)(unsafe.Pointer(deletable)))
	return &LinkedList{ctx: out}, err.Error()
}

type ReadNVLDirectoryHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, specs *LinkedList, deletable bool)

var mapReadNVLDirectoryHandlers = sync.Map{}

//export fReadNVLDirectoryHandlerGo
func fReadNVLDirectoryHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, specs C.LinkedList, deletable C.bool) {
	mapReadNVLDirectoryHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(ReadNVLDirectoryHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), &LinkedList{ctx: specs}, bool(deletable))
			}
			mapReadNVLDirectoryHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) ReadNamedVariableListDirectoryAsync(domainId string, listName string, handler ReadNVLDirectoryHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readNamedVariableListDirectoryAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(listName), (C.MmsConnection_ReadNVLDirectoryHandler)(C.fReadNVLDirectoryHandlerGo), parameter)
	mapReadNVLDirectoryHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadNamedVariableListDirectoryAssociationSpecific(listName string, deletable *bool) (*LinkedList, error) {
	err := MMS_ERROR_NONE
	out := C.MmsConnection_readNamedVariableListDirectoryAssociationSpecific(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(listName), (*C.bool)(unsafe.Pointer(deletable)))
	return &LinkedList{ctx: out}, err.Error()
}

func (x *MmsConnection) ReadNamedVariableListDirectoryAssociationSpecificAsync(listName string, handler ReadNVLDirectoryHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readNamedVariableListDirectoryAssociationSpecificAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(listName), (C.MmsConnection_ReadNVLDirectoryHandler)(C.fReadNVLDirectoryHandlerGo), parameter)
	mapReadNVLDirectoryHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) DeleteNamedVariableList(domainId string, listName string) (bool, error) {
	err := MMS_ERROR_NONE
	success := C.MmsConnection_deleteNamedVariableList(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(listName))
	return bool(success), err.Error()
}

func (x *MmsConnection) DeleteNamedVariableListAsync(domainId string, listName string, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_deleteNamedVariableListAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(listName), (C.MmsConnection_GenericServiceHandler)(C.fMmsConnectionGenericServiceHandlerGo), parameter)
	mapMmsConnectionGenericServiceHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) DeleteAssociationSpecificNamedVariableList(listName string) (bool, error) {
	err := MMS_ERROR_NONE
	success := C.MmsConnection_deleteAssociationSpecificNamedVariableList(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(listName))
	return bool(success), err.Error()
}

func (x *MmsConnection) DeleteAssociationSpecificNamedVariableListAsync(listName string, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_deleteAssociationSpecificNamedVariableListAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(listName), (C.MmsConnection_GenericServiceHandler)(C.fMmsConnectionGenericServiceHandlerGo), parameter)
	mapMmsConnectionGenericServiceHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func VariableAccessSpecificationCreate(domainId string, itemId string) *MmsVariableAccessSpecification {
	return &MmsVariableAccessSpecification{
		ctx: C.MmsVariableAccessSpecification_create(StringData(domainId), StringData(itemId)),
	}
}

func VariableAccessSpecificationCreateAlternateAccess(domainId string, itemId string, index int32, componentName string) *MmsVariableAccessSpecification {
	return &MmsVariableAccessSpecification{
		ctx: C.MmsVariableAccessSpecification_createAlternateAccess(StringData(domainId), StringData(itemId), C.int32_t(index), StringData(componentName)),
	}
}

func (x *MmsVariableAccessSpecification) Destroy() {
	C.MmsVariableAccessSpecification_destroy(x.ctx)
}

func (x *MmsConnection) SetLocalDetail(localDetail int32) {
	C.MmsConnection_setLocalDetail(x.ctx, C.int32_t(localDetail))
}

func (x *MmsConnection) GetLocalDetail() int32 {
	return int32(C.MmsConnection_getLocalDetail(x.ctx))
}

func (x *MmsConnection) Identify() *MmsServerIdentity {
	err := MMS_ERROR_NONE
	cout := C.MmsConnection_identify(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)))
	return &MmsServerIdentity{
		ctx: cout,
	}
}

type IdentifyHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, vendorName string, modelName string, revision string)

var mapIdentifyHandlers = sync.Map{}

//export fIdentifyHandlerGo
func fIdentifyHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, vendorName *C.char, modelName *C.char, revision *C.char) {
	mapIdentifyHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(IdentifyHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), C.GoString(vendorName), C.GoString(modelName), C.GoString(revision))
			}
			mapIdentifyHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) IdentifyAsync(handler IdentifyHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_identifyAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), (C.MmsConnection_IdentifyHandler)(C.fIdentifyHandlerGo), parameter)
	mapIdentifyHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsServerIdentity) Destroy() {
	C.MmsServerIdentity_destroy(x.ctx)
}

func (x *MmsConnection) GetServerStatus(extendedDerivation bool) (int, int, error) {
	err := MMS_ERROR_NONE
	vmdLogicalStatus := C.int(0)
	vmdPhysicalStatus := C.int(0)
	C.MmsConnection_getServerStatus(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), (*C.int)(unsafe.Pointer(&vmdLogicalStatus)), (*C.int)(unsafe.Pointer(&vmdPhysicalStatus)), C.bool(extendedDerivation))
	return int(vmdLogicalStatus), int(vmdPhysicalStatus), err.Error()
}

type GetServerStatusHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, vmdLogicalStatus int, vmdPhysicalStatus int)

var mapGetServerStatusHandlers = sync.Map{}

//export fGetServerStatusHandlerGo
func fGetServerStatusHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, vmdLogicalStatus C.int, vmdPhysicalStatus C.int) {
	mapGetServerStatusHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(GetServerStatusHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), int(vmdLogicalStatus), int(vmdPhysicalStatus))
			}
			mapGetServerStatusHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) GetServerStatusAsync(extendedDerivation bool, handler GetServerStatusHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getServerStatusAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), C.bool(extendedDerivation), (C.MmsConnection_GetServerStatusHandler)(C.fGetServerStatusHandlerGo), parameter)
	mapGetServerStatusHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

type MmsFileDirectoryHandler func(parameter unsafe.Pointer, filename string, size uint32, lastModified uint64)

var mapMmsFileDirectoryHandlers = sync.Map{}

//export fMmsFileDirectoryHandlerGo
func fMmsFileDirectoryHandlerGo(parameter unsafe.Pointer, filename *C.char, size C.uint32_t, lastModified C.uint64_t) {
	mapMmsFileDirectoryHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(MmsFileDirectoryHandler); ok {
			fn(parameter, C.GoString(filename), uint32(size), uint64(lastModified))
		}
		return true
	})
}

type FileDirectoryHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, filename string, size uint32, lastModified uint64, moreFollows bool)

var mapFileDirectoryHandlers = sync.Map{}

//export fFileDirectoryHandlerGo
func fFileDirectoryHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, filename *C.char, size C.uint32_t, lastModified C.uint64_t, moreFollows C.bool) {
	mapFileDirectoryHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(FileDirectoryHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), C.GoString(filename), uint32(size), uint64(lastModified), bool(moreFollows))
			}
			mapFileDirectoryHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

type MmsFileReadHandler func(parameter unsafe.Pointer, frsmId int32, buffer []byte)

var mapMmsFileReadHandlers = sync.Map{}

//export fMmsFileReadHandlerGo
func fMmsFileReadHandlerGo(parameter unsafe.Pointer, frsmId C.int32_t, buffer *C.uint8_t, bytesReceived C.uint32_t) {
	mapMmsFileReadHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(MmsFileReadHandler); ok {
			fn(parameter, int32(frsmId), C.GoBytes(unsafe.Pointer(buffer), C.int(bytesReceived)))
		}
		return true
	})
}

type FileReadHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, frsmId int32, buffer []byte, moreFollows bool)

var mapFileReadHandlers = sync.Map{}

//export fFileReadHandlerGo
func fFileReadHandlerGo(invokeId uint32, parameter unsafe.Pointer, mmsError C.MmsError, frsmId C.int32_t, buffer *C.uint8_t, byteReceived C.uint32_t, moreFollows C.bool) {
	mapFileReadHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(FileReadHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), int32(frsmId), C.GoBytes(unsafe.Pointer(buffer), C.int(byteReceived)), bool(moreFollows))
			}
			mapFileReadHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) FileOpen(filename string, initialPosition uint32) (uint32, uint32, uint64, error) {
	err := MMS_ERROR_NONE
	fileSize := C.uint32_t(0)
	lastModified := C.uint64_t(0)
	frsmId := C.MmsConnection_fileOpen(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(filename), C.uint32_t(initialPosition), (*C.uint32_t)(unsafe.Pointer(&fileSize)), (*C.uint64_t)(unsafe.Pointer(&lastModified)))
	return uint32(frsmId), uint32(fileSize), uint64(lastModified), err.Error()
}

type FileOpenHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, frsmId int32, fileSize uint32, lastModified uint64)

var mapFileOpenHandlers = sync.Map{}

//export fFileOpenHandlerGo
func fFileOpenHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, frsmId C.int32_t, fileSize C.uint32_t, lastModified C.uint64_t) {
	mapFileOpenHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(FileOpenHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), int32(frsmId), uint32(fileSize), uint64(lastModified))
			}
			mapFileOpenHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) FileOpenAsync(filename string, initialPosition uint32, handler FileOpenHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_fileOpenAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(filename), C.uint32_t(initialPosition), (C.MmsConnection_FileOpenHandler)(C.fFileOpenHandlerGo), parameter)
	mapFileOpenHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) FileRead(frsmId int32, handler MmsFileReadHandler, parameter unsafe.Pointer) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_fileRead(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), C.int32_t(frsmId), (C.MmsConnection_FileReadHandler)(C.fFileReadHandlerGo), parameter)
	mapMmsFileReadHandlers.Store(x.ctx, handler)
	return err.Error()
}

func (x *MmsConnection) FileReadAsync(frsmId int32, handler FileReadHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_fileReadAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), C.int32_t(frsmId), (C.MmsConnection_FileReadHandler)(C.fFileReadHandlerGo), parameter)
	mapFileReadHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) FileClose(frsmId int32) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_fileClose(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), C.int32_t(frsmId))
	return err.Error()
}

func (x *MmsConnection) FileCloseAsync(frsmId int32, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_fileCloseAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), C.uint32_t(frsmId), (C.MmsConnection_GenericServiceHandler)(C.fMmsConnectionGenericServiceHandlerGo), parameter)
	mapMmsConnectionGenericServiceHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) FileDelete(filename string) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_fileDelete(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(filename))
	return err.Error()
}

func (x *MmsConnection) FileDeleteAsync(filename string, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_fileDeleteAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(filename), (C.MmsConnection_GenericServiceHandler)(C.fMmsConnectionGenericServiceHandlerGo), parameter)
	mapMmsConnectionGenericServiceHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) FileRename(currentFileName string, newFileName string) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_fileRename(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(currentFileName), StringData(newFileName))
	return err.Error()
}

func (x *MmsConnection) FileRenameAsync(currentFileName string, newFileName string, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_fileRenameAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(currentFileName), StringData(newFileName), (C.MmsConnection_GenericServiceHandler)(C.fMmsConnectionGenericServiceHandlerGo), parameter)
	mapMmsConnectionGenericServiceHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ObtainFile(sourceFile string, destinationFile string) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_obtainFile(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(sourceFile), StringData(destinationFile))
	return err.Error()
}

func (x *MmsConnection) ObtainFileAsync(sourceFile string, destinationFile string, handler GenericServiceHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_obtainFileAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(sourceFile), StringData(destinationFile), (C.MmsConnection_GenericServiceHandler)(C.fMmsConnectionGenericServiceHandlerGo), parameter)
	mapMmsConnectionGenericServiceHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) GetFileDirectory(fileSpecification string, continueAfter string, handler FileDirectoryHandler, parameter unsafe.Pointer) error {
	err := MMS_ERROR_NONE
	C.MmsConnection_getFileDirectory(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(fileSpecification), StringData(continueAfter), (C.MmsConnection_FileDirectoryHandler)(C.fFileDirectoryHandlerGo), parameter)
	mapMmsFileDirectoryHandlers.Store(x.ctx, handler)
	return err.Error()
}

func (x *MmsConnection) GetFileDirectoryAsync(fileSpecification string, continueAfter string, handler FileDirectoryHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_getFileDirectoryAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(fileSpecification), StringData(continueAfter), (C.MmsConnection_FileDirectoryHandler)(C.fFileDirectoryHandlerGo), parameter)
	mapFileDirectoryHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

type MmsJournalEntry struct {
	ctx C.MmsJournalEntry
}

type MmsJournalVariable struct {
	ctx C.MmsJournalVariable
}

func (x *MmsJournalEntry) Destroy() {
	C.MmsJournalEntry_destroy(x.ctx)
}

func (x *MmsJournalEntry) GetEntryID() *MmsValue {
	return &MmsValue{ctx: C.MmsJournalEntry_getEntryID(x.ctx)}
}

func (x *MmsJournalEntry) GetOccurenceTime() *MmsValue {
	return &MmsValue{ctx: C.MmsJournalEntry_getOccurenceTime(x.ctx)}
}

func (x *MmsJournalEntry) GetJournalVariables() *LinkedList {
	return &LinkedList{ctx: C.MmsJournalEntry_getJournalVariables(x.ctx)}
}

func (x *MmsJournalVariable) GetTag() string {
	return C.GoString(C.MmsJournalVariable_getTag(x.ctx))
}

func (x *MmsJournalVariable) GetValue() *MmsValue {
	return &MmsValue{ctx: C.MmsJournalVariable_getValue(x.ctx)}
}

type ReadJournalHandler func(invokeId uint32, parameter unsafe.Pointer, mmsError error, journalEntries *LinkedList, moreFollows bool)

var mapReadJournalHandlers = sync.Map{}

//export fReadJournalHandlerGo
func fReadJournalHandlerGo(invokeId C.uint32_t, parameter unsafe.Pointer, mmsError C.MmsError, journalEntries C.LinkedList, moreFollows C.bool) {
	mapReadJournalHandlers.Range(func(k, v any) bool {
		if k.(uint32) == uint32(invokeId) {
			if fn, ok := v.(ReadJournalHandler); ok {
				fn(uint32(invokeId), parameter, MmsError(mmsError).Error(), &LinkedList{ctx: journalEntries}, bool(moreFollows))
			}
			mapReadJournalHandlers.Delete(uint32(invokeId))
			return false
		}
		return true
	})
}

func (x *MmsConnection) ReadJournalTimeRange(domainId string, itemId string, startTime *MmsValue, endTime *MmsValue) (*LinkedList, bool, error) {
	err := MMS_ERROR_NONE
	moreFollows := C.bool(false)
	list := C.MmsConnection_readJournalTimeRange(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), startTime.ctx, endTime.ctx, (*C.bool)(unsafe.Pointer(&moreFollows)))
	return &LinkedList{ctx: list}, bool(moreFollows), err.Error()
}

func (x *MmsConnection) ReadJournalTimeRangeAsync(domainId string, itemId string, startTime *MmsValue, endTime *MmsValue, handler ReadJournalHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readJournalTimeRangeAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), startTime.ctx, endTime.ctx, (C.MmsConnection_ReadJournalHandler)(C.fReadJournalHandlerGo), parameter)
	mapReadJournalHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}

func (x *MmsConnection) ReadJournalStartAfter(domainId string, itemId string, timeSpecification *MmsValue, entrySpecification *MmsValue) (*LinkedList, bool, error) {
	err := MMS_ERROR_NONE
	moreFollows := C.bool(false)
	list := C.MmsConnection_readJournalStartAfter(x.ctx, (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), timeSpecification.ctx, entrySpecification.ctx, (*C.bool)(unsafe.Pointer(&moreFollows)))
	return &LinkedList{ctx: list}, bool(moreFollows), err.Error()
}

func (x *MmsConnection) ReadJournalStartAfterAsync(domainId string, itemId string, timeSpecification *MmsValue, entrySpecification *MmsValue, handler ReadJournalHandler, parameter unsafe.Pointer) (uint32, error) {
	err := MMS_ERROR_NONE
	usedInvokeId := C.uint32_t(0)
	C.MmsConnection_readJournalStartAfterAsync(x.ctx, (*C.uint32_t)(unsafe.Pointer(&usedInvokeId)), (*C.MmsError)(unsafe.Pointer(&err)), StringData(domainId), StringData(itemId), timeSpecification.ctx, entrySpecification.ctx, (C.MmsConnection_ReadJournalHandler)(C.fReadJournalHandlerGo), parameter)
	mapReadJournalHandlers.Store(uint32(usedInvokeId), handler)
	return uint32(usedInvokeId), err.Error()
}
