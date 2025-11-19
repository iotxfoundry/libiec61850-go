package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_server.h"

extern bool fAcseAuthenticatorGo(void* parameter, AcseAuthenticationParameter authParameter, void** securityToken, IsoApplicationReference* appReference);
extern void fIedConnectionIndicationHandlerGo(IedServer self, ClientConnection connection, bool connected, void* parameter);
extern bool fActiveSettingGroupChangedHandlerGo(void* parameter, SettingGroupControlBlock* sgcb, uint8_t newActSg, ClientConnection connection);
extern bool fEditSettingGroupChangedHandlerGo(void* parameter, SettingGroupControlBlock* sgcb, uint8_t newEditSg, ClientConnection connection);
extern void fEditSettingGroupConfirmationHandlerGo(void* parameter, SettingGroupControlBlock* sgcb, uint8_t editSg);
extern CheckHandlerResult fControlPerformCheckHandlerGo(ControlAction action, void* parameter, MmsValue* ctlVal, bool test, bool interlockCheck);
extern ControlHandlerResult fControlWaitForExecutionHandlerGo(ControlAction action, void* parameter, MmsValue* ctlVal, bool test, bool interlockCheck);
extern ControlHandlerResult fControlHandlerGo(ControlAction action, void* parameter, MmsValue* ctlVal, bool test);
extern void fControlSelectStateChangedHandlerGo(ControlAction action, void* parameter, bool isSelected, SelectStateChangedReason reason);
extern void fRCBEventHandlerGo(void* parameter, ReportControlBlock* rcb, ClientConnection connection, IedServer_RCBEventType event, char* parameterName, MmsDataAccessError serviceError);
extern void fSVCBEventHandlerGo(SVControlBlock* svcb, int event, void* parameter);
extern void fGoCBEventHandlerGo(MmsGooseControlBlock goCb, int event, void* parameter);
extern MmsDataAccessError fWriteAccessHandlerGo(DataAttribute* dataAttribute, MmsValue* value, ClientConnection connection, void* parameter);
extern MmsDataAccessError fReadAccessHandlerGo (LogicalDevice* ld, LogicalNode* ln, DataObject* dataObject, FunctionalConstraint fc, ClientConnection connection, void* parameter);
extern bool fDataSetAccessHandlerGo(void* parameter, ClientConnection connection, IedServer_DataSetOperation operation, char* datasetRef);
extern bool fListObjectsAccessHandlerGo(void* parameter, ClientConnection connection, ACSIClass acsiClass, LogicalDevice* ld, LogicalNode* ln, char* objectName, char* subObjectName, FunctionalConstraint fc);
extern bool fControlBlockAccessHandlerGo(void* parameter, ClientConnection connection, ACSIClass acsiClass, LogicalDevice* ld, LogicalNode* ln, char* objectName, char* subObjectName, IedServer_ControlBlockAccessType accessType);
extern bool fDirectoryAccessHandlerGo(void* parameter, ClientConnection connection, IedServer_DirectoryCategory category, LogicalDevice* logicalDevice);
*/
import "C"
import (
	"sync"
	"unsafe"
)

const (
	IEC61850_REPORTSETTINGS_RPT_ID     = C.IEC61850_REPORTSETTINGS_RPT_ID
	IEC61850_REPORTSETTINGS_BUF_TIME   = C.IEC61850_REPORTSETTINGS_BUF_TIME
	IEC61850_REPORTSETTINGS_DATSET     = C.IEC61850_REPORTSETTINGS_DATSET
	IEC61850_REPORTSETTINGS_TRG_OPS    = C.IEC61850_REPORTSETTINGS_TRG_OPS
	IEC61850_REPORTSETTINGS_OPT_FIELDS = C.IEC61850_REPORTSETTINGS_OPT_FIELDS
	IEC61850_REPORTSETTINGS_INTG_PD    = C.IEC61850_REPORTSETTINGS_INTG_PD
)

type IedServerConfig struct {
	ctx C.IedServerConfig
}

func IedServerConfigCreate() *IedServerConfig {
	return &IedServerConfig{ctx: C.IedServerConfig_create()}
}

func (x *IedServerConfig) Destroy() {
	C.IedServerConfig_destroy(x.ctx)
}

func (x *IedServerConfig) SetEdition(edition uint8) {
	C.IedServerConfig_setEdition(x.ctx, C.uint8_t(edition))
}

func (x *IedServerConfig) GetEdition() uint8 {
	return uint8(C.IedServerConfig_getEdition(x.ctx))
}

func (x *IedServerConfig) SetReportBufferSize(reportBufferSize int) {
	C.IedServerConfig_setReportBufferSize(x.ctx, C.int(reportBufferSize))
}

func (x *IedServerConfig) GetReportBufferSize() int {
	return int(C.IedServerConfig_getReportBufferSize(x.ctx))
}

func (x *IedServerConfig) SetReportBufferSizeForURCBs(reportBufferSize int) {
	C.IedServerConfig_setReportBufferSizeForURCBs(x.ctx, C.int(reportBufferSize))
}

func (x *IedServerConfig) GetReportBufferSizeForURCBs() int {
	return int(C.IedServerConfig_getReportBufferSizeForURCBs(x.ctx))
}

func (x *IedServerConfig) SetMaxMmsConnections(maxConnections int) {
	C.IedServerConfig_setMaxMmsConnections(x.ctx, C.int(maxConnections))
}

func (x *IedServerConfig) GetMaxMmsConnections() int {
	return int(C.IedServerConfig_getMaxMmsConnections(x.ctx))
}

