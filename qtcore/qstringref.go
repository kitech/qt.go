package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QStringRef struct {
	*qtrt.CObject
}
type QStringRef_ITF interface {
	QStringRef_PTR() *QStringRef
}

func (ptr *QStringRef) QStringRef_PTR() *QStringRef { return ptr }

func (this *QStringRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringRefFromPointer(cthis unsafe.Pointer) *QStringRef {
	return &QStringRef{&qtrt.CObject{cthis}}
}
func (*QStringRef) NewFromPointer(cthis unsafe.Pointer) *QStringRef {
	return NewQStringRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstring.h:1420
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QStringRef()
func NewQStringRef() *QStringRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1421
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QStringRef(const QString *, int, int)
func NewQStringRef_1(string string, position int, size int) *QStringRef {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QStringii", qtrt.FFI_TYPE_POINTER, convArg0, position, size)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1422
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QStringRef(const QString *)
func NewQStringRef_2(string string) *QStringRef {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1431
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef & operator=(QStringRef &&)
func (this *QStringRef) Operator_equal(other unsafe.Pointer /*333*/) *QStringRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1433
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QStringRef & operator=(const QStringRef &)
func (this *QStringRef) Operator_equal_1(other QStringRef_ITF) *QStringRef {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStringRef_PTR() != nil {
		convArg0 = other.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1505
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QStringRef & operator=(const QString *)
func (this *QStringRef) Operator_equal_2(string string) *QStringRef {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefaSEPK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1438
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QStringRef()
func DeleteQStringRef(this *QStringRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstring.h:1441
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QString * string() const
func (this *QStringRef) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:1442
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int position() const
func (this *QStringRef) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1443
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const
func (this *QStringRef) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1444
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count() const
func (this *QStringRef) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1461
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) Count_1(s string, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1461
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) Count_1_(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1462
// index:2
// Public Visibility=Default Availability=Available
// [4] int count(QChar, Qt::CaseSensitivity) const
func (this *QStringRef) Count_2(c QChar_ITF /*123*/, cs int) int {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1462
// index:2
// Public Visibility=Default Availability=Available
// [4] int count(QChar, Qt::CaseSensitivity) const
func (this *QStringRef) Count_2_(c QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1463
// index:3
// Public Visibility=Default Availability=Available
// [4] int count(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) Count_3(s QStringRef_ITF, cs int) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1463
// index:3
// Public Visibility=Default Availability=Available
// [4] int count(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) Count_3_(s QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1445
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length() const
func (this *QStringRef) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf(str string, from int, cs int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf__(str string) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	from := 0
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf__1(str string, from int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(QChar, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_1(ch QChar_ITF /*123*/, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(QChar, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_1_(ch QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := 0
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(QChar, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_1_1(ch QChar_ITF /*123*/, from int) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(QLatin1String, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_2(str QLatin1String_ITF /*123*/, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(QLatin1String, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_2_(str QLatin1String_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := 0
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(QLatin1String, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_2_1(str QLatin1String_ITF /*123*/, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QStringRef &, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_3(str QStringRef_ITF, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QStringRef &, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_3_(str QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := 0
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QStringRef &, int, Qt::CaseSensitivity) const
func (this *QStringRef) IndexOf_3_1(str QStringRef_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf(str string, from int, cs int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf__(str string) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	from := -1
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf__1(str string, from int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QChar, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_1(ch QChar_ITF /*123*/, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QChar, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_1_(ch QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := -1
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QChar, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_1_1(ch QChar_ITF /*123*/, from int) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QLatin1String, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_2(str QLatin1String_ITF /*123*/, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QLatin1String, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_2_(str QLatin1String_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := -1
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QLatin1String, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_2_1(str QLatin1String_ITF /*123*/, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QStringRef &, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_3(str QStringRef_ITF, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QStringRef &, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_3_(str QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := -1
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QStringRef &, int, Qt::CaseSensitivity) const
func (this *QStringRef) LastIndexOf_3_1(str QStringRef_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1456
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) Contains(str string, cs int) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1456
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) Contains__(str string) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1457
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QChar, Qt::CaseSensitivity) const
func (this *QStringRef) Contains_1(ch QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1457
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QChar, Qt::CaseSensitivity) const
func (this *QStringRef) Contains_1_(ch QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1458
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringRef) Contains_2(str QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1458
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringRef) Contains_2_(str QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1459
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) Contains_3(str QStringRef_ITF, cs int) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1459
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) Contains_3_(str QStringRef_ITF) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1470
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef left(int) const
func (this *QStringRef) Left(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4leftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1471
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef right(int) const
func (this *QStringRef) Right(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5rightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1472
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef mid(int, int) const
func (this *QStringRef) Mid(pos int, n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1472
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef mid(int, int) const
func (this *QStringRef) Mid__(pos int) *QStringRef /*123*/ {
	// arg: 1, int=Int, =Invalid,
	n := -1
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1473
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef chopped(int) const
func (this *QStringRef) Chopped(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7choppedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1476
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void truncate(int)
func (this *QStringRef) Truncate(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef8truncateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1477
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void chop(int)
func (this *QStringRef) Chop(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef4chopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1485
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRightToLeft() const
func (this *QStringRef) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef13isRightToLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1487
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1487
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith__(s QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1489
// index:1
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1489
// index:1
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith_1_(s QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1490
// index:2
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith_2(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1490
// index:2
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith_2_(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1492
// index:3
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith_3(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1492
// index:3
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith_3_(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1493
// index:4
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith_4(c QStringRef_ITF, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QStringRef_PTR() != nil {
		convArg0 = c.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1493
// index:4
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) StartsWith_4_(c QStringRef_ITF) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QStringRef_PTR() != nil {
		convArg0 = c.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1496
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1496
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith__(s QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1498
// index:1
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1498
// index:1
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith_1_(s QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1499
// index:2
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith_2(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1499
// index:2
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith_2_(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1501
// index:3
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith_3(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1501
// index:3
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith_3_(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1502
// index:4
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith_4(c QStringRef_ITF, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QStringRef_PTR() != nil {
		convArg0 = c.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1502
// index:4
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) EndsWith_4_(c QStringRef_ITF) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QStringRef_PTR() != nil {
		convArg0 = c.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1507
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * unicode() const
func (this *QStringRef) Unicode() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1513
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * data() const
func (this *QStringRef) Data() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1514
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * constData() const
func (this *QStringRef) ConstData() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1516
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator begin() const
func (this *QStringRef) Begin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1517
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator cbegin() const
func (this *QStringRef) Cbegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1518
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator constBegin() const
func (this *QStringRef) ConstBegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1519
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator end() const
func (this *QStringRef) End() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1520
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator cend() const
func (this *QStringRef) Cend() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1521
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator constEnd() const
func (this *QStringRef) ConstEnd() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1531
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toLatin1() const
func (this *QStringRef) ToLatin1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1532
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toUtf8() const
func (this *QStringRef) ToUtf8() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUtf8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1533
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toLocal8Bit() const
func (this *QStringRef) ToLocal8Bit() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toLocal8BitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1536
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clear()
func (this *QStringRef) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1537
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const
func (this *QStringRef) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:1538
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QStringRef) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1539
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QStringRef) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1541
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef appendTo(QString *) const
func (this *QStringRef) AppendTo(string string) *QStringRef /*123*/ {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8appendToEP7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1543
// index:0
// Public inline Visibility=Default Availability=Available
// [2] const QChar at(int) const
func (this *QStringRef) At(i int) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1545
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar operator[](int) const
func (this *QStringRef) Operator_get_index(i int) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1546
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar front() const
func (this *QStringRef) Front() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1547
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar back() const
func (this *QStringRef) Back() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1551
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const char *) const
func (this *QStringRef) Operator_equal_equal(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefeqEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1552
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const char *) const
func (this *QStringRef) Operator_not_equal(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefneEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1553
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const char *) const
func (this *QStringRef) Operator_less_than(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefltEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1554
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const char *) const
func (this *QStringRef) Operator_less_than_equal(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefleEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1555
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const char *) const
func (this *QStringRef) Operator_greater_than(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefgtEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1556
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const char *) const
func (this *QStringRef) Operator_greater_than_equal(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefgeEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1559
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) Compare(s string, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1559
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QString &, Qt::CaseSensitivity) const
func (this *QStringRef) Compare__(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1560
// index:1
// Public Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) Compare_1(s QStringRef_ITF, cs int) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1560
// index:1
// Public Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, Qt::CaseSensitivity) const
func (this *QStringRef) Compare_1_(s QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1561
// index:2
// Public Visibility=Default Availability=Available
// [4] int compare(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringRef) Compare_2(s QLatin1String_ITF /*123*/, cs int) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1561
// index:2
// Public Visibility=Default Availability=Available
// [4] int compare(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringRef) Compare_2_(s QLatin1String_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1563
// index:3
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QByteArray &, Qt::CaseSensitivity) const
func (this *QStringRef) Compare_3(s QByteArray_ITF, cs int) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK10QByteArrayN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1563
// index:3
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QByteArray &, Qt::CaseSensitivity) const
func (this *QStringRef) Compare_3_(s QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK10QByteArrayN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1566
// index:4
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QString &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_4(s1 QStringRef_ITF, s2 string, arg2 int) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_RK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_4(s1 QStringRef_ITF, s2 string, arg2 int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_4(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1566
// index:4
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QString &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_4_(s1 QStringRef_ITF, s2 string) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	arg2 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_RK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1568
// index:5
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_5(s1 QStringRef_ITF, s2 QStringRef_ITF, arg2 int) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QStringRef_PTR() != nil {
		convArg1 = s2.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_S1_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_5(s1 QStringRef_ITF, s2 QStringRef_ITF, arg2 int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_5(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1568
// index:5
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_5_(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QStringRef_PTR() != nil {
		convArg1 = s2.QStringRef_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	arg2 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_S1_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1570
// index:6
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Compare_6(s1 QStringRef_ITF, s2 QLatin1String_ITF /*123*/, cs int) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QLatin1String_PTR() != nil {
		convArg1 = s2.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_6(s1 QStringRef_ITF, s2 QLatin1String_ITF /*123*/, cs int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_6(s1, s2, cs)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1570
// index:6
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Compare_6_(s1 QStringRef_ITF, s2 QLatin1String_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QLatin1String_PTR() != nil {
		convArg1 = s2.QLatin1String_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1573
// index:0
// Public Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QString &) const
func (this *QStringRef) LocaleAwareCompare(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1574
// index:1
// Public Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &) const
func (this *QStringRef) LocaleAwareCompare_1(s QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1575
// index:2
// Public static Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &, const QString &)
func (this *QStringRef) LocaleAwareCompare_2(s1 QStringRef_ITF, s2 string) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_LocaleAwareCompare_2(s1 QStringRef_ITF, s2 string) int {
	var nilthis *QStringRef
	rv := nilthis.LocaleAwareCompare_2(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1576
// index:3
// Public static Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &, const QStringRef &)
func (this *QStringRef) LocaleAwareCompare_3(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QStringRef_PTR() != nil {
		convArg1 = s2.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_LocaleAwareCompare_3(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	var nilthis *QStringRef
	rv := nilthis.LocaleAwareCompare_3(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1578
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef trimmed() const
func (this *QStringRef) Trimmed() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(_Bool *, int) const
func (this *QStringRef) ToShort(ok *bool, base int) int16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(_Bool *, int) const
func (this *QStringRef) ToShort__() int16 {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(_Bool *, int) const
func (this *QStringRef) ToShort__1(ok *bool) int16 {
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(_Bool *, int) const
func (this *QStringRef) ToUShort(ok *bool, base int) uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(_Bool *, int) const
func (this *QStringRef) ToUShort__() uint16 {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(_Bool *, int) const
func (this *QStringRef) ToUShort__1(ok *bool) uint16 {
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(_Bool *, int) const
func (this *QStringRef) ToInt(ok *bool, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(_Bool *, int) const
func (this *QStringRef) ToInt__() int {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(_Bool *, int) const
func (this *QStringRef) ToInt__1(ok *bool) int {
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(_Bool *, int) const
func (this *QStringRef) ToUInt(ok *bool, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(_Bool *, int) const
func (this *QStringRef) ToUInt__() uint {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(_Bool *, int) const
func (this *QStringRef) ToUInt__1(ok *bool) uint {
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(_Bool *, int) const
func (this *QStringRef) ToLong(ok *bool, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(_Bool *, int) const
func (this *QStringRef) ToLong__() int {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(_Bool *, int) const
func (this *QStringRef) ToLong__1(ok *bool) int {
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(_Bool *, int) const
func (this *QStringRef) ToULong(ok *bool, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(_Bool *, int) const
func (this *QStringRef) ToULong__() uint {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(_Bool *, int) const
func (this *QStringRef) ToULong__1(ok *bool) uint {
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(_Bool *, int) const
func (this *QStringRef) ToLongLong(ok *bool, base int) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(_Bool *, int) const
func (this *QStringRef) ToLongLong__() int64 {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(_Bool *, int) const
func (this *QStringRef) ToLongLong__1(ok *bool) int64 {
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(_Bool *, int) const
func (this *QStringRef) ToULongLong(ok *bool, base int) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(_Bool *, int) const
func (this *QStringRef) ToULongLong__() uint64 {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(_Bool *, int) const
func (this *QStringRef) ToULongLong__1(ok *bool) uint64 {
	// arg: 1, int=Int, =Invalid,
	base := 10
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1587
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(_Bool *) const
func (this *QStringRef) ToFloat(ok *bool) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1587
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(_Bool *) const
func (this *QStringRef) ToFloat__() float32 {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1588
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(_Bool *) const
func (this *QStringRef) ToDouble(ok *bool) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1588
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(_Bool *) const
func (this *QStringRef) ToDouble__() float64 {
	// arg: 0, bool *=Pointer, =Invalid,
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
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
