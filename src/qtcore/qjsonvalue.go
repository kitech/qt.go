package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern void* C_ZN10QJsonValueC2EPKc(void* arg0); // 1
  // proto:  void QJsonValue::QJsonValue(qint64 n);
extern void* C_ZN10QJsonValueC2Ex(int64_t arg0); // 3
  // proto: static QJsonValue QJsonValue::fromVariant(const QVariant & variant);
extern void C_ZN10QJsonValue11fromVariantERK8QVariant(void* arg0); // 4
  // proto:  void QJsonValue::~QJsonValue();
extern void C_ZN10QJsonValueD2Ev(void* qthis); // 4
  // proto:  QJsonObject QJsonValue::toObject();
extern void C_ZNK10QJsonValue8toObjectEv(void* qthis); // 4
  // proto:  bool QJsonValue::toBool(bool defaultValue);
extern bool C_ZNK10QJsonValue6toBoolEb(void* qthis, bool arg0); // 4
  // proto:  QVariant QJsonValue::toVariant();
extern void* C_ZNK10QJsonValue9toVariantEv(void* qthis); // 4
  // proto:  QString QJsonValue::toString(const QString & defaultValue);
extern void* C_ZNK10QJsonValue8toStringERK7QString(void* qthis, void* arg0); // 4
  // proto:  QJsonValue::Type QJsonValue::type();
extern void C_ZNK10QJsonValue4typeEv(void* qthis); // 4
  // proto:  bool QJsonValue::isUndefined();
extern bool C_ZNK10QJsonValue11isUndefinedEv(void* qthis); // 2
  // proto:  bool QJsonValue::isArray();
extern bool C_ZNK10QJsonValue7isArrayEv(void* qthis); // 2
  // proto:  bool QJsonValue::isBool();
extern bool C_ZNK10QJsonValue6isBoolEv(void* qthis); // 2
  // proto:  bool QJsonValue::isObject();
extern bool C_ZNK10QJsonValue8isObjectEv(void* qthis); // 2
  // proto:  bool QJsonValue::isDouble();
extern bool C_ZNK10QJsonValue8isDoubleEv(void* qthis); // 2
  // proto:  QJsonArray QJsonValue::toArray();
extern void C_ZNK10QJsonValue7toArrayEv(void* qthis); // 4
  // proto:  bool QJsonValue::isString();
extern bool C_ZNK10QJsonValue8isStringEv(void* qthis); // 2
  // proto:  int QJsonValue::toInt(int defaultValue);
extern int32_t C_ZNK10QJsonValue5toIntEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QJsonValue::isNull();
extern bool C_ZNK10QJsonValue6isNullEv(void* qthis); // 2
  // proto:  double QJsonValue::toDouble(double defaultValue);
extern double C_ZNK10QJsonValue8toDoubleEd(void* qthis, double arg0); // 4
  // proto:  bool QJsonValueRef::isArray();
extern bool C_ZNK13QJsonValueRef7isArrayEv(void* qthis); // 2
  // proto:  QJsonArray QJsonValueRef::toArray();
extern void C_ZNK13QJsonValueRef7toArrayEv(void* qthis); // 4
  // proto:  bool QJsonValueRef::isString();
extern bool C_ZNK13QJsonValueRef8isStringEv(void* qthis); // 2
  // proto:  QJsonObject QJsonValueRef::toObject();
extern void C_ZNK13QJsonValueRef8toObjectEv(void* qthis); // 4
  // proto:  int QJsonValueRef::toInt(int defaultValue);
extern int32_t C_ZNK13QJsonValueRef5toIntEi(void* qthis, int32_t arg0); // 2
  // proto:  int QJsonValueRef::toInt();
extern int32_t C_ZNK13QJsonValueRef5toIntEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::isBool();
extern bool C_ZNK13QJsonValueRef6isBoolEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::isNull();
extern bool C_ZNK13QJsonValueRef6isNullEv(void* qthis); // 2
  // proto:  double QJsonValueRef::toDouble(double defaultValue);