func (x *IedServerConfig) SetSyncIntegrityReportTimes(enable bool) {
	C.IedServerConfig_setSyncIntegrityReportTimes(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) GetSyncIntegrityReportTimes() bool {
	return bool(C.IedServerConfig_getSyncIntegrityReportTimes(x.ctx))
}

func (x *IedServerConfig) SetFileServiceBasePath(basepath string) {
	C.IedServerConfig_setFileServiceBasePath(x.ctx, StringData(basepath))
}

func (x *IedServerConfig) GetFileServiceBasePath() string {
	return C.GoString(C.IedServerConfig_getFileServiceBasePath(x.ctx))
}

func (x *IedServerConfig) EnableFileService(enable bool) {
	C.IedServerConfig_enableFileService(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) IsFileServiceEnabled() bool {
	return bool(C.IedServerConfig_isFileServiceEnabled(x.ctx))
}

func (x *IedServerConfig) EnableDynamicDataSetService(enable bool) {
	C.IedServerConfig_enableDynamicDataSetService(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) IsDynamicDataSetServiceEnabled() bool {
	return bool(C.IedServerConfig_isDynamicDataSetServiceEnabled(x.ctx))
}

func (x *IedServerConfig) SetMaxAssociationSpecificDataSets(maxDataSets int) {
	C.IedServerConfig_setMaxAssociationSpecificDataSets(x.ctx, C.int(maxDataSets))
}

func (x *IedServerConfig) GetMaxAssociationSpecificDataSets() int {
	return int(C.IedServerConfig_getMaxAssociationSpecificDataSets(x.ctx))
}

func (x *IedServerConfig) SetMaxDomainSpecificDataSets(maxDataSets int) {
	C.IedServerConfig_setMaxDomainSpecificDataSets(x.ctx, C.int(maxDataSets))
}

func (x *IedServerConfig) GetMaxDomainSpecificDataSets() int {
	return int(C.IedServerConfig_getMaxDomainSpecificDataSets(x.ctx))
}

func (x *IedServerConfig) SetMaxDataSetEntries(maxDataSetEntries int) {
	C.IedServerConfig_setMaxDataSetEntries(x.ctx, C.int(maxDataSetEntries))
}

func (x *IedServerConfig) GetMaxDatasSetEntries() int {
	return int(C.IedServerConfig_getMaxDatasSetEntries(x.ctx))
}

func (x *IedServerConfig) EnableLogService(enable bool) {
	C.IedServerConfig_enableLogService(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) EnableEditSG(enable bool) {
	C.IedServerConfig_enableEditSG(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) EnableResvTmsForSGCB(enable bool) {
	C.IedServerConfig_enableResvTmsForSGCB(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) EnableResvTmsForBRCB(enable bool) {
	C.IedServerConfig_enableResvTmsForBRCB(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) IsResvTmsForBRCBEnabled() bool {
	return bool(C.IedServerConfig_isResvTmsForBRCBEnabled(x.ctx))
}

func (x *IedServerConfig) EnableOwnerForRCB(enable bool) {
	C.IedServerConfig_enableOwnerForRCB(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) IsOwnerForRCBEnabled() bool {
	return bool(C.IedServerConfig_isOwnerForRCBEnabled(x.ctx))
}

func (x *IedServerConfig) UseIntegratedGoosePublisher(enable bool) {
	C.IedServerConfig_useIntegratedGoosePublisher(x.ctx, C.bool(enable))
}

func (x *IedServerConfig) IsLogServiceEnabled() bool {
	return bool(C.IedServerConfig_isLogServiceEnabled(x.ctx))
}

func (x *IedServerConfig) SetReportSetting(setting uint8, isDyn bool) {
	C.IedServerConfig_setReportSetting(x.ctx, C.uint8_t(setting), C.bool(isDyn))
}

func (x *IedServerConfig) GetReportSetting(setting uint8) bool {
	return bool(C.IedServerConfig_getReportSetting(x.ctx, C.uint8_t(setting)))
}

type IedServer struct {
	ctx C.IedServer
}

type ClientConnection struct {
	ctx C.ClientConnection
}

func IedServerCreate(dataModel *IedModel) *IedServer {
	return &IedServer{ctx: C.IedServer_create(dataModel.ctx)}
}

func IedServerCreateWithTlsSupport(dataModel *IedModel, tlsConfiguration *TLSConfiguration) *IedServer {
	return &IedServer{ctx: C.IedServer_createWithTlsSupport(dataModel.ctx, tlsConfiguration.ctx)}
}

func IedServerCreateWithConfig(dataModel *IedModel, tlsConfiguration *TLSConfiguration, serverConfiguration *IedServerConfig) *IedServer {
	return &IedServer{ctx: C.IedServer_createWithConfig(dataModel.ctx, tlsConfiguration.ctx, serverConfiguration.ctx)}
}

func (x *IedServer) Destroy() {
	C.IedServer_destroy(x.ctx)
}

func (x *IedServer) AddAccessPoint(ipAddr string, tcpPort int, tlsConfiguration *TLSConfiguration) bool {
	return bool(C.IedServer_addAccessPoint(x.ctx, StringData(ipAddr), C.int(tcpPort), tlsConfiguration.ctx))
}

func (x *IedServer) SetLocalIpAddress(localIpAddress string) {
	C.IedServer_setLocalIpAddress(x.ctx, StringData(localIpAddress))
}

func (x *IedServer) SetServerIdentity(vendor string, model string, revision string) {
	C.IedServer_setServerIdentity(x.ctx, StringData(vendor), StringData(model), StringData(revision))
}

func (x *IedServer) SetFilestoreBasepath(basepath string) {
	C.IedServer_setFilestoreBasepath(x.ctx, StringData(basepath))
}

func (x *IedServer) SetLogStorage(logRef string, logStorage *LogStorage) {
	C.IedServer_setLogStorage(x.ctx, StringData(logRef), logStorage.ctx)
}

func (x *IedServer) Start(tcpPort int) {
	C.IedServer_start(x.ctx, C.int(tcpPort))
}

func (x *IedServer) Stop() {
	C.IedServer_stop(x.ctx)
}

func (x *IedServer) StartThreadless(tcpPort int) {
	C.IedServer_startThreadless(x.ctx, C.int(tcpPort))
}

func (x *IedServer) WaitReady(timeoutMs uint) int {
	return int(C.IedServer_waitReady(x.ctx, C.uint(timeoutMs)))
}

func (x *IedServer) ProcessIncomingData() {
	C.IedServer_processIncomingData(x.ctx)
}

func (x *IedServer) PerformPeriodicTasks() {
	C.IedServer_performPeriodicTasks(x.ctx)
}

func (x *IedServer) StopThreadless() {
	C.IedServer_stopThreadless(x.ctx)
}

func (x *IedServer) GetDataModel() *IedModel {
	return &IedModel{ctx: C.IedServer_getDataModel(x.ctx)}
}

func (x *IedServer) IsRunning() bool {
	return bool(C.IedServer_isRunning(x.ctx))
}

func (x *IedServer) GetNumberOfOpenConnections() int {
	return int(C.IedServer_getNumberOfOpenConnections(x.ctx))
}

func (x *IedServer) GetMmsServer() *MmsServer {
	return &MmsServer{ctx: C.IedServer_getMmsServer(x.ctx)}
}

func (x *IedServer) EnableGoosePublishing() {
	C.IedServer_enableGoosePublishing(x.ctx)
}

func (x *IedServer) DisableGoosePublishing() {
	C.IedServer_disableGoosePublishing(x.ctx)
}

func (x *IedServer) SetGooseInterfaceId(interfaceId string) {
	C.IedServer_setGooseInterfaceId(x.ctx, StringData(interfaceId))
}

func (x *IedServer) SetGooseInterfaceIdEx(ln *LogicalNode, gcbName string, interfaceId string) {
	C.IedServer_setGooseInterfaceIdEx(x.ctx, ln.ctx, StringData(gcbName), StringData(interfaceId))
}

func (x *IedServer) UseGooseVlanTag(ln *LogicalNode, gcbName string, useVlanTag bool) {
	C.IedServer_useGooseVlanTag(x.ctx, ln.ctx, StringData(gcbName), C.bool(useVlanTag))
}

func (x *IedServer) SetTimeQuality(leapSecondKnown bool, clockFailure bool, clockNotSynchronized bool, subsecondPrecision int) {
	C.IedServer_setTimeQuality(x.ctx, C.bool(leapSecondKnown), C.bool(clockFailure), C.bool(clockNotSynchronized), C.int(subsecondPrecision))
}

func (x *IedServer) SetAuthenticator(authenticator AcseAuthenticator, authenticatorParameter unsafe.Pointer) {
	C.IedServer_setAuthenticator(x.ctx, C.AcseAuthenticator(C.fAcseAuthenticatorGo), authenticatorParameter)
}

func (x *ClientConnection) GetPeerAddress() string {
	return C.GoString(C.ClientConnection_getPeerAddress(x.ctx))
}

func (x *ClientConnection) GetLocalAddress() string {
	return C.GoString(C.ClientConnection_getLocalAddress(x.ctx))
}

func (x *ClientConnection) GetSecurityToken() unsafe.Pointer {
	return C.ClientConnection_getSecurityToken(x.ctx)
}

func (x *ClientConnection) Abort() bool {
	return bool(C.ClientConnection_abort(x.ctx))
}

func (x *ClientConnection) ClaimOwnership() *ClientConnection {
	return &ClientConnection{ctx: C.ClientConnection_claimOwnership(x.ctx)}
}

func (x *ClientConnection) Release() {
	C.ClientConnection_release(x.ctx)
}

type IedConnectionIndicationHandler func(self *IedServer, connection *ClientConnection, connected bool, parameter unsafe.Pointer)

var mapIedConnectionIndicationHandlers = sync.Map{}

//export fIedConnectionIndicationHandlerGo
func fIedConnectionIndicationHandlerGo(self C.IedServer, connection C.ClientConnection, connected C._Bool, parameter unsafe.Pointer) {
	mapIedConnectionIndicationHandlers.Range(func(k, v any) bool {
		if k.(C.IedServer) == self {
			if fn, ok := v.(IedConnectionIndicationHandler); ok {
				fn(&IedServer{ctx: self}, &ClientConnection{ctx: connection}, bool(connected), parameter)
			}
		}
		return true
	})
}

func (x *IedServer) SetConnectionIndicationHandler(handler IedConnectionIndicationHandler, parameter unsafe.Pointer) {
	mapIedConnectionIndicationHandlers.Store(x.ctx, handler)
	C.IedServer_setConnectionIndicationHandler(x.ctx, C.IedConnectionIndicationHandler(C.fIedConnectionIndicationHandlerGo), parameter)
}

func (x *IedServer) IgnoreClientRequests(enable bool) {
	C.IedServer_ignoreClientRequests(x.ctx, C.bool(enable))
}

func (x *IedServer) LockDataModel() {
	C.IedServer_lockDataModel(x.ctx)
}

func (x *IedServer) UnlockDataModel() {
	C.IedServer_unlockDataModel(x.ctx)
}

func (x *IedServer) GetAttributeValue(da *DataAttribute) *MmsValue {
	return &MmsValue{ctx: C.IedServer_getAttributeValue(x.ctx, da.ctx)}
}

func (x *IedServer) GetBooleanAttributeValue(da *DataAttribute) bool {
	return bool(C.IedServer_getBooleanAttributeValue(x.ctx, da.ctx))
}

func (x *IedServer) GetInt32AttributeValue(da *DataAttribute) int32 {
	return int32(C.IedServer_getInt32AttributeValue(x.ctx, da.ctx))
}

func (x *IedServer) GetInt64AttributeValue(da *DataAttribute) int64 {
	return int64(C.IedServer_getInt64AttributeValue(x.ctx, da.ctx))
}

func (x *IedServer) GetUInt32AttributeValue(da *DataAttribute) uint32 {
	return uint32(C.IedServer_getUInt32AttributeValue(x.ctx, da.ctx))
}

func (x *IedServer) GetFloatAttributeValue(da *DataAttribute) float32 {
	return float32(C.IedServer_getFloatAttributeValue(x.ctx, da.ctx))
}

func (x *IedServer) GetUTCTimeAttributeValue(da *DataAttribute) uint64 {
	return uint64(C.IedServer_getUTCTimeAttributeValue(x.ctx, da.ctx))
}

func (x *IedServer) GetBitStringAttributeValue(da *DataAttribute) uint32 {
	return uint32(C.IedServer_getBitStringAttributeValue(x.ctx, da.ctx))
}

func (x *IedServer) GetStringAttributeValue(da *DataAttribute) string {
	return C.GoString(C.IedServer_getStringAttributeValue(x.ctx, da.ctx))
}

func (x *IedServer) GetFunctionalConstrainedData(do *DataObject, fc FunctionalConstraint) *MmsValue {
	return &MmsValue{ctx: C.IedServer_getFunctionalConstrainedData(x.ctx, do.ctx, C.FunctionalConstraint(fc))}
}

func (x *IedServer) UpdateAttributeValue(da *DataAttribute, value *MmsValue) {
	C.IedServer_updateAttributeValue(x.ctx, da.ctx, value.ctx)
}

func (x *IedServer) UpdateFloatAttributeValue(da *DataAttribute, value float32) {
	C.IedServer_updateFloatAttributeValue(x.ctx, da.ctx, C.float(value))
}

func (x *IedServer) UpdateInt32AttributeValue(da *DataAttribute, value int32) {
	C.IedServer_updateInt32AttributeValue(x.ctx, da.ctx, C.int32_t(value))
}

func (x *IedServer) UpdateDbposAttributeValue(da *DataAttribute, value Dbpos) {
	C.IedServer_updateDbposValue(x.ctx, da.ctx, C.Dbpos(value))
}

func (x *IedServer) UpdateInt64AttributeValue(da *DataAttribute, value int64) {
	C.IedServer_updateInt64AttributeValue(x.ctx, da.ctx, C.int64_t(value))
}

func (x *IedServer) UpdateUInt32AttributeValue(da *DataAttribute, value uint32) {
	C.IedServer_updateUnsignedAttributeValue(x.ctx, da.ctx, C.uint32_t(value))
}

func (x *IedServer) UpdateBitStringAttributeValue(da *DataAttribute, value uint32) {
	C.IedServer_updateBitStringAttributeValue(x.ctx, da.ctx, C.uint32_t(value))
}

func (x *IedServer) UpdateBooleanAttributeValue(da *DataAttribute, value bool) {
	C.IedServer_updateBooleanAttributeValue(x.ctx, da.ctx, C._Bool(value))
}

func (x *IedServer) UpdateVisibleStringAttributeValue(da *DataAttribute, value string) {
	C.IedServer_updateVisibleStringAttributeValue(x.ctx, da.ctx, StringData(value))
}

func (x *IedServer) UpdateUTCTimeAttributeValue(da *DataAttribute, value uint64) {
	C.IedServer_updateUTCTimeAttributeValue(x.ctx, da.ctx, C.uint64_t(value))
}

func (x *IedServer) UpdateTimestampAttributeValue(da *DataAttribute, value *Timestamp) {
	C.IedServer_updateTimestampAttributeValue(x.ctx, da.ctx, value.ctx)
}

func (x *IedServer) UpdateQuality(da *DataAttribute, value Quality) {
	C.IedServer_updateQuality(x.ctx, da.ctx, C.Quality(value))
}

func (x *IedServer) ChangeActiveSettingGroup(sgcb *SettingGroupControlBlock, newActiveSg uint8) {
	C.IedServer_changeActiveSettingGroup(x.ctx, sgcb.ctx, C.uint8_t(newActiveSg))
}

func (x *IedServer) GetActiveSettingGroup(sgcb *SettingGroupControlBlock) uint8 {
	return uint8(C.IedServer_getActiveSettingGroup(x.ctx, sgcb.ctx))
}

type ActiveSettingGroupChangedHandler func(parameter unsafe.Pointer, sgcb *SettingGroupControlBlock, newActSg uint8, connection *ClientConnection) bool

var mapActiveSettingGroupChangedHandlers = sync.Map{}

//export fActiveSettingGroupChangedHandlerGo
func fActiveSettingGroupChangedHandlerGo(parameter unsafe.Pointer, sgcb *C.SettingGroupControlBlock, newActSg C.uint8_t, connection C.ClientConnection) C._Bool {
	ret := false
	mapActiveSettingGroupChangedHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(ActiveSettingGroupChangedHandler); ok {
			ret = fn(parameter, &SettingGroupControlBlock{ctx: sgcb}, uint8(newActSg), &ClientConnection{ctx: connection})
		}
		return true
	})
	return C._Bool(ret)
}

func (x *IedServer) SetActiveSettingGroupChangedHandler(sgcb *SettingGroupControlBlock, handler ActiveSettingGroupChangedHandler, parameter unsafe.Pointer) {
	C.IedServer_setActiveSettingGroupChangedHandler(x.ctx, sgcb.ctx, C.ActiveSettingGroupChangedHandler(C.fActiveSettingGroupChangedHandlerGo), parameter)
}

type EditSettingGroupChangedHandler func(parameter unsafe.Pointer, sgcb *SettingGroupControlBlock, newEditSg uint8, connection *ClientConnection) bool

var mapEditSettingGroupChangedHandlers = sync.Map{}

//export fEditSettingGroupChangedHandlerGo
func fEditSettingGroupChangedHandlerGo(parameter unsafe.Pointer, sgcb *C.SettingGroupControlBlock, newEditSg C.uint8_t, connection C.ClientConnection) C._Bool {
	ret := false
	mapEditSettingGroupChangedHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(EditSettingGroupChangedHandler); ok {
			ret = fn(parameter, &SettingGroupControlBlock{ctx: sgcb}, uint8(newEditSg), &ClientConnection{ctx: connection})
		}
		return true
	})
	return C._Bool(ret)
}

func (x *IedServer) SetEditSettingGroupChangedHandler(sgcb *SettingGroupControlBlock, handler EditSettingGroupChangedHandler, parameter unsafe.Pointer) {
	C.IedServer_setEditSettingGroupChangedHandler(x.ctx, sgcb.ctx, C.EditSettingGroupChangedHandler(C.fEditSettingGroupChangedHandlerGo), parameter)
}

type EditSettingGroupConfirmationHandler func(parameter unsafe.Pointer, sgcb *SettingGroupControlBlock, editSg uint8)

var mapEditSettingGroupConfirmationHandlers = sync.Map{}

//export fEditSettingGroupConfirmationHandlerGo
func fEditSettingGroupConfirmationHandlerGo(parameter unsafe.Pointer, sgcb *C.SettingGroupControlBlock, editSg C.uint8_t) {
	mapEditSettingGroupConfirmationHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(EditSettingGroupConfirmationHandler); ok {
			fn(parameter, &SettingGroupControlBlock{ctx: sgcb}, uint8(editSg))
		}
		return true
	})
}

func (x *IedServer) SetEditSettingGroupConfirmationHandler(sgcb *SettingGroupControlBlock, handler EditSettingGroupConfirmationHandler, parameter unsafe.Pointer) {
	C.IedServer_setEditSettingGroupConfirmationHandler(x.ctx, sgcb.ctx, C.EditSettingGroupConfirmationHandler(C.fEditSettingGroupConfirmationHandlerGo), parameter)
}

type CheckHandlerResult int32

const (
	CONTROL_ACCEPTED                CheckHandlerResult = C.CONTROL_ACCEPTED
	CONTROL_WAITING_FOR_SELECT      CheckHandlerResult = C.CONTROL_WAITING_FOR_SELECT
	CONTROL_HARDWARE_FAULT          CheckHandlerResult = C.CONTROL_HARDWARE_FAULT
	CONTROL_TEMPORARILY_UNAVAILABLE CheckHandlerResult = C.CONTROL_TEMPORARILY_UNAVAILABLE
	CONTROL_OBJECT_ACCESS_DENIED    CheckHandlerResult = C.CONTROL_OBJECT_ACCESS_DENIED
	CONTROL_OBJECT_UNDEFINED        CheckHandlerResult = C.CONTROL_OBJECT_UNDEFINED
	CONTROL_VALUE_INVALID           CheckHandlerResult = C.CONTROL_VALUE_INVALID
)

type ControlHandlerResult int32

const (
	CONTROL_RESULT_FAILED  ControlHandlerResult = C.CONTROL_RESULT_FAILED
	CONTROL_RESULT_OK      ControlHandlerResult = C.CONTROL_RESULT_OK
	CONTROL_RESULT_WAITING ControlHandlerResult = C.CONTROL_RESULT_WAITING
)

type ControlAction struct {
	ctx C.ControlAction
}

func (x *ControlAction) SetError(error ControlLastApplError) {
	C.ControlAction_setError(x.ctx, C.ControlLastApplError(error))
}

func (x *ControlAction) SetAddCause(addCause ControlAddCause) {
	C.ControlAction_setAddCause(x.ctx, C.ControlAddCause(addCause))
}

func (x *ControlAction) GetOrCat() int {
	return int(C.ControlAction_getOrCat(x.ctx))
}

func (x *ControlAction) GetOrIdent() []byte {
	var orIdentSize C.int
	orIdent := C.ControlAction_getOrIdent(x.ctx, &orIdentSize)
	return C.GoBytes(unsafe.Pointer(orIdent), C.int(orIdentSize))
}

func (x *ControlAction) GetCtlNum() int {
	return int(C.ControlAction_getCtlNum(x.ctx))
}

func (x *ControlAction) GetSynchroCheck() bool {
	return bool(C.ControlAction_getSynchroCheck(x.ctx))
}

func (x *ControlAction) GetInterlockCheck() bool {
	return bool(C.ControlAction_getInterlockCheck(x.ctx))
}

func (x *ControlAction) IsSelect() bool {
	return bool(C.ControlAction_isSelect(x.ctx))
}

func (x *ControlAction) GetClientConnection() *ClientConnection {
	return &ClientConnection{ctx: C.ControlAction_getClientConnection(x.ctx)}
}

func (x *ControlAction) GetControlObject() *DataObject {
	return &DataObject{ctx: C.ControlAction_getControlObject(x.ctx)}
}

func (x *ControlAction) GetControlTime() uint64 {
	return uint64(C.ControlAction_getControlTime(x.ctx))
}

func (x *ControlAction) GetT() *Timestamp {
	return &Timestamp{ctx: C.ControlAction_getT(x.ctx)}
}

type ControlPerformCheckHandler func(action *ControlAction, parameter unsafe.Pointer, ctlVal *MmsValue, test bool, interlockCheck bool) CheckHandlerResult

var mapControlPerformCheckHandlers = sync.Map{}

//export fControlPerformCheckHandlerGo
func fControlPerformCheckHandlerGo(action C.ControlAction, parameter unsafe.Pointer, ctlVal *C.MmsValue, test C._Bool, interlockCheck C._Bool) C.CheckHandlerResult {
	ret := CONTROL_ACCEPTED
	mapControlPerformCheckHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(ControlPerformCheckHandler); ok {
			ret = fn(&ControlAction{ctx: action}, parameter, &MmsValue{ctx: ctlVal}, bool(test), bool(interlockCheck))
		}
		return true
	})
	return C.CheckHandlerResult(ret)
}

