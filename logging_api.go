package libiec61850go

/*
#include <stdlib.h>
#include "logging_api.h"

*/
import "C"

type LogStorage struct {
	ctx C.LogStorage
}
