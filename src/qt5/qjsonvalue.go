package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qjsonvalue.h
// dst-file: /src/core/qjsonvalue.go
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
  // proto:  void QJsonValue::QJsonValue(double n);
extern void* C_ZN10QJsonValueC2Ed(double arg0); // 3
  // proto:  void QJsonValue::QJsonValue(bool b);
extern void* C_ZN10QJsonValueC2Eb(bool arg0); // 3
  // proto:  void QJsonValue::QJsonValue(int n);
extern void* C_ZN10QJsonValueC2Ei(int32_t arg0); // 3
  // proto:  void QJsonValue::QJsonValue(const QString & s);
extern void* C_ZN10QJsonValueC2ERK7QString(void* arg0); // 3
  // proto:  void QJsonValue::QJsonValue(const char * s);
extern void* C_ZN10QJsonValueC2EPKc(unsigned char* arg0); // 1
  // proto:  void QJsonValue::QJsonValue(qint64 n);
extern void* C_ZN10QJsonValueC2Ex(int64_t arg0); // 3
  // proto: static QJsonValue QJsonValue::fromVariant(const QVariant & variant);
extern void C_ZN10QJsonValue11fromVariantERK8QVariant(void* arg0); // 4
  // proto:  void QJsonValue::~QJsonValue();
extern void C_ZN10QJsonValueD2Ev(void* qthis); // 4
  // proto:  QJsonObject QJsonValue::toObject();
extern void C_ZNK10QJsonValue8toObjectEv(void* qthis); // 4
  // proto:  bool QJsonValue::toBool(bool defaultValue);
extern void C_ZNK10QJsonValue6toBoolEb(void* qthis, bool arg0); // 4
  // proto:  QVariant QJsonValue::toVariant();
extern void C_ZNK10QJsonValue9toVariantEv(void* qthis); // 4
  // proto:  QString QJsonValue::toString(const QString & defaultValue);
extern void C_ZNK10QJsonValue8toStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  QJsonValue::Type QJsonValue::type();
extern void C_ZNK10QJsonValue4typeEv(void* qthis); // 4
  // proto:  bool QJsonValue::isUndefined();
extern void C_ZNK10QJsonValue11isUndefinedEv(void* qthis); // 2
  // proto:  bool QJsonValue::isArray();
extern void C_ZNK10QJsonValue7isArrayEv(void* qthis); // 2
  // proto:  bool QJsonValue::isBool();
extern void C_ZNK10QJsonValue6isBoolEv(void* qthis); // 2
  // proto:  bool QJsonValue::isObject();
extern void C_ZNK10QJsonValue8isObjectEv(void* qthis); // 2
  // proto:  bool QJsonValue::isDouble();
extern void C_ZNK10QJsonValue8isDoubleEv(void* qthis); // 2
  // proto:  QJsonArray QJsonValue::toArray();
extern void C_ZNK10QJsonValue7toArrayEv(void* qthis); // 4
  // proto:  bool QJsonValue::isString();
extern void C_ZNK10QJsonValue8isStringEv(void* qthis); // 2
  // proto:  int QJsonValue::toInt(int defaultValue);
extern void C_ZNK10QJsonValue5toIntEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QJsonValue::isNull();
extern void C_ZNK10QJsonValue6isNullEv(void* qthis); // 2
  // proto:  double QJsonValue::toDouble(double defaultValue);
extern void C_ZNK10QJsonValue8toDoubleEd(void* qthis, double arg0); // 4
  // proto:  bool QJsonValueRef::isArray();
extern void C_ZNK13QJsonValueRef7isArrayEv(void* qthis); // 2
  // proto:  QJsonArray QJsonValueRef::toArray();
extern void C_ZNK13QJsonValueRef7toArrayEv(void* qthis); // 4
  // proto:  bool QJsonValueRef::isString();
extern void C_ZNK13QJsonValueRef8isStringEv(void* qthis); // 2
  // proto:  QJsonObject QJsonValueRef::toObject();
extern void C_ZNK13QJsonValueRef8toObjectEv(void* qthis); // 4
  // proto:  int QJsonValueRef::toInt(int defaultValue);
