package xxtea

//#include <string.h>
//#include <stdlib.h>
//#include <strings.h>
//#include "xxtea.h"
import "C"

import (
	"errors"
	"unsafe"
)

func xxteaRun(src []byte, key []byte, isEncode bool) ([]byte, error) {
	l := len(src)
	if l == 0 {
		return nil, errors.New("src data error")
	}
	srcptr := (*C.uchar)(unsafe.Pointer(&src[0]))
	srclen := C.xxtea_long(len(src))
	keyptr := (*C.uchar)(unsafe.Pointer(&key[0]))
	keylen := C.xxtea_long(len(key))
	retlen := C.xxtea_long(0)
	retptr := (*C.uchar)(nil)
	if isEncode {
		retptr = C.xxtea_encrypt(srcptr, srclen, keyptr, keylen, &retlen)
	} else {
		retptr = C.xxtea_decrypt(srcptr, srclen, keyptr, keylen, &retlen)
	}
	if retlen <= 0 || retptr == nil {
		return nil, errors.New("xxtea error")
	}
	d := make([]byte, retlen)
	C.memcpy(unsafe.Pointer(&d[0]), unsafe.Pointer(retptr), C.size_t(retlen))
	C.free(unsafe.Pointer(retptr))
	return d, nil
}

func Encode(src []byte, key []byte) ([]byte, error) {
	return xxteaRun(src, key, true)
}

func Decode(src []byte, key []byte) ([]byte, error) {
	return xxteaRun(src, key, false)
}