type ControlWaitForExecutionHandler func(action *ControlAction, parameter unsafe.Pointer, ctlVal *MmsValue, test bool, synchroCheck bool) ControlHandlerResult

var mapControlWaitForExecutionHandlers = sync.Map{}

//export fControlWaitForExecutionHandlerGo
func fControlWaitForExecutionHandlerGo(action C.ControlAction, parameter unsafe.Pointer, ctlVal *C.MmsValue, test C._Bool, synchroCheck C._Bool) C.ControlHandlerResult {
	ret := CONTROL_RESULT_OK
	mapControlWaitForExecutionHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(ControlWaitForExecutionHandler); ok {
			ret = fn(&ControlAction{ctx: action}, parameter, &MmsValue{ctx: ctlVal}, bool(test), bool(synchroCheck))
		}
		return true
	})
	return C.ControlHandlerResult(ret)
}

type ControlHandler func(action *ControlAction, parameter unsafe.Pointer, ctlVal *MmsValue, test bool) ControlHandlerResult

var mapControlHandlers = sync.Map{}

//export fControlHandlerGo
func fControlHandlerGo(action C.ControlAction, parameter unsafe.Pointer, ctlVal *C.MmsValue, test C._Bool) C.ControlHandlerResult {
	ret := CONTROL_RESULT_OK
	mapControlHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(ControlHandler); ok {
			ret = fn(&ControlAction{ctx: action}, parameter, &MmsValue{ctx: ctlVal}, bool(test))
		}
		return true
	})
	return C.ControlHandlerResult(ret)
}

