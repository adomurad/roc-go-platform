package roc

// #include "app.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func Main() int {

	size := C.roc__mainForHost_1_exposed_size()
	capturePtr := roc_alloc(size, 0)
	defer roc_dealloc(capturePtr, 0)

	// C.roc__mainForHost_1_exposed_generic(capturePtr)

	result := C.roc__mainForHost_1_exposed()

	// var result C.struct_ResultVoidI64
	// C.roc__mainForHost_0_caller(nil, capturePtr, &result)

	return (*(*int)(unsafe.Pointer(&result)))

	// switch result.disciminant {
	// case 1: // Ok
	// 	return 0
	// case 0: // Err
	// 	return (*(*int)(unsafe.Pointer(&result.payload)))
	// default:
	// 	panic("invalid disciminat")
	// }
}

//export roc_fx_stdoutLine
func roc_fx_stdoutLine(msg *RocStr) {
	fmt.Println(msg)
}

type RocResultType int

const (
	RocErr = iota
	RocOk
)

func createRocResultStr(resultType RocResultType, str string) C.struct_ResultVoidStr {
	rocStr := NewRocStr(str)

	var result C.struct_ResultVoidStr

	result.disciminant = C.uchar(resultType)

	payloadPtr := unsafe.Pointer(&result.payload)
	*(*C.struct_RocStr)(payloadPtr) = rocStr.C()

	return result
}
