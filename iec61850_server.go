package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_server.h"

extern bool fAcseAuthenticatorGo(void* parameter, AcseAuthenticationParameter authParameter, void** securityToken, IsoApplicationReference* appReference);
extern void fIedConnectionIndicationHandlerGo(IedServer self, ClientConnection connection, bool connected, void* parameter);
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
	cs := C.CString(basepath)
	defer C.free(unsafe.Pointer(cs))
	C.IedServerConfig_setFileServiceBasePath(x.ctx, cs)
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
	cs := C.CString(ipAddr)
	defer C.free(unsafe.Pointer(cs))
	return bool(C.IedServer_addAccessPoint(x.ctx, cs, C.int(tcpPort), tlsConfiguration.ctx))
}

func (x *IedServer) SetLocalIpAddress(localIpAddress string) {
	cs := C.CString(localIpAddress)
	defer C.free(unsafe.Pointer(cs))
	C.IedServer_setLocalIpAddress(x.ctx, cs)
}

func (x *IedServer) SetServerIdentity(vendor string, model string, revision string) {
	cs_vendor := C.CString(vendor)
	defer C.free(unsafe.Pointer(cs_vendor))
	cs_model := C.CString(model)
	defer C.free(unsafe.Pointer(cs_model))
	cs_revision := C.CString(revision)
	defer C.free(unsafe.Pointer(cs_revision))
	C.IedServer_setServerIdentity(x.ctx, cs_vendor, cs_model, cs_revision)
}

func (x *IedServer) SetFilestoreBasepath(basepath string) {
	cs := C.CString(basepath)
	defer C.free(unsafe.Pointer(cs))
	C.IedServer_setFilestoreBasepath(x.ctx, cs)
}

func (x *IedServer) SetLogStorage(logRef string, logStorage *LogStorage) {
	cs := C.CString(logRef)
	defer C.free(unsafe.Pointer(cs))
	C.IedServer_setLogStorage(x.ctx, cs, logStorage.ctx)
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
	cs := C.CString(interfaceId)
	defer C.free(unsafe.Pointer(cs))
	C.IedServer_setGooseInterfaceId(x.ctx, cs)
}

func (x *IedServer) SetGooseInterfaceIdEx(ln *LogicalNode, gcbName string, interfaceId string) {
	cs_gcbName := C.CString(gcbName)
	defer C.free(unsafe.Pointer(cs_gcbName))
	cs_interfaceId := C.CString(interfaceId)
	defer C.free(unsafe.Pointer(cs_interfaceId))
	C.IedServer_setGooseInterfaceIdEx(x.ctx, ln.ctx, cs_gcbName, cs_interfaceId)
}

func (x *IedServer) UseGooseVlanTag(ln *LogicalNode, gcbName string, useVlanTag bool) {
	cs_gcbName := C.CString(gcbName)
	defer C.free(unsafe.Pointer(cs_gcbName))
	C.IedServer_useGooseVlanTag(x.ctx, ln.ctx, cs_gcbName, C.bool(useVlanTag))
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
		if fn, ok := v.(IedConnectionIndicationHandler); ok {
			fn(&IedServer{ctx: self}, &ClientConnection{ctx: connection}, bool(connected), parameter)
		}
		return true
	})
}

func (x *IedServer) SetConnectionIndicationHandler(handler IedConnectionIndicationHandler, parameter unsafe.Pointer) {
	mapIedConnectionIndicationHandlers.Store(handler, handler)
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
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.IedServer_updateVisibleStringAttributeValue(x.ctx, da.ctx, cstr)
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
