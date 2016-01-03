package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qvariant.h
// dst-file: /src/core/qvariant.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QVariantComparisonHelper::QVariantComparisonHelper(const QVariant & var);
extern void* dector_ZN24QVariantComparisonHelperC1ERK8QVariant(void* arg0);
extern void demth_ZN24QVariantComparisonHelperC1ERK8QVariant(void* qthis, void* arg0);
  // proto:  double QVariant::toDouble(bool * ok);
extern void _ZNK8QVariant8toDoubleEPb(void* qthis, bool* arg0);
  // proto:  void QVariant::QVariant(const char * str);
extern void* dector_ZN8QVariantC1EPKc(unsigned char* arg0);
extern void _ZN8QVariantC1EPKc(void* qthis, unsigned char* arg0);
  // proto:  qlonglong QVariant::toLongLong(bool * ok);
extern void _ZNK8QVariant10toLongLongEPb(void* qthis, bool* arg0);
  // proto:  void QVariant::QVariant(const QPointF & pt);
extern void* dector_ZN8QVariantC1ERK7QPointF(void* arg0);
extern void _ZN8QVariantC1ERK7QPointF(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(const QPoint & pt);
extern void* dector_ZN8QVariantC1ERK6QPoint(void* arg0);
extern void _ZN8QVariantC1ERK6QPoint(void* qthis, void* arg0);
  // proto:  QSize QVariant::toSize();
extern void _ZNK8QVariant6toSizeEv(void* qthis);
  // proto:  QString QVariant::toString();
extern void _ZNK8QVariant8toStringEv(void* qthis);
  // proto:  qreal QVariant::toReal(bool * ok);
extern void _ZNK8QVariant6toRealEPb(void* qthis, bool* arg0);
  // proto:  float QVariant::toFloat(bool * ok);
extern void _ZNK8QVariant7toFloatEPb(void* qthis, bool* arg0);
  // proto:  void QVariant::QVariant(const QString & string);
extern void* dector_ZN8QVariantC1ERK7QString(void* arg0);
extern void _ZN8QVariantC1ERK7QString(void* qthis, void* arg0);
  // proto:  QByteArray QVariant::toByteArray();
extern void _ZNK8QVariant11toByteArrayEv(void* qthis);
  // proto:  QLocale QVariant::toLocale();
extern void _ZNK8QVariant8toLocaleEv(void* qthis);
  // proto:  QUrl QVariant::toUrl();
extern void _ZNK8QVariant5toUrlEv(void* qthis);
  // proto:  QLine QVariant::toLine();
extern void _ZNK8QVariant6toLineEv(void* qthis);
  // proto:  void QVariant::QVariant(const QSize & size);
extern void* dector_ZN8QVariantC1ERK5QSize(void* arg0);
extern void _ZN8QVariantC1ERK5QSize(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(const QLineF & line);
extern void* dector_ZN8QVariantC1ERK6QLineF(void* arg0);
extern void _ZN8QVariantC1ERK6QLineF(void* qthis, void* arg0);
  // proto:  const char * QVariant::typeName();
extern void _ZNK8QVariant8typeNameEv(void* qthis);
  // proto:  QJsonArray QVariant::toJsonArray();
extern void _ZNK8QVariant11toJsonArrayEv(void* qthis);
  // proto:  void QVariant::QVariant(const QLocale & locale);
extern void* dector_ZN8QVariantC1ERK7QLocale(void* arg0);
extern void _ZN8QVariantC1ERK7QLocale(void* qthis, void* arg0);
  // proto:  QStringList QVariant::toStringList();
extern void _ZNK8QVariant12toStringListEv(void* qthis);
  // proto:  QList<QVariant> QVariant::toList();
extern void _ZNK8QVariant6toListEv(void* qthis);
  // proto:  uint QVariant::toUInt(bool * ok);
extern void _ZNK8QVariant6toUIntEPb(void* qthis, bool* arg0);
  // proto:  QUuid QVariant::toUuid();
extern void _ZNK8QVariant6toUuidEv(void* qthis);
  // proto:  void QVariant::QVariant(const QPersistentModelIndex & modelIndex);
extern void* dector_ZN8QVariantC1ERK21QPersistentModelIndex(void* arg0);
extern void _ZN8QVariantC1ERK21QPersistentModelIndex(void* qthis, void* arg0);
  // proto:  QJsonDocument QVariant::toJsonDocument();
extern void _ZNK8QVariant14toJsonDocumentEv(void* qthis);
  // proto:  void QVariant::QVariant(QDataStream & s);
extern void* dector_ZN8QVariantC1ER11QDataStream(void* arg0);
extern void _ZN8QVariantC1ER11QDataStream(void* qthis, void* arg0);
  // proto:  QPoint QVariant::toPoint();
extern void _ZNK8QVariant7toPointEv(void* qthis);
  // proto:  int QVariant::toInt(bool * ok);
extern void _ZNK8QVariant5toIntEPb(void* qthis, bool* arg0);
  // proto:  bool QVariant::isValid();
extern void demth_ZNK8QVariant7isValidEv(void* qthis);
  // proto:  void QVariant::QVariant(const QUuid & uuid);
extern void* dector_ZN8QVariantC1ERK5QUuid(void* arg0);
extern void _ZN8QVariantC1ERK5QUuid(void* qthis, void* arg0);
  // proto:  void QVariant::detach();
extern void _ZN8QVariant6detachEv(void* qthis);
  // proto:  void QVariant::QVariant(const QRegExp & regExp);
extern void* dector_ZN8QVariantC1ERK7QRegExp(void* arg0);
extern void _ZN8QVariantC1ERK7QRegExp(void* qthis, void* arg0);
  // proto:  QModelIndex QVariant::toModelIndex();
extern void _ZNK8QVariant12toModelIndexEv(void* qthis);
  // proto:  QHash<QString, QVariant> QVariant::toHash();
extern void _ZNK8QVariant6toHashEv(void* qthis);
  // proto:  QMap<QString, QVariant> QVariant::toMap();
extern void _ZNK8QVariant5toMapEv(void* qthis);
  // proto:  bool QVariant::canConvert(int targetTypeId);
extern void _ZNK8QVariant10canConvertEi(void* qthis, int32_t arg0);
  // proto:  void QVariant::QVariant(const QRectF & rect);
extern void* dector_ZN8QVariantC1ERK6QRectF(void* arg0);
extern void _ZN8QVariantC1ERK6QRectF(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(const QRect & rect);
extern void* dector_ZN8QVariantC1ERK5QRect(void* arg0);
extern void _ZN8QVariantC1ERK5QRect(void* qthis, void* arg0);
  // proto:  int QVariant::userType();
extern void _ZNK8QVariant8userTypeEv(void* qthis);
  // proto:  const void * QVariant::constData();
extern void _ZNK8QVariant9constDataEv(void* qthis);
  // proto:  void QVariant::QVariant(void * );
extern void* dector_ZN8QVariantC1EPv(void* arg0);
extern void demth_ZN8QVariantC1EPv(void* qthis, void* arg0);
  // proto:  QPersistentModelIndex QVariant::toPersistentModelIndex();
extern void _ZNK8QVariant22toPersistentModelIndexEv(void* qthis);
  // proto:  void QVariant::QVariant(int typeId, const void * copy, uint flags);
extern void* dector_ZN8QVariantC1EiPKvj(int32_t arg0, void* arg1, int32_t arg2);
extern void _ZN8QVariantC1EiPKvj(void* qthis, int32_t arg0, void* arg1, int32_t arg2);
  // proto:  QLineF QVariant::toLineF();
extern void _ZNK8QVariant7toLineFEv(void* qthis);
  // proto:  QJsonObject QVariant::toJsonObject();
extern void _ZNK8QVariant12toJsonObjectEv(void* qthis);
  // proto:  void QVariant::load(QDataStream & ds);
extern void _ZN8QVariant4loadER11QDataStream(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(const QSizeF & size);
extern void* dector_ZN8QVariantC1ERK6QSizeF(void* arg0);
extern void _ZN8QVariantC1ERK6QSizeF(void* qthis, void* arg0);
  // proto:  QChar QVariant::toChar();
extern void _ZNK8QVariant6toCharEv(void* qthis);
  // proto:  bool QVariant::isNull();
extern void _ZNK8QVariant6isNullEv(void* qthis);
  // proto:  void QVariant::QVariant(const QDate & date);
extern void* dector_ZN8QVariantC1ERK5QDate(void* arg0);
extern void _ZN8QVariantC1ERK5QDate(void* qthis, void* arg0);
  // proto:  QRectF QVariant::toRectF();
extern void _ZNK8QVariant7toRectFEv(void* qthis);
  // proto:  void QVariant::QVariant(const QBitArray & bitarray);
extern void* dector_ZN8QVariantC1ERK9QBitArray(void* arg0);
extern void _ZN8QVariantC1ERK9QBitArray(void* qthis, void* arg0);
  // proto:  QDate QVariant::toDate();
extern void _ZNK8QVariant6toDateEv(void* qthis);
  // proto:  void QVariant::QVariant(const QModelIndex & modelIndex);
extern void* dector_ZN8QVariantC1ERK11QModelIndex(void* arg0);
extern void _ZN8QVariantC1ERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QVariant::~QVariant();
extern void _ZN8QVariantD0Ev(void* qthis);
  // proto:  void QVariant::save(QDataStream & ds);
extern void _ZNK8QVariant4saveER11QDataStream(void* qthis, void* arg0);
  // proto:  QTime QVariant::toTime();
extern void _ZNK8QVariant6toTimeEv(void* qthis);
  // proto:  void QVariant::QVariant(const QLine & line);
extern void* dector_ZN8QVariantC1ERK5QLine(void* arg0);
extern void _ZN8QVariantC1ERK5QLine(void* qthis, void* arg0);
  // proto:  void * QVariant::data();
extern void _ZN8QVariant4dataEv(void* qthis);
  // proto:  void QVariant::QVariant(const QTime & time);
extern void* dector_ZN8QVariantC1ERK5QTime(void* arg0);
extern void _ZN8QVariantC1ERK5QTime(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(const QDateTime & datetime);
extern void* dector_ZN8QVariantC1ERK9QDateTime(void* arg0);
extern void _ZN8QVariantC1ERK9QDateTime(void* qthis, void* arg0);
  // proto:  bool QVariant::convert(int targetTypeId);
extern void _ZN8QVariant7convertEi(void* qthis, int32_t arg0);
  // proto:  QRegExp QVariant::toRegExp();
extern void _ZNK8QVariant8toRegExpEv(void* qthis);
  // proto:  QPointF QVariant::toPointF();
extern void _ZNK8QVariant8toPointFEv(void* qthis);
  // proto:  void QVariant::QVariant(QChar qchar);
extern void* dector_ZN8QVariantC1E5QChar(void* arg0);
extern void _ZN8QVariantC1E5QChar(void* qthis, void* arg0);
  // proto: static const char * QVariant::typeToName(int typeId);
extern void _ZN8QVariant10typeToNameEi(int32_t arg0);
  // proto:  QSizeF QVariant::toSizeF();
extern void _ZNK8QVariant7toSizeFEv(void* qthis);
  // proto:  void QVariant::swap(QVariant & other);
extern void demth_ZN8QVariant4swapERS_(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(int typeId, const void * copy);
extern void* dector_ZN8QVariantC1EiPKv(int32_t arg0, void* arg1);
extern void _ZN8QVariantC1EiPKv(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QVariant::QVariant(const QEasingCurve & easing);
extern void* dector_ZN8QVariantC1ERK12QEasingCurve(void* arg0);
extern void _ZN8QVariantC1ERK12QEasingCurve(void* qthis, void* arg0);
  // proto:  void QVariant::clear();
extern void _ZN8QVariant5clearEv(void* qthis);
  // proto:  QRect QVariant::toRect();
extern void _ZNK8QVariant6toRectEv(void* qthis);
  // proto:  void QVariant::QVariant(const QByteArray & bytearray);
extern void* dector_ZN8QVariantC1ERK10QByteArray(void* arg0);
extern void _ZN8QVariantC1ERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(qlonglong ll);
extern void* dector_ZN8QVariantC1Ex(int64_t arg0);
extern void _ZN8QVariantC1Ex(void* qthis, int64_t arg0);
  // proto:  void QVariant::QVariant(qulonglong ull);
extern void* dector_ZN8QVariantC1Ey(int64_t arg0);
extern void _ZN8QVariantC1Ey(void* qthis, int64_t arg0);
  // proto:  void QVariant::QVariant();
extern void* dector_ZN8QVariantC1Ev();
extern void _ZN8QVariantC1Ev(void* qthis);
  // proto:  bool QVariant::toBool();
extern void _ZNK8QVariant6toBoolEv(void* qthis);
  // proto:  void QVariant::QVariant(uint ui);
extern void* dector_ZN8QVariantC1Ej(int32_t arg0);
extern void _ZN8QVariantC1Ej(void* qthis, int32_t arg0);
  // proto:  void QVariant::QVariant(int i);
extern void* dector_ZN8QVariantC1Ei(int32_t arg0);
extern void _ZN8QVariantC1Ei(void* qthis, int32_t arg0);
  // proto:  void QVariant::QVariant(float f);
extern void* dector_ZN8QVariantC1Ef(float arg0);
extern void _ZN8QVariantC1Ef(void* qthis, float arg0);
  // proto:  void QVariant::QVariant(double d);
extern void* dector_ZN8QVariantC1Ed(double arg0);
extern void _ZN8QVariantC1Ed(void* qthis, double arg0);
  // proto:  void QVariant::QVariant(bool b);
extern void* dector_ZN8QVariantC1Eb(bool arg0);
extern void _ZN8QVariantC1Eb(void* qthis, bool arg0);
  // proto:  qulonglong QVariant::toULongLong(bool * ok);
extern void _ZNK8QVariant11toULongLongEPb(void* qthis, bool* arg0);
  // proto:  QJsonValue QVariant::toJsonValue();
extern void _ZNK8QVariant11toJsonValueEv(void* qthis);
  // proto:  QDateTime QVariant::toDateTime();
extern void _ZNK8QVariant10toDateTimeEv(void* qthis);
  // proto:  bool QVariant::isDetached();
extern void demth_ZNK8QVariant10isDetachedEv(void* qthis);
  // proto:  QEasingCurve QVariant::toEasingCurve();
extern void _ZNK8QVariant13toEasingCurveEv(void* qthis);
  // proto:  void QVariant::QVariant(const QUrl & url);
extern void* dector_ZN8QVariantC1ERK4QUrl(void* arg0);
extern void _ZN8QVariantC1ERK4QUrl(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(const QVariant & other);
extern void* dector_ZN8QVariantC1ERKS_(void* arg0);
extern void _ZN8QVariantC1ERKS_(void* qthis, void* arg0);
  // proto:  QBitArray QVariant::toBitArray();
extern void _ZNK8QVariant10toBitArrayEv(void* qthis);
  // proto:  QRegularExpression QVariant::toRegularExpression();
extern void _ZNK8QVariant19toRegularExpressionEv(void* qthis);
  // proto:  void QVariant::QVariant(const QRegularExpression & re);
extern void* dector_ZN8QVariantC1ERK18QRegularExpression(void* arg0);
extern void _ZN8QVariantC1ERK18QRegularExpression(void* qthis, void* arg0);
  // proto:  void QVariant::QVariant(const QStringList & stringlist);
extern void* dector_ZN8QVariantC1ERK11QStringList(void* arg0);
extern void _ZN8QVariantC1ERK11QStringList(void* qthis, void* arg0);
  // proto:  int QSequentialIterable::size();
extern void _ZNK19QSequentialIterable4sizeEv(void* qthis);
  // proto:  bool QSequentialIterable::canReverseIterate();
extern void _ZNK19QSequentialIterable17canReverseIterateEv(void* qthis);
  // proto:  QVariant QSequentialIterable::at(int idx);
extern void _ZNK19QSequentialIterable2atEi(void* qthis, int32_t arg0);
  // proto:  int QAssociativeIterable::size();
extern void _ZNK20QAssociativeIterable4sizeEv(void* qthis);
  // proto:  QVariant QAssociativeIterable::value(const QVariant & key);
extern void _ZNK20QAssociativeIterable5valueERK8QVariant(void* qthis, void* arg0);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QVariantComparisonHelper)=8
type QVariantComparisonHelper struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QVariant)=16
type QVariant struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QSequentialIterable)=104
type QSequentialIterable struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAssociativeIterable)=112
type QAssociativeIterable struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QVariantComparisonHelper::QVariantComparisonHelper(const QVariant & var);
func NewQVariantComparisonHelper(args ...interface{}) QVariantComparisonHelper {
  return QVariantComparisonHelper{}
}

  // proto:  double QVariant::toDouble(bool * ok);
func (this *QVariant) toDouble(args ...interface{}) () {
  // toDouble(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant8toDoubleEPb
    // invoke: double toDouble(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant8toDoubleEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "toDouble", args)
  }

}

  // proto:  void QVariant::QVariant(const char * str);
func NewQVariant(args ...interface{}) QVariant {
  return QVariant{}
}

  // proto:  qlonglong QVariant::toLongLong(bool * ok);
func (this *QVariant) toLongLong(args ...interface{}) () {
  // toLongLong(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant10toLongLongEPb
    // invoke: qlonglong toLongLong(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant10toLongLongEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "toLongLong", args)
  }

}

  // proto:  QSize QVariant::toSize();
func (this *QVariant) toSize(args ...interface{}) () {
  // toSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toSizeEv
    // invoke: QSize toSize()
    C._ZNK8QVariant6toSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toSize", args)
  }

}

  // proto:  QString QVariant::toString();
func (this *QVariant) toString(args ...interface{}) () {
  // toString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant8toStringEv
    // invoke: QString toString()
    C._ZNK8QVariant8toStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toString", args)
  }

}

  // proto:  qreal QVariant::toReal(bool * ok);
