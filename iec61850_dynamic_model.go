package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_dynamic_model.h"

*/
import "C"
import "unsafe"

func IedModelCreate(name string) *IedModel {
	return &IedModel{ctx: C.IedModel_create((*C.char)(unsafe.Pointer(unsafe.StringData(name + "\x00"))))}
}

func (x *IedModel) SetIedNameForDynamicModel(name string) {
	C.IedModel_setIedNameForDynamicModel(x.ctx, StringData(name))
}

func (x *IedModel) Destroy() {
	C.IedModel_destroy(x.ctx)
}

func (x *IedModel) LogicalDeviceCreate(name string) *LogicalDevice {
	return &LogicalDevice{ctx: C.LogicalDevice_create(StringData(name), x.ctx)}
}

func LogicalDeviceCreateEx(inst string, parent *IedModel, ldName string) *LogicalDevice {
	return &LogicalDevice{ctx: C.LogicalDevice_createEx(StringData(inst), parent.ctx, StringData(ldName))}
}

func LogicalNodeCreate(name string, parent *LogicalDevice) *LogicalNode {
	return &LogicalNode{ctx: C.LogicalNode_create(StringData(name), parent.ctx)}
}

func DataObjectCreate(name string, parent *ModelNode, arrayElements int) *DataObject {
	return &DataObject{ctx: C.DataObject_create(StringData(name), parent.ctx, C.int(arrayElements))}
}

func DataAttributeCreate(name string, parent *ModelNode, type_ DataAttributeType, fc FunctionalConstraint, triggerOptions uint8, arrayElements int, sAddr uint32) *DataAttribute {
	return &DataAttribute{ctx: C.DataAttribute_create(StringData(name), parent.ctx, C.DataAttributeType(type_),
		C.FunctionalConstraint(fc), C.uint8_t(triggerOptions), C.int(arrayElements), C.uint32_t(sAddr))}
}

func (x *DataAttribute) GetType() DataAttributeType {
	return DataAttributeType(C.DataAttribute_getType(x.ctx))
}

func (x *DataAttribute) GetFC() FunctionalConstraint {
	return FunctionalConstraint(C.DataAttribute_getFC(x.ctx))
}

func (x *DataAttribute) GetTrgOps() uint8 {
	return uint8(C.DataAttribute_getTrgOps(x.ctx))
}

func (x *DataAttribute) SetValue(value *MmsValue) {
	C.DataAttribute_setValue(x.ctx, value.ctx)
}

func ReportControlBlockCreate(name string, parent *LogicalNode, rptId string, isBuffered bool, dataSetName string,
	confRef uint32, trgOps uint8, options uint8, bufTm uint32, intgPd uint32) *ReportControlBlock {
	return &ReportControlBlock{ctx: C.ReportControlBlock_create(StringData(name), parent.ctx, StringData(rptId), C.bool(isBuffered),
		StringData(dataSetName), C.uint32_t(confRef), C.uint8_t(trgOps), C.uint8_t(options), C.uint32_t(bufTm), C.uint32_t(intgPd))}
}

func (x *ReportControlBlock) SetPreconfiguredClient(clientType uint8, clientAddress []uint8) {
	C.ReportControlBlock_setPreconfiguredClient(x.ctx, C.uint8_t(clientType), SliceData(clientAddress))
}

func (x *ReportControlBlock) GetName() string {
	return C.GoString(C.ReportControlBlock_getName(x.ctx))
}

func (x *ReportControlBlock) IsBuffered() bool {
	return bool(C.ReportControlBlock_isBuffered(x.ctx))
}

func (x *ReportControlBlock) GetParent() *LogicalNode {
	return &LogicalNode{ctx: C.ReportControlBlock_getParent(x.ctx)}
}

func (x *ReportControlBlock) GetRptId() string {
	return C.GoString(C.ReportControlBlock_getRptID(x.ctx))
}

func (x *ReportControlBlock) GetRptEna() bool {
	return bool(C.ReportControlBlock_getRptEna(x.ctx))
}

func (x *ReportControlBlock) GetDataSetName() string {
	return C.GoString(C.ReportControlBlock_getDataSet(x.ctx))
}

func (x *ReportControlBlock) GetConfRef() uint32 {
	return uint32(C.ReportControlBlock_getConfRev(x.ctx))
}

func (x *ReportControlBlock) GetOptFlds() uint8 {
	return uint8(C.ReportControlBlock_getOptFlds(x.ctx))
}

func (x *ReportControlBlock) GetBufTm() uint32 {
	return uint32(C.ReportControlBlock_getBufTm(x.ctx))
}

func (x *ReportControlBlock) GetSqNum() uint16 {
	return uint16(C.ReportControlBlock_getSqNum(x.ctx))
}

func (x *ReportControlBlock) GetTrgOps() uint8 {
	return uint8(C.ReportControlBlock_getTrgOps(x.ctx))
}

func (x *ReportControlBlock) GetIntgPd() uint32 {
	return uint32(C.ReportControlBlock_getIntgPd(x.ctx))
}

