package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qjsondocument.h
// dst-file: /src/core/qjsondocument.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QJsonDocument QJsonDocument::fromVariant(const QVariant & variant);
extern void C_ZN13QJsonDocument11fromVariantERK8QVariant(void* arg0); // 4
  // proto:  void QJsonDocument::~QJsonDocument();
extern void C_ZN13QJsonDocumentD2Ev(void* qthis); // 4
  // proto:  void QJsonDocument::QJsonDocument();
extern void* C_ZN13QJsonDocumentC2Ev(); // 3
  // proto:  QJsonObject QJsonDocument::object();
extern void C_ZNK13QJsonDocument6objectEv(void* qthis); // 4
  // proto:  bool QJsonDocument::isArray();
extern bool C_ZNK13QJsonDocument7isArrayEv(void* qthis); // 4
  // proto:  bool QJsonDocument::isNull();
extern bool C_ZNK13QJsonDocument6isNullEv(void* qthis); // 4
  // proto:  bool QJsonDocument::isEmpty();
extern bool C_ZNK13QJsonDocument7isEmptyEv(void* qthis); // 4
  // proto:  const char * QJsonDocument::rawData(int * size);
extern void* C_ZNK13QJsonDocument7rawDataEPi(void* qthis, void* arg0); // 4
  // proto:  QVariant QJsonDocument::toVariant();
extern void* C_ZNK13QJsonDocument9toVariantEv(void* qthis); // 4
  // proto:  QJsonArray QJsonDocument::array();
extern void C_ZNK13QJsonDocument5arrayEv(void* qthis); // 4
  // proto:  QByteArray QJsonDocument::toJson();
extern void* C_ZNK13QJsonDocument6toJsonEv(void* qthis); // 4
  // proto:  QByteArray QJsonDocument::toBinaryData();
extern void* C_ZNK13QJsonDocument12toBinaryDataEv(void* qthis); // 4
  // proto:  bool QJsonDocument::isObject();
extern bool C_ZNK13QJsonDocument8isObjectEv(void* qthis); // 4
  // proto:  QString QJsonParseError::errorString();
extern void* C_ZNK15QJsonParseError11errorStringEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QJsonDocument)=8
type QJsonDocument struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QJsonParseError)=8
type QJsonParseError struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// fromVariant(const class QVariant &)
func (this *QJsonDocument) FromVariant_s(args ...interface{}) () {
  // fromVariant(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QJsonDocument11fromVariantERK8QVariant
    // invoke: QJsonDocument fromVariant(const class QVariant &)
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QJsonDocument11fromVariantERK8QVariant(arg0)
  default:
    qtrt.ErrorResolve("QJsonDocument", "fromVariant", args)
  }

  return
}

// ~QJsonDocument()
func (this *QJsonDocument) Free(args ...interface{}) () {
  // ~QJsonDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QJsonDocumentD0Ev
    // invoke: void ~QJsonDocument()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN13QJsonDocumentD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "~QJsonDocument", args)
  }

  return
}

// QJsonDocument()
func GcfreeQJsonDocument(this *QJsonDocument) {
  qtrt.UniverseFree(this)
}
func NewQJsonDocument(args ...interface{}) *QJsonDocument {
  // QJsonDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QJsonDocumentC1Ev
    // invoke: void QJsonDocument()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QJsonDocumentC2Ev()
    this := &QJsonDocument{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQJsonDocument)
    return this
  default:
    qtrt.ErrorResolve("QJsonDocument", "QJsonDocument", args)
  }

  return nil // QJsonDocument{Qclsinst:qthis}
}

// object()
func (this *QJsonDocument) Object(args ...interface{}) () {
  // object()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument6objectEv
    // invoke: QJsonObject object()
    C.C_ZNK13QJsonDocument6objectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonDocument", "object", args)
  }

  return
}

// isArray()
func (this *QJsonDocument) IsArray(args ...interface{}) (ret interface{}) {
  // isArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument7isArrayEv
    // invoke: bool isArray()
    var ret0 = C.C_ZNK13QJsonDocument7isArrayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "isArray", args)
  }

  return
}

// isNull()
func (this *QJsonDocument) IsNull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK13QJsonDocument6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "isNull", args)
  }

  return
}

// isEmpty()
func (this *QJsonDocument) IsEmpty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK13QJsonDocument7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "isEmpty", args)
  }

  return
}

// rawData(int *)
func (this *QJsonDocument) RawData(args ...interface{}) (ret interface{}) {
  // rawData(int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument7rawDataEPi
    // invoke: const char * rawData(int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QJsonDocument7rawDataEPi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "rawData", args)
  }

  return
}

// toVariant()
func (this *QJsonDocument) ToVariant(args ...interface{}) (ret interface{}) {
  // toVariant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument9toVariantEv
    // invoke: QVariant toVariant()
    var ret0 = C.C_ZNK13QJsonDocument9toVariantEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "toVariant", args)
  }

  return
}

// array()
func (this *QJsonDocument) Array(args ...interface{}) () {
  // array()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument5arrayEv
    // invoke: QJsonArray array()
    C.C_ZNK13QJsonDocument5arrayEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonDocument", "array", args)
  }

  return
}

// toJson()
func (this *QJsonDocument) ToJson(args ...interface{}) (ret interface{}) {
  // toJson()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument6toJsonEv
    // invoke: QByteArray toJson()
    var ret0 = C.C_ZNK13QJsonDocument6toJsonEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "toJson", args)
  }

  return
}

// toBinaryData()
func (this *QJsonDocument) ToBinaryData(args ...interface{}) (ret interface{}) {
  // toBinaryData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument12toBinaryDataEv
    // invoke: QByteArray toBinaryData()
    var ret0 = C.C_ZNK13QJsonDocument12toBinaryDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "toBinaryData", args)
  }

  return
}

// isObject()
func (this *QJsonDocument) IsObject(args ...interface{}) (ret interface{}) {
  // isObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument8isObjectEv
    // invoke: bool isObject()
    var ret0 = C.C_ZNK13QJsonDocument8isObjectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonDocument", "isObject", args)
  }

  return
}

// errorString()
func (this *QJsonParseError) ErrorString(args ...interface{}) (ret interface{}) {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QJsonParseError11errorStringEv
    // invoke: QString errorString()
    var ret0 = C.C_ZNK15QJsonParseError11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonParseError", "errorString", args)
  }

  return
}

// <= body block end

