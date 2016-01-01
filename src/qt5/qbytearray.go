package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qbytearray.h
// dst-file: /src/core/qbytearray.go
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
// class sizeof(QByteRef)=16
type QByteRef struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QByteArray)=8
type QByteArray struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QByteArrayDataPtr)=8
type QByteArrayDataPtr struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQByteRef(args ...interface{})() {
}


func (this *QByteArray) insert(args ...interface{}) () {
  // insert(int, char)
  // insert(int, const char *, int)
  // insert(int, const class QString &)
  // insert(int, const class QByteArray &)
  // insert(int, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6insertEic
  case 1:
    // invoke: _ZN10QByteArray6insertEiPKci
  case 2:
    // invoke: _ZN10QByteArray6insertEiRK7QString
  case 3:
    // invoke: _ZN10QByteArray6insertEiRKS_
  case 4:
    // invoke: _ZN10QByteArray6insertEiPKc
  default:
    qtrt.ErrorResolve("QByteArray", "insert", args)
 }

}


func (this *QByteArray) cend(args ...interface{}) () {
  // cend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4cendEv
  default:
    qtrt.ErrorResolve("QByteArray", "cend", args)
 }

}


func (this *QByteArray) lastIndexOf(args ...interface{}) () {
  // lastIndexOf(const class QByteArray &, int)
  // lastIndexOf(const class QString &, int)
  // lastIndexOf(char, int)
  // lastIndexOf(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11lastIndexOfERKS_i
  case 1:
    // invoke: _ZNK10QByteArray11lastIndexOfERK7QStringi
  case 2:
    // invoke: _ZNK10QByteArray11lastIndexOfEci
  case 3:
    // invoke: _ZNK10QByteArray11lastIndexOfEPKci
  default:
    qtrt.ErrorResolve("QByteArray", "lastIndexOf", args)
 }

}


func (this *QByteArray) push_front(args ...interface{}) () {
  // push_front(const class QByteArray &)
  // push_front(char)
  // push_front(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10push_frontERKS_
  case 1:
    // invoke: _ZN10QByteArray10push_frontEc
  case 2:
    // invoke: _ZN10QByteArray10push_frontEPKc
  default:
    qtrt.ErrorResolve("QByteArray", "push_front", args)
 }

}


func (this *QByteArray) toULong(args ...interface{}) () {
  // toULong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toULongEPbi
  default:
    qtrt.ErrorResolve("QByteArray", "toULong", args)
 }

}


func (this *QByteArray) replace(args ...interface{}) () {
  // replace(const char *, const char *)
  // replace(const class QByteArray &, const class QByteArray &)
  // replace(const char *, const class QByteArray &)
  // replace(int, int, const char *, int)
  // replace(char, const class QByteArray &)
  // replace(char, char)
  // replace(int, int, const class QByteArray &)
  // replace(char, const class QString &)
  // replace(const class QString &, const char *)
  // replace(const char *, int, const char *, int)
  // replace(char, const char *)
  // replace(int, int, const char *)
  // replace(const class QByteArray &, const char *)
  // replace(const class QString &, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.ByteTy(true) // "const char *"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(false) // "char"
  vtys[4][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(false) // "char"
  vtys[5][1] = qtrt.ByteTy(false) // "char"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "int"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[6][2] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.ByteTy(false) // "char"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][1] = qtrt.ByteTy(true) // "const char *"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = qtrt.Int32Ty(false) // "int"
  vtys[9][2] = qtrt.ByteTy(true) // "const char *"
  vtys[9][3] = qtrt.Int32Ty(false) // "int"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(false) // "char"
  vtys[10][1] = qtrt.ByteTy(true) // "const char *"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "int"
  vtys[11][1] = qtrt.Int32Ty(false) // "int"
  vtys[11][2] = qtrt.ByteTy(true) // "const char *"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[12][1] = qtrt.ByteTy(true) // "const char *"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[13][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7replaceEPKcS1_
  case 1:
    // invoke: _ZN10QByteArray7replaceERKS_S1_
  case 2:
    // invoke: _ZN10QByteArray7replaceEPKcRKS_
  case 3:
    // invoke: _ZN10QByteArray7replaceEiiPKci
  case 4:
    // invoke: _ZN10QByteArray7replaceEcRKS_
  case 5:
    // invoke: _ZN10QByteArray7replaceEcc
  case 6:
    // invoke: _ZN10QByteArray7replaceEiiRKS_
  case 7:
    // invoke: _ZN10QByteArray7replaceEcRK7QString
  case 8:
    // invoke: _ZN10QByteArray7replaceERK7QStringPKc
  case 9:
    // invoke: _ZN10QByteArray7replaceEPKciS1_i
  case 10:
    // invoke: _ZN10QByteArray7replaceEcPKc
  case 11:
    // invoke: _ZN10QByteArray7replaceEiiPKc
  case 12:
    // invoke: _ZN10QByteArray7replaceERKS_PKc
  case 13:
    // invoke: _ZN10QByteArray7replaceERK7QStringRKS_
  default:
    qtrt.ErrorResolve("QByteArray", "replace", args)
 }

}


