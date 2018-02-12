package qtcore

// /usr/include/qt/QtCore/qsize.h
// #include <qsize.h>
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
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QSizeF struct {
	*qtrt.CObject
}

func (this *QSizeF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSizeF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSizeFFromPointer(cthis unsafe.Pointer) *QSizeF {
	return &QSizeF{&qtrt.CObject{cthis}}
}
func (*QSizeF) NewFromPointer(cthis unsafe.Pointer) *QSizeF {
	return NewQSizeFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsize.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSizeF()
func NewQSizeF() *QSizeF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizeF)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:219
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSizeF(const QSize &)
func NewQSizeF_1(sz *QSize) *QSizeF {
	var convArg0 = sz.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFC2ERK5QSize", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizeF)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:220
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QSizeF(qreal, qreal)
func NewQSizeF_2(w float64, h float64) *QSizeF {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFC2Edd", qtrt.FFI_TYPE_POINTER, w, h)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizeF)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QSizeF) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QSizeF) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:224
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QSizeF) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal width()
func (this *QSizeF) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qsize.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal height()
func (this *QSizeF) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qsize.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(qreal)
func (this *QSizeF) SetWidth(w float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(qreal)
func (this *QSizeF) SetHeight(h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void transpose()
func (this *QSizeF) Transpose() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF9transposeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF transposed()
func (this *QSizeF) Transposed() *QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void scale(qreal, qreal, Qt::AspectRatioMode)
func (this *QSizeF) Scale(w float64, h float64, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF5scaleEddN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:234
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void scale(const QSizeF &, Qt::AspectRatioMode)
func (this *QSizeF) Scale_1(s *QSizeF, mode int) {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF5scaleERKS_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:235
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF scaled(qreal, qreal, Qt::AspectRatioMode)
func (this *QSizeF) Scaled(w float64, h float64, mode int) *QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6scaledEddN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:236
// index:1
// Public Visibility=Default Availability=Available
// [16] QSizeF scaled(const QSizeF &, Qt::AspectRatioMode)
func (this *QSizeF) Scaled_1(s *QSizeF, mode int) *QSizeF /*123*/ {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6scaledERKS_N2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF expandedTo(const QSizeF &)
func (this *QSizeF) ExpandedTo(arg0 *QSizeF) *QSizeF /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF10expandedToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:239
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QSizeF boundedTo(const QSizeF &)
func (this *QSizeF) BoundedTo(arg0 *QSizeF) *QSizeF /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF9boundedToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:241
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal & rwidth()
func (this *QSizeF) Rwidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF6rwidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float64", rv).(float64) // 3331
}

// /usr/include/qt/QtCore/qsize.h:242
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal & rheight()
func (this *QSizeF) Rheight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeF7rheightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float64", rv).(float64) // 3331
}

// /usr/include/qt/QtCore/qsize.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize toSize()
func (this *QSizeF) ToSize() *QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QSizeF6toSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

func DeleteQSizeF(this *QSizeF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QSizeFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		qtrt.KeepMe()
	}
}

//  keep block end