extern double C_ZNK13QJsonValueRef8toDoubleEd(void* qthis, double arg0); // 2
  // proto:  double QJsonValueRef::toDouble();
extern double C_ZNK13QJsonValueRef8toDoubleEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::toBool();
extern bool C_ZNK13QJsonValueRef6toBoolEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::toBool(bool defaultValue);
extern bool C_ZNK13QJsonValueRef6toBoolEb(void* qthis, bool arg0); // 2
  // proto:  QVariant QJsonValueRef::toVariant();
extern void* C_ZNK13QJsonValueRef9toVariantEv(void* qthis); // 4
  // proto:  QString QJsonValueRef::toString();
extern void* C_ZNK13QJsonValueRef8toStringEv(void* qthis); // 2
  // proto:  QString QJsonValueRef::toString(const QString & defaultValue);
extern void* C_ZNK13QJsonValueRef8toStringERK7QString(void* qthis, void* arg0); // 2
  // proto:  bool QJsonValueRef::isUndefined();
extern bool C_ZNK13QJsonValueRef11isUndefinedEv(void* qthis); // 2
  // proto:  QJsonValue::Type QJsonValueRef::type();
extern void C_ZNK13QJsonValueRef4typeEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::isObject();
extern bool C_ZNK13QJsonValueRef8isObjectEv(void* qthis); // 2
  // proto:  bool QJsonValueRef::isDouble();
extern bool C_ZNK13QJsonValueRef8isDoubleEv(void* qthis); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QJsonValuePtr)=24
type QJsonValuePtr struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QJsonValue)=24
type QJsonValue struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QJsonValueRef)=16
type QJsonValueRef struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2Ed(arg0)
    return &QJsonValue{Qclsinst:qthis}
  case 1:
    // invoke: _ZN10QJsonValueC1Eb
    // invoke: void QJsonValue(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2Eb(arg0)
    return &QJsonValue{Qclsinst:qthis}
  case 2:
    // invoke: _ZN10QJsonValueC1Ei
    // invoke: void QJsonValue(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2Ei(arg0)
    return &QJsonValue{Qclsinst:qthis}
  case 3:
    // invoke: _ZN10QJsonValueC1ERK7QString
    // invoke: void QJsonValue(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2ERK7QString(arg0)
    return &QJsonValue{Qclsinst:qthis}
  case 4:
    // invoke: _ZN10QJsonValueC1EPKc
    // invoke: void QJsonValue(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[4][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2EPKc(arg0)
    return &QJsonValue{Qclsinst:qthis}
  case 5:
    // invoke: _ZN10QJsonValueC1Ex
    // invoke: void QJsonValue(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QJsonValueC2Ex(arg0)
    return &QJsonValue{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QJsonValue", "QJsonValue", args)
  }

  return nil // QJsonValue{Qclsinst:qthis}
}

// fromVariant(const class QVariant &)
func (this *QJsonValue) Fromvariant_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QJsonValue11fromVariantERK8QVariant(arg0)
  default:
    qtrt.ErrorResolve("QJsonValue", "fromVariant", args)
  }

  return
}

// ~QJsonValue()
func (this *QJsonValue) Freeqjsonvalue(args ...interface{}) () {
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
    C.C_ZN10QJsonValueD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "~QJsonValue", args)
  }

  return
}

// toObject()
func (this *QJsonValue) Toobject(args ...interface{}) () {
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
    C.C_ZNK10QJsonValue8toObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "toObject", args)
  }

  return
}

// toBool(_Bool)
func (this *QJsonValue) Tobool(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue6toBoolEb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "toBool", args)
  }

  return
}

// toVariant()
func (this *QJsonValue) Tovariant(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue9toVariantEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "toVariant", args)
  }

  return
}

// toString(const class QString &)
func (this *QJsonValue) Tostring(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QJsonValue8toStringERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "toString", args)
  }

  return
}