func (this *QByteArray) fromHex_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromHex", args)
 }

}


func (this *QByteArray) prepend(args ...interface{}) () {
  // prepend(const char *)
  // prepend(char)
  // prepend(const class QByteArray &)
  // prepend(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7prependEPKc
  case 1:
    // invoke: _ZN10QByteArray7prependEc
  case 2:
    // invoke: _ZN10QByteArray7prependERKS_
  case 3:
    // invoke: _ZN10QByteArray7prependEPKci
  default:
    qtrt.ErrorResolve("QByteArray", "prepend", args)
 }

}


func (this *QByteArray) count(args ...interface{}) () {
  // count(const class QByteArray &)
  // count()
  // count(char)
  // count(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5countERKS_
  case 1:
    // invoke: _ZNK10QByteArray5countEv
  case 2:
    // invoke: _ZNK10QByteArray5countEc
  case 3:
    // invoke: _ZNK10QByteArray5countEPKc
  default:
    qtrt.ErrorResolve("QByteArray", "count", args)
 }

}


func (this *QByteArray) FreeQByteArray(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "~QByteArray", args)
 }

}


func (this *QByteArray) end(args ...interface{}) () {
  // end()
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray3endEv
  case 1:
    // invoke: _ZNK10QByteArray3endEv
  default:
    qtrt.ErrorResolve("QByteArray", "end", args)
 }

}


func NewQByteArray(args ...interface{})() {
}


func (this *QByteArray) toFloat(args ...interface{}) () {
  // toFloat(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toFloatEPb
  default:
    qtrt.ErrorResolve("QByteArray", "toFloat", args)
 }

}


func (this *QByteArray) truncate(args ...interface{}) () {
  // truncate(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray8truncateEi
  default:
    qtrt.ErrorResolve("QByteArray", "truncate", args)
 }

}


func (this *QByteArray) toBase64(args ...interface{}) () {
  // toBase64(Base64Options)
  // toBase64()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "Base64Options"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toBase64E6QFlagsINS_12Base64OptionEE
  case 1:
    // invoke: _ZNK10QByteArray8toBase64Ev
  default:
    qtrt.ErrorResolve("QByteArray", "toBase64", args)
 }

}


func (this *QByteArray) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7isEmptyEv
  default:
    qtrt.ErrorResolve("QByteArray", "isEmpty", args)
 }

}


func (this *QByteArray) resize(args ...interface{}) () {
  // resize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6resizeEi
  default:
    qtrt.ErrorResolve("QByteArray", "resize", args)
 }

}


func (this *QByteArray) toHex(args ...interface{}) () {
  // toHex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5toHexEv
  default:
    qtrt.ErrorResolve("QByteArray", "toHex", args)
 }

}


