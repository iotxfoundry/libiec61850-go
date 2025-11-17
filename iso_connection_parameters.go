package libiec61850go

/*
#include <stdlib.h>
#include "libiec61850_common_api.h"
#include "iso_connection_parameters.h"

*/
import "C"

type IsoConnectionParameters struct {
	ctx C.IsoConnectionParameters
}