func (this *QVariant) toReal(args ...interface{}) () {
  // toReal(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toRealEPb
    // invoke: qreal toReal(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant6toRealEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "toReal", args)
  }

}

  // proto:  float QVariant::toFloat(bool * ok);
func (this *QVariant) toFloat(args ...interface{}) () {
  // toFloat(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant7toFloatEPb
    // invoke: float toFloat(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant7toFloatEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "toFloat", args)
  }

}

  // proto:  QByteArray QVariant::toByteArray();
func (this *QVariant) toByteArray(args ...interface{}) () {
  // toByteArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant11toByteArrayEv
    // invoke: QByteArray toByteArray()
    C._ZNK8QVariant11toByteArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toByteArray", args)
  }

}

  // proto:  QLocale QVariant::toLocale();
func (this *QVariant) toLocale(args ...interface{}) () {
  // toLocale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant8toLocaleEv
    // invoke: QLocale toLocale()
    C._ZNK8QVariant8toLocaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toLocale", args)
  }

}

  // proto:  QUrl QVariant::toUrl();
func (this *QVariant) toUrl(args ...interface{}) () {
  // toUrl()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant5toUrlEv
    // invoke: QUrl toUrl()
    C._ZNK8QVariant5toUrlEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toUrl", args)
  }

}

  // proto:  QLine QVariant::toLine();
