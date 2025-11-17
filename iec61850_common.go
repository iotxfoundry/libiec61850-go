package libiec61850go

/*
#include <stdlib.h>
#include "iec61850_common.h"

*/
import "C"
import (
	"time"
	"unsafe"
)

const (
	IEC_61850_EDITION_1   = C.IEC_61850_EDITION_1
	IEC_61850_EDITION_2   = C.IEC_61850_EDITION_2
	IEC_61850_EDITION_2_1 = C.IEC_61850_EDITION_2_1
)

type PhyComAddress struct {
	ctx *C.PhyComAddress
}

func (x *PhyComAddress) VlanPriority() uint8 {
	return uint8(x.ctx.vlanPriority)
}

func (x *PhyComAddress) VlanId() uint16 {
	return uint16(x.ctx.vlanId)
}

func (x *PhyComAddress) AppId() uint16 {
	return uint16(x.ctx.appId)
}

func (x *PhyComAddress) DstAddress() [6]uint8 {
	res := [6]uint8{}
	for i := 0; i < 6; i++ {
		res[i] = uint8(x.ctx.dstAddress[i])
	}
	return res
}

type ACSIClass int32

const (
	ACSI_CLASS_DATA_OBJECT ACSIClass = C.ACSI_CLASS_DATA_OBJECT
	ACSI_CLASS_DATA_SET    ACSIClass = C.ACSI_CLASS_DATA_SET
	ACSI_CLASS_BRCB        ACSIClass = C.ACSI_CLASS_BRCB
	ACSI_CLASS_URCB        ACSIClass = C.ACSI_CLASS_URCB
	ACSI_CLASS_LCB         ACSIClass = C.ACSI_CLASS_LCB
	ACSI_CLASS_LOG         ACSIClass = C.ACSI_CLASS_LOG
	ACSI_CLASS_SGCB        ACSIClass = C.ACSI_CLASS_SGCB
	ACSI_CLASS_GoCB        ACSIClass = C.ACSI_CLASS_GoCB
	ACSI_CLASS_GsCB        ACSIClass = C.ACSI_CLASS_GsCB
	ACSI_CLASS_MSVCB       ACSIClass = C.ACSI_CLASS_MSVCB
	ACSI_CLASS_USVCB       ACSIClass = C.ACSI_CLASS_USVCB
)

type ControlModel int32

const (
	/**
	 * No support for control functions. Control object only support status information.
	 */
	CONTROL_MODEL_STATUS_ONLY ControlModel = C.CONTROL_MODEL_STATUS_ONLY

	/**
	 * Direct control with normal security: Supports Operate, TimeActivatedOperate (optional),
	 * and Cancel (optional).
	 */
	CONTROL_MODEL_DIRECT_NORMAL ControlModel = C.CONTROL_MODEL_DIRECT_NORMAL

	/**
	 * Select before operate (SBO) with normal security: Supports Select, Operate, TimeActivatedOperate (optional),
	 * and Cancel (optional).
	 */
	CONTROL_MODEL_SBO_NORMAL ControlModel = C.CONTROL_MODEL_SBO_NORMAL

	/**
	 * Direct control with enhanced security (enhanced security includes the CommandTermination service)
	 */
	CONTROL_MODEL_DIRECT_ENHANCED ControlModel = C.CONTROL_MODEL_DIRECT_ENHANCED

	/**
	 * Select before operate (SBO) with enhanced security (enhanced security includes the CommandTermination service)
	 */
	CONTROL_MODEL_SBO_ENHANCED ControlModel = C.CONTROL_MODEL_SBO_ENHANCED
)

const (
	/** Report will be triggered when data changes */
	TRG_OPT_DATA_CHANGED = C.TRG_OPT_DATA_CHANGED
	/** Report will be triggered when quality changes */
	TRG_OPT_QUALITY_CHANGED = C.TRG_OPT_QUALITY_CHANGED
	/** Report will be triggered when data is updated */
	TRG_OPT_DATA_UPDATE = C.TRG_OPT_DATA_UPDATE
	/** Report will be triggered periodically */
	TRG_OPT_INTEGRITY = C.TRG_OPT_INTEGRITY
	/** Report will be triggered by GI (general interrogation) request */
	TRG_OPT_GI = C.TRG_OPT_GI
	/** Report will be triggered only on rising edge (transient variable */
	TRG_OPT_TRANSIENT = C.TRG_OPT_TRANSIENT
)