extern void C_ZNK13QJsonValueRef5toIntEi(void* qthis, int32_t arg0); // 2
  // proto:  int QJsonValueRef::toInt();
extern void C_ZNK13QJsonValueRef5toIntEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::isBool();
extern void C_ZNK13QJsonValueRef6isBoolEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::isNull();
extern void C_ZNK13QJsonValueRef6isNullEv(void* qthis); // 2
  // proto:  double QJsonValueRef::toDouble(double defaultValue);
extern void C_ZNK13QJsonValueRef8toDoubleEd(void* qthis, double arg0); // 2
  // proto:  double QJsonValueRef::toDouble();
extern void C_ZNK13QJsonValueRef8toDoubleEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::toBool();
extern void C_ZNK13QJsonValueRef6toBoolEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::toBool(bool defaultValue);
extern void C_ZNK13QJsonValueRef6toBoolEb(void* qthis, bool arg0); // 2
  // proto:  QVariant QJsonValueRef::toVariant();
extern void C_ZNK13QJsonValueRef9toVariantEv(void* qthis); // 4
  // proto:  QString QJsonValueRef::toString();
extern void C_ZNK13QJsonValueRef8toStringEv(void* qthis); // 2
  // proto:  QString QJsonValueRef::toString(const QString & defaultValue);
extern void C_ZNK13QJsonValueRef8toStringERK7QString(void* qthis, void* arg0); // 2
  // proto:  bool QJsonValueRef::isUndefined();
extern void C_ZNK13QJsonValueRef11isUndefinedEv(void* qthis); // 2
  // proto:  QJsonValue::Type QJsonValueRef::type();
extern void C_ZNK13QJsonValueRef4typeEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::isObject();
extern void C_ZNK13QJsonValueRef8isObjectEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::isDouble();
extern void C_ZNK13QJsonValueRef8isDoubleEv(void* qthis); // 2
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

// class sizeof(QJsonValueRefPtr)=16
type QJsonValueRefPtr struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QJsonValuePtr)=24
type QJsonValuePtr struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QJsonValue)=24
type QJsonValue struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QJsonValueRef)=16
type QJsonValueRef struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QJsonValue(double)
func NewQJsonValue(args ...interface{}) *QJsonValue {
  // QJsonValue(double)
  // QJsonValue(_Bool)
  // QJsonValue(int)
  // QJsonValue(const class QString &)
  // QJsonValue(const char *)
  // QJsonValue(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.BoolTy(false) // "bool"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(true) // "const char *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonValueC1Ed
    // invoke: void QJsonValue(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2Ed(arg0)
    return &QJsonValue{qclsinst:qthis}
  case 1:
    // invoke: _ZN10QJsonValueC1Eb
    // invoke: void QJsonValue(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2Eb(arg0)
    return &QJsonValue{qclsinst:qthis}
  case 2:
    // invoke: _ZN10QJsonValueC1Ei
    // invoke: void QJsonValue(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2Ei(arg0)
    return &QJsonValue{qclsinst:qthis}
  case 3:
    // invoke: _ZN10QJsonValueC1ERK7QString
    // invoke: void QJsonValue(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2ERK7QString(arg0)
    return &QJsonValue{qclsinst:qthis}
  case 4:
    // invoke: _ZN10QJsonValueC1EPKc
    // invoke: void QJsonValue(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2EPKc(arg0)
    return &QJsonValue{qclsinst:qthis}
  case 5:
    // invoke: _ZN10QJsonValueC1Ex
    // invoke: void QJsonValue(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2Ex(arg0)
    return &QJsonValue{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QJsonValue", "QJsonValue", args)
  }

  return nil // QJsonValue{qclsinst:qthis}
}

// fromVariant(const class QVariant &)
func (this *QJsonValue) fromVariant_s(args ...interface{}) () {
  // fromVariant(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonValue11fromVariantERK8QVariant
    // invoke: QJsonValue fromVariant(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QJsonValue11fromVariantERK8QVariant(arg0)
  default:
    qtrt.ErrorResolve("QJsonValue", "fromVariant", args)
  }

}

