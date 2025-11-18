package libiec61850go

/*
#include <stdlib.h>
#include "tls_config.h"

extern void fEventHandlerGo(void* parameter, TLSEventLevel eventLevel, int eventCode, char* message, TLSConnection con);
*/
import "C"
import (
	"sync"
	"unsafe"
)

type TLSConfiguration struct {
	ctx C.TLSConfiguration
}

func TLSConfigurationCreate() *TLSConfiguration {
	return &TLSConfiguration{
		ctx: C.TLSConfiguration_create(),
	}
}

func (x *TLSConfiguration) SetClientMode() {
	C.TLSConfiguration_setClientMode(x.ctx)
}

type TLSConfigVersion int32

const (
	TLS_VERSION_NOT_SELECTED TLSConfigVersion = C.TLS_VERSION_NOT_SELECTED
	TLS_VERSION_SSL_3_0      TLSConfigVersion = C.TLS_VERSION_SSL_3_0
	TLS_VERSION_TLS_1_0      TLSConfigVersion = C.TLS_VERSION_TLS_1_0
	TLS_VERSION_TLS_1_1      TLSConfigVersion = C.TLS_VERSION_TLS_1_1
	TLS_VERSION_TLS_1_2      TLSConfigVersion = C.TLS_VERSION_TLS_1_2
	TLS_VERSION_TLS_1_3      TLSConfigVersion = C.TLS_VERSION_TLS_1_3
)

func (x TLSConfigVersion) String() string {
	return C.GoString(C.TLSConfigVersion_toString(C.TLSConfigVersion(x)))
}

type TLSEventLevel int32

const (
	TLS_SEC_EVT_INFO     TLSEventLevel = C.TLS_SEC_EVT_INFO
	TLS_SEC_EVT_WARNING  TLSEventLevel = C.TLS_SEC_EVT_WARNING
	TLS_SEC_EVT_INCIDENT TLSEventLevel = C.TLS_SEC_EVT_INCIDENT
)

const (
	TLS_EVENT_CODE_ALM_ALGO_NOT_SUPPORTED              = C.TLS_EVENT_CODE_ALM_ALGO_NOT_SUPPORTED
	TLS_EVENT_CODE_ALM_UNSECURE_COMMUNICATION          = C.TLS_EVENT_CODE_ALM_UNSECURE_COMMUNICATION
	TLS_EVENT_CODE_ALM_CERT_UNAVAILABLE                = C.TLS_EVENT_CODE_ALM_CERT_UNAVAILABLE
	TLS_EVENT_CODE_ALM_BAD_CERT                        = C.TLS_EVENT_CODE_ALM_BAD_CERT
	TLS_EVENT_CODE_ALM_CERT_SIZE_EXCEEDED              = C.TLS_EVENT_CODE_ALM_CERT_SIZE_EXCEEDED
	TLS_EVENT_CODE_ALM_CERT_VALIDATION_FAILED          = C.TLS_EVENT_CODE_ALM_CERT_VALIDATION_FAILED
	TLS_EVENT_CODE_ALM_CERT_REQUIRED                   = C.TLS_EVENT_CODE_ALM_CERT_REQUIRED
	TLS_EVENT_CODE_ALM_HANDSHAKE_FAILED_UNKNOWN_REASON = C.TLS_EVENT_CODE_ALM_HANDSHAKE_FAILED_UNKNOWN_REASON
	TLS_EVENT_CODE_WRN_INSECURE_TLS_VERSION            = C.TLS_EVENT_CODE_WRN_INSECURE_TLS_VERSION
	TLS_EVENT_CODE_INF_SESSION_RENEGOTIATION           = C.TLS_EVENT_CODE_INF_SESSION_RENEGOTIATION
	TLS_EVENT_CODE_ALM_CERT_EXPIRED                    = C.TLS_EVENT_CODE_ALM_CERT_EXPIRED
	TLS_EVENT_CODE_ALM_CERT_REVOKED                    = C.TLS_EVENT_CODE_ALM_CERT_REVOKED
	TLS_EVENT_CODE_ALM_CERT_NOT_CONFIGURED             = C.TLS_EVENT_CODE_ALM_CERT_NOT_CONFIGURED
	TLS_EVENT_CODE_ALM_CERT_NOT_TRUSTED                = C.TLS_EVENT_CODE_ALM_CERT_NOT_TRUSTED
	TLS_EVENT_CODE_ALM_NO_CIPHER                       = C.TLS_EVENT_CODE_ALM_NO_CIPHER
	TLS_EVENT_CODE_INF_SESSION_ESTABLISHED             = C.TLS_EVENT_CODE_INF_SESSION_ESTABLISHED
	TLS_EVENT_CODE_WRN_CERT_EXPIRED                    = C.TLS_EVENT_CODE_WRN_CERT_EXPIRED
	TLS_EVENT_CODE_WRN_CERT_NOT_YET_VALID              = C.TLS_EVENT_CODE_WRN_CERT_NOT_YET_VALID
	TLS_EVENT_CODE_WRN_CRL_EXPIRED                     = C.TLS_EVENT_CODE_WRN_CRL_EXPIRED
	TLS_EVENT_CODE_WRN_CRL_NOT_YET_VALID               = C.TLS_EVENT_CODE_WRN_CRL_NOT_YET_VALID
)