func (this *QVariant) toLine(args ...interface{}) () {
  // toLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toLineEv
    // invoke: QLine toLine()
    C._ZNK8QVariant6toLineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toLine", args)
  }

}

  // proto:  const char * QVariant::typeName();
func (this *QVariant) typeName(args ...interface{}) () {
  // typeName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant8typeNameEv
    // invoke: const char * typeName()
    C._ZNK8QVariant8typeNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "typeName", args)
  }

}

  // proto:  QJsonArray QVariant::toJsonArray();
func (this *QVariant) toJsonArray(args ...interface{}) () {
  // toJsonArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant11toJsonArrayEv
    // invoke: QJsonArray toJsonArray()
    C._ZNK8QVariant11toJsonArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toJsonArray", args)
  }

}

  // proto:  QStringList QVariant::toStringList();
func (this *QVariant) toStringList(args ...interface{}) () {
  // toStringList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant12toStringListEv
    // invoke: QStringList toStringList()
    C._ZNK8QVariant12toStringListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toStringList", args)
  }

}

  // proto:  QList<QVariant> QVariant::toList();
func (this *QVariant) toList(args ...interface{}) () {
  // toList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toListEv
    // invoke: QList<QVariant> toList()
    C._ZNK8QVariant6toListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toList", args)
  }

}

  // proto:  uint QVariant::toUInt(bool * ok);
