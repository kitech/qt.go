//  header block begin
// /usr/include/qt/QtCore/qsize.h
// #include <qsize.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QSize struct {
	*qtrt.CObject
}

func (this *QSize) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qsize.h:55
// index:0
// inline
// void QSize()
func NewQSize() *QSize {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSizeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizeFromPointer(cthis)
	return gothis
}
func NewQSizeFromPointer(cthis unsafe.Pointer) *QSize {
	return &QSize{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qsize.h:56
// index:1
// inline
// void QSize(int, int)
func NewQSize_1(w int, h int) *QSize {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSizeC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &w, &h)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:58
// index:0
// inline
// bool isNull()
func (this *QSize) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:59
// index:0
// inline
// bool isEmpty()
func (this *QSize) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:60
// index:0
// inline
// bool isValid()
func (this *QSize) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:62
// index:0
// inline
// int width()
func (this *QSize) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:63
// index:0
// inline
// int height()
func (this *QSize) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:64
// index:0
// inline
// void setWidth(int)
func (this *QSize) SetWidth(w int) {
	// 0: (, w int), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize8setWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:65
// index:0
// inline
// void setHeight(int)
func (this *QSize) SetHeight(h int) {
	// 0: (, h int), (&h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize9setHeightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:66
// index:0
// void transpose()
func (this *QSize) Transpose() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize9transposeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:67
// index:0
// inline
// QSize transposed()
func (this *QSize) Transposed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize10transposedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:69
// index:0
// inline
// void scale(int, int, Qt::AspectRatioMode)
func (this *QSize) Scale(w int, h int, mode int) {
	// 0: (, w int, h int, mode Qt::AspectRatioMode), (&w, &h, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize5scaleEiiN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:70
// index:1
// inline
// void scale(const class QSize &, Qt::AspectRatioMode)
func (this *QSize) Scale_1(s unsafe.Pointer, mode int) {
	// 1: (, s const QSize &, mode Qt::AspectRatioMode), (s, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize5scaleERKS_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:71
// index:0
// QSize scaled(int, int, Qt::AspectRatioMode)
func (this *QSize) Scaled(w int, h int, mode int) {
	// 0: (, w int, h int, mode Qt::AspectRatioMode), (&w, &h, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize6scaledEiiN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:72
// index:1
// QSize scaled(const class QSize &, Qt::AspectRatioMode)
func (this *QSize) Scaled_1(s unsafe.Pointer, mode int) {
	// 1: (, s const QSize &, mode Qt::AspectRatioMode), (s, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize6scaledERKS_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:74
// index:0
// inline
// QSize expandedTo(const class QSize &)
func (this *QSize) ExpandedTo(arg0 unsafe.Pointer) {
	// 0: (, const QSize & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize10expandedToERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:75
// index:0
// inline
// QSize boundedTo(const class QSize &)
func (this *QSize) BoundedTo(arg0 unsafe.Pointer) {
	// 0: (, const QSize & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize9boundedToERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:77
// index:0
// inline
// int & rwidth()
func (this *QSize) Rwidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize6rwidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:78
// index:0
// inline
// int & rheight()
func (this *QSize) Rheight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize7rheightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
