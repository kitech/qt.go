package qtcore

// /usr/include/qt/QtCore/qcborcommon.h
// #include <qcborcommon.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QCborError struct {
	*qtrt.CObject
}
type QCborError_ITF interface {
	QCborError_PTR() *QCborError
}

func (ptr *QCborError) QCborError_PTR() *QCborError { return ptr }

func (this *QCborError) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCborError) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCborErrorFromPointer(cthis unsafe.Pointer) *QCborError {
	return &QCborError{&qtrt.CObject{cthis}}
}
func (*QCborError) NewFromPointer(cthis unsafe.Pointer) *QCborError {
	return NewQCborErrorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcborcommon.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*

 */
func (this *QCborError) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborError8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

func DeleteQCborError(this *QCborError) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborErrorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QCborError__Code = int

//
const QCborError__UnknownError QCborError__Code = 1

//
const QCborError__AdvancePastEnd QCborError__Code = 3

//
const QCborError__InputOutputError QCborError__Code = 4

//
const QCborError__GarbageAtEnd QCborError__Code = 256

//
const QCborError__EndOfFile QCborError__Code = 257

//
const QCborError__UnexpectedBreak QCborError__Code = 258

//
const QCborError__UnknownType QCborError__Code = 259

//
const QCborError__IllegalType QCborError__Code = 260

//
const QCborError__IllegalNumber QCborError__Code = 261

//
const QCborError__IllegalSimpleType QCborError__Code = 262

//
const QCborError__InvalidUtf8String QCborError__Code = 516

//
const QCborError__DataTooLarge QCborError__Code = 1024

//
const QCborError__NestingTooDeep QCborError__Code = 1025

//
const QCborError__UnsupportedType QCborError__Code = 1026

//
const QCborError__NoError QCborError__Code = 0

func (this *QCborError) CodeItemName(val int) string {
	switch val {
	case QCborError__UnknownError: // 1
		return "UnknownError"
	case QCborError__AdvancePastEnd: // 3
		return "AdvancePastEnd"
	case QCborError__InputOutputError: // 4
		return "InputOutputError"
	case QCborError__GarbageAtEnd: // 256
		return "GarbageAtEnd"
	case QCborError__EndOfFile: // 257
		return "EndOfFile"
	case QCborError__UnexpectedBreak: // 258
		return "UnexpectedBreak"
	case QCborError__UnknownType: // 259
		return "UnknownType"
	case QCborError__IllegalType: // 260
		return "IllegalType"
	case QCborError__IllegalNumber: // 261
		return "IllegalNumber"
	case QCborError__IllegalSimpleType: // 262
		return "IllegalSimpleType"
	case QCborError__InvalidUtf8String: // 516
		return "InvalidUtf8String"
	case QCborError__DataTooLarge: // 1024
		return "DataTooLarge"
	case QCborError__NestingTooDeep: // 1025
		return "NestingTooDeep"
	case QCborError__UnsupportedType: // 1026
		return "UnsupportedType"
	case QCborError__NoError: // 0
		return "NoError"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCborError_CodeItemName(val int) string {
	var nilthis *QCborError
	return nilthis.CodeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10337() {
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
