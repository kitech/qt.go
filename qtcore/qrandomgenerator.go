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
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QRandomGenerator struct {
	*qtrt.CObject
}
type QRandomGenerator_ITF interface {
	QRandomGenerator_PTR() *QRandomGenerator
}

func (ptr *QRandomGenerator) QRandomGenerator_PTR() *QRandomGenerator { return ptr }

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

// /usr/include/qt/QtCore/qrandom.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator(quint32)

/*

 */
func NewQRandomGenerator(seedValue uint) *QRandomGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2Ej", qtrt.FFI_TYPE_POINTER, seedValue)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator(quint32)

/*

 */
func NewQRandomGenerator__() *QRandomGenerator {
	// arg: 0, quint32=Typedef, quint32=Typedef, unsigned int, UInt
	seedValue := uint(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2Ej", qtrt.FFI_TYPE_POINTER, seedValue)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:68
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QRandomGenerator(const quint32 *, qsizetype)

/*

 */
func NewQRandomGenerator_1(seedBuffer unsafe.Pointer /*666*/, len_ int64) *QRandomGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2EPKjx", qtrt.FFI_TYPE_POINTER, seedBuffer, len_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:72
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRandomGenerator(const quint32 *, const quint32 *)

/*

 */
func NewQRandomGenerator_2(begin unsafe.Pointer /*666*/, end unsafe.Pointer /*666*/) *QRandomGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2EPKjS1_", qtrt.FFI_TYPE_POINTER, begin, end)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:179
// index:3
// Protected Visibility=Default Availability=Available
// [-2] void QRandomGenerator(QRandomGenerator::System)

/*

 */
func NewQRandomGenerator_3(arg0 int) *QRandomGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorC2ENS_6SystemE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRandomGenerator)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:76
// index:0
// Public Visibility=Default Availability=Available
// [2512] QRandomGenerator & operator=(const QRandomGenerator &)

/*

 */
func (this *QRandomGenerator) Operator_equal(other QRandomGenerator_ITF) *QRandomGenerator {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRandomGenerator_PTR() != nil {
		convArg0 = other.QRandomGenerator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratoraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRandomGenerator)
	return rv2
}

// /usr/include/qt/QtCore/qrandom.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [4] quint32 generate()

/*

 */
func (this *QRandomGenerator) Generate() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator8generateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:159
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void generate(quint32 *, quint32 *)

/*

 */
func (this *QRandomGenerator) Generate_1(begin unsafe.Pointer /*666*/, end unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator8generateEPjS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), begin, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] quint64 generate64()

/*

 */
func (this *QRandomGenerator) Generate64() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator10generate64Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double generateDouble()

/*

 */
func (this *QRandomGenerator) GenerateDouble() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator14generateDoubleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrandom.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [8] double bounded(double)

/*

 */
func (this *QRandomGenerator) Bounded(highest float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), highest)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qrandom.h:117
// index:1
// Public inline Visibility=Default Availability=Available
// [4] quint32 bounded(quint32)

/*

 */
func (this *QRandomGenerator) Bounded_1(highest uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), highest)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:125
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int bounded(int)

/*

 */
func (this *QRandomGenerator) Bounded_2(highest int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), highest)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrandom.h:130
// index:3
// Public inline Visibility=Default Availability=Available
// [4] quint32 bounded(quint32, quint32)

/*

 */
func (this *QRandomGenerator) Bounded_3(lowest uint, highest uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEjj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), lowest, highest)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:135
// index:4
// Public inline Visibility=Default Availability=Available
// [4] int bounded(int, int)

/*

 */
func (this *QRandomGenerator) Bounded_4(lowest int, highest int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), lowest, highest)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qrandom.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QRandomGenerator::result_type operator()()

/*

 */
func (this *QRandomGenerator) Operator_fncall() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGeneratorclEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:167
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void seed(quint32)

/*

 */
func (this *QRandomGenerator) Seed(s uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator4seedEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), s)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:167
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void seed(quint32)

/*

 */
func (this *QRandomGenerator) Seed__() {
	// arg: 0, quint32=Typedef, quint32=Typedef, unsigned int, UInt
	s := uint(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator4seedEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), s)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void discard(unsigned long long)

/*

 */
func (this *QRandomGenerator) Discard(z uint64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator7discardEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:170
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] QRandomGenerator::result_type min()

/*

 */
func (this *QRandomGenerator) Min() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator3minEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QRandomGenerator_Min() uint {
	var nilthis *QRandomGenerator
	rv := nilthis.Min()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:171
// index:0
// Public static inline Visibility=Default Availability=Available
// [4] QRandomGenerator::result_type max()

/*

 */
func (this *QRandomGenerator) Max() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator3maxEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QRandomGenerator_Max() uint {
	var nilthis *QRandomGenerator
	rv := nilthis.Max()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:173
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRandomGenerator * system()

/*

 */
func (this *QRandomGenerator) System() *QRandomGenerator /*777 QRandomGenerator **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator6systemEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QRandomGenerator_System() *QRandomGenerator /*777 QRandomGenerator **/ {
	var nilthis *QRandomGenerator
	rv := nilthis.System()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:174
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QRandomGenerator * global()

/*

 */
func (this *QRandomGenerator) Global() *QRandomGenerator /*777 QRandomGenerator **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator6globalEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QRandomGenerator_Global() *QRandomGenerator /*777 QRandomGenerator **/ {
	var nilthis *QRandomGenerator
	rv := nilthis.Global()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:175
// index:0
// Public static inline Visibility=Default Availability=Available
// [2512] QRandomGenerator securelySeeded()

/*

 */
func (this *QRandomGenerator) SecurelySeeded() *QRandomGenerator /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QRandomGenerator14securelySeededEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QRandomGenerator__System = int

func (this *QRandomGenerator) SystemItemName(val int) string {
	switch val {
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QRandomGenerator_SystemItemName(val int) string {
	var nilthis *QRandomGenerator
	return nilthis.SystemItemName(val)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