const (
	/** Report contains sequence number */
	RPT_OPT_SEQ_NUM = C.RPT_OPT_SEQ_NUM
	/** Report contains a report timestamp */
	RPT_OPT_TIME_STAMP = C.RPT_OPT_TIME_STAMP
	/** Report contains reason for inclusion value for each included data set member */
	RPT_OPT_REASON_FOR_INCLUSION = C.RPT_OPT_REASON_FOR_INCLUSION
	/** Report contains data set object reference */
	RPT_OPT_DATA_SET = C.RPT_OPT_DATA_SET
	/** Report contains data reference for each included data set member */
	RPT_OPT_DATA_REFERENCE = C.RPT_OPT_DATA_REFERENCE
	/** Report contains buffer overflow flag */
	RPT_OPT_BUFFER_OVERFLOW = C.RPT_OPT_BUFFER_OVERFLOW
	/** Report contains entry id */
	RPT_OPT_ENTRY_ID = C.RPT_OPT_ENTRY_ID
	/** Report contains configuration revision */
	RPT_OPT_CONF_REV = C.RPT_OPT_CONF_REV
)

const (

	/** Not supported - should not be used */
	CONTROL_ORCAT_NOT_SUPPORTED = C.CONTROL_ORCAT_NOT_SUPPORTED

	/** Control operation issued from an operator using a client located at bay level */
	CONTROL_ORCAT_BAY_CONTROL = C.CONTROL_ORCAT_BAY_CONTROL

	/** Control operation issued from an operator using a client located at station level */
	CONTROL_ORCAT_STATION_CONTROL = C.CONTROL_ORCAT_STATION_CONTROL

	/** Control operation from a remote operator outside the substation (for example network control center) */
	CONTROL_ORCAT_REMOTE_CONTROL = C.CONTROL_ORCAT_REMOTE_CONTROL

	/** Control operation issued from an automatic function at bay level */
	CONTROL_ORCAT_AUTOMATIC_BAY = C.CONTROL_ORCAT_AUTOMATIC_BAY

	/** Control operation issued from an automatic function at station level */
	CONTROL_ORCAT_AUTOMATIC_STATION = C.CONTROL_ORCAT_AUTOMATIC_STATION

	/** Control operation issued from a automatic function outside of the substation */
	CONTROL_ORCAT_AUTOMATIC_REMOTE = C.CONTROL_ORCAT_AUTOMATIC_REMOTE

	/** Control operation issued from a maintenance/service tool */
	CONTROL_ORCAT_MAINTENANCE = C.CONTROL_ORCAT_MAINTENANCE

	/** Status change occurred without control action (for example external trip of a circuit breaker or failure inside the breaker) */
	CONTROL_ORCAT_PROCESS = C.CONTROL_ORCAT_PROCESS
)

type ControlAddCause int32