func (this *QByteArray) indexOf(args ...interface{}) () {
  // indexOf(const char *, int)
  // indexOf(const class QByteArray &, int)
  // indexOf(char, int)
  // indexOf(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7indexOfEPKci
  case 1:
    // invoke: _ZNK10QByteArray7indexOfERKS_i
  case 2:
    // invoke: _ZNK10QByteArray7indexOfEci
  case 3:
    // invoke: _ZNK10QByteArray7indexOfERK7QStringi
  default:
    qtrt.ErrorResolve("QByteArray", "indexOf", args)
 }

}


func (this *QByteArray) toUInt(args ...interface{}) () {
  // toUInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6toUIntEPbi
  default:
    qtrt.ErrorResolve("QByteArray", "toUInt", args)
 }

}


func (this *QByteArray) fromStdString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromStdString", args)
 }

}


func (this *QByteArray) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6isNullEv
  default:
    qtrt.ErrorResolve("QByteArray", "isNull", args)
 }

}


func (this *QByteArray) reserve(args ...interface{}) () {
  // reserve(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7reserveEi
  default:
    qtrt.ErrorResolve("QByteArray", "reserve", args)
 }

}


func (this *QByteArray) cbegin(args ...interface{}) () {
  // cbegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6cbeginEv
  default:
    qtrt.ErrorResolve("QByteArray", "cbegin", args)
 }

}


func (this *QByteArray) fromRawData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromRawData", args)
 }

}


func (this *QByteArray) contains(args ...interface{}) () {
  // contains(char)
  // contains(const char *)
  // contains(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8containsEc
  case 1:
    // invoke: _ZNK10QByteArray8containsEPKc
  case 2:
    // invoke: _ZNK10QByteArray8containsERKS_
  default:
    qtrt.ErrorResolve("QByteArray", "contains", args)
 }

}


func (this *QByteArray) toLong(args ...interface{}) () {
  // toLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6toLongEPbi
  default:
    qtrt.ErrorResolve("QByteArray", "toLong", args)
 }

}


func (this *QByteArray) data(args ...interface{}) () {
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
    // invoke: _ZN10QByteArray4dataEv
  case 1:
    // invoke: _ZNK10QByteArray4dataEv
  default:
    qtrt.ErrorResolve("QByteArray", "data", args)
 }

}


func (this *QByteArray) number_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "number", args)
 }

}


func (this *QByteArray) capacity(args ...interface{}) () {
  // capacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8capacityEv
  default:
    qtrt.ErrorResolve("QByteArray", "capacity", args)
 }

}


func (this *QByteArray) fromBase64_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromBase64", args)
 }

}


func (this *QByteArray) left(args ...interface{}) () {
  // left(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4leftEi
  default:
    qtrt.ErrorResolve("QByteArray", "left", args)
 }

}


func (this *QByteArray) append(args ...interface{}) () {
  // append(char)
  // append(const class QByteArray &)
  // append(const char *)
  // append(const class QString &)
  // append(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(true) // "const char *"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6appendEc
  case 1:
    // invoke: _ZN10QByteArray6appendERKS_
  case 2:
    // invoke: _ZN10QByteArray6appendEPKc
  case 3:
    // invoke: _ZN10QByteArray6appendERK7QString
  case 4:
    // invoke: _ZN10QByteArray6appendEPKci
  default:
    qtrt.ErrorResolve("QByteArray", "append", args)
 }

}


func (this *QByteArray) startsWith(args ...interface{}) () {
  // startsWith(const char *)
  // startsWith(const class QByteArray &)
  // startsWith(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10startsWithEPKc
  case 1:
    // invoke: _ZNK10QByteArray10startsWithERKS_
  case 2:
    // invoke: _ZNK10QByteArray10startsWithEc
  default:
    qtrt.ErrorResolve("QByteArray", "startsWith", args)
 }

}


