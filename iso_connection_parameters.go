package libiec61850go

/*
#include <stdlib.h>
#include "libiec61850_common_api.h"
#include "iso_connection_parameters.h"

extern bool fAcseAuthenticatorGo(void* parameter, AcseAuthenticationParameter authParameter, void** securityToken, IsoApplicationReference* appReference);
*/
import "C"
import (
	"sync"
	"unsafe"
)

type AcseAuthenticationMechanism int32

const (
	/** Neither ACSE nor TLS authentication used */
	ACSE_AUTH_NONE AcseAuthenticationMechanism = C.ACSE_AUTH_NONE

	/** Use ACSE password for client authentication */
	ACSE_AUTH_PASSWORD AcseAuthenticationMechanism = C.ACSE_AUTH_PASSWORD

	/** Use ACSE certificate for client authentication */
	ACSE_AUTH_CERTIFICATE AcseAuthenticationMechanism = C.ACSE_AUTH_CERTIFICATE

	/** Use TLS certificate for client authentication */
	ACSE_AUTH_TLS AcseAuthenticationMechanism = C.ACSE_AUTH_TLS
)

type AcseAuthenticationParameter struct {
	ctx C.AcseAuthenticationParameter
}

func AcseAuthenticationParameter_create() *AcseAuthenticationParameter {
	return &AcseAuthenticationParameter{ctx: C.AcseAuthenticationParameter_create()}
}

func (x *AcseAuthenticationParameter) Destroy() {
	C.AcseAuthenticationParameter_destroy(x.ctx)
}

func (x *AcseAuthenticationParameter) SetAuthMechanism(mechanism AcseAuthenticationMechanism) {
	C.AcseAuthenticationParameter_setAuthMechanism(x.ctx, C.AcseAuthenticationMechanism(mechanism))
}

func (x *AcseAuthenticationParameter) GetAuthMechanism() AcseAuthenticationMechanism {
	return AcseAuthenticationMechanism(C.AcseAuthenticationParameter_getAuthMechanism(x.ctx))
}

func (x *AcseAuthenticationParameter) SetPassword(password string) {
	cpassword := C.CString(password)
	defer C.free(unsafe.Pointer(cpassword))
	C.AcseAuthenticationParameter_setPassword(x.ctx, cpassword)
}

func (x *AcseAuthenticationParameter) GetPassword() string {
	return C.GoString(C.AcseAuthenticationParameter_getPassword(x.ctx))
}

func (x *AcseAuthenticationParameter) GetPasswordLength() int {
	return int(C.AcseAuthenticationParameter_getPasswordLength(x.ctx))
}

func (x *IsoApplicationReference) GetAeQualifier() int {
	return int(C.IsoApplicationReference_getAeQualifier(*x.ctx))
}

func (x *IsoApplicationReference) GetApTitle() *ItuObjectIdentifier {
	return &ItuObjectIdentifier{ctx: C.IsoApplicationReference_getApTitle(x.ctx)}
}

func (x *ItuObjectIdentifier) GetArcCount() int {
	return int(C.ItuObjectIdentifier_getArcCount(x.ctx))
}

func (x *ItuObjectIdentifier) GetArc() []uint16 {
	count := x.GetArcCount()
	arc := C.ItuObjectIdentifier_getArc(x.ctx)
	al := unsafe.Slice(arc, count)
	out := make([]uint16, count)
	for i := range count {
		out[i] = uint16(al[i])
	}
	return out
}

type AcseAuthenticator func(parameter unsafe.Pointer, authParameter AcseAuthenticationParameter, securityToken *unsafe.Pointer, appReference *IsoApplicationReference) bool

var mapAcseAuthenticators = sync.Map{}

//export fAcseAuthenticatorGo
func fAcseAuthenticatorGo(parameter unsafe.Pointer, authParameter C.AcseAuthenticationParameter, securityToken *unsafe.Pointer, appReference *C.IsoApplicationReference) C._Bool {
	ret := false
	mapAcseAuthenticators.Range(func(k, v any) bool {
		if fn, ok := v.(AcseAuthenticator); ok {
			ret = fn(parameter, AcseAuthenticationParameter{ctx: authParameter}, securityToken, &IsoApplicationReference{ctx: appReference})
		}
		return true
	})
	return C._Bool(ret)
}

