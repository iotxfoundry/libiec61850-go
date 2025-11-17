package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_model.h"

extern void fIedModelInitializerGo();
*/
import "C"
import "sync"

type ModelNode struct {
	ctx *C.ModelNode
}

func (x *ModelNode) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *ModelNode) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *ModelNode) Parent() *ModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *ModelNode) Sibling() *ModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *ModelNode) FirstChild() *ModelNode {
	return &ModelNode{ctx: x.ctx.firstChild}
}

type DataAttribute struct {
	ctx *C.DataAttribute
}

func (x *DataAttribute) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *DataAttribute) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *DataAttribute) Parent() *ModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *DataAttribute) Sibling() *ModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *DataAttribute) FirstChild() *ModelNode {
	return &ModelNode{ctx: x.ctx.firstChild}
}

func (x *DataAttribute) ElementCount() int {
	return int(x.ctx.elementCount)
}

func (x *DataAttribute) ArrayIndex() int {
	return int(x.ctx.arrayIndex)
}

func (x *DataAttribute) Fc() FunctionalConstraint {
	return FunctionalConstraint(x.ctx.fc)
}

func (x *DataAttribute) Type() DataAttributeType {
	return DataAttributeType(x.ctx._type)
}

func (x *DataAttribute) TriggerOptions() uint8 {
	return uint8(x.ctx.triggerOptions)
}

func (x *DataAttribute) MmsValue() *MmsValue {
	return &MmsValue{ctx: x.ctx.mmsValue}
}

func (x *DataAttribute) Addr() uint32 {
	return uint32(x.ctx.sAddr)
}

type DataObject struct {
	ctx *C.DataObject
}

func (x *DataObject) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *DataObject) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *DataObject) Parent() *ModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *DataObject) Sibling() *ModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *DataObject) FirstChild() *ModelNode {
	return &ModelNode{ctx: x.ctx.firstChild}
}

func (x *DataObject) ElementCount() int {
	return int(x.ctx.elementCount)
}

func (x *DataObject) ArrayIndex() int {
	return int(x.ctx.arrayIndex)
}

type LogicalNode struct {
	ctx *C.LogicalNode
}

func (x *LogicalNode) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *LogicalNode) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *LogicalNode) Parent() *ModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *LogicalNode) Sibling() *ModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *LogicalNode) FirstChild() *ModelNode {
	return &ModelNode{ctx: x.ctx.firstChild}
}

type LogicalDevice struct {
	ctx *C.LogicalDevice
}

func (x *LogicalDevice) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *LogicalDevice) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *LogicalDevice) Parent() *ModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *LogicalDevice) Sibling() *ModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *LogicalDevice) FirstChild() *ModelNode {
	return &ModelNode{ctx: x.ctx.firstChild}
}

func (x *LogicalDevice) LdName() string {
	return C.GoString(x.ctx.ldName)
}

type IedModel struct {
	ctx *C.IedModel
}

func (x *IedModel) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *IedModel) FirstChild() *LogicalDevice {
	return &LogicalDevice{ctx: x.ctx.firstChild}
}

func (x *IedModel) DataSets() *DataSet {
	return &DataSet{ctx: x.ctx.dataSets}
}

func (x *IedModel) Rcbs() *ReportControlBlock {
	return &ReportControlBlock{ctx: x.ctx.rcbs}
}

func (x *IedModel) GseCBs() *GSEControlBlock {
	return &GSEControlBlock{ctx: x.ctx.gseCBs}
}

func (x *IedModel) SvCBs() *SVControlBlock {
	return &SVControlBlock{ctx: x.ctx.svCBs}
}

func (x *IedModel) Sgcbs() *SettingGroupControlBlock {
	return &SettingGroupControlBlock{ctx: x.ctx.sgcbs}
}

func (x *IedModel) Lcbs() *LogControlBlock {
	return &LogControlBlock{ctx: x.ctx.lcbs}
}

func (x *IedModel) Logs() *Log {
	return &Log{ctx: x.ctx.logs}
}

var mapIedModelInitializerCallbacks = sync.Map{}

