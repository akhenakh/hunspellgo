package hunspellgo

// #cgo LDFLAGS: -lhunspell
// #include <stdlib.h>
// #include <hunspell/hunspell.h>
import "C"
import "unsafe"

func Hunspell(affpath string, dpath string) (*Hunhandle, error) {
	affpathcs := C.CString(affpath)
	defer C.free(unsafe.Pointer(affpathcs))
	dpathcs := C.CString(dpath)
	defer C.free(unsafe.Pointer(dpathcs))

	var handle = C.Hunspell_create(affpathcs, dpathcs)
	return &Hunhandle{handle}, nil
}

func Suggest(handle *Hunhandle, word string) {
	wordcs := C.CString(word)
	defer C.free(unsafe.Pointer(wordcs))
  var carray **C.char
  length := C.Hunspell_suggest(handle.handle, &carray, wordcs)
  C.Hunspell_free_list(handle.handle, &carray, length)
}

func Spell(handle *Hunhandle, word string) bool {
  wordcs := C.CString(word)
  defer C.free(unsafe.Pointer(wordcs))
  res := C.Hunspell_spell(handle.handle, wordcs)
  if int(res) == 0 {
    return false
  }
  return true
}

func Destroy(handle *Hunhandle) {
  C.Hunspell_destroy(handle.handle)
  handle.handle = nil;
}

type Hunhandle struct {
	handle *C.Hunhandle
}
