package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qframe.h
// dst-file: /src/widgets/qframe.go
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
// class sizeof(QFrame)=1
type QFrame struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QFrame) setFrameRect(args ...interface{}) () {
  // setFrameRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QFrame12setFrameRectERK5QRect
  default:
    qtrt.ErrorResolve("QFrame", "setFrameRect", args)
 }

}


func (this *QFrame) lineWidth(args ...interface{}) () {
  // lineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame9lineWidthEv
  default:
    qtrt.ErrorResolve("QFrame", "lineWidth", args)
 }

}


func (this *QFrame) setFrameStyle(args ...interface{}) () {
  // setFrameStyle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QFrame13setFrameStyleEi
  default:
    qtrt.ErrorResolve("QFrame", "setFrameStyle", args)
 }

}


func (this *QFrame) frameRect(args ...interface{}) () {
  // frameRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame9frameRectEv
  default:
    qtrt.ErrorResolve("QFrame", "frameRect", args)
 }

}


func (this *QFrame) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame8sizeHintEv
  default:
    qtrt.ErrorResolve("QFrame", "sizeHint", args)
 }

}


func NewQFrame(args ...interface{})() {
}


func (this *QFrame) frameStyle(args ...interface{}) () {
  // frameStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10frameStyleEv
  default:
    qtrt.ErrorResolve("QFrame", "frameStyle", args)
 }

}


func (this *QFrame) midLineWidth(args ...interface{}) () {
  // midLineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame12midLineWidthEv
  default:
    qtrt.ErrorResolve("QFrame", "midLineWidth", args)
 }

}


func (this *QFrame) setLineWidth(args ...interface{}) () {
  // setLineWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QFrame12setLineWidthEi
  default:
    qtrt.ErrorResolve("QFrame", "setLineWidth", args)
 }

}


func (this *QFrame) setMidLineWidth(args ...interface{}) () {
  // setMidLineWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN6QFrame15setMidLineWidthEi
  default:
    qtrt.ErrorResolve("QFrame", "setMidLineWidth", args)
 }

}


func (this *QFrame) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10metaObjectEv
  default:
    qtrt.ErrorResolve("QFrame", "metaObject", args)
 }

}


func (this *QFrame) frameWidth(args ...interface{}) () {
  // frameWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10frameWidthEv
  default:
    qtrt.ErrorResolve("QFrame", "frameWidth", args)
 }

}


func (this *QFrame) FreeQFrame(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QFrame", "~QFrame", args)
 }

}

// <= body block end

