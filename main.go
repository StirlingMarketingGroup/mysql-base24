package main

// #include <string.h>
// #include <stdbool.h>
// #include <mysql.h>
// #cgo CFLAGS: -O3 -I/usr/include/mysql -fno-omit-frame-pointer
import "C"
import (
	"unsafe"

	"github.com/eknkc/basex"
)

func msg(message *C.char, s string) {
	m := C.CString(s)
	defer C.free(unsafe.Pointer(m))

	C.strcpy(message, m)
}

var base24, _ = basex.NewEncoding("2346789bcdfghjkmpqrtvwxy")

//export to_base24_init
func to_base24_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 1 {
		msg(message, "`to_base24` requires 1 parameter: the string to be encoded")
		return C.bool(true)
	}

	argsTypes := (*[2]uint32)(unsafe.Pointer(args.arg_type))

	argsTypes[0] = C.STRING_RESULT
	initid.maybe_null = 1

	return C.bool(false)
}

//export to_base24
func to_base24(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, message *C.char) *C.char {
	*isNull = 1

	c := 1
	argsArgs := (*[1 << 30]*C.char)(unsafe.Pointer(args.args))[:c:c]
	argsLengths := (*[1 << 30]uint64)(unsafe.Pointer(args.lengths))[:c:c]

	if argsArgs[0] == nil {
		*length = 0
		*isNull = 1
		return nil
	}

	a := make([]string, c, c)
	for i, argsArg := range argsArgs {
		a[i] = C.GoStringN(argsArg, C.int(argsLengths[i]))
	}

	encoded := base24.Encode([]byte(a[0]))
	*length = uint64(len(encoded))
	*isNull = 0
	return C.CString(encoded)
}

//export from_base24_init
func from_base24_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 1 {
		msg(message, "`to_base24` requires 1 parameter: the string to be decoded")
		return C.bool(true)
	}

	argsTypes := (*[2]uint32)(unsafe.Pointer(args.arg_type))

	argsTypes[0] = C.STRING_RESULT
	initid.maybe_null = 1

	return C.bool(false)
}

//export from_base24
func from_base24(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, message *C.char) *C.char {
	*isNull = 1

	c := 1
	argsArgs := (*[1 << 30]*C.char)(unsafe.Pointer(args.args))[:c:c]
	argsLengths := (*[1 << 30]uint64)(unsafe.Pointer(args.lengths))[:c:c]

	if argsArgs[0] == nil {
		*length = 0
		*isNull = 1
		return nil
	}

	a := make([]string, c, c)
	for i, argsArg := range argsArgs {
		a[i] = C.GoStringN(argsArg, C.int(argsLengths[i]))
	}

	decoded, err := base24.Decode(a[0])
	if err != nil {
		*length = 0
		*isNull = 1
		return nil
	}

	*length = uint64(len(decoded))
	*isNull = 0
	return C.CString(string(decoded))
}

func main() {}
