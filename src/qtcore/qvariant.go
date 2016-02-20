package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern void* C_ZN24QVariantComparisonHelperC2ERK8QVariant(void* arg0); // 1
  // proto:  void QVariant::load(QDataStream & ds);
extern void C_ZN8QVariant4loadER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  int QVariant::toInt(bool * ok);
extern int32_t C_ZNK8QVariant5toIntEPb(void* qthis, void* arg0); // 4
  // proto:  uint QVariant::toUInt(bool * ok);
extern int32_t C_ZNK8QVariant6toUIntEPb(void* qthis, void* arg0); // 4
  // proto:  float QVariant::toFloat(bool * ok);
extern float C_ZNK8QVariant7toFloatEPb(void* qthis, void* arg0); // 4
  // proto:  QChar QVariant::toChar();
extern void* C_ZNK8QVariant6toCharEv(void* qthis); // 4
  // proto:  const char * QVariant::typeName();
extern void* C_ZNK8QVariant8typeNameEv(void* qthis); // 4
  // proto:  QRegularExpression QVariant::toRegularExpression();
extern void* C_ZNK8QVariant19toRegularExpressionEv(void* qthis); // 4
  // proto: static QVariant::Type QVariant::nameToType(const char * name);
extern void C_ZN8QVariant10nameToTypeEPKc(void* arg0); // 4
  // proto:  QEasingCurve QVariant::toEasingCurve();
extern void* C_ZNK8QVariant13toEasingCurveEv(void* qthis); // 4
  // proto:  QList<QVariant> QVariant::toList();
extern void C_ZNK8QVariant6toListEv(void* qthis); // 4
  // proto:  QMap<QString, QVariant> QVariant::toMap();
extern void C_ZNK8QVariant5toMapEv(void* qthis); // 4
  // proto: static const char * QVariant::typeToName(int typeId);
extern void* C_ZN8QVariant10typeToNameEi(int32_t arg0); // 4
  // proto:  bool QVariant::toBool();
extern bool C_ZNK8QVariant6toBoolEv(void* qthis); // 4
  // proto:  QBitArray QVariant::toBitArray();
extern void* C_ZNK8QVariant10toBitArrayEv(void* qthis); // 4
  // proto:  int QVariant::userType();
extern int32_t C_ZNK8QVariant8userTypeEv(void* qthis); // 4
  // proto:  qlonglong QVariant::toLongLong(bool * ok);
extern int64_t C_ZNK8QVariant10toLongLongEPb(void* qthis, void* arg0); // 4
  // proto:  const void * QVariant::constData();
extern void C_ZNK8QVariant9constDataEv(void* qthis); // 4
  // proto:  QString QVariant::toString();
extern void* C_ZNK8QVariant8toStringEv(void* qthis); // 4
  // proto:  void QVariant::swap(QVariant & other);
extern void C_ZN8QVariant4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QJsonObject QVariant::toJsonObject();
extern void C_ZNK8QVariant12toJsonObjectEv(void* qthis); // 4
  // proto:  QVariant::Type QVariant::type();
extern void C_ZNK8QVariant4typeEv(void* qthis); // 4
  // proto:  QSize QVariant::toSize();
extern void* C_ZNK8QVariant6toSizeEv(void* qthis); // 4
  // proto:  bool QVariant::isDetached();
extern bool C_ZNK8QVariant10isDetachedEv(void* qthis); // 2
  // proto:  QLocale QVariant::toLocale();
extern void* C_ZNK8QVariant8toLocaleEv(void* qthis); // 4
  // proto:  void QVariant::QVariant(const QRectF & rect);
extern void* C_ZN8QVariantC2ERK6QRectF(void* arg0); // 3
  // proto:  void QVariant::QVariant(QChar qchar);
extern void* C_ZN8QVariantC2E5QChar(void* arg0); // 3
  // proto:  void QVariant::QVariant(const char * str);
extern void* C_ZN8QVariantC2EPKc(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QPointF & pt);
extern void* C_ZN8QVariantC2ERK7QPointF(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QRect & rect);
extern void* C_ZN8QVariantC2ERK5QRect(void* arg0); // 3
  // proto:  void QVariant::QVariant(int typeId, const void * copy);
