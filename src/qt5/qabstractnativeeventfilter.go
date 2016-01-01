package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qabstractnativeeventfilter.h
// dst-file: /src/core/qabstractnativeeventfilter.go
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
// class sizeof(QAbstractNativeEventFilter)=16
type QAbstractNativeEventFilter struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QAbstractNativeEventFilter) nativeEventFilter(args ...interface{}) () {
  // nativeEventFilter(const class QByteArray &, void *, long *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  vtys[0][2] = qtrt.Int32Ty(true) // "long *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAbstractNativeEventFilter17nativeEventFilterERK10QByteArrayPvPl
  default:
    qtrt.ErrorResolve("QAbstractNativeEventFilter", "nativeEventFilter", args)
  }

}


func NewQAbstractNativeEventFilter(args ...interface{}) QAbstractNativeEventFilter {
  return QAbstractNativeEventFilter{}
}


func (this *QAbstractNativeEventFilter) FreeQAbstractNativeEventFilter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractNativeEventFilter", "~QAbstractNativeEventFilter", args)
  }

}

// <= body block end

