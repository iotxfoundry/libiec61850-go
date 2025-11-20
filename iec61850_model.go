package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_model.h"

extern void fIedModelInitializerGo();
*/
import "C"
import (
	"sync"
	"unsafe"
)

type IModelNode interface {
	ModelType() ModelNodeType
	Name() string
	Parent() IModelNode
	Sibling() IModelNode
	FirstChild() IModelNode
	Context() unsafe.Pointer
}

type ModelNode struct {
	ctx *C.ModelNode
}

func NewModelNode() *ModelNode {
	return &ModelNode{ctx: &C.ModelNode{}}
}

func (x *ModelNode) Initialize(modelType ModelNodeType, name string, parent *ModelNode, sibling *ModelNode, firstChild *ModelNode) {
	var cparent *C.ModelNode
	if parent != nil {
		cparent = parent.ctx
	}
	var csibling *C.ModelNode
	if sibling != nil {
		csibling = sibling.ctx
	}
	var cfirstChild *C.ModelNode
	if firstChild != nil {
		cfirstChild = firstChild.ctx
	}
	if x.ctx == nil {
		x.ctx = &C.ModelNode{}
	}
	x.ctx.modelType = C.ModelNodeType(modelType)
	x.ctx.name = C.CString(name)
	x.ctx.parent = cparent
	x.ctx.sibling = csibling
	x.ctx.firstChild = cfirstChild
}

func (x *ModelNode) Context() unsafe.Pointer {
	return unsafe.Pointer(x.ctx)
}

func (x *ModelNode) GetChildCount() int {
	return int(C.ModelNode_getChildCount(x.ctx))
}

func (x *ModelNode) GetChild(name string) *ModelNode {
	return &ModelNode{ctx: C.ModelNode_getChild(x.ctx, StringData(name))}
}

func (x *ModelNode) GetChildWithIdx(idx int) *ModelNode {
	return &ModelNode{ctx: C.ModelNode_getChildWithIdx(x.ctx, C.int(idx))}
}

func (x *ModelNode) GetChildWithFc(name string, fc FunctionalConstraint) *ModelNode {
	return &ModelNode{ctx: C.ModelNode_getChildWithFc(x.ctx, StringData(name), C.FunctionalConstraint(fc))}
}

func (x *ModelNode) GetObjectReference(objectReference string) string {
	return C.GoString(C.ModelNode_getObjectReference(x.ctx, StringData(objectReference)))
}

func (x *ModelNode) GetObjectReferenceEx(objectReference string, withoutIedName bool) string {
	return C.GoString(C.ModelNode_getObjectReferenceEx(x.ctx, StringData(objectReference), C.bool(withoutIedName)))
}

func (x *ModelNode) GetType() ModelNodeType {
	return ModelNodeType(C.ModelNode_getType(x.ctx))
}

func (x *ModelNode) GetName() string {
	return C.GoString(C.ModelNode_getName(x.ctx))
}

func (x *ModelNode) GetParent() *ModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *ModelNode) GetChildren() *LinkedList {
	return &LinkedList{ctx: C.ModelNode_getChildren(x.ctx)}
}

func (x *ModelNode) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *ModelNode) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *ModelNode) Parent() IModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *ModelNode) Sibling() IModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *ModelNode) FirstChild() IModelNode {
	return &ModelNode{ctx: x.ctx.firstChild}
}

type DataAttribute struct {
	ctx *C.DataAttribute
}

func NewDataAttribute() *DataAttribute {
	return &DataAttribute{ctx: &C.DataAttribute{}}
}

func (x *DataAttribute) Initialize(
	modelType ModelNodeType,
	name string,
	parent IModelNode,
	sibling IModelNode,
	firstChild IModelNode,
	elementCount int,
	arrayIndex int,
	fc FunctionalConstraint,
	type_ DataAttributeType,
	triggerOptions uint8,
	mmsValue *MmsValue,
	addr uint32,
) {
	var cparent *C.ModelNode
	if parent != nil {
		cparent = (*C.ModelNode)(parent.Context())
	}
	var csibling *C.ModelNode
	if sibling != nil {
		csibling = (*C.ModelNode)(sibling.Context())
	}
	var cfirstChild *C.ModelNode
	if firstChild != nil {
		cfirstChild = (*C.ModelNode)(firstChild.Context())
	}
	if x.ctx == nil {
		x.ctx = &C.DataAttribute{}
	}
	var cmmsValue *C.MmsValue
	if mmsValue != nil {
		cmmsValue = mmsValue.ctx
	}
	x.ctx.modelType = C.ModelNodeType(modelType)
	x.ctx.name = C.CString(name)
	x.ctx.parent = cparent
	x.ctx.sibling = csibling
	x.ctx.firstChild = cfirstChild
	x.ctx.elementCount = C.int(elementCount)
	x.ctx.arrayIndex = C.int(arrayIndex)
	x.ctx.fc = C.FunctionalConstraint(fc)
	x.ctx._type = C.DataAttributeType(type_)
	x.ctx.triggerOptions = C.uint8_t(triggerOptions)
	x.ctx.mmsValue = cmmsValue
	x.ctx.sAddr = C.uint32_t(addr)
}