type TLSConnection struct {
	ctx C.TLSConnection
}

func (x *TLSConnection) GetPeerAddress(peerAddrBuf string) string {
	out := C.TLSConnection_getPeerAddress(x.ctx, (*C.char)(unsafe.Pointer(unsafe.StringData(peerAddrBuf))))
	return C.GoString(out)
}

func (x *TLSConnection) GetPeerCertificate() []byte {
	certSize := int32(0)
	out := C.TLSConnection_getPeerCertificate(x.ctx, (*C.int)(unsafe.Pointer(&certSize)))
	return C.GoBytes(unsafe.Pointer(out), C.int(certSize))
}

func (x *TLSConnection) GetTLSVersion() TLSConfigVersion {
	return TLSConfigVersion(C.TLSConnection_getTLSVersion(x.ctx))
}

type EventHandler func(parameter unsafe.Pointer, eventLevel TLSEventLevel, eventCode int, message string, con *TLSConnection)

var mapEventHandlers = sync.Map{}

//export fEventHandlerGo
func fEventHandlerGo(parameter unsafe.Pointer, eventLevel C.TLSEventLevel, eventCode C.int, message *C.char, con C.TLSConnection) {
	mapEventHandlers.Range(func(k, v any) bool {
		if fn, ok := v.(EventHandler); ok {
			fn(parameter, TLSEventLevel(eventLevel), int(eventCode), C.GoString(message), &TLSConnection{ctx: con})
		}
		return true
	})
}

func (x *TLSConfiguration) SetEventHandler(handler EventHandler, parameter unsafe.Pointer) {
	C.TLSConfiguration_setEventHandler(x.ctx, C.TLSConfiguration_EventHandler(C.TLSConfiguration_EventHandler(C.fEventHandlerGo)), parameter)
	mapEventHandlers.Store(parameter, handler)
}

func (x *TLSConfiguration) EnableSessionResumption(enable bool) {
	C.TLSConfiguration_enableSessionResumption(x.ctx, C.bool(enable))
}

func (x *TLSConfiguration) SetSessionResumptionInterval(intervalInSeconds int) {
	C.TLSConfiguration_setSessionResumptionInterval(x.ctx, C.int(intervalInSeconds))
}

func (x *TLSConfiguration) SetChainValidation(value bool) {
	C.TLSConfiguration_setChainValidation(x.ctx, C.bool(value))
}

func (x *TLSConfiguration) SetTimeValidation(value bool) {
	C.TLSConfiguration_setTimeValidation(x.ctx, C.bool(value))
}

