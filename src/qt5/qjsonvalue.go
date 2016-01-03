package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  QJsonObject QJsonValue::toObject();
extern void _ZNK10QJsonValue8toObjectEv(void* qthis);
  // proto:  bool QJsonValue::isDouble();
extern void demth_ZNK10QJsonValue8isDoubleEv(void* qthis);
  // proto:  void QJsonValue::QJsonValue(const QString & s);
extern void* dector_ZN10QJsonValueC1ERK7QString(void* arg0);
extern void _ZN10QJsonValueC1ERK7QString(void* qthis, void* arg0);
  // proto:  int QJsonValue::toInt(int defaultValue);
extern void _ZNK10QJsonValue5toIntEi(void* qthis, int32_t arg0);
  // proto:  void QJsonValue::QJsonValue(const void * );
extern void* dector_ZN10QJsonValueC1EPKv(void* arg0);
extern void demth_ZN10QJsonValueC1EPKv(void* qthis, void* arg0);
  // proto:  QJsonArray QJsonValue::toArray();
extern void _ZNK10QJsonValue7toArrayEv(void* qthis);
  // proto:  bool QJsonValue::isArray();
extern void demth_ZNK10QJsonValue7isArrayEv(void* qthis);
  // proto:  void QJsonValue::QJsonValue(const char * s);
extern void* dector_ZN10QJsonValueC1EPKc(unsigned char* arg0);
extern void demth_ZN10QJsonValueC1EPKc(void* qthis, unsigned char* arg0);
  // proto:  QString QJsonValue::toString(const QString & defaultValue);
extern void _ZNK10QJsonValue8toStringERK7QString(void* qthis, void* arg0);
  // proto:  double QJsonValue::toDouble(double defaultValue);
extern void _ZNK10QJsonValue8toDoubleEd(void* qthis, double arg0);
  // proto:  void QJsonValue::~QJsonValue();
extern void _ZN10QJsonValueD0Ev(void* qthis);
  // proto:  QVariant QJsonValue::toVariant();
extern void _ZNK10QJsonValue9toVariantEv(void* qthis);
  // proto:  bool QJsonValue::isObject();
extern void demth_ZNK10QJsonValue8isObjectEv(void* qthis);
  // proto: static QJsonValue QJsonValue::fromVariant(const QVariant & variant);
extern void _ZN10QJsonValue11fromVariantERK8QVariant(void* arg0);
  // proto:  bool QJsonValue::toBool(bool defaultValue);
extern void _ZNK10QJsonValue6toBoolEb(void* qthis, bool arg0);
  // proto:  void QJsonValue::QJsonValue(double n);
extern void* dector_ZN10QJsonValueC1Ed(double arg0);
extern void _ZN10QJsonValueC1Ed(void* qthis, double arg0);
  // proto:  bool QJsonValue::isBool();
extern void demth_ZNK10QJsonValue6isBoolEv(void* qthis);
  // proto:  void QJsonValue::QJsonValue(bool b);
extern void* dector_ZN10QJsonValueC1Eb(bool arg0);
extern void _ZN10QJsonValueC1Eb(void* qthis, bool arg0);
  // proto:  bool QJsonValue::isUndefined();
extern void demth_ZNK10QJsonValue11isUndefinedEv(void* qthis);
  // proto:  bool QJsonValue::isNull();
extern void demth_ZNK10QJsonValue6isNullEv(void* qthis);
  // proto:  bool QJsonValue::isString();
extern void demth_ZNK10QJsonValue8isStringEv(void* qthis);
  // proto:  void QJsonValue::QJsonValue(int n);
extern void* dector_ZN10QJsonValueC1Ei(int32_t arg0);
extern void _ZN10QJsonValueC1Ei(void* qthis, int32_t arg0);
  // proto:  void QJsonValue::QJsonValue(qint64 n);
extern void* dector_ZN10QJsonValueC1Ex(int64_t arg0);
extern void _ZN10QJsonValueC1Ex(void* qthis, int64_t arg0);
  // proto:  QJsonArray QJsonValueRef::toArray();
extern void _ZNK13QJsonValueRef7toArrayEv(void* qthis);
  // proto:  QJsonObject QJsonValueRef::toObject();
extern void _ZNK13QJsonValueRef8toObjectEv(void* qthis);
  // proto:  bool QJsonValueRef::isBool();