func (x *DataAttribute) Context() unsafe.Pointer {
	return unsafe.Pointer(x.ctx)
}

func (x *DataAttribute) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *DataAttribute) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *DataAttribute) Parent() IModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *DataAttribute) Sibling() IModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *DataAttribute) FirstChild() IModelNode {
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

func (x *DataAttribute) SetMmsValue(mmsValue *MmsValue) {
	if x.ctx.mmsValue != nil {
		(&MmsValue{ctx: x.ctx.mmsValue}).Delete()
		x.ctx.mmsValue = nil
	}
	x.ctx.mmsValue = mmsValue.ctx
}

func (x *DataAttribute) Addr() uint32 {
	return uint32(x.ctx.sAddr)
}

type DataObject struct {
	ctx *C.DataObject
}

func NewDataObject() *DataObject {
	return &DataObject{ctx: &C.DataObject{}}
}

func (x *DataObject) Initialize(
	modelType ModelNodeType,
	name string,
	parent IModelNode,
	sibling IModelNode,
	firstChild IModelNode,
	elementCount int,
	arrayIndex int,
) {
	var cparent *C.ModelNode
	if parent != nil {
		cparent = (*C.ModelNode)(parent.Context())
	}
	var csibling *C.ModelNode
	if sibling != nil {
		csibling = (*C.ModelNode)(sibling.Context())
	}
	var cfirstChild *C.ModelNode
	if firstChild != nil {
		cfirstChild = (*C.ModelNode)(firstChild.Context())
	}
	if x.ctx == nil {
		x.ctx = &C.DataObject{}
	}
	x.ctx.modelType = C.ModelNodeType(modelType)
	x.ctx.name = C.CString(name)
	x.ctx.parent = cparent
	x.ctx.sibling = csibling
	x.ctx.firstChild = cfirstChild
	x.ctx.elementCount = C.int(elementCount)
	x.ctx.arrayIndex = C.int(arrayIndex)
}

func (x *DataObject) Context() unsafe.Pointer {
	return unsafe.Pointer(x.ctx)
}

func (x *DataObject) HasFCData(fc FunctionalConstraint) bool {
	return bool(C.DataObject_hasFCData(x.ctx, C.FunctionalConstraint(fc)))
}

func (x *DataObject) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *DataObject) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *DataObject) Parent() IModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *DataObject) Sibling() IModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *DataObject) FirstChild() IModelNode {
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

func NewLogicalNode() *LogicalNode {
	return &LogicalNode{ctx: &C.LogicalNode{}}
}

func (x *LogicalNode) Initialize(
	modelType ModelNodeType,
	name string,
	parent IModelNode,
	sibling IModelNode,
	firstChild IModelNode,
) {
	var cparent *C.ModelNode
	if parent != nil {
		cparent = (*C.ModelNode)(parent.Context())
	}
	var csibling *C.ModelNode
	if sibling != nil {
		csibling = (*C.ModelNode)(sibling.Context())
	}
	var cfirstChild *C.ModelNode
	if firstChild != nil {
		cfirstChild = (*C.ModelNode)(firstChild.Context())
	}
	if x.ctx == nil {
		x.ctx = &C.LogicalNode{}
	}
	x.ctx.modelType = C.ModelNodeType(modelType)
	x.ctx.name = C.CString(name)
	x.ctx.parent = cparent
	x.ctx.sibling = csibling
	x.ctx.firstChild = cfirstChild
}

func (x *LogicalNode) Context() unsafe.Pointer {
	return unsafe.Pointer(x.ctx)
}

func (x *LogicalNode) HasFCData(fc FunctionalConstraint) bool {
	return bool(C.LogicalNode_hasFCData(x.ctx, C.FunctionalConstraint(fc)))
}

// func (x *LogicalNode) HasBufferedReports() bool {
// 	return bool(C.LogicalNode_hasBufferedReports(x.ctx))
// }

// func (x *LogicalNode) HasUnbufferedReports() bool {
// 	return bool(C.LogicalNode_hasUnbufferedReports(x.ctx))
// }

func (x *LogicalNode) GetDataSet(dataSetName string) *DataSet {
	return &DataSet{ctx: C.LogicalNode_getDataSet(x.ctx, StringData(dataSetName))}
}

func (x *LogicalNode) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *LogicalNode) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *LogicalNode) Parent() IModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *LogicalNode) Sibling() IModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *LogicalNode) FirstChild() IModelNode {
	return &ModelNode{ctx: x.ctx.firstChild}
}

