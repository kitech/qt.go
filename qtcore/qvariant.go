package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
// void create(int, const void *)
func (this *QVariant) InheritCreate(f func(type_ int, copy unsafe.Pointer /*666*/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "create", f)
}

// bool cmp(const class QVariant &)
func (this *QVariant) InheritCmp(f func(other *QVariant) bool) {
	qtrt.SetAllInheritCallback(this, "cmp", f)
}

// int compare(const class QVariant &)
func (this *QVariant) InheritCompare(f func(other *QVariant) int) {
	qtrt.SetAllInheritCallback(this, "compare", f)
}

// bool convert(const int, void *)
func (this *QVariant) InheritConvert(f func(t int, ptr unsafe.Pointer /*666*/) bool) {
	qtrt.SetAllInheritCallback(this, "convert", f)
}

type QVariant struct {
	*qtrt.CObject
}

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
func NewQVariant() *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:201
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QVariant(enum QVariant::Type)
func NewQVariant_1(type_ int) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:202
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int, const void *)
func NewQVariant_2(typeId int, copy unsafe.Pointer /*666*/) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2EiPKv", qtrt.FFI_TYPE_POINTER, typeId, copy)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:203
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int, const void *, uint)
func NewQVariant_3(typeId int, copy unsafe.Pointer /*666*/, flags uint) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2EiPKvj", qtrt.FFI_TYPE_POINTER, typeId, copy, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:207
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QVariant(QDataStream &)
func NewQVariant_4(s *QDataStream) *QVariant {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ER11QDataStream", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:210
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVariant(int)
func NewQVariant_5(i int) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ei", qtrt.FFI_TYPE_POINTER, i)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:211
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVariant(uint)
func NewQVariant_6(ui uint) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ej", qtrt.FFI_TYPE_POINTER, ui)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:212
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QVariant(qlonglong)
func NewQVariant_7(ll int64) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ex", qtrt.FFI_TYPE_POINTER, ll)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:213
// index:8
// Public Visibility=Default Availability=Available
// [-2] void QVariant(qulonglong)
func NewQVariant_8(ull uint64) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ey", qtrt.FFI_TYPE_POINTER, ull)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:214
// index:9
// Public Visibility=Default Availability=Available
// [-2] void QVariant(_Bool)
func NewQVariant_9(b bool) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Eb", qtrt.FFI_TYPE_POINTER, b)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:215
// index:10
// Public Visibility=Default Availability=Available
// [-2] void QVariant(double)
func NewQVariant_10(d float64) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ed", qtrt.FFI_TYPE_POINTER, d)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:216
// index:11
// Public Visibility=Default Availability=Available
// [-2] void QVariant(float)
func NewQVariant_11(f float32) *QVariant {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2Ef", qtrt.FFI_TYPE_POINTER, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:218
// index:12
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const char *)
func NewQVariant_12(str string) *QVariant {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:221
// index:13
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QByteArray &)
func NewQVariant_13(bytearray *QByteArray) *QVariant {
	var convArg0 = bytearray.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:222
// index:14
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QBitArray &)
func NewQVariant_14(bitarray *QBitArray) *QVariant {
	var convArg0 = bitarray.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK9QBitArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:223
// index:15
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QString &)
func NewQVariant_15(string string) *QVariant {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:224
// index:16
// Public Visibility=Default Availability=Available
// [-2] void QVariant(QLatin1String)
func NewQVariant_16(string *QLatin1String /*123*/) *QVariant {
	var convArg0 = string.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2E13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:225
// index:17
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QStringList &)
func NewQVariant_17(stringlist *QStringList) *QVariant {
	var convArg0 = stringlist.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:226
// index:18
// Public Visibility=Default Availability=Available
// [-2] void QVariant(QChar)
func NewQVariant_18(qchar *QChar /*123*/) *QVariant {
	var convArg0 = qchar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2E5QChar", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:227
// index:19
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QDate &)
func NewQVariant_19(date *QDate) *QVariant {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QDate", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:228
// index:20
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QTime &)
func NewQVariant_20(time *QTime) *QVariant {
	var convArg0 = time.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QTime", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:229
// index:21
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QDateTime &)
func NewQVariant_21(datetime *QDateTime) *QVariant {
	var convArg0 = datetime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK9QDateTime", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:234
// index:22
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QSize &)
func NewQVariant_22(size *QSize) *QVariant {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QSize", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:235
// index:23
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QSizeF &)
func NewQVariant_23(size *QSizeF) *QVariant {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK6QSizeF", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:236
// index:24
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QPoint &)
func NewQVariant_24(pt *QPoint) *QVariant {
	var convArg0 = pt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:237
// index:25
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QPointF &)
func NewQVariant_25(pt *QPointF) *QVariant {
	var convArg0 = pt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:238
// index:26
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QLine &)
func NewQVariant_26(line *QLine) *QVariant {
	var convArg0 = line.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QLine", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:239
// index:27
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QLineF &)
func NewQVariant_27(line *QLineF) *QVariant {
	var convArg0 = line.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK6QLineF", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:240
// index:28
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRect &)
func NewQVariant_28(rect *QRect) *QVariant {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QRect", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:241
// index:29
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRectF &)
func NewQVariant_29(rect *QRectF) *QVariant {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK6QRectF", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:243
// index:30
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QLocale &)
func NewQVariant_30(locale *QLocale) *QVariant {
	var convArg0 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK7QLocale", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:245
// index:31
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRegExp &)
func NewQVariant_31(regExp *QRegExp) *QVariant {
	var convArg0 = regExp.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK7QRegExp", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:249
// index:32
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRegularExpression &)
func NewQVariant_32(re *QRegularExpression) *QVariant {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:251
// index:33
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QUrl &)
func NewQVariant_33(url *QUrl) *QVariant {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:252
// index:34
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QEasingCurve &)
func NewQVariant_34(easing *QEasingCurve) *QVariant {
	var convArg0 = easing.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK12QEasingCurve", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:253
// index:35
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QUuid &)
func NewQVariant_35(uuid *QUuid) *QVariant {
	var convArg0 = uuid.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK5QUuid", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:254
// index:36
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QModelIndex &)
func NewQVariant_36(modelIndex *QModelIndex) *QVariant {
	var convArg0 = modelIndex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK11QModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:255
// index:37
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QPersistentModelIndex &)
func NewQVariant_37(modelIndex *QPersistentModelIndex) *QVariant {
	var convArg0 = modelIndex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK21QPersistentModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:256
// index:38
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QJsonValue &)
func NewQVariant_38(jsonValue *QJsonValue) *QVariant {
	var convArg0 = jsonValue.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK10QJsonValue", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:257
// index:39
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QJsonObject &)
func NewQVariant_39(jsonObject *QJsonObject) *QVariant {
	var convArg0 = jsonObject.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK11QJsonObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:258
// index:40
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QJsonArray &)
func NewQVariant_40(jsonArray *QJsonArray) *QVariant {
	var convArg0 = jsonArray.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK10QJsonArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:259
// index:41
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QJsonDocument &)
func NewQVariant_41(jsonDocument *QJsonDocument) *QVariant {
	var convArg0 = jsonDocument.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK13QJsonDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QVariant()
func DeleteQVariant(this *QVariant) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qvariant.h:270
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QVariant &)
func (this *QVariant) Swap(other *QVariant) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:272
// index:0
// Public Visibility=Default Availability=Available
// [4] QVariant::Type type()
func (this *QVariant) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qvariant.h:273
// index:0
// Public Visibility=Default Availability=Available
// [4] int userType()
func (this *QVariant) UserType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8userTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:274
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * typeName()
func (this *QVariant) TypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8typeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qvariant.h:276
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canConvert(int)
func (this *QVariant) CanConvert(targetTypeId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10canConvertEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), targetTypeId)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:277
// index:0
// Public Visibility=Default Availability=Available
// [1] bool convert(int)
func (this *QVariant) Convert(targetTypeId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant7convertEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), targetTypeId)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:467
// index:1
// Protected Visibility=Default Availability=Available
// [1] bool convert(const int, void *)
func (this *QVariant) Convert_1(t int, ptr unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7convertEiPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t, ptr)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:279
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QVariant) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:280
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QVariant) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QVariant) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:284
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QVariant) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:285
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QVariant) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:287
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(_Bool *)
func (this *QVariant) ToInt(ok unsafe.Pointer /*666*/) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant5toIntEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:288
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(_Bool *)
func (this *QVariant) ToUInt(ok unsafe.Pointer /*666*/) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toUIntEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:289
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(_Bool *)
func (this *QVariant) ToLongLong(ok unsafe.Pointer /*666*/) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10toLongLongEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:290
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(_Bool *)
func (this *QVariant) ToULongLong(ok unsafe.Pointer /*666*/) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toULongLongEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qvariant.h:291
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toBool()
func (this *QVariant) ToBool() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toBoolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:292
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(_Bool *)
func (this *QVariant) ToDouble(ok unsafe.Pointer /*666*/) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:293
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(_Bool *)
func (this *QVariant) ToFloat(ok unsafe.Pointer /*666*/) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:294
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal toReal(_Bool *)
func (this *QVariant) ToReal(ok unsafe.Pointer /*666*/) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toRealEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qvariant.h:295
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toByteArray()
func (this *QVariant) ToByteArray() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toByteArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:296
// index:0
// Public Visibility=Default Availability=Available
// [8] QBitArray toBitArray()
func (this *QVariant) ToBitArray() *QBitArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10toBitArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitArray)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:297
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QVariant) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qvariant.h:299
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar toChar()
func (this *QVariant) ToChar() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:300
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate toDate()
func (this *QVariant) ToDate() *QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDate)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:301
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime toTime()
func (this *QVariant) ToTime() *QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTime)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:302
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime toDateTime()
func (this *QVariant) ToDateTime() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant10toDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:308
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint toPoint()
func (this *QVariant) ToPoint() *QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:309
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF toPointF()
func (this *QVariant) ToPointF() *QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toPointFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:310
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect toRect()
func (this *QVariant) ToRect() *QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRect)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:311
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize toSize()
func (this *QVariant) ToSize() *QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:312
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF toSizeF()
func (this *QVariant) ToSizeF() *QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toSizeFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:313
// index:0
// Public Visibility=Default Availability=Available
// [16] QLine toLine()
func (this *QVariant) ToLine() *QLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLine)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:314
// index:0
// Public Visibility=Default Availability=Available
// [32] QLineF toLineF()
func (this *QVariant) ToLineF() *QLineF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toLineFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLineFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLineF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:315
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF toRectF()
func (this *QVariant) ToRectF() *QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7toRectFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:317
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale toLocale()
func (this *QVariant) ToLocale() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toLocaleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegExp toRegExp()
func (this *QVariant) ToRegExp() *QRegExp /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant8toRegExpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegExp)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:323
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression toRegularExpression()
func (this *QVariant) ToRegularExpression() *QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant19toRegularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:325
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl toUrl()
func (this *QVariant) ToUrl() *QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant5toUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:326
// index:0
// Public Visibility=Default Availability=Available
// [8] QEasingCurve toEasingCurve()
func (this *QVariant) ToEasingCurve() *QEasingCurve /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant13toEasingCurveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQEasingCurve)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:327
// index:0
// Public Visibility=Default Availability=Available
// [16] QUuid toUuid()
func (this *QVariant) ToUuid() *QUuid /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant6toUuidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUuidFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUuid)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:328
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex toModelIndex()
func (this *QVariant) ToModelIndex() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant12toModelIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:329
// index:0
// Public Visibility=Default Availability=Available
// [8] QPersistentModelIndex toPersistentModelIndex()
func (this *QVariant) ToPersistentModelIndex() *QPersistentModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant22toPersistentModelIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:330
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue toJsonValue()
func (this *QVariant) ToJsonValue() *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toJsonValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:331
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject toJsonObject()
func (this *QVariant) ToJsonObject() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant12toJsonObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:332
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray toJsonArray()
func (this *QVariant) ToJsonArray() *QJsonArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant11toJsonArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:333
// index:0
// Public Visibility=Default Availability=Available
// [8] QJsonDocument toJsonDocument()
func (this *QVariant) ToJsonDocument() *QJsonDocument /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant14toJsonDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonDocument)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:337
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(QDataStream &)
func (this *QVariant) Load(ds *QDataStream) {
	var convArg0 = ds.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant4loadER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void save(QDataStream &)
func (this *QVariant) Save(ds *QDataStream) {
	var convArg0 = ds.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant4saveER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:340
// index:0
// Public static Visibility=Default Availability=Available
// [8] const char * typeToName(int)
func (this *QVariant) TypeToName(typeId int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant10typeToNameEi", qtrt.FFI_TYPE_POINTER, typeId)
	gopp.ErrPrint(err, rv)
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
func (this *QVariant) NameToType(name string) int {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant10nameToTypeEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
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
func (this *QVariant) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:345
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const void * data()
func (this *QVariant) Data_1() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:344
// index:0
// Public Visibility=Default Availability=Available
// [8] const void * constData()
func (this *QVariant) ConstData() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:464
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void create(int, const void *)
func (this *QVariant) Create(type_ int, copy unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariant6createEiPKv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, copy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:465
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool cmp(const QVariant &)
func (this *QVariant) Cmp(other *QVariant) bool {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant3cmpERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qvariant.h:466
// index:0
// Protected Visibility=Default Availability=Available
// [4] int compare(const QVariant &)
func (this *QVariant) Compare(other *QVariant) int {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant7compareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QVariant__Type = int

const QVariant__Invalid QVariant__Type = 0
const QVariant__Bool QVariant__Type = 1
const QVariant__Int QVariant__Type = 2
const QVariant__UInt QVariant__Type = 3
const QVariant__LongLong QVariant__Type = 4
const QVariant__ULongLong QVariant__Type = 5
const QVariant__Double QVariant__Type = 6
const QVariant__Char QVariant__Type = 7
const QVariant__Map QVariant__Type = 8
const QVariant__List QVariant__Type = 9
const QVariant__String QVariant__Type = 10
const QVariant__StringList QVariant__Type = 11
const QVariant__ByteArray QVariant__Type = 12
const QVariant__BitArray QVariant__Type = 13
const QVariant__Date QVariant__Type = 14
const QVariant__Time QVariant__Type = 15
const QVariant__DateTime QVariant__Type = 16
const QVariant__Url QVariant__Type = 17
const QVariant__Locale QVariant__Type = 18
const QVariant__Rect QVariant__Type = 19
const QVariant__RectF QVariant__Type = 20
const QVariant__Size QVariant__Type = 21
const QVariant__SizeF QVariant__Type = 22
const QVariant__Line QVariant__Type = 23
const QVariant__LineF QVariant__Type = 24
const QVariant__Point QVariant__Type = 25
const QVariant__PointF QVariant__Type = 26
const QVariant__RegExp QVariant__Type = 27
const QVariant__RegularExpression QVariant__Type = 44
const QVariant__Hash QVariant__Type = 28
const QVariant__EasingCurve QVariant__Type = 29
const QVariant__Uuid QVariant__Type = 30
const QVariant__ModelIndex QVariant__Type = 42
const QVariant__PersistentModelIndex QVariant__Type = 50
const QVariant__LastCoreType QVariant__Type = 51
const QVariant__Font QVariant__Type = 64
const QVariant__Pixmap QVariant__Type = 65
const QVariant__Brush QVariant__Type = 66
const QVariant__Color QVariant__Type = 67
const QVariant__Palette QVariant__Type = 68
const QVariant__Image QVariant__Type = 70
const QVariant__Polygon QVariant__Type = 71
const QVariant__Region QVariant__Type = 72
const QVariant__Bitmap QVariant__Type = 73
const QVariant__Cursor QVariant__Type = 74
const QVariant__KeySequence QVariant__Type = 75
const QVariant__Pen QVariant__Type = 76
const QVariant__TextLength QVariant__Type = 77
const QVariant__TextFormat QVariant__Type = 78
const QVariant__Matrix QVariant__Type = 79
const QVariant__Transform QVariant__Type = 80
const QVariant__Matrix4x4 QVariant__Type = 81
const QVariant__Vector2D QVariant__Type = 82
const QVariant__Vector3D QVariant__Type = 83
const QVariant__Vector4D QVariant__Type = 84
const QVariant__Quaternion QVariant__Type = 85
const QVariant__PolygonF QVariant__Type = 86
const QVariant__Icon QVariant__Type = 69
const QVariant__LastGuiType QVariant__Type = 86
const QVariant__SizePolicy QVariant__Type = 121
const QVariant__UserType QVariant__Type = 1024
const QVariant__LastType QVariant__Type = -1

//  body block end