extern void* C_ZN8QVariantC2EiPKv(int32_t arg0, void* arg1); // 3
  // proto:  void QVariant::QVariant(const QPoint & pt);
extern void* C_ZN8QVariantC2ERK6QPoint(void* arg0); // 3
  // proto:  void QVariant::QVariant(uint ui);
extern void* C_ZN8QVariantC2Ej(int32_t arg0); // 3
  // proto:  void QVariant::QVariant(const QLocale & locale);
extern void* C_ZN8QVariantC2ERK7QLocale(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QByteArray & bytearray);
extern void* C_ZN8QVariantC2ERK10QByteArray(void* arg0); // 3
  // proto:  void QVariant::QVariant(qlonglong ll);
extern void* C_ZN8QVariantC2Ex(int64_t arg0); // 3
  // proto:  void QVariant::QVariant(qulonglong ull);
extern void* C_ZN8QVariantC2Ey(int64_t arg0); // 3
  // proto:  void QVariant::QVariant();
extern void* C_ZN8QVariantC2Ev(); // 1
  // proto:  void QVariant::QVariant(const QPersistentModelIndex & modelIndex);
extern void* C_ZN8QVariantC2ERK21QPersistentModelIndex(void* arg0); // 3
  // proto:  void QVariant::QVariant(int typeId, const void * copy, uint flags);