type LogicalDevice struct {
	ctx *C.LogicalDevice
}

func NewLogicalDevice() *LogicalDevice {
	return &LogicalDevice{ctx: &C.LogicalDevice{}}
}

func (x *LogicalDevice) Initialize(
	modelType ModelNodeType,
	name string,
	parent IModelNode,
	sibling IModelNode,
	firstChild IModelNode,
	ldName *string,
) {
	var cparent *C.ModelNode
	if parent != nil {
		cparent = (*C.ModelNode)(parent.Context())
	}
	var csibling *C.ModelNode
	if sibling != nil {
		csibling = (*C.ModelNode)(sibling.Context())
	}
	var cfirstChild *C.ModelNode
	if firstChild != nil {
		cfirstChild = (*C.ModelNode)(firstChild.Context())
	}
	var cldName *C.char
	if ldName != nil {
		cldName = C.CString(*ldName)
	}
	if x.ctx == nil {
		x.ctx = &C.LogicalDevice{}
	}
	x.ctx.modelType = C.ModelNodeType(modelType)
	x.ctx.name = C.CString(name)
	x.ctx.parent = cparent
	x.ctx.sibling = csibling
	x.ctx.firstChild = cfirstChild
	x.ctx.ldName = cldName
}

func (x *LogicalDevice) Context() unsafe.Pointer {
	return unsafe.Pointer(x.ctx)
}

func (x *LogicalDevice) GetLogicalNodeCount() int {
	return int(C.LogicalDevice_getLogicalNodeCount(x.ctx))
}

func (x *LogicalDevice) GetChildByMmsVariableName(mmsVariableName string) *ModelNode {
	return &ModelNode{ctx: C.LogicalDevice_getChildByMmsVariableName(x.ctx, StringData(mmsVariableName))}
}

func (x *LogicalDevice) GetSettingGroupControlBlock() *SettingGroupControlBlock {
	return &SettingGroupControlBlock{ctx: C.LogicalDevice_getSettingGroupControlBlock(x.ctx)}
}

func (x *LogicalDevice) GetLogicalNode(lnName string) *LogicalNode {
	return &LogicalNode{ctx: C.LogicalDevice_getLogicalNode(x.ctx, StringData(lnName))}
}

func (x *LogicalDevice) ModelType() ModelNodeType {
	return ModelNodeType(x.ctx.modelType)
}

func (x *LogicalDevice) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *LogicalDevice) Parent() IModelNode {
	return &ModelNode{ctx: x.ctx.parent}
}

func (x *LogicalDevice) Sibling() IModelNode {
	return &ModelNode{ctx: x.ctx.sibling}
}

func (x *LogicalDevice) FirstChild() IModelNode {
	return &ModelNode{ctx: x.ctx.firstChild}
}

func (x *LogicalDevice) LdName() string {
	return C.GoString(x.ctx.ldName)
}

type IedModel struct {
	ctx *C.IedModel
}

func NewIedModel() *IedModel {
	return &IedModel{ctx: &C.IedModel{}}
}

