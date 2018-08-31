package qtcore

// /usr/include/qt/QtCore/qmessageauthenticationcode.h
// #include <qmessageauthenticationcode.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QMessageAuthenticationCode struct {
	*qtrt.CObject
}
type QMessageAuthenticationCode_ITF interface {
	QMessageAuthenticationCode_PTR() *QMessageAuthenticationCode
}

func (ptr *QMessageAuthenticationCode) QMessageAuthenticationCode_PTR() *QMessageAuthenticationCode {
	return ptr
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

/*
Constructs an object that can be used to create a cryptographic hash from data using method method and key key.
*/
func (*QMessageAuthenticationCode) NewForInherit(method int, key QByteArray_ITF) *QMessageAuthenticationCode {
	return NewQMessageAuthenticationCode(method, key)
}
func NewQMessageAuthenticationCode(method int, key QByteArray_ITF) *QMessageAuthenticationCode {
	var convArg1 unsafe.Pointer
	if key != nil && key.QByteArray_PTR() != nil {
		convArg1 = key.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeC2EN18QCryptographicHash9AlgorithmERK10QByteArray", qtrt.FFI_TYPE_POINTER, method, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMessageAuthenticationCodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMessageAuthenticationCode)
	return gothis
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMessageAuthenticationCode(QCryptographicHash::Algorithm, const QByteArray &)

/*
Constructs an object that can be used to create a cryptographic hash from data using method method and key key.
*/
func (*QMessageAuthenticationCode) NewForInherit__(method int) *QMessageAuthenticationCode {
	return NewQMessageAuthenticationCode__(method)
}
func NewQMessageAuthenticationCode__(method int) *QMessageAuthenticationCode {
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeC2EN18QCryptographicHash9AlgorithmERK10QByteArray", qtrt.FFI_TYPE_POINTER, method, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMessageAuthenticationCodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMessageAuthenticationCode)
	return gothis
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMessageAuthenticationCode()

/*

 */
func DeleteQMessageAuthenticationCode(this *QMessageAuthenticationCode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*
Resets message data. Calling this method doesn't affect the key.
*/
func (this *QMessageAuthenticationCode) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QByteArray &)

/*
Sets secret key. Calling this method automatically resets the object state.
*/
func (this *QMessageAuthenticationCode) SetKey(key QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if key != nil && key.QByteArray_PTR() != nil {
		convArg0 = key.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode6setKeyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addData(const char *, int)

/*
Adds the first length chars of data to the message.
*/
func (this *QMessageAuthenticationCode) AddData(data string, length int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, length)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addData(const QByteArray &)

/*
Adds the first length chars of data to the message.
*/
func (this *QMessageAuthenticationCode) AddData_1(data QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:64
// index:2
// Public Visibility=Default Availability=Available
// [1] bool addData(QIODevice *)

/*
Adds the first length chars of data to the message.
*/
func (this *QMessageAuthenticationCode) AddData_2(device QIODevice_ITF /*777 QIODevice **/) bool {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode7addDataEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray result() const

/*
Returns the final authentication code.

See also QByteArray::toHex().
*/
func (this *QMessageAuthenticationCode) Result() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QMessageAuthenticationCode6resultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmessageauthenticationcode.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray hash(const QByteArray &, const QByteArray &, QCryptographicHash::Algorithm)

/*
Returns the authentication code for the message message using the key key and the method method.
*/
func (this *QMessageAuthenticationCode) Hash(message QByteArray_ITF, key QByteArray_ITF, method int) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if message != nil && message.QByteArray_PTR() != nil {
		convArg0 = message.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if key != nil && key.QByteArray_PTR() != nil {
		convArg1 = key.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMessageAuthenticationCode4hashERK10QByteArrayS2_N18QCryptographicHash9AlgorithmE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, method)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QMessageAuthenticationCode_Hash(message QByteArray_ITF, key QByteArray_ITF, method int) *QByteArray /*123*/ {
	var nilthis *QMessageAuthenticationCode
	rv := nilthis.Hash(message, key, method)
	return rv
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
