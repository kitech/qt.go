package qtcore

// /usr/include/qt/QtCore/qsize.h
// #include <qsize.h>
// #include <QtCore>

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
import "mkuse/cffiqt"
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSize) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSizeFromPointer(cthis unsafe.Pointer) *QSize {
	return &QSize{&qtrt.CObject{cthis}}
}
func (*QSize) NewFromPointer(cthis unsafe.Pointer) *QSize {
	return NewQSizeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsize.h:55
// index:0
// Public inline
// void QSize()
func NewQSize() *QSize {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSizeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:56
// index:1
// Public inline
// void QSize(int, int)
func NewQSize_1(w int, h int) *QSize {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSizeC2Eii", ffiqt.FFI_TYPE_VOID, cthis, w, h)
	gopp.ErrPrint(err, rv)
	gothis := NewQSizeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsize.h:58
// index:0
// Public inline
// bool isNull()
func (this *QSize) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:59
// index:0
// Public inline
// bool isEmpty()
func (this *QSize) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:60
// index:0
// Public inline
// bool isValid()
func (this *QSize) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsize.h:62
// index:0
// Public inline
// int width()
func (this *QSize) Width() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qsize.h:63
// index:0
// Public inline
// int height()
func (this *QSize) Height() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qsize.h:64
// index:0
// Public inline
// void setWidth(int)
func (this *QSize) SetWidth(w int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize8setWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:65
// index:0
// Public inline
// void setHeight(int)
func (this *QSize) SetHeight(h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize9setHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:66
// index:0
// Public
// void transpose()
func (this *QSize) Transpose() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize9transposeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:67
// index:0
// Public inline
// QSize transposed()
func (this *QSize) Transposed() *QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize10transposedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:69
// index:0
// Public inline
// void scale(int, int, Qt::AspectRatioMode)
func (this *QSize) Scale(w int, h int, mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize5scaleEiiN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:70
// index:1
// Public inline
// void scale(const class QSize &, Qt::AspectRatioMode)
func (this *QSize) Scale_1(s *QSize, mode int) {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize5scaleERKS_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsize.h:71
// index:0
// Public
// QSize scaled(int, int, Qt::AspectRatioMode)
func (this *QSize) Scaled(w int, h int, mode int) *QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize6scaledEiiN2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), w, h, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:72
// index:1
// Public
// QSize scaled(const class QSize &, Qt::AspectRatioMode)
func (this *QSize) Scaled_1(s *QSize, mode int) *QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize6scaledERKS_N2Qt15AspectRatioModeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:74
// index:0
// Public inline
// QSize expandedTo(const class QSize &)
func (this *QSize) ExpandedTo(arg0 *QSize) *QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize10expandedToERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:75
// index:0
// Public inline
// QSize boundedTo(const class QSize &)
func (this *QSize) BoundedTo(arg0 *QSize) *QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QSize9boundedToERKS_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qsize.h:77
// index:0
// Public inline
// int & rwidth()
func (this *QSize) Rwidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize6rwidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 3331
}

// /usr/include/qt/QtCore/qsize.h:78
// index:0
// Public inline
// int & rheight()
func (this *QSize) Rheight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QSize7rheightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 3331
}

//  body block end