func (this *QVariant) toUInt(args ...interface{}) () {
  // toUInt(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toUIntEPb
    // invoke: uint toUInt(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant6toUIntEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "toUInt", args)
  }

}

  // proto:  QUuid QVariant::toUuid();
func (this *QVariant) toUuid(args ...interface{}) () {
  // toUuid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toUuidEv
    // invoke: QUuid toUuid()
    C._ZNK8QVariant6toUuidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toUuid", args)
  }

}

  // proto:  QJsonDocument QVariant::toJsonDocument();
func (this *QVariant) toJsonDocument(args ...interface{}) () {
  // toJsonDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant14toJsonDocumentEv
    // invoke: QJsonDocument toJsonDocument()
    C._ZNK8QVariant14toJsonDocumentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toJsonDocument", args)
  }

}

  // proto:  QPoint QVariant::toPoint();
func (this *QVariant) toPoint(args ...interface{}) () {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant7toPointEv
    // invoke: QPoint toPoint()
    C._ZNK8QVariant7toPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toPoint", args)
  }

}

  // proto:  int QVariant::toInt(bool * ok);
func (this *QVariant) toInt(args ...interface{}) () {
  // toInt(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant5toIntEPb
    // invoke: int toInt(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant5toIntEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "toInt", args)
  }

}

  // proto:  bool QVariant::isValid();