// ~QJsonValue()
func (this *QJsonValue) FreeQJsonValue(args ...interface{}) () {
  // ~QJsonValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QJsonValueD0Ev
    // invoke: void ~QJsonValue()
    C.C_ZN10QJsonValueD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "~QJsonValue", args)
  }

}

// toObject()
func (this *QJsonValue) toObject(args ...interface{}) () {
  // toObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue8toObjectEv
    // invoke: QJsonObject toObject()
    C.C_ZNK10QJsonValue8toObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "toObject", args)
  }

}

// toBool(_Bool)
func (this *QJsonValue) toBool(args ...interface{}) () {
  // toBool(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue6toBoolEb
    // invoke: bool toBool(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QJsonValue6toBoolEb(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "toBool", args)
  }

}

// toVariant()
func (this *QJsonValue) toVariant(args ...interface{}) () {
  // toVariant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue9toVariantEv
    // invoke: QVariant toVariant()
    var ret = C.C_ZNK10QJsonValue9toVariantEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "toVariant", args)
  }

}

// toString(const class QString &)
func (this *QJsonValue) toString(args ...interface{}) () {
  // toString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QJsonValue8toStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "toString", args)
  }

}

// type()
func (this *QJsonValue) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue4typeEv
    // invoke: QJsonValue::Type type()
    C.C_ZNK10QJsonValue4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "type", args)
  }

}

// isUndefined()
func (this *QJsonValue) isUndefined(args ...interface{}) () {
  // isUndefined()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue11isUndefinedEv
    // invoke: bool isUndefined()
    var ret = C.C_ZNK10QJsonValue11isUndefinedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "isUndefined", args)
  }

}

// isArray()
func (this *QJsonValue) isArray(args ...interface{}) () {
  // isArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue7isArrayEv
    // invoke: bool isArray()
    var ret = C.C_ZNK10QJsonValue7isArrayEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "isArray", args)
  }

}

// isBool()
func (this *QJsonValue) isBool(args ...interface{}) () {
  // isBool()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue6isBoolEv
    // invoke: bool isBool()
    var ret = C.C_ZNK10QJsonValue6isBoolEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "isBool", args)
  }

}

// isObject()
func (this *QJsonValue) isObject(args ...interface{}) () {
  // isObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue8isObjectEv
    // invoke: bool isObject()
    var ret = C.C_ZNK10QJsonValue8isObjectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "isObject", args)
  }

}

// isDouble()
func (this *QJsonValue) isDouble(args ...interface{}) () {
  // isDouble()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue8isDoubleEv
    // invoke: bool isDouble()
    var ret = C.C_ZNK10QJsonValue8isDoubleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "isDouble", args)
  }

}

// toArray()
func (this *QJsonValue) toArray(args ...interface{}) () {
  // toArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue7toArrayEv
    // invoke: QJsonArray toArray()
    C.C_ZNK10QJsonValue7toArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "toArray", args)
  }

}

// isString()
func (this *QJsonValue) isString(args ...interface{}) () {
  // isString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue8isStringEv
    // invoke: bool isString()
    var ret = C.C_ZNK10QJsonValue8isStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "isString", args)
  }

}

// toInt(int)
func (this *QJsonValue) toInt(args ...interface{}) () {
  // toInt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue5toIntEi
    // invoke: int toInt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QJsonValue5toIntEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "toInt", args)
  }

}

// isNull()
func (this *QJsonValue) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue6isNullEv
    // invoke: bool isNull()
    var ret = C.C_ZNK10QJsonValue6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "isNull", args)
  }

}

// toDouble(double)
func (this *QJsonValue) toDouble(args ...interface{}) () {
  // toDouble(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue8toDoubleEd
    // invoke: double toDouble(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QJsonValue8toDoubleEd(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValue", "toDouble", args)
  }

}

// isArray()
func (this *QJsonValueRef) isArray(args ...interface{}) () {
  // isArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef7isArrayEv
    // invoke: bool isArray()
    var ret = C.C_ZNK13QJsonValueRef7isArrayEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isArray", args)
  }

}