//export fIedModelInitializerGo
func fIedModelInitializerGo() {
	mapIedModelInitializerCallbacks.Range(func(k, v any) bool {
		cb, ok := v.(func())
		if ok {
			cb()
		}
		return true
	})
}

func (x *IedModel) SetInitializer(initializer func()) {
	// Register the initializer callback
	mapIedModelInitializerCallbacks.Store(x, initializer)
	x.ctx.initializer = C.IedModelInitializer(C.fIedModelInitializerGo)
}

type DataSet struct {
	ctx *C.DataSet
}

func (x *DataSet) LogicalDeviceName() string {
	return C.GoString(x.ctx.logicalDeviceName)
}

func (x *DataSet) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *DataSet) Fcdas() *DataSetEntry {
	return &DataSetEntry{ctx: x.ctx.fcdas}
}

func (x *DataSet) Sibling() *DataSet {
	return &DataSet{ctx: x.ctx.sibling}
}

type DataSetEntry struct {
	ctx *C.DataSetEntry
}

func (x *DataSetEntry) LogicalDeviceName() string {
	return C.GoString(x.ctx.logicalDeviceName)
}

func (x *DataSetEntry) IsLDNameDynamicallyAllocated() bool {
	return bool(x.ctx.isLDNameDynamicallyAllocated)
}

func (x *DataSetEntry) Value() *MmsValue {
	return &MmsValue{ctx: x.ctx.value}
}

func (x *DataSetEntry) Sibling() *DataSetEntry {
	return &DataSetEntry{ctx: x.ctx.sibling}
}

func (x *DataSetEntry) Index() int {
	return int(x.ctx.index)
}

func (x *DataSetEntry) ComponentName() string {
	return C.GoString(x.ctx.componentName)
}

func (x *DataSetEntry) VariableName() string {
	return C.GoString(x.ctx.variableName)
}

type ReportControlBlock struct {
	ctx *C.ReportControlBlock
}

func (x *ReportControlBlock) Parent() *LogicalNode {
	return &LogicalNode{ctx: x.ctx.parent}
}

func (x *ReportControlBlock) RptId() string {
	return C.GoString(x.ctx.rptId)
}

func (x *ReportControlBlock) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *ReportControlBlock) Buffered() bool {
	return bool(x.ctx.buffered)
}

func (x *ReportControlBlock) DataSetName() string {
	return C.GoString(x.ctx.dataSetName)
}

func (x *ReportControlBlock) ConfRef() uint32 {
	return uint32(x.ctx.confRef)
}

func (x *ReportControlBlock) TrgOps() uint8 {
	return uint8(x.ctx.trgOps)
}

func (x *ReportControlBlock) OptFlds() uint8 {
	return uint8(x.ctx.options)
}

func (x *ReportControlBlock) BufferTime() uint32 {
	return uint32(x.ctx.bufferTime)
}

func (x *ReportControlBlock) IntPeriod() uint32 {
	return uint32(x.ctx.intPeriod)
}

func (x *ReportControlBlock) ClientReservation() [17]uint8 {
	res := [17]uint8{}
	for i := 0; i < 17; i++ {
		res[i] = uint8(x.ctx.clientReservation[i])
	}
	return res
}

func (x *ReportControlBlock) Sibling() *ReportControlBlock {
	return &ReportControlBlock{ctx: x.ctx.sibling}
}

type SettingGroupControlBlock struct {
	ctx *C.SettingGroupControlBlock
}

func (x *SettingGroupControlBlock) Parent() *LogicalNode {
	return &LogicalNode{ctx: x.ctx.parent}
}

func (x *SettingGroupControlBlock) ActSG() uint8 {
	return uint8(x.ctx.actSG)
}

func (x *SettingGroupControlBlock) NumOfSGs() uint8 {
	return uint8(x.ctx.numOfSGs)
}

func (x *SettingGroupControlBlock) EditSG() uint8 {
	return uint8(x.ctx.editSG)
}

func (x *SettingGroupControlBlock) CnfEdit() bool {
	return bool(x.ctx.cnfEdit)
}