func (this *QByteArray) remove(args ...interface{}) () {
  // remove(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6removeEii
  default:
    qtrt.ErrorResolve("QByteArray", "remove", args)
 }

}


func (this *QByteArray) endsWith(args ...interface{}) () {
  // endsWith(const char *)
  // endsWith(char)
  // endsWith(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8endsWithEPKc
  case 1:
    // invoke: _ZNK10QByteArray8endsWithEc
  case 2:
    // invoke: _ZNK10QByteArray8endsWithERKS_
  default:
    qtrt.ErrorResolve("QByteArray", "endsWith", args)
 }

}


func (this *QByteArray) squeeze(args ...interface{}) () {
  // squeeze()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7squeezeEv
  default:
    qtrt.ErrorResolve("QByteArray", "squeeze", args)
 }

}


func (this *QByteArray) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6detachEv
  default:
    qtrt.ErrorResolve("QByteArray", "detach", args)
 }

}


func (this *QByteArray) repeated(args ...interface{}) () {
  // repeated(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8repeatedEi
  default:
    qtrt.ErrorResolve("QByteArray", "repeated", args)
 }

}


func (this *QByteArray) setNum(args ...interface{}) () {
  // setNum(float, char, int)
  // setNum(short, int)
  // setNum(double, char, int)
  // setNum(ushort, int)
  // setNum(qlonglong, int)
  // setNum(qulonglong, int)
  // setNum(uint, int)
  // setNum(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int16Ty(false) // "short"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "double"
  vtys[2][1] = qtrt.ByteTy(false) // "char"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "uint"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6setNumEfci
  case 1:
    // invoke: _ZN10QByteArray6setNumEsi
  case 2:
    // invoke: _ZN10QByteArray6setNumEdci
  case 3:
    // invoke: _ZN10QByteArray6setNumEti
  case 4:
    // invoke: _ZN10QByteArray6setNumExi
  case 5:
    // invoke: _ZN10QByteArray6setNumEyi
  case 6:
    // invoke: _ZN10QByteArray6setNumEji
  case 7:
    // invoke: _ZN10QByteArray6setNumEii
  default:
    qtrt.ErrorResolve("QByteArray", "setNum", args)
 }

}


func (this *QByteArray) toShort(args ...interface{}) () {
  // toShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toShortEPbi
  default:
    qtrt.ErrorResolve("QByteArray", "toShort", args)
 }

}


func (this *QByteArray) toInt(args ...interface{}) () {
  // toInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5toIntEPbi
  default:
    qtrt.ErrorResolve("QByteArray", "toInt", args)
 }

}


func (this *QByteArray) constBegin(args ...interface{}) () {
  // constBegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10constBeginEv
  default:
    qtrt.ErrorResolve("QByteArray", "constBegin", args)
 }

}


func (this *QByteArray) push_back(args ...interface{}) () {
  // push_back(char)
  // push_back(const class QByteArray &)
  // push_back(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray9push_backEc
  case 1:
    // invoke: _ZN10QByteArray9push_backERKS_
  case 2:
    // invoke: _ZN10QByteArray9push_backEPKc
  default:
    qtrt.ErrorResolve("QByteArray", "push_back", args)
 }

}


func (this *QByteArray) isSharedWith(args ...interface{}) () {
  // isSharedWith(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray12isSharedWithERKS_
  default:
    qtrt.ErrorResolve("QByteArray", "isSharedWith", args)
 }

}


func (this *QByteArray) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4sizeEv
  default:
    qtrt.ErrorResolve("QByteArray", "size", args)
 }

}


func (this *QByteArray) leftJustified(args ...interface{}) () {
  // leftJustified(int, char, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray13leftJustifiedEicb
  default:
    qtrt.ErrorResolve("QByteArray", "leftJustified", args)
 }

}