const (
	ADD_CAUSE_UNKNOWN                        ControlAddCause = C.ADD_CAUSE_UNKNOWN
	ADD_CAUSE_NOT_SUPPORTED                  ControlAddCause = C.ADD_CAUSE_NOT_SUPPORTED
	ADD_CAUSE_BLOCKED_BY_SWITCHING_HIERARCHY ControlAddCause = C.ADD_CAUSE_BLOCKED_BY_SWITCHING_HIERARCHY
	ADD_CAUSE_SELECT_FAILED                  ControlAddCause = C.ADD_CAUSE_SELECT_FAILED
	ADD_CAUSE_INVALID_POSITION               ControlAddCause = C.ADD_CAUSE_INVALID_POSITION
	ADD_CAUSE_POSITION_REACHED               ControlAddCause = C.ADD_CAUSE_POSITION_REACHED
	ADD_CAUSE_PARAMETER_CHANGE_IN_EXECUTION  ControlAddCause = C.ADD_CAUSE_PARAMETER_CHANGE_IN_EXECUTION
	ADD_CAUSE_STEP_LIMIT                     ControlAddCause = C.ADD_CAUSE_STEP_LIMIT
	ADD_CAUSE_BLOCKED_BY_MODE                ControlAddCause = C.ADD_CAUSE_BLOCKED_BY_MODE
	ADD_CAUSE_BLOCKED_BY_PROCESS             ControlAddCause = C.ADD_CAUSE_BLOCKED_BY_PROCESS
	ADD_CAUSE_BLOCKED_BY_INTERLOCKING        ControlAddCause = C.ADD_CAUSE_BLOCKED_BY_INTERLOCKING
	ADD_CAUSE_BLOCKED_BY_SYNCHROCHECK        ControlAddCause = C.ADD_CAUSE_BLOCKED_BY_SYNCHROCHECK
	ADD_CAUSE_COMMAND_ALREADY_IN_EXECUTION   ControlAddCause = C.ADD_CAUSE_COMMAND_ALREADY_IN_EXECUTION
	ADD_CAUSE_BLOCKED_BY_HEALTH              ControlAddCause = C.ADD_CAUSE_BLOCKED_BY_HEALTH
	ADD_CAUSE_1_OF_N_CONTROL                 ControlAddCause = C.ADD_CAUSE_1_OF_N_CONTROL
	ADD_CAUSE_ABORTION_BY_CANCEL             ControlAddCause = C.ADD_CAUSE_ABORTION_BY_CANCEL
	ADD_CAUSE_TIME_LIMIT_OVER                ControlAddCause = C.ADD_CAUSE_TIME_LIMIT_OVER
	ADD_CAUSE_ABORTION_BY_TRIP               ControlAddCause = C.ADD_CAUSE_ABORTION_BY_TRIP
	ADD_CAUSE_OBJECT_NOT_SELECTED            ControlAddCause = C.ADD_CAUSE_OBJECT_NOT_SELECTED
	ADD_CAUSE_OBJECT_ALREADY_SELECTED        ControlAddCause = C.ADD_CAUSE_OBJECT_ALREADY_SELECTED
	ADD_CAUSE_NO_ACCESS_AUTHORITY            ControlAddCause = C.ADD_CAUSE_NO_ACCESS_AUTHORITY
	ADD_CAUSE_ENDED_WITH_OVERSHOOT           ControlAddCause = C.ADD_CAUSE_ENDED_WITH_OVERSHOOT
	ADD_CAUSE_ABORTION_DUE_TO_DEVIATION      ControlAddCause = C.ADD_CAUSE_ABORTION_DUE_TO_DEVIATION
	ADD_CAUSE_ABORTION_BY_COMMUNICATION_LOSS ControlAddCause = C.ADD_CAUSE_ABORTION_BY_COMMUNICATION_LOSS
	ADD_CAUSE_ABORTION_BY_COMMAND            ControlAddCause = C.ADD_CAUSE_ABORTION_BY_COMMAND
	ADD_CAUSE_NONE                           ControlAddCause = C.ADD_CAUSE_NONE
	ADD_CAUSE_INCONSISTENT_PARAMETERS        ControlAddCause = C.ADD_CAUSE_INCONSISTENT_PARAMETERS
	ADD_CAUSE_LOCKED_BY_OTHER_CLIENT         ControlAddCause = C.ADD_CAUSE_LOCKED_BY_OTHER_CLIENT
)

type ControlLastApplError int32

const (
	CONTROL_ERROR_NO_ERROR      ControlLastApplError = C.CONTROL_ERROR_NO_ERROR
	CONTROL_ERROR_UNKNOWN       ControlLastApplError = C.CONTROL_ERROR_UNKNOWN
	CONTROL_ERROR_TIMEOUT_TEST  ControlLastApplError = C.CONTROL_ERROR_TIMEOUT_TEST
	CONTROL_ERROR_OPERATOR_TEST ControlLastApplError = C.CONTROL_ERROR_OPERATOR_TEST
)

type FunctionalConstraint int32

const (
	/** Status information */
	IEC61850_FC_ST FunctionalConstraint = C.IEC61850_FC_ST
	/** Measurands - analog values */
	IEC61850_FC_MX FunctionalConstraint = C.IEC61850_FC_MX
	/** Setpoint */
	IEC61850_FC_SP FunctionalConstraint = C.IEC61850_FC_SP
	/** Substitution */
	IEC61850_FC_SV FunctionalConstraint = C.IEC61850_FC_SV
	/** Configuration */
	IEC61850_FC_CF FunctionalConstraint = C.IEC61850_FC_CF
	/** Description */
	IEC61850_FC_DC FunctionalConstraint = C.IEC61850_FC_DC
	/** Setting group */
	IEC61850_FC_SG FunctionalConstraint = C.IEC61850_FC_SG
	/** Setting group editable */
	IEC61850_FC_SE FunctionalConstraint = C.IEC61850_FC_SE
	/** Service response / Service tracking */
	IEC61850_FC_SR FunctionalConstraint = C.IEC61850_FC_SR
	/** Operate received */
	IEC61850_FC_OR FunctionalConstraint = C.IEC61850_FC_OR
	/** Blocking */
	IEC61850_FC_BL FunctionalConstraint = C.IEC61850_FC_BL
	/** Extended definition */
	IEC61850_FC_EX FunctionalConstraint = C.IEC61850_FC_EX
	/** Control */
	IEC61850_FC_CO FunctionalConstraint = C.IEC61850_FC_CO
	/** Unicast SV */
	IEC61850_FC_US FunctionalConstraint = C.IEC61850_FC_US
	/** Multicast SV */
	IEC61850_FC_MS FunctionalConstraint = C.IEC61850_FC_MS
	/** Unbuffered report */
	IEC61850_FC_RP FunctionalConstraint = C.IEC61850_FC_RP
	/** Buffered report */
	IEC61850_FC_BR FunctionalConstraint = C.IEC61850_FC_BR
	/** Log control blocks */
	IEC61850_FC_LG FunctionalConstraint = C.IEC61850_FC_LG
	/** Goose control blocks */
	IEC61850_FC_GO FunctionalConstraint = C.IEC61850_FC_GO

	/** All FCs - wildcard value */
	IEC61850_FC_ALL  FunctionalConstraint = C.IEC61850_FC_ALL
	IEC61850_FC_NONE FunctionalConstraint = C.IEC61850_FC_NONE
)

