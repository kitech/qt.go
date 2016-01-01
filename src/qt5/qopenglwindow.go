package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qopenglwindow.h
// dst-file: /src/gui/qopenglwindow.go
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
// class sizeof(QOpenGLWindow)=1
type QOpenGLWindow struct {
  /*qbase*/ QPaintDeviceWindow;
  qclsinst uint64 /* *mut c_void*/;
//  _frameSwapped QOpenGLWindow_frameSwapped_signal;
}


func (this *QOpenGLWindow) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow7isValidEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "isValid", args)
 }

}


func (this *QOpenGLWindow) grabFramebuffer(args ...interface{}) () {
  // grabFramebuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow15grabFramebufferEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "grabFramebuffer", args)
 }

}


func NewQOpenGLWindow(args ...interface{})() {
}


func (this *QOpenGLWindow) shareContext(args ...interface{}) () {
  // shareContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow12shareContextEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "shareContext", args)
 }

}


func (this *QOpenGLWindow) makeCurrent(args ...interface{}) () {
  // makeCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow11makeCurrentEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "makeCurrent", args)
 }

}


func (this *QOpenGLWindow) context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow7contextEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "context", args)
 }

}


func (this *QOpenGLWindow) doneCurrent(args ...interface{}) () {
  // doneCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow11doneCurrentEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "doneCurrent", args)
 }

}


func (this *QOpenGLWindow) defaultFramebufferObject(args ...interface{}) () {
  // defaultFramebufferObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow24defaultFramebufferObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "defaultFramebufferObject", args)
 }

}


func (this *QOpenGLWindow) FreeQOpenGLWindow(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "~QOpenGLWindow", args)
 }

}


func (this *QOpenGLWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "metaObject", args)
 }

}

// <= body block end

