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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMessageAuthenticationCodeFromPointer(cthis unsafe.Pointer) *QMessageAuthenticationCode {
	return &QMessageAuthenticationCode{&qtrt.CObject{cthis}}
}
func (*QMessageAuthenticationCode) NewFromPointer(cthis unsafe.Pointer) *QMessageAuthenticationCode {
	return NewQMessageAuthenticationCodeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMessageAuthenticationCode(QCryptographicHash::Algorithm, const QByteArray &)
func NewQMessageAuthenticationCode(method int, key *QByteArray) *QMessageAuthenticationCode {
	var convArg1 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeC2EN18QCryptographicHash9AlgorithmERK10QByteArray", qtrt.FFI_TYPE_POINTER, method, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageAuthenticationCodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMessageAuthenticationCode)
	return gothis
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMessageAuthenticationCode()
func DeleteQMessageAuthenticationCode(this *QMessageAuthenticationCode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QMessageAuthenticationCode) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QByteArray &)
func (this *QMessageAuthenticationCode) SetKey(key *QByteArray) {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode6setKeyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addData(const char *, int)
func (this *QMessageAuthenticationCode) AddData(data string, length int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addData(const QByteArray &)
func (this *QMessageAuthenticationCode) AddData_1(data *QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:64
// index:2
// Public Visibility=Default Availability=Available
// [1] bool addData(QIODevice *)
func (this *QMessageAuthenticationCode) AddData_2(device *QIODevice /*777 QIODevice **/) bool {
	var convArg0 = device.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray result()
func (this *QMessageAuthenticationCode) Result() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QMessageAuthenticationCode6resultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray hash(const QByteArray &, const QByteArray &, QCryptographicHash::Algorithm)
func (this *QMessageAuthenticationCode) Hash(message *QByteArray, key *QByteArray, method int) *QByteArray /*123*/ {
	var convArg0 = message.GetCthis()
	var convArg1 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode4hashERK10QByteArrayS2_N18QCryptographicHash9AlgorithmE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, method)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QMessageAuthenticationCode_Hash(message *QByteArray, key *QByteArray, method int) *QByteArray /*123*/ {
	var nilthis *QMessageAuthenticationCode
	rv := nilthis.Hash(message, key, method)
	return rv
}

//  body block end