func (x *TLSConfiguration) SetAllowOnlyKnownCertificates(value bool) {
	C.TLSConfiguration_setAllowOnlyKnownCertificates(x.ctx, C.bool(value))
}

func (x *TLSConfiguration) SetOwnCertificate(certificate []byte) bool {
	return bool(C.TLSConfiguration_setOwnCertificate(x.ctx, (*C.uint8_t)(unsafe.SliceData(certificate)), C.int(len(certificate))))
}

func (x *TLSConfiguration) SetOwnCertificateFromFile(filename string) bool {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return bool(C.TLSConfiguration_setOwnCertificateFromFile(x.ctx, cfilename))
}

func (x *TLSConfiguration) SetOwnKey(key []byte, keyPassword string) bool {
	ckeyPassword := C.CString(keyPassword)
	defer C.free(unsafe.Pointer(ckeyPassword))
	return bool(C.TLSConfiguration_setOwnKey(x.ctx, (*C.uint8_t)(unsafe.SliceData(key)), C.int(len(key)), ckeyPassword))
}

func (x *TLSConfiguration) SetOwnKeyFromFile(filename string, keyPassword string) bool {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	ckeyPassword := C.CString(keyPassword)
	defer C.free(unsafe.Pointer(ckeyPassword))
	return bool(C.TLSConfiguration_setOwnKeyFromFile(x.ctx, cfilename, ckeyPassword))
}

func (x *TLSConfiguration) AddAllowedCertificate(certificate []byte) bool {
	return bool(C.TLSConfiguration_addAllowedCertificate(x.ctx, (*C.uint8_t)(unsafe.SliceData(certificate)), C.int(len(certificate))))
}

func (x *TLSConfiguration) AddAllowedCertificateFromFile(filename string) bool {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return bool(C.TLSConfiguration_addAllowedCertificateFromFile(x.ctx, cfilename))
}

func (x *TLSConfiguration) AddCACertificate(certificate []byte) bool {
	return bool(C.TLSConfiguration_addCACertificate(x.ctx, (*C.uint8_t)(unsafe.SliceData(certificate)), C.int(len(certificate))))
}

func (x *TLSConfiguration) AddCACertificateFromFile(filename string) bool {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return bool(C.TLSConfiguration_addCACertificateFromFile(x.ctx, cfilename))
}

func (x *TLSConfiguration) SetRenegotiationTime(timeInMs int) {
	C.TLSConfiguration_setRenegotiationTime(x.ctx, C.int(timeInMs))
}

func (x *TLSConfiguration) SetMinTlsVersion(version TLSConfigVersion) {
	C.TLSConfiguration_setMinTlsVersion(x.ctx, C.TLSConfigVersion(version))
}

func (x *TLSConfiguration) SetMaxTlsVersion(version TLSConfigVersion) {
	C.TLSConfiguration_setMaxTlsVersion(x.ctx, C.TLSConfigVersion(version))
}

func (x *TLSConfiguration) AddCRL(crl []byte) bool {
	return bool(C.TLSConfiguration_addCRL(x.ctx, (*C.uint8_t)(unsafe.SliceData(crl)), C.int(len(crl))))
}

func (x *TLSConfiguration) AddCRLFromFile(filename string) bool {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return bool(C.TLSConfiguration_addCRLFromFile(x.ctx, cfilename))
}

func (x *TLSConfiguration) ResetCRL() {
	C.TLSConfiguration_resetCRL(x.ctx)
}

func (x *TLSConfiguration) AddCipherSuite(ciphersuite int) {
	C.TLSConfiguration_addCipherSuite(x.ctx, C.int(ciphersuite))
}

func (x *TLSConfiguration) ClearCipherSuiteList() {
	C.TLSConfiguration_clearCipherSuiteList(x.ctx)
}

func (x *TLSConfiguration) Destroy() {
	C.TLSConfiguration_destroy(x.ctx)
}