type SelectStateChangedReason int32

const (
	SELECT_STATE_REASON_SELECTED       SelectStateChangedReason = C.SELECT_STATE_REASON_SELECTED
	SELECT_STATE_REASON_CANCELED       SelectStateChangedReason = C.SELECT_STATE_REASON_CANCELED
	SELECT_STATE_REASON_TIMEOUT        SelectStateChangedReason = C.SELECT_STATE_REASON_TIMEOUT
	SELECT_STATE_REASON_OPERATED       SelectStateChangedReason = C.SELECT_STATE_REASON_OPERATED
	SELECT_STATE_REASON_OPERATE_FAILED SelectStateChangedReason = C.SELECT_STATE_REASON_OPERATE_FAILED
	SELECT_STATE_REASON_DISCONNECTED   SelectStateChangedReason = C.SELECT_STATE_REASON_DISCONNECTED
)

type ControlSelectStateChangedHandler func(action *ControlAction, parameter unsafe.Pointer, isSelected bool, reason SelectStateChangedReason)

var mapControlSelectStateChangedHandlers = sync.Map{}

//export fControlSelectStateChangedHandlerGo
func fControlSelectStateChangedHandlerGo(action C.ControlAction, parameter unsafe.Pointer, isSelected C._Bool, reason C.SelectStateChangedReason) {
	mapControlSelectStateChangedHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(ControlSelectStateChangedHandler); ok {
			fn(&ControlAction{ctx: action}, parameter, bool(isSelected), SelectStateChangedReason(reason))
		}
		return true
	})
}