extern void* C_ZN8QVariantC2EiPKvj(int32_t arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QVariant::QVariant(const QLineF & line);
extern void* C_ZN8QVariantC2ERK6QLineF(void* arg0); // 3
  // proto:  void QVariant::QVariant(int i);
extern void* C_ZN8QVariantC2Ei(int32_t arg0); // 3
  // proto:  void QVariant::QVariant(float f);
extern void* C_ZN8QVariantC2Ef(float arg0); // 3
  // proto:  void QVariant::QVariant(const QString & string);
extern void* C_ZN8QVariantC2ERK7QString(void* arg0); // 3
  // proto:  void QVariant::QVariant(double d);
extern void* C_ZN8QVariantC2Ed(double arg0); // 3
  // proto:  void QVariant::QVariant(bool b);
extern void* C_ZN8QVariantC2Eb(bool arg0); // 3
  // proto:  void QVariant::QVariant(const QEasingCurve & easing);
extern void* C_ZN8QVariantC2ERK12QEasingCurve(void* arg0); // 3
  // proto:  void QVariant::QVariant(QDataStream & s);
extern void* C_ZN8QVariantC2ER11QDataStream(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QSizeF & size);
extern void* C_ZN8QVariantC2ERK6QSizeF(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QStringList & stringlist);
extern void* C_ZN8QVariantC2ERK11QStringList(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QUuid & uuid);
extern void* C_ZN8QVariantC2ERK5QUuid(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QDate & date);
extern void* C_ZN8QVariantC2ERK5QDate(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QBitArray & bitarray);
extern void* C_ZN8QVariantC2ERK9QBitArray(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QModelIndex & modelIndex);
extern void* C_ZN8QVariantC2ERK11QModelIndex(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QUrl & url);
extern void* C_ZN8QVariantC2ERK4QUrl(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QVariant & other);
extern void* C_ZN8QVariantC2ERKS_(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QSize & size);
extern void* C_ZN8QVariantC2ERK5QSize(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QLine & line);
extern void* C_ZN8QVariantC2ERK5QLine(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QRegularExpression & re);
extern void* C_ZN8QVariantC2ERK18QRegularExpression(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QTime & time);
extern void* C_ZN8QVariantC2ERK5QTime(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QDateTime & datetime);
extern void* C_ZN8QVariantC2ERK9QDateTime(void* arg0); // 3
  // proto:  void QVariant::QVariant(const QRegExp & regExp);
extern void* C_ZN8QVariantC2ERK7QRegExp(void* arg0); // 3
  // proto:  QDateTime QVariant::toDateTime();
extern void* C_ZNK8QVariant10toDateTimeEv(void* qthis); // 4
  // proto:  QJsonArray QVariant::toJsonArray();
extern void C_ZNK8QVariant11toJsonArrayEv(void* qthis); // 4
  // proto:  bool QVariant::isValid();
extern bool C_ZNK8QVariant7isValidEv(void* qthis); // 2
  // proto:  QLineF QVariant::toLineF();
extern void* C_ZNK8QVariant7toLineFEv(void* qthis); // 4
  // proto:  QJsonValue QVariant::toJsonValue();
extern void C_ZNK8QVariant11toJsonValueEv(void* qthis); // 4
  // proto:  bool QVariant::canConvert(int targetTypeId);
extern bool C_ZNK8QVariant10canConvertEi(void* qthis, int32_t arg0); // 4
  // proto:  QPersistentModelIndex QVariant::toPersistentModelIndex();
extern void* C_ZNK8QVariant22toPersistentModelIndexEv(void* qthis); // 4
  // proto:  QPointF QVariant::toPointF();
extern void* C_ZNK8QVariant8toPointFEv(void* qthis); // 4
  // proto:  QJsonDocument QVariant::toJsonDocument();
extern void C_ZNK8QVariant14toJsonDocumentEv(void* qthis); // 4
  // proto:  QModelIndex QVariant::toModelIndex();
extern void* C_ZNK8QVariant12toModelIndexEv(void* qthis); // 4
  // proto:  void QVariant::detach();
extern void C_ZN8QVariant6detachEv(void* qthis); // 4
  // proto:  void * QVariant::data();
extern void C_ZN8QVariant4dataEv(void* qthis); // 4
  // proto:  qulonglong QVariant::toULongLong(bool * ok);
extern int64_t C_ZNK8QVariant11toULongLongEPb(void* qthis, void* arg0); // 4
  // proto:  QPoint QVariant::toPoint();
extern void* C_ZNK8QVariant7toPointEv(void* qthis); // 4
  // proto:  QHash<QString, QVariant> QVariant::toHash();
extern void C_ZNK8QVariant6toHashEv(void* qthis); // 4
  // proto:  bool QVariant::convert(int targetTypeId);
extern bool C_ZN8QVariant7convertEi(void* qthis, int32_t arg0); // 4
  // proto:  void QVariant::save(QDataStream & ds);
extern void C_ZNK8QVariant4saveER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  QRegExp QVariant::toRegExp();
extern void* C_ZNK8QVariant8toRegExpEv(void* qthis); // 4
  // proto:  QLine QVariant::toLine();
extern void* C_ZNK8QVariant6toLineEv(void* qthis); // 4
  // proto:  QStringList QVariant::toStringList();
extern void C_ZNK8QVariant12toStringListEv(void* qthis); // 4
  // proto:  QSizeF QVariant::toSizeF();
extern void* C_ZNK8QVariant7toSizeFEv(void* qthis); // 4
  // proto:  QRectF QVariant::toRectF();
extern void* C_ZNK8QVariant7toRectFEv(void* qthis); // 4
  // proto:  QUrl QVariant::toUrl();
extern void* C_ZNK8QVariant5toUrlEv(void* qthis); // 4
  // proto:  QUuid QVariant::toUuid();
extern void* C_ZNK8QVariant6toUuidEv(void* qthis); // 4
  // proto:  qreal QVariant::toReal(bool * ok);
extern double C_ZNK8QVariant6toRealEPb(void* qthis, void* arg0); // 4
  // proto:  void QVariant::~QVariant();
extern void C_ZN8QVariantD2Ev(void* qthis); // 4
  // proto:  bool QVariant::isNull();
extern bool C_ZNK8QVariant6isNullEv(void* qthis); // 4
  // proto:  QByteArray QVariant::toByteArray();
extern void* C_ZNK8QVariant11toByteArrayEv(void* qthis); // 4
  // proto:  double QVariant::toDouble(bool * ok);
extern double C_ZNK8QVariant8toDoubleEPb(void* qthis, void* arg0); // 4
  // proto:  QTime QVariant::toTime();
extern void* C_ZNK8QVariant6toTimeEv(void* qthis); // 4
  // proto:  QDate QVariant::toDate();
extern void* C_ZNK8QVariant6toDateEv(void* qthis); // 4
  // proto:  void QVariant::clear();
extern void C_ZN8QVariant5clearEv(void* qthis); // 4
  // proto:  QRect QVariant::toRect();
extern void* C_ZNK8QVariant6toRectEv(void* qthis); // 4
  // proto:  QSequentialIterable::const_iterator QSequentialIterable::begin();
extern void C_ZNK19QSequentialIterable5beginEv(void* qthis); // 4
  // proto:  QSequentialIterable::const_iterator QSequentialIterable::end();
extern void C_ZNK19QSequentialIterable3endEv(void* qthis); // 4
  // proto:  bool QSequentialIterable::canReverseIterate();
extern bool C_ZNK19QSequentialIterable17canReverseIterateEv(void* qthis); // 4
  // proto:  QVariant QSequentialIterable::at(int idx);
extern void* C_ZNK19QSequentialIterable2atEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSequentialIterable::size();
extern int32_t C_ZNK19QSequentialIterable4sizeEv(void* qthis); // 4
  // proto:  QAssociativeIterable::const_iterator QAssociativeIterable::begin();
extern void C_ZNK20QAssociativeIterable5beginEv(void* qthis); // 4
  // proto:  QVariant QAssociativeIterable::value(const QVariant & key);
extern void* C_ZNK20QAssociativeIterable5valueERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  QAssociativeIterable::const_iterator QAssociativeIterable::end();
extern void C_ZNK20QAssociativeIterable3endEv(void* qthis); // 4
  // proto:  QAssociativeIterable::const_iterator QAssociativeIterable::find(const QVariant & key);
extern void C_ZNK20QAssociativeIterable4findERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  int QAssociativeIterable::size();
extern int32_t C_ZNK20QAssociativeIterable4sizeEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QVariant)=16
type QVariant struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QSequentialIterable)=104
type QSequentialIterable struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAssociativeIterable)=112
type QAssociativeIterable struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QVariantComparisonHelper(const class QVariant &)
func NewQVariantComparisonHelper(args ...interface{}) *QVariantComparisonHelper {
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
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QVariantComparisonHelperC2ERK8QVariant(arg0)
    return &QVariantComparisonHelper{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVariantComparisonHelper", "QVariantComparisonHelper", args)
  }

  return nil // QVariantComparisonHelper{Qclsinst:qthis}
}