func (x FunctionalConstraint) String() string {
	return C.GoString(C.FunctionalConstraint_toString(C.FunctionalConstraint(x)))
}

func NewFunctionalConstraint(s string) FunctionalConstraint {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return FunctionalConstraint(C.FunctionalConstraint_fromString(cs))
}

type Quality uint16
type Validity uint16

const (
	QUALITY_VALIDITY_GOOD         Quality = C.QUALITY_VALIDITY_GOOD
	QUALITY_VALIDITY_INVALID      Quality = C.QUALITY_VALIDITY_INVALID
	QUALITY_VALIDITY_RESERVED     Quality = C.QUALITY_VALIDITY_RESERVED
	QUALITY_VALIDITY_QUESTIONABLE Quality = C.QUALITY_VALIDITY_QUESTIONABLE
	QUALITY_DETAIL_OVERFLOW       Quality = C.QUALITY_DETAIL_OVERFLOW
	QUALITY_DETAIL_OUT_OF_RANGE   Quality = C.QUALITY_DETAIL_OUT_OF_RANGE
	QUALITY_DETAIL_BAD_REFERENCE  Quality = C.QUALITY_DETAIL_BAD_REFERENCE
	QUALITY_DETAIL_OSCILLATORY    Quality = C.QUALITY_DETAIL_OSCILLATORY
	QUALITY_DETAIL_FAILURE        Quality = C.QUALITY_DETAIL_FAILURE
	QUALITY_DETAIL_OLD_DATA       Quality = C.QUALITY_DETAIL_OLD_DATA
	QUALITY_DETAIL_INCONSISTENT   Quality = C.QUALITY_DETAIL_INCONSISTENT
	QUALITY_DETAIL_INACCURATE     Quality = C.QUALITY_DETAIL_INACCURATE
	QUALITY_SOURCE_SUBSTITUTED    Quality = C.QUALITY_SOURCE_SUBSTITUTED
	QUALITY_TEST                  Quality = C.QUALITY_TEST
	QUALITY_OPERATOR_BLOCKED      Quality = C.QUALITY_OPERATOR_BLOCKED
	QUALITY_DERIVED               Quality = C.QUALITY_DERIVED
)

func NewQualityFromMmsValue(mmsValue *MmsValue) Quality {
	return Quality(C.Quality_fromMmsValue(mmsValue.ctx))
}

func (x *Quality) GetValidity() Validity {
	return Validity(C.Quality_getValidity((*C.Quality)(unsafe.Pointer(x))))
}

func (x *Quality) SetValidity(validity Validity) {
	C.Quality_setValidity((*C.Quality)(unsafe.Pointer(x)), C.Validity(validity))
}

func (x *Quality) SetFlag(flag int) {
	C.Quality_setFlag((*C.Quality)(unsafe.Pointer(x)), C.int(flag))
}

func (x *Quality) UnsetFlag(flag int) {
	C.Quality_unsetFlag((*C.Quality)(unsafe.Pointer(x)), C.int(flag))
}

func (x *Quality) IsFlagSet(flag int) bool {
	return bool(C.Quality_isFlagSet((*C.Quality)(unsafe.Pointer(x)), C.int(flag)))
}

func (x *Quality) ToMmsValue(mmsValue *MmsValue) {
	C.Quality_toMmsValue((*C.Quality)(unsafe.Pointer(x)), mmsValue.ctx)
}

type Dbpos int32

const (
	DBPOS_INTERMEDIATE_STATE Dbpos = C.DBPOS_INTERMEDIATE_STATE
	DBPOS_OFF                Dbpos = C.DBPOS_OFF
	DBPOS_ON                 Dbpos = C.DBPOS_ON
	DBPOS_BAD_STATE          Dbpos = C.DBPOS_BAD_STATE
)

