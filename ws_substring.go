/*
 MySQL UDF ws substring
 */
package main

import (
	/*
		#cgo CFLAGS: -I/usr/include/mysql
		#include <stdio.h>
		#include <stdlib.h>
		#include <string.h>
		#include <mysql.h>
		#include <limits.h>
	*/
	"C"
	"github.com/shigenobu/mysql_ws_substring/func"
	"strconv"
	"unsafe"
)

// convert argc, argv into go structure
func argToGostrings(count C.uint, args **C.char, lengths *C.ulong) []string {
	// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	length := count
	argslice := (*[1 << 30]*C.char)(unsafe.Pointer(args))[:length:length]
	lengthsslice := (*[1 << 3]C.ulong)(unsafe.Pointer(lengths))[:length:length]

	gostrings := make([]string, count)

	for i, s := range argslice {
		l := C.int(lengthsslice[i])
		gostrings[i] = C.GoStringN(s, l)
	}
	return gostrings
}

// get pointer value
func getUintPointerValue(pointer *uint32, offset int) *C.uint {
	// https://www.programminghunter.com/article/4762726307/
	return (*C.uint)(unsafe.Pointer(uintptr(unsafe.Pointer(pointer)) + uintptr(offset)*uintptr(unsafe.Sizeof(C.uint(0)))))
}

//export ws_substring_init
func ws_substring_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	length := int(args.arg_count)
	if length != 3 {
		return 1
	}

	typeslice := (*[3]C.uint)(unsafe.Pointer(args.arg_type))[:length:length]
	if typeslice[0] != C.STRING_RESULT {
		return 1
	}
	if typeslice[1] != C.INT_RESULT {
		return 1
	}
	if typeslice[2] != C.INT_RESULT {
		return 1
	}

	// resolve int value as string
	*getUintPointerValue(args.arg_type, 1) = C.STRING_RESULT
	*getUintPointerValue(args.arg_type, 2) = C.STRING_RESULT

	return 0
}

//export ws_substring
func ws_substring(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *C.ulong, isNull *C.char, error *C.char) *C.char {
	C.free(unsafe.Pointer(initid.ptr))

	argsString := argToGostrings(args.arg_count, args.args, args.lengths)

	// -----
	// ここが処理本体
	start, _ := strconv.Atoi(argsString[1])
	leng, _ := strconv.Atoi(argsString[2])
	rslt := _func.Substring(argsString[0], start, leng)
	// -----

	initid.ptr = C.CString(rslt)
	*length = C.ulong(len(rslt))
	return initid.ptr
}

//export ws_substring_deinit
func ws_substring_deinit(initid *C.UDF_INIT) {
	C.free(unsafe.Pointer(initid.ptr))
}

func main() {
}
