//  header block begin
// /usr/include/qt/QtCore/qrandom.h
// #include <qrandom.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QRandomGenerator64 struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qrandom.h:209
// index:0
// inline
// quint64 generate()
func (this *QRandomGenerator64) Generate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator648generateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:215
// index:0
// inline
// void QRandomGenerator64(quint32)
func NewQRandomGenerator64(seedValue uint) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2Ej", ffiqt.FFI_TYPE_VOID, cthis, &seedValue)
	gopp.ErrPrint(err, rv)
	return &QRandomGenerator64{cthis}
}

// /usr/include/qt/QtCore/qrandom.h:221
// index:1
// inline
// void QRandomGenerator64(const quint32 *, qsizetype)
func NewQRandomGenerator64_1(seedBuffer unsafe.Pointer, len int64) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2EPKjx", ffiqt.FFI_TYPE_VOID, cthis, seedBuffer, &len)
	gopp.ErrPrint(err, rv)
	return &QRandomGenerator64{cthis}
}

// /usr/include/qt/QtCore/qrandom.h:224
// index:2
// inline
// void QRandomGenerator64(std::seed_seq &)
func NewQRandomGenerator64_2(sseq int) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2ERSt8seed_seq", ffiqt.FFI_TYPE_VOID, cthis, &sseq)
	gopp.ErrPrint(err, rv)
	return &QRandomGenerator64{cthis}
}

// /usr/include/qt/QtCore/qrandom.h:227
// index:3
// inline
// void QRandomGenerator64(const quint32 *, const quint32 *)
func NewQRandomGenerator64_3(begin unsafe.Pointer, end unsafe.Pointer) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2EPKjS1_", ffiqt.FFI_TYPE_VOID, cthis, begin, end)
	gopp.ErrPrint(err, rv)
	return &QRandomGenerator64{cthis}
}

// /usr/include/qt/QtCore/qrandom.h:230
// index:4
// inline
// void QRandomGenerator64(const class QRandomGenerator &)
func NewQRandomGenerator64_4(other unsafe.Pointer) *QRandomGenerator64 {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator64C2ERK16QRandomGenerator", ffiqt.FFI_TYPE_VOID, cthis, other)
	gopp.ErrPrint(err, rv)
	return &QRandomGenerator64{cthis}
}

// /usr/include/qt/QtCore/qrandom.h:232
// index:0
// inline
// void discard(unsigned long long)
func (this *QRandomGenerator64) Discard(z uint64) {
	// 0: (, unsigned long long z), (&z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator647discardEy", ffiqt.FFI_TYPE_VOID, this.cthis, &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:239
// index:0
// static inline
// QRandomGenerator64::result_type min()
func (this *QRandomGenerator64) Min() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator643minEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator64_Min() {
	// 0: (), ()
	var nilthis *QRandomGenerator64
	nilthis.Min()
}

// /usr/include/qt/QtCore/qrandom.h:240
// index:0
// static inline
// QRandomGenerator64::result_type max()
func (this *QRandomGenerator64) Max() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator643maxEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator64_Max() {
	// 0: (), ()
	var nilthis *QRandomGenerator64
	nilthis.Max()
}

// /usr/include/qt/QtCore/qrandom.h:241
// index:0
// static
// QRandomGenerator64 * system()
func (this *QRandomGenerator64) System() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator646systemEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator64_System() {
	// 0: (), ()
	var nilthis *QRandomGenerator64
	nilthis.System()
}

// /usr/include/qt/QtCore/qrandom.h:242
// index:0
// static
// QRandomGenerator64 * global()
func (this *QRandomGenerator64) Global() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator646globalEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator64_Global() {
	// 0: (), ()
	var nilthis *QRandomGenerator64
	nilthis.Global()
}

// /usr/include/qt/QtCore/qrandom.h:243
// index:0
// static
// QRandomGenerator64 securelySeeded()
func (this *QRandomGenerator64) SecurelySeeded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRandomGenerator6414securelySeededEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator64_SecurelySeeded() {
	// 0: (), ()
	var nilthis *QRandomGenerator64
	nilthis.SecurelySeeded()
}

//  body block end
