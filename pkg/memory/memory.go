package memory

import "unsafe"

func ReadLoc(addr uintptr) {
	//data := Extract(addr, )
}

func Extract(ptr unsafe.Pointer, size uintptr) []byte {
	out := make([]byte, size)
	for i := range out {
		out[i] = *((*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))))
	}
	return out
}
