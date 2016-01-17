package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QVariantComparisonHelper::QVariantComparisonHelper(const QVariant & var);
extern void _ZN24QVariantComparisonHelperC2ERK8QVariant(void* qthis, void* arg0); // 1
  // proto:  void QVariant::load(QDataStream & ds);
extern void _ZN8QVariant4loadER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  int QVariant::toInt(bool * ok);
extern void _ZNK8QVariant5toIntEPb(void* qthis, bool* arg0); // 4
  // proto:  uint QVariant::toUInt(bool * ok);
extern void _ZNK8QVariant6toUIntEPb(void* qthis, bool* arg0); // 4
  // proto:  float QVariant::toFloat(bool * ok);
extern void _ZNK8QVariant7toFloatEPb(void* qthis, bool* arg0); // 4
  // proto:  QChar QVariant::toChar();
extern void _ZNK8QVariant6toCharEv(void* qthis); // 4
  // proto:  const char * QVariant::typeName();
extern void _ZNK8QVariant8typeNameEv(void* qthis); // 4
  // proto:  QRegularExpression QVariant::toRegularExpression();
extern void _ZNK8QVariant19toRegularExpressionEv(void* qthis); // 4
  // proto: static QVariant::Type QVariant::nameToType(const char * name);
extern void _ZN8QVariant10nameToTypeEPKc(unsigned char* arg0); // 4
  // proto:  QEasingCurve QVariant::toEasingCurve();
extern void _ZNK8QVariant13toEasingCurveEv(void* qthis); // 4
  // proto:  QList<QVariant> QVariant::toList();
extern void _ZNK8QVariant6toListEv(void* qthis); // 4
  // proto:  QMap<QString, QVariant> QVariant::toMap();
extern void _ZNK8QVariant5toMapEv(void* qthis); // 4
  // proto: static const char * QVariant::typeToName(int typeId);
extern void _ZN8QVariant10typeToNameEi(int32_t arg0); // 4
  // proto:  bool QVariant::toBool();
extern void _ZNK8QVariant6toBoolEv(void* qthis); // 4
  // proto:  QBitArray QVariant::toBitArray();
extern void _ZNK8QVariant10toBitArrayEv(void* qthis); // 4
  // proto:  int QVariant::userType();
extern void _ZNK8QVariant8userTypeEv(void* qthis); // 4
  // proto:  qlonglong QVariant::toLongLong(bool * ok);
extern void _ZNK8QVariant10toLongLongEPb(void* qthis, bool* arg0); // 4
  // proto:  const void * QVariant::constData();
extern void _ZNK8QVariant9constDataEv(void* qthis); // 4
  // proto:  QString QVariant::toString();
extern void _ZNK8QVariant8toStringEv(void* qthis); // 4
  // proto:  void QVariant::swap(QVariant & other);
extern void _ZN8QVariant4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QJsonObject QVariant::toJsonObject();
extern void _ZNK8QVariant12toJsonObjectEv(void* qthis); // 4
  // proto:  QVariant::Type QVariant::type();
extern void _ZNK8QVariant4typeEv(void* qthis); // 4
  // proto:  QSize QVariant::toSize();
extern void _ZNK8QVariant6toSizeEv(void* qthis); // 4
  // proto:  bool QVariant::isDetached();
extern void _ZNK8QVariant10isDetachedEv(void* qthis); // 2
  // proto:  QLocale QVariant::toLocale();
extern void _ZNK8QVariant8toLocaleEv(void* qthis); // 4
  // proto:  void QVariant::QVariant(const QRectF & rect);
