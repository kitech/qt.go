package qtcore

// /usr/include/qt/QtCore/qcborvalue.h
// #include <qcborvalue.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QCborParserError struct {
	*qtrt.CObject
}
type QCborParserError_ITF interface {
	QCborParserError_PTR() *QCborParserError
}

func (ptr *QCborParserError) QCborParserError_PTR() *QCborParserError { return ptr }

func (this *QCborParserError) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCborParserError) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCborParserErrorFromPointer(cthis unsafe.Pointer) *QCborParserError {
	return &QCborParserError{&qtrt.CObject{cthis}}
}
func (*QCborParserError) NewFromPointer(cthis unsafe.Pointer) *QCborParserError {
	return NewQCborParserErrorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcborvalue.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString errorString() const

/*

 */
func (this *QCborParserError) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCborParserError11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

func DeleteQCborParserError(this *QCborParserError) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCborParserErrorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10349() {
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