// load(class QDataStream &)
func (this *QVariant) Load(args ...interface{}) () {
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
    var arg0 = args[0].(*QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QVariant4loadER11QDataStream(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "load", args)
  }

  return
}

// toInt(_Bool *)
func (this *QVariant) Toint(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QVariant5toIntEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toInt", args)
  }

  return
}

// toUInt(_Bool *)
func (this *QVariant) Touint(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QVariant6toUIntEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toUInt", args)
  }

  return
}

// toFloat(_Bool *)
func (this *QVariant) Tofloat(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QVariant7toFloatEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toFloat", args)
  }

  return
}

// toChar()
func (this *QVariant) Tochar(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6toCharEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toChar", args)
  }

  return
}

// typeName()
func (this *QVariant) Typename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant8typeNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "typeName", args)
  }

  return
}

// toRegularExpression()
func (this *QVariant) Toregularexpression(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant19toRegularExpressionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegularExpression{}) // "QRegularExpression"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toRegularExpression", args)
  }

  return
}

// nameToType(const char *)
func (this *QVariant) Nametotype_S(args ...interface{}) () {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN8QVariant10nameToTypeEPKc(arg0)
  default:
    qtrt.ErrorResolve("QVariant", "nameToType", args)
  }

  return
}

// toEasingCurve()
func (this *QVariant) Toeasingcurve(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant13toEasingCurveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QEasingCurve{}) // "QEasingCurve"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toEasingCurve", args)
  }

  return
}

// toList()
func (this *QVariant) Tolist(args ...interface{}) () {
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
    C.C_ZNK8QVariant6toListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toList", args)
  }

  return
}

// toMap()
func (this *QVariant) Tomap(args ...interface{}) () {
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
    C.C_ZNK8QVariant5toMapEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toMap", args)
  }

  return
}