func (x *IedServer) SetControlHandler(node *DataObject, handler ControlHandler, parameter unsafe.Pointer) {
	C.IedServer_setControlHandler(x.ctx, node.ctx, C.ControlHandler(C.fControlHandlerGo), parameter)
}

func (x *IedServer) SetPerformCheckHandler(node *DataObject, handler ControlPerformCheckHandler, parameter unsafe.Pointer) {
	C.IedServer_setPerformCheckHandler(x.ctx, node.ctx, C.ControlPerformCheckHandler(C.fControlPerformCheckHandlerGo), parameter)
}

func (x *IedServer) SetWaitForExecutionHandler(node *DataObject, handler ControlWaitForExecutionHandler, parameter unsafe.Pointer) {
	C.IedServer_setWaitForExecutionHandler(x.ctx, node.ctx, C.ControlWaitForExecutionHandler(C.fControlWaitForExecutionHandlerGo), parameter)
}

func (x *IedServer) SetSelectStateChangedHandler(node *DataObject, handler ControlSelectStateChangedHandler, parameter unsafe.Pointer) {
	C.IedServer_setSelectStateChangedHandler(x.ctx, node.ctx, C.ControlSelectStateChangedHandler(C.fControlSelectStateChangedHandlerGo), parameter)
}

func (x *IedServer) UpdateCtlModel(node *DataObject, value ControlModel) {
	C.IedServer_updateCtlModel(x.ctx, node.ctx, C.ControlModel(value))
}

