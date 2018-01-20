//  header block begin
// /usr/include/qt/QtCore/qrandom.h
// #include <qrandom.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QRandomGenerator struct {
	*qtrt.CObject
}

func (this *QRandomGenerator) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qrandom.h:55
// index:0
// inline
// void QRandomGenerator(quint32)
func NewQRandomGenerator(seedValue uint) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2Ej", ffiqt.FFI_TYPE_VOID, cthis, &seedValue)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}
func NewQRandomGeneratorFromPointer(cthis unsafe.Pointer) *QRandomGenerator {
	return &QRandomGenerator{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qrandom.h:61
// index:1
// inline
// void QRandomGenerator(const quint32 *, qsizetype)
func NewQRandomGenerator_1(seedBuffer unsafe.Pointer, len int64) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2EPKjx", ffiqt.FFI_TYPE_VOID, cthis, seedBuffer, &len)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:64
// index:2
// void QRandomGenerator(std::seed_seq &)
func NewQRandomGenerator_2(sseq int) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2ERSt8seed_seq", ffiqt.FFI_TYPE_VOID, cthis, &sseq)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:65
// index:3
// void QRandomGenerator(const quint32 *, const quint32 *)
func NewQRandomGenerator_3(begin unsafe.Pointer, end unsafe.Pointer) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2EPKjS1_", ffiqt.FFI_TYPE_VOID, cthis, begin, end)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:172
// index:4
// void QRandomGenerator(enum QRandomGenerator::System)
func NewQRandomGenerator_4(arg0 int) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2ENS_6SystemE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:77
// index:0
// inline
// quint32 generate()
func (this *QRandomGenerator) Generate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator8generateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:152
// index:1
// inline
// void generate(quint32 *, quint32 *)
func (this *QRandomGenerator) Generate_1(begin unsafe.Pointer, end unsafe.Pointer) {
	// 1: (, begin quint32 *, end quint32 *), (begin, end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator8generateEPjS0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), begin, end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:84
// index:0
// inline
// quint64 generate64()
func (this *QRandomGenerator) Generate64() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator10generate64Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:91
// index:0
// inline
// double generateDouble()
func (this *QRandomGenerator) GenerateDouble() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator14generateDoubleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:105
// index:0
// inline
// double bounded(double)
func (this *QRandomGenerator) Bounded(highest float64) {
	// 0: (, highest double), (&highest)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &highest)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:110
// index:1
// inline
// quint32 bounded(quint32)
func (this *QRandomGenerator) Bounded_1(highest uint) {
	// 1: (, highest quint32), (&highest)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &highest)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:118
// index:2
// inline
// int bounded(int)
func (this *QRandomGenerator) Bounded_2(highest int) {
	// 2: (, highest int), (&highest)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &highest)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:123
// index:3
// inline
// quint32 bounded(quint32, quint32)
func (this *QRandomGenerator) Bounded_3(lowest uint, highest uint) {
	// 3: (, lowest quint32, highest quint32), (&lowest, &highest)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEjj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &lowest, &highest)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:128
// index:4
// inline
// int bounded(int, int)
func (this *QRandomGenerator) Bounded_4(lowest int, highest int) {
	// 4: (, lowest int, highest int), (&lowest, &highest)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &lowest, &highest)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:160
// index:0
// inline
// void seed(quint32)
func (this *QRandomGenerator) Seed(s uint) {
	// 0: (, s quint32), (&s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator4seedEj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:161
// index:1
// inline
// void seed(std::seed_seq &)
func (this *QRandomGenerator) Seed_1(sseq int) {
	// 1: (, sseq std::seed_seq &), (&sseq)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator4seedERSt8seed_seq", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sseq)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:162
// index:0
// void discard(unsigned long long)
func (this *QRandomGenerator) Discard(z uint64) {
	// 0: (, z unsigned long long), (&z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7discardEy", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:163
// index:0
// static inline
// QRandomGenerator::result_type min()
func (this *QRandomGenerator) Min() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator3minEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator_Min() {
	// 0: (), ()
	var nilthis *QRandomGenerator
	nilthis.Min()
}

// /usr/include/qt/QtCore/qrandom.h:164
// index:0
// static inline
// QRandomGenerator::result_type max()
func (this *QRandomGenerator) Max() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator3maxEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator_Max() {
	// 0: (), ()
	var nilthis *QRandomGenerator
	nilthis.Max()
}

// /usr/include/qt/QtCore/qrandom.h:166
// index:0
// static inline
// QRandomGenerator * system()
func (this *QRandomGenerator) System() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator6systemEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator_System() {
	// 0: (), ()
	var nilthis *QRandomGenerator
	nilthis.System()
}

// /usr/include/qt/QtCore/qrandom.h:167
// index:0
// static inline
// QRandomGenerator * global()
func (this *QRandomGenerator) Global() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator6globalEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator_Global() {
	// 0: (), ()
	var nilthis *QRandomGenerator
	nilthis.Global()
}

// /usr/include/qt/QtCore/qrandom.h:168
// index:0
// static inline
// QRandomGenerator securelySeeded()
func (this *QRandomGenerator) SecurelySeeded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator14securelySeededEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRandomGenerator_SecurelySeeded() {
	// 0: (), ()
	var nilthis *QRandomGenerator
	nilthis.SecurelySeeded()
}

//  body block end