func (x *SettingGroupControlBlock) Timestamp() uint64 {
	return uint64(x.ctx.timestamp)
}

func (x *SettingGroupControlBlock) ResvTms() uint16 {
	return uint16(x.ctx.resvTms)
}

func (x *SettingGroupControlBlock) Sibling() *SettingGroupControlBlock {
	return &SettingGroupControlBlock{ctx: x.ctx.sibling}
}

type GSEControlBlock struct {
	ctx *C.GSEControlBlock
}

func (x *GSEControlBlock) Parent() *LogicalNode {
	return &LogicalNode{ctx: x.ctx.parent}
}

func (x *GSEControlBlock) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *GSEControlBlock) AppId() string {
	return C.GoString(x.ctx.appId)
}

func (x *GSEControlBlock) DataSetName() string {
	return C.GoString(x.ctx.dataSetName)
}

func (x *GSEControlBlock) ConfRev() uint32 {
	return uint32(x.ctx.confRev)
}

func (x *GSEControlBlock) FixedOffs() bool {
	return bool(x.ctx.fixedOffs)
}

func (x *GSEControlBlock) Address() *PhyComAddress {
	return &PhyComAddress{ctx: x.ctx.address}
}

func (x *GSEControlBlock) MinTime() int32 {
	return int32(x.ctx.minTime)
}

func (x *GSEControlBlock) MaxTime() int32 {
	return int32(x.ctx.maxTime)
}

func (x *GSEControlBlock) Sibling() *GSEControlBlock {
	return &GSEControlBlock{ctx: x.ctx.sibling}
}

type SVControlBlock struct {
	ctx *C.SVControlBlock
}

func (x *SVControlBlock) Parent() *LogicalNode {
	return &LogicalNode{ctx: x.ctx.parent}
}

func (x *SVControlBlock) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *SVControlBlock) SvId() string {
	return C.GoString(x.ctx.svId)
}

func (x *SVControlBlock) DataSetName() string {
	return C.GoString(x.ctx.dataSetName)
}

func (x *SVControlBlock) OptFlds() uint8 {
	return uint8(x.ctx.optFlds)
}

func (x *SVControlBlock) SmpMod() uint8 {
	return uint8(x.ctx.smpMod)
}

func (x *SVControlBlock) SmpRate() uint16 {
	return uint16(x.ctx.smpRate)
}

func (x *SVControlBlock) ConfRev() uint32 {
	return uint32(x.ctx.confRev)
}

func (x *SVControlBlock) DstAddress() *PhyComAddress {
	return &PhyComAddress{ctx: x.ctx.dstAddress}
}

func (x *SVControlBlock) IsUnicast() bool {
	return bool(x.ctx.isUnicast)
}

func (x *SVControlBlock) NoASDU() int {
	return int(x.ctx.noASDU)
}

func (x *SVControlBlock) Sibling() *SVControlBlock {
	return &SVControlBlock{ctx: x.ctx.sibling}
}

type LogControlBlock struct {
	ctx *C.LogControlBlock
}

func (x *LogControlBlock) Parent() *LogicalNode {
	return &LogicalNode{ctx: x.ctx.parent}
}

func (x *LogControlBlock) Sibling() *LogControlBlock {
	return &LogControlBlock{ctx: x.ctx.sibling}
}

func (x *LogControlBlock) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *LogControlBlock) LogRef() string {
	return C.GoString(x.ctx.logRef)
}

func (x *LogControlBlock) TrgOps() uint8 {
	return uint8(x.ctx.trgOps)
}

func (x *LogControlBlock) IntPeriod() uint32 {
	return uint32(x.ctx.intPeriod)
}

func (x *LogControlBlock) LogEna() bool {
	return bool(x.ctx.logEna)
}

func (x *LogControlBlock) ReasonCode() bool {
	return bool(x.ctx.reasonCode)
}

func (x *LogControlBlock) DataSetName() string {
	return C.GoString(x.ctx.dataSetName)
}

type Log struct {
	ctx *C.Log
}