type RCBEventType int32

const (
	RCB_EVENT_GET_PARAMETER  RCBEventType = C.RCB_EVENT_GET_PARAMETER
	RCB_EVENT_SET_PARAMETER  RCBEventType = C.RCB_EVENT_SET_PARAMETER
	RCB_EVENT_UNRESERVED     RCBEventType = C.RCB_EVENT_UNRESERVED
	RCB_EVENT_RESERVED       RCBEventType = C.RCB_EVENT_RESERVED
	RCB_EVENT_ENABLE         RCBEventType = C.RCB_EVENT_ENABLE
	RCB_EVENT_DISABLE        RCBEventType = C.RCB_EVENT_DISABLE
	RCB_EVENT_GI             RCBEventType = C.RCB_EVENT_GI
	RCB_EVENT_PURGEBUF       RCBEventType = C.RCB_EVENT_PURGEBUF
	RCB_EVENT_OVERFLOW       RCBEventType = C.RCB_EVENT_OVERFLOW
	RCB_EVENT_REPORT_CREATED RCBEventType = C.RCB_EVENT_REPORT_CREATED
)

type RCBEventHandler func(parameter unsafe.Pointer, rcb *ReportControlBlock, connection *ClientConnection, event RCBEventType, parameterName string, serviceError MmsDataAccessError)

var mapRCBEventHandlers = sync.Map{}

//export fRCBEventHandlerGo
func fRCBEventHandlerGo(parameter unsafe.Pointer, rcb *C.ReportControlBlock, connection C.ClientConnection, event C.IedServer_RCBEventType, parameterName *C.char, serviceError C.MmsDataAccessError) {
	mapRCBEventHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(RCBEventHandler); ok {
			fn(parameter, &ReportControlBlock{ctx: rcb}, &ClientConnection{ctx: connection}, RCBEventType(event), C.GoString(parameterName), MmsDataAccessError(serviceError))
		}
		return true
	})
}

func (x *IedServer) SetRCBEventHandler(handler RCBEventHandler, parameter unsafe.Pointer) {
	C.IedServer_setRCBEventHandler(x.ctx, C.IedServer_RCBEventHandler(C.fRCBEventHandlerGo), parameter)
}

const (
	IEC61850_SVCB_EVENT_ENABLE  = C.IEC61850_SVCB_EVENT_ENABLE
	IEC61850_SVCB_EVENT_DISABLE = C.IEC61850_SVCB_EVENT_DISABLE
)

type SVCBEventHandler func(svcb *SVControlBlock, event int32, parameter unsafe.Pointer)

var mapSVCBEventHandlers = sync.Map{}

//export fSVCBEventHandlerGo
func fSVCBEventHandlerGo(svcb *C.SVControlBlock, event C.int32_t, parameter unsafe.Pointer) {
	mapSVCBEventHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(SVCBEventHandler); ok {
			fn(&SVControlBlock{ctx: svcb}, int32(event), parameter)
		}
		return true
	})
}

func (x *IedServer) SetSVCBEventHandler(svcb *SVControlBlock, handler SVCBEventHandler, parameter unsafe.Pointer) {
	C.IedServer_setSVCBHandler(x.ctx, svcb.ctx, C.SVCBEventHandler(C.fSVCBEventHandlerGo), parameter)
}

type MmsGooseControlBlock struct {
	ctx C.MmsGooseControlBlock
}