func (x *IedModel) Initialize(name string,
	firstChild *LogicalDevice,
	dataSets *DataSet,
	rcbs *ReportControlBlock,
	gseCBs *GSEControlBlock,
	svCBs *SVControlBlock,
	sgcbs *SettingGroupControlBlock,
	lcbs *LogControlBlock,
	logs *Log,
	initializer func(),
) {
	var cfirstChild *C.LogicalDevice
	if firstChild != nil {
		cfirstChild = firstChild.ctx
	}
	var cdataSets *C.DataSet
	if dataSets != nil {
		cdataSets = dataSets.ctx
	}
	var crcbs *C.ReportControlBlock
	if rcbs != nil {
		crcbs = rcbs.ctx
	}
	var cgseCBs *C.GSEControlBlock
	if gseCBs != nil {
		cgseCBs = gseCBs.ctx
	}
	var csvCBs *C.SVControlBlock
	if svCBs != nil {
		csvCBs = svCBs.ctx
	}
	var csgcbs *C.SettingGroupControlBlock
	if sgcbs != nil {
		csgcbs = sgcbs.ctx
	}
	var clcbs *C.LogControlBlock
	if lcbs != nil {
		clcbs = lcbs.ctx
	}
	var clogs *C.Log
	if logs != nil {
		clogs = logs.ctx
	}
	if x.ctx == nil {
		x.ctx = &C.IedModel{}
	}
	x.ctx.name = C.CString(name)
	x.ctx.firstChild = cfirstChild
	x.ctx.dataSets = cdataSets
	x.ctx.rcbs = crcbs
	x.ctx.gseCBs = cgseCBs
	x.ctx.svCBs = csvCBs
	x.ctx.sgcbs = csgcbs
	x.ctx.lcbs = clcbs
	x.ctx.logs = clogs
	x.SetInitializer(initializer)
}

func (x *IedModel) Context() unsafe.Pointer {
	return unsafe.Pointer(x.ctx)
}

func (x *IedModel) ModelType() ModelNodeType {
	return -1
}

func (x *IedModel) Parent() IModelNode {
	return nil
}

func (x *IedModel) Sibling() IModelNode {
	return nil
}

func (x *IedModel) SetAttributeValuesToNull() {
	C.IedModel_setAttributeValuesToNull(x.ctx)
}

func (x *IedModel) GetDevice(ldName string) *LogicalDevice {
	return &LogicalDevice{ctx: C.IedModel_getDevice(x.ctx, StringData(ldName))}
}

func (x *IedModel) LookupDataSet(dataSetReference string) *DataSet {
	return &DataSet{ctx: C.IedModel_lookupDataSet(x.ctx, StringData(dataSetReference))}
}

func (x *IedModel) LookupDataAttributeByMmsValue(mmsValue *MmsValue) *DataAttribute {
	return &DataAttribute{ctx: C.IedModel_lookupDataAttributeByMmsValue(x.ctx, mmsValue.ctx)}
}

func (x *IedModel) GetLogicalDeviceCount() int {
	return int(C.IedModel_getLogicalDeviceCount(x.ctx))
}

func (x *IedModel) SetIedName(iedName string) {
	C.IedModel_setIedName(x.ctx, StringData(iedName))
}

func (x *IedModel) GetModelNodeByObjectReference(objectReference string) *ModelNode {
	return &ModelNode{ctx: C.IedModel_getModelNodeByObjectReference(x.ctx, StringData(objectReference))}
}

func (x *IedModel) GetSVControlBlock(parentLN *LogicalNode, svcbName string) *SVControlBlock {
	return &SVControlBlock{ctx: C.IedModel_getSVControlBlock(x.ctx, parentLN.ctx, StringData(svcbName))}
}

func (x *IedModel) GetModelNodeByShortObjectReference(objectReference string) *ModelNode {
	return &ModelNode{ctx: C.IedModel_getModelNodeByShortObjectReference(x.ctx, StringData(objectReference))}
}

func (x *IedModel) GetModelNodeByShortAddress(shortAddress uint32) *ModelNode {
	return &ModelNode{ctx: C.IedModel_getModelNodeByShortAddress(x.ctx, C.uint32_t(shortAddress))}
}

func (x *IedModel) GetDeviceByInst(ldInst string) *LogicalDevice {
	return &LogicalDevice{ctx: C.IedModel_getDeviceByInst(x.ctx, StringData(ldInst))}
}

func (x *IedModel) GetDeviceByIndex(index int) *LogicalDevice {
	return &LogicalDevice{ctx: C.IedModel_getDeviceByIndex(x.ctx, C.int(index))}
}

func (x *IedModel) Name() string {
	return C.GoString(x.ctx.name)
}

