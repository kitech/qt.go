package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  QJsonObject QJsonDocument::object();
extern void _ZNK13QJsonDocument6objectEv(void* qthis);
  // proto: static QJsonDocument QJsonDocument::fromVariant(const QVariant & variant);
extern void _ZN13QJsonDocument11fromVariantERK8QVariant(void* arg0);
  // proto:  QJsonArray QJsonDocument::array();
extern void _ZNK13QJsonDocument5arrayEv(void* qthis);
  // proto:  QByteArray QJsonDocument::toJson();
extern void _ZNK13QJsonDocument6toJsonEv(void* qthis);
  // proto:  bool QJsonDocument::isNull();
extern void _ZNK13QJsonDocument6isNullEv(void* qthis);
  // proto:  void QJsonDocument::QJsonDocument();
extern void* dector_ZN13QJsonDocumentC1Ev();
extern void _ZN13QJsonDocumentC1Ev(void* qthis);
  // proto:  QVariant QJsonDocument::toVariant();
extern void _ZNK13QJsonDocument9toVariantEv(void* qthis);
  // proto:  bool QJsonDocument::isEmpty();
extern void _ZNK13QJsonDocument7isEmptyEv(void* qthis);
  // proto:  const char * QJsonDocument::rawData(int * size);
extern void _ZNK13QJsonDocument7rawDataEPi(void* qthis, int* arg0);
  // proto:  bool QJsonDocument::isObject();
extern void _ZNK13QJsonDocument8isObjectEv(void* qthis);
  // proto:  void QJsonDocument::~QJsonDocument();
extern void _ZN13QJsonDocumentD0Ev(void* qthis);
  // proto:  bool QJsonDocument::isArray();
extern void _ZNK13QJsonDocument7isArrayEv(void* qthis);
  // proto:  QByteArray QJsonDocument::toBinaryData();
extern void _ZNK13QJsonDocument12toBinaryDataEv(void* qthis);
  // proto:  QString QJsonParseError::errorString();
extern void _ZNK15QJsonParseError11errorStringEv(void* qthis);
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

// class sizeof(QJsonDocument)=8
type QJsonDocument struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QJsonParseError)=8
type QJsonParseError struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QJsonObject QJsonDocument::object();
func (this *QJsonDocument) object(args ...interface{}) () {
  // object()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument6objectEv
  default:
    qtrt.ErrorResolve("QJsonDocument", "object", args)
  }

}

  // proto: static QJsonDocument QJsonDocument::fromVariant(const QVariant & variant);
func (this *QJsonDocument) fromVariant_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QJsonDocument", "fromVariant", args)
  }

}

  // proto:  QJsonArray QJsonDocument::array();
func (this *QJsonDocument) array(args ...interface{}) () {
  // array()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument5arrayEv
  default:
    qtrt.ErrorResolve("QJsonDocument", "array", args)
  }

}

  // proto:  QByteArray QJsonDocument::toJson();
func (this *QJsonDocument) toJson(args ...interface{}) () {
  // toJson()
  // toJson(enum QJsonDocument::JsonFormat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QJsonDocument::JsonFormat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument6toJsonEv
  case 1:
    // invoke: _ZNK13QJsonDocument6toJsonENS_10JsonFormatE
    var arg0 = C.int(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QJsonDocument", "toJson", args)
  }

}

  // proto:  bool QJsonDocument::isNull();
func (this *QJsonDocument) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument6isNullEv
  default:
    qtrt.ErrorResolve("QJsonDocument", "isNull", args)
  }

}

  // proto:  void QJsonDocument::QJsonDocument();
func NewQJsonDocument(args ...interface{}) QJsonDocument {
  return QJsonDocument{}
}

  // proto:  QVariant QJsonDocument::toVariant();
func (this *QJsonDocument) toVariant(args ...interface{}) () {
  // toVariant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument9toVariantEv
  default:
    qtrt.ErrorResolve("QJsonDocument", "toVariant", args)
  }

}

  // proto:  bool QJsonDocument::isEmpty();
func (this *QJsonDocument) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument7isEmptyEv
  default:
    qtrt.ErrorResolve("QJsonDocument", "isEmpty", args)
  }

}

  // proto:  const char * QJsonDocument::rawData(int * size);
func (this *QJsonDocument) rawData(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QJsonDocument", "rawData", args)
  }

}

  // proto:  bool QJsonDocument::isObject();
func (this *QJsonDocument) isObject(args ...interface{}) () {
  // isObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument8isObjectEv
  default:
    qtrt.ErrorResolve("QJsonDocument", "isObject", args)
  }

}

  // proto:  void QJsonDocument::~QJsonDocument();
func (this *QJsonDocument) FreeQJsonDocument(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QJsonDocument", "~QJsonDocument", args)
  }

}

  // proto:  bool QJsonDocument::isArray();
func (this *QJsonDocument) isArray(args ...interface{}) () {
  // isArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument7isArrayEv
  default:
    qtrt.ErrorResolve("QJsonDocument", "isArray", args)
  }

}

  // proto:  QByteArray QJsonDocument::toBinaryData();
func (this *QJsonDocument) toBinaryData(args ...interface{}) () {
  // toBinaryData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QJsonDocument12toBinaryDataEv
  default:
    qtrt.ErrorResolve("QJsonDocument", "toBinaryData", args)
  }

}

  // proto:  QString QJsonParseError::errorString();
func (this *QJsonParseError) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QJsonParseError11errorStringEv
  default:
    qtrt.ErrorResolve("QJsonParseError", "errorString", args)
  }

}

// <= body block end

