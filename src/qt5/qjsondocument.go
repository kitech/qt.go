package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qjsondocument.h
// dst-file: /src/core/qjsondocument.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
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
  default:
    qtrt.ErrorResolve("QJsonDocument", "toJson", args)
  }

}


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


func NewQJsonDocument(args ...interface{}) QJsonDocument {
  return QJsonDocument{}
}


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
  default:
    qtrt.ErrorResolve("QJsonDocument", "rawData", args)
  }

}


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