// typeToName(int)
func (this *QVariant) Typetoname_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QVariant10typeToNameEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "typeToName", args)
  }

  return
}

// toBool()
func (this *QVariant) Tobool(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6toBoolEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toBool", args)
  }

  return
}

// toBitArray()
func (this *QVariant) Tobitarray(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant10toBitArrayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBitArray{}) // "QBitArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toBitArray", args)
  }

  return
}

// userType()
func (this *QVariant) Usertype(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant8userTypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "userType", args)
  }

  return
}

// toLongLong(_Bool *)
func (this *QVariant) Tolonglong(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QVariant10toLongLongEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qlonglong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toLongLong", args)
  }

  return
}

// constData()
func (this *QVariant) Constdata(args ...interface{}) () {
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
    C.C_ZNK8QVariant9constDataEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "constData", args)
  }

  return
}

// toString()
func (this *QVariant) Tostring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant8toStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toString", args)
  }

  return
}

// swap(class QVariant &)
func (this *QVariant) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QVariant4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "swap", args)
  }

  return
}

// toJsonObject()
func (this *QVariant) Tojsonobject(args ...interface{}) () {
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
    C.C_ZNK8QVariant12toJsonObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toJsonObject", args)
  }

  return
}

// type()
func (this *QVariant) Type_(args ...interface{}) () {
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
    C.C_ZNK8QVariant4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "type", args)
  }

  return
}

// toSize()
func (this *QVariant) Tosize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6toSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toSize", args)
  }

  return
}

// isDetached()
func (this *QVariant) Isdetached(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "isDetached", args)
  }

  return
}

// toLocale()
func (this *QVariant) Tolocale(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant8toLocaleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLocale{}) // "QLocale"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toLocale", args)
  }

  return
}

