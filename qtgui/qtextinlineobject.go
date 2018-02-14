package qtgui

// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 80
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextInlineObject struct {
	*qtrt.CObject
}
type QTextInlineObject_ITF interface {
	QTextInlineObject_PTR() *QTextInlineObject
}

func (ptr *QTextInlineObject) QTextInlineObject_PTR() *QTextInlineObject { return ptr }

func (this *QTextInlineObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextInlineObject) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextInlineObjectFromPointer(cthis unsafe.Pointer) *QTextInlineObject {
	return &QTextInlineObject{&qtrt.CObject{cthis}}
}
func (*QTextInlineObject) NewFromPointer(cthis unsafe.Pointer) *QTextInlineObject {
	return NewQTextInlineObjectFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextlayout.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextInlineObject()
func NewQTextInlineObject() *QTextInlineObject {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QTextInlineObjectC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextInlineObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextInlineObject)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextInlineObject) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect()
func (this *QTextInlineObject) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width()
func (this *QTextInlineObject) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent()
func (this *QTextInlineObject) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent()
func (this *QTextInlineObject) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal height()
func (this *QTextInlineObject) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection textDirection()
func (this *QTextInlineObject) TextDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject13textDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(qreal)
func (this *QTextInlineObject) SetWidth(w float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QTextInlineObject8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAscent(qreal)
func (this *QTextInlineObject) SetAscent(a float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QTextInlineObject9setAscentEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescent(qreal)
func (this *QTextInlineObject) SetDescent(d float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QTextInlineObject10setDescentEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), d)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int textPosition()
func (this *QTextInlineObject) TextPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject12textPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int formatIndex()
func (this *QTextInlineObject) FormatIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject11formatIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:90
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFormat format()
func (this *QTextInlineObject) Format() *QTextFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QTextInlineObject6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFormat)
	return rv2
}

func DeleteQTextInlineObject(this *QTextInlineObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QTextInlineObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
