package qtcore

// /usr/include/qt/QtCore/qcryptographichash.h
// #include <qcryptographichash.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QCryptographicHash struct {
	*qtrt.CObject
}
type QCryptographicHash_ITF interface {
	QCryptographicHash_PTR() *QCryptographicHash
}

func (ptr *QCryptographicHash) QCryptographicHash_PTR() *QCryptographicHash { return ptr }

func (this *QCryptographicHash) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCryptographicHash) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCryptographicHashFromPointer(cthis unsafe.Pointer) *QCryptographicHash {
	return &QCryptographicHash{&qtrt.CObject{cthis}}
}
func (*QCryptographicHash) NewFromPointer(cthis unsafe.Pointer) *QCryptographicHash {
	return NewQCryptographicHashFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcryptographichash.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCryptographicHash(QCryptographicHash::Algorithm)

/*
Constructs an object that can be used to create a cryptographic hash from data using method.
*/
func (*QCryptographicHash) NewForInherit(method int) *QCryptographicHash {
	return NewQCryptographicHash(method)
}
func NewQCryptographicHash(method int) *QCryptographicHash {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCryptographicHashC2ENS_9AlgorithmE", qtrt.FFI_TYPE_POINTER, method)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCryptographicHashFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCryptographicHash)
	return gothis
}

// /usr/include/qt/QtCore/qcryptographichash.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCryptographicHash()

/*

 */
func DeleteQCryptographicHash(this *QCryptographicHash) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCryptographicHashD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcryptographichash.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*
Resets the object.
*/
func (this *QCryptographicHash) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCryptographicHash5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addData(const char *, int)

/*
Adds the first length chars of data to the cryptographic hash.
*/
func (this *QCryptographicHash) AddData(data string, length int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, length)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:98
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addData(const QByteArray &)

/*
Adds the first length chars of data to the cryptographic hash.
*/
func (this *QCryptographicHash) AddData_1(data QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:99
// index:2
// Public Visibility=Default Availability=Available
// [1] bool addData(QIODevice *)

/*
Adds the first length chars of data to the cryptographic hash.
*/
func (this *QCryptographicHash) AddData_2(device QIODevice_ITF /*777 QIODevice **/) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcryptographichash.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray result() const

/*
Returns the final hash value.

See also QByteArray::toHex().
*/
func (this *QCryptographicHash) Result() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCryptographicHash6resultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qcryptographichash.h:103
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray hash(const QByteArray &, QCryptographicHash::Algorithm)

/*
Returns the hash of data using method.
*/
func (this *QCryptographicHash) Hash(data QByteArray_ITF, method int) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCryptographicHash4hashERK10QByteArrayNS_9AlgorithmE", qtrt.FFI_TYPE_POINTER, convArg0, method)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QCryptographicHash_Hash(data QByteArray_ITF, method int) *QByteArray /*123*/ {
	var nilthis *QCryptographicHash
	rv := nilthis.Hash(data, method)
	return rv
}

/*
Note: In Qt versions before 5.9, when asked to generate a SHA3 hash sum, QCryptographicHash actually calculated Keccak. If you need compatibility with SHA-3 hashes produced by those versions of Qt, use the Keccak_ enumerators. Alternatively, if source compatibility is required, define the macro QT_SHA3_KECCAK_COMPAT.


*/
type QCryptographicHash__Algorithm = int

//
const QCryptographicHash__Md4 QCryptographicHash__Algorithm = 0

//
const QCryptographicHash__Md5 QCryptographicHash__Algorithm = 1

//
const QCryptographicHash__Sha1 QCryptographicHash__Algorithm = 2

//
const QCryptographicHash__Sha224 QCryptographicHash__Algorithm = 3

//
const QCryptographicHash__Sha256 QCryptographicHash__Algorithm = 4

//
const QCryptographicHash__Sha384 QCryptographicHash__Algorithm = 5

//
const QCryptographicHash__Sha512 QCryptographicHash__Algorithm = 6

//
const QCryptographicHash__Keccak_224 QCryptographicHash__Algorithm = 7

//
const QCryptographicHash__Keccak_256 QCryptographicHash__Algorithm = 8

//
const QCryptographicHash__Keccak_384 QCryptographicHash__Algorithm = 9

//
const QCryptographicHash__Keccak_512 QCryptographicHash__Algorithm = 10

//
const QCryptographicHash__RealSha3_224 QCryptographicHash__Algorithm = 11

//
const QCryptographicHash__RealSha3_256 QCryptographicHash__Algorithm = 12

//
const QCryptographicHash__RealSha3_384 QCryptographicHash__Algorithm = 13

//
const QCryptographicHash__RealSha3_512 QCryptographicHash__Algorithm = 14

//
const QCryptographicHash__Sha3_224 QCryptographicHash__Algorithm = 11

//
const QCryptographicHash__Sha3_256 QCryptographicHash__Algorithm = 12

//
const QCryptographicHash__Sha3_384 QCryptographicHash__Algorithm = 13

//
const QCryptographicHash__Sha3_512 QCryptographicHash__Algorithm = 14

func (this *QCryptographicHash) AlgorithmItemName(val int) string {
	switch val {
	case QCryptographicHash__Md4: // 0
		return "Md4"
	case QCryptographicHash__Md5: // 1
		return "Md5"
	case QCryptographicHash__Sha1: // 2
		return "Sha1"
	case QCryptographicHash__Sha224: // 3
		return "Sha224"
	case QCryptographicHash__Sha256: // 4
		return "Sha256"
	case QCryptographicHash__Sha384: // 5
		return "Sha384"
	case QCryptographicHash__Sha512: // 6
		return "Sha512"
	case QCryptographicHash__Keccak_224: // 7
		return "Keccak_224"
	case QCryptographicHash__Keccak_256: // 8
		return "Keccak_256"
	case QCryptographicHash__Keccak_384: // 9
		return "Keccak_384"
	case QCryptographicHash__Keccak_512: // 10
		return "Keccak_512"
	case QCryptographicHash__RealSha3_224: // 11
		return "RealSha3_224,Sha3_224"
	case QCryptographicHash__RealSha3_256: // 12
		return "RealSha3_256,Sha3_256"
	case QCryptographicHash__RealSha3_384: // 13
		return "RealSha3_384,Sha3_384"
	case QCryptographicHash__RealSha3_512: // 14
		return "RealSha3_512,Sha3_512"
		// case QCryptographicHash__Sha3_224: // 11
		// return ""
		// case QCryptographicHash__Sha3_256: // 12
		// return ""
		// case QCryptographicHash__Sha3_384: // 13
		// return ""
		// case QCryptographicHash__Sha3_512: // 14
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCryptographicHash_AlgorithmItemName(val int) string {
	var nilthis *QCryptographicHash
	return nilthis.AlgorithmItemName(val)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