// QVariant(const class QRectF &)
func NewQVariant(args ...interface{}) *QVariant {
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
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK6QRectF(arg0)
    return &QVariant{Qclsinst:qthis}
  case 1:
    // invoke: _ZN8QVariantC1E5QChar
    // invoke: void QVariant(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2E5QChar(arg0)
    return &QVariant{Qclsinst:qthis}
  case 2:
    // invoke: _ZN8QVariantC1EPKc
    // invoke: void QVariant(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2EPKc(arg0)
    return &QVariant{Qclsinst:qthis}
  case 3:
    // invoke: _ZN8QVariantC1ERK7QPointF
    // invoke: void QVariant(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK7QPointF(arg0)
    return &QVariant{Qclsinst:qthis}
  case 4:
    // invoke: _ZN8QVariantC1ERK5QRect
    // invoke: void QVariant(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK5QRect(arg0)
    return &QVariant{Qclsinst:qthis}
  case 5:
    // invoke: _ZN8QVariantC1EiPKv
    // invoke: void QVariant(int, const void *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2EiPKv(arg0, arg1)
    return &QVariant{Qclsinst:qthis}
  case 6:
    // invoke: _ZN8QVariantC1ERK6QPoint
    // invoke: void QVariant(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK6QPoint(arg0)
    return &QVariant{Qclsinst:qthis}
  case 7:
    // invoke: _ZN8QVariantC1Ej
    // invoke: void QVariant(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2Ej(arg0)
    return &QVariant{Qclsinst:qthis}
  case 8:
    // invoke: _ZN8QVariantC1ERK7QLocale
    // invoke: void QVariant(const class QLocale &)
    var arg0 = args[0].(*QLocale).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK7QLocale(arg0)
    return &QVariant{Qclsinst:qthis}
  case 9:
    // invoke: _ZN8QVariantC1ERK10QByteArray
    // invoke: void QVariant(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK10QByteArray(arg0)
    return &QVariant{Qclsinst:qthis}
  case 10:
    // invoke: _ZN8QVariantC1Ex
    // invoke: void QVariant(qlonglong)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2Ex(arg0)
    return &QVariant{Qclsinst:qthis}
  case 11:
    // invoke: _ZN8QVariantC1Ey
    // invoke: void QVariant(qulonglong)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2Ey(arg0)
    return &QVariant{Qclsinst:qthis}
  case 12:
    // invoke: _ZN8QVariantC1Ev
    // invoke: void QVariant()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2Ev()
    return &QVariant{Qclsinst:qthis}
  case 13:
    // invoke: _ZN8QVariantC1ERK21QPersistentModelIndex
    // invoke: void QVariant(const class QPersistentModelIndex &)
    var arg0 = args[0].(*QPersistentModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK21QPersistentModelIndex(arg0)
    return &QVariant{Qclsinst:qthis}
  case 14:
    // invoke: _ZN8QVariantC1EiPKvj
    // invoke: void QVariant(int, const void *, uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2EiPKvj(arg0, arg1, arg2)
    return &QVariant{Qclsinst:qthis}
  case 15:
    // invoke: _ZN8QVariantC1ERK6QLineF
    // invoke: void QVariant(const class QLineF &)
    var arg0 = args[0].(*QLineF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK6QLineF(arg0)
    return &QVariant{Qclsinst:qthis}
  case 16:
    // invoke: _ZN8QVariantC1Ei
    // invoke: void QVariant(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2Ei(arg0)
    return &QVariant{Qclsinst:qthis}
  case 17:
    // invoke: _ZN8QVariantC1Ef
    // invoke: void QVariant(float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2Ef(arg0)
    return &QVariant{Qclsinst:qthis}
  case 18:
    // invoke: _ZN8QVariantC1ERK7QString
    // invoke: void QVariant(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK7QString(arg0)
    return &QVariant{Qclsinst:qthis}
  case 19:
    // invoke: _ZN8QVariantC1Ed
    // invoke: void QVariant(double)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2Ed(arg0)
    return &QVariant{Qclsinst:qthis}
  case 20:
    // invoke: _ZN8QVariantC1Eb
    // invoke: void QVariant(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2Eb(arg0)
    return &QVariant{Qclsinst:qthis}
  case 21:
    // invoke: _ZN8QVariantC1ERK12QEasingCurve
    // invoke: void QVariant(const class QEasingCurve &)
    var arg0 = args[0].(*QEasingCurve).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK12QEasingCurve(arg0)
    return &QVariant{Qclsinst:qthis}
  case 22:
    // invoke: _ZN8QVariantC1ER11QDataStream
    // invoke: void QVariant(class QDataStream &)
    var arg0 = args[0].(*QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ER11QDataStream(arg0)
    return &QVariant{Qclsinst:qthis}
  case 23:
    // invoke: _ZN8QVariantC1ERK6QSizeF
    // invoke: void QVariant(const class QSizeF &)
    var arg0 = args[0].(*QSizeF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK6QSizeF(arg0)
    return &QVariant{Qclsinst:qthis}
  case 24:
    // invoke: _ZN8QVariantC1ERK11QStringList
    // invoke: void QVariant(const class QStringList &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK11QStringList(arg0)
    return &QVariant{Qclsinst:qthis}
  case 25:
    // invoke: _ZN8QVariantC1ERK5QUuid
    // invoke: void QVariant(const class QUuid &)
    var arg0 = args[0].(*QUuid).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK5QUuid(arg0)
    return &QVariant{Qclsinst:qthis}
  case 26:
    // invoke: _ZN8QVariantC1ERK5QDate
    // invoke: void QVariant(const class QDate &)
    var arg0 = args[0].(*QDate).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK5QDate(arg0)
    return &QVariant{Qclsinst:qthis}
  case 27:
    // invoke: _ZN8QVariantC1ERK9QBitArray
    // invoke: void QVariant(const class QBitArray &)
    var arg0 = args[0].(*QBitArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK9QBitArray(arg0)
    return &QVariant{Qclsinst:qthis}
  case 28:
    // invoke: _ZN8QVariantC1ERK11QModelIndex
    // invoke: void QVariant(const class QModelIndex &)
    var arg0 = args[0].(*QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK11QModelIndex(arg0)
    return &QVariant{Qclsinst:qthis}
  case 29:
    // invoke: _ZN8QVariantC1ERK4QUrl
    // invoke: void QVariant(const class QUrl &)
    var arg0 = args[0].(*QUrl).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK4QUrl(arg0)
    return &QVariant{Qclsinst:qthis}
  case 30:
    // invoke: _ZN8QVariantC1ERKS_
    // invoke: void QVariant(const class QVariant &)
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERKS_(arg0)
    return &QVariant{Qclsinst:qthis}
  case 31:
    // invoke: _ZN8QVariantC1ERK5QSize
    // invoke: void QVariant(const class QSize &)
    var arg0 = args[0].(*QSize).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK5QSize(arg0)
    return &QVariant{Qclsinst:qthis}
  case 32:
    // invoke: _ZN8QVariantC1ERK5QLine
    // invoke: void QVariant(const class QLine &)
    var arg0 = args[0].(*QLine).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK5QLine(arg0)
    return &QVariant{Qclsinst:qthis}
  case 33:
    // invoke: _ZN8QVariantC1ERK18QRegularExpression
    // invoke: void QVariant(const class QRegularExpression &)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK18QRegularExpression(arg0)
    return &QVariant{Qclsinst:qthis}
  case 34:
    // invoke: _ZN8QVariantC1ERK5QTime
    // invoke: void QVariant(const class QTime &)
    var arg0 = args[0].(*QTime).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK5QTime(arg0)
    return &QVariant{Qclsinst:qthis}
  case 35:
    // invoke: _ZN8QVariantC1ERK9QDateTime
    // invoke: void QVariant(const class QDateTime &)
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK9QDateTime(arg0)
    return &QVariant{Qclsinst:qthis}
  case 36:
    // invoke: _ZN8QVariantC1ERK7QRegExp
    // invoke: void QVariant(const class QRegExp &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QVariantC2ERK7QRegExp(arg0)
    return &QVariant{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVariant", "QVariant", args)
  }

  return nil // QVariant{Qclsinst:qthis}
}

// toDateTime()
func (this *QVariant) Todatetime(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant10toDateTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toDateTime", args)
  }

  return
}

// toJsonArray()
func (this *QVariant) Tojsonarray(args ...interface{}) () {
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
    C.C_ZNK8QVariant11toJsonArrayEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toJsonArray", args)
  }

  return
}

// isValid()
func (this *QVariant) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "isValid", args)
  }

  return
}

// toLineF()
func (this *QVariant) Tolinef(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant7toLineFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLineF{}) // "QLineF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toLineF", args)
  }

  return
}