func (this *QVariant) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant7isValidEv
    // invoke: bool isValid()
    C.demth_ZNK8QVariant7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "isValid", args)
  }

}

  // proto:  void QVariant::detach();
func (this *QVariant) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariant6detachEv
    // invoke: void detach()
    C._ZN8QVariant6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "detach", args)
  }

}

  // proto:  QModelIndex QVariant::toModelIndex();
func (this *QVariant) toModelIndex(args ...interface{}) () {
  // toModelIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant12toModelIndexEv
    // invoke: QModelIndex toModelIndex()
    C._ZNK8QVariant12toModelIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toModelIndex", args)
  }

}

  // proto:  QHash<QString, QVariant> QVariant::toHash();
func (this *QVariant) toHash(args ...interface{}) () {
  // toHash()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toHashEv
    // invoke: QHash<QString, QVariant> toHash()
    C._ZNK8QVariant6toHashEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toHash", args)
  }

}

  // proto:  QMap<QString, QVariant> QVariant::toMap();
func (this *QVariant) toMap(args ...interface{}) () {
  // toMap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant5toMapEv
    // invoke: QMap<QString, QVariant> toMap()
    C._ZNK8QVariant5toMapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toMap", args)
  }

}

  // proto:  bool QVariant::canConvert(int targetTypeId);
