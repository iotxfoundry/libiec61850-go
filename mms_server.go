package libiec61850go

/*
#include <stdlib.h>
#include "mms_server.h"

*/
import "C"

type MmsServer struct {
	ctx C.MmsServer
}