// toJsonValue()
func (this *QVariant) Tojsonvalue(args ...interface{}) () {
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
    C.C_ZNK8QVariant11toJsonValueEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toJsonValue", args)
  }

  return
}

// canConvert(int)
func (this *QVariant) Canconvert(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QVariant10canConvertEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "canConvert", args)
  }

  return
}

// toPersistentModelIndex()
func (this *QVariant) Topersistentmodelindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant22toPersistentModelIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPersistentModelIndex{}) // "QPersistentModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toPersistentModelIndex", args)
  }

  return
}

// toPointF()
func (this *QVariant) Topointf(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant8toPointFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toPointF", args)
  }

  return
}

// toJsonDocument()
func (this *QVariant) Tojsondocument(args ...interface{}) () {
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
    C.C_ZNK8QVariant14toJsonDocumentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toJsonDocument", args)
  }

  return
}

// toModelIndex()
func (this *QVariant) Tomodelindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant12toModelIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toModelIndex", args)
  }

  return
}

// detach()
func (this *QVariant) Detach(args ...interface{}) () {
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
    C.C_ZN8QVariant6detachEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "detach", args)
  }

  return
}

// data()
func (this *QVariant) Data(args ...interface{}) () {
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
    C.C_ZN8QVariant4dataEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "data", args)
  }

  return
}

// toULongLong(_Bool *)
func (this *QVariant) Toulonglong(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QVariant11toULongLongEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qulonglong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toULongLong", args)
  }

  return
}

// toPoint()
func (this *QVariant) Topoint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant7toPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toPoint", args)
  }

  return
}

// toHash()
func (this *QVariant) Tohash(args ...interface{}) () {
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
    C.C_ZNK8QVariant6toHashEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toHash", args)
  }

  return
}

// convert(int)
func (this *QVariant) Convert(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QVariant7convertEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "convert", args)
  }

  return
}

