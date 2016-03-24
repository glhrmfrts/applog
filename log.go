// Copyright 2016 Guilherme Nemeth.

package applog

import "fmt"

// D is equivalent to Log.d in Java, it logs with the debug flag.
func D(tag, str string) {
	printDebug(tag, str)
}

// DPrintf logs a formatted string with the debug flag.
func DPrintf(tag string, format string, values ...interface{}) {
	printDebug(tag, fmt.Sprintf(format, values...))
}

// TODO: implement logging for Info, Fatal, etc...
