package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
// size 16
type QVariant struct {
	*qtrt.CObject
}
type QVariant_ITF interface {
	QVariant_PTR() *QVariant
}

func (ptr *QVariant) QVariant_PTR() *QVariant { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QVariantFromptr(cthis Voidptr) *QVariant {
	return &QVariant{&qtrt.CObject{cthis}}
}
func (*QVariant) Fromptr(cthis Voidptr) *QVariant {
	return QVariantFromptr(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVariant()

/*
 */
func (*QVariant) NewForInherit() *QVariant {
	return NewQVariant()
}
func NewQVariant() *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(93067653, "_ZN8QVariantC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:210
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QVariant(QVariant::Type)

/*
 */
func (*QVariant) NewForInherit1(type_ int) *QVariant {
	return NewQVariant1(type_)
}
func NewQVariant1(type_ int) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3094501539, "_ZN8QVariantC2ENS_4TypeE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&type_))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:211
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int, const void *)

/*
 */
func (*QVariant) NewForInherit2(typeId int, copy unsafe.Pointer /*666*/) *QVariant {
	return NewQVariant2(typeId, copy)
}
func NewQVariant2(typeId int, copy unsafe.Pointer /*666*/) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(2845593041, "_ZN8QVariantC2EiPKv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&typeId), Voidptr(&copy))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:212
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int, const void *, uint)

/*
 */
func (*QVariant) NewForInherit3(typeId int, copy unsafe.Pointer /*666*/, flags uint) *QVariant {
	return NewQVariant3(typeId, copy, flags)
}
func NewQVariant3(typeId int, copy unsafe.Pointer /*666*/, flags uint) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(2383944132, "_ZN8QVariantC2EiPKvj", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&typeId), Voidptr(&copy), Voidptr(&flags))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:219
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int)

/*
 */
func (*QVariant) NewForInherit4(i int) *QVariant {
	return NewQVariant4(i)
}
func NewQVariant4(i int) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(2290357360, "_ZN8QVariantC2Ei", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&cthis), Voidptr(&i))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:220
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVariant(uint)

/*
 */
func (*QVariant) NewForInherit5(ui uint) *QVariant {
	return NewQVariant5(ui)
}
func NewQVariant5(ui uint) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(294471114, "_ZN8QVariantC2Ej", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&ui))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:221
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVariant(qlonglong)

/*
 */
func (*QVariant) NewForInherit6(ll int64) *QVariant {
	return NewQVariant6(ll)
}
func NewQVariant6(ll int64) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3795072130, "_ZN8QVariantC2Ex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&ll))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:222
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QVariant(qulonglong)

/*
 */
func (*QVariant) NewForInherit7(ull uint64) *QVariant {
	return NewQVariant7(ull)
}
func NewQVariant7(ull uint64) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(2503148564, "_ZN8QVariantC2Ey", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&ull))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:223
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QVariant(bool)

/*
 */
func (*QVariant) NewForInherit8(b bool) *QVariant {
	return NewQVariant8(b)
}
func NewQVariant8(b bool) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(525782520, "_ZN8QVariantC2Eb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&b))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:224
// index:9
// Public Visibility=Default Availability=Available
// [-2] void QVariant(double)

/*
 */
func (*QVariant) NewForInherit9(d float64) *QVariant {
	return NewQVariant9(d)
}
func NewQVariant9(d float64) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(4130695373, "_ZN8QVariantC2Ed", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_DOUBLE, Voidptr(&cthis), Voidptr(&d))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:225
// index:10
// Public Visibility=Default Availability=Available
// [-2] void QVariant(float)

/*
 */
func (*QVariant) NewForInherit10(f float32) *QVariant {
	return NewQVariant10(f)
}
func NewQVariant10(f float32) *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(406522337, "_ZN8QVariantC2Ef", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_FLOAT, Voidptr(&cthis), Voidptr(&f))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:227
// index:11
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const char *)

/*
 */
