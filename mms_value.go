package libiec61850go

/*
#include <stdlib.h>
#include "mms_value.h"

*/
import "C"

type MmsValue struct {
	ctx *C.MmsValue
}