func NewDbposFromMmsValue(mmsValue *MmsValue) Dbpos {
	return Dbpos(C.Dbpos_fromMmsValue(mmsValue.ctx))
}

func (x Dbpos) ToMmsValue(mmsValue *MmsValue) (out *MmsValue) {
	out = &MmsValue{ctx: mmsValue.ctx}
	out.ctx = C.Dbpos_toMmsValue(mmsValue.ctx, C.Dbpos(x))
	return
}

type Timestamp struct {
	ctx *C.Timestamp
}

func TimestampCreate() *Timestamp {
	return &Timestamp{ctx: C.Timestamp_create()}
}

func NewTimestampFromByteArray(byteArray []byte) *Timestamp {
	array := C.CBytes(byteArray)
	defer C.free(unsafe.Pointer(array))
	return &Timestamp{ctx: C.Timestamp_createFromByteArray((*C.uint8_t)(array))}
}

func (x *Timestamp) Destroy() {
	C.Timestamp_destroy(x.ctx)
}

func (x *Timestamp) ClearFlags() {
	C.Timestamp_clearFlags(x.ctx)
}

func (x *Timestamp) GetTimeInSeconds() time.Duration {
	return time.Duration(C.Timestamp_getTimeInSeconds(x.ctx)) * time.Second
}

func (x *Timestamp) GetTimeInMs() time.Duration {
	return time.Duration(C.Timestamp_getTimeInMs(x.ctx)) * time.Millisecond
}

func (x *Timestamp) GetTimeInNs() time.Duration {
	return time.Duration(C.Timestamp_getTimeInNs(x.ctx)) * time.Nanosecond
}

func (x *Timestamp) IsLeapSecondKnown() bool {
	return bool(C.Timestamp_isLeapSecondKnown(x.ctx))
}

func (x *Timestamp) SetLeapSecondKnown(value bool) {
	C.Timestamp_setLeapSecondKnown(x.ctx, C.bool(value))
}

func (x *Timestamp) HasClockFailure() bool {
	return bool(C.Timestamp_hasClockFailure(x.ctx))
}

func (x *Timestamp) SetClockFailure(value bool) {
	C.Timestamp_setClockFailure(x.ctx, C.bool(value))
}

func (x *Timestamp) IsClockNotSynchronized() bool {
	return bool(C.Timestamp_isClockNotSynchronized(x.ctx))
}

func (x *Timestamp) SetClockNotSynchronized(value bool) {
	C.Timestamp_setClockNotSynchronized(x.ctx, C.bool(value))
}

func (x *Timestamp) GetSubsecondPrecision() int {
	return int(C.Timestamp_getSubsecondPrecision(x.ctx))
}

func (x *Timestamp) SetFractionOfSecondPart(value uint32) {
	C.Timestamp_setFractionOfSecondPart(x.ctx, C.uint32_t(value))
}

func (x *Timestamp) GetFractionOfSecondPart() int {
	return int(C.Timestamp_getFractionOfSecondPart(x.ctx))
}

func (x *Timestamp) GetFractionOfSecond() time.Duration {
	return time.Duration(C.Timestamp_getFractionOfSecond(x.ctx)) * time.Second
}

func (x *Timestamp) SetSubsecondPrecision(value int) {
	C.Timestamp_setSubsecondPrecision(x.ctx, C.int(value))
}

func (x *Timestamp) SetTimeInSeconds(value time.Duration) {
	C.Timestamp_setTimeInSeconds(x.ctx, C.uint32_t(value.Seconds()))
}

func (x *Timestamp) SetTimeInMilliseconds(value time.Duration) {
	C.Timestamp_setTimeInMilliseconds(x.ctx, C.msSinceEpoch(value.Milliseconds()))
}

func (x *Timestamp) SetTimeInNanoseconds(value time.Duration) {
	C.Timestamp_setTimeInNanoseconds(x.ctx, C.nsSinceEpoch(value.Nanoseconds()))
}

func (x *Timestamp) SetByMmsUtcTime(mmsValue *MmsValue) {
	C.Timestamp_setByMmsUtcTime(x.ctx, mmsValue.ctx)
}

func (x *Timestamp) ToMmsValue(mmsValue *MmsValue) {
	C.Timestamp_toMmsValue(x.ctx, mmsValue.ctx)
}

func (x *Timestamp) FromMmsValue(mmsValue *MmsValue) {
	C.Timestamp_fromMmsValue(x.ctx, mmsValue.ctx)
}

func GetVersionString() string {
	return C.GoString(C.LibIEC61850_getVersionString())
}