type TSelector []byte
type SSelector []byte
type PSelector []byte

type IsoConnectionParameters struct {
	ctx C.IsoConnectionParameters
}

func IsoConnectionParametersCreate() *IsoConnectionParameters {
	return &IsoConnectionParameters{ctx: C.IsoConnectionParameters_create()}
}

func (x *IsoConnectionParameters) Destroy() {
	C.IsoConnectionParameters_destroy(x.ctx)
}

func (x *IsoConnectionParameters) SetTlsConfiguration(tlsConfig *TLSConfiguration) {
	C.IsoConnectionParameters_setTlsConfiguration(x.ctx, tlsConfig.ctx)
}

func (x *IsoConnectionParameters) SetAcseAuthenticationParameter(acseAuthParameter *AcseAuthenticationParameter) {
	C.IsoConnectionParameters_setAcseAuthenticationParameter(x.ctx, acseAuthParameter.ctx)
}

func (x *IsoConnectionParameters) SetTcpParameters(hostname string, tcpPort int) {
	chostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(chostname))
	C.IsoConnectionParameters_setTcpParameters(x.ctx, chostname, C.int(tcpPort))
}

func (x *IsoConnectionParameters) SetLocalTcpParameters(localIpAddress string, localTcpPort int) {
	clocalIpAddress := C.CString(localIpAddress)
	defer C.free(unsafe.Pointer(clocalIpAddress))
	C.IsoConnectionParameters_setLocalTcpParameters(x.ctx, clocalIpAddress, C.int(localTcpPort))
}

func (x *IsoConnectionParameters) SetRemoteApTitle(apTitle string, aeQualifier int) {
	capTitle := C.CString(apTitle)
	defer C.free(unsafe.Pointer(capTitle))
	C.IsoConnectionParameters_setRemoteApTitle(x.ctx, capTitle, C.int(aeQualifier))
}

func (x *IsoConnectionParameters) SetRemoteAddresses(pSelector PSelector, sSelector SSelector, tSelector TSelector) {
	in1 := C.PSelector{
		size: C.uint8_t(len(pSelector)),
	}
	for i := 0; i < len(pSelector) && i < 16; i++ {
		in1.value[i] = C.uint8_t(pSelector[i])
	}
	in2 := C.SSelector{
		size: C.uint8_t(len(sSelector)),
	}
	for i := 0; i < len(sSelector) && i < 16; i++ {
		in2.value[i] = C.uint8_t(sSelector[i])
	}
	in3 := C.TSelector{
		size: C.uint8_t(len(tSelector)),
	}
	for i := 0; i < len(tSelector) && i < 4; i++ {
		in3.value[i] = C.uint8_t(tSelector[i])
	}
	C.IsoConnectionParameters_setRemoteAddresses(x.ctx, in1, in2, in3)
}

func (x *IsoConnectionParameters) SetLocalApTitle(apTitle string, aeQualifier int) {
	capTitle := C.CString(apTitle)
	defer C.free(unsafe.Pointer(capTitle))
	C.IsoConnectionParameters_setLocalApTitle(x.ctx, capTitle, C.int(aeQualifier))
}

func (x *IsoConnectionParameters) SetLocalAddresses(pSelector PSelector, sSelector SSelector, tSelector TSelector) {
	in1 := C.PSelector{
		size: C.uint8_t(len(pSelector)),
	}
	for i := 0; i < len(pSelector) && i < 16; i++ {
		in1.value[i] = C.uint8_t(pSelector[i])
	}
	in2 := C.SSelector{
		size: C.uint8_t(len(sSelector)),
	}
	for i := 0; i < len(sSelector) && i < 16; i++ {
		in2.value[i] = C.uint8_t(sSelector[i])
	}
	in3 := C.TSelector{
		size: C.uint8_t(len(tSelector)),
	}
	for i := 0; i < len(tSelector) && i < 4; i++ {
		in3.value[i] = C.uint8_t(tSelector[i])
	}
	C.IsoConnectionParameters_setLocalAddresses(x.ctx, in1, in2, in3)
}