// save(class QDataStream &)
func (this *QVariant) Save(args ...interface{}) () {
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
    var arg0 = args[0].(*QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK8QVariant4saveER11QDataStream(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QVariant", "save", args)
  }

  return
}

// toRegExp()
func (this *QVariant) Toregexp(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant8toRegExpEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegExp{}) // "QRegExp"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toRegExp", args)
  }

  return
}

// toLine()
func (this *QVariant) Toline(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6toLineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLine{}) // "QLine"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toLine", args)
  }

  return
}

// toStringList()
func (this *QVariant) Tostringlist(args ...interface{}) () {
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
    C.C_ZNK8QVariant12toStringListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "toStringList", args)
  }

  return
}

// toSizeF()
func (this *QVariant) Tosizef(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant7toSizeFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toSizeF", args)
  }

  return
}

// toRectF()
func (this *QVariant) Torectf(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant7toRectFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toRectF", args)
  }

  return
}

// toUrl()
func (this *QVariant) Tourl(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant5toUrlEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUrl{}) // "QUrl"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toUrl", args)
  }

  return
}

// toUuid()
func (this *QVariant) Touuid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6toUuidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUuid{}) // "QUuid"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toUuid", args)
  }

  return
}

// toReal(_Bool *)
func (this *QVariant) Toreal(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QVariant6toRealEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toReal", args)
  }

  return
}

// ~QVariant()
func (this *QVariant) Freeqvariant(args ...interface{}) () {
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
    C.C_ZN8QVariantD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "~QVariant", args)
  }

  return
}

// isNull()
func (this *QVariant) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "isNull", args)
  }

  return
}

// toByteArray()
func (this *QVariant) Tobytearray(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant11toByteArrayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toByteArray", args)
  }

  return
}

// toDouble(_Bool *)
func (this *QVariant) Todouble(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QVariant8toDoubleEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toDouble", args)
  }

  return
}

// toTime()
func (this *QVariant) Totime(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6toTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toTime", args)
  }

  return
}

// toDate()
func (this *QVariant) Todate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6toDateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toDate", args)
  }

  return
}

// clear()
func (this *QVariant) Clear(args ...interface{}) () {
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
    C.C_ZN8QVariant5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVariant", "clear", args)
  }

  return
}

// toRect()
func (this *QVariant) Torect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QVariant6toRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVariant", "toRect", args)
  }

  return
}

// begin()
func (this *QSequentialIterable) Begin(args ...interface{}) () {
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
    C.C_ZNK19QSequentialIterable5beginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialIterable", "begin", args)
  }

  return
}

// end()
func (this *QSequentialIterable) End(args ...interface{}) () {
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
    C.C_ZNK19QSequentialIterable3endEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSequentialIterable", "end", args)
  }

  return
}

// canReverseIterate()
func (this *QSequentialIterable) Canreverseiterate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QSequentialIterable17canReverseIterateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSequentialIterable", "canReverseIterate", args)
  }

  return
}

// at(int)
func (this *QSequentialIterable) At(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QSequentialIterable2atEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSequentialIterable", "at", args)
  }

  return
}

// size()
func (this *QSequentialIterable) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QSequentialIterable4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSequentialIterable", "size", args)
  }

  return
}

// begin()
func (this *QAssociativeIterable) Begin(args ...interface{}) () {
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
    C.C_ZNK20QAssociativeIterable5beginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "begin", args)
  }

  return
}

// value(const class QVariant &)
func (this *QAssociativeIterable) Value(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QAssociativeIterable5valueERK8QVariant(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "value", args)
  }

  return
}

// end()
func (this *QAssociativeIterable) End(args ...interface{}) () {
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
    C.C_ZNK20QAssociativeIterable3endEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "end", args)
  }

  return
}

// find(const class QVariant &)
func (this *QAssociativeIterable) Find(args ...interface{}) () {
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
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK20QAssociativeIterable4findERK8QVariant(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "find", args)
  }

  return
}

// size()
func (this *QAssociativeIterable) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK20QAssociativeIterable4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "size", args)
  }

  return
}

// <= body block end

