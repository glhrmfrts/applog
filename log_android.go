// Copyright 2016 Guilherme Nemeth.

package applog

/*
#cgo LDFLAGS: -landroid -llog
#include <android/log.h>
#include <stdlib.h>

void log_print_debug(char *tag, char *str) {
  __android_log_print(ANDROID_LOG_DEBUG, tag, str);
}

*/
import "C"

import (
  "unsafe"
)

func callFn(fn func(*C.char, *C.char), tag, str string) {
  ctag, cstr := C.CString(tag), C.CString(str)
  defer C.free(unsafe.Pointer(ctag))
  defer C.free(unsafe.Pointer(cstr))

  fn(ctag, cstr)
}

func printDebug(tag, str *C.char) {
	C.log_print_debug(tag, str)
}
