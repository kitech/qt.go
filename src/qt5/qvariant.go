package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qvariant.h
// dst-file: /src/core/qvariant.go
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
// class sizeof(QVariantComparisonHelper)=8
type QVariantComparisonHelper struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QVariant)=16
type QVariant struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QSequentialIterable)=104
type QSequentialIterable struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAssociativeIterable)=112
type QAssociativeIterable struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQVariantComparisonHelper(args ...interface{}) QVariantComparisonHelper {
  return QVariantComparisonHelper{}
}


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
  default:
    qtrt.ErrorResolve("QVariant", "toDouble", args)
  }

}


func NewQVariant(args ...interface{}) QVariant {
  return QVariant{}
}


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
  default:
    qtrt.ErrorResolve("QVariant", "toLongLong", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toSize", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toString", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toReal", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toFloat", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toByteArray", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toLocale", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toUrl", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toLine", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "typeName", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toJsonArray", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toStringList", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toList", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toUInt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toUuid", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toJsonDocument", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toPoint", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toInt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "isValid", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "detach", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toModelIndex", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toHash", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toMap", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "canConvert", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "userType", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "constData", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toPersistentModelIndex", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toLineF", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toJsonObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "load", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toChar", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "isNull", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toRectF", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toDate", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "save", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toTime", args)
  }

}


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
    // invoke: _ZNK8QVariant4dataEv
  case 1:
    // invoke: _ZN8QVariant4dataEv
  default:
    qtrt.ErrorResolve("QVariant", "data", args)
  }

}


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
    // invoke: _ZNK8QVariant7convertEiPv
  case 1:
    // invoke: _ZN8QVariant7convertEi
  default:
    qtrt.ErrorResolve("QVariant", "convert", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toRegExp", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toPointF", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toSizeF", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "swap", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "clear", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toRect", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toBool", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toULongLong", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toJsonValue", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toDateTime", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "isDetached", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toEasingCurve", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toBitArray", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QVariant", "toRegularExpression", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSequentialIterable", "size", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSequentialIterable", "canReverseIterate", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSequentialIterable", "at", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "size", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAssociativeIterable", "value", args)
  }

}

// <= body block end

