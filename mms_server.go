package libiec61850go

/*
#include <stdlib.h>
#include "mms_server.h"

*/
import "C"

type MmsServerEvent int32

const (
	MMS_SERVER_NEW_CONNECTION    MmsServerEvent = C.MMS_SERVER_NEW_CONNECTION
	MMS_SERVER_CONNECTION_CLOSED MmsServerEvent = C.MMS_SERVER_CONNECTION_CLOSED
	MMS_SERVER_CONNECTION_TICK   MmsServerEvent = C.MMS_SERVER_CONNECTION_TICK
)

type MmsServer struct {
	ctx C.MmsServer
}

type MmsServerConnection struct {
	ctx C.MmsServerConnection
}

type MmsVariableListType int32

const (
	MMS_DOMAIN_SPECIFIC      MmsVariableListType = C.MMS_DOMAIN_SPECIFIC
	MMS_ASSOCIATION_SPECIFIC MmsVariableListType = C.MMS_ASSOCIATION_SPECIFIC
	MMS_VMD_SPECIFIC         MmsVariableListType = C.MMS_VMD_SPECIFIC
)

type MmsVariableListAccessType int32

const (
	MMS_VARLIST_CREATE        MmsVariableListAccessType = C.MMS_VARLIST_CREATE
	MMS_VARLIST_DELETE        MmsVariableListAccessType = C.MMS_VARLIST_DELETE
	MMS_VARLIST_READ          MmsVariableListAccessType = C.MMS_VARLIST_READ
	MMS_VARLIST_WRITE         MmsVariableListAccessType = C.MMS_VARLIST_WRITE
	MMS_VARLIST_GET_DIRECTORY MmsVariableListAccessType = C.MMS_VARLIST_GET_DIRECTORY
)

type MmsFileServiceType int32

const (
	MMS_FILE_ACCESS_TYPE_READ_DIRECTORY MmsFileServiceType = C.MMS_FILE_ACCESS_TYPE_READ_DIRECTORY
	MMS_FILE_ACCESS_TYPE_OPEN           MmsFileServiceType = C.MMS_FILE_ACCESS_TYPE_OPEN
	MMS_FILE_ACCESS_TYPE_OBTAIN         MmsFileServiceType = C.MMS_FILE_ACCESS_TYPE_OBTAIN
	MMS_FILE_ACCESS_TYPE_DELETE         MmsFileServiceType = C.MMS_FILE_ACCESS_TYPE_DELETE
	MMS_FILE_ACCESS_TYPE_RENAME         MmsFileServiceType = C.MMS_FILE_ACCESS_TYPE_RENAME
)
