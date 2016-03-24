// Copyright 2016 Guilherme Nemeth.

package applog

/*
#cgo LDFLAGS: -landroid -llog
#include <android/log.h>

void log_print_debug(char *tag, char *str) {
  __android_log_print(ANDROID_LOG_DEBUG, tag, str);
}

*/
import "C"

func printDebug(tag, str string) {
	C.log_print_debug(C.CString(tag), C.CString(str))
}
