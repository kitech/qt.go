//  header block begin
// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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

// /usr/include/qt/QtCore/qvariant.h:199
// index:0
// inline
// void QVariant()
func NewQVariant() *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}
func NewQVariantFromPointer(cthis unsafe.Pointer) *QVariant {
	return &QVariant{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qvariant.h:201
// index:1
// void QVariant(enum QVariant::Type)
func NewQVariant_1(type_ int) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ENS_4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:202
// index:2
// void QVariant(int, const void *)
func NewQVariant_2(typeId int, copy unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2EiPKv", ffiqt.FFI_TYPE_VOID, cthis, &typeId, copy)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:203
// index:3
// void QVariant(int, const void *, uint)
func NewQVariant_3(typeId int, copy unsafe.Pointer, flags uint) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2EiPKvj", ffiqt.FFI_TYPE_VOID, cthis, &typeId, copy, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:207
// index:4
// void QVariant(class QDataStream &)
func NewQVariant_4(s unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ER11QDataStream", ffiqt.FFI_TYPE_VOID, cthis, s)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:210
// index:5
// void QVariant(int)
func NewQVariant_5(i int) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &i)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:211
// index:6
// void QVariant(uint)
func NewQVariant_6(ui uint) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ej", ffiqt.FFI_TYPE_VOID, cthis, &ui)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:212
// index:7
// void QVariant(qlonglong)
func NewQVariant_7(ll int64) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ex", ffiqt.FFI_TYPE_VOID, cthis, &ll)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:213
// index:8
// void QVariant(qulonglong)
func NewQVariant_8(ull uint64) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ey", ffiqt.FFI_TYPE_VOID, cthis, &ull)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:214
// index:9
// void QVariant(_Bool)
func NewQVariant_9(b bool) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Eb", ffiqt.FFI_TYPE_VOID, cthis, &b)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:215
// index:10
// void QVariant(double)
func NewQVariant_10(d float64) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ed", ffiqt.FFI_TYPE_VOID, cthis, &d)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:216
// index:11
// void QVariant(float)
func NewQVariant_11(f float32) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2Ef", ffiqt.FFI_TYPE_VOID, cthis, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:218
// index:12
// void QVariant(const char *)
func NewQVariant_12(str unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, str)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:221
// index:13
// void QVariant(const class QByteArray &)
func NewQVariant_13(bytearray unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, bytearray)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:222
// index:14
// void QVariant(const class QBitArray &)
func NewQVariant_14(bitarray unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK9QBitArray", ffiqt.FFI_TYPE_VOID, cthis, bitarray)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:223
// index:15
// void QVariant(const class QString &)
func NewQVariant_15(string unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, string)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:224
// index:16
// void QVariant(class QLatin1String)
func NewQVariant_16(string unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2E13QLatin1String", ffiqt.FFI_TYPE_VOID, cthis, string)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:225
// index:17
// void QVariant(const class QStringList &)
func NewQVariant_17(stringlist unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK11QStringList", ffiqt.FFI_TYPE_VOID, cthis, stringlist)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:226
// index:18
// void QVariant(class QChar)
func NewQVariant_18(qchar unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2E5QChar", ffiqt.FFI_TYPE_VOID, cthis, qchar)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:227
// index:19
// void QVariant(const class QDate &)
func NewQVariant_19(date unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QDate", ffiqt.FFI_TYPE_VOID, cthis, date)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:228
// index:20
// void QVariant(const class QTime &)
func NewQVariant_20(time unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QTime", ffiqt.FFI_TYPE_VOID, cthis, time)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:229
// index:21
// void QVariant(const class QDateTime &)
func NewQVariant_21(datetime unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK9QDateTime", ffiqt.FFI_TYPE_VOID, cthis, datetime)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:234
// index:22
// void QVariant(const class QSize &)
func NewQVariant_22(size unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QSize", ffiqt.FFI_TYPE_VOID, cthis, size)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:235
// index:23
// void QVariant(const class QSizeF &)
func NewQVariant_23(size unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK6QSizeF", ffiqt.FFI_TYPE_VOID, cthis, size)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:236
// index:24
// void QVariant(const class QPoint &)
func NewQVariant_24(pt unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, pt)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:237
// index:25
// void QVariant(const class QPointF &)
func NewQVariant_25(pt unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, pt)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:238
// index:26
// void QVariant(const class QLine &)
func NewQVariant_26(line unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QLine", ffiqt.FFI_TYPE_VOID, cthis, line)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:239
// index:27
// void QVariant(const class QLineF &)
func NewQVariant_27(line unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK6QLineF", ffiqt.FFI_TYPE_VOID, cthis, line)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:240
// index:28
// void QVariant(const class QRect &)
func NewQVariant_28(rect unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QRect", ffiqt.FFI_TYPE_VOID, cthis, rect)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:241
// index:29
// void QVariant(const class QRectF &)
func NewQVariant_29(rect unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK6QRectF", ffiqt.FFI_TYPE_VOID, cthis, rect)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:243
// index:30
// void QVariant(const class QLocale &)
func NewQVariant_30(locale unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK7QLocale", ffiqt.FFI_TYPE_VOID, cthis, locale)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:245
// index:31
// void QVariant(const class QRegExp &)
func NewQVariant_31(regExp unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK7QRegExp", ffiqt.FFI_TYPE_VOID, cthis, regExp)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:249
// index:32
// void QVariant(const class QRegularExpression &)
func NewQVariant_32(re unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK18QRegularExpression", ffiqt.FFI_TYPE_VOID, cthis, re)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:251
// index:33
// void QVariant(const class QUrl &)
func NewQVariant_33(url unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK4QUrl", ffiqt.FFI_TYPE_VOID, cthis, url)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:252
// index:34
// void QVariant(const class QEasingCurve &)
func NewQVariant_34(easing unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK12QEasingCurve", ffiqt.FFI_TYPE_VOID, cthis, easing)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:253
// index:35
// void QVariant(const class QUuid &)
func NewQVariant_35(uuid unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK5QUuid", ffiqt.FFI_TYPE_VOID, cthis, uuid)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:254
// index:36
// void QVariant(const class QModelIndex &)
func NewQVariant_36(modelIndex unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK11QModelIndex", ffiqt.FFI_TYPE_VOID, cthis, modelIndex)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:255
// index:37
// void QVariant(const class QPersistentModelIndex &)
func NewQVariant_37(modelIndex unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK21QPersistentModelIndex", ffiqt.FFI_TYPE_VOID, cthis, modelIndex)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:256
// index:38
// void QVariant(const class QJsonValue &)
func NewQVariant_38(jsonValue unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK10QJsonValue", ffiqt.FFI_TYPE_VOID, cthis, jsonValue)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:257
// index:39
// void QVariant(const class QJsonObject &)
func NewQVariant_39(jsonObject unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK11QJsonObject", ffiqt.FFI_TYPE_VOID, cthis, jsonObject)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:258
// index:40
// void QVariant(const class QJsonArray &)
func NewQVariant_40(jsonArray unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK10QJsonArray", ffiqt.FFI_TYPE_VOID, cthis, jsonArray)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:259
// index:41
// void QVariant(const class QJsonDocument &)
func NewQVariant_41(jsonDocument unsafe.Pointer) *QVariant {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantC2ERK13QJsonDocument", ffiqt.FFI_TYPE_VOID, cthis, jsonDocument)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:200
// index:0
// void ~QVariant()
func DeleteQVariant(*QVariant) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariantD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:270
// index:0
// inline
// void swap(class QVariant &)
func (this *QVariant) Swap(other unsafe.Pointer) {
	// 0: (, other QVariant &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:272
// index:0
// QVariant::Type type()
func (this *QVariant) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:273
// index:0
// int userType()
func (this *QVariant) UserType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8userTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:274
// index:0
// const char * typeName()
func (this *QVariant) TypeName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8typeNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:276
// index:0
// bool canConvert(int)
func (this *QVariant) CanConvert(targetTypeId int) {
	// 0: (, targetTypeId int), (&targetTypeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10canConvertEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &targetTypeId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:277
// index:0
// bool convert(int)
func (this *QVariant) Convert(targetTypeId int) {
	// 0: (, targetTypeId int), (&targetTypeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant7convertEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &targetTypeId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:467
// index:1
// bool convert(const int, void *)
func (this *QVariant) Convert_1(t int, ptr unsafe.Pointer) {
	// 1: (, t const int, ptr void *), (&t, ptr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7convertEiPv", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &t, ptr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:279
// index:0
// inline
// bool isValid()
func (this *QVariant) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:280
// index:0
// bool isNull()
func (this *QVariant) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:282
// index:0
// void clear()
func (this *QVariant) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:284
// index:0
// void detach()
func (this *QVariant) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant6detachEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:285
// index:0
// inline
// bool isDetached()
func (this *QVariant) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:287
// index:0
// int toInt(_Bool *)
func (this *QVariant) ToInt(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant5toIntEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:288
// index:0
// uint toUInt(_Bool *)
func (this *QVariant) ToUInt(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toUIntEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:289
// index:0
// qlonglong toLongLong(_Bool *)
func (this *QVariant) ToLongLong(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10toLongLongEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:290
// index:0
// qulonglong toULongLong(_Bool *)
func (this *QVariant) ToULongLong(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant11toULongLongEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:291
// index:0
// bool toBool()
func (this *QVariant) ToBool() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toBoolEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:292
// index:0
// double toDouble(_Bool *)
func (this *QVariant) ToDouble(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toDoubleEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:293
// index:0
// float toFloat(_Bool *)
func (this *QVariant) ToFloat(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toFloatEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:294
// index:0
// qreal toReal(_Bool *)
func (this *QVariant) ToReal(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toRealEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:295
// index:0
// QByteArray toByteArray()
func (this *QVariant) ToByteArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant11toByteArrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:296
// index:0
// QBitArray toBitArray()
func (this *QVariant) ToBitArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10toBitArrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:297
// index:0
// QString toString()
func (this *QVariant) ToString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:298
// index:0
// QStringList toStringList()
func (this *QVariant) ToStringList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant12toStringListEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:299
// index:0
// QChar toChar()
func (this *QVariant) ToChar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toCharEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:300
// index:0
// QDate toDate()
func (this *QVariant) ToDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toDateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:301
// index:0
// QTime toTime()
func (this *QVariant) ToTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:302
// index:0
// QDateTime toDateTime()
func (this *QVariant) ToDateTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant10toDateTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:303
// index:0
// QList<QVariant> toList()
func (this *QVariant) ToList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toListEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:304
// index:0
// QMap<QString, QVariant> toMap()
func (this *QVariant) ToMap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant5toMapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:305
// index:0
// QHash<QString, QVariant> toHash()
func (this *QVariant) ToHash() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toHashEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:308
// index:0
// QPoint toPoint()
func (this *QVariant) ToPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toPointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:309
// index:0
// QPointF toPointF()
func (this *QVariant) ToPointF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toPointFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:310
// index:0
// QRect toRect()
func (this *QVariant) ToRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:311
// index:0
// QSize toSize()
func (this *QVariant) ToSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:312
// index:0
// QSizeF toSizeF()
func (this *QVariant) ToSizeF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toSizeFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:313
// index:0
// QLine toLine()
func (this *QVariant) ToLine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toLineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:314
// index:0
// QLineF toLineF()
func (this *QVariant) ToLineF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toLineFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:315
// index:0
// QRectF toRectF()
func (this *QVariant) ToRectF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7toRectFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:317
// index:0
// QLocale toLocale()
func (this *QVariant) ToLocale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toLocaleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:319
// index:0
// QRegExp toRegExp()
func (this *QVariant) ToRegExp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant8toRegExpEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:323
// index:0
// QRegularExpression toRegularExpression()
func (this *QVariant) ToRegularExpression() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant19toRegularExpressionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:325
// index:0
// QUrl toUrl()
func (this *QVariant) ToUrl() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant5toUrlEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:326
// index:0
// QEasingCurve toEasingCurve()
func (this *QVariant) ToEasingCurve() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant13toEasingCurveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:327
// index:0
// QUuid toUuid()
func (this *QVariant) ToUuid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant6toUuidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:328
// index:0
// QModelIndex toModelIndex()
func (this *QVariant) ToModelIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant12toModelIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:329
// index:0
// QPersistentModelIndex toPersistentModelIndex()
func (this *QVariant) ToPersistentModelIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant22toPersistentModelIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:330
// index:0
// QJsonValue toJsonValue()
func (this *QVariant) ToJsonValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant11toJsonValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:331
// index:0
// QJsonObject toJsonObject()
func (this *QVariant) ToJsonObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant12toJsonObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:332
// index:0
// QJsonArray toJsonArray()
func (this *QVariant) ToJsonArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant11toJsonArrayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:333
// index:0
// QJsonDocument toJsonDocument()
func (this *QVariant) ToJsonDocument() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant14toJsonDocumentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:337
// index:0
// void load(class QDataStream &)
func (this *QVariant) Load(ds unsafe.Pointer) {
	// 0: (, ds QDataStream &), (ds)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant4loadER11QDataStream", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ds)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:338
// index:0
// void save(class QDataStream &)
func (this *QVariant) Save(ds unsafe.Pointer) {
	// 0: (, ds QDataStream &), (ds)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant4saveER11QDataStream", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ds)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:340
// index:0
// static
// const char * typeToName(int)
func (this *QVariant) TypeToName(typeId int) {
	// 0: (typeId int), (typeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant10typeToNameEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVariant_TypeToName(typeId int) {
	// 0: (typeId int), (typeId)
	var nilthis *QVariant
	nilthis.TypeToName(typeId)
}

// /usr/include/qt/QtCore/qvariant.h:341
// index:0
// static
// QVariant::Type nameToType(const char *)
func (this *QVariant) NameToType(name unsafe.Pointer) {
	// 0: (name const char *), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant10nameToTypeEPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QVariant_NameToType(name unsafe.Pointer) {
	// 0: (name const char *), (name)
	var nilthis *QVariant
	nilthis.NameToType(name)
}

// /usr/include/qt/QtCore/qvariant.h:343
// index:0
// void * data()
func (this *QVariant) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant4dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:345
// index:1
// inline
// const void * data()
func (this *QVariant) Data_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant4dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:344
// index:0
// const void * constData()
func (this *QVariant) ConstData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant9constDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:464
// index:0
// void create(int, const void *)
func (this *QVariant) Create(type_ int, copy unsafe.Pointer) {
	// 0: (, type int, copy const void *), (&type_, copy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QVariant6createEiPKv", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_, copy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:465
// index:0
// bool cmp(const class QVariant &)
func (this *QVariant) Cmp(other unsafe.Pointer) {
	// 0: (, other const QVariant &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant3cmpERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qvariant.h:466
// index:0
// int compare(const class QVariant &)
func (this *QVariant) Compare(other unsafe.Pointer) {
	// 0: (, other const QVariant &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QVariant7compareERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

//  body block end