func (x *Log) Parent() *LogicalNode {
	return &LogicalNode{ctx: x.ctx.parent}
}

func (x *Log) Sibling() *Log {
	return &Log{ctx: x.ctx.sibling}
}

func (x *Log) Name() string {
	return C.GoString(x.ctx.name)
}

type DataAttributeType int32

const (
	IEC61850_UNKNOWN_TYPE       DataAttributeType = C.IEC61850_UNKNOWN_TYPE
	IEC61850_BOOLEAN            DataAttributeType = C.IEC61850_BOOLEAN
	IEC61850_INT8               DataAttributeType = C.IEC61850_INT8
	IEC61850_INT16              DataAttributeType = C.IEC61850_INT16
	IEC61850_INT32              DataAttributeType = C.IEC61850_INT32
	IEC61850_INT64              DataAttributeType = C.IEC61850_INT64
	IEC61850_INT128             DataAttributeType = C.IEC61850_INT128
	IEC61850_INT8U              DataAttributeType = C.IEC61850_INT8U
	IEC61850_INT16U             DataAttributeType = C.IEC61850_INT16U
	IEC61850_INT24U             DataAttributeType = C.IEC61850_INT24U
	IEC61850_INT32U             DataAttributeType = C.IEC61850_INT32U
	IEC61850_FLOAT32            DataAttributeType = C.IEC61850_FLOAT32
	IEC61850_FLOAT64            DataAttributeType = C.IEC61850_FLOAT64
	IEC61850_ENUMERATED         DataAttributeType = C.IEC61850_ENUMERATED
	IEC61850_OCTET_STRING_64    DataAttributeType = C.IEC61850_OCTET_STRING_64
	IEC61850_OCTET_STRING_6     DataAttributeType = C.IEC61850_OCTET_STRING_6
	IEC61850_OCTET_STRING_8     DataAttributeType = C.IEC61850_OCTET_STRING_8
	IEC61850_VISIBLE_STRING_32  DataAttributeType = C.IEC61850_VISIBLE_STRING_32
	IEC61850_VISIBLE_STRING_64  DataAttributeType = C.IEC61850_VISIBLE_STRING_64
	IEC61850_VISIBLE_STRING_65  DataAttributeType = C.IEC61850_VISIBLE_STRING_65
	IEC61850_VISIBLE_STRING_129 DataAttributeType = C.IEC61850_VISIBLE_STRING_129
	IEC61850_VISIBLE_STRING_255 DataAttributeType = C.IEC61850_VISIBLE_STRING_255
	IEC61850_UNICODE_STRING_255 DataAttributeType = C.IEC61850_UNICODE_STRING_255
	IEC61850_TIMESTAMP          DataAttributeType = C.IEC61850_TIMESTAMP
	IEC61850_QUALITY            DataAttributeType = C.IEC61850_QUALITY
	IEC61850_CHECK              DataAttributeType = C.IEC61850_CHECK
	IEC61850_CODEDENUM          DataAttributeType = C.IEC61850_CODEDENUM
	IEC61850_GENERIC_BITSTRING  DataAttributeType = C.IEC61850_GENERIC_BITSTRING
	IEC61850_CONSTRUCTED        DataAttributeType = C.IEC61850_CONSTRUCTED
	IEC61850_ENTRY_TIME         DataAttributeType = C.IEC61850_ENTRY_TIME
	IEC61850_PHYCOMADDR         DataAttributeType = C.IEC61850_PHYCOMADDR
	IEC61850_CURRENCY           DataAttributeType = C.IEC61850_CURRENCY
	IEC61850_OPTFLDS            DataAttributeType = C.IEC61850_OPTFLDS
	IEC61850_TRGOPS             DataAttributeType = C.IEC61850_TRGOPS
)

type ModelNodeType int32

const (
	LogicalDeviceModelType ModelNodeType = C.LogicalDeviceModelType
	LogicalNodeModelType   ModelNodeType = C.LogicalNodeModelType
	DataObjectModelType    ModelNodeType = C.DataObjectModelType
	DataAttributeModelType ModelNodeType = C.DataAttributeModelType
)