func (*QVariant) NewForInherit11(str string) *QVariant {
	return NewQVariant11(str)
}
func NewQVariant11(str string) *QVariant {
	var convArg0 = qtrt.CStringRef(&str)
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(2561691587, "_ZN8QVariantC2EPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:283
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QVariant::Type type() const

/*
 */
func (this *QVariant) Type() int {
	rv, err := qtrt.Qtcc3(1071828666, "_ZNK8QVariant4typeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int()
}

// /usr/include/qt/QtCore/qvariant.h:284
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int userType() const

/*
 */
func (this *QVariant) UserType() int {
	rv, err := qtrt.Qtcc3(2949716066, "_ZNK8QVariant8userTypeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:285
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] const char * typeName() const

/*
 */
func (this *QVariant) TypeName() string {
	rv, err := qtrt.Qtcc3(330372657, "_ZNK8QVariant8typeNameEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return qtrt.GoStringI(rv.Uint64())
}

// /usr/include/qt/QtCore/qvariant.h:290
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isValid() const

/*
 */
func (this *QVariant) IsValid() bool {
	rv, err := qtrt.Qtcc3(2204930347, "_ZNK8QVariant7isValidEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qvariant.h:291
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isNull() const

/*
 */
func (this *QVariant) IsNull() bool {
	rv, err := qtrt.Qtcc3(3898904494, "_ZNK8QVariant6isNullEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qvariant.h:293
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear()

/*
 */
func (this *QVariant) Clear() {
	rv, err := qtrt.Qtcc3(126984843, "_ZN8QVariant5clearEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:295
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void detach()

/*
 */
func (this *QVariant) Detach() {
	rv, err := qtrt.Qtcc3(1393086918, "_ZN8QVariant6detachEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:296
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isDetached() const

/*
 */
func (this *QVariant) IsDetached() bool {
	rv, err := qtrt.Qtcc3(3109827116, "_ZNK8QVariant10isDetachedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qvariant.h:298
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int toInt(bool *) const

/*
 */
func (this *QVariant) ToInt(ok *bool) int {
	rv, err := qtrt.Qtcc3(323143542, "_ZNK8QVariant5toIntEPb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:298
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int toInt(bool *) const

/*
 */
func (this *QVariant) ToIntp() int {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok Voidptr
	rv, err := qtrt.Qtcc3(323143542, "_ZNK8QVariant5toIntEPb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:299
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] uint toUInt(bool *) const

/*
 */
func (this *QVariant) ToUInt(ok *bool) uint {
	rv, err := qtrt.Qtcc3(208741389, "_ZNK8QVariant6toUIntEPb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Uint() // 222
}

// /usr/include/qt/QtCore/qvariant.h:299
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] uint toUInt(bool *) const

/*
 */
func (this *QVariant) ToUIntp() uint {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok Voidptr
	rv, err := qtrt.Qtcc3(208741389, "_ZNK8QVariant6toUIntEPb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Uint() // 222
}

// /usr/include/qt/QtCore/qvariant.h:300
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *) const

/*
 */
func (this *QVariant) ToLongLong(ok *bool) int64 {
	rv, err := qtrt.Qtcc3(1775070467, "_ZNK8QVariant10toLongLongEPb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Int64() // 222
}

// /usr/include/qt/QtCore/qvariant.h:300
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *) const

/*
 */
func (this *QVariant) ToLongLongp() int64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok Voidptr
	rv, err := qtrt.Qtcc3(1775070467, "_ZNK8QVariant10toLongLongEPb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Int64() // 222
}

// /usr/include/qt/QtCore/qvariant.h:301
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *) const

/*
 */
func (this *QVariant) ToULongLong(ok *bool) uint64 {
	rv, err := qtrt.Qtcc3(4247319954, "_ZNK8QVariant11toULongLongEPb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Uint64() // 222
}

// /usr/include/qt/QtCore/qvariant.h:301
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *) const

/*
 */
func (this *QVariant) ToULongLongp() uint64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok Voidptr
	rv, err := qtrt.Qtcc3(4247319954, "_ZNK8QVariant11toULongLongEPb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Uint64() // 222
}

// /usr/include/qt/QtCore/qvariant.h:302
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool toBool() const

/*
 */
func (this *QVariant) ToBool() bool {
	rv, err := qtrt.Qtcc3(3766730072, "_ZNK8QVariant6toBoolEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qvariant.h:303
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
 */
func (this *QVariant) ToDouble(ok *bool) float64 {
	rv, err := qtrt.Qtcc3(982518782, "_ZNK8QVariant8toDoubleEPb", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Float64() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:303
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
 */
func (this *QVariant) ToDoublep() float64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok Voidptr
	rv, err := qtrt.Qtcc3(982518782, "_ZNK8QVariant8toDoubleEPb", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Float64() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:304
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
 */
func (this *QVariant) ToFloat(ok *bool) float32 {
	rv, err := qtrt.Qtcc3(2104130602, "_ZNK8QVariant7toFloatEPb", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Float32() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:304
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
 */
func (this *QVariant) ToFloatp() float32 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok Voidptr
	rv, err := qtrt.Qtcc3(2104130602, "_ZNK8QVariant7toFloatEPb", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Float32() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:305
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qreal toReal(bool *) const

/*
 */
func (this *QVariant) ToReal(ok *bool) float64 {
	rv, err := qtrt.Qtcc3(2923771759, "_ZNK8QVariant6toRealEPb", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Float64() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:305
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qreal toReal(bool *) const

/*
 */
func (this *QVariant) ToRealp() float64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok Voidptr
	rv, err := qtrt.Qtcc3(2923771759, "_ZNK8QVariant6toRealEPb", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ok))
	qtrt.ErrPrint3(err, rv)
	return rv.Float64() // 1111
}

// /usr/include/qt/QtCore/qvariant.h:319
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QPoint toPoint() const

/*
 */
func (this *QVariant) ToPoint() *QPoint /*123*/ {
	rv, err := qtrt.Qtcc3(3845935258, "_ZNK8QVariant7toPointEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(8)
	qtrt.Cmemcpy(cthis, rv.Addr(), 8)
	rv2 := /*==*/ QPointFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:320
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QPointF toPointF() const

/*
 */
func (this *QVariant) ToPointF() *QPointF /*123*/ {
	rv, err := qtrt.Qtcc3(512643617, "_ZNK8QVariant8toPointFEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(16)
	qtrt.Cmemcpy(cthis, rv.Addr(), 16)
	rv2 := /*==*/ QPointFFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:321
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QRect toRect() const

/*
 */
func (this *QVariant) ToRect() *QRect /*123*/ {
	rv, err := qtrt.Qtcc3(4043762194, "_ZNK8QVariant6toRectEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(16)
	qtrt.Cmemcpy(cthis, rv.Addr(), 16)
	rv2 := /*==*/ QRectFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:322
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QSize toSize() const

/*
 */
func (this *QVariant) ToSize() *QSize /*123*/ {
	rv, err := qtrt.Qtcc3(3479111140, "_ZNK8QVariant6toSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(8)
	qtrt.Cmemcpy(cthis, rv.Addr(), 8)
	rv2 := /*==*/ QSizeFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:323
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QSizeF toSizeF() const

/*
 */
func (this *QVariant) ToSizeF() *QSizeF /*123*/ {
	rv, err := qtrt.Qtcc3(3575902903, "_ZNK8QVariant7toSizeFEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(16)
	qtrt.Cmemcpy(cthis, rv.Addr(), 16)
	rv2 := /*==*/ QSizeFFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:324
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QLine toLine() const

/*
 */
func (this *QVariant) ToLine() *QLine /*123*/ {
	rv, err := qtrt.Qtcc3(3802428770, "_ZNK8QVariant6toLineEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(16)
	qtrt.Cmemcpy(cthis, rv.Addr(), 16)
	rv2 := /*==*/ QLineFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLine)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:325
// index:0
// Public Indirect Visibility=Default Availability=Available
// [32] QLineF toLineF() const

/*
 */
func (this *QVariant) ToLineF() *QLineF /*123*/ {
	sretobj := qtrt.Malloc(32) // QLineF
	rv, err := qtrt.Qtcc3(3520409554, "_ZNK8QVariant7toLineFEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QLineFFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:326
// index:0
// Public Indirect Visibility=Default Availability=Available
// [32] QRectF toRectF() const

/*
 */
func (this *QVariant) ToRectF() *QRectF /*123*/ {
	sretobj := qtrt.Malloc(32) // QRectF
	rv, err := qtrt.Qtcc3(2177105009, "_ZNK8QVariant7toRectFEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QRectFFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

func DeleteQVariant(this *QVariant) {
	rv, err := qtrt.Qtcc3(2556109116, "_ZN8QVariantD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QVariant__Type = int

//
const QVariant__Invalid QVariant__Type = 0

//
const QVariant__Bool QVariant__Type = 1

//
const QVariant__Int QVariant__Type = 2

//
const QVariant__UInt QVariant__Type = 3

//
const QVariant__LongLong QVariant__Type = 4

//
const QVariant__ULongLong QVariant__Type = 5

//
const QVariant__Double QVariant__Type = 6

//
const QVariant__Char QVariant__Type = 7

//
const QVariant__Map QVariant__Type = 8

//
const QVariant__List QVariant__Type = 9

//
const QVariant__String QVariant__Type = 10

//
const QVariant__StringList QVariant__Type = 11

//
const QVariant__ByteArray QVariant__Type = 12

//
const QVariant__BitArray QVariant__Type = 13

//
const QVariant__Date QVariant__Type = 14

//
const QVariant__Time QVariant__Type = 15

//
const QVariant__DateTime QVariant__Type = 16

//
const QVariant__Url QVariant__Type = 17

//
const QVariant__Locale QVariant__Type = 18

//
const QVariant__Rect QVariant__Type = 19

//
const QVariant__RectF QVariant__Type = 20

//
const QVariant__Size QVariant__Type = 21

//
const QVariant__SizeF QVariant__Type = 22

//
const QVariant__Line QVariant__Type = 23

//
const QVariant__LineF QVariant__Type = 24

//
const QVariant__Point QVariant__Type = 25

//
const QVariant__PointF QVariant__Type = 26

//
const QVariant__RegExp QVariant__Type = 27

//
const QVariant__RegularExpression QVariant__Type = 44

//
const QVariant__Hash QVariant__Type = 28

//
const QVariant__EasingCurve QVariant__Type = 29

//
const QVariant__Uuid QVariant__Type = 30

//
const QVariant__ModelIndex QVariant__Type = 42

//
const QVariant__PersistentModelIndex QVariant__Type = 50

//
const QVariant__LastCoreType QVariant__Type = 55

//
const QVariant__Font QVariant__Type = 64

//
const QVariant__Pixmap QVariant__Type = 65

//
const QVariant__Brush QVariant__Type = 66

//
const QVariant__Color QVariant__Type = 67

//
const QVariant__Palette QVariant__Type = 68

//
const QVariant__Image QVariant__Type = 70

//
const QVariant__Polygon QVariant__Type = 71

//
const QVariant__Region QVariant__Type = 72

//
const QVariant__Bitmap QVariant__Type = 73

//
const QVariant__Cursor QVariant__Type = 74

//
const QVariant__KeySequence QVariant__Type = 75

//
const QVariant__Pen QVariant__Type = 76

//
const QVariant__TextLength QVariant__Type = 77

//
const QVariant__TextFormat QVariant__Type = 78

//
const QVariant__Matrix QVariant__Type = 79

//
const QVariant__Transform QVariant__Type = 80

//
const QVariant__Matrix4x4 QVariant__Type = 81

//
const QVariant__Vector2D QVariant__Type = 82

//
const QVariant__Vector3D QVariant__Type = 83

//
const QVariant__Vector4D QVariant__Type = 84

//
const QVariant__Quaternion QVariant__Type = 85

//
const QVariant__PolygonF QVariant__Type = 86

//
const QVariant__Icon QVariant__Type = 69

//
const QVariant__LastGuiType QVariant__Type = 87

//
const QVariant__SizePolicy QVariant__Type = 121

//
const QVariant__UserType QVariant__Type = 1024

//
const QVariant__LastType QVariant__Type = -1

func (this *QVariant) TypeItemName(val int) string {
	switch val {
	case QVariant__Invalid: // 0
		return "Invalid"
	case QVariant__Bool: // 1
		return "Bool"
	case QVariant__Int: // 2
		return "Int"
	case QVariant__UInt: // 3
		return "UInt"
	case QVariant__LongLong: // 4
		return "LongLong"
	case QVariant__ULongLong: // 5
		return "ULongLong"
	case QVariant__Double: // 6
		return "Double"
	case QVariant__Char: // 7
		return "Char"
	case QVariant__Map: // 8
		return "Map"
	case QVariant__List: // 9
		return "List"
	case QVariant__String: // 10
		return "String"
	case QVariant__StringList: // 11
		return "StringList"
	case QVariant__ByteArray: // 12
		return "ByteArray"
	case QVariant__BitArray: // 13
		return "BitArray"
	case QVariant__Date: // 14
		return "Date"
	case QVariant__Time: // 15
		return "Time"
	case QVariant__DateTime: // 16
		return "DateTime"
	case QVariant__Url: // 17
		return "Url"
	case QVariant__Locale: // 18
		return "Locale"
	case QVariant__Rect: // 19
		return "Rect"
	case QVariant__RectF: // 20
		return "RectF"
	case QVariant__Size: // 21
		return "Size"
	case QVariant__SizeF: // 22
		return "SizeF"
	case QVariant__Line: // 23
		return "Line"
	case QVariant__LineF: // 24
		return "LineF"
	case QVariant__Point: // 25
		return "Point"
	case QVariant__PointF: // 26
		return "PointF"
	case QVariant__RegExp: // 27
		return "RegExp"
	case QVariant__RegularExpression: // 44
		return "RegularExpression"
	case QVariant__Hash: // 28
		return "Hash"
	case QVariant__EasingCurve: // 29
		return "EasingCurve"
	case QVariant__Uuid: // 30
		return "Uuid"
	case QVariant__ModelIndex: // 42
		return "ModelIndex"
	case QVariant__PersistentModelIndex: // 50
		return "PersistentModelIndex"
	case QVariant__LastCoreType: // 55
		return "LastCoreType"
	case QVariant__Font: // 64
		return "Font"
	case QVariant__Pixmap: // 65
		return "Pixmap"
	case QVariant__Brush: // 66
		return "Brush"
	case QVariant__Color: // 67
		return "Color"
	case QVariant__Palette: // 68
		return "Palette"
	case QVariant__Image: // 70
		return "Image"
	case QVariant__Polygon: // 71
		return "Polygon"
	case QVariant__Region: // 72
		return "Region"
	case QVariant__Bitmap: // 73
		return "Bitmap"
	case QVariant__Cursor: // 74
		return "Cursor"
	case QVariant__KeySequence: // 75
		return "KeySequence"
	case QVariant__Pen: // 76
		return "Pen"
	case QVariant__TextLength: // 77
		return "TextLength"
	case QVariant__TextFormat: // 78
		return "TextFormat"
	case QVariant__Matrix: // 79
		return "Matrix"
	case QVariant__Transform: // 80
		return "Transform"
	case QVariant__Matrix4x4: // 81
		return "Matrix4x4"
	case QVariant__Vector2D: // 82
		return "Vector2D"
	case QVariant__Vector3D: // 83
		return "Vector3D"
	case QVariant__Vector4D: // 84
		return "Vector4D"
	case QVariant__Quaternion: // 85
		return "Quaternion"
	case QVariant__PolygonF: // 86
		return "PolygonF"
	case QVariant__Icon: // 69
		return "Icon"
	case QVariant__LastGuiType: // 87
		return "LastGuiType"
	case QVariant__SizePolicy: // 121
		return "SizePolicy"
	case QVariant__UserType: // 1024
		return "UserType"
	case QVariant__LastType: // -1
		return "LastType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QVariant_TypeItemName(val int) string {
	var nilthis *QVariant
	return nilthis.TypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10013() {
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
