package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_model.h"

*/
import "C"

type ModelNode struct {
	ModelType  ModelNodeType
	Name       string
	Parent     *ModelNode
	Sibling    *ModelNode
	FirstChild *ModelNode
}

type DataAttribute struct {
	ModelType  ModelNodeType
	Name       string
	Parent     *ModelNode
	Sibling    *ModelNode
	FirstChild *ModelNode

	ElementCount int /* value > 0 if this is an array */
	ArrayIndex   int /* value > -1 when this is an array element */

	Fc   FunctionalConstraint
	Type DataAttributeType

	TriggerOptions uint8 /* TRG_OPT_DATA_CHANGED | TRG_OPT_QUALITY_CHANGED | TRG_OPT_DATA_UPDATE */

	MmsValue *MmsValue

	Addr uint32 /* TODO remove in version 2.0 */
}

type DataObject struct {
	ModelType  ModelNodeType
	Name       string
	Parent     *ModelNode
	Sibling    *ModelNode
	FirstChild *ModelNode

	ElementCount int /* value > 0 if this is an array */
	ArrayIndex   int /* value > -1 when this is an array element */
}

type LogicalNode struct {
	ModelType  ModelNodeType
	Name       string
	Parent     *ModelNode
	Sibling    *ModelNode
	FirstChild *ModelNode
}

type LogicalDevice struct {
	ModelType  ModelNodeType
	Name       string
	Parent     *ModelNode
	Sibling    *ModelNode
	FirstChild *ModelNode
	LdName     string /* ldName (when using functional naming) */
}

type IedModel struct {
	Name        string
	FirstChild  *LogicalDevice
	DataSets    []DataSet
	Rcbs        []ReportControlBlock
	GseCBs      []GSEControlBlock
	SvCBs       []SVControlBlock
	Sgcbs       []SettingGroupControlBlock
	Lcbs        []LogControlBlock
	Logs        []Log
	Initializer func()
}

type DataSet struct {
	LogicalDeviceName string /* logical device instance name */
	Name              string /* eg. MMXU1$dataset1 */
	Fcdas             []DataSetEntry
	Sibling           *DataSet
}

type DataSetEntry struct {
	LogicalDeviceName            string /* logical device instance name */
	IsLDNameDynamicallyAllocated bool
	VariableName                 string
	Index                        int
	ComponentName                string
	Value                        *MmsValue
	Sibling                      *DataSetEntry
}

type ReportControlBlock struct {
	Parent      *LogicalNode
	Name        string
	RptId       string
	Buffered    bool
	DataSetName string /* pre loaded with relative name in logical node */

	ConfRef    uint32 /* ConfRef - configuration revision */
	TrgOps     uint8  /* TrgOps - trigger conditions */
	OptFlds    uint8  /* OptFlds */
	BufferTime uint32 /* BufTm - time to buffer events until a report is generated */
	IntPeriod  uint32 /* IntgPd - integrity period */

	/* type (first byte) and address of the pre-configured client
	   type can be one of (0 - no reservation, 4 - IPv4 client, 6 - IPv6 client) */
	ClientReservation [17]uint8 /* clientReservation */

	/*
	 * next control block in list or NULL if this is the last entry
	 * at runtime reuse as pointer to ReportControl instance!
	 **/
	Sibling *ReportControlBlock
}

type SettingGroupControlBlock struct {
	Parent    *LogicalNode
	Name      string
	ActSG     uint8 /* value from SCL file */
	NumOfSGs  uint8 /* value from SCL file */
	EditSG    uint8 /* 0 at power-up */
	CnfEdit   bool  /* false at power-up */
	Timestamp uint64
	ResvTms   uint16 /* value from SCL file */

	/*
	 * next control block in list or NULL if this is the last entry
	 * at runtime reuse as pointer to SettingGroupControl instance!
	 **/
	Sibling *SettingGroupControlBlock
}

type GSEControlBlock struct {
	Parent      *LogicalNode
	Name        string
	AppId       string
	DataSetName string           /* pre loaded with relative name in logical node */
	ConfRev     uint32           /* ConfRev - configuration revision */
	FixedOffs   bool             /* fixed offsets */
	Address     *PhyComAddress   /* GSE communication parameters */
	MinTime     int32            /* optional minTime parameter --> -1 if not present */
	MaxTime     int32            /* optional maxTime parameter --> -1 if not present */
	Sibling     *GSEControlBlock /* next control block in list or NULL if this is the last entry */
}

type SVControlBlock struct {
	Parent      *LogicalNode
	Name        string
	SvId        string          /* MsvUD/UsvID */
	DataSetName string          /* pre loaded with relative name in logical node */
	OptFlds     uint8           /* OptFlds */
	SmpMod      uint8           /* SmpMod */
	SmpRate     uint16          /* SmpRate */
	ConfRev     uint32          /* ConfRev - configuration revision */
	DstAddress  *PhyComAddress  /* SV communication parameters */
	IsUnicast   bool            /* true if unicast */
	NoASDU      int             /* number of ASDU in a message */
	Sibling     *SVControlBlock /* next control block in list or NULL if this is the last entry */
}

type LogControlBlock struct {
	Parent      *LogicalNode
	Name        string
	DataSetName string           /* pre loaded with relative name in logical node */
	LogRef      string           /* object reference to the journal */
	TrgOps      uint8            /* TrgOps - trigger conditions */
	IntPeriod   uint32           /* IntgPd - integrity period */
	LogEna      bool             /* enable log by default */
	ReasonCode  bool             /* include reason code in log */
	Sibling     *LogControlBlock /* next control block in list or NULL if this is the last entry */
}

type Log struct {
	Parent  *LogicalNode
	Name    string
	Sibling *Log
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
