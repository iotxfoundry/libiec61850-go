package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
	"unsafe"

	libiec61850go "github.com/iotxfoundry/libiec61850-go"
)

// inject by go build
var (
	Version   = "0.0.0"
	BuildTime = "2020-01-13-0802 UTC"
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
	iedServer.SetControlHandler(iedModel_GenericIO_GGIO1_SPCSO1, controlHandlerForBinaryOutput, unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO1))
	iedServer.SetControlHandler(iedModel_GenericIO_GGIO1_SPCSO2, controlHandlerForBinaryOutput, unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO2))
	iedServer.SetControlHandler(iedModel_GenericIO_GGIO1_SPCSO3, controlHandlerForBinaryOutput, unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO3))
	iedServer.SetControlHandler(iedModel_GenericIO_GGIO1_SPCSO4, controlHandlerForBinaryOutput, unsafe.Pointer(iedModel_GenericIO_GGIO1_SPCSO4))

	iedServer.SetConnectionIndicationHandler(connectionHandler, nil)

	iedServer.SetRCBEventHandler(rcbEventHandler, nil)

	iedServer.SetWriteAccessPolicy(libiec61850go.IEC61850_FC_DC, libiec61850go.ACCESS_POLICY_ALLOW)

	iedServer.Start(*portFlag)

	if !iedServer.IsRunning() {
		fmt.Printf("Starting server failed (maybe need root permissions or another server is already using the port)! Exit.\n")
		iedServer.Destroy()
		os.Exit(1)
	}
	t := float32(0.0)
	timer := time.NewTimer(100 * time.Millisecond)
	for range timer.C {
		timestamp := uint64(time.Now().UnixNano() / int64(time.Millisecond))
		t += 0.1

		an1 := float32(math.Sin(float64(t)))
		an2 := float32(math.Sin(float64(t + 1.)))
		an3 := float32(math.Sin(float64(t + 2.)))
		an4 := float32(math.Sin(float64(t + 3.)))

		iecTimestamp := libiec61850go.NewTimestamp()
		iecTimestamp.ClearFlags()
		iecTimestamp.SetTimeInMilliseconds(timestamp)
		iecTimestamp.SetLeapSecondKnown(true)

		/* toggle clock-not-synchronized flag in timestamp */
		if int(t)%2 == 0 {
			iecTimestamp.SetClockNotSynchronized(true)
		} else {
			iecTimestamp.SetClockNotSynchronized(false)
		}

		iedServer.LockDataModel()
		iedServer.UpdateTimestampAttributeValue(iedModel_GenericIO_GGIO1_AnIn1_t, iecTimestamp)
		iedServer.UpdateFloatAttributeValue(iedModel_GenericIO_GGIO1_AnIn1_mag_f, an1)

		iedServer.UpdateTimestampAttributeValue(iedModel_GenericIO_GGIO1_AnIn2_t, iecTimestamp)
		iedServer.UpdateFloatAttributeValue(iedModel_GenericIO_GGIO1_AnIn2_mag_f, an2)

		iedServer.UpdateTimestampAttributeValue(iedModel_GenericIO_GGIO1_AnIn3_t, iecTimestamp)
		iedServer.UpdateFloatAttributeValue(iedModel_GenericIO_GGIO1_AnIn3_mag_f, an3)

		iedServer.UpdateTimestampAttributeValue(iedModel_GenericIO_GGIO1_AnIn4_t, iecTimestamp)
		iedServer.UpdateFloatAttributeValue(iedModel_GenericIO_GGIO1_AnIn4_mag_f, an4)
		iedServer.UnlockDataModel()
	}

	iedServer.Stop()
	iedServer.Destroy()
}

func rcbEventHandler(parameter unsafe.Pointer, rcb *libiec61850go.ReportControlBlock, connection *libiec61850go.ClientConnection, event libiec61850go.RCBEventType, parameterName string, serviceError libiec61850go.MmsDataAccessError) {
	fmt.Printf("RCB: %s event: %d\n", rcb.Name(), event)

	if (event == libiec61850go.RCB_EVENT_SET_PARAMETER) || (event == libiec61850go.RCB_EVENT_GET_PARAMETER) {
		fmt.Printf("  param:  %s\n", parameterName)
		fmt.Printf("  result: %d\n", serviceError)
	}

	if event == libiec61850go.RCB_EVENT_ENABLE {
		rptId := rcb.GetRptId()
		fmt.Printf("   rptID:  %s\n", rptId)
		dataSet := rcb.GetDataSetName()
		fmt.Printf("   datSet: %s\n", dataSet)
	}
}

func connectionHandler(self *libiec61850go.IedServer, connection *libiec61850go.ClientConnection, connected bool, parameter unsafe.Pointer) {
	if connected {
		fmt.Println("Connection opened")
	} else {
		fmt.Println("Connection closed")
	}
}

func controlHandlerForBinaryOutput(action *libiec61850go.ControlAction, parameter unsafe.Pointer, ctlVal *libiec61850go.MmsValue, test bool) libiec61850go.ControlHandlerResult {
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