const (
	IEC61850_GOCB_EVENT_ENABLE  = C.IEC61850_GOCB_EVENT_ENABLE
	IEC61850_GOCB_EVENT_DISABLE = C.IEC61850_GOCB_EVENT_DISABLE
)

type GoCBEventHandler func(goCb *MmsGooseControlBlock, event int32, parameter unsafe.Pointer)

var mapGoCBEventHandlers = sync.Map{}

//export fGoCBEventHandlerGo
func fGoCBEventHandlerGo(goCb C.MmsGooseControlBlock, event C.int32_t, parameter unsafe.Pointer) {
	mapGoCBEventHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(GoCBEventHandler); ok {
			fn(&MmsGooseControlBlock{ctx: goCb}, int32(event), parameter)
		}
		return true
	})
}

func (x *IedServer) SetGoCBEventHandler(handler GoCBEventHandler, parameter unsafe.Pointer) {
	C.IedServer_setGoCBHandler(x.ctx, C.GoCBEventHandler(C.fGoCBEventHandlerGo), parameter)
}

func (x *MmsGooseControlBlock) GetName() string {
	return C.GoString(C.MmsGooseControlBlock_getName(x.ctx))
}

func (x *MmsGooseControlBlock) GetLogicalNode() *LogicalNode {
	return &LogicalNode{ctx: C.MmsGooseControlBlock_getLogicalNode(x.ctx)}
}

func (x *MmsGooseControlBlock) GetDataSet() *DataSet {
	return &DataSet{ctx: C.MmsGooseControlBlock_getDataSet(x.ctx)}
}

func (x *MmsGooseControlBlock) GetGoEna() bool {
	return bool(C.MmsGooseControlBlock_getGoEna(x.ctx))
}
func (x *MmsGooseControlBlock) GetMinTime() int {
	return int(C.MmsGooseControlBlock_getMinTime(x.ctx))
}
func (x *MmsGooseControlBlock) GetMaxTime() int {
	return int(C.MmsGooseControlBlock_getMaxTime(x.ctx))
}

func (x *MmsGooseControlBlock) GetFixedOffs() bool {
	return bool(C.MmsGooseControlBlock_getFixedOffs(x.ctx))
}
func (x *MmsGooseControlBlock) GetNdsCom() bool {
	return bool(C.MmsGooseControlBlock_getNdsCom(x.ctx))
}

type WriteAccessHandler func(dataAttribute *DataAttribute, value *MmsValue, connection *ClientConnection, parameter unsafe.Pointer) MmsDataAccessError

var mapWriteAccessHandlers = sync.Map{}

//export fWriteAccessHandlerGo
func fWriteAccessHandlerGo(dataAttribute *C.DataAttribute, value *C.MmsValue, connection C.ClientConnection, parameter unsafe.Pointer) C.MmsDataAccessError {
	ret := DATA_ACCESS_ERROR_UNKNOWN
	mapWriteAccessHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(WriteAccessHandler); ok {
			ret = fn(&DataAttribute{ctx: dataAttribute}, &MmsValue{ctx: value}, &ClientConnection{ctx: connection}, parameter)
		}
		return true
	})
	return C.MmsDataAccessError(ret)
}

func (x *IedServer) HandleWriteAccess(dataAttribute *DataAttribute, handler WriteAccessHandler, parameter unsafe.Pointer) {
	C.IedServer_handleWriteAccess(x.ctx, dataAttribute.ctx, C.WriteAccessHandler(C.fWriteAccessHandlerGo), parameter)
}

func (x *IedServer) HandleWriteAccessForComplexAttribute(dataAttribute *DataAttribute, handler WriteAccessHandler, parameter unsafe.Pointer) {
	C.IedServer_handleWriteAccessForComplexAttribute(x.ctx, dataAttribute.ctx, C.WriteAccessHandler(C.fWriteAccessHandlerGo), parameter)
}

func (x *IedServer) HandleWriteAccessForDataObject(dataObject *DataObject, fc FunctionalConstraint, handler WriteAccessHandler, parameter unsafe.Pointer) {
	C.IedServer_handleWriteAccessForDataObject(x.ctx, dataObject.ctx, C.FunctionalConstraint(fc), C.WriteAccessHandler(C.fWriteAccessHandlerGo), parameter)
}

type AccessPolicy int32

const (
	ACCESS_POLICY_ALLOW AccessPolicy = C.ACCESS_POLICY_ALLOW
	ACCESS_POLICY_DENY  AccessPolicy = C.ACCESS_POLICY_DENY
)

func (x *IedServer) SetWriteAccessPolicy(fc FunctionalConstraint, policy AccessPolicy) {
	C.IedServer_setWriteAccessPolicy(x.ctx, C.FunctionalConstraint(fc), C.AccessPolicy(policy))
}

type ReadAccessHandler func(ld *LogicalDevice, ln *LogicalNode, dataObject *DataObject, fc FunctionalConstraint, connection *ClientConnection, parameter unsafe.Pointer) MmsDataAccessError

var mapReadAccessHandlers = sync.Map{}

//export fReadAccessHandlerGo
func fReadAccessHandlerGo(ld *C.LogicalDevice, ln *C.LogicalNode, dataObject *C.DataObject, fc C.FunctionalConstraint, connection C.ClientConnection, parameter unsafe.Pointer) C.MmsDataAccessError {
	ret := DATA_ACCESS_ERROR_UNKNOWN
	mapReadAccessHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(ReadAccessHandler); ok {
			ret = fn(&LogicalDevice{ctx: ld}, &LogicalNode{ctx: ln}, &DataObject{ctx: dataObject}, FunctionalConstraint(fc), &ClientConnection{ctx: connection}, parameter)
		}
		return true
	})
	return C.MmsDataAccessError(ret)
}

func (x *IedServer) SetReadAccessHandler(handler ReadAccessHandler, parameter unsafe.Pointer) {
	C.IedServer_setReadAccessHandler(x.ctx, C.ReadAccessHandler(C.fReadAccessHandlerGo), parameter)
}

type DataSetOperation int32

