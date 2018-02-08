package qtcore

// /usr/include/qt/QtCore/qrandom.h
// #include <qrandom.h>
// #include <QtCore>

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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QRandomGenerator struct {
	*qtrt.CObject
}

func (this *QRandomGenerator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRandomGenerator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRandomGeneratorFromPointer(cthis unsafe.Pointer) *QRandomGenerator {
	return &QRandomGenerator{&qtrt.CObject{cthis}}
}
func (*QRandomGenerator) NewFromPointer(cthis unsafe.Pointer) *QRandomGenerator {
	return NewQRandomGeneratorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrandom.h:55
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator(quint32)
func NewQRandomGenerator(seedValue uint) *QRandomGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2Ej", qtrt.FFI_TYPE_POINTER, seedValue)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator(const quint32 *, qsizetype)
func NewQRandomGenerator_1(seedBuffer unsafe.Pointer /*666*/, len int64) *QRandomGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2EPKjx", qtrt.FFI_TYPE_POINTER, &seedBuffer, len)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:65
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRandomGenerator(const quint32 *, const quint32 *)
func NewQRandomGenerator_2(begin unsafe.Pointer /*666*/, end unsafe.Pointer /*666*/) *QRandomGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2EPKjS1_", qtrt.FFI_TYPE_POINTER, &begin, &end)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:172
// index:3
// Protected Visibility=Default Availability=Available
// [-2] void QRandomGenerator(enum QRandomGenerator::System)
func NewQRandomGenerator_3(arg0 int) *QRandomGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2ENS_6SystemE", qtrt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [4] quint32 generate()
func (this *QRandomGenerator) Generate() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator8generateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:152
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void generate(quint32 *, quint32 *)
func (this *QRandomGenerator) Generate_1(begin unsafe.Pointer /*666*/, end unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator8generateEPjS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &begin, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [8] quint64 generate64()
func (this *QRandomGenerator) Generate64() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator10generate64Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double generateDouble()
func (this *QRandomGenerator) GenerateDouble() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator14generateDoubleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrandom.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double bounded(double)
func (this *QRandomGenerator) Bounded(highest float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), highest)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrandom.h:110
// index:1
// Public inline Visibility=Default Availability=Available
// [4] quint32 bounded(quint32)
func (this *QRandomGenerator) Bounded_1(highest uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), highest)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:118
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int bounded(int)
func (this *QRandomGenerator) Bounded_2(highest int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), highest)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrandom.h:123
// index:3
// Public inline Visibility=Default Availability=Available
// [4] quint32 bounded(quint32, quint32)
func (this *QRandomGenerator) Bounded_3(lowest uint, highest uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEjj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), lowest, highest)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:128
// index:4
// Public inline Visibility=Default Availability=Available
// [4] int bounded(int, int)
func (this *QRandomGenerator) Bounded_4(lowest int, highest int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), lowest, highest)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrandom.h:160
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void seed(quint32)
func (this *QRandomGenerator) Seed(s uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator4seedEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void discard(unsigned long long)
func (this *QRandomGenerator) Discard(z uint64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7discardEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:163
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] QRandomGenerator::result_type min()
func (this *QRandomGenerator) Min() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator3minEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QRandomGenerator_Min() uint {
	var nilthis *QRandomGenerator
	rv := nilthis.Min()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:164
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] QRandomGenerator::result_type max()
func (this *QRandomGenerator) Max() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator3maxEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QRandomGenerator_Max() uint {
	var nilthis *QRandomGenerator
	rv := nilthis.Max()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:166
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRandomGenerator * system()
func (this *QRandomGenerator) System() *QRandomGenerator /*777 QRandomGenerator **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator6systemEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QRandomGenerator_System() *QRandomGenerator /*777 QRandomGenerator **/ {
	var nilthis *QRandomGenerator
	rv := nilthis.System()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:167
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRandomGenerator * global()
func (this *QRandomGenerator) Global() *QRandomGenerator /*777 QRandomGenerator **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator6globalEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QRandomGenerator_Global() *QRandomGenerator /*777 QRandomGenerator **/ {
	var nilthis *QRandomGenerator
	rv := nilthis.Global()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:168
// index:0
// Public static inline Visibility=Default Availability=Available
// [2512] QRandomGenerator securelySeeded()
func (this *QRandomGenerator) SecurelySeeded() *QRandomGenerator /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator14securelySeededEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRandomGenerator)
	return rv2
}
func QRandomGenerator_SecurelySeeded() *QRandomGenerator /*123*/ {
	var nilthis *QRandomGenerator
	rv := nilthis.SecurelySeeded()
	return rv
}

func DeleteQRandomGenerator(this *QRandomGenerator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QRandomGenerator__System = int

//  body block end