func (this *QByteArray) begin(args ...interface{}) () {
  // begin()
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5beginEv
  case 1:
    // invoke: _ZN10QByteArray5beginEv
  default:
    qtrt.ErrorResolve("QByteArray", "begin", args)
 }

}


func (this *QByteArray) toDouble(args ...interface{}) () {
  // toDouble(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toDoubleEPb
  default:
    qtrt.ErrorResolve("QByteArray", "toDouble", args)
 }

}


func (this *QByteArray) toULongLong(args ...interface{}) () {
  // toULongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11toULongLongEPbi
  default:
    qtrt.ErrorResolve("QByteArray", "toULongLong", args)
 }

}


func (this *QByteArray) fromPercentEncoding_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromPercentEncoding", args)
 }

}


func (this *QByteArray) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray5clearEv
  default:
    qtrt.ErrorResolve("QByteArray", "clear", args)
 }

}


func (this *QByteArray) toLongLong(args ...interface{}) () {
  // toLongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10toLongLongEPbi
  default:
    qtrt.ErrorResolve("QByteArray", "toLongLong", args)
 }

}


func (this *QByteArray) constData(args ...interface{}) () {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray9constDataEv
  default:
    qtrt.ErrorResolve("QByteArray", "constData", args)
 }

}


func (this *QByteArray) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6lengthEv
  default:
    qtrt.ErrorResolve("QByteArray", "length", args)
 }

}


func (this *QByteArray) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray2atEi
  default:
    qtrt.ErrorResolve("QByteArray", "at", args)
 }

}


func (this *QByteArray) swap(args ...interface{}) () {
  // swap(class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4swapERS_
  default:
    qtrt.ErrorResolve("QByteArray", "swap", args)
 }

}


func (this *QByteArray) split(args ...interface{}) () {
  // split(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5splitEc
  default:
    qtrt.ErrorResolve("QByteArray", "split", args)
 }

}


func (this *QByteArray) right(args ...interface{}) () {
  // right(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5rightEi
  default:
    qtrt.ErrorResolve("QByteArray", "right", args)
 }

}


func (this *QByteArray) chop(args ...interface{}) () {
  // chop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4chopEi
  default:
    qtrt.ErrorResolve("QByteArray", "chop", args)
 }

}


func (this *QByteArray) toPercentEncoding(args ...interface{}) () {
  // toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][2] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray17toPercentEncodingERKS_S1_c
  default:
    qtrt.ErrorResolve("QByteArray", "toPercentEncoding", args)
 }

}


func (this *QByteArray) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10isDetachedEv
  default:
    qtrt.ErrorResolve("QByteArray", "isDetached", args)
 }

}


func (this *QByteArray) constEnd(args ...interface{}) () {
  // constEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8constEndEv
  default:
    qtrt.ErrorResolve("QByteArray", "constEnd", args)
 }

}


func (this *QByteArray) setRawData(args ...interface{}) () {
  // setRawData(const char *, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10setRawDataEPKcj
  default:
    qtrt.ErrorResolve("QByteArray", "setRawData", args)
 }

}


func (this *QByteArray) mid(args ...interface{}) () {
  // mid(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray3midEii
  default:
    qtrt.ErrorResolve("QByteArray", "mid", args)
 }

}


func (this *QByteArray) fill(args ...interface{}) () {
  // fill(char, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4fillEci
  default:
    qtrt.ErrorResolve("QByteArray", "fill", args)
 }

}


func (this *QByteArray) toUShort(args ...interface{}) () {
  // toUShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toUShortEPbi
  default:
    qtrt.ErrorResolve("QByteArray", "toUShort", args)
 }

}


func (this *QByteArray) rightJustified(args ...interface{}) () {
  // rightJustified(int, char, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray14rightJustifiedEicb
  default:
    qtrt.ErrorResolve("QByteArray", "rightJustified", args)
 }

}

// <= body block end

