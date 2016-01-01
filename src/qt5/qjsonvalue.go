package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qjsonvalue.h
// dst-file: /src/core/qjsonvalue.go
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
// class sizeof(QJsonValueRefPtr)=16
type QJsonValueRefPtr struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QJsonValuePtr)=24
type QJsonValuePtr struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QJsonValue)=24
type QJsonValue struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QJsonValueRef)=16
type QJsonValueRef struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  case 1:
    // invoke: _ZNK10QJsonValue8toObjectERK11QJsonObject
  default:
    qtrt.ErrorResolve("QJsonValue", "toObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "isDouble", args)
  }

}


func NewQJsonValue(args ...interface{}) QJsonValue {
  return QJsonValue{}
}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "toInt", args)
  }

}


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
    // invoke: _ZNK10QJsonValue7toArrayERK10QJsonArray
  case 1:
    // invoke: _ZNK10QJsonValue7toArrayEv
  default:
    qtrt.ErrorResolve("QJsonValue", "toArray", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "isArray", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "toString", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "toDouble", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "toVariant", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "isObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "toBool", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "isBool", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "isUndefined", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "isNull", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValue", "isString", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toArray", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isBool", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isDouble", args)
  }

}


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
  case 1:
    // invoke: _ZNK13QJsonValueRef8toDoubleEd
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toDouble", args)
  }

}


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
  case 1:
    // invoke: _ZNK13QJsonValueRef6toBoolEv
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toBool", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toVariant", args)
  }

}


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
  case 1:
    // invoke: _ZNK13QJsonValueRef8toStringEv
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toString", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isString", args)
  }

}


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
  case 1:
    // invoke: _ZNK13QJsonValueRef5toIntEv
  default:
    qtrt.ErrorResolve("QJsonValueRef", "toInt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isArray", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isNull", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QJsonValueRef", "isUndefined", args)
  }

}

// <= body block end