extern void _ZN8QVariantC2ERK6QRectF(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(QChar qchar);
extern void _ZN8QVariantC2E5QChar(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const char * str);
extern void _ZN8QVariantC2EPKc(void* qthis, unsigned char* arg0); // 3
  // proto:  void QVariant::QVariant(const QPointF & pt);
extern void _ZN8QVariantC2ERK7QPointF(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QRect & rect);
extern void _ZN8QVariantC2ERK5QRect(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(int typeId, const void * copy);
extern void _ZN8QVariantC2EiPKv(void* qthis, int32_t arg0, void* arg1); // 3
  // proto:  void QVariant::QVariant(const QPoint & pt);
extern void _ZN8QVariantC2ERK6QPoint(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(uint ui);
extern void _ZN8QVariantC2Ej(void* qthis, int32_t arg0); // 3
  // proto:  void QVariant::QVariant(const QLocale & locale);
extern void _ZN8QVariantC2ERK7QLocale(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QByteArray & bytearray);
extern void _ZN8QVariantC2ERK10QByteArray(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(qlonglong ll);
extern void _ZN8QVariantC2Ex(void* qthis, int64_t arg0); // 3
  // proto:  void QVariant::QVariant(qulonglong ull);
extern void _ZN8QVariantC2Ey(void* qthis, int64_t arg0); // 3
  // proto:  void QVariant::QVariant();
extern void _ZN8QVariantC2Ev(void* qthis); // 1
  // proto:  void QVariant::QVariant(const QPersistentModelIndex & modelIndex);
extern void _ZN8QVariantC2ERK21QPersistentModelIndex(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(int typeId, const void * copy, uint flags);
extern void _ZN8QVariantC2EiPKvj(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QVariant::QVariant(const QLineF & line);
extern void _ZN8QVariantC2ERK6QLineF(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(int i);
extern void _ZN8QVariantC2Ei(void* qthis, int32_t arg0); // 3
  // proto:  void QVariant::QVariant(float f);
extern void _ZN8QVariantC2Ef(void* qthis, float arg0); // 3
  // proto:  void QVariant::QVariant(const QString & string);
extern void _ZN8QVariantC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(double d);
extern void _ZN8QVariantC2Ed(void* qthis, double arg0); // 3
  // proto:  void QVariant::QVariant(bool b);
extern void _ZN8QVariantC2Eb(void* qthis, bool arg0); // 3
  // proto:  void QVariant::QVariant(const QEasingCurve & easing);
extern void _ZN8QVariantC2ERK12QEasingCurve(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(QDataStream & s);
extern void _ZN8QVariantC2ER11QDataStream(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QSizeF & size);
extern void _ZN8QVariantC2ERK6QSizeF(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QStringList & stringlist);
extern void _ZN8QVariantC2ERK11QStringList(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QUuid & uuid);
extern void _ZN8QVariantC2ERK5QUuid(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QDate & date);
extern void _ZN8QVariantC2ERK5QDate(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QBitArray & bitarray);
extern void _ZN8QVariantC2ERK9QBitArray(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QModelIndex & modelIndex);
extern void _ZN8QVariantC2ERK11QModelIndex(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QUrl & url);
extern void _ZN8QVariantC2ERK4QUrl(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QVariant & other);
extern void _ZN8QVariantC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QSize & size);
extern void _ZN8QVariantC2ERK5QSize(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QLine & line);
extern void _ZN8QVariantC2ERK5QLine(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QRegularExpression & re);
extern void _ZN8QVariantC2ERK18QRegularExpression(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QTime & time);
extern void _ZN8QVariantC2ERK5QTime(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QDateTime & datetime);
extern void _ZN8QVariantC2ERK9QDateTime(void* qthis, void* arg0); // 3
  // proto:  void QVariant::QVariant(const QRegExp & regExp);
extern void _ZN8QVariantC2ERK7QRegExp(void* qthis, void* arg0); // 3
  // proto:  QDateTime QVariant::toDateTime();
extern void _ZNK8QVariant10toDateTimeEv(void* qthis); // 4
  // proto:  QJsonArray QVariant::toJsonArray();
extern void _ZNK8QVariant11toJsonArrayEv(void* qthis); // 4
  // proto:  bool QVariant::isValid();
extern void _ZNK8QVariant7isValidEv(void* qthis); // 2
  // proto:  QLineF QVariant::toLineF();
extern void _ZNK8QVariant7toLineFEv(void* qthis); // 4
  // proto:  QJsonValue QVariant::toJsonValue();
extern void _ZNK8QVariant11toJsonValueEv(void* qthis); // 4
  // proto:  bool QVariant::canConvert(int targetTypeId);
extern void _ZNK8QVariant10canConvertEi(void* qthis, int32_t arg0); // 4
  // proto:  QPersistentModelIndex QVariant::toPersistentModelIndex();
extern void _ZNK8QVariant22toPersistentModelIndexEv(void* qthis); // 4
  // proto:  QPointF QVariant::toPointF();
extern void _ZNK8QVariant8toPointFEv(void* qthis); // 4
  // proto:  QJsonDocument QVariant::toJsonDocument();
extern void _ZNK8QVariant14toJsonDocumentEv(void* qthis); // 4
  // proto:  QModelIndex QVariant::toModelIndex();
extern void _ZNK8QVariant12toModelIndexEv(void* qthis); // 4
  // proto:  void QVariant::detach();
extern void _ZN8QVariant6detachEv(void* qthis); // 4
  // proto:  void * QVariant::data();
extern void _ZN8QVariant4dataEv(void* qthis); // 4
  // proto:  qulonglong QVariant::toULongLong(bool * ok);
extern void _ZNK8QVariant11toULongLongEPb(void* qthis, bool* arg0); // 4
  // proto:  QPoint QVariant::toPoint();
extern void _ZNK8QVariant7toPointEv(void* qthis); // 4
  // proto:  QHash<QString, QVariant> QVariant::toHash();
extern void _ZNK8QVariant6toHashEv(void* qthis); // 4
  // proto:  bool QVariant::convert(int targetTypeId);
extern void _ZN8QVariant7convertEi(void* qthis, int32_t arg0); // 4
  // proto:  void QVariant::save(QDataStream & ds);
extern void _ZNK8QVariant4saveER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  QRegExp QVariant::toRegExp();
extern void _ZNK8QVariant8toRegExpEv(void* qthis); // 4
  // proto:  QLine QVariant::toLine();
extern void _ZNK8QVariant6toLineEv(void* qthis); // 4
  // proto:  QStringList QVariant::toStringList();
extern void _ZNK8QVariant12toStringListEv(void* qthis); // 4
  // proto:  QSizeF QVariant::toSizeF();
extern void _ZNK8QVariant7toSizeFEv(void* qthis); // 4
  // proto:  QRectF QVariant::toRectF();
extern void _ZNK8QVariant7toRectFEv(void* qthis); // 4
  // proto:  QUrl QVariant::toUrl();
extern void _ZNK8QVariant5toUrlEv(void* qthis); // 4
  // proto:  QUuid QVariant::toUuid();
extern void _ZNK8QVariant6toUuidEv(void* qthis); // 4
  // proto:  qreal QVariant::toReal(bool * ok);
extern void _ZNK8QVariant6toRealEPb(void* qthis, bool* arg0); // 4
  // proto:  void QVariant::~QVariant();
extern void _ZN8QVariantD2Ev(void* qthis); // 4
  // proto:  bool QVariant::isNull();
extern void _ZNK8QVariant6isNullEv(void* qthis); // 4
  // proto:  QByteArray QVariant::toByteArray();
extern void _ZNK8QVariant11toByteArrayEv(void* qthis); // 4
  // proto:  double QVariant::toDouble(bool * ok);
extern void _ZNK8QVariant8toDoubleEPb(void* qthis, bool* arg0); // 4
  // proto:  QTime QVariant::toTime();
extern void _ZNK8QVariant6toTimeEv(void* qthis); // 4
  // proto:  QDate QVariant::toDate();
extern void _ZNK8QVariant6toDateEv(void* qthis); // 4
  // proto:  void QVariant::clear();
extern void _ZN8QVariant5clearEv(void* qthis); // 4
  // proto:  QRect QVariant::toRect();
extern void _ZNK8QVariant6toRectEv(void* qthis); // 4
  // proto:  QSequentialIterable::const_iterator QSequentialIterable::begin();
extern void _ZNK19QSequentialIterable5beginEv(void* qthis); // 4
  // proto:  QSequentialIterable::const_iterator QSequentialIterable::end();
extern void _ZNK19QSequentialIterable3endEv(void* qthis); // 4
  // proto:  bool QSequentialIterable::canReverseIterate();
extern void _ZNK19QSequentialIterable17canReverseIterateEv(void* qthis); // 4
  // proto:  QVariant QSequentialIterable::at(int idx);
extern void _ZNK19QSequentialIterable2atEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSequentialIterable::size();
extern void _ZNK19QSequentialIterable4sizeEv(void* qthis); // 4
  // proto:  QAssociativeIterable::const_iterator QAssociativeIterable::begin();
extern void _ZNK20QAssociativeIterable5beginEv(void* qthis); // 4
  // proto:  QVariant QAssociativeIterable::value(const QVariant & key);
extern void _ZNK20QAssociativeIterable5valueERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  QAssociativeIterable::const_iterator QAssociativeIterable::end();
extern void _ZNK20QAssociativeIterable3endEv(void* qthis); // 4
  // proto:  QAssociativeIterable::const_iterator QAssociativeIterable::find(const QVariant & key);
extern void _ZNK20QAssociativeIterable4findERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  int QAssociativeIterable::size();
extern void _ZNK20QAssociativeIterable4sizeEv(void* qthis); // 4
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

// QVariantComparisonHelper(const class QVariant &)
func NewQVariantComparisonHelper(args ...interface{}) QVariantComparisonHelper {
  // QVariantComparisonHelper(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QVariantComparisonHelperC1ERK8QVariant
    // invoke: void QVariantComparisonHelper(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN24QVariantComparisonHelperC2ERK8QVariant(qthis, arg0)
  default:
    qtrt.ErrorResolve("QVariantComparisonHelper", "QVariantComparisonHelper", args)
  }

  return QVariantComparisonHelper{}
}

// load(class QDataStream &)
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

// toInt(_Bool *)
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

// toUInt(_Bool *)
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

// toFloat(_Bool *)
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

// toChar()
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

// typeName()
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

// toRegularExpression()
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

// nameToType(const char *)
func (this *QVariant) nameToType_s(args ...interface{}) () {
  // nameToType(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariant10nameToTypeEPKc
    // invoke: QVariant::Type nameToType(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN8QVariant10nameToTypeEPKc(arg0)
  default:
    qtrt.ErrorResolve("QVariant", "nameToType", args)
  }

}

// toEasingCurve()
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

// toList()
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

// toMap()
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

// typeToName(int)
func (this *QVariant) typeToName_s(args ...interface{}) () {
  // typeToName(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariant10typeToNameEi
    // invoke: const char * typeToName(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QVariant10typeToNameEi(arg0)
  default:
    qtrt.ErrorResolve("QVariant", "typeToName", args)
  }

}

// toBool()
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

// toBitArray()
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

// userType()
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

// toLongLong(_Bool *)
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

// constData()
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

// toString()
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

// swap(class QVariant &)
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
    C._ZN8QVariant4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "swap", args)
  }

}

// toJsonObject()
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

// type()
func (this *QVariant) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QVariant4typeEv
    // invoke: QVariant::Type type()
    C._ZNK8QVariant4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "type", args)
  }

}