// toArray()
func (this *QJsonValueRef) toArray(args ...interface{}) () {
  // toArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef7toArrayEv
    // invoke: QJsonArray toArray()
    C.C_ZNK13QJsonValueRef7toArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toArray", args)
  }

}

// isString()
func (this *QJsonValueRef) isString(args ...interface{}) () {
  // isString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef8isStringEv
    // invoke: bool isString()
    var ret = C.C_ZNK13QJsonValueRef8isStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isString", args)
  }

}

// toObject()
func (this *QJsonValueRef) toObject(args ...interface{}) () {
  // toObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef8toObjectEv
    // invoke: QJsonObject toObject()
    C.C_ZNK13QJsonValueRef8toObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toObject", args)
  }

}

// toInt(int)
func (this *QJsonValueRef) toInt(args ...interface{}) () {
  // toInt(int)
  // toInt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef5toIntEi
    // invoke: int toInt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QJsonValueRef5toIntEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QJsonValueRef5toIntEv
    // invoke: int toInt()
    var ret = C.C_ZNK13QJsonValueRef5toIntEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toInt", args)
  }

}

// isBool()
func (this *QJsonValueRef) isBool(args ...interface{}) () {
  // isBool()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef6isBoolEv
    // invoke: bool isBool()
    var ret = C.C_ZNK13QJsonValueRef6isBoolEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isBool", args)
  }

}

// isNull()
func (this *QJsonValueRef) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef6isNullEv
    // invoke: bool isNull()
    var ret = C.C_ZNK13QJsonValueRef6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isNull", args)
  }

}

// toDouble(double)
func (this *QJsonValueRef) toDouble(args ...interface{}) () {
  // toDouble(double)
  // toDouble()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "double"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef8toDoubleEd
    // invoke: double toDouble(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QJsonValueRef8toDoubleEd(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QJsonValueRef8toDoubleEv
    // invoke: double toDouble()
    var ret = C.C_ZNK13QJsonValueRef8toDoubleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toDouble", args)
  }

}

// toBool()
func (this *QJsonValueRef) toBool(args ...interface{}) () {
  // toBool()
  // toBool(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef6toBoolEv
    // invoke: bool toBool()
    var ret = C.C_ZNK13QJsonValueRef6toBoolEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QJsonValueRef6toBoolEb
    // invoke: bool toBool(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QJsonValueRef6toBoolEb(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toBool", args)
  }

}

// toVariant()
func (this *QJsonValueRef) toVariant(args ...interface{}) () {
  // toVariant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef9toVariantEv
    // invoke: QVariant toVariant()
    var ret = C.C_ZNK13QJsonValueRef9toVariantEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toVariant", args)
  }

}

// toString()
func (this *QJsonValueRef) toString(args ...interface{}) () {
  // toString()
  // toString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef8toStringEv
    // invoke: QString toString()
    var ret = C.C_ZNK13QJsonValueRef8toStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK13QJsonValueRef8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK13QJsonValueRef8toStringERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toString", args)
  }

}

// isUndefined()
func (this *QJsonValueRef) isUndefined(args ...interface{}) () {
  // isUndefined()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef11isUndefinedEv
    // invoke: bool isUndefined()
    var ret = C.C_ZNK13QJsonValueRef11isUndefinedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isUndefined", args)
  }

}

// type()
func (this *QJsonValueRef) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef4typeEv
    // invoke: QJsonValue::Type type()
    C.C_ZNK13QJsonValueRef4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "type", args)
  }

}

// isObject()
func (this *QJsonValueRef) isObject(args ...interface{}) () {
  // isObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef8isObjectEv
    // invoke: bool isObject()
    var ret = C.C_ZNK13QJsonValueRef8isObjectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isObject", args)
  }

}

// isDouble()
func (this *QJsonValueRef) isDouble(args ...interface{}) () {
  // isDouble()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef8isDoubleEv
    // invoke: bool isDouble()
    var ret = C.C_ZNK13QJsonValueRef8isDoubleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isDouble", args)
  }

}

// <= body block end

