//  header block begin
// /usr/include/qt/QtCore/qcryptographichash.h
// #include <qcryptographichash.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
	return this.Cthis
}

// /usr/include/qt/QtCore/qcryptographichash.h:92
// index:0
// void QCryptographicHash(enum QCryptographicHash::Algorithm)
func NewQCryptographicHash(method int) *QCryptographicHash {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHashC2ENS_9AlgorithmE", ffiqt.FFI_TYPE_VOID, cthis, &method)
	gopp.ErrPrint(err, rv)
	gothis := NewQCryptographicHashFromPointer(cthis)
	return gothis
}
func NewQCryptographicHashFromPointer(cthis unsafe.Pointer) *QCryptographicHash {
	return &QCryptographicHash{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qcryptographichash.h:93
// index:0
// void ~QCryptographicHash()
func DeleteQCryptographicHash(*QCryptographicHash) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHashD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:95
// index:0
// void reset()
func (this *QCryptographicHash) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:97
// index:0
// void addData(const char *, int)
func (this *QCryptographicHash) AddData(data unsafe.Pointer, length int) {
	// 0: (, data const char *, length int), (data, &length)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataEPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:98
// index:1
// void addData(const class QByteArray &)
func (this *QCryptographicHash) AddData_1(data unsafe.Pointer) {
	// 1: (, data const QByteArray &), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:99
// index:2
// bool addData(class QIODevice *)
func (this *QCryptographicHash) AddData_2(device unsafe.Pointer) {
	// 2: (, device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash7addDataEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:101
// index:0
// QByteArray result()
func (this *QCryptographicHash) Result() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCryptographicHash6resultEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcryptographichash.h:103
// index:0
// static
// QByteArray hash(const class QByteArray &, enum QCryptographicHash::Algorithm)
func (this *QCryptographicHash) Hash(data unsafe.Pointer, method int) {
	// 0: (data const QByteArray &, method QCryptographicHash::Algorithm), (data, method)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCryptographicHash4hashERK10QByteArrayNS_9AlgorithmE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCryptographicHash_Hash(data unsafe.Pointer, method int) {
	// 0: (data const QByteArray &, method QCryptographicHash::Algorithm), (data, method)
	var nilthis *QCryptographicHash
	nilthis.Hash(data, method)
}

//  body block end