func (x *ReportControlBlock) GetGI() bool {
	return bool(C.ReportControlBlock_getGI(x.ctx))
}

func (x *ReportControlBlock) GetPurgeBuf() bool {
	return bool(C.ReportControlBlock_getPurgeBuf(x.ctx))
}

func (x *ReportControlBlock) GetEntryId() *MmsValue {
	return &MmsValue{ctx: C.ReportControlBlock_getEntryId(x.ctx)}
}

func (x *ReportControlBlock) GetTimeofEntry() uint64 {
	return uint64(C.ReportControlBlock_getTimeofEntry(x.ctx))
}

func (x *ReportControlBlock) GetResvTms() int16 {
	return int16(C.ReportControlBlock_getResvTms(x.ctx))
}

func (x *ReportControlBlock) GetResv() bool {
	return bool(C.ReportControlBlock_getResv(x.ctx))
}

func (x *ReportControlBlock) GetOwner() *MmsValue {
	return &MmsValue{ctx: C.ReportControlBlock_getOwner(x.ctx)}
}

func LogControlBlockCreate(name string, parent *LogicalNode, dataSetName string, logRef string, trgOps uint8,
	intgPd uint32, logEna bool, reasonCode bool) *LogControlBlock {
	return &LogControlBlock{ctx: C.LogControlBlock_create(StringData(name), parent.ctx, StringData(dataSetName), StringData(logRef), C.uint8_t(trgOps),
		C.uint32_t(intgPd), C.bool(logEna), C.bool(reasonCode))}
}

func (x *LogControlBlock) GetName() string {
	return C.GoString(C.LogControlBlock_getName(x.ctx))
}

func (x *LogControlBlock) GetParent() *LogicalNode {
	return &LogicalNode{ctx: C.LogControlBlock_getParent(x.ctx)}
}

func LogCreate(name string, parent *LogicalNode) *Log {
	return &Log{ctx: C.Log_create(StringData(name), parent.ctx)}
}

func SettingGroupControlBlockCreate(parent *LogicalNode, actSG uint8, numOfSGs uint8) *SettingGroupControlBlock {
	return &SettingGroupControlBlock{ctx: C.SettingGroupControlBlock_create(parent.ctx, C.uint8_t(actSG), C.uint8_t(numOfSGs))}
}

func GSEControlBlockCreate(name string, parent *LogicalNode, appId string, dataSet string, confRev uint32,
	fixedOffs bool, minTime int, maxTime int) *GSEControlBlock {
	return &GSEControlBlock{ctx: C.GSEControlBlock_create(StringData(name), parent.ctx, StringData(appId), StringData(dataSet), C.uint32_t(confRev),
		C.bool(fixedOffs), C.int(minTime), C.int(maxTime))}
}

func SVControlBlockCreate(name string, parent *LogicalNode, svID string, dataSet string, confRev uint32, smpMod uint8,
	smpRate uint16, optFlds uint8, isUnicast bool) *SVControlBlock {
	return &SVControlBlock{ctx: C.SVControlBlock_create(StringData(name), parent.ctx, StringData(svID), StringData(dataSet), C.uint32_t(confRev), C.uint8_t(smpMod),
		C.uint16_t(smpRate), C.uint8_t(optFlds), C.bool(isUnicast))}
}

func (x *SVControlBlock) GetName() string {
	return C.GoString(C.SVControlBlock_getName(x.ctx))
}

func (x *SVControlBlock) AddPhyComAddress(phyComAddress *PhyComAddress) {
	C.SVControlBlock_addPhyComAddress(x.ctx, phyComAddress.ctx)
}

func (x *GSEControlBlock) AddPhyComAddress(phyComAddress *PhyComAddress) {
	C.GSEControlBlock_addPhyComAddress(x.ctx, phyComAddress.ctx)
}

func PhyComAddressCreate(vlanPriority uint8, vlanId uint16, appId uint16, dstAddress []uint8) *PhyComAddress {
	return &PhyComAddress{ctx: C.PhyComAddress_create(C.uint8_t(vlanPriority), C.uint16_t(vlanId), C.uint16_t(appId), SliceData(dstAddress))}
}

func DataSetCreate(name string, parent *LogicalNode) *DataSet {
	return &DataSet{ctx: C.DataSet_create(StringData(name), parent.ctx)}
}

func (x *DataSet) GetName() string {
	return C.GoString(C.DataSet_getName(x.ctx))
}

func (x *DataSet) GetSize() int {
	return int(C.DataSet_getSize(x.ctx))
}

func (x *DataSet) GetFirstEntry() *DataSetEntry {
	return &DataSetEntry{ctx: C.DataSet_getFirstEntry(x.ctx)}
}

func (x *DataSetEntry) GetNext() *DataSetEntry {
	return &DataSetEntry{ctx: C.DataSetEntry_getNext(x.ctx)}
}

func DataSetEntryCreate(dataSet *DataSet, variable string, index int, component string) *DataSetEntry {
	return &DataSetEntry{ctx: C.DataSetEntry_create(dataSet.ctx, StringData(variable), C.int(index), StringData(component))}
}
