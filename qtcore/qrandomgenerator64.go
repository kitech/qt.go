package qtcore

// /usr/include/qt/QtCore/qrandom.h
// #include <qrandom.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QRandomGenerator64 struct {
	*QRandomGenerator
}

func (this *QRandomGenerator64) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QRandomGenerator.GetCthis()
	}
}
func (this *QRandomGenerator64) SetCthis(cthis unsafe.Pointer) {
	this.QRandomGenerator = NewQRandomGeneratorFromPointer(cthis)
}
func NewQRandomGenerator64FromPointer(cthis unsafe.Pointer) *QRandomGenerator64 {
	bcthis0 := NewQRandomGeneratorFromPointer(cthis)
	return &QRandomGenerator64{bcthis0}
}
func (*QRandomGenerator64) NewFromPointer(cthis unsafe.Pointer) *QRandomGenerator64 {
	return NewQRandomGenerator64FromPointer(cthis)
}

// /usr/include/qt/QtCore/qrandom.h:209
// index:0
// Public inline
// quint64 generate()
func (this *QRandomGenerator64) Generate() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator648generateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:215
// index:0
// Public inline
// void QRandomGenerator64(quint32)
func NewQRandomGenerator64(seedValue uint) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256) // 2512
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2Ej", ffiqt.FFI_TYPE_VOID, cthis, seedValue)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGenerator64FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:221
// index:1
// Public inline
// void QRandomGenerator64(const quint32 *, qsizetype)
func NewQRandomGenerator64_1(seedBuffer unsafe.Pointer /*666*/, len int64) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256) // 2512
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2EPKjx", ffiqt.FFI_TYPE_VOID, cthis, &seedBuffer, len)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGenerator64FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:227
// index:2
// Public inline
// void QRandomGenerator64(const quint32 *, const quint32 *)
func NewQRandomGenerator64_2(begin unsafe.Pointer /*666*/, end unsafe.Pointer /*666*/) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256) // 2512
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2EPKjS1_", ffiqt.FFI_TYPE_VOID, cthis, &begin, &end)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGenerator64FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:230
// index:3
// Public inline
// void QRandomGenerator64(const class QRandomGenerator &)
func NewQRandomGenerator64_3(other *QRandomGenerator) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256) // 2512
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2ERK16QRandomGenerator", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGenerator64FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:232
// index:0
// Public inline
// void discard(unsigned long long)
func (this *QRandomGenerator64) Discard(z uint64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator647discardEy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:239
// index:0
// Public static inline
// QRandomGenerator64::result_type min()
func (this *QRandomGenerator64) Min() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator643minEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint64(rv) // 222
}
func QRandomGenerator64_Min() uint64 {
	var nilthis *QRandomGenerator64
	rv := nilthis.Min()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:240
// index:0
// Public static inline
// QRandomGenerator64::result_type max()
func (this *QRandomGenerator64) Max() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator643maxEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint64(rv) // 222
}
func QRandomGenerator64_Max() uint64 {
	var nilthis *QRandomGenerator64
	rv := nilthis.Max()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:241
// index:0
// Public static
// QRandomGenerator64 * system()
func (this *QRandomGenerator64) System() *QRandomGenerator64 /*777 QRandomGenerator64 **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator646systemEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QRandomGenerator64_System() *QRandomGenerator64 /*777 QRandomGenerator64 **/ {
	var nilthis *QRandomGenerator64
	rv := nilthis.System()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:242
// index:0
// Public static
// QRandomGenerator64 * global()
func (this *QRandomGenerator64) Global() *QRandomGenerator64 /*777 QRandomGenerator64 **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator646globalEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QRandomGenerator64_Global() *QRandomGenerator64 /*777 QRandomGenerator64 **/ {
	var nilthis *QRandomGenerator64
	rv := nilthis.Global()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:243
// index:0
// Public static
// QRandomGenerator64 securelySeeded()
func (this *QRandomGenerator64) SecurelySeeded() *QRandomGenerator64 /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator6414securelySeededEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QRandomGenerator64_SecurelySeeded() *QRandomGenerator64 /*123*/ {
	var nilthis *QRandomGenerator64
	rv := nilthis.SecurelySeeded()
	return rv
}

//  body block end