extern void demth_ZNK13QJsonValueRef6isBoolEv(void* qthis);
  // proto:  bool QJsonValueRef::isDouble();
extern void demth_ZNK13QJsonValueRef8isDoubleEv(void* qthis);
  // proto:  double QJsonValueRef::toDouble();
extern void demth_ZNK13QJsonValueRef8toDoubleEv(void* qthis);
  // proto:  bool QJsonValueRef::toBool(bool defaultValue);
extern void demth_ZNK13QJsonValueRef6toBoolEb(void* qthis, bool arg0);
  // proto:  double QJsonValueRef::toDouble(double defaultValue);
extern void demth_ZNK13QJsonValueRef8toDoubleEd(void* qthis, double arg0);
  // proto:  bool QJsonValueRef::toBool();
extern void demth_ZNK13QJsonValueRef6toBoolEv(void* qthis);
  // proto:  QVariant QJsonValueRef::toVariant();
extern void _ZNK13QJsonValueRef9toVariantEv(void* qthis);
  // proto:  QString QJsonValueRef::toString(const QString & defaultValue);
extern void demth_ZNK13QJsonValueRef8toStringERK7QString(void* qthis, void* arg0);
  // proto:  bool QJsonValueRef::isObject();
extern void demth_ZNK13QJsonValueRef8isObjectEv(void* qthis);
  // proto:  bool QJsonValueRef::isString();
extern void demth_ZNK13QJsonValueRef8isStringEv(void* qthis);
  // proto:  QString QJsonValueRef::toString();
extern void demth_ZNK13QJsonValueRef8toStringEv(void* qthis);
  // proto:  int QJsonValueRef::toInt(int defaultValue);
extern void demth_ZNK13QJsonValueRef5toIntEi(void* qthis, int32_t arg0);
  // proto:  bool QJsonValueRef::isArray();
extern void demth_ZNK13QJsonValueRef7isArrayEv(void* qthis);
  // proto:  bool QJsonValueRef::isNull();
extern void demth_ZNK13QJsonValueRef6isNullEv(void* qthis);
  // proto:  int QJsonValueRef::toInt();
extern void demth_ZNK13QJsonValueRef5toIntEv(void* qthis);
  // proto:  bool QJsonValueRef::isUndefined();
extern void demth_ZNK13QJsonValueRef11isUndefinedEv(void* qthis);
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

  // proto:  QJsonObject QJsonValue::toObject();
func (this *QJsonValue) toObject(args ...interface{}) () {
  // toObject()
  // toObject(const class QJsonObject &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QJsonObject{}) // "const QJsonObject &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue8toObjectEv
    // invoke: QJsonObject toObject()
    C._ZNK10QJsonValue8toObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "toObject", args)
  }

}

  // proto:  bool QJsonValue::isDouble();
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
    C.demth_ZNK10QJsonValue8isDoubleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "isDouble", args)
  }

}

  // proto:  void QJsonValue::QJsonValue(const QString & s);
func NewQJsonValue(args ...interface{}) QJsonValue {
  return QJsonValue{}
}

  // proto:  int QJsonValue::toInt(int defaultValue);
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
    C._ZNK10QJsonValue5toIntEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonValue", "toInt", args)
  }

}

  // proto:  QJsonArray QJsonValue::toArray();
func (this *QJsonValue) toArray(args ...interface{}) () {
  // toArray(const class QJsonArray &)
  // toArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QJsonArray{}) // "const QJsonArray &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QJsonValue7toArrayEv
    // invoke: QJsonArray toArray()
    C._ZNK10QJsonValue7toArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "toArray", args)
  }

}

  // proto:  bool QJsonValue::isArray();
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
    C.demth_ZNK10QJsonValue7isArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "isArray", args)
  }

}

  // proto:  QString QJsonValue::toString(const QString & defaultValue);
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
    C._ZNK10QJsonValue8toStringERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonValue", "toString", args)
  }

}

  // proto:  double QJsonValue::toDouble(double defaultValue);
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
    C._ZNK10QJsonValue8toDoubleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonValue", "toDouble", args)
  }

}

  // proto:  void QJsonValue::~QJsonValue();
func (this *QJsonValue) FreeQJsonValue(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QJsonValue", "~QJsonValue", args)
  }

}

  // proto:  QVariant QJsonValue::toVariant();
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
    C._ZNK10QJsonValue9toVariantEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "toVariant", args)
  }

}

  // proto:  bool QJsonValue::isObject();
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
    C.demth_ZNK10QJsonValue8isObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "isObject", args)
  }

}

  // proto: static QJsonValue QJsonValue::fromVariant(const QVariant & variant);
