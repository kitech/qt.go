//  header block begin
// /usr/include/qt/QtCore/qsize.h
// #include <qsize.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QSizeF struct {
	*qtrt.CObject
}

func (this *QSizeF) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQSizeFFromPointer(cthis unsafe.Pointer) *QSizeF {
	return &QSizeF{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qsize.h:218
// index:0
// Public inline
// void QSizeF()
func NewQSizeF() *QSizeF {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeFC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:219
// index:1
// Public inline
// void QSizeF(const class QSize &)
func NewQSizeF_1(sz *QSize) *QSizeF {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = sz.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeFC2ERK5QSize", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:220
// index:2
// Public inline
// void QSizeF(qreal, qreal)
func NewQSizeF_2(w float64, h float64) *QSizeF {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeFC2Edd", ffiqt.FFI_TYPE_VOID, cthis, &w, &h)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:222
// index:0
// Public inline
// bool isNull()
func (this *QSizeF) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:223
// index:0
// Public inline
// bool isEmpty()
func (this *QSizeF) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:224
// index:0
// Public inline
// bool isValid()
func (this *QSizeF) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:226
// index:0
// Public inline
// qreal width()
func (this *QSizeF) Width() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:227
// index:0
// Public inline
// qreal height()
func (this *QSizeF) Height() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:228
// index:0
// Public inline
// void setWidth(qreal)
func (this *QSizeF) SetWidth(w float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeF8setWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:229
// index:0
// Public inline
// void setHeight(qreal)
func (this *QSizeF) SetHeight(h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeF9setHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:230
// index:0
// Public
// void transpose()
func (this *QSizeF) Transpose() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeF9transposeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:231
// index:0
// Public inline
// QSizeF transposed()
func (this *QSizeF) Transposed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF10transposedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:233
// index:0
// Public inline
// void scale(qreal, qreal, Qt::AspectRatioMode)
func (this *QSizeF) Scale(w float64, h float64, mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeF5scaleEddN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w, &h, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:234
// index:1
// Public inline
// void scale(const class QSizeF &, Qt::AspectRatioMode)
func (this *QSizeF) Scale_1(s *QSizeF, mode int) {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeF5scaleERKS_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:235
// index:0
// Public
// QSizeF scaled(qreal, qreal, Qt::AspectRatioMode)
func (this *QSizeF) Scaled(w float64, h float64, mode int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF6scaledEddN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w, &h, &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:236
// index:1
// Public
// QSizeF scaled(const class QSizeF &, Qt::AspectRatioMode)
func (this *QSizeF) Scaled_1(s *QSizeF, mode int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF6scaledERKS_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:238
// index:0
// Public inline
// QSizeF expandedTo(const class QSizeF &)
func (this *QSizeF) ExpandedTo(arg0 *QSizeF) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF10expandedToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:239
// index:0
// Public inline
// QSizeF boundedTo(const class QSizeF &)
func (this *QSizeF) BoundedTo(arg0 *QSizeF) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF9boundedToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:241
// index:0
// Public inline
// qreal & rwidth()
func (this *QSizeF) Rwidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeF6rwidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:242
// index:0
// Public inline
// qreal & rheight()
func (this *QSizeF) Rheight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QSizeF7rheightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsize.h:257
// index:0
// Public inline
// QSize toSize()
func (this *QSizeF) ToSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QSizeF6toSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