func (x *IedModel) FirstChild() IModelNode {
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
	mapIedModelInitializerCallbacks.Store(x, initializer)
	x.ctx.initializer = C.IedModelInitializer(C.fIedModelInitializerGo)
}

type DataSet struct {
	ctx *C.DataSet
}

func NewDataSet() *DataSet {
	return &DataSet{ctx: &C.DataSet{}}
}

func (x *DataSet) Initialize(logicalDeviceName string, name string, elementCount int, fcdas *DataSetEntry, sibling *DataSet) {
	var cfcdas *C.DataSetEntry
	if fcdas != nil {
		cfcdas = fcdas.ctx
	}
	var csibling *C.DataSet
	if sibling != nil {
		csibling = sibling.ctx
	}
	if x.ctx == nil {
		x.ctx = &C.DataSet{}
	}
	x.ctx.logicalDeviceName = C.CString(logicalDeviceName)
	x.ctx.name = C.CString(name)
	x.ctx.elementCount = C.int(elementCount)
	x.ctx.fcdas = cfcdas
	x.ctx.sibling = csibling
}

func (x *DataSet) Context() unsafe.Pointer {
	return unsafe.Pointer(x.ctx)
}

func (x *DataSet) ModelType() ModelNodeType {
	return -1
}

func (x *DataSet) Parent() IModelNode {
	return nil
}

func (x *DataSet) FirstChild() IModelNode {
	return nil
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

func (x *DataSet) Sibling() IModelNode {
	return &DataSet{ctx: x.ctx.sibling}
}

type DataSetEntry struct {
	ctx *C.DataSetEntry
}

func NewDataSetEntry() *DataSetEntry {
	return &DataSetEntry{ctx: &C.DataSetEntry{}}
}

func (x *DataSetEntry) Initialize(logicalDeviceName string, isLDNameDynamicallyAllocated bool, variableName string, index int, componentName *string, value *MmsValue, sibling *DataSetEntry) {
	var ccomponentName *C.char
	if componentName != nil {
		ccomponentName = C.CString(*componentName)
	}
	var cvalue *C.MmsValue
	if value != nil {
		cvalue = value.ctx
	}
	var csibling *C.DataSetEntry
	if sibling != nil {
		csibling = sibling.ctx
	}
	if x.ctx == nil {
		x.ctx = &C.DataSetEntry{}
	}
	x.ctx.logicalDeviceName = C.CString(logicalDeviceName)
	x.ctx.isLDNameDynamicallyAllocated = C.bool(isLDNameDynamicallyAllocated)
	x.ctx.variableName = C.CString(variableName)
	x.ctx.index = C.int(index)
	x.ctx.componentName = ccomponentName
	x.ctx.value = cvalue
	x.ctx.sibling = csibling
}

func (x *DataSetEntry) Context() unsafe.Pointer {
	return unsafe.Pointer(x.ctx)
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

func NewReportControlBlock() *ReportControlBlock {
	return &ReportControlBlock{ctx: &C.ReportControlBlock{}}
}

func (x *ReportControlBlock) Initialize(
	parent *LogicalNode,
	name string,
	rptId string,
	buffered bool,
	dataSetName string,
	confRef uint32,
	trgOps uint8,
	options uint8,
	bufferTime uint32,
	intPeriod uint32,
	clientReservation []byte,
	sibling *ReportControlBlock,
) {
	var cparent *C.LogicalNode
	if parent != nil {
		cparent = parent.ctx
	}
	var csibling *C.ReportControlBlock
	if sibling != nil {
		csibling = sibling.ctx
	}
	if x.ctx == nil {
		x.ctx = &C.ReportControlBlock{}
	}
	x.ctx.parent = cparent
	x.ctx.name = C.CString(name)
	x.ctx.rptId = C.CString(rptId)
	x.ctx.buffered = C.bool(buffered)
	x.ctx.dataSetName = C.CString(dataSetName)
	x.ctx.confRef = C.uint32_t(confRef)
	x.ctx.trgOps = C.uint8_t(trgOps)
	x.ctx.options = C.uint8_t(options)
	x.ctx.bufferTime = C.uint32_t(bufferTime)
	x.ctx.intPeriod = C.uint32_t(intPeriod)
	for i := 0; i < 17 && i < len(clientReservation); i++ {
		x.ctx.clientReservation[i] = C.uint8_t(clientReservation[i])
	}
	x.ctx.sibling = csibling
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
	for i := range 17 {
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
