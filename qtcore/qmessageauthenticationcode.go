package qtcore

// /usr/include/qt/QtCore/qmessageauthenticationcode.h
// #include <qmessageauthenticationcode.h>
// #include <QtCore>

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
import "mkuse/cffiqt"
import "gopp"
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
type QMessageAuthenticationCode struct {
	*qtrt.CObject
}

func (this *QMessageAuthenticationCode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMessageAuthenticationCode) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMessageAuthenticationCodeFromPointer(cthis unsafe.Pointer) *QMessageAuthenticationCode {
	return &QMessageAuthenticationCode{&qtrt.CObject{cthis}}
}
func (*QMessageAuthenticationCode) NewFromPointer(cthis unsafe.Pointer) *QMessageAuthenticationCode {
	return NewQMessageAuthenticationCodeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:54
// index:0
// Public
// void QMessageAuthenticationCode(class QCryptographicHash::Algorithm, const class QByteArray &)
func NewQMessageAuthenticationCode(method int, key *QByteArray) *QMessageAuthenticationCode {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg1 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeC2EN18QCryptographicHash9AlgorithmERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, method, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageAuthenticationCodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:56
// index:0
// Public
// void ~QMessageAuthenticationCode()
func DeleteQMessageAuthenticationCode(*QMessageAuthenticationCode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:58
// index:0
// Public
// void reset()
func (this *QMessageAuthenticationCode) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:60
// index:0
// Public
// void setKey(const class QByteArray &)
func (this *QMessageAuthenticationCode) SetKey(key *QByteArray) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode6setKeyERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:62
// index:0
// Public
// void addData(const char *, int)
func (this *QMessageAuthenticationCode) AddData(data string, length int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:63
// index:1
// Public
// void addData(const class QByteArray &)
func (this *QMessageAuthenticationCode) AddData_1(data *QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:64
// index:2
// Public
// bool addData(class QIODevice *)
func (this *QMessageAuthenticationCode) AddData_2(device *QIODevice /*777 QIODevice **/) bool {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:66
// index:0
// Public
// QByteArray result()
func (this *QMessageAuthenticationCode) Result() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QMessageAuthenticationCode6resultEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:68
// index:0
// Public static
// QByteArray hash(const class QByteArray &, const class QByteArray &, class QCryptographicHash::Algorithm)
func (this *QMessageAuthenticationCode) Hash(message *QByteArray, key *QByteArray, method int) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode4hashERK10QByteArrayS2_N18QCryptographicHash9AlgorithmE", ffiqt.FFI_TYPE_POINTER, message, key, method)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QMessageAuthenticationCode_Hash(message *QByteArray, key *QByteArray, method int) *QByteArray /*123*/ {
	var nilthis *QMessageAuthenticationCode
	rv := nilthis.Hash(message, key, method)
	return rv
}

//  body block end
