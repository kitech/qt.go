//  header block begin
// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QVariant struct {
	*qtrt.CObject
}

func (this *QVariant) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQVariantFromPointer(cthis unsafe.Pointer) *QVariant {
	return &QVariant{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qvariant.h:199
// index:0
// Public inline
// void QVariant()
func NewQVariant() *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:201
// index:1
// Public
// void QVariant(enum QVariant::Type)
func NewQVariant_1(type_ int) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ENS_4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:202
// index:2
// Public
// void QVariant(int, const void *)
func NewQVariant_2(typeId int, copy unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2EiPKv", ffiqt.FFI_TYPE_VOID, cthis, &typeId, copy)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:203
// index:3
// Public
// void QVariant(int, const void *, uint)
func NewQVariant_3(typeId int, copy unsafe.Pointer, flags uint) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2EiPKvj", ffiqt.FFI_TYPE_VOID, cthis, &typeId, copy, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:207
// index:4
// Public
// void QVariant(class QDataStream &)
func NewQVariant_4(s *QDataStream) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ER11QDataStream", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:210
// index:5
// Public
// void QVariant(int)
func NewQVariant_5(i int) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &i)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:211
// index:6
// Public
// void QVariant(uint)
func NewQVariant_6(ui uint) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ej", ffiqt.FFI_TYPE_VOID, cthis, &ui)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:212
// index:7
// Public
// void QVariant(qlonglong)
func NewQVariant_7(ll int64) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ex", ffiqt.FFI_TYPE_VOID, cthis, &ll)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:213
// index:8
// Public
// void QVariant(qulonglong)
func NewQVariant_8(ull uint64) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ey", ffiqt.FFI_TYPE_VOID, cthis, &ull)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:214
// index:9
// Public
// void QVariant(_Bool)
func NewQVariant_9(b bool) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Eb", ffiqt.FFI_TYPE_VOID, cthis, &b)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:215
// index:10
// Public
// void QVariant(double)
func NewQVariant_10(d float64) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ed", ffiqt.FFI_TYPE_VOID, cthis, &d)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:216
// index:11
// Public
// void QVariant(float)
func NewQVariant_11(f float32) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ef", ffiqt.FFI_TYPE_VOID, cthis, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:218
// index:12
// Public
// void QVariant(const char *)
func NewQVariant_12(str string) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:221
// index:13
// Public
// void QVariant(const class QByteArray &)
func NewQVariant_13(bytearray *QByteArray) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = bytearray.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:222
// index:14
// Public
// void QVariant(const class QBitArray &)
func NewQVariant_14(bitarray *QBitArray) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = bitarray.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK9QBitArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:223
// index:15
// Public
// void QVariant(const class QString &)
func NewQVariant_15(string *QString) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:224
// index:16
// Public
// void QVariant(class QLatin1String)
func NewQVariant_16(string *QLatin1String) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2E13QLatin1String", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:225
// index:17
// Public
// void QVariant(const class QStringList &)
func NewQVariant_17(stringlist *QStringList) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = stringlist.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK11QStringList", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:226
// index:18
// Public
// void QVariant(class QChar)
func NewQVariant_18(qchar *QChar) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = qchar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2E5QChar", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:227
// index:19
// Public
// void QVariant(const class QDate &)
func NewQVariant_19(date *QDate) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = date.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QDate", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:228
// index:20
// Public
// void QVariant(const class QTime &)
func NewQVariant_20(time *QTime) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = time.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QTime", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:229
// index:21
// Public
// void QVariant(const class QDateTime &)
func NewQVariant_21(datetime *QDateTime) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = datetime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK9QDateTime", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:234
// index:22
// Public
// void QVariant(const class QSize &)
func NewQVariant_22(size *QSize) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QSize", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:235
// index:23
// Public
// void QVariant(const class QSizeF &)
func NewQVariant_23(size *QSizeF) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK6QSizeF", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:236
// index:24
// Public
// void QVariant(const class QPoint &)
func NewQVariant_24(pt *QPoint) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = pt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:237
// index:25
// Public
// void QVariant(const class QPointF &)
func NewQVariant_25(pt *QPointF) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = pt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:238
// index:26
// Public
// void QVariant(const class QLine &)
func NewQVariant_26(line *QLine) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = line.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QLine", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:239
// index:27
// Public
// void QVariant(const class QLineF &)
func NewQVariant_27(line *QLineF) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = line.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK6QLineF", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:240
// index:28
// Public
// void QVariant(const class QRect &)
func NewQVariant_28(rect *QRect) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QRect", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:241
// index:29
// Public
// void QVariant(const class QRectF &)
func NewQVariant_29(rect *QRectF) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK6QRectF", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:243
// index:30
// Public
// void QVariant(const class QLocale &)
func NewQVariant_30(locale *QLocale) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK7QLocale", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:245
// index:31
// Public
// void QVariant(const class QRegExp &)
func NewQVariant_31(regExp *QRegExp) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = regExp.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK7QRegExp", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:249
// index:32
// Public
// void QVariant(const class QRegularExpression &)
func NewQVariant_32(re *QRegularExpression) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK18QRegularExpression", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:251
// index:33
// Public
// void QVariant(const class QUrl &)
func NewQVariant_33(url *QUrl) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK4QUrl", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:252
// index:34
// Public
// void QVariant(const class QEasingCurve &)
func NewQVariant_34(easing *QEasingCurve) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = easing.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK12QEasingCurve", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:253
// index:35
// Public
// void QVariant(const class QUuid &)
func NewQVariant_35(uuid *QUuid) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = uuid.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QUuid", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:254
// index:36
// Public
// void QVariant(const class QModelIndex &)
func NewQVariant_36(modelIndex *QModelIndex) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = modelIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK11QModelIndex", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:255
// index:37
// Public
// void QVariant(const class QPersistentModelIndex &)
func NewQVariant_37(modelIndex *QPersistentModelIndex) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = modelIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK21QPersistentModelIndex", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:256
// index:38
// Public
// void QVariant(const class QJsonValue &)
func NewQVariant_38(jsonValue *QJsonValue) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = jsonValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK10QJsonValue", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:257
// index:39
// Public
// void QVariant(const class QJsonObject &)
func NewQVariant_39(jsonObject *QJsonObject) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = jsonObject.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK11QJsonObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:258
// index:40
// Public
// void QVariant(const class QJsonArray &)
func NewQVariant_40(jsonArray *QJsonArray) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = jsonArray.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK10QJsonArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:259
// index:41
// Public
// void QVariant(const class QJsonDocument &)
func NewQVariant_41(jsonDocument *QJsonDocument) *QVariant {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = jsonDocument.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK13QJsonDocument", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:200
// index:0
// Public
// void ~QVariant()
func DeleteQVariant(*QVariant) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:270
// index:0
// Public inline
// void swap(class QVariant &)
func (this *QVariant) Swap(other *QVariant) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:272
// index:0
// Public
// QVariant::Type type()
func (this *QVariant) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:273
// index:0
// Public
// int userType()
func (this *QVariant) UserType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8userTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:274
// index:0
// Public
// const char * typeName()
func (this *QVariant) TypeName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8typeNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:276
// index:0
// Public
// bool canConvert(int)
func (this *QVariant) CanConvert(targetTypeId int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10canConvertEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &targetTypeId)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:277
// index:0
// Public
// bool convert(int)
func (this *QVariant) Convert(targetTypeId int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant7convertEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &targetTypeId)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:467
// index:1
// Protected
// bool convert(const int, void *)
func (this *QVariant) Convert_1(t int, ptr unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7convertEiPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t, ptr)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:279
// index:0
// Public inline
// bool isValid()
func (this *QVariant) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:280
// index:0
// Public
// bool isNull()
func (this *QVariant) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:282
// index:0
// Public
// void clear()
func (this *QVariant) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:284
// index:0
// Public
// void detach()
func (this *QVariant) Detach() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:285
// index:0
// Public inline
// bool isDetached()
func (this *QVariant) IsDetached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:287
// index:0
// Public
// int toInt(_Bool *)
func (this *QVariant) ToInt(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant5toIntEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:288
// index:0
// Public
// uint toUInt(_Bool *)
func (this *QVariant) ToUInt(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toUIntEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:289
// index:0
// Public
// qlonglong toLongLong(_Bool *)
func (this *QVariant) ToLongLong(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10toLongLongEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:290
// index:0
// Public
// qulonglong toULongLong(_Bool *)
func (this *QVariant) ToULongLong(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant11toULongLongEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:291
// index:0
// Public
// bool toBool()
func (this *QVariant) ToBool() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toBoolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:292
// index:0
// Public
// double toDouble(_Bool *)
func (this *QVariant) ToDouble(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toDoubleEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:293
// index:0
// Public
// float toFloat(_Bool *)
func (this *QVariant) ToFloat(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toFloatEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:294
// index:0
// Public
// qreal toReal(_Bool *)
func (this *QVariant) ToReal(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toRealEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:295
// index:0
// Public
// QByteArray toByteArray()
func (this *QVariant) ToByteArray() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant11toByteArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:296
// index:0
// Public
// QBitArray toBitArray()
func (this *QVariant) ToBitArray() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10toBitArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:297
// index:0
// Public
// QString toString()
func (this *QVariant) ToString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:298
// index:0
// Public
// QStringList toStringList()
func (this *QVariant) ToStringList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant12toStringListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:299
// index:0
// Public
// QChar toChar()
func (this *QVariant) ToChar() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toCharEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:300
// index:0
// Public
// QDate toDate()
func (this *QVariant) ToDate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toDateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:301
// index:0
// Public
// QTime toTime()
func (this *QVariant) ToTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:302
// index:0
// Public
// QDateTime toDateTime()
func (this *QVariant) ToDateTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10toDateTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:303
// index:0
// Public
// QList<QVariant> toList()
func (this *QVariant) ToList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:304
// index:0
// Public
// QMap<QString, QVariant> toMap()
func (this *QVariant) ToMap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant5toMapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:305
// index:0
// Public
// QHash<QString, QVariant> toHash()
func (this *QVariant) ToHash() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toHashEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:308
// index:0
// Public
// QPoint toPoint()
func (this *QVariant) ToPoint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:309
// index:0
// Public
// QPointF toPointF()
func (this *QVariant) ToPointF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toPointFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:310
// index:0
// Public
// QRect toRect()
func (this *QVariant) ToRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:311
// index:0
// Public
// QSize toSize()
func (this *QVariant) ToSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:312
// index:0
// Public
// QSizeF toSizeF()
func (this *QVariant) ToSizeF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toSizeFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:313
// index:0
// Public
// QLine toLine()
func (this *QVariant) ToLine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:314
// index:0
// Public
// QLineF toLineF()
func (this *QVariant) ToLineF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toLineFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:315
// index:0
// Public
// QRectF toRectF()
func (this *QVariant) ToRectF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toRectFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:317
// index:0
// Public
// QLocale toLocale()
func (this *QVariant) ToLocale() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toLocaleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:319
// index:0
// Public
// QRegExp toRegExp()
func (this *QVariant) ToRegExp() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toRegExpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:323
// index:0
// Public
// QRegularExpression toRegularExpression()
func (this *QVariant) ToRegularExpression() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant19toRegularExpressionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:325
// index:0
// Public
// QUrl toUrl()
func (this *QVariant) ToUrl() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant5toUrlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:326
// index:0
// Public
// QEasingCurve toEasingCurve()
func (this *QVariant) ToEasingCurve() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant13toEasingCurveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:327
// index:0
// Public
// QUuid toUuid()
func (this *QVariant) ToUuid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toUuidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:328
// index:0
// Public
// QModelIndex toModelIndex()
func (this *QVariant) ToModelIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant12toModelIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:329
// index:0
// Public
// QPersistentModelIndex toPersistentModelIndex()
func (this *QVariant) ToPersistentModelIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant22toPersistentModelIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:330
// index:0
// Public
// QJsonValue toJsonValue()
func (this *QVariant) ToJsonValue() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant11toJsonValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:331
// index:0
// Public
// QJsonObject toJsonObject()
func (this *QVariant) ToJsonObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant12toJsonObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:332
// index:0
// Public
// QJsonArray toJsonArray()
func (this *QVariant) ToJsonArray() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant11toJsonArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:333
// index:0
// Public
// QJsonDocument toJsonDocument()
func (this *QVariant) ToJsonDocument() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant14toJsonDocumentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:337
// index:0
// Public
// void load(class QDataStream &)
func (this *QVariant) Load(ds *QDataStream) {
	var convArg0 = ds.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant4loadER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:338
// index:0
// Public
// void save(class QDataStream &)
func (this *QVariant) Save(ds *QDataStream) {
	var convArg0 = ds.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant4saveER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:340
// index:0
// Public static
// const char * typeToName(int)
func (this *QVariant) TypeToName(typeId int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant10typeToNameEi", ffiqt.FFI_TYPE_POINTER, typeId)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVariant_TypeToName(typeId int) {
	var nilthis *QVariant
	nilthis.TypeToName(typeId)
}

// /usr/include/qt/QtCore/qvariant.h:341
// index:0
// Public static
// QVariant::Type nameToType(const char *)
func (this *QVariant) NameToType(name string) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant10nameToTypeEPKc", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVariant_NameToType(name string) {
	var nilthis *QVariant
	nilthis.NameToType(name)
}

// /usr/include/qt/QtCore/qvariant.h:343
// index:0
// Public
// void * data()
func (this *QVariant) Data() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:345
// index:1
// Public inline
// const void * data()
func (this *QVariant) Data_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:344
// index:0
// Public
// const void * constData()
func (this *QVariant) ConstData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant9constDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:464
// index:0
// Protected
// void create(int, const void *)
func (this *QVariant) Create(type_ int, copy unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant6createEiPKv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_, copy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:465
// index:0
// Protected
// bool cmp(const class QVariant &)
func (this *QVariant) Cmp(other *QVariant) interface{} {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant3cmpERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qvariant.h:466
// index:0
// Protected
// int compare(const class QVariant &)
func (this *QVariant) Compare(other *QVariant) interface{} {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7compareERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
