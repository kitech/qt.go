//  header block begin
// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 74
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextlayout.h:71
// index:0
// inline
// void QTextInlineObject(int, class QTextEngine *)
func NewQTextInlineObject(i int, e unsafe.Pointer) *QTextInlineObject {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObjectC2EiP11QTextEngine", ffiqt.FFI_TYPE_VOID, cthis, &i, e)
	gopp.ErrPrint(err, rv)
	return &QTextInlineObject{cthis}
}

// /usr/include/qt/QtGui/qtextlayout.h:72
// index:1
// inline
// void QTextInlineObject()
func NewQTextInlineObject_1() *QTextInlineObject {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObjectC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextInlineObject{cthis}
}

// /usr/include/qt/QtGui/qtextlayout.h:73
// index:0
// inline
// bool isValid()
func (this *QTextInlineObject) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:75
// index:0
// QRectF rect()
func (this *QTextInlineObject) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject4rectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:76
// index:0
// qreal width()
func (this *QTextInlineObject) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject5widthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:77
// index:0
// qreal ascent()
func (this *QTextInlineObject) Ascent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject6ascentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:78
// index:0
// qreal descent()
func (this *QTextInlineObject) Descent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject7descentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:79
// index:0
// qreal height()
func (this *QTextInlineObject) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject6heightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:81
// index:0
// Qt::LayoutDirection textDirection()
func (this *QTextInlineObject) TextDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject13textDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:83
// index:0
// void setWidth(qreal)
func (this *QTextInlineObject) SetWidth(w float64) {
	// 0: (, qreal w), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObject8setWidthEd", ffiqt.FFI_TYPE_VOID, this.cthis, &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:84
// index:0
// void setAscent(qreal)
func (this *QTextInlineObject) SetAscent(a float64) {
	// 0: (, qreal a), (&a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObject9setAscentEd", ffiqt.FFI_TYPE_VOID, this.cthis, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:85
// index:0
// void setDescent(qreal)
func (this *QTextInlineObject) SetDescent(d float64) {
	// 0: (, qreal d), (&d)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QTextInlineObject10setDescentEd", ffiqt.FFI_TYPE_VOID, this.cthis, &d)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:87
// index:0
// int textPosition()
func (this *QTextInlineObject) TextPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject12textPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:89
// index:0
// int formatIndex()
func (this *QTextInlineObject) FormatIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject11formatIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:90
// index:0
// QTextFormat format()
func (this *QTextInlineObject) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QTextInlineObject6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