func (this *QVariant) canConvert(args ...interface{}) () {
  // canConvert(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant10canConvertEi
    // invoke: bool canConvert(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant10canConvertEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "canConvert", args)
  }

}

  // proto:  int QVariant::userType();
func (this *QVariant) userType(args ...interface{}) () {
  // userType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant8userTypeEv
    // invoke: int userType()
    C._ZNK8QVariant8userTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "userType", args)
  }

}

  // proto:  const void * QVariant::constData();
func (this *QVariant) constData(args ...interface{}) () {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant9constDataEv
    // invoke: const void * constData()
    C._ZNK8QVariant9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "constData", args)
  }

}

  // proto:  QPersistentModelIndex QVariant::toPersistentModelIndex();
func (this *QVariant) toPersistentModelIndex(args ...interface{}) () {
  // toPersistentModelIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant22toPersistentModelIndexEv
    // invoke: QPersistentModelIndex toPersistentModelIndex()
    C._ZNK8QVariant22toPersistentModelIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toPersistentModelIndex", args)
  }

}

  // proto:  QLineF QVariant::toLineF();
func (this *QVariant) toLineF(args ...interface{}) () {
  // toLineF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant7toLineFEv
    // invoke: QLineF toLineF()
    C._ZNK8QVariant7toLineFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toLineF", args)
  }

}

  // proto:  QJsonObject QVariant::toJsonObject();
