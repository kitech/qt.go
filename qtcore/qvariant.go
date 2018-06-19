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
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void create(int, const void *)
func (this *QVariant) InheritCreate(f func(type_ int, copy unsafe.Pointer /*666*/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "create", f)
}

// bool cmp(const QVariant &)
func (this *QVariant) InheritCmp(f func(other *QVariant) bool) {
	qtrt.SetAllInheritCallback(this, "cmp", f)
}

// int compare(const QVariant &)
func (this *QVariant) InheritCompare(f func(other *QVariant) int) {
	qtrt.SetAllInheritCallback(this, "compare", f)
}

// bool convert(const int, void *)
func (this *QVariant) InheritConvert(f func(t int, ptr unsafe.Pointer /*666*/) bool) {
	qtrt.SetAllInheritCallback(this, "convert", f)
}

/*

 */
type QVariant struct {
	*qtrt.CObject
}
type QVariant_ITF interface {
	QVariant_PTR() *QVariant
}

func (ptr *QVariant) QVariant_PTR() *QVariant { return ptr }

func (this *QVariant) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVariant) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVariantFromPointer(cthis unsafe.Pointer) *QVariant {
	return &QVariant{&qtrt.CObject{cthis}}
}
func (*QVariant) NewFromPointer(cthis unsafe.Pointer) *QVariant {
	return NewQVariantFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:199
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVariant()

/*
Constructs an invalid variant.
*/
func NewQVariant() *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:201
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QVariant(QVariant::Type)

/*
Constructs an invalid variant.
*/
func NewQVariant_1(type_ int) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:202
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int, const void *)

/*
Constructs an invalid variant.
*/
func NewQVariant_2(typeId int, copy unsafe.Pointer /*666*/) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2EiPKv", qtrt.FFI_TYPE_POINTER, typeId, copy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:203
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int, const void *, uint)

/*
Constructs an invalid variant.
*/
func NewQVariant_3(typeId int, copy unsafe.Pointer /*666*/, flags uint) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2EiPKvj", qtrt.FFI_TYPE_POINTER, typeId, copy, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:207
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QVariant(QDataStream &)

/*
Constructs an invalid variant.
*/
func NewQVariant_4(s QDataStream_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if s != nil && s.QDataStream_PTR() != nil {
		convArg0 = s.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ER11QDataStream", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:210
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int)

/*
Constructs an invalid variant.
*/
func NewQVariant_5(i int) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ei", qtrt.FFI_TYPE_POINTER, i)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:211
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVariant(uint)

/*
Constructs an invalid variant.
*/
func NewQVariant_6(ui uint) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ej", qtrt.FFI_TYPE_POINTER, ui)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:212
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QVariant(qlonglong)

/*
Constructs an invalid variant.
*/
func NewQVariant_7(ll int64) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ex", qtrt.FFI_TYPE_POINTER, ll)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:213
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QVariant(qulonglong)

/*
Constructs an invalid variant.
*/
func NewQVariant_8(ull uint64) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ey", qtrt.FFI_TYPE_POINTER, ull)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:214
// index:9
// Public Visibility=Default Availability=Available
// [-2] void QVariant(bool)

/*
Constructs an invalid variant.
*/
func NewQVariant_9(b bool) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Eb", qtrt.FFI_TYPE_POINTER, b)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:215
// index:10
// Public Visibility=Default Availability=Available
// [-2] void QVariant(double)

/*
Constructs an invalid variant.
*/
func NewQVariant_10(d float64) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ed", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:216
// index:11
// Public Visibility=Default Availability=Available
// [-2] void QVariant(float)

/*
Constructs an invalid variant.
*/
func NewQVariant_11(f float32) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ef", qtrt.FFI_TYPE_POINTER, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:218
// index:12
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const char *)

/*
Constructs an invalid variant.
*/
func NewQVariant_12(str string) *QVariant {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:221
// index:13
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QByteArray &)

