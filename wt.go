package wavelib

/*
#cgo LDFLAGS: -lm -lwavelib
#include <wavelib.h>
*/
import "C"

import (
	"reflect"
	"runtime"
	"unsafe"
)

type WTObject struct {
	C_WTObject C.wt_object
}

func WTInit(wave WaveObject, method string, sigLength int, j int) WTObject {
	return WTObject{C_WTObject: C.wt_init(wave.C_WaveObject, C.CString(method), C.int(sigLength), C.int(j))}
}

func (wt WTObject) Free() {
	WTFree(wt)
}

func (wt WTObject) OutLength() int {
	return int(wt.C_WTObject.outlength)
}

func (wt WTObject) Output() []float64 {
	length := wt.OutLength()
	sliceHdr := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(wt.C_WTObject.output)),
		Cap:  length,
		Len:  length,
	}
	return *(*[]float64)(unsafe.Pointer(sliceHdr))
}

func (wt WTObject) SigLength() int {
	return int(wt.C_WTObject.siglength)
}

func WTFree(wt WTObject) {
	C.wt_free(wt.C_WTObject)
}

func SetWTConv(wt WTObject, cmethod string) {
	C.setWTConv(wt.C_WTObject, C.CString(cmethod))
}

func SWT(wt WTObject, inp []float64) {
	C.swt(wt.C_WTObject, (*C.double)((unsafe.Pointer)((*reflect.SliceHeader)((unsafe.Pointer)(&inp)).Data)))
	runtime.KeepAlive(inp)
}

func ISWT(wt WTObject, swtop []float64) {
	C.iswt(wt.C_WTObject, (*C.double)((unsafe.Pointer)((*reflect.SliceHeader)((unsafe.Pointer)(&swtop)).Data)))
	runtime.KeepAlive(swtop)
}

func WTSummary(wt WTObject) {
	C.wt_summary(wt.C_WTObject)
}
