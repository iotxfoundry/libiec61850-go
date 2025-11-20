package main

import (
	"flag"
	"fmt"
	"time"
	"unsafe"

	libiec61850go "github.com/iotxfoundry/libiec61850-go"
)

var (
	portFlag = flag.Int("port", 102, "port number to listen on")
)

var iedServer *libiec61850go.IedServer

func main() {
	flag.Parse()

	fmt.Printf("Using libIEC61850 version %s\n", libiec61850go.GetVersionString())

	// Create new server configuration object
	config := libiec61850go.IedServerConfigCreate()
	// Set buffer size for buffered report control blocks to 200000 bytes
	config.SetReportBufferSize(200000)
	// Set stack compliance to a specific edition of the standard (WARNING: data model has also to be checked for compliance) */
	config.SetEdition(libiec61850go.IEC_61850_EDITION_2)
	// Set the base path for the MMS file services
	config.SetFileServiceBasePath("./vmd-filestore/")
	// Disable MMS file service
	config.EnableFileService(false)
	// Enable dynamic data set service
	config.EnableDynamicDataSetService(true)
	// Disable log service
	config.EnableLogService(false)
	// Set maximum number of MMS connections to 2
	config.SetMaxMmsConnections(2)

	// Create a new IEC 61850 server instance
	iedServer = libiec61850go.IedServerCreateWithConfig(iedModel, nil, config)

	// configuration object is no longer required
	config.Destroy()

	// Set the identity values for MMS identify service
	iedServer.SetServerIdentity("MZ", "basic io", "1.6.0")

	// Install handler for operate command
	iedServer.SetControlHandler(iedModel_GenericIO_GGIO1_SPCSO1, ControlHandler, unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO1))
}

func ControlHandler(action *libiec61850go.ControlAction, parameter unsafe.Pointer, ctlVal *libiec61850go.MmsValue, test bool) libiec61850go.ControlHandlerResult {
	if test {
		return libiec61850go.CONTROL_RESULT_FAILED
	}

	if ctlVal.GetType() == libiec61850go.MMS_BOOLEAN {
		fmt.Printf("received binary control command: ")

		if ctlVal.GetBoolean() {
			fmt.Println("on")
		} else {
			fmt.Println("off")
		}
	} else {
		return libiec61850go.CONTROL_RESULT_FAILED
	}

	timeStamp := time.Now().UnixNano() / int64(time.Millisecond)

	if parameter == unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO1) {
		iedServer.UpdateUTCTimeAttributeValue(iedModel_GenericIO_GGIO1_SPCSO1_t, uint64(timeStamp))
		iedServer.UpdateAttributeValue(iedModel_GenericIO_GGIO1_SPCSO1_stVal, ctlVal)
	}

	if parameter == unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO2) {
		iedServer.UpdateUTCTimeAttributeValue(iedModel_GenericIO_GGIO1_SPCSO2_t, uint64(timeStamp))
		iedServer.UpdateAttributeValue(iedModel_GenericIO_GGIO1_SPCSO2_stVal, ctlVal)
	}

	if parameter == unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO3) {
		iedServer.UpdateUTCTimeAttributeValue(iedModel_GenericIO_GGIO1_SPCSO3_t, uint64(timeStamp))
		iedServer.UpdateAttributeValue(iedModel_GenericIO_GGIO1_SPCSO3_stVal, ctlVal)
	}

	if parameter == unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO4) {
		iedServer.UpdateUTCTimeAttributeValue(iedModel_GenericIO_GGIO1_SPCSO4_t, uint64(timeStamp))
		iedServer.UpdateAttributeValue(iedModel_GenericIO_GGIO1_SPCSO4_stVal, ctlVal)
	}

	return libiec61850go.CONTROL_RESULT_OK
}
