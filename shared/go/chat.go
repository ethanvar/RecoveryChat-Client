// Package main exposes the chat helpers used by NativeChatModule.cpp.

package main

/*
#include <stdlib.h>
*/
import "C"

import "unsafe"

func ChatReverse(input *C.char) *C.char {
	if input == nil {
		return nil
	}

	runes := []rune(C.GoString(input))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return C.CString(string(runes))
}

// ChatFree releases a string previously returned by ChatReverse.
func ChatFree(ptr *C.char) {
	if ptr != nil {
		C.free(unsafe.Pointer(ptr))
	}
}

// main is required by -buildmode=c-archive. It is never called.
func main() {}