// toSize()
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

// isDetached()
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
    C._ZNK8QVariant10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "isDetached", args)
  }

}

// toLocale()
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

// QVariant(const class QRectF &)
func NewQVariant(args ...interface{}) QVariant {
  // QVariant(const class QRectF &)
  // QVariant(class QChar)
  // QVariant(const char *)
  // QVariant(const class QPointF &)
  // QVariant(const class QRect &)
  // QVariant(int, const void *)
  // QVariant(const class QPoint &)
  // QVariant(uint)
  // QVariant(const class QLocale &)
  // QVariant(const class QByteArray &)
  // QVariant(qlonglong)
  // QVariant(qulonglong)
  // QVariant()
  // QVariant(const class QPersistentModelIndex &)
  // QVariant(int, const void *, uint)
  // QVariant(const class QLineF &)
  // QVariant(int)
  // QVariant(float)
  // QVariant(const class QString &)
  // QVariant(double)
  // QVariant(_Bool)
  // QVariant(const class QEasingCurve &)
  // QVariant(class QDataStream &)
  // QVariant(const class QSizeF &)
  // QVariant(const class QStringList &)
  // QVariant(const class QUuid &)
  // QVariant(const class QDate &)
  // QVariant(const class QBitArray &)
  // QVariant(const class QModelIndex &)
  // QVariant(const class QUrl &)
  // QVariant(const class QVariant &)
  // QVariant(const class QSize &)
  // QVariant(const class QLine &)
  // QVariant(const class QRegularExpression &)
  // QVariant(const class QTime &)
  // QVariant(const class QDateTime &)
  // QVariant(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.VoidpTy() // "const void *"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "uint"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = reflect.TypeOf(QPersistentModelIndex{}) // "const QPersistentModelIndex &"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.Int32Ty(false) // "int"
  vtys[14][1] = qtrt.VoidpTy() // "const void *"
  vtys[14][2] = qtrt.Int32Ty(false) // "uint"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.FloatTy(false) // "float"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = qtrt.DoubleTy(false) // "double"
  vtys[20] = make(map[int32]reflect.Type)
  vtys[20][0] = qtrt.BoolTy(false) // "bool"
  vtys[21] = make(map[int32]reflect.Type)
  vtys[21][0] = reflect.TypeOf(QEasingCurve{}) // "const QEasingCurve &"
  vtys[22] = make(map[int32]reflect.Type)
  vtys[22][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"
  vtys[23] = make(map[int32]reflect.Type)
  vtys[23][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[24] = make(map[int32]reflect.Type)
  vtys[24][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[25] = make(map[int32]reflect.Type)
  vtys[25][0] = reflect.TypeOf(QUuid{}) // "const QUuid &"
  vtys[26] = make(map[int32]reflect.Type)
  vtys[26][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[27] = make(map[int32]reflect.Type)
  vtys[27][0] = reflect.TypeOf(QBitArray{}) // "const QBitArray &"
  vtys[28] = make(map[int32]reflect.Type)
  vtys[28][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[29] = make(map[int32]reflect.Type)
  vtys[29][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"
  vtys[30] = make(map[int32]reflect.Type)
  vtys[30][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[31] = make(map[int32]reflect.Type)
  vtys[31][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[32] = make(map[int32]reflect.Type)
  vtys[32][0] = reflect.TypeOf(QLine{}) // "const QLine &"
  vtys[33] = make(map[int32]reflect.Type)
  vtys[33][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[34] = make(map[int32]reflect.Type)
  vtys[34][0] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[35] = make(map[int32]reflect.Type)
  vtys[35][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[36] = make(map[int32]reflect.Type)
  vtys[36][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariantC1ERK6QRectF
    // invoke: void QVariant(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK6QRectF(qthis, arg0)
  case 1:
    // invoke: _ZN8QVariantC1E5QChar
    // invoke: void QVariant(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2E5QChar(qthis, arg0)
  case 2:
    // invoke: _ZN8QVariantC1EPKc
    // invoke: void QVariant(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2EPKc(qthis, arg0)
  case 3:
    // invoke: _ZN8QVariantC1ERK7QPointF
    // invoke: void QVariant(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK7QPointF(qthis, arg0)
  case 4:
    // invoke: _ZN8QVariantC1ERK5QRect
    // invoke: void QVariant(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK5QRect(qthis, arg0)
  case 5:
    // invoke: _ZN8QVariantC1EiPKv
    // invoke: void QVariant(int, const void *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2EiPKv(qthis, arg0, arg1)
  case 6:
    // invoke: _ZN8QVariantC1ERK6QPoint
    // invoke: void QVariant(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK6QPoint(qthis, arg0)
  case 7:
    // invoke: _ZN8QVariantC1Ej
    // invoke: void QVariant(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2Ej(qthis, arg0)
  case 8:
    // invoke: _ZN8QVariantC1ERK7QLocale
    // invoke: void QVariant(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK7QLocale(qthis, arg0)
  case 9:
    // invoke: _ZN8QVariantC1ERK10QByteArray
    // invoke: void QVariant(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK10QByteArray(qthis, arg0)
  case 10:
    // invoke: _ZN8QVariantC1Ex
    // invoke: void QVariant(qlonglong)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2Ex(qthis, arg0)
  case 11:
    // invoke: _ZN8QVariantC1Ey
    // invoke: void QVariant(qulonglong)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2Ey(qthis, arg0)
  case 12:
    // invoke: _ZN8QVariantC1Ev
    // invoke: void QVariant()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2Ev(qthis)
  case 13:
    // invoke: _ZN8QVariantC1ERK21QPersistentModelIndex
    // invoke: void QVariant(const class QPersistentModelIndex &)
    var arg0 = args[0].(QPersistentModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK21QPersistentModelIndex(qthis, arg0)
  case 14:
    // invoke: _ZN8QVariantC1EiPKvj
    // invoke: void QVariant(int, const void *, uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2EiPKvj(qthis, arg0, arg1, arg2)
  case 15:
    // invoke: _ZN8QVariantC1ERK6QLineF
    // invoke: void QVariant(const class QLineF &)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK6QLineF(qthis, arg0)
  case 16:
    // invoke: _ZN8QVariantC1Ei
    // invoke: void QVariant(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2Ei(qthis, arg0)
  case 17:
    // invoke: _ZN8QVariantC1Ef
    // invoke: void QVariant(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2Ef(qthis, arg0)
  case 18:
    // invoke: _ZN8QVariantC1ERK7QString
    // invoke: void QVariant(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK7QString(qthis, arg0)
  case 19:
    // invoke: _ZN8QVariantC1Ed
    // invoke: void QVariant(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2Ed(qthis, arg0)
  case 20:
    // invoke: _ZN8QVariantC1Eb
    // invoke: void QVariant(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2Eb(qthis, arg0)
  case 21:
    // invoke: _ZN8QVariantC1ERK12QEasingCurve
    // invoke: void QVariant(const class QEasingCurve &)
    var arg0 = args[0].(QEasingCurve).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK12QEasingCurve(qthis, arg0)
  case 22:
    // invoke: _ZN8QVariantC1ER11QDataStream
    // invoke: void QVariant(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ER11QDataStream(qthis, arg0)
  case 23:
    // invoke: _ZN8QVariantC1ERK6QSizeF
    // invoke: void QVariant(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK6QSizeF(qthis, arg0)
  case 24:
    // invoke: _ZN8QVariantC1ERK11QStringList
    // invoke: void QVariant(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK11QStringList(qthis, arg0)
  case 25:
    // invoke: _ZN8QVariantC1ERK5QUuid
    // invoke: void QVariant(const class QUuid &)
    var arg0 = args[0].(QUuid).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK5QUuid(qthis, arg0)
  case 26:
    // invoke: _ZN8QVariantC1ERK5QDate
    // invoke: void QVariant(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK5QDate(qthis, arg0)
  case 27:
    // invoke: _ZN8QVariantC1ERK9QBitArray
    // invoke: void QVariant(const class QBitArray &)
    var arg0 = args[0].(QBitArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK9QBitArray(qthis, arg0)
  case 28:
    // invoke: _ZN8QVariantC1ERK11QModelIndex
    // invoke: void QVariant(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK11QModelIndex(qthis, arg0)
  case 29:
    // invoke: _ZN8QVariantC1ERK4QUrl
    // invoke: void QVariant(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK4QUrl(qthis, arg0)
  case 30:
    // invoke: _ZN8QVariantC1ERKS_
    // invoke: void QVariant(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERKS_(qthis, arg0)
  case 31:
    // invoke: _ZN8QVariantC1ERK5QSize
    // invoke: void QVariant(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK5QSize(qthis, arg0)
  case 32:
    // invoke: _ZN8QVariantC1ERK5QLine
    // invoke: void QVariant(const class QLine &)
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK5QLine(qthis, arg0)
  case 33:
    // invoke: _ZN8QVariantC1ERK18QRegularExpression
    // invoke: void QVariant(const class QRegularExpression &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK18QRegularExpression(qthis, arg0)
  case 34:
    // invoke: _ZN8QVariantC1ERK5QTime
    // invoke: void QVariant(const class QTime &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK5QTime(qthis, arg0)
  case 35:
    // invoke: _ZN8QVariantC1ERK9QDateTime
    // invoke: void QVariant(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK9QDateTime(qthis, arg0)
  case 36:
    // invoke: _ZN8QVariantC1ERK7QRegExp
    // invoke: void QVariant(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QVariantC2ERK7QRegExp(qthis, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "QVariant", args)
  }

  return QVariant{}
}

// toDateTime()
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

// toJsonArray()
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

// isValid()
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
    C._ZNK8QVariant7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "isValid", args)
  }

}

// toLineF()
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

// toJsonValue()
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

// canConvert(int)
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

// toPersistentModelIndex()
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

// toPointF()
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

// toJsonDocument()
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

// toModelIndex()
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

// detach()
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

// data()
func (this *QVariant) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

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

// toULongLong(_Bool *)
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

// toPoint()
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

// toHash()
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

// convert(int)
func (this *QVariant) convert(args ...interface{}) () {
  // convert(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

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

// save(class QDataStream &)
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

// toRegExp()
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

// toLine()
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

// toStringList()
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

// toSizeF()
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

// toRectF()
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

// toUrl()
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

// toUuid()
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

// toReal(_Bool *)
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

// ~QVariant()
func (this *QVariant) FreeQVariant(args ...interface{}) () {
  // ~QVariant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QVariantD0Ev
    // invoke: void ~QVariant()
    C._ZN8QVariantD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "~QVariant", args)
  }

}

// isNull()
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

// toByteArray()
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

// toDouble(_Bool *)
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

// toTime()
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

// toDate()
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

// clear()
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

// toRect()
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

// begin()
func (this *QSequentialIterable) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QSequentialIterable5beginEv
    // invoke: QSequentialIterable::const_iterator begin()
    C._ZNK19QSequentialIterable5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialIterable", "begin", args)
  }

}

// end()
func (this *QSequentialIterable) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QSequentialIterable3endEv
    // invoke: QSequentialIterable::const_iterator end()
    C._ZNK19QSequentialIterable3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialIterable", "end", args)
  }

}

// canReverseIterate()
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

// at(int)
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

// size()
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

// begin()
func (this *QAssociativeIterable) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAssociativeIterable5beginEv
    // invoke: QAssociativeIterable::const_iterator begin()
    C._ZNK20QAssociativeIterable5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "begin", args)
  }

}

// value(const class QVariant &)
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

// end()
func (this *QAssociativeIterable) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAssociativeIterable3endEv
    // invoke: QAssociativeIterable::const_iterator end()
    C._ZNK20QAssociativeIterable3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "end", args)
  }

}

// find(const class QVariant &)
func (this *QAssociativeIterable) find(args ...interface{}) () {
  // find(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAssociativeIterable4findERK8QVariant
    // invoke: QAssociativeIterable::const_iterator find(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK20QAssociativeIterable4findERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "find", args)
  }

}

// size()
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

// <= body block end

