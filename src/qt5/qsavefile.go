package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qsavefile.h
// dst-file: /src/core/qsavefile.go
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
// class sizeof(QSaveFile)=1
type QSaveFile struct {
  /*qbase*/ QFileDevice;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QSaveFile) cancelWriting(args ...interface{}) () {
  // cancelWriting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFile13cancelWritingEv
  default:
    qtrt.ErrorResolve("QSaveFile", "cancelWriting", args)
 }

}


func NewQSaveFile(args ...interface{})() {
}


func (this *QSaveFile) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QSaveFile8fileNameEv
  default:
    qtrt.ErrorResolve("QSaveFile", "fileName", args)
 }

}


func (this *QSaveFile) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QSaveFile10metaObjectEv
  default:
    qtrt.ErrorResolve("QSaveFile", "metaObject", args)
 }

}


func (this *QSaveFile) commit(args ...interface{}) () {
  // commit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFile6commitEv
  default:
    qtrt.ErrorResolve("QSaveFile", "commit", args)
 }

}


func (this *QSaveFile) FreeQSaveFile(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSaveFile", "~QSaveFile", args)
 }

}


func (this *QSaveFile) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFile11setFileNameERK7QString
  default:
    qtrt.ErrorResolve("QSaveFile", "setFileName", args)
 }

}


func (this *QSaveFile) directWriteFallback(args ...interface{}) () {
  // directWriteFallback()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QSaveFile19directWriteFallbackEv
  default:
    qtrt.ErrorResolve("QSaveFile", "directWriteFallback", args)
 }

}


func (this *QSaveFile) setDirectWriteFallback(args ...interface{}) () {
  // setDirectWriteFallback(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFile22setDirectWriteFallbackEb
  default:
    qtrt.ErrorResolve("QSaveFile", "setDirectWriteFallback", args)
 }

}

// <= body block end