/*
Constructs an invalid variant.
*/
func NewQVariant_13(bytearray QByteArray_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if bytearray != nil && bytearray.QByteArray_PTR() != nil {
		convArg0 = bytearray.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:222
// index:14
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QBitArray &)

/*
Constructs an invalid variant.
*/
func NewQVariant_14(bitarray QBitArray_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if bitarray != nil && bitarray.QBitArray_PTR() != nil {
		convArg0 = bitarray.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK9QBitArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:223
// index:15
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QString &)

/*
Constructs an invalid variant.
*/
func NewQVariant_15(string string) *QVariant {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:224
// index:16
// Public Visibility=Default Availability=Available
// [-2] void QVariant(QLatin1String)

/*
Constructs an invalid variant.
*/
func NewQVariant_16(string QLatin1String_ITF /*123*/) *QVariant {
	var convArg0 unsafe.Pointer
	if string != nil && string.QLatin1String_PTR() != nil {
		convArg0 = string.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2E13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:225
// index:17
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QStringList &)

/*
Constructs an invalid variant.
*/
func NewQVariant_17(stringlist QStringList_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if stringlist != nil && stringlist.QStringList_PTR() != nil {
		convArg0 = stringlist.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:226
// index:18
// Public Visibility=Default Availability=Available
// [-2] void QVariant(QChar)

/*
Constructs an invalid variant.
*/
func NewQVariant_18(qchar QChar_ITF /*123*/) *QVariant {
	var convArg0 unsafe.Pointer
	if qchar != nil && qchar.QChar_PTR() != nil {
		convArg0 = qchar.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2E5QChar", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:227
// index:19
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QDate &)

/*
Constructs an invalid variant.
*/
func NewQVariant_19(date QDate_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QDate", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:228
// index:20
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QTime &)

/*
Constructs an invalid variant.
*/
func NewQVariant_20(time QTime_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QTime", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:229
// index:21
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QDateTime &)

/*
Constructs an invalid variant.
*/
func NewQVariant_21(datetime QDateTime_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if datetime != nil && datetime.QDateTime_PTR() != nil {
		convArg0 = datetime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK9QDateTime", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:234
// index:22
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QSize &)

/*
Constructs an invalid variant.
*/
func NewQVariant_22(size QSize_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QSize", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:235
// index:23
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QSizeF &)

/*
Constructs an invalid variant.
*/
func NewQVariant_23(size QSizeF_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK6QSizeF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:236
// index:24
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QPoint &)

/*
Constructs an invalid variant.
*/
func NewQVariant_24(pt QPoint_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:237
// index:25
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QPointF &)

/*
Constructs an invalid variant.
*/
func NewQVariant_25(pt QPointF_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPointF_PTR() != nil {
		convArg0 = pt.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:238
// index:26
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QLine &)

/*
Constructs an invalid variant.
*/
func NewQVariant_26(line QLine_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if line != nil && line.QLine_PTR() != nil {
		convArg0 = line.QLine_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QLine", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:239
// index:27
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QLineF &)

/*
Constructs an invalid variant.
*/
func NewQVariant_27(line QLineF_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if line != nil && line.QLineF_PTR() != nil {
		convArg0 = line.QLineF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK6QLineF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:240
// index:28
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRect &)

/*
Constructs an invalid variant.
*/
func NewQVariant_28(rect QRect_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QRect", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:241
// index:29
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRectF &)

/*
Constructs an invalid variant.
*/
func NewQVariant_29(rect QRectF_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK6QRectF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:243
// index:30
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QLocale &)

/*
Constructs an invalid variant.
*/
func NewQVariant_30(locale QLocale_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK7QLocale", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:245
// index:31
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRegExp &)

/*
Constructs an invalid variant.
*/
func NewQVariant_31(regExp QRegExp_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if regExp != nil && regExp.QRegExp_PTR() != nil {
		convArg0 = regExp.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK7QRegExp", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:249
// index:32
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRegularExpression &)

/*
Constructs an invalid variant.
*/
func NewQVariant_32(re QRegularExpression_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:251
// index:33
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QUrl &)

/*
Constructs an invalid variant.
*/
func NewQVariant_33(url QUrl_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:252
// index:34
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QEasingCurve &)

/*
Constructs an invalid variant.
*/
func NewQVariant_34(easing QEasingCurve_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if easing != nil && easing.QEasingCurve_PTR() != nil {
		convArg0 = easing.QEasingCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK12QEasingCurve", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:253
// index:35
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QUuid &)

/*
Constructs an invalid variant.
*/
func NewQVariant_35(uuid QUuid_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if uuid != nil && uuid.QUuid_PTR() != nil {
		convArg0 = uuid.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QUuid", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:254
// index:36
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QModelIndex &)

/*
Constructs an invalid variant.
*/
func NewQVariant_36(modelIndex QModelIndex_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if modelIndex != nil && modelIndex.QModelIndex_PTR() != nil {
		convArg0 = modelIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK11QModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:255
// index:37
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QPersistentModelIndex &)

/*
Constructs an invalid variant.
*/
func NewQVariant_37(modelIndex QPersistentModelIndex_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if modelIndex != nil && modelIndex.QPersistentModelIndex_PTR() != nil {
		convArg0 = modelIndex.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK21QPersistentModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:256
// index:38
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QJsonValue &)

/*
Constructs an invalid variant.
*/
func NewQVariant_38(jsonValue QJsonValue_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if jsonValue != nil && jsonValue.QJsonValue_PTR() != nil {
		convArg0 = jsonValue.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK10QJsonValue", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:257
// index:39
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QJsonObject &)

/*
Constructs an invalid variant.
*/
func NewQVariant_39(jsonObject QJsonObject_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if jsonObject != nil && jsonObject.QJsonObject_PTR() != nil {
		convArg0 = jsonObject.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK11QJsonObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:258
// index:40
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QJsonArray &)

/*
Constructs an invalid variant.
*/
func NewQVariant_40(jsonArray QJsonArray_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if jsonArray != nil && jsonArray.QJsonArray_PTR() != nil {
		convArg0 = jsonArray.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK10QJsonArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:259
// index:41
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QJsonDocument &)

/*
Constructs an invalid variant.
*/
func NewQVariant_41(jsonDocument QJsonDocument_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if jsonDocument != nil && jsonDocument.QJsonDocument_PTR() != nil {
		convArg0 = jsonDocument.QJsonDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK13QJsonDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QVariant()

/*

 */
func DeleteQVariant(this *QVariant) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qvariant.h:262
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant & operator=(const QVariant &)

/*

 */
func (this *QVariant) Operator_equal(other QVariant_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVariant_PTR() != nil {
		convArg0 = other.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:266
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QVariant & operator=(QVariant &&)

/*

 */
func (this *QVariant) Operator_equal_1(other unsafe.Pointer /*333*/) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:270
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QVariant &)

/*
Swaps variant other with this variant. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QVariant) Swap(other QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVariant_PTR() != nil {
		convArg0 = other.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:272
// index:0
// Public Visibility=Default Availability=Available
// [4] QVariant::Type type() const

/*
Returns the storage type of the value stored in the variant. Although this function is declared as returning QVariant::Type, the return value should be interpreted as QMetaType::Type. In particular, QVariant::UserType is returned here only if the value is equal or greater than QMetaType::User.

Note that return values in the ranges QVariant::Char through QVariant::RegExp and QVariant::Font through QVariant::Transform correspond to the values in the ranges QMetaType::QChar through QMetaType::QRegExp and QMetaType::QFont through QMetaType::QQuaternion.

Pay particular attention when working with char and QChar variants. Note that there is no QVariant constructor specifically for type char, but there is one for QChar. For a variant of type QChar, this function returns QVariant::Char, which is the same as QMetaType::QChar, but for a variant of type char, this function returns QMetaType::Char, which is not the same as QVariant::Char.

Also note that the types void*, long, short, unsigned long, unsigned short, unsigned char, float, QObject*, and QWidget* are represented in QMetaType::Type but not in QVariant::Type, and they can be returned by this function. However, they are considered to be user defined types when tested against QVariant::Type.

To test whether an instance of QVariant contains a data type that is compatible with the data type you are interested in, use canConvert().
*/
func (this *QVariant) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qvariant.h:273
// index:0
// Public Visibility=Default Availability=Available
// [4] int userType() const

/*
Returns the storage type of the value stored in the variant. For non-user types, this is the same as type().

See also type().
*/
func (this *QVariant) UserType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8userTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:274
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * typeName() const

/*
Returns the name of the type stored in the variant. The returned strings describe the C++ datatype used to store the data: for example, "QFont", "QString", or "QVariantList". An Invalid variant returns 0.
*/
func (this *QVariant) TypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8typeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qvariant.h:276
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canConvert(int) const

/*
Returns true if the variant's type can be cast to the requested type, targetTypeId. Such casting is done automatically when calling the toInt(), toBool(), ... methods.

The following casts are done automatically:


 TypeAutomatically Cast To
QMetaType::BoolQMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, QMetaType::ULongLong
QMetaType::QByteArrayQMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, QMetaType::ULongLong, QMetaType::QUuid
QMetaType::QCharQMetaType::Bool, QMetaType::Int, QMetaType::UInt, QMetaType::LongLong, QMetaType::ULongLong
QMetaType::QColorQMetaType::QString
QMetaType::QDateQMetaType::QDateTime, QMetaType::QString
QMetaType::QDateTimeQMetaType::QDate, QMetaType::QString, QMetaType::QTime
QMetaType::DoubleQMetaType::Bool, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, QMetaType::ULongLong
QMetaType::QFontQMetaType::QString
QMetaType::IntQMetaType::Bool, QMetaType::QChar, QMetaType::Double, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, QMetaType::ULongLong
QMetaType::QKeySequenceQMetaType::Int, QMetaType::QString
QMetaType::QVariantListQMetaType::QStringList (if the list's items can be converted to QStrings)
QMetaType::LongLongQMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::QString, QMetaType::UInt, QMetaType::ULongLong
QMetaType::QPointQMetaType::QPointF
QMetaType::QRectQMetaType::QRectF
QMetaType::QStringQMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::QColor, QMetaType::QDate, QMetaType::QDateTime, QMetaType::Double, QMetaType::QFont, QMetaType::Int, QMetaType::QKeySequence, QMetaType::LongLong, QMetaType::QStringList, QMetaType::QTime, QMetaType::UInt, QMetaType::ULongLong, QMetaType::QUuid
QMetaType::QStringListQMetaType::QVariantList, QMetaType::QString (if the list contains exactly one item)
QMetaType::QTimeQMetaType::QString
QMetaType::UIntQMetaType::Bool, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::ULongLong
QMetaType::ULongLongQMetaType::Bool, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt
QMetaType::QUuidQMetaType::QByteArray, QMetaType::QString


A QVariant containing a pointer to a type derived from QObject will also return true for this function if a qobject_cast to the type described by targetTypeId would succeed. Note that this only works for QObject subclasses which use the Q_OBJECT macro.

A QVariant containing a sequential container will also return true for this function if the targetTypeId is QVariantList. It is possible to iterate over the contents of the container without extracting it as a (copied) QVariantList:


  QList<int> intList = {7, 11, 42};

  QVariant variant = QVariant::fromValue(intList);
  if (variant.canConvert<QVariantList>()) {
      QSequentialIterable iterable = variant.value<QSequentialIterable>();
      // Can use foreach:
      foreach (const QVariant &v, iterable) {
          qDebug() << v;
      }
      // Can use C++11 range-for:
      for (const QVariant &v : iterable) {
          qDebug() << v;
      }
      // Can use iterators:
      QSequentialIterable::const_iterator it = iterable.begin();
      const QSequentialIterable::const_iterator end = iterable.end();
      for ( ; it != end; ++it) {
          qDebug() << *it;
      }
  }



This requires that the value_type of the container is itself a metatype.

Similarly, a QVariant containing a sequential container will also return true for this function the targetTypeId is QVariantHash or QVariantMap. It is possible to iterate over the contents of the container without extracting it as a (copied) QVariantHash or QVariantMap:


  QHash<int, QString> mapping;
  mapping.insert(7, "Seven");
  mapping.insert(11, "Eleven");
  mapping.insert(42, "Forty-two");

  QVariant variant = QVariant::fromValue(mapping);
  if (variant.canConvert<QVariantHash>()) {
      QAssociativeIterable iterable = variant.value<QAssociativeIterable>();
      // Can use foreach over the values:
      foreach (const QVariant &v, iterable) {
          qDebug() << v;
      }
      // Can use C++11 range-for over the values:
      for (const QVariant &v : iterable) {
          qDebug() << v;
      }
      // Can use iterators:
      QAssociativeIterable::const_iterator it = iterable.begin();
      const QAssociativeIterable::const_iterator end = iterable.end();
      for ( ; it != end; ++it) {
          qDebug() << *it; // The current value
          qDebug() << it.key();
          qDebug() << it.value();
      }
  }



See also convert(), QSequentialIterable, Q_DECLARE_SEQUENTIAL_CONTAINER_METATYPE(), QAssociativeIterable, and Q_DECLARE_ASSOCIATIVE_CONTAINER_METATYPE().
*/
func (this *QVariant) CanConvert(targetTypeId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10canConvertEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), targetTypeId)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:277
// index:0
// Public Visibility=Default Availability=Available
// [1] bool convert(int)

/*
Casts the variant to the requested type, targetTypeId. If the cast cannot be done, the variant is still changed to the requested type, but is left in a cleared null state similar to that constructed by QVariant(Type).

Returns true if the current type of the variant was successfully cast; otherwise returns false.

A QVariant containing a pointer to a type derived from QObject will also convert and return true for this function if a qobject_cast to the type described by targetTypeId would succeed. Note that this only works for QObject subclasses which use the Q_OBJECT macro.

Note: converting QVariants that are null due to not being initialized or having failed a previous conversion will always fail, changing the type, remaining null, and returning false.

See also canConvert() and clear().
*/
func (this *QVariant) Convert(targetTypeId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant7convertEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), targetTypeId)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:467
// index:1
// Protected Visibility=Default Availability=Available
// [1] bool convert(const int, void *) const

/*
Casts the variant to the requested type, targetTypeId. If the cast cannot be done, the variant is still changed to the requested type, but is left in a cleared null state similar to that constructed by QVariant(Type).

Returns true if the current type of the variant was successfully cast; otherwise returns false.

A QVariant containing a pointer to a type derived from QObject will also convert and return true for this function if a qobject_cast to the type described by targetTypeId would succeed. Note that this only works for QObject subclasses which use the Q_OBJECT macro.

Note: converting QVariants that are null due to not being initialized or having failed a previous conversion will always fail, changing the type, remaining null, and returning false.

See also canConvert() and clear().
*/
func (this *QVariant) Convert_1(t int, ptr unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7convertEiPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t, ptr)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:279
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the storage type of this variant is not QMetaType::UnknownType; otherwise returns false.
*/
func (this *QVariant) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:280
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this is a null variant, false otherwise. A variant is considered null if it contains no initialized value, or the contained value is a null pointer or is an instance of a built-in type that has an isNull method, in which case the result would be the same as calling isNull on the wrapped object.

Warning: Null variants is not a single state and two null variants may easily return false on the == operator if they do not contain similar null values.

See also QVariant(Type) and convert(int).
*/
func (this *QVariant) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Convert this variant to type QMetaType::UnknownType and free up any resources used.
*/
func (this *QVariant) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:284
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()

/*

 */
func (this *QVariant) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:285
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QVariant) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:287
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *) const

/*
Returns the variant as an int if the variant has userType() QMetaType::Int, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.

If ok is non-null: *ok is set to true if the value could be converted to an int; otherwise *ok is set to false.

Warning: If the value is convertible to a QMetaType::LongLong but is too large to be represented in an int, the resulting arithmetic overflow will not be reflected in ok. A simple workaround is to use QString::toInt().

See also canConvert() and convert().
*/
func (this *QVariant) ToInt(ok *bool) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant5toIntEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:287
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *) const

/*
Returns the variant as an int if the variant has userType() QMetaType::Int, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.

If ok is non-null: *ok is set to true if the value could be converted to an int; otherwise *ok is set to false.

Warning: If the value is convertible to a QMetaType::LongLong but is too large to be represented in an int, the resulting arithmetic overflow will not be reflected in ok. A simple workaround is to use QString::toInt().

See also canConvert() and convert().
*/
func (this *QVariant) ToInt__() int {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant5toIntEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:288
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *) const

/*
Returns the variant as an unsigned int if the variant has userType() QMetaType::UInt, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, or QMetaType::ULongLong; otherwise returns 0.

If ok is non-null: *ok is set to true if the value could be converted to an unsigned int; otherwise *ok is set to false.

Warning: If the value is convertible to a QMetaType::ULongLong but is too large to be represented in an unsigned int, the resulting arithmetic overflow will not be reflected in ok. A simple workaround is to use QString::toUInt().

See also canConvert() and convert().
*/
func (this *QVariant) ToUInt(ok *bool) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toUIntEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:288
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *) const

/*
Returns the variant as an unsigned int if the variant has userType() QMetaType::UInt, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, or QMetaType::ULongLong; otherwise returns 0.

If ok is non-null: *ok is set to true if the value could be converted to an unsigned int; otherwise *ok is set to false.

Warning: If the value is convertible to a QMetaType::ULongLong but is too large to be represented in an unsigned int, the resulting arithmetic overflow will not be reflected in ok. A simple workaround is to use QString::toUInt().

See also canConvert() and convert().
*/
func (this *QVariant) ToUInt__() uint {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toUIntEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:289
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *) const

/*
Returns the variant as a long long int if the variant has userType() QMetaType::LongLong, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.

If ok is non-null: *ok is set to true if the value could be converted to an int; otherwise *ok is set to false.

See also canConvert() and convert().
*/
func (this *QVariant) ToLongLong(ok *bool) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10toLongLongEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:289
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *) const

/*
Returns the variant as a long long int if the variant has userType() QMetaType::LongLong, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.

If ok is non-null: *ok is set to true if the value could be converted to an int; otherwise *ok is set to false.

See also canConvert() and convert().
*/
func (this *QVariant) ToLongLong__() int64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10toLongLongEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:290
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *) const

/*
Returns the variant as an unsigned long long int if the variant has type() QMetaType::ULongLong, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, or QMetaType::UInt; otherwise returns 0.

If ok is non-null: *ok is set to true if the value could be converted to an int; otherwise *ok is set to false.

See also canConvert() and convert().
*/
func (this *QVariant) ToULongLong(ok *bool) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toULongLongEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:290
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *) const

/*
Returns the variant as an unsigned long long int if the variant has type() QMetaType::ULongLong, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, or QMetaType::UInt; otherwise returns 0.

If ok is non-null: *ok is set to true if the value could be converted to an int; otherwise *ok is set to false.

See also canConvert() and convert().
*/
func (this *QVariant) ToULongLong__() uint64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toULongLongEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:291
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toBool() const

/*
Returns the variant as a bool if the variant has userType() Bool.

Returns true if the variant has userType() QMetaType::Bool, QMetaType::QChar, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::UInt, or QMetaType::ULongLong and the value is non-zero, or if the variant has type QMetaType::QString or QMetaType::QByteArray and its lower-case content is not one of the following: empty, "0" or "false"; otherwise returns false.

See also canConvert() and convert().
*/
func (this *QVariant) ToBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:292
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
Returns the variant as a double if the variant has userType() QMetaType::Double, QMetaType::Float, QMetaType::Bool, QMetaType::QByteArray, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.0.

If ok is non-null: *ok is set to true if the value could be converted to a double; otherwise *ok is set to false.

See also canConvert() and convert().
*/
func (this *QVariant) ToDouble(ok *bool) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:292
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
Returns the variant as a double if the variant has userType() QMetaType::Double, QMetaType::Float, QMetaType::Bool, QMetaType::QByteArray, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.0.

If ok is non-null: *ok is set to true if the value could be converted to a double; otherwise *ok is set to false.

See also canConvert() and convert().
*/
func (this *QVariant) ToDouble__() float64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:293
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
Returns the variant as a float if the variant has userType() QMetaType::Double, QMetaType::Float, QMetaType::Bool, QMetaType::QByteArray, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.0.

If ok is non-null: *ok is set to true if the value could be converted to a double; otherwise *ok is set to false.

This function was introduced in  Qt 4.6.

See also canConvert() and convert().
*/
func (this *QVariant) ToFloat(ok *bool) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:293
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
Returns the variant as a float if the variant has userType() QMetaType::Double, QMetaType::Float, QMetaType::Bool, QMetaType::QByteArray, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.0.

If ok is non-null: *ok is set to true if the value could be converted to a double; otherwise *ok is set to false.

This function was introduced in  Qt 4.6.

See also canConvert() and convert().
*/
func (this *QVariant) ToFloat__() float32 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:294
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal toReal(bool *) const

/*
Returns the variant as a qreal if the variant has userType() QMetaType::Double, QMetaType::Float, QMetaType::Bool, QMetaType::QByteArray, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.0.

If ok is non-null: *ok is set to true if the value could be converted to a double; otherwise *ok is set to false.

This function was introduced in  Qt 4.6.

See also canConvert() and convert().
*/
func (this *QVariant) ToReal(ok *bool) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toRealEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:294
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal toReal(bool *) const

/*
Returns the variant as a qreal if the variant has userType() QMetaType::Double, QMetaType::Float, QMetaType::Bool, QMetaType::QByteArray, QMetaType::Int, QMetaType::LongLong, QMetaType::QString, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns 0.0.

If ok is non-null: *ok is set to true if the value could be converted to a double; otherwise *ok is set to false.

This function was introduced in  Qt 4.6.

See also canConvert() and convert().
*/
func (this *QVariant) ToReal__() float64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toRealEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:295
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toByteArray() const

/*
Returns the variant as a QByteArray if the variant has userType() QMetaType::QByteArray or QMetaType::QString (converted using QString::fromUtf8()); otherwise returns an empty byte array.

See also canConvert() and convert().
*/
func (this *QVariant) ToByteArray() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toByteArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:296
// index:0
// Public Visibility=Default Availability=Available
// [8] QBitArray toBitArray() const

/*
Returns the variant as a QBitArray if the variant has userType() QMetaType::QBitArray; otherwise returns an empty bit array.

See also canConvert() and convert().
*/
func (this *QVariant) ToBitArray() *QBitArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10toBitArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitArray)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:297
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the variant as a QString if the variant has userType() QMetaType::QString, QMetaType::Bool, QMetaType::QByteArray, QMetaType::QChar, QMetaType::QDate, QMetaType::QDateTime, QMetaType::Double, QMetaType::Int, QMetaType::LongLong, QMetaType::QStringList, QMetaType::QTime, QMetaType::UInt, or QMetaType::ULongLong; otherwise returns an empty string.

See also canConvert() and convert().
*/
func (this *QVariant) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qvariant.h:298
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList toStringList() const

/*
Returns the variant as a QStringList if the variant has userType() QMetaType::QStringList, QMetaType::QString, or QMetaType::QVariantList of a type that can be converted to QString; otherwise returns an empty list.

See also canConvert() and convert().
*/
func (this *QVariant) ToStringList() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant12toStringListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:299
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar toChar() const

/*
Returns the variant as a QChar if the variant has userType() QMetaType::QChar, QMetaType::Int, or QMetaType::UInt; otherwise returns an invalid QChar.

See also canConvert() and convert().
*/
func (this *QVariant) ToChar() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:300
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate toDate() const

/*
Returns the variant as a QDate if the variant has userType() QMetaType::QDate, QMetaType::QDateTime, or QMetaType::QString; otherwise returns an invalid date.

If the type() is QMetaType::QString, an invalid date will be returned if the string cannot be parsed as a Qt::ISODate format date.

See also canConvert() and convert().
*/
func (this *QVariant) ToDate() *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:301
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime toTime() const

/*
Returns the variant as a QTime if the variant has userType() QMetaType::QTime, QMetaType::QDateTime, or QMetaType::QString; otherwise returns an invalid time.

If the type() is QMetaType::QString, an invalid time will be returned if the string cannot be parsed as a Qt::ISODate format time.

See also canConvert() and convert().
*/
func (this *QVariant) ToTime() *QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:302
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime() const

/*
Returns the variant as a QDateTime if the variant has userType() QMetaType::QDateTime, QMetaType::QDate, or QMetaType::QString; otherwise returns an invalid date/time.

If the type() is QMetaType::QString, an invalid date/time will be returned if the string cannot be parsed as a Qt::ISODate format date/time.

See also canConvert() and convert().
*/
func (this *QVariant) ToDateTime() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10toDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:308
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint toPoint() const

/*
Returns the variant as a QPoint if the variant has userType() QMetaType::QPoint or QMetaType::QPointF; otherwise returns a null QPoint.

See also canConvert() and convert().
*/
func (this *QVariant) ToPoint() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:309
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF toPointF() const

/*
Returns the variant as a QPointF if the variant has userType() QMetaType::QPoint or QMetaType::QPointF; otherwise returns a null QPointF.

See also canConvert() and convert().
*/
func (this *QVariant) ToPointF() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toPointFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:310
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect toRect() const

/*
Returns the variant as a QRect if the variant has userType() QMetaType::QRect; otherwise returns an invalid QRect.

See also canConvert() and convert().
*/
func (this *QVariant) ToRect() *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:311
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize toSize() const

/*
Returns the variant as a QSize if the variant has userType() QMetaType::QSize; otherwise returns an invalid QSize.

See also canConvert() and convert().
*/
func (this *QVariant) ToSize() *QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:312
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF toSizeF() const

/*
Returns the variant as a QSizeF if the variant has userType() QMetaType::QSizeF; otherwise returns an invalid QSizeF.

See also canConvert() and convert().
*/
func (this *QVariant) ToSizeF() *QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toSizeFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:313
// index:0
// Public Visibility=Default Availability=Available
// [16] QLine toLine() const

/*
Returns the variant as a QLine if the variant has userType() QMetaType::QLine; otherwise returns an invalid QLine.

See also canConvert() and convert().
*/
func (this *QVariant) ToLine() *QLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLine)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:314
// index:0
// Public Visibility=Default Availability=Available
// [32] QLineF toLineF() const

/*
Returns the variant as a QLineF if the variant has userType() QMetaType::QLineF; otherwise returns an invalid QLineF.

See also canConvert() and convert().
*/
func (this *QVariant) ToLineF() *QLineF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toLineFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:315
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF toRectF() const

/*
Returns the variant as a QRectF if the variant has userType() QMetaType::QRect or QMetaType::QRectF; otherwise returns an invalid QRectF.

See also canConvert() and convert().
*/
func (this *QVariant) ToRectF() *QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toRectFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:317
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale toLocale() const

/*
Returns the variant as a QLocale if the variant has userType() QMetaType::QLocale; otherwise returns an invalid QLocale.

See also canConvert() and convert().
*/
func (this *QVariant) ToLocale() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toLocaleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegExp toRegExp() const

/*
Returns the variant as a QRegExp if the variant has userType() QMetaType::QRegExp; otherwise returns an empty QRegExp.

This function was introduced in  Qt 4.1.

See also canConvert() and convert().
*/
func (this *QVariant) ToRegExp() *QRegExp /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toRegExpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegExp)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:323
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression toRegularExpression() const

/*
Returns the variant as a QRegularExpression if the variant has userType() QRegularExpression; otherwise returns an empty QRegularExpression.

This function was introduced in  Qt 5.0.

See also canConvert() and convert().
*/
func (this *QVariant) ToRegularExpression() *QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant19toRegularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:325
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl toUrl() const

/*
Returns the variant as a QUrl if the variant has userType() QMetaType::QUrl; otherwise returns an invalid QUrl.

See also canConvert() and convert().
*/
func (this *QVariant) ToUrl() *QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant5toUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:326
// index:0
// Public Visibility=Default Availability=Available
// [8] QEasingCurve toEasingCurve() const

/*
Returns the variant as a QEasingCurve if the variant has userType() QMetaType::QEasingCurve; otherwise returns a default easing curve.

This function was introduced in  Qt 4.7.

See also canConvert() and convert().
*/
func (this *QVariant) ToEasingCurve() *QEasingCurve /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant13toEasingCurveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQEasingCurve)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:327
// index:0
// Public Visibility=Default Availability=Available
// [16] QUuid toUuid() const

/*
Returns the variant as a QUuid if the variant has type() QMetaType::QUuid, QMetaType::QByteArray or QMetaType::QString; otherwise returns a default-constructed QUuid.

This function was introduced in  Qt 5.0.

See also canConvert() and convert().
*/
func (this *QVariant) ToUuid() *QUuid /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toUuidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:328
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex toModelIndex() const

/*
Returns the variant as a QModelIndex if the variant has userType() QModelIndex; otherwise returns a default constructed QModelIndex.

This function was introduced in  Qt 5.0.

See also canConvert(), convert(), and toPersistentModelIndex().
*/
func (this *QVariant) ToModelIndex() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant12toModelIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:329
// index:0
// Public Visibility=Default Availability=Available
// [8] QPersistentModelIndex toPersistentModelIndex() const

/*
Returns the variant as a QPersistentModelIndex if the variant has userType() QPersistentModelIndex; otherwise returns a default constructed QPersistentModelIndex.

This function was introduced in  Qt 5.5.

See also canConvert(), convert(), and toModelIndex().
*/
func (this *QVariant) ToPersistentModelIndex() *QPersistentModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant22toPersistentModelIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:330
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue toJsonValue() const

/*
Returns the variant as a QJsonValue if the variant has userType() QJsonValue; otherwise returns a default constructed QJsonValue.

This function was introduced in  Qt 5.0.

See also canConvert() and convert().
*/
func (this *QVariant) ToJsonValue() *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toJsonValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:331
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject toJsonObject() const

/*
Returns the variant as a QJsonObject if the variant has userType() QJsonObject; otherwise returns a default constructed QJsonObject.

This function was introduced in  Qt 5.0.

See also canConvert() and convert().
*/
func (this *QVariant) ToJsonObject() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant12toJsonObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:332
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray toJsonArray() const

/*
Returns the variant as a QJsonArray if the variant has userType() QJsonArray; otherwise returns a default constructed QJsonArray.

This function was introduced in  Qt 5.0.

See also canConvert() and convert().
*/
func (this *QVariant) ToJsonArray() *QJsonArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toJsonArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:333
// index:0
// Public Visibility=Default Availability=Available
// [8] QJsonDocument toJsonDocument() const

/*
Returns the variant as a QJsonDocument if the variant has userType() QJsonDocument; otherwise returns a default constructed QJsonDocument.

This function was introduced in  Qt 5.0.

See also canConvert() and convert().
*/
func (this *QVariant) ToJsonDocument() *QJsonDocument /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant14toJsonDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:337
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(QDataStream &)

/*

 */
func (this *QVariant) Load(ds QDataStream_ITF) {
	var convArg0 unsafe.Pointer
	if ds != nil && ds.QDataStream_PTR() != nil {
		convArg0 = ds.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant4loadER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void save(QDataStream &) const

/*

 */
func (this *QVariant) Save(ds QDataStream_ITF) {
	var convArg0 unsafe.Pointer
	if ds != nil && ds.QDataStream_PTR() != nil {
		convArg0 = ds.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant4saveER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:340
// index:0
// Public static Visibility=Default Availability=Available
// [8] const char * typeToName(int)

/*
Converts the int representation of the storage type, typeId, to its string representation.

Returns a null pointer if the type is QMetaType::UnknownType or doesn't exist.
*/
func (this *QVariant) TypeToName(typeId int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant10typeToNameEi", qtrt.FFI_TYPE_POINTER, typeId)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}
func QVariant_TypeToName(typeId int) string {
	var nilthis *QVariant
	rv := nilthis.TypeToName(typeId)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:341
// index:0
// Public static Visibility=Default Availability=Available
// [4] QVariant::Type nameToType(const char *)

/*
Converts the string representation of the storage type given in name, to its enum representation.

If the string representation cannot be converted to any enum representation, the variant is set to Invalid.
*/
func (this *QVariant) NameToType(name string) int {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant10nameToTypeEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QVariant_NameToType(name string) int {
	var nilthis *QVariant
	rv := nilthis.NameToType(name)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:343
// index:0
// Public Visibility=Default Availability=Available
// [8] void * data()

/*

 */
func (this *QVariant) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:345
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const void * data() const

/*

 */
func (this *QVariant) Data_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:344
// index:0
// Public Visibility=Default Availability=Available
// [8] const void * constData() const

/*

 */
func (this *QVariant) ConstData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:436
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QVariant &) const

/*

 */
func (this *QVariant) Operator_equal_equal(v QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if v != nil && v.QVariant_PTR() != nil {
		convArg0 = v.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVarianteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:438
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QVariant &) const

/*

 */
func (this *QVariant) Operator_not_equal(v QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if v != nil && v.QVariant_PTR() != nil {
		convArg0 = v.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariantneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:440
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QVariant &) const

/*

 */
func (this *QVariant) Operator_less_than(v QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if v != nil && v.QVariant_PTR() != nil {
		convArg0 = v.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariantltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:442
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QVariant &) const

/*

 */
func (this *QVariant) Operator_less_than_equal(v QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if v != nil && v.QVariant_PTR() != nil {
		convArg0 = v.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariantleERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:444
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QVariant &) const

/*

 */
func (this *QVariant) Operator_greater_than(v QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if v != nil && v.QVariant_PTR() != nil {
		convArg0 = v.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariantgtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:446
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QVariant &) const

/*

 */
func (this *QVariant) Operator_greater_than_equal(v QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if v != nil && v.QVariant_PTR() != nil {
		convArg0 = v.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariantgeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:464
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void create(int, const void *)

/*

 */
func (this *QVariant) Create(type_ int, copy unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant6createEiPKv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, copy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:465
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool cmp(const QVariant &) const

/*

 */
func (this *QVariant) Cmp(other QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVariant_PTR() != nil {
		convArg0 = other.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant3cmpERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:466
// index:0
// Protected Visibility=Default Availability=Available
// [4] int compare(const QVariant &) const

/*

 */
func (this *QVariant) Compare(other QVariant_ITF) int {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVariant_PTR() != nil {
		convArg0 = other.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7compareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
const QVariant__LastCoreType QVariant__Type = 51

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
const QVariant__LastGuiType QVariant__Type = 86

//
const QVariant__SizePolicy QVariant__Type = 121

//
const QVariant__UserType QVariant__Type = 1024

//
const QVariant__LastType QVariant__Type = -1

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