func (this *QVariant) toJsonObject(args ...interface{}) () {
  // toJsonObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant12toJsonObjectEv
    // invoke: QJsonObject toJsonObject()
    C._ZNK8QVariant12toJsonObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toJsonObject", args)
  }

}

  // proto:  void QVariant::load(QDataStream & ds);
func (this *QVariant) load(args ...interface{}) () {
  // load(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariant4loadER11QDataStream
    // invoke: void load(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QVariant4loadER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "load", args)
  }

}

  // proto:  QChar QVariant::toChar();
func (this *QVariant) toChar(args ...interface{}) () {
  // toChar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toCharEv
    // invoke: QChar toChar()
    C._ZNK8QVariant6toCharEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toChar", args)
  }

}

  // proto:  bool QVariant::isNull();
func (this *QVariant) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6isNullEv
    // invoke: bool isNull()
    C._ZNK8QVariant6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "isNull", args)
  }

}

  // proto:  QRectF QVariant::toRectF();
func (this *QVariant) toRectF(args ...interface{}) () {
  // toRectF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant7toRectFEv
    // invoke: QRectF toRectF()
    C._ZNK8QVariant7toRectFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toRectF", args)
  }

}

  // proto:  QDate QVariant::toDate();
func (this *QVariant) toDate(args ...interface{}) () {
  // toDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toDateEv
    // invoke: QDate toDate()
    C._ZNK8QVariant6toDateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toDate", args)
  }

}

  // proto:  void QVariant::~QVariant();
func (this *QVariant) FreeQVariant(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVariant", "~QVariant", args)
  }

}

  // proto:  void QVariant::save(QDataStream & ds);
func (this *QVariant) save(args ...interface{}) () {
  // save(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant4saveER11QDataStream
    // invoke: void save(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant4saveER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "save", args)
  }

}

  // proto:  QTime QVariant::toTime();
func (this *QVariant) toTime(args ...interface{}) () {
  // toTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toTimeEv
    // invoke: QTime toTime()
    C._ZNK8QVariant6toTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toTime", args)
  }

}

  // proto:  void * QVariant::data();
func (this *QVariant) data(args ...interface{}) () {
  // data()
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariant4dataEv
    // invoke: void * data()
    C._ZN8QVariant4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "data", args)
  }

}

  // proto:  bool QVariant::convert(int targetTypeId);
func (this *QVariant) convert(args ...interface{}) () {
  // convert(const int, void *)
  // convert(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "const int"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariant7convertEi
    // invoke: bool convert(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QVariant7convertEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "convert", args)
  }

}

  // proto:  QRegExp QVariant::toRegExp();
func (this *QVariant) toRegExp(args ...interface{}) () {
  // toRegExp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant8toRegExpEv
    // invoke: QRegExp toRegExp()
    C._ZNK8QVariant8toRegExpEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toRegExp", args)
  }

}

  // proto:  QPointF QVariant::toPointF();
func (this *QVariant) toPointF(args ...interface{}) () {
  // toPointF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant8toPointFEv
    // invoke: QPointF toPointF()
    C._ZNK8QVariant8toPointFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toPointF", args)
  }

}

  // proto: static const char * QVariant::typeToName(int typeId);
func (this *QVariant) typeToName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QVariant", "typeToName", args)
  }

}

  // proto:  QSizeF QVariant::toSizeF();
func (this *QVariant) toSizeF(args ...interface{}) () {
  // toSizeF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant7toSizeFEv
    // invoke: QSizeF toSizeF()
    C._ZNK8QVariant7toSizeFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toSizeF", args)
  }

}

  // proto:  void QVariant::swap(QVariant & other);
func (this *QVariant) swap(args ...interface{}) () {
  // swap(class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariant4swapERS_
    // invoke: void swap(class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QVariant4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "swap", args)
  }

}

  // proto:  void QVariant::clear();