// type()
func (this *QJsonValue) Type_(args ...interface{}) () {
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
    C.C_ZNK10QJsonValue4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "type", args)
  }

  return
}

// isUndefined()
func (this *QJsonValue) Isundefined(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue11isUndefinedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "isUndefined", args)
  }

  return
}

// isArray()
func (this *QJsonValue) Isarray(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue7isArrayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "isArray", args)
  }

  return
}

// isBool()
func (this *QJsonValue) Isbool(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue6isBoolEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "isBool", args)
  }

  return
}

// isObject()
func (this *QJsonValue) Isobject(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue8isObjectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "isObject", args)
  }

  return
}

// isDouble()
func (this *QJsonValue) Isdouble(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue8isDoubleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "isDouble", args)
  }

  return
}

// toArray()
func (this *QJsonValue) Toarray(args ...interface{}) () {
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
    C.C_ZNK10QJsonValue7toArrayEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "toArray", args)
  }

  return
}

// isString()
func (this *QJsonValue) Isstring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue8isStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "isString", args)
  }

  return
}

// toInt(int)
func (this *QJsonValue) Toint(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QJsonValue5toIntEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "toInt", args)
  }

  return
}

// isNull()
func (this *QJsonValue) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QJsonValue6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "isNull", args)
  }

  return
}

// toDouble(double)
func (this *QJsonValue) Todouble(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QJsonValue8toDoubleEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValue", "toDouble", args)
  }

  return
}

// isArray()
func (this *QJsonValueRef) Isarray(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef7isArrayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isArray", args)
  }

  return
}

// toArray()
func (this *QJsonValueRef) Toarray(args ...interface{}) () {
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
    C.C_ZNK13QJsonValueRef7toArrayEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toArray", args)
  }

  return
}

// isString()
func (this *QJsonValueRef) Isstring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef8isStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isString", args)
  }

  return
}

// toObject()
func (this *QJsonValueRef) Toobject(args ...interface{}) () {
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
    C.C_ZNK13QJsonValueRef8toObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toObject", args)
  }

  return
}

// toInt(int)
func (this *QJsonValueRef) Toint(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QJsonValueRef5toIntEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QJsonValueRef5toIntEv
    // invoke: int toInt()
    var ret0 = C.C_ZNK13QJsonValueRef5toIntEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toInt", args)
  }

  return
}

// isBool()
func (this *QJsonValueRef) Isbool(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef6isBoolEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isBool", args)
  }

  return
}

// isNull()
func (this *QJsonValueRef) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isNull", args)
  }

  return
}

// toDouble(double)
func (this *QJsonValueRef) Todouble(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QJsonValueRef8toDoubleEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QJsonValueRef8toDoubleEv
    // invoke: double toDouble()
    var ret0 = C.C_ZNK13QJsonValueRef8toDoubleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toDouble", args)
  }

  return
}

// toBool()
func (this *QJsonValueRef) Tobool(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef6toBoolEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QJsonValueRef6toBoolEb
    // invoke: bool toBool(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QJsonValueRef6toBoolEb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toBool", args)
  }

  return
}

// toVariant()
func (this *QJsonValueRef) Tovariant(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef9toVariantEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toVariant", args)
  }

  return
}

// toString()
func (this *QJsonValueRef) Tostring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef8toStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QJsonValueRef8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QJsonValueRef8toStringERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toString", args)
  }

  return
}

// isUndefined()
func (this *QJsonValueRef) Isundefined(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef11isUndefinedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isUndefined", args)
  }

  return
}

// type()
func (this *QJsonValueRef) Type_(args ...interface{}) () {
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
    C.C_ZNK13QJsonValueRef4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "type", args)
  }

  return
}

// isObject()
func (this *QJsonValueRef) Isobject(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef8isObjectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isObject", args)
  }

  return
}

// isDouble()
func (this *QJsonValueRef) Isdouble(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QJsonValueRef8isDoubleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isDouble", args)
  }

  return
}

// <= body block end

