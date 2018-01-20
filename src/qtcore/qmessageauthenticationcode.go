//  header block begin
// /usr/include/qt/QtCore/qmessageauthenticationcode.h
// #include <qmessageauthenticationcode.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QMessageAuthenticationCode struct {
	*qtrt.CObject
}

func (this *QMessageAuthenticationCode) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:54
// index:0
// void QMessageAuthenticationCode(class QCryptographicHash::Algorithm, const class QByteArray &)
func NewQMessageAuthenticationCode(method int, key unsafe.Pointer) *QMessageAuthenticationCode {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeC2EN18QCryptographicHash9AlgorithmERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, &method, key)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageAuthenticationCodeFromPointer(cthis)
	return gothis
}
func NewQMessageAuthenticationCodeFromPointer(cthis unsafe.Pointer) *QMessageAuthenticationCode {
	return &QMessageAuthenticationCode{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:56
// index:0
// void ~QMessageAuthenticationCode()
func DeleteQMessageAuthenticationCode(*QMessageAuthenticationCode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:58
// index:0
// void reset()
func (this *QMessageAuthenticationCode) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:60
// index:0
// void setKey(const class QByteArray &)
func (this *QMessageAuthenticationCode) SetKey(key unsafe.Pointer) {
	// 0: (, key const QByteArray &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode6setKeyERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:62
// index:0
// void addData(const char *, int)
func (this *QMessageAuthenticationCode) AddData(data unsafe.Pointer, length int) {
	// 0: (, data const char *, length int), (data, &length)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataEPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:63
// index:1
// void addData(const class QByteArray &)
func (this *QMessageAuthenticationCode) AddData_1(data unsafe.Pointer) {
	// 1: (, data const QByteArray &), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:64
// index:2
// bool addData(class QIODevice *)
func (this *QMessageAuthenticationCode) AddData_2(device unsafe.Pointer) {
	// 2: (, device QIODevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataEP9QIODevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:66
// index:0
// QByteArray result()
func (this *QMessageAuthenticationCode) Result() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QMessageAuthenticationCode6resultEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:68
// index:0
// static
// QByteArray hash(const class QByteArray &, const class QByteArray &, class QCryptographicHash::Algorithm)
func (this *QMessageAuthenticationCode) Hash(message unsafe.Pointer, key unsafe.Pointer, method int) {
	// 0: (message const QByteArray &, key const QByteArray &, method QCryptographicHash::Algorithm), (message, key, method)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode4hashERK10QByteArrayS2_N18QCryptographicHash9AlgorithmE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QMessageAuthenticationCode_Hash(message unsafe.Pointer, key unsafe.Pointer, method int) {
	// 0: (message const QByteArray &, key const QByteArray &, method QCryptographicHash::Algorithm), (message, key, method)
	var nilthis *QMessageAuthenticationCode
	nilthis.Hash(message, key, method)
}

//  body block end