func (this *QJsonValue) fromVariant_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QJsonValue", "fromVariant", args)
  }

}

  // proto:  bool QJsonValue::toBool(bool defaultValue);
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
    C._ZNK10QJsonValue6toBoolEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonValue", "toBool", args)
  }

}

  // proto:  bool QJsonValue::isBool();
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
    C.demth_ZNK10QJsonValue6isBoolEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "isBool", args)
  }

}

  // proto:  bool QJsonValue::isUndefined();
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
    C.demth_ZNK10QJsonValue11isUndefinedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "isUndefined", args)
  }

}

  // proto:  bool QJsonValue::isNull();
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
    C.demth_ZNK10QJsonValue6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "isNull", args)
  }

}

  // proto:  bool QJsonValue::isString();
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
    C.demth_ZNK10QJsonValue8isStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValue", "isString", args)
  }

}

  // proto:  QJsonArray QJsonValueRef::toArray();
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
    C._ZNK13QJsonValueRef7toArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toArray", args)
  }

}

  // proto:  QJsonObject QJsonValueRef::toObject();
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
    C._ZNK13QJsonValueRef8toObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toObject", args)
  }

}

  // proto:  bool QJsonValueRef::isBool();
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
    C.demth_ZNK13QJsonValueRef6isBoolEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isBool", args)
  }

}

  // proto:  bool QJsonValueRef::isDouble();
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
    C.demth_ZNK13QJsonValueRef8isDoubleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isDouble", args)
  }

}

  // proto:  double QJsonValueRef::toDouble();
func (this *QJsonValueRef) toDouble(args ...interface{}) () {
  // toDouble()
  // toDouble(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef8toDoubleEv
    // invoke: double toDouble()
    C.demth_ZNK13QJsonValueRef8toDoubleEv(this.qclsinst)
  case 1:
    // invoke: _ZNK13QJsonValueRef8toDoubleEd
    // invoke: double toDouble(double)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.demth_ZNK13QJsonValueRef8toDoubleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toDouble", args)
  }

}

  // proto:  bool QJsonValueRef::toBool(bool defaultValue);
func (this *QJsonValueRef) toBool(args ...interface{}) () {
  // toBool(_Bool)
  // toBool()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef6toBoolEb
    // invoke: bool toBool(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.demth_ZNK13QJsonValueRef6toBoolEb(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK13QJsonValueRef6toBoolEv
    // invoke: bool toBool()
    C.demth_ZNK13QJsonValueRef6toBoolEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toBool", args)
  }

}

  // proto:  QVariant QJsonValueRef::toVariant();
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
    C._ZNK13QJsonValueRef9toVariantEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toVariant", args)
  }

}

  // proto:  QString QJsonValueRef::toString(const QString & defaultValue);
func (this *QJsonValueRef) toString(args ...interface{}) () {
  // toString(const class QString &)
  // toString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonValueRef8toStringERK7QString
    // invoke: QString toString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK13QJsonValueRef8toStringERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK13QJsonValueRef8toStringEv
    // invoke: QString toString()
    C.demth_ZNK13QJsonValueRef8toStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toString", args)
  }

}

  // proto:  bool QJsonValueRef::isObject();
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
    C.demth_ZNK13QJsonValueRef8isObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isObject", args)
  }

}

  // proto:  bool QJsonValueRef::isString();
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
    C.demth_ZNK13QJsonValueRef8isStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isString", args)
  }

}

  // proto:  int QJsonValueRef::toInt(int defaultValue);
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
    C.demth_ZNK13QJsonValueRef5toIntEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK13QJsonValueRef5toIntEv
    // invoke: int toInt()
    C.demth_ZNK13QJsonValueRef5toIntEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toInt", args)
  }

}

  // proto:  bool QJsonValueRef::isArray();
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
    C.demth_ZNK13QJsonValueRef7isArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isArray", args)
  }

}

  // proto:  bool QJsonValueRef::isNull();
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
    C.demth_ZNK13QJsonValueRef6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isNull", args)
  }

}

  // proto:  bool QJsonValueRef::isUndefined();
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
    C.demth_ZNK13QJsonValueRef11isUndefinedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isUndefined", args)
  }

}

// <= body block end