func (this *QVariant) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariant5clearEv
    // invoke: void clear()
    C._ZN8QVariant5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "clear", args)
  }

}

  // proto:  QRect QVariant::toRect();
func (this *QVariant) toRect(args ...interface{}) () {
  // toRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toRectEv
    // invoke: QRect toRect()
    C._ZNK8QVariant6toRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toRect", args)
  }

}

  // proto:  bool QVariant::toBool();
func (this *QVariant) toBool(args ...interface{}) () {
  // toBool()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant6toBoolEv
    // invoke: bool toBool()
    C._ZNK8QVariant6toBoolEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toBool", args)
  }

}

  // proto:  qulonglong QVariant::toULongLong(bool * ok);
func (this *QVariant) toULongLong(args ...interface{}) () {
  // toULongLong(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant11toULongLongEPb
    // invoke: qulonglong toULongLong(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK8QVariant11toULongLongEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "toULongLong", args)
  }

}

  // proto:  QJsonValue QVariant::toJsonValue();
func (this *QVariant) toJsonValue(args ...interface{}) () {
  // toJsonValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant11toJsonValueEv
    // invoke: QJsonValue toJsonValue()
    C._ZNK8QVariant11toJsonValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toJsonValue", args)
  }

}

  // proto:  QDateTime QVariant::toDateTime();
func (this *QVariant) toDateTime(args ...interface{}) () {
  // toDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant10toDateTimeEv
    // invoke: QDateTime toDateTime()
    C._ZNK8QVariant10toDateTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toDateTime", args)
  }

}

  // proto:  bool QVariant::isDetached();
func (this *QVariant) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant10isDetachedEv
    // invoke: bool isDetached()
    C.demth_ZNK8QVariant10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "isDetached", args)
  }

}

  // proto:  QEasingCurve QVariant::toEasingCurve();
func (this *QVariant) toEasingCurve(args ...interface{}) () {
  // toEasingCurve()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant13toEasingCurveEv
    // invoke: QEasingCurve toEasingCurve()
    C._ZNK8QVariant13toEasingCurveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toEasingCurve", args)
  }

}

  // proto:  QBitArray QVariant::toBitArray();
func (this *QVariant) toBitArray(args ...interface{}) () {
  // toBitArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant10toBitArrayEv
    // invoke: QBitArray toBitArray()
    C._ZNK8QVariant10toBitArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toBitArray", args)
  }

}

  // proto:  QRegularExpression QVariant::toRegularExpression();
func (this *QVariant) toRegularExpression(args ...interface{}) () {
  // toRegularExpression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant19toRegularExpressionEv
    // invoke: QRegularExpression toRegularExpression()
    C._ZNK8QVariant19toRegularExpressionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toRegularExpression", args)
  }

}

  // proto:  int QSequentialIterable::size();
func (this *QSequentialIterable) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QSequentialIterable4sizeEv
    // invoke: int size()
    C._ZNK19QSequentialIterable4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialIterable", "size", args)
  }

}

  // proto:  bool QSequentialIterable::canReverseIterate();
func (this *QSequentialIterable) canReverseIterate(args ...interface{}) () {
  // canReverseIterate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QSequentialIterable17canReverseIterateEv
    // invoke: bool canReverseIterate()
    C._ZNK19QSequentialIterable17canReverseIterateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialIterable", "canReverseIterate", args)
  }

}

  // proto:  QVariant QSequentialIterable::at(int idx);
func (this *QSequentialIterable) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QSequentialIterable2atEi
    // invoke: QVariant at(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK19QSequentialIterable2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSequentialIterable", "at", args)
  }

}

  // proto:  int QAssociativeIterable::size();
func (this *QAssociativeIterable) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAssociativeIterable4sizeEv
    // invoke: int size()
    C._ZNK20QAssociativeIterable4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "size", args)
  }

}

  // proto:  QVariant QAssociativeIterable::value(const QVariant & key);
func (this *QAssociativeIterable) value(args ...interface{}) () {
  // value(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAssociativeIterable5valueERK8QVariant
    // invoke: QVariant value(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK20QAssociativeIterable5valueERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "value", args)
  }

}

// <= body block end

