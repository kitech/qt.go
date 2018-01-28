package qtcore

// /usr/include/qt/QtCore/qcryptographichash.h
// #include <qcryptographichash.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QCryptographicHash struct {
	*qtrt.CObject
}

func (this *QCryptographicHash) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCryptographicHash) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQCryptographicHashFromPointer(cthis unsafe.Pointer) *QCryptographicHash {
	return &QCryptographicHash{&qtrt.CObject{cthis}}
}
func (*QCryptographicHash) NewFromPointer(cthis unsafe.Pointer) *QCryptographicHash {
	return NewQCryptographicHashFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcryptographichash.h:92
// index:0
// Public
// void QCryptographicHash(enum QCryptographicHash::Algorithm)
func NewQCryptographicHash(method int) *QCryptographicHash {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHashC2ENS_9AlgorithmE", ffiqt.FFI_TYPE_VOID, cthis, method)
	gopp.ErrPrint(err, rv)
	gothis := NewQCryptographicHashFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcryptographichash.h:93
// index:0
// Public
// void ~QCryptographicHash()
func DeleteQCryptographicHash(*QCryptographicHash) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHashD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:95
// index:0
// Public
// void reset()
func (this *QCryptographicHash) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:97
// index:0
// Public
// void addData(const char *, int)
func (this *QCryptographicHash) AddData(data string, length int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:98
// index:1
// Public
// void addData(const QByteArray &)
func (this *QCryptographicHash) AddData_1(data *QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:99
// index:2
// Public
// bool addData(QIODevice *)
func (this *QCryptographicHash) AddData_2(device *QIODevice /*777 QIODevice **/) bool {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcryptographichash.h:101
// index:0
// Public
// QByteArray result()
func (this *QCryptographicHash) Result() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCryptographicHash6resultEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qcryptographichash.h:103
// index:0
// Public static
// QByteArray hash(const QByteArray &, enum QCryptographicHash::Algorithm)
func (this *QCryptographicHash) Hash(data *QByteArray, method int) *QByteArray /*123*/ {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash4hashERK10QByteArrayNS_9AlgorithmE", ffiqt.FFI_TYPE_POINTER, convArg0, method)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QCryptographicHash_Hash(data *QByteArray, method int) *QByteArray /*123*/ {
	var nilthis *QCryptographicHash
	rv := nilthis.Hash(data, method)
	return rv
}

type QCryptographicHash__Algorithm = int

const QCryptographicHash__Md4 QCryptographicHash__Algorithm = 0
const QCryptographicHash__Md5 QCryptographicHash__Algorithm = 1
const QCryptographicHash__Sha1 QCryptographicHash__Algorithm = 2
const QCryptographicHash__Sha224 QCryptographicHash__Algorithm = 3
const QCryptographicHash__Sha256 QCryptographicHash__Algorithm = 4
const QCryptographicHash__Sha384 QCryptographicHash__Algorithm = 5
const QCryptographicHash__Sha512 QCryptographicHash__Algorithm = 6
const QCryptographicHash__Keccak_224 QCryptographicHash__Algorithm = 7
const QCryptographicHash__Keccak_256 QCryptographicHash__Algorithm = 8
const QCryptographicHash__Keccak_384 QCryptographicHash__Algorithm = 9
const QCryptographicHash__Keccak_512 QCryptographicHash__Algorithm = 10
const QCryptographicHash__RealSha3_224 QCryptographicHash__Algorithm = 11
const QCryptographicHash__RealSha3_256 QCryptographicHash__Algorithm = 12
const QCryptographicHash__RealSha3_384 QCryptographicHash__Algorithm = 13
const QCryptographicHash__RealSha3_512 QCryptographicHash__Algorithm = 14
const QCryptographicHash__Sha3_224 QCryptographicHash__Algorithm = 11
const QCryptographicHash__Sha3_256 QCryptographicHash__Algorithm = 12
const QCryptographicHash__Sha3_384 QCryptographicHash__Algorithm = 13
const QCryptographicHash__Sha3_512 QCryptographicHash__Algorithm = 14

//  body block end
