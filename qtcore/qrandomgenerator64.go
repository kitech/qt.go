package qtcore

// /usr/include/qt/QtCore/qrandom.h
// #include <qrandom.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QRandomGenerator64 struct {
	*QRandomGenerator
}
type QRandomGenerator64_ITF interface {
	QRandomGenerator_ITF
	QRandomGenerator64_PTR() *QRandomGenerator64
}

func (ptr *QRandomGenerator64) QRandomGenerator64_PTR() *QRandomGenerator64 { return ptr }

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
// Public inline Visibility=Default Availability=Available
// [8] quint64 generate()
func (this *QRandomGenerator64) Generate() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator648generateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:215
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator64(quint32)
func NewQRandomGenerator64(seedValue uint) *QRandomGenerator64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator64C2Ej", qtrt.FFI_TYPE_POINTER, seedValue)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator64)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:221
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator64(const quint32 *, qsizetype)
func NewQRandomGenerator64_1(seedBuffer unsafe.Pointer /*666*/, len int64) *QRandomGenerator64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator64C2EPKjx", qtrt.FFI_TYPE_POINTER, &seedBuffer, len)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator64)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:227
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator64(const quint32 *, const quint32 *)
func NewQRandomGenerator64_2(begin unsafe.Pointer /*666*/, end unsafe.Pointer /*666*/) *QRandomGenerator64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator64C2EPKjS1_", qtrt.FFI_TYPE_POINTER, &begin, &end)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator64)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:230
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator64(const QRandomGenerator &)
func NewQRandomGenerator64_3(other *QRandomGenerator) *QRandomGenerator64 {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator64C2ERK16QRandomGenerator", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator64)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void discard(unsigned long long)
func (this *QRandomGenerator64) Discard(z uint64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator647discardEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:239
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRandomGenerator64::result_type min()
func (this *QRandomGenerator64) Min() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator643minEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}
func QRandomGenerator64_Min() uint64 {
	var nilthis *QRandomGenerator64
	rv := nilthis.Min()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:240
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRandomGenerator64::result_type max()
func (this *QRandomGenerator64) Max() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator643maxEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}
func QRandomGenerator64_Max() uint64 {
	var nilthis *QRandomGenerator64
	rv := nilthis.Max()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:241
// index:0
// Public static Visibility=Default Availability=Available
// [8] QRandomGenerator64 * system()
func (this *QRandomGenerator64) System() *QRandomGenerator64 /*777 QRandomGenerator64 **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator646systemEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QRandomGenerator64_System() *QRandomGenerator64 /*777 QRandomGenerator64 **/ {
	var nilthis *QRandomGenerator64
	rv := nilthis.System()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:242
// index:0
// Public static Visibility=Default Availability=Available
// [8] QRandomGenerator64 * global()
func (this *QRandomGenerator64) Global() *QRandomGenerator64 /*777 QRandomGenerator64 **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator646globalEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QRandomGenerator64_Global() *QRandomGenerator64 /*777 QRandomGenerator64 **/ {
	var nilthis *QRandomGenerator64
	rv := nilthis.Global()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:243
// index:0
// Public static Visibility=Default Availability=Available
// [2512] QRandomGenerator64 securelySeeded()
func (this *QRandomGenerator64) SecurelySeeded() *QRandomGenerator64 /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator6414securelySeededEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRandomGenerator64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRandomGenerator64)
	return rv2
}
func QRandomGenerator64_SecurelySeeded() *QRandomGenerator64 /*123*/ {
	var nilthis *QRandomGenerator64
	rv := nilthis.SecurelySeeded()
	return rv
}

func DeleteQRandomGenerator64(this *QRandomGenerator64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRandomGenerator64D2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
