package wavelib

/*
#cgo LDFLAGS: -lm -lwavelib
#include <wavelib.h>
*/
import "C"

type WaveObject struct {
	C_WaveObject C.wave_object
}

func WaveInit(name string) WaveObject {
	return WaveObject{C_WaveObject: C.wave_init(C.CString(name))}
}

func (obj WaveObject) Free() {
	WaveFree(obj)
}

func WaveFree(obj WaveObject) {
	C.wave_free(obj.C_WaveObject)
}