const (
	DATASET_CREATE        DataSetOperation = C.DATASET_CREATE
	DATASET_DELETE        DataSetOperation = C.DATASET_DELETE
	DATASET_READ          DataSetOperation = C.DATASET_READ
	DATASET_WRITE         DataSetOperation = C.DATASET_WRITE
	DATASET_GET_DIRECTORY DataSetOperation = C.DATASET_GET_DIRECTORY
)

type DataSetAccessHandler func(parameter unsafe.Pointer, connection *ClientConnection, operation DataSetOperation, datasetRef string) bool

var mapDataSetAccessHandlers = sync.Map{}

//export fDataSetAccessHandlerGo
func fDataSetAccessHandlerGo(parameter unsafe.Pointer, connection C.ClientConnection, operation C.IedServer_DataSetOperation, datasetRef *C.char) C._Bool {
	ret := false
	mapDataSetAccessHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(DataSetAccessHandler); ok {
			ret = fn(parameter, &ClientConnection{ctx: connection}, DataSetOperation(operation), C.GoString(datasetRef))
		}
		return true
	})
	return C.bool(ret)
}

func (x *IedServer) SetDataSetAccessHandler(handler DataSetAccessHandler, parameter unsafe.Pointer) {
	C.IedServer_setDataSetAccessHandler(x.ctx, C.IedServer_DataSetAccessHandler(C.fDataSetAccessHandlerGo), parameter)
}

type DirectoryCategory int32

const (
	DIRECTORY_CAT_LD_LIST      DirectoryCategory = C.DIRECTORY_CAT_LD_LIST
	DIRECTORY_CAT_DATA_LIST    DirectoryCategory = C.DIRECTORY_CAT_DATA_LIST
	DIRECTORY_CAT_DATASET_LIST DirectoryCategory = C.DIRECTORY_CAT_DATASET_LIST
	DIRECTORY_CAT_LOG_LIST     DirectoryCategory = C.DIRECTORY_CAT_LOG_LIST
)

type DirectoryAccessHandler func(parameter unsafe.Pointer, connection *ClientConnection, category DirectoryCategory, logicalDevice *LogicalDevice) bool

var mapDirectoryAccessHandlers = sync.Map{}

//export fDirectoryAccessHandlerGo
func fDirectoryAccessHandlerGo(parameter unsafe.Pointer, connection C.ClientConnection, category C.IedServer_DirectoryCategory, logicalDevice *C.LogicalDevice) C._Bool {
	ret := false
	mapDirectoryAccessHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(DirectoryAccessHandler); ok {
			ret = fn(parameter, &ClientConnection{ctx: connection}, DirectoryCategory(category), &LogicalDevice{ctx: logicalDevice})
		}
		return true
	})
	return C._Bool(ret)
}

func (x *IedServer) SetDirectoryAccessHandler(handler DirectoryAccessHandler, parameter unsafe.Pointer) {
	C.IedServer_setDirectoryAccessHandler(x.ctx, C.IedServer_DirectoryAccessHandler(C.fDirectoryAccessHandlerGo), parameter)
}

type ListObjectsAccessHandler func(parameter unsafe.Pointer, connection *ClientConnection, acsiClass ACSIClass, ld *LogicalDevice, ln *LogicalNode, objectName string, subObjectName string, fc FunctionalConstraint) bool

var mapListObjectsAccessHandlers = sync.Map{}

//export fListObjectsAccessHandlerGo
func fListObjectsAccessHandlerGo(parameter unsafe.Pointer, connection C.ClientConnection, acsiClass C.ACSIClass, ld *C.LogicalDevice, ln *C.LogicalNode, objectName *C.char, subObjectName *C.char, fc C.FunctionalConstraint) C._Bool {
	ret := false
	mapListObjectsAccessHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(ListObjectsAccessHandler); ok {
			ret = fn(parameter, &ClientConnection{ctx: connection}, ACSIClass(acsiClass), &LogicalDevice{ctx: ld}, &LogicalNode{ctx: ln}, C.GoString(objectName), C.GoString(subObjectName), FunctionalConstraint(fc))
		}
		return true
	})
	return C._Bool(ret)
}

func (x *IedServer) SetListObjectsAccessHandler(handler ListObjectsAccessHandler, parameter unsafe.Pointer) {
	C.IedServer_setListObjectsAccessHandler(x.ctx, C.IedServer_ListObjectsAccessHandler(C.fListObjectsAccessHandlerGo), parameter)
}

type ControlBlockAccessType int32

const (
	IEC61850_CB_ACCESS_TYPE_READ  ControlBlockAccessType = C.IEC61850_CB_ACCESS_TYPE_READ
	IEC61850_CB_ACCESS_TYPE_WRITE ControlBlockAccessType = C.IEC61850_CB_ACCESS_TYPE_WRITE
)

type ControlBlockAccessHandler func(parameter unsafe.Pointer, connection *ClientConnection, acsiClass ACSIClass, ld *LogicalDevice, ln *LogicalNode, objectName string, subObjectName string, accessType ControlBlockAccessType) bool

var mapControlBlockAccessHandlers = sync.Map{}

//export fControlBlockAccessHandlerGo
func fControlBlockAccessHandlerGo(parameter unsafe.Pointer, connection C.ClientConnection, acsiClass C.ACSIClass, ld *C.LogicalDevice, ln *C.LogicalNode, objectName *C.char, subObjectName *C.char, accessType C.IedServer_ControlBlockAccessType) C._Bool {
	ret := false
	mapControlBlockAccessHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(ControlBlockAccessHandler); ok {
			ret = fn(parameter, &ClientConnection{ctx: connection}, ACSIClass(acsiClass), &LogicalDevice{ctx: ld}, &LogicalNode{ctx: ln}, C.GoString(objectName), C.GoString(subObjectName), ControlBlockAccessType(accessType))
		}
		return true
	})
	return C._Bool(ret)
}

func (x *IedServer) SetControlBlockAccessHandler(handler ControlBlockAccessHandler, parameter unsafe.Pointer) {
	C.IedServer_setControlBlockAccessHandler(x.ctx, C.IedServer_ControlBlockAccessHandler(C.fControlBlockAccessHandlerGo), parameter)
}

func (x *IedServer) IgnoreReadAccess(ignore bool) {
	C.IedServer_ignoreReadAccess(x.ctx, C._Bool(ignore))
}
