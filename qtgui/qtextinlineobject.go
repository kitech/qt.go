package qtgui

// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 72
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTextInlineObject struct {
	*qtrt.CObject
}

func (this *QTextInlineObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextInlineObject) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextInlineObjectFromPointer(cthis unsafe.Pointer) *QTextInlineObject {
	return &QTextInlineObject{&qtrt.CObject{cthis}}
}
func (*QTextInlineObject) NewFromPointer(cthis unsafe.Pointer) *QTextInlineObject {
	return NewQTextInlineObjectFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextlayout.h:72
// index:0
// Public inline
// void QTextInlineObject()
func NewQTextInlineObject() *QTextInlineObject {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObjectC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextInlineObjectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:73
// index:0
// Public inline
// bool isValid()
func (this *QTextInlineObject) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:75
// index:0
// Public
// QRectF rect()
func (this *QTextInlineObject) Rect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject4rectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:76
// index:0
// Public
// qreal width()
func (this *QTextInlineObject) Width() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:77
// index:0
// Public
// qreal ascent()
func (this *QTextInlineObject) Ascent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject6ascentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:78
// index:0
// Public
// qreal descent()
func (this *QTextInlineObject) Descent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject7descentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:79
// index:0
// Public
// qreal height()
func (this *QTextInlineObject) Height() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:81
// index:0
// Public
// Qt::LayoutDirection textDirection()
func (this *QTextInlineObject) TextDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject13textDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:83
// index:0
// Public
// void setWidth(qreal)
func (this *QTextInlineObject) SetWidth(w float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObject8setWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:84
// index:0
// Public
// void setAscent(qreal)
func (this *QTextInlineObject) SetAscent(a float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObject9setAscentEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:85
// index:0
// Public
// void setDescent(qreal)
func (this *QTextInlineObject) SetDescent(d float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObject10setDescentEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), d)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:87
// index:0
// Public
// int textPosition()
func (this *QTextInlineObject) TextPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject12textPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextlayout.h:89
// index:0
// Public
// int formatIndex()
func (this *QTextInlineObject) FormatIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject11formatIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextlayout.h:90
// index:0
// Public
// QTextFormat format()
func (this *QTextInlineObject) Format() *QTextFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject6formatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
